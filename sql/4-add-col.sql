

ALTER TABLE chair_locations
ADD COLUMN location POINT AS (POINT(longitude, latitude)) STORED NOT NULL,
ADD SPATIAL INDEX idx_location (location);

ALTER TABLE rides
ADD COLUMN pickup_location POINT AS (POINT(pickup_longitude, pickup_latitude)) STORED NOT NULL,
ADD SPATIAL INDEX idx_pickup_location (pickup_location);
