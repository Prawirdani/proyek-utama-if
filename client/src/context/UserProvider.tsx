import { Fetch } from '@/api/fetcher';
import { UserRegisterSchema, UserResetPasswordSchema, UserUpdateSchema } from '@/lib/schemas/user';
import { createContext, useContext, useEffect, useState } from 'react';

type Context = {
  // Fetch State
  loading: boolean;
  // Data State
  users: User[] | null;
  // Revalidate Data
  invalidate: () => Promise<void>;
  // add new meja
  registerUser: (data: UserRegisterSchema) => Promise<Response>;
  // update meja
  updateUser: (data: UserUpdateSchema) => Promise<Response>;
  // deactivate user
  deactivateUser: (id: number) => Promise<Response>;
  // activate user
  activateUser: (id: number) => Promise<Response>;
  // reset password
  resetPassword: (data: UserResetPasswordSchema) => Promise<Response>;
};

export const UsersContext = createContext<Context>({} as Context);
export const useUsers = () => useContext(UsersContext);

export default function UsersProvider({ children }: { children: React.ReactNode }) {
  const [users, setUsers] = useState<User[] | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    Fetch(fetchUsers)()
      .then((data) => setUsers(data))
      .finally(() => setLoading(false));
  }, []);

  const invalidate = async () => {
    await Fetch(fetchUsers)().then((data) => setUsers(data));
  };

  const fetchUsers = async () => {
    const response = await fetch('/api/v1/users', {
      credentials: 'include',
    });
    const resBody = (await response.json()) as ApiResponse<User[]>;
    return resBody.data;
  };

  const registerUser = async (data: UserRegisterSchema) => {
    return await fetch('/api/v1/auth/register', {
      method: 'POST',
      credentials: 'include',
      body: JSON.stringify({
        nama: data.nama,
        username: data.username,
        password: data.password,
      }),
    });
  };

  const updateUser = async (data: UserUpdateSchema) => {
    return await fetch(`/api/v1/users/${data.id}`, {
      method: 'PUT',
      credentials: 'include',
      body: JSON.stringify({
        nama: data.nama,
        username: data.username,
      }),
    });
  };

  const deactivateUser = async (id: number) => {
    return await fetch(`/api/v1/users/${id}/deactivate`, {
      method: 'DELETE',
      credentials: 'include',
    });
  };

  const activateUser = async (id: number) => {
    return await fetch(`/api/v1/users/${id}/activate`, {
      method: 'PUT',
      credentials: 'include',
    });
  };

  const resetPassword = async (data: UserResetPasswordSchema) => {
    return await fetch(`/api/v1/users/${data.id}/reset-password`, {
      method: 'PUT',
      credentials: 'include',
      body: JSON.stringify({
        password: data.newPassword,
      }),
    });
  };

  return (
    <UsersContext.Provider
      value={{
        loading,
        users,
        invalidate,
        registerUser,
        updateUser,
        deactivateUser,
        activateUser,
        resetPassword,
      }}
    >
      {children}
    </UsersContext.Provider>
  );
}
