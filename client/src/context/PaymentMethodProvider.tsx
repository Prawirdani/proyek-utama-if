import { Fetch } from '@/api/fetcher';
import { fetchPaymentMethods } from '@/api/payment_method';
import { AddPaymentMethodSchema, UpdatePaymentMethodSchema } from '@/lib/schemas/payment';
import { createContext, useContext, useEffect, useState } from 'react';

type PaymentMethodContext = {
  // Fetch State
  loading: boolean;
  // Data State
  payment_methods: MetodePembayaran[] | null;
  // Tipe Pembayaran Form Options
  tipe_pembayaran_opts: TipePembayaran[];
  // Revalidate Data
  invalidate: () => Promise<void>;
  // add new metode pembayaran
  addMetodePembayaran: (data: AddPaymentMethodSchema) => Promise<Response>;
  // update metode pembayaran
  updateMetodePembayaran: (data: UpdatePaymentMethodSchema) => Promise<Response>;
  // delete metode pembayaran
  deleteMetodePembayaran: (id: number) => Promise<Response>;
};

const PaymentMethodCtx = createContext<PaymentMethodContext | undefined>(undefined);
export const usePaymentMethod = () => {
  const ctx = useContext(PaymentMethodCtx);
  if (ctx === undefined) {
    throw new Error('Component is not wrapped with PaymentMethodsProvider');
  }
  return ctx;
};

export default function PaymentMethodProvider({ children }: { children: React.ReactNode }) {
  const tipe_pembayaran_opts: TipePembayaran[] = ['CARD', 'MOBILE'];
  const [payment_methods, set_payment_methods] = useState<MetodePembayaran[] | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    Fetch(fetchPaymentMethods)()
      .then((data) => set_payment_methods(data))
      .finally(() => setLoading(false));
  }, []);

  const invalidate = async () => {
    await Fetch(fetchPaymentMethods)().then((data) => set_payment_methods(data));
  };

  const addMetodePembayaran = async (data: AddPaymentMethodSchema) => {
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

  const updateMetodePembayaran = async (data: UpdatePaymentMethodSchema) => {
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
    <PaymentMethodCtx.Provider
      value={{
        loading,
        payment_methods,
        tipe_pembayaran_opts,
        invalidate,
        addMetodePembayaran,
        updateMetodePembayaran,
        deleteMetodePembayaran,
      }}
    >
      {children}
    </PaymentMethodCtx.Provider>
  );
}
