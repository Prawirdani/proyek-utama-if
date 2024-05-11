DROP TYPE IF EXISTS TipePesanan;
DROP TYPE IF EXISTS StatusPesanan;
CREATE TYPE TipePesanan AS ENUM ('Dine In', 'Take Away');
CREATE TYPE StatusPesanan AS ENUM ('Diproses', 'Dihidangkan');
CREATE TABLE pesanan (
	id SERIAL PRIMARY KEY,
	nama_pelanggan VARCHAR(100) NOT NULL,
	kasir_id INTEGER NOT NULL,
	meja_id INTEGER,
	total BIGINT NOT NULL,
	tipe_pesanan TipePesanan NOT NULL,
	status_pesanan StatusPesanan NOT NULL DEFAULT 'Diproses',
	catatan TEXT,
	waktu_pesanan TIMESTAMPTZ DEFAULT current_timestamp,
	CONSTRAINT fk_kasir_id FOREIGN KEY(kasir_id) REFERENCES users(id),
	CONSTRAINT fk_meja_id FOREIGN KEY(meja_id) REFERENCES meja(id)
);
