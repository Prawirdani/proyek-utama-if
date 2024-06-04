import { titleCase } from '@/lib/utils';
import { Badge } from '../ui/badge';

interface Props {
  status: Transaksi['status'];
}

const statusVariants: Record<Transaksi['status'], 'default' | 'secondary' | 'destructive'> = {
  Diproses: 'secondary',
  Selesai: 'default',
  Batal: 'destructive',
};

export default function StatusBadge({ status }: Props) {
  const variant = statusVariants[status];
  return (
    <Badge className="w-[4.5rem] flex justify-center" variant={variant}>
      <p>{titleCase(status)}</p>
    </Badge>
  );
}
