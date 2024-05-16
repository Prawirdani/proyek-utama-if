import { fetchTables } from '@/api/table';
import TitleSetter from '@/components/pageTitle';
import { H2 } from '@/components/typography';
import Loader from '@/components/ui/loader';
import { useEffect, useState } from 'react';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Card } from '@/components/ui/card';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { useForm } from 'react-hook-form';
import { Button } from '@/components/ui/button';
import { Loader2, Plus } from 'lucide-react';
import { Input } from '@/components/ui/input';
import { toast } from '@/components/ui/use-toast';

export default function TablePage() {
  const [tables, setTables] = useState<Meja[] | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchTables().then((data) => {
      setTables(data);
      setLoading(false);
    });
  }, []);
  return loading ? (
    <Loader />
  ) : (
    <section className="relative">
      <TitleSetter title="Menu" />
      <div className="-space-y-1 mb-4">
        <H2>Meja</H2>
        <p>Manajemen Meja</p>
      </div>

      <div className="flex justify-end mb-4">
        <TableForm setTables={setTables} />
      </div>
      <Card className="p-4">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Nomor Meja</TableHead>
              <TableHead>Status</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {tables?.map((table) => (
              <TableRow key={table.id}>
                <TableCell>{table.nomor}</TableCell>
                <TableCell>{table.status}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </Card>
    </section>
  );
}

const formSchema = z.object({
  nomor: z.string().min(1, { message: 'Mohon isi kolom nomor meja' }),
});

interface TableFormProps {
  setTables: (ts: Meja[] | null) => void;
}
const TableForm = ({ setTables }: TableFormProps) => {
  const [open, setOpen] = useState(false);
  const [apiError, setApiError] = useState<string | null>(null);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      nomor: '',
    },
  });

  const {
    handleSubmit,
    control,
    reset,
    formState: { isSubmitting },
  } = form;

  useEffect(() => {
    reset();
    setApiError(null);
  }, [open]);

  const onSubmit = async (data: z.infer<typeof formSchema>) => {
    const res = await fetch('/api/v1/tables', {
      method: 'POST',
      body: JSON.stringify({
        nomor: data.nomor,
      }),
    });
    const resBody = await res.json();
    if (res.ok) {
      reset();
      toast({
        description: 'Berhasil menambahkan meja.',
      });
      setOpen(false);
      setApiError(null);
      fetchTables().then(setTables);
      return;
    }
    res.status === 409
      ? setApiError((resBody as { error: ErrorResponse }).error.message)
      : setApiError('Terjadi kesalahan');
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      {/* Dialog Trigger Button */}
      <Button className="space-x-1" onClick={() => setOpen(true)}>
        <Plus />
        <span>Meja</span>
      </Button>
      {/* Dialog Trigger Button */}

      <DialogContent className="sm:max-w-[525px]">
        <Form {...form}>
          <form onSubmit={handleSubmit(onSubmit)}>
            <DialogHeader className="mb-4">
              <DialogTitle>Tambah meja baru</DialogTitle>
            </DialogHeader>
            <div className="mb-4 space-y-2">
              <FormField
                control={control}
                name="nomor"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel htmlFor="nomor">Nama Menu</FormLabel>
                    <FormControl>
                      <Input id="nomor" placeholder="Masukkan nama menu" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <p className="text-sm text-destructive">{apiError}</p>
            </div>
            <div className="flex justify-end">
              <Button type="submit">
                {isSubmitting && <Loader2 />}
                <span>Simpan</span>
              </Button>
            </div>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
};
