CREATE TYPE TipePembayaran AS ENUM ('CASH', 'CARD', 'MOBILE');
CREATE TABLE IF NOT EXISTS metode_pembayaran (
	id SERIAL PRIMARY KEY,
	tipe_pembayaran TipePembayaran NOT NULL,
	metode VARCHAR(100) UNIQUE,
	deskripsi TEXT,
	deleted_at TIMESTAMPTZ
);
