-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TABLE groups (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    short_name VARCHAR(255) NOT NULL,
    description TEXT,
    color VARCHAR(128) NOT NULL,
    image_url TEXT,
    website_url TEXT,
    established_at DATE NOT NULL,
    dissolved_at DATE,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    image_url TEXT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE positions (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    group_id UUID NOT NULL REFERENCES groups(id) ON DELETE CASCADE,

    established_at DATE NOT NULL,
    dissolved_at DATE,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE trustees (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    start_date DATE NOT NULL,
    end_date DATE,
    position_id UUID NOT NULL REFERENCES positions(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TRIGGER update_timestamp_group
BEFORE UPDATE ON groups
FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_timestamp_user
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_timestamp_position
BEFORE UPDATE ON positions
FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_timestamp_trustee
BEFORE UPDATE ON trustees
FOR EACH ROW EXECUTE FUNCTION update_timestamp();

-- +goose Down
DROP TRIGGER IF EXISTS update_timestamp_trustee ON trustees;
DROP TRIGGER IF EXISTS update_timestamp_position ON positions;
DROP TRIGGER IF EXISTS update_timestamp_user ON users;
DROP TRIGGER IF EXISTS update_timestamp_group ON groups;
DROP TABLE IF EXISTS trustees;
DROP TABLE IF EXISTS positions;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS groups;
DROP FUNCTION IF EXISTS update_timestamp();
