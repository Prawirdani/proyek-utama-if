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
