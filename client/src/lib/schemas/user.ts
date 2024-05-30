import { z } from 'zod';

export const userRegisterSchema = z
  .object({
    nama: z.string().min(1, { message: 'Isi nama lengkap pengguna.' }),
    username: z.string().min(1, { message: 'Isi username pengguna.' }),
    password: z.string().min(1, { message: 'Isi password pengguna.' }),
    repeatPassword: z.string().min(1, { message: 'Isi ulang password pengguna.' }),
  })
  .refine((data) => data.password === data.repeatPassword, {
    message: 'Password tidak sama.',
    path: ['repeatPassword'],
  });

export const userUpdateSchema = z.object({
  id: z.number(),
  nama: z.string().min(1, { message: 'Isi nama lengkap pengguna.' }),
  username: z.string().min(1, { message: 'Isi username pengguna.' }),
});

export const userResetPasswordSchema = z
  .object({
    id: z.number(),
    newPassword: z.string().min(1, { message: 'Isi kolom password.' }),
    repeatPassword: z.string().min(1, { message: 'Isi kolom ulangi password.' }),
  })
  .refine((data) => data.newPassword === data.repeatPassword, {
    message: 'Password tidak sama.',
    path: ['repeatPassword'],
  });

export type UserRegisterSchema = z.infer<typeof userRegisterSchema>;
export type UserUpdateSchema = z.infer<typeof userUpdateSchema>;
export type UserResetPasswordSchema = z.infer<typeof userResetPasswordSchema>;
