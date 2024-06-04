import { Fetch } from '@/api/fetcher';
import { fetchTables } from '@/api/table';
import { AddTableSchema, UpdateTableSchema } from '@/lib/schemas/table';
import { createContext, useContext, useEffect, useState } from 'react';

type TableContext = {
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

const TableCtx = createContext<TableContext | undefined>(undefined);
export const useTable = () => {
  const ctx = useContext(TableCtx);
  if (ctx === undefined) {
    throw new Error('Component is not wrapped with TableProvider');
  }
  return ctx;
};

export default function TableProvider({ children }: { children: React.ReactNode }) {
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
    <TableCtx.Provider
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
    </TableCtx.Provider>
  );
}
