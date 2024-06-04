import { RouterProvider, createBrowserRouter } from 'react-router-dom';
import './App.css';
import AuthProvider from './context/AuthProvider';
import { Toaster } from './components/ui/toaster';
import { routes } from './pages';

export default function App() {
  const router = createBrowserRouter(routes);
  return (
    <AuthProvider>
      <RouterProvider router={router} />
      <Toaster />
    </AuthProvider>
  );
}
