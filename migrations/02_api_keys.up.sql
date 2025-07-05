CREATE TABLE api_key (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    value VARCHAR(32) NOT NULL
);
