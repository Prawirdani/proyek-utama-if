import { z } from 'zod';

export const addTableSchema = z.object({
  nomor: z.string().min(1, { message: 'Mohon isi kolom nomor meja' }),
});

export const updateTableSchema = z.object({
  ...addTableSchema.shape,
  id: z.number(),
});

export type AddTableSchema = z.infer<typeof addTableSchema>;
export type UpdateTableSchema = z.infer<typeof updateTableSchema>;
