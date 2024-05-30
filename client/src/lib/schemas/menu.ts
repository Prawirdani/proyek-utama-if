import { z } from 'zod';

export const addMenuSchema = z.object({
  nama: z.string().min(1, { message: 'Nama menu harus di isi.' }).trim(),
  harga: z.number().min(1, { message: 'Tentukan harga menu' }),
  kategori_id: z.string().min(1, { message: 'Tentukan kategori menu' }),
  deskripsi: z.string().min(1, { message: 'Isi deskripsi menu' }),
  image: z.any().refine((files) => files?.length >= 1, { message: 'Photo is required.' }),
});

export type AddMenuSchema = z.infer<typeof addMenuSchema>;
