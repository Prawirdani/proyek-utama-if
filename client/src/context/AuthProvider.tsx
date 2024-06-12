import { createContext, useContext, useState } from 'react';

type AuthContext = {
  user: AuthUser;
  isAuthenticated: boolean;
  login: (username: string, password: string) => Promise<Response>;
  identify: () => Promise<void>;
  logout: () => Promise<void>;
};

const AuthCtx = createContext<AuthContext | undefined>(undefined);
export const useAuth = () => {
  const ctx = useContext(AuthCtx);
  if (ctx === undefined) {
    throw new Error('Component is not wrapped with AuthProvider');
  }
  return ctx;
};

export default function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<AuthUser>({} as AuthUser);
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  const login = async (username: string, password: string) => {
    const res = await fetch('/api/v1/auth/login?web=true', {
      method: 'POST',
      credentials: 'include',
      body: JSON.stringify({
        username: username,
        password: password,
      }),
    });
    return res;
  };

  const logout = async () => {
    await fetch('/api/v1/auth/logout', {
      method: 'DELETE',
      credentials: 'include',
    });
    setIsAuthenticated(false);
    setUser({} as AuthUser);
  };

  const identify = async () => {
    const res = await fetch('/api/v1/auth/current', {
      method: 'GET',
      credentials: 'include',
    });
    if (res.ok) {
      const resBody = (await res.json()) as { data: AuthUser };
      setIsAuthenticated(true);
      setUser(resBody.data);
    }
  };

  return (
    <AuthCtx.Provider value={{ user, isAuthenticated, login, identify, logout }}>
      <>{children}</>
    </AuthCtx.Provider>
  );
}
