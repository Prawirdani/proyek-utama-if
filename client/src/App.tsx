import { RouterProvider, createBrowserRouter } from 'react-router-dom';
import './App.css';
import Dashboard from './layout/Dashboard';
import LoginPage from './pages/LoginPage';

export default function App() {
  const router = createBrowserRouter([
    {
      path: '/',
      element: <Dashboard />,
    },
    {
      path: '/login',
      element: <LoginPage />,
    },
  ]);
  return <RouterProvider router={router} />;
}
