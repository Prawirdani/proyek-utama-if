DROP TYPE IF EXISTS StatusPesanan;
CREATE TYPE StatusPesanan AS ENUM ('Diproses', 'Dihidangkan');
CREATE TABLE IF NOT EXISTS pesanan(
	id SERIAL PRIMARY KEY,
	transaksi_id INTEGER NOT NULL,
	menu_id INTEGER NOT NULL,
	harga INTEGER NOT NULL,
	kuantitas INTEGER NOT NULL,
	subtotal BIGINT NOT NULL,	
	status_pesanan StatusPesanan NOT NULL DEFAULT 'Diproses',
	CONSTRAINT fk_transaksi_id FOREIGN KEY(transaksi_id) REFERENCES transaksi(id),
	CONSTRAINT fk_menu_id FOREIGN KEY(menu_id) REFERENCES menus(id)
);
