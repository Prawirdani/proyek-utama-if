import { useState } from 'react';
import Header from './Header';
import Sidebar from './Sidebar';
import { Outlet } from 'react-router-dom';

export default function Dashboard() {
  const [sidebarOpen, setSidebarOpen] = useState(false);

  return (
    <div className="flex h-screen overflow-hidden">
      {/* Sidebar component */}
      <Sidebar sideBarOpen={sidebarOpen} />
      {/* Sidebar component */}

      <div className="relative flex-1 flex flex-col overflow-y-auto overflow-x-hidden">
        {/* Header component */}
        <Header sidebarOpen={sidebarOpen} setSidebarOpen={setSidebarOpen} />
        {/* Header component */}
        {/* Main content */}
        <main className="bg-secondary h-full p-4">
          <p>Content Goes Here..</p>
          <Outlet />
        </main>
        {/* Main content */}
      </div>
    </div>
  );
}
