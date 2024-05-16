import { Fetch } from '@/api/fetcher';
import { fetchPaymentMethods } from '@/api/payment_method';
import { createContext, useContext, useEffect, useState } from 'react';
import { z } from 'zod';

export const addSchema = z.object({
  tipePembayaran: z.string().min(1, { message: 'Mohon pilih tipe pembayaran' }),
  metode: z.string().min(1, { message: 'Mohon isi kolom nama metode' }),
  deskripsi: z.string().min(1, { message: 'Mohon isi kolom deskripsi' }),
});

export const updateSchema = z.object({
  ...addSchema.shape,
  id: z.number(),
});

export type AddSchema = z.infer<typeof addSchema>;
export type UpdateSchema = z.infer<typeof updateSchema>;

type PaymentMethodsCtxType = {
  // Fetch State
  loading: boolean;
  // Data State
  payment_methods: MetodePembayaran[] | null;
  // Tipe Pembayaran Form Options
  tipe_pembayaran_opts: TipePembayaran[];
  // Revalidate Data
  revalidate: () => Promise<void>;
  // add new metode pembayaran
  addMetodePembayaran: (data: AddSchema) => Promise<Response>;
  // update metode pembayaran
  updateMetodePembayaran: (data: UpdateSchema) => Promise<Response>;
  // delete metode pembayaran
  deleteMetodePembayaran: (id: number) => Promise<Response>;
};

export const PaymentMethodsContext = createContext<PaymentMethodsCtxType>({} as PaymentMethodsCtxType);
export const usePaymentMethods = () => useContext(PaymentMethodsContext);

export default function PaymentMethodsProvider({ children }: { children: React.ReactNode }) {
  const tipe_pembayaran_opts: TipePembayaran[] = ['CARD', 'MOBILE'];
  const [payment_methods, set_payment_methods] = useState<MetodePembayaran[] | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    Fetch(fetchPaymentMethods)()
      .then((data) => set_payment_methods(data))
      .finally(() => setLoading(false));
  }, []);

  const revalidate = async () => {
    await Fetch(fetchPaymentMethods)().then((data) => set_payment_methods(data));
  };

  const addMetodePembayaran = async (data: AddSchema) => {
    return await fetch('/api/v1/payments/methods', {
      method: 'POST',
      credentials: 'include',
      body: JSON.stringify({
        tipePembayaran: data.tipePembayaran,
        metode: data.metode,
        deskripsi: data.deskripsi,
      }),
    });
  };

  const updateMetodePembayaran = async (data: UpdateSchema) => {
    return await fetch(`/api/v1/payments/methods/${data.id}`, {
      method: 'PUT',
      credentials: 'include',
      body: JSON.stringify({
        tipePembayaran: data.tipePembayaran,
        metode: data.metode,
        deskripsi: data.deskripsi,
      }),
    });
  };

  const deleteMetodePembayaran = async (id: number) => {
    return await fetch(`/api/v1/payments/methods/${id}`, {
      method: 'DELETE',
      credentials: 'include',
    });
  };
  return (
    <PaymentMethodsContext.Provider
      value={{
        loading,
        payment_methods,
        tipe_pembayaran_opts,
        revalidate,
        addMetodePembayaran,
        updateMetodePembayaran,
        deleteMetodePembayaran,
      }}
    >
      {children}
    </PaymentMethodsContext.Provider>
  );
}
