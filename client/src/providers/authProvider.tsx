import { createContext, useState } from 'react';

type AuthCtxType = {
  user: AuthUser | undefined;
  setUser: (user: AuthUser | undefined) => void;
  login: (username: string, password: string) => Promise<void>;
  identify: () => Promise<void>;
  logout: () => Promise<void>;
};

type Props = {
  children: React.ReactNode;
};

export const AuthCtx = createContext<AuthCtxType>({} as AuthCtxType);

export default function AuthProvider({ children }: Props) {
  const [user, setUser] = useState<AuthUser | undefined>(undefined);

  const login = async (username: string, password: string) => {
    const url = '/api/v1/auth/login';
    const res = await fetch(url, {
      method: 'POST',
      credentials: 'include',
      body: JSON.stringify({
        username,
        password,
      }),
    });
    if (res.ok) {
      await identify();
    }
  };

  const logout = async () => {
    const url = '/api/v1/auth/logout';
    await fetch(url, {
      method: 'DELETE',
      credentials: 'include',
    });
    setUser(undefined);
  };

  const identify = async () => {
    if (!user) {
      const url = '/api/v1/auth/current';
      const res = await fetch(url, {
        method: 'GET',
        credentials: 'include',
      });
      if (res.ok) {
        const resBody = (await res.json()) as { data: { user: AuthUser } };
        return setUser(resBody.data.user);
      }
      setUser(undefined);
    }
  };

  return (
    <AuthCtx.Provider value={{ user, setUser, login, identify, logout }}>
      <>{children}</>
    </AuthCtx.Provider>
  );
}
