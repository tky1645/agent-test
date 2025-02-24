CREATE TABLE plant (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    wateringDate DATE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);