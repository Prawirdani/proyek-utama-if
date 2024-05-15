import { Button } from '@/components/ui/button';
import { useAuth } from '@/context/useAuth';
import { ChevronDown, LogOut, Menu } from 'lucide-react';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import { useNavigate } from 'react-router-dom';

interface HeaderProps {
  sidebarOpen: boolean;
  setSidebarOpen: (open: boolean) => void;
}

export default function Header({ sidebarOpen, setSidebarOpen }: HeaderProps) {
  const navigate = useNavigate();
  const { user, logout } = useAuth();
  const toggleSidebar = () => {
    setSidebarOpen(!sidebarOpen);
  };

  const handleLogout = async () => {
    await logout().finally(() => {
      navigate('/login', { replace: true });
    });
  };

  return (
    <header className="border-b shadow-sm fixed top-0 left-0 right-0 bg-primary-foreground z-40 p-2 lg:p-4 flex justify-between lg:justify-end items-center">
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
      <div className="flex items-center gap-2">
        <div className="-space-y-1 [&>p]:text-sm">
          <p className="font-medium tracking-wide">{user.nama}</p>
          <p className="text-end text-sm text-muted-foreground">{user.role}</p>
        </div>
        <Dropdown handleLogout={handleLogout} />
      </div>
    </header>
  );
}

function Dropdown({ handleLogout }: { handleLogout: () => void }) {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger>
        <ChevronDown className="h-5 w-5" />
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" className="mt-3">
        <DropdownMenuItem className="space-x-2" onClick={handleLogout}>
          <LogOut />
          <span>Logout</span>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
