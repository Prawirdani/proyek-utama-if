import { Button } from '@/components/ui/button';
import { X } from 'lucide-react';
import { useRef } from 'react';

interface SidebarProps {
  sidebarOpen: boolean;
  setSidebarOpen: (open: boolean) => void;
}

export default function Sidebar({ sidebarOpen, setSidebarOpen }: SidebarProps) {
  const sideBarRef = useRef(null);

  // className={`border lg:relative left-0 top-0 z-50 h-screen lg:w-72 duration-200 ease-linear shadow-xl ${sideBarOpen ? 'w-72' : 'w-0 overflow-hidden'}`}
  return (
    <aside
      ref={sideBarRef}
      id="sidebar"
      aria-labelledby="sidebar-toggle"
      className={`border fixed lg:relative top-0 z-50 h-screen w-72 duration-200 ease-linear shadow-xl transform bg-white ${
        sidebarOpen ? 'translate-x-0' : '-translate-x-full'
      } lg:translate-x-0`}
    >
      <div className="h-full flex flex-col p-2">
        <Button
          onClick={() => setSidebarOpen(!sidebarOpen)}
          variant="ghost"
          className="lg:hidden px-2 w-fit self-end"
        >
          <X />
        </Button>

        <h2 className="text-center">My Sidebar</h2>
      </div>
    </aside>
  );
}
