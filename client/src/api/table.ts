export const fetchTables = async () => {
  const res = await fetch('/api/v1/tables', {
    method: 'GET',
    credentials: 'include',
  });

  const resBody = (await res.json()) as { data: Meja[] | null };
  return resBody.data;
};
