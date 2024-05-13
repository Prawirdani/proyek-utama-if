import { Loader2 } from 'lucide-react';

export default function Loader() {
  return (
    <div className="h-screen flex place-items-center">
      <Loader2 className="mx-auto w-24 h-24 duration-1000 animate-spin text-primary" />
    </div>
  );
}
