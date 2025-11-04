-- Add additional fields to users table for the Users module

ALTER TABLE users ADD COLUMN IF NOT EXISTS phone VARCHAR(20);
ALTER TABLE users ADD COLUMN IF NOT EXISTS avatar TEXT;
ALTER TABLE users ADD COLUMN IF NOT EXISTS department VARCHAR(100);

-- Create index on department
CREATE INDEX IF NOT EXISTS idx_users_department ON users(department);
