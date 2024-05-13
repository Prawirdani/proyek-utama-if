import { Button } from '@/components/ui/button';
import { Menu } from 'lucide-react';

interface HeaderProps {
  sidebarOpen: boolean;
  setSidebarOpen: (open: boolean) => void;
}

export default function Header({ sidebarOpen, setSidebarOpen }: HeaderProps) {
  const toggleSidebar = () => {
    setSidebarOpen(!sidebarOpen);
  };
  return (
    <header className="border-b shadow-sm top-0 sticky z-999 p-4 flex justify-between md:justify-end  items-center">
      <Button
        id="sidebar-toggle"
        aria-controls="sidebar"
        aria-expanded={sidebarOpen}
        onClick={toggleSidebar}
        variant="ghost"
        className="md:hidden p-2 h-auto w-auto rounded-md"
      >
        <Menu />
      </Button>
      <div>
        <h1 className="text-center">My Header</h1>
      </div>
    </header>
  );
}
