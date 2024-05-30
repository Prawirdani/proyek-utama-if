import { Card } from '@/components/ui/card';
import { SquarePen } from 'lucide-react';

interface Props {
  menu: Menu;
}
export default function MenuCard({ menu }: Props) {
  return (
    <Card className="shadow-lg hover:cursor-pointer relative">
      <div className="space-y-1">
        {/* Image */}
        <div className="pt-2 px-2">
          <img className="rounded-t-lg object-cover aspect-16/9" src={`/api/images/${menu.url}`} alt="" />
        </div>
        {/* Image */}
        {/* Content */}
        <div className="space-y-1 pb-2 px-4">
          <div className="flex justify-between items-center">
            <p className="font-medium">{menu.nama}</p>
            <p className="text-muted-foreground text-sm">{menu.kategori.nama}</p>
          </div>
          <p className="leading-tight text-sm text-muted-foreground">
            {menu.deskripsi} Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur
            cupidatat.
          </p>
          <p className="font-medium text-end">Rp. {menu.harga}</p>
        </div>
        {/* Content */}
      </div>
      {/* Hover Hint Edit */}
      <div className="absolute top-0 w-full h-full flex items-center justify-center opacity-0 hover:opacity-100 transition-opacity duration-300 bg-black bg-opacity-[25%] rounded-md">
        <p className="text-white font-semibold flex gap-2">
          <SquarePen />
          Klik untuk edit
        </p>
      </div>
      {/* Hover Hint Edit */}
    </Card>
  );
}
