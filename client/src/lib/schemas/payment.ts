import { z } from 'zod';

export const addPaymentMethodSchema = z.object({
  tipePembayaran: z.string().min(1, { message: 'Mohon pilih tipe pembayaran' }),
  metode: z.string().min(1, { message: 'Mohon isi kolom nama metode' }),
  deskripsi: z.string().min(1, { message: 'Mohon isi kolom deskripsi' }),
});

export const updatePaymentMethodSchema = z.object({
  ...addPaymentMethodSchema.shape,
  id: z.number(),
});

export type AddPaymentMethodSchema = z.infer<typeof addPaymentMethodSchema>;
export type UpdatePaymentMethodSchema = z.infer<typeof updatePaymentMethodSchema>;
