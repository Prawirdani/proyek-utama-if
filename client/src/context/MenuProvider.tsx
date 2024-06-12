import { MenuCategorySchema } from '@/lib/schemas/menu';
import { createContext, useContext, useEffect, useState } from 'react';

type MenuContext = {
  loading: boolean;
  menus: Menu[];
  categories: Kategori[];
  invalidate: () => Promise<void>;
  addMenu(data: FormData): Promise<Response>;
  updateMenu(id: number, data: FormData): Promise<Response>;
  deleteMenu(id: number): Promise<Response>;
  createMenuCategory(data: MenuCategorySchema): Promise<Response>;
  updateMenuCategory(id: number, data: MenuCategorySchema): Promise<Response>;
  deleteMenuCategory(id: number): Promise<Response>;
};

const MenuCtx = createContext<MenuContext | undefined>(undefined);
export const useMenu = () => {
  const ctx = useContext(MenuCtx);
  if (ctx === undefined) {
    throw new Error('Component is not wrapped with MenuProvider');
  }
  return ctx;
};

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

  async function addMenu(data: FormData) {
    return await fetch('/api/v1/menus', {
      method: 'POST',
      body: data,
      credentials: 'include',
    });
  }

  async function updateMenu(id: number, data: FormData) {
    return await fetch(`/api/v1/menus/${id}`, {
      method: 'PUT',
      body: data,
      credentials: 'include',
    });
  }

  async function deleteMenu(id: number) {
    return await fetch(`/api/v1/menus/${id}`, {
      method: 'DELETE',
      credentials: 'include',
    });
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

  async function createMenuCategory(data: MenuCategorySchema) {
    return await fetch('/api/v1/menus/categories', {
      method: 'POST',
      body: JSON.stringify(data),
      credentials: 'include',
    });
  }

  async function updateMenuCategory(id: Number, data: MenuCategorySchema) {
    return await fetch(`/api/v1/menus/categories/${id}`, {
      method: 'PUT',
      body: JSON.stringify(data),
      credentials: 'include',
    });
  }

  async function deleteMenuCategory(id: number) {
    return await fetch(`/api/v1/menus/categories/${id}`, {
      method: 'DELETE',
      credentials: 'include',
    });
  }

  async function invalidate() {
    await Promise.all([fetchMenus(), fetchMenuCategories()]);
  }

  return (
    <MenuCtx.Provider
      value={{
        loading,
        menus,
        categories,
        invalidate,
        addMenu,
        updateMenu,
        deleteMenu,
        createMenuCategory,
        updateMenuCategory,
        deleteMenuCategory,
      }}
    >
      <>{children}</>
    </MenuCtx.Provider>
  );
}
