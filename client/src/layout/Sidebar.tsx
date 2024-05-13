import { useRef } from 'react';

interface SidebarProps {
  sideBarOpen: boolean;
}

export default function Sidebar({ sideBarOpen }: SidebarProps) {
  const sideBarRef = useRef(null);

  return (
    <aside
      ref={sideBarRef}
      id="sidebar"
      aria-labelledby="sidebar-toggle"
      className={`border relative left-0 top-0 z-9999 h-screen md:w-72 duration-200 ease-linear shadow-xl ${sideBarOpen ? 'w-72' : 'w-0 overflow-hidden'}`}
    >
      <h2 className="text-center">My Sidebar</h2>
    </aside>
  );
}
