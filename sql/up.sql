-- Create the appointment table
CREATE TABLE IF NOT EXISTS appointment (
    id serial PRIMARY KEY,
    start_date timestamp with time zone UNIQUE NOT NULL,
    end_date timestamp with time zone UNIQUE NOT NULL,
    note varchar(65),
    created timestamp with time zone NOT NULL DEFAULT NOW()
)
