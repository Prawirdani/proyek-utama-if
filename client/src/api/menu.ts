export const fetchMenus = async () => {
  const res = await fetch('/api/v1/menus', {
    credentials: 'include',
  });
  const resBody = (await res.json()) as { data: Menu[] | null };
  return resBody.data;
};

export const fetchMenuCategories = async () => {
  const res = await fetch('/api/v1/menus/categories', {
    credentials: 'include',
  });
  const resBody = (await res.json()) as { data: Kategori[] };
  return resBody.data;
};
