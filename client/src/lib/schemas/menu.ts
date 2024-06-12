import { z } from 'zod';

export const addMenuSchema = z.object({
  nama: z.string().min(1, { message: 'Nama menu harus di isi.' }).trim(),
  harga: z.number().min(1, { message: 'Tentukan harga menu' }),
  kategori_id: z.string().min(1, { message: 'Tentukan kategori menu' }),
  deskripsi: z.string().min(1, { message: 'Isi deskripsi menu' }),
  image: z.any().refine((files) => files?.length >= 1, { message: 'Photo is required.' }),
});

export type AddMenuSchema = z.infer<typeof addMenuSchema>;

export const updateMenuSchema = z.object({
  nama: z.string().min(1, { message: 'Nama menu harus di isi.' }).trim(),
  harga: z.number().min(1, { message: 'Tentukan harga menu' }),
  kategori_id: z.string().min(1, { message: 'Tentukan kategori menu' }),
  deskripsi: z.string().min(1, { message: 'Isi deskripsi menu' }),
  image: z.any().optional(),
});

export type UpdateMenuSchema = z.infer<typeof updateMenuSchema>;

// create and update menu category
export const menuCategorySchema = z.object({
  nama: z.string().min(1, { message: 'Nama kategori harus di isi.' }).trim(),
});

export type MenuCategorySchema = z.infer<typeof menuCategorySchema>;
