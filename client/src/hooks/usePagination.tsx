import { useEffect, useState } from 'react';

export type PaginationState = {
  page: number;
  limit: number;
  nextPage: () => void;
  prevPage: () => void;
  setLimit: (limit: number) => void;
};

export const usePagination = () => {
  const [query, setQuery] = useState<string>(`?page=1&limit=10`);
  const [pagination, setPagination] = useState<PaginationState>({
    page: 1,
    limit: 10,
    nextPage: () => {
      setPagination((prev) => ({ ...prev, page: prev.page + 1 }));
    },
    prevPage: () => {
      setPagination((prev) => ({ ...prev, page: prev.page - 1 }));
    },
    setLimit: (limit) => {
      setPagination((prev) => ({ ...prev, page: 1, limit }));
    },
  });

  useEffect(() => {
    setQuery(`?page=${pagination.page}&limit=${pagination.limit}`);
  }, [pagination]);

  return { query, pagination };
};
