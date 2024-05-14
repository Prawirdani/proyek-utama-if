import { Outlet, RouterProvider, createBrowserRouter } from 'react-router-dom';
import './App.css';
import Dashboard from './layout/Dashboard';
import LoginPage from './pages/LoginPage';
import AuthProvider from './context/authProvider';
import { useEffect, useState } from 'react';
import Loader from '@/components/ui/loader';
import Index from './pages/dashboard/Index';
import { useAuth } from './context/useAuth';
import MenuPage from './pages/dashboard/MenuPage';
import TablePage from './pages/dashboard/TablePage';
import PaymentPage from './pages/dashboard/PaymentPage';
import ReportPage from './pages/dashboard/ReportPage';
import UserPage from './pages/dashboard/UserPage';

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
              element: <Index />,
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
