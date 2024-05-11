CREATE TYPE Role AS ENUM ('Kasir', 'Manajer');
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	nama VARCHAR(100) NOT NULL,
	username VARCHAR(100) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	aktif BOOL DEFAULT true,
	role Role NOT NULL default 'Kasir',
	created_at TIMESTAMPTZ DEFAULT current_timestamp,
	updated_at TIMESTAMPTZ DEFAULT current_timestamp
);
