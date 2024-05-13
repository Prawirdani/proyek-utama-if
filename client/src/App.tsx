import { Outlet, RouterProvider, createBrowserRouter } from 'react-router-dom';
import './App.css';
import Dashboard from './layout/Dashboard';
import LoginPage from './pages/LoginPage';
import AuthProvider from './providers/authProvider';
import { useAuth } from './hooks/useAuth';
import { useEffect, useState } from 'react';
import Loader from '@/components/ui/loader';

export default function App() {
  const router = createBrowserRouter([
    {
      element: <PersistLogin />,
      children: [
        {
          path: '/',
          element: <Dashboard />,
        },
      ],
    },
    {
      path: '/login',
      element: <LoginPage />,
    },
  ]);
  return (
    <AuthProvider>
      <RouterProvider router={router} />
    </AuthProvider>
  );
}

const PersistLogin = () => {
  const [isLoading, setIsLoading] = useState(true);
  const { identify } = useAuth();

  useEffect(() => {
    const identifyUser = async () => {
      await identify().finally(() => setIsLoading(false));
    };

    identifyUser();
  }, []);

  return isLoading ? <Loader /> : <Outlet />;
};
