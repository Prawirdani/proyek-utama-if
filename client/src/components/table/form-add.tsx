import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { Button } from '@/components/ui/button';
import { Loader2, Plus } from 'lucide-react';
import { Input } from '@/components/ui/input';
import { toast } from '@/components/ui/use-toast';
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { useEffect, useState } from 'react';
import { useTables } from '@/context/TableProvider';
import { AddTableSchema, addTableSchema } from '@/lib/schemas/table';

export default function FormAdd() {
  const { invalidate, addMeja } = useTables();
  const [open, setOpen] = useState(false);
  const [apiError, setApiError] = useState<string | null>(null);

  const form = useForm<AddTableSchema>({
    resolver: zodResolver(addTableSchema),
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

  const onSubmit = async (data: AddTableSchema) => {
    const res = await addMeja(data);
    if (!res.ok) {
      const resBody = (await res.json()) as ErrorResponse;
      res.status === 409 ? setApiError(resBody.error.message) : setApiError('Terjadi kesalahan');
      return;
    }
    invalidate();
    reset();
    toast({ description: 'Berhasil menambahkan meja.' });
    setOpen(false);
    setApiError(null);
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
                    <FormLabel htmlFor="nomor">Nomor Meja</FormLabel>
                    <FormControl>
                      <Input id="nomor" placeholder="Masukkan nomor meja" {...field} />
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
}
