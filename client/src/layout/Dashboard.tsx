import { useState } from 'react';
import Header from './Header';
import Sidebar from './Sidebar';
import { Navigate, Outlet } from 'react-router-dom';
import { useAuth } from '@/context/useAuth';

export default function Dashboard() {
	const [sidebarOpen, setSidebarOpen] = useState(false);
	const { isAuthenticated } = useAuth();

	return isAuthenticated ? (
		<div className="flex min-h-screen overflow-hidden">
			{/* Sidebar component */}
			<Sidebar sidebarOpen={sidebarOpen} setSidebarOpen={setSidebarOpen} />
			{/* Sidebar component */}

			<div className="flex-1 flex flex-col overflow-y-auto overflow-x-hidden pl-0 lg:pl-72 pt-12 lg:pt-18">
				{/* Header component */}
				<Header sidebarOpen={sidebarOpen} setSidebarOpen={setSidebarOpen} />
				{/* Header component */}
				{/* Main content */}
				<main className="bg-secondary h-full p-8 pt-10 lg:pt-14">
					<Outlet />
				</main>
				{/* Main content */}
			</div>
		</div>
	) : (
		<Navigate to="/login" replace />
	);
}
