import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { Button } from '@/components/ui/button';
import { Loader2 } from 'lucide-react';
import { Input } from '@/components/ui/input';
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { useEffect, useState } from 'react';
import { toast } from '@/components/ui/use-toast';
import { useTables } from '@/context/TableProvider';
import { UpdateTableSchema, updateTableSchema } from '@/lib/schemas/table';

interface Props {
  open: boolean;
  setOpen: (open: boolean) => void;
  updateTarget: Meja;
}
export default function FormUpdate({ open, setOpen, updateTarget }: Props) {
  const [apiError, setApiError] = useState<string | null>(null);

  useEffect(() => {}, [open, updateTarget]);

  const { updateMeja, invalidate } = useTables();

  const form = useForm<UpdateTableSchema>({
    resolver: zodResolver(updateTableSchema),
  });

  const {
    handleSubmit,
    control,
    reset,
    formState: { isSubmitting },
  } = form;

  useEffect(() => {
    reset({
      id: updateTarget.id,
      nomor: updateTarget.nomor,
    });
    setApiError(null);
  }, [open, updateTarget]);

  const onSubmit = async (data: UpdateTableSchema) => {
    const res = await updateMeja(data);
    if (!res.ok) {
      const resBody = (await res.json()) as ErrorResponse;
      res.status === 409 ? setApiError(resBody.error.message) : setApiError('Terjadi kesalahan');
      return;
    }
    invalidate();
    reset();
    toast({
      description: 'Berhasil update meja.',
    });
    setOpen(false);
    setApiError(null);
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent className="sm:max-w-[525px]">
        <Form {...form}>
          <form onSubmit={handleSubmit(onSubmit)}>
            <DialogHeader className="mb-4">
              <DialogTitle>Update Meja</DialogTitle>
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
