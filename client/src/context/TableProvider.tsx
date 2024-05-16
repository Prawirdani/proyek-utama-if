import { Fetch } from '@/api/fetcher';
import { fetchTables } from '@/api/table';
import { createContext, useContext, useEffect, useState } from 'react';
import { z } from 'zod';

export const addSchema = z.object({
  nomor: z.string().min(1, { message: 'Mohon isi kolom nomor meja' }),
});

export const updateSchema = z.object({
  ...addSchema.shape,
  id: z.number(),
});

export type AddSchema = z.infer<typeof addSchema>;
export type UpdateSchema = z.infer<typeof updateSchema>;

type Context = {
  // Fetch State
  loading: boolean;
  // Data State
  tables: Meja[] | null;
  // Revalidate Data
  revalidate: () => Promise<void>;
  // add new meja
  addMeja: (data: AddSchema) => Promise<Response>;
  // update meja
  updateMeja: (data: UpdateSchema) => Promise<Response>;
  // delete meja
  deleteMeja: (id: number) => Promise<Response>;
};

export const TableContext = createContext<Context>({} as Context);
export const useTables = () => useContext(TableContext);

export default function TablesProvider({ children }: { children: React.ReactNode }) {
  const [tables, setTables] = useState<Meja[] | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    Fetch(fetchTables)()
      .then((data) => setTables(data))
      .finally(() => setLoading(false));
  }, []);

  const revalidate = async () => {
    await Fetch(fetchTables)().then((data) => setTables(data));
  };

  const addMeja = async (data: AddSchema) => {
    return await fetch('/api/v1/tables', {
      method: 'POST',
      credentials: 'include',
      body: JSON.stringify({
        nomor: data.nomor,
      }),
    });
  };

  const updateMeja = async (data: UpdateSchema) => {
    return await fetch(`/api/v1/tables/${data.id}`, {
      method: 'PUT',
      credentials: 'include',
      body: JSON.stringify({
        nomor: data.nomor,
      }),
    });
  };

  const deleteMeja = async (id: number) => {
    return await fetch(`/api/v1/tables/${id}`, {
      method: 'DELETE',
      credentials: 'include',
    });
  };
  return (
    <TableContext.Provider
      value={{
        loading,
        tables,
        revalidate,
        addMeja,
        updateMeja,
        deleteMeja,
      }}
    >
      {children}
    </TableContext.Provider>
  );
}
