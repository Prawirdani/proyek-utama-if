CREATE TABLE IF NOT EXISTS pembayaran (
	id SERIAL PRIMARY KEY,
	pesanan_id INTEGER NOT NULL,
	metode_pembayaran_id INTEGER NOT NULL,
	jumlah BIGINT NOT NULL,
	waktu_pembayaran TIMESTAMPTZ DEFAULT current_timestamp,
	CONSTRAINT fk_pesanan_id FOREIGN KEY(pesanan_id) REFERENCES pesanan(id),
	CONSTRAINT fk_metode_pembayaran_id FOREIGN KEY(metode_pembayaran_id) REFERENCES metode_pembayaran(id)
);
