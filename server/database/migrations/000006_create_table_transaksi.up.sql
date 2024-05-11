DROP TYPE IF EXISTS TipeTransaksi;
CREATE TYPE TipeTransaksi AS ENUM ('Dine In', 'Take Away');
CREATE TABLE IF NOT EXISTS transaksi (
	id SERIAL PRIMARY KEY,
	nama_pelanggan VARCHAR(100) NOT NULL,
	kasir_id INTEGER NOT NULL,
	meja_id INTEGER,
	total BIGINT NOT NULL,
	tipe_transaksi TipeTransaksi NOT NULL,
	catatan TEXT,
	waktu_transaksi TIMESTAMPTZ DEFAULT current_timestamp,
	CONSTRAINT fk_kasir_id FOREIGN KEY(kasir_id) REFERENCES users(id),
	CONSTRAINT fk_meja_id FOREIGN KEY(meja_id) REFERENCES meja(id)
);
