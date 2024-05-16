export const fetchTables = async () => {
  const res = await fetch('/api/v1/tables', {
    method: 'GET',
    credentials: 'include',
  });

  if (!res.ok) {
    const errorBody = (await res.json()) as ErrorResponse;
    throw new Error(errorBody.error.message);
  }

  const resBody = (await res.json()) as { data: Meja[] | null };
  return resBody.data;
};
