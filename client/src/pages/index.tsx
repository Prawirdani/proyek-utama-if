import IndexPage from './IndexPage';
import MenuPage from './MenuPage';
import UserPage from './UserPage';
import TablePage from './TablePage';
import PaymentPage from './PaymentPage';
import { Outlet, RouteObject } from 'react-router-dom';
import { useEffect, useState } from 'react';
import { useAuth } from '@/context/AuthProvider';
import LoginPage from './LoginPage';
import Loader from '@/components/ui/loader';
import Dashboard from '@/components/layout/Dashboard';
import TransactionPage from './TransactionPage';

export const routes: RouteObject[] = [
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
            path: '/transactions',
            element: <TransactionPage />,
          },
        ],
      },
    ],
  },
  {
    path: '/login',
    element: <LoginPage />,
  },
];

function PersistLogin() {
  const [isLoading, setIsLoading] = useState(true);
  const { identify } = useAuth();

  useEffect(() => {
    const identifyUser = async () => {
      await identify().finally(() => setIsLoading(false));
    };

    identifyUser();
  }, []);

  return isLoading ? (
    <div className="h-screen">
      <Loader />
    </div>
  ) : (
    <Outlet />
  );
}
