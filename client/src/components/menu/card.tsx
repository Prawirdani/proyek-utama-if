import { Card } from '@/components/ui/card';

interface Props {
  menu: Menu;
}
export default function MenuCard({ menu }: Props) {
  return (
    <Card className="shadow-lg hover:cursor-pointer hover:bg-gray-200">
      {/* Image */}
      <div className="px-2 pt-2">
        <img className="rounded-lg object-cover aspect-16/9" src={`/api/images/${menu.url}`} alt="" />
      </div>
      {/* Image */}

      {/* Content */}
      <div className="space-y-1 pb-2 py-3 px-5">
        <div className="flex justify-between items-center">
          <p className="font-medium">{menu.nama}</p>
          <p className="text-muted-foreground text-sm">{menu.kategori.nama}</p>
        </div>
        <div className="space-y-4">
          <p className="h-12 leading-tight text-sm text-muted-foreground line-clamp-3">{menu.deskripsi}</p>
          <p className="font-medium text-end">Rp. {menu.harga}</p>
        </div>
      </div>
      {/* Content */}
    </Card>
  );
}
