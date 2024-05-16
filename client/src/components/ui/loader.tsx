import { Loader2 } from 'lucide-react';

export default function Loader() {
  return (
    <div className="h-full flex place-items-center">
      <Loader2 className="mx-auto w-12 md:w-20 h-12 md:h-20 duration-1000 animate-spin text-primary" />
    </div>
  );
}
