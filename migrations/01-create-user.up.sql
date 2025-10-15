-- +migrate Up

CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  external_id VARCHAR(255),                  
  email VARCHAR(255) UNIQUE,                 
  name VARCHAR(255) NOT NULL,
  password_hash VARCHAR(255),               
  role VARCHAR(50) NOT NULL DEFAULT 'user', 
  immutable BOOLEAN NOT NULL DEFAULT FALSE,  
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);


CREATE UNIQUE INDEX IF NOT EXISTS ux_users_externalid 
ON users(external_id) 
WHERE external_id IS NOT NULL;

CREATE UNIQUE INDEX IF NOT EXISTS ux_users_email 
ON users(email) 
WHERE email IS NOT NULL;