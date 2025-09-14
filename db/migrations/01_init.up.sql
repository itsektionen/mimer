CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE committee (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,

    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    short_name VARCHAR(255) NOT NULL,
    description TEXT,
    color VARCHAR(128) NOT NULL,
    image_url TEXT,
    website_url TEXT,
    active BOOLEAN NOT NULL DEFAULT false,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE person (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,

    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    image_url TEXT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE position (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,

    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT false,

    committee_id UUID NOT NULL REFERENCES committee(id),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE trustee (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,

    start_date DATE NOT NULL,
    end_date DATE NOT NULL,

    position_id UUID NOT NULL REFERENCES position(id),
    person_id UUID NOT NULL REFERENCES person(id),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TRIGGER update_timestamp_trigger
BEFORE UPDATE ON committee
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_timestamp_trigger
BEFORE UPDATE ON person
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_timestamp_trigger
BEFORE UPDATE ON position
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_timestamp_trigger
BEFORE UPDATE ON trustee
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
