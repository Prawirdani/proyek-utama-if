import { createContext, useContext, useEffect, useState } from 'react';

type MenuContext = {
  loading: boolean;
  menus: Menu[];
  categories: Kategori[];
  invalidate: () => Promise<void>;
};

const MenuCtx = createContext<MenuContext>({} as MenuContext);
export const useMenu = () => useContext(MenuCtx);

export default function MenuProvider({ children }: { children: React.ReactNode }) {
  const [loading, setLoading] = useState(true);
  const [menus, setMenus] = useState<Menu[]>([]);
  const [categories, setCategories] = useState<Kategori[]>([]);

  useEffect(() => {
    (async () => {
      await Promise.all([fetchMenus(), fetchMenuCategories()]);
    })()
      .then(() => setLoading(false))
      .catch((err) => {
        console.error(err);
      });
  }, []);

  async function fetchMenus() {
    const res = await fetch('/api/v1/menus', {
      credentials: 'include',
    });
    if (!res.ok) {
      const errorBody = (await res.json()) as ErrorResponse;
      throw new Error(errorBody.error.message);
    }
    const resBody = (await res.json()) as ApiResponse<Menu[]>;
    setMenus(resBody.data!);
  }

  async function fetchMenuCategories() {
    const res = await fetch('/api/v1/menus/categories', {
      credentials: 'include',
    });
    if (!res.ok) {
      const errorBody = (await res.json()) as ErrorResponse;
      throw new Error(errorBody.error.message);
    }
    const resBody = (await res.json()) as ApiResponse<Kategori[]>;
    setCategories(resBody.data!);
  }

  async function invalidate() {
    setLoading(true);
    await Promise.all([fetchMenus(), fetchMenuCategories()]);
    setLoading(false);
  }

  return (
    <MenuCtx.Provider value={{ loading, menus, categories, invalidate }}>
      <>{children}</>
    </MenuCtx.Provider>
  );
}
