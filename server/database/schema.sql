CREATE TYPE Role AS ENUM ('Kasir', 'Manajer');
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	nama VARCHAR(100) NOT NULL,
	username VARCHAR(100) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	aktif BOOL DEFAULT true,
	role Role NOT NULL default 'Kasir',
	created_at TIMESTAMPTZ DEFAULT current_timestamp,
	updated_at TIMESTAMPTZ DEFAULT current_timestamp
);

CREATE TABLE kategori_menu (
	id SERIAL PRIMARY KEY,
	nama VARCHAR(30) UNIQUE,
	deleted_at TIMESTAMPTZ
);

CREATE TABLE menus (
	id SERIAL PRIMARY KEY,
	nama VARCHAR(100) NOT NULL,
	deskripsi TEXT,
	harga BIGINT NOT NULL,
	kategori_id INTEGER NOT NULL,
	url_foto TEXT,
	deleted_at TIMESTAMPTZ,	
	created_at TIMESTAMPTZ DEFAULT current_timestamp,
	updated_at TIMESTAMPTZ DEFAULT current_timestamp,
	CONSTRAINT fk_kategori_menu_id FOREIGN KEY(kategori_id) REFERENCES kategori_menu(id)
);

CREATE TYPE StatusMeja AS ENUM ('Tersedia', 'Reserved', 'Terisi');
CREATE TABLE meja (
	id SERIAL PRIMARY KEY,
	nomor VARCHAR(30) UNIQUE NOT NULL,
	status StatusMeja NOT NULL DEFAULT 'Tersedia',
	deleted_at TIMESTAMPTZ
);

CREATE TYPE TipePembayaran AS ENUM ('CASH', 'CARD', 'MOBILE');
CREATE TABLE metode_pembayaran (
	id SERIAL PRIMARY KEY,
	tipe_pembayaran TipePembayaran NOT NULL,
	metode VARCHAR(100) UNIQUE,
	deskripsi TEXT,
	deleted_at TIMESTAMPTZ
);

CREATE TYPE TipePesanan AS ENUM ('Dine In', 'Take Away');
CREATE TYPE StatusPesanan AS ENUM ('Diproses', 'Disajikan', 'Selesai', 'Batal');
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

CREATE TABLE detail_pesanan(
	id SERIAL PRIMARY KEY,
	pesanan_id INTEGER NOT NULL,
	menu_id INTEGER NOT NULL,
	harga INTEGER NOT NULL,
	kuantitas INTEGER NOT NULL,
	subtotal BIGINT NOT NULL,	
	CONSTRAINT fk_pesanan_id FOREIGN KEY(pesanan_id) REFERENCES pesanan(id),
	CONSTRAINT fk_menu_id FOREIGN KEY(menu_id) REFERENCES menus(id)
);

CREATE TABLE pembayaran (
	id SERIAL PRIMARY KEY,
	pesanan_id INTEGER NOT NULL,
	metode_pembayaran_id INTEGER NOT NULL,
	jumlah BIGINT NOT NULL,
	waktu_pembayaran TIMESTAMPTZ DEFAULT current_timestamp,
	CONSTRAINT fk_pesanan_id FOREIGN KEY(pesanan_id) REFERENCES pesanan(id),
	CONSTRAINT fk_metode_pembayaran_id FOREIGN KEY(metode_pembayaran_id) REFERENCES metode_pembayaran(id)
);

