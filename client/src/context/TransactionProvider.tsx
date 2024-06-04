import { createContext, useContext, useEffect, useState } from 'react';
import { debounce } from 'ts-debounce';
import { Fetch } from '@/api/fetcher';

type TransactionContext = {
  // Fetch State
  loading: boolean;
  // Data State
  transactions: Transaksi[] | null;
  // Revalidate Data
  invalidate: () => Promise<void>;
  pageLoading: boolean;
  pagination: Pagination;
};

const TransactionCtx = createContext<TransactionContext | undefined>(undefined);
export const useTransaction = () => {
  const ctx = useContext(TransactionCtx);
  if (!ctx) throw new Error('useTransaction must be used within TransactionProvider');
  return ctx;
};

type Pagination = {
  page: number;
  limit: number;
  nextPage: () => void;
  prevPage: () => void;
  setLimit: (limit: number) => void;
};

export default function TrasactionProvider({ children }: { children: React.ReactNode }) {
  const [pageLoading, setPageLoading] = useState<boolean>(false);
  const [pagination, setPagination] = useState<Pagination>({
    page: 1,
    limit: 10,
    nextPage: () => {
      setPagination((prev) => ({ ...prev, page: prev.page + 1 }));
    },
    prevPage: () => {
      setPagination((prev) => ({ ...prev, page: prev.page - 1 }));
    },
    setLimit: (limit) => {
      setPagination((prev) => ({ ...prev, limit }));
    },
  });
  const [query, setQuery] = useState<string>(`?page=1&limit=10`);
  const [loading, setLoading] = useState<boolean>(true);
  const [transactions, setTransactions] = useState<Transaksi[]>([]);

  useEffect(() => {
    console.log('Pagination Changed');
    setQuery(`?page=${pagination.page}&limit=${pagination.limit}`);
  }, [pagination]);

  useEffect(() => {
    console.log('Query Changed');
    setPageLoading(true);
    const debounced = debounce(async () => {
      await Fetch(fetchTransactions)()
        .then(() => setLoading(false))
        .catch((err) => console.error(err));
    }, 300);
    debounced().then(() => setPageLoading(false));
  }, [query]);

  async function fetchTransactions() {
    const res = await fetch(`/api/v1/orders${query}&sort=datetime&order=desc`, {
      credentials: 'include',
    });
    const resBody = (await res.json()) as ApiResponse<Transaksi[]>;
    setTransactions(resBody.data ?? []);
  }

  async function invalidate() {
    await fetchTransactions();
  }

  return (
    <TransactionCtx.Provider value={{ loading, transactions, invalidate, pageLoading, pagination }}>
      {children}
    </TransactionCtx.Provider>
  );
}
