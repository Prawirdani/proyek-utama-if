import { Fetch } from '@/api/fetcher';
import { fetchTables } from '@/api/table';
import { AddTableSchema, UpdateTableSchema } from '@/lib/schemas/table';
import { createContext, useContext, useEffect, useState } from 'react';

type Context = {
  // Fetch State
  loading: boolean;
  // Data State
  tables: Meja[] | null;
  // Revalidate Data
  invalidate: () => Promise<void>;
  // add new meja
  addMeja: (data: AddTableSchema) => Promise<Response>;
  // update meja
  updateMeja: (data: UpdateTableSchema) => Promise<Response>;
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

  const invalidate = async () => {
    await Fetch(fetchTables)().then((data) => setTables(data));
  };

  const addMeja = async (data: AddTableSchema) => {
    return await fetch('/api/v1/tables', {
      method: 'POST',
      credentials: 'include',
      body: JSON.stringify({
        nomor: data.nomor,
      }),
    });
  };

  const updateMeja = async (data: UpdateTableSchema) => {
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
        invalidate,
        addMeja,
        updateMeja,
        deleteMeja,
      }}
    >
      {children}
    </TableContext.Provider>
  );
}
