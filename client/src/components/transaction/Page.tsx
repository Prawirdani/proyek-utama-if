import { useTransaction } from '@/context/TransactionProvider';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import TitleSetter from '../pageTitle';
import { H2 } from '../typography';
import Loader from '../ui/loader';
import { Card } from '../ui/card';
import { titleCase } from '@/lib/utils';
import { formatDateTime, formatIDR } from '@/lib/formatter';
import { Button } from '../ui/button';
import { ChevronLeft, ChevronRight, Loader2 } from 'lucide-react';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import StatusBadge from './status-badge';

export default function Page() {
  const { loading, transactions, pagination, pageLoading } = useTransaction();
  const { page, nextPage, prevPage, limit, setLimit } = pagination;

  return loading ? (
    <Loader />
  ) : (
    <section>
      <TitleSetter title="Meja" />
      <div className="-space-y-1 mb-4">
        <H2>Transaksi</H2>
        <p>Daftar Transaksi</p>
      </div>
      <Card className="p-8">
        <div className="flex justify-end mb-4 font-medium">
          <Select onValueChange={(e) => setLimit(Number(e))} defaultValue={String(limit)}>
            <SelectTrigger className="w-[120px] shadow-sm">
              <SelectValue placeholder="Baris Data" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="10">10 Baris</SelectItem>
              <SelectItem value="15">15 Baris</SelectItem>
              <SelectItem value="20">20 Baris</SelectItem>
              <SelectItem value="25">25 Baris</SelectItem>
              <SelectItem value="30">30 Baris</SelectItem>
            </SelectContent>
          </Select>
        </div>
        <div className="mb-4 rounded-md border">
          <Table className="2xl:table-fixed">
            <TableHeader>
              <TableRow className="[&>th]:whitespace-nowrap">
                <TableHead>Waktu Pesanan</TableHead>
                <TableHead>Pelanggan</TableHead>
                <TableHead>Tipe Pesanan</TableHead>
                <TableHead>Kasir</TableHead>
                <TableHead>Status</TableHead>
                <TableHead>Total</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {transactions?.map((tx) => (
                <TableRow key={tx.id} className="[&>td]:whitespace-nowrap">
                  <TableCell>{formatDateTime(new Date(tx.waktuPesanan))}</TableCell>
                  <TableCell>{tx.namaPelanggan}</TableCell>
                  <TableCell>{titleCase(tx.tipe)}</TableCell>
                  <TableCell>{tx.kasir}</TableCell>
                  <TableCell>
                    <StatusBadge status={tx.status} />
                  </TableCell>
                  <TableCell className="font-medium">{formatIDR(tx.total)}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </div>
        <div className="flex justify-end gap-4 items-center">
          {pageLoading && <Loader2 className="animate-spin text-primary" />}
          <Button size="icon" className="shadow-lg" onClick={prevPage} disabled={page === 1 || pageLoading}>
            <ChevronLeft />
          </Button>
          <Button size="icon" className="shadow-lg" onClick={nextPage} disabled={pageLoading}>
            <ChevronRight />
          </Button>
        </div>
      </Card>
    </section>
  );
}
