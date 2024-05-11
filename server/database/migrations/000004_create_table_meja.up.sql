CREATE TYPE StatusMeja AS ENUM ('Tersedia', 'Reserved', 'Terisi');
CREATE TABLE IF NOT EXISTS meja (
	id SERIAL PRIMARY KEY,
	nomor VARCHAR(30) UNIQUE NOT NULL,
	status StatusMeja NOT NULL DEFAULT 'Tersedia',
	deleted_at TIMESTAMPTZ
);
