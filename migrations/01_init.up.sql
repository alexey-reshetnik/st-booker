CREATE TABLE IF NOT EXISTS bookings (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name VARCHAR,
    last_name VARCHAR,
    gender VARCHAR,
    birthday DATE,
    launchpad_id VARCHAR, // TODO?
    destination_id VARCHAR,
    launch_date DATE,
);

