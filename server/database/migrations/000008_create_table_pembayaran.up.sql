CREATE TABLE IF NOT EXISTS pembayaran (
	id SERIAL PRIMARY KEY,
	transaksi_id INTEGER NOT NULL,
	metode_pembayaran_id INTEGER NOT NULL,
	jumlah BIGINT NOT NULL,
	waktu_pembayaran TIMESTAMPTZ DEFAULT current_timestamp,
	CONSTRAINT fk_transaksi_id FOREIGN KEY(transaksi_id) REFERENCES transaksi(id),
	CONSTRAINT fk_metode_pembayaran_id FOREIGN KEY(metode_pembayaran_id) REFERENCES metode_pembayaran(id)
);
