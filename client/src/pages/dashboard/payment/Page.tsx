import TitleSetter from '@/components/pageTitle';
import { H2 } from '@/components/typography';
import Loader from '@/components/ui/loader';
import { Card } from '@/components/ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { usePaymentMethods } from '@/context/PaymentMethodsProvider';
import { PaymentMethodAddForm } from './AddForm';
import { useState } from 'react';
import { PaymentMethodUpdateForm } from './UpdateForm';
import { Button } from '@/components/ui/button';
import Delete from './Delete';
import { SquarePen, Trash } from 'lucide-react';

export default function Page() {
  const { loading, payment_methods } = usePaymentMethods();
  const [openUpdateDialog, setOpenUpdateDialog] = useState(false);
  const [openDeleteDialog, setOpenDeleteDialog] = useState(false);
  const [updateTarget, setUpdateTarget] = useState<MetodePembayaran>({} as MetodePembayaran);

  const triggerUpdateForm = (m: MetodePembayaran) => {
    setUpdateTarget(m);
    setOpenUpdateDialog(true);
  };

  const triggerDeleteForm = (m: MetodePembayaran) => {
    setUpdateTarget(m);
    setOpenDeleteDialog(true);
  };

  return loading ? (
    <Loader />
  ) : (
    <section>
      <TitleSetter title="Menu" />
      <div className="-space-y-1 mb-4">
        <H2>Pembayaran</H2>
        <p>Manajemen metode pembayaran</p>
      </div>

      <div className="flex justify-end mb-4">
        <PaymentMethodAddForm />
        {/* Update Form, Triggered by each row edit button */}
        <PaymentMethodUpdateForm updateTarget={updateTarget} open={openUpdateDialog} setOpen={setOpenUpdateDialog} />
        <Delete id={updateTarget.id} open={openDeleteDialog} setOpen={setOpenDeleteDialog} />
      </div>
      <Card className="p-8">
        <Table>
          <TableHeader>
            <TableRow className="[&>th]:text-medium">
              <TableHead>Nama Metode</TableHead>
              <TableHead>Tipe Pembayaran</TableHead>
              <TableHead>Deskripsi</TableHead>
              <TableHead className="w-[10%]"></TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {payment_methods?.map((m) => (
              <TableRow key={m.id}>
                <TableCell>{m.metode}</TableCell>
                <TableCell>{m.tipePembayaran}</TableCell>
                <TableCell className="max-w-lg">{m.deskripsi}</TableCell>
                <TableCell className="w-fit flex gap-4 [&>button]:shadow-md [&>button]:w-12 [&>button]:p-0">
                  <Button onClick={() => triggerUpdateForm(m)} variant="outline">
                    <SquarePen className="h-4 w-4" />
                  </Button>
                  <Button onClick={() => triggerDeleteForm(m)} variant="destructive">
                    <Trash className="h-4 w-4" />
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </Card>
    </section>
  );
}
