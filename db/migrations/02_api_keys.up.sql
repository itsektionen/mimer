CREATE TABLE api_key (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    value VARCHAR(32) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT false,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TRIGGER update_timestamp_trigger
BEFORE UPDATE ON api_key
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
