CREATE TABLE IF NOT EXISTS bookings(
    id SERIAL PRIMARY KEY,
    first_name VARCHAR,
    last_name VARCHAR,
    gender VARCHAR,
    birthday DATE,
    launchpad_id VARCHAR,
    destination_id VARCHAR,
    launch_date DATE
);
