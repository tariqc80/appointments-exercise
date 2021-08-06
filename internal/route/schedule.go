package route

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/tariqc80/appointments-exercise/internal/data"
)

// ParseTimeslot middleware to set start and end times from request.
func ParseTimeslot(c *gin.Context) {
	log.Print("Parsing timeslot request parameters")

	start := c.PostForm("start")

	log.Print("Request param 'start': ", start)

	startDate, err := time.Parse(time.RFC3339, start)

	if err != nil {
		message := fmt.Sprintf("Error parsing 'start' of value '%s'", start)
		c.AbortWithStatusJSON(400, gin.H{"message": message})
		return
	}

	durationParam := c.PostForm("duration")
	duration, err := strconv.Atoi(durationParam)

	if err != nil {
		message := fmt.Sprintf("Error parsing 'duration' of value '%s'", durationParam)
		c.AbortWithStatusJSON(400, gin.H{"message": message})
		return
	}

	log.Print("Request param 'duration': ", duration)

	endDate := startDate.Add(time.Minute * time.Duration(duration))

	c.Set("start", startDate)
	c.Set("end", endDate)

	log.Print("Completed parsing, the parsed input is:")
	log.Print("start: ", startDate)
	log.Print("end: ", endDate)

	c.Next()
}

// Schedule route handler struct
type Schedule struct {
	data *data.Schedule
}

// NewSchedule creates and returns a route handler
func NewSchedule(d *data.Schedule) *Schedule {
	return &Schedule{
		data: d,
	}
}

// ConfigRoutes sets up all the routes
func (s *Schedule) ConfigRoutes(e *gin.Engine) {
	// create route group for /schedule
	g := e.Group("/schedule")

	// Add parameter validation and parsing middleware
	g.Use(ParseTimeslot)

	// Add all the routes
	g.POST("/available", s.IsAvailable)
	g.POST("/create", s.Create)
	g.POST("/cancel", s.Cancel)
}

// Create adds a new appointment
func (s *Schedule) Create(c *gin.Context) {
	startDate := c.MustGet("start").(time.Time)
	endDate := c.MustGet("end").(time.Time)

	// check new appointment times are available
	isAvailable, err := s.data.IsAvailable(startDate, endDate)

	if err != nil {
		// if we encounter an error end the request
		c.AbortWithError(500, err)
		return
	}

	if isAvailable == false {
		log.Print("Could not create new appointment; time slot unavailable")
		c.AbortWithStatusJSON(400, gin.H{"message": "Time slot is not available for appointment"})
	} else {
		log.Print("Creating a new appointment")
		log.Print(startDate, " - ", endDate)

		err = s.data.Insert(startDate, endDate)

		if err != nil {
			c.AbortWithError(500, err)
		} else {
			log.Print("Successfully created an appointment")
			c.JSON(200, gin.H{"message": "New appointment created"})
		}
	}
}

// Cancel removes an existing appointment
func (s *Schedule) Cancel(c *gin.Context) {
	startDate := c.MustGet("start").(time.Time)
	endDate := c.MustGet("end").(time.Time)

	log.Print("Cancelling an appointment")
	log.Print(startDate, " - ", endDate)

	deleted, err := s.data.Delete(startDate, endDate)

	if err != nil {
		c.AbortWithError(500, err)
	} else {
		if deleted == false {
			log.Print("Could not cancel appointment; time slot not found")
			c.JSON(400, gin.H{"message": "Provided timeslot not found"})
		} else {
			log.Print("Successfully cancelled appointment")
			c.JSON(200, gin.H{"message": "Appointment cancelled"})
		}
	}
}

// IsAvailable checks if timeslots are available for booking
func (s *Schedule) IsAvailable(c *gin.Context) {
	startDate := c.MustGet("start").(time.Time)
	endDate := c.MustGet("end").(time.Time)

	log.Print("Checking if times are available")
	log.Print(startDate, " - ", endDate)

	available, err := s.data.IsAvailable(startDate, endDate)

	if err != nil {
		c.AbortWithError(500, err)
	} else {
		var message string

		if available {
			message = "Time slot is available"
		} else {
			message = "Time slot is unavailable"
		}

		c.JSON(200, gin.H{
			"available": available,
			"message":   message,
		})
	}
}
