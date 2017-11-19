CREATE DATABASE rs_db;

USE rs_db;

CREATE TABLE IF NOT EXISTS properties(
    id INT NOT NULL AUTO_INCREMENT, 
    description TEXT, 
    type VARCHAR(256) NOT NULL, 
    orientation VARCHAR(256),
    courtyard_size INT,
    bedrooms INT NOT NULL,
    living_room_size VARCHAR(64),
    kitchen_size VARCHAR(64),
    bedrooms_sizes VARCHAR(256),
    bathrooms INT NOT NULL,
    showers INT NOT NULL,
    size INT NOT NULL,
    construction_year INT,
    padron VARCHAR(256),
    building_name VARCHAR(256),
    apartments_per_floor INT,
    floors INT NOT NULL,
    terrace_size VARCHAR(64),
    balcony_size VARCHAR(64),
    expenses BIGINT NOT NULL,
    amenities TEXT,
    created_date DATETIME,
    address_street VARCHAR(2048) NOT NULL,
    address_street_number VARCHAR(256),
    address_apartment_number VARCHAR(256),
    address_neighborhood VARCHAR(2048) NOT NULL,
    address_city VARCHAR(512) NOT NULL,
    address_country VARCHAR(512) NOT NULL,
    address_postal_code VARCHAR(512) NOT NULL,
    address_latitude VARCHAR(512) NOT NULL,
    address_longitude VARCHAR(512) NOT NULL,
    elevators INT,
    PRIMARY KEY(id)
    -- CONSTRAINT fk_address FOREIGN KEY (address_id) REFERENCES addresses(id)
); 

CREATE TABLE IF NOT EXISTS properties_state(
    id INT NOT NULL AUTO_INCREMENT,
    property_id INT NOT NULL,
    overall INT,
    safety INT,
    windows INT,
    floors INT,
    kitchen INT,
    kitchen_furniture INT,
    natural_light INT,
    appearence INT,
    garage INT,
    water_preasure INT,
    created_date DATETIME,
    PRIMARY KEY(id),
    CONSTRAINT fk_property FOREIGN KEY (property_id) REFERENCES properties(id)
);