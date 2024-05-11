CREATE TABLE IF NOT EXISTS menus (
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
