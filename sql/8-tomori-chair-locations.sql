-- chair_locationsのchair_idをユニークにする
ALTER TABLE chair_locations ADD CONSTRAINT unique_chair_id UNIQUE (chair_id);
