import { Button } from '@/components/ui/button';
import { LayoutDashboard, ReceiptText, TowerControl, User2, Utensils, Wallet2, X } from 'lucide-react';
import { useRef } from 'react';
import { Link, useLocation } from 'react-router-dom';

interface SidebarProps {
  sidebarOpen: boolean;
  setSidebarOpen: (open: boolean) => void;
}

const sidebarItems = [
  {
    name: 'Dashboard',
    icon: <LayoutDashboard />,
    path: '/',
  },
  {
    name: 'Menu',
    icon: <Utensils />,
    path: '/menus',
  },
  {
    name: 'Meja',
    icon: <TowerControl />,
    path: '/tables',
  },
  {
    name: 'Pembayaran',
    icon: <Wallet2 />,
    path: '/payments',
  },
  {
    name: 'Pengguna',
    icon: <User2 />,
    path: '/users',
  },
  {
    name: 'Transaksi',
    icon: <ReceiptText />,
    path: '/reports',
  },
];

export default function Sidebar({ sidebarOpen, setSidebarOpen }: SidebarProps) {
  const sideBarRef = useRef(null);

  return (
    <aside
      ref={sideBarRef}
      id="sidebar"
      aria-labelledby="sidebar-toggle"
      className={`border-r fixed top-0 left-0 z-50 h-screen w-72 duration-200 ease-linear shadow-xl transform bg-primary-foreground ${
        sidebarOpen ? 'translate-x-0' : '-translate-x-full'
      } lg:translate-x-0`}
    >
      <div className="h-full flex flex-col p-2 mb-8">
        <Button
          onClick={() => setSidebarOpen(!sidebarOpen)}
          variant="ghost"
          className="lg:hidden px-2 w-fit self-end mb-4"
        >
          <X />
        </Button>

        <div className="lg:mt-10 flex flex-col gap-y-3 p-2">
          {sidebarItems.map((item) => (
            <SidebarNavItem key={item.name} link={item.path} icon={item.icon} onClick={() => setSidebarOpen(false)}>
              {item.name}
            </SidebarNavItem>
          ))}
        </div>
      </div>
    </aside>
  );
}

interface SidebarNavItemProps extends React.HTMLAttributes<HTMLAnchorElement> {
  children: React.ReactNode;
  icon: React.ReactNode;
  link: string;
}

function SidebarNavItem({ children, icon, link, ...props }: SidebarNavItemProps) {
  const location = useLocation();
  const isActive = location.pathname === link;
  return (
    <Link
      to={link}
      {...props}
      className={`flex justify-start  w-full p-4 rounded-md text-left tracking-wide hover:bg-muted transition-colors duration-300 ${isActive && 'bg-muted text-primary'}`}
    >
      {icon}
      <p className="ml-2 font-medium">{children}</p>
    </Link>
  );
}
