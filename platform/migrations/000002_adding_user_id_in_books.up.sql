-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create books table
ALTER TABLE books
ADD COLUMN user_id UUID NOT NULL;
