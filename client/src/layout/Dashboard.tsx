import Header from './Header';
import Sidebar from './Sidebar';

export default function Dashboard() {
  return (
    <div className="flex h-screen overflow-hidden">
      <Sidebar />
      <div className="relative flex-1 flex flex-col overflow-y-auto overflow-x-hidden">
        <Header />
        <main className="bg-secondary h-full p-4">
          <p>Content Goes Here..</p>
        </main>
      </div>
    </div>
  );
}
