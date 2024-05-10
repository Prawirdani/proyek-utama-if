CREATE TABLE kategori_menu (
	id SERIAL PRIMARY KEY,
	nama VARCHAR(30) UNIQUE,
	deleted_at TIMESTAMPTZ
);
