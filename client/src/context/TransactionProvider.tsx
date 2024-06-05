import { createContext, useContext, useEffect, useState } from 'react';
import { debounce } from 'ts-debounce';
import { Fetch } from '@/api/fetcher';
import { PaginationState, usePagination } from '@/hooks/usePagination';

// TODO: Debounce should happen on nextPage, prevPage, and setLimit not the fetch itself
type TransactionContext = {
  // Fetch State
  loading: boolean;
  // Data State
  transactions: Transaksi[] | null;
  // Revalidate Data
  invalidate: () => Promise<void>;
  pageLoading: boolean;
  pagination: PaginationState;
  maxPage: number;
};

const TransactionCtx = createContext<TransactionContext | undefined>(undefined);
export const useTransaction = () => {
  const ctx = useContext(TransactionCtx);
  if (!ctx) throw new Error('useTransaction must be used within TransactionProvider');
  return ctx;
};

export default function TrasactionProvider({ children }: { children: React.ReactNode }) {
  const [pageLoading, setPageLoading] = useState<boolean>(false);
  const [maxPage, setMaxPage] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(true);
  const [transactions, setTransactions] = useState<Transaksi[]>([]);
  const [mounted, setMounted] = useState<boolean>(false);
  const { query, pagination } = usePagination();

  useEffect(() => {
    // Non-Debounced Fetch on First Load
    if (!mounted) {
      fetchTransactions().then(() => {
        setLoading(false);
        setMounted(true);
      });
    } else {
      console.log('debounced');
      setPageLoading(true);
      const debounced = debounce(Fetch(fetchTransactions), 100);
      debounced().then(() => setPageLoading(false));
      return () => {
        debounced.cancel();
      };
    }
  }, [query]);

  async function fetchTransactions() {
    const res = await fetch(`/api/v1/orders${query}&sort=datetime&order=desc`, {
      credentials: 'include',
    });
    const resBody = (await res.json()) as ApiResponse<{
      results: Transaksi[];
      pagination: Pagination;
    }>;
    setMaxPage(resBody.data?.pagination.maxPage ?? 0);
    setTransactions(resBody.data?.results ?? []);
  }

  async function invalidate() {
    await fetchTransactions();
  }

  return (
    <TransactionCtx.Provider value={{ loading, transactions, invalidate, pageLoading, pagination, maxPage }}>
      {children}
    </TransactionCtx.Provider>
  );
}
