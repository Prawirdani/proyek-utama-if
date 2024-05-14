import { Button } from '@/components/ui/button';
import { useAuth } from '@/context/useAuth';
import { Menu } from 'lucide-react';

interface HeaderProps {
  sidebarOpen: boolean;
  setSidebarOpen: (open: boolean) => void;
}

export default function Header({ sidebarOpen, setSidebarOpen }: HeaderProps) {
  const { user } = useAuth();
  const toggleSidebar = () => {
    setSidebarOpen(!sidebarOpen);
  };

  return (
    <header className="border-b shadow-sm top-0 sticky bg-white z-40 p-2 lg:p-4 flex justify-between lg:justify-end  items-center">
      <Button
        id="sidebar-toggle"
        aria-controls="sidebar"
        aria-expanded={sidebarOpen}
        onClick={toggleSidebar}
        variant="ghost"
        className="lg:hidden p-2 h-auto w-auto rounded-md"
      >
        <Menu />
      </Button>
      <div>
        <p>{user.nama}</p>
      </div>
    </header>
  );
}
