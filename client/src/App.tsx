import { Outlet, RouterProvider, createBrowserRouter } from 'react-router-dom';
import './App.css';
import Dashboard from './layout/Dashboard';
import LoginPage from './pages/LoginPage';
import AuthProvider from './context/authProvider';
import { useEffect, useState } from 'react';
import Loader from '@/components/ui/loader';
import { useAuth } from './context/useAuth';
import { IndexPage, MenuPage, TablePage, PaymentPage, UserPage, ReportPage } from './pages/dashboard';
import { Toaster } from './components/ui/toaster';

export default function App() {
  const router = createBrowserRouter([
    {
      element: <PersistLogin />,
      children: [
        {
          path: '/',
          element: <Dashboard />,
          children: [
            {
              path: '/',
              element: <IndexPage />,
            },
            {
              path: '/menus',
              element: <MenuPage />,
            },
            {
              path: '/tables',
              element: <TablePage />,
            },
            {
              path: '/payments',
              element: <PaymentPage />,
            },
            {
              path: '/users',
              element: <UserPage />,
            },
            {
              path: '/reports',
              element: <ReportPage />,
            },
          ],
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
      <Toaster />
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
