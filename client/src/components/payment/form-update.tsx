import { Textarea } from '@/components/ui/textarea';
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { Button } from '@/components/ui/button';
import { Loader2 } from 'lucide-react';
import { Input } from '@/components/ui/input';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { useEffect, useState } from 'react';
import { usePaymentMethod } from '@/context/PaymentMethodProvider';
import { toast } from '@/components/ui/use-toast';
import { UpdatePaymentMethodSchema, updatePaymentMethodSchema } from '@/lib/schemas/payment';
import { isErrorResponse } from '@/api/fetcher';

interface Props {
  open: boolean;
  setOpen: (open: boolean) => void;
  updateTarget: MetodePembayaran;
}
export default function FormUpdate({ open, setOpen, updateTarget }: Props) {
  const [apiError, setApiError] = useState<string | null>(null);

  useEffect(() => {}, [open, updateTarget]);

  const { tipe_pembayaran_opts, updateMetodePembayaran, invalidate } = usePaymentMethod();

  const form = useForm<UpdatePaymentMethodSchema>({
    resolver: zodResolver(updatePaymentMethodSchema),
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
      tipePembayaran: updateTarget.tipePembayaran,
      metode: updateTarget.metode,
      deskripsi: updateTarget.deskripsi,
    });
    setApiError(null);
  }, [open, updateTarget]);

  const onSubmit = async (data: UpdatePaymentMethodSchema) => {
    const res = await updateMetodePembayaran(data);
    if (!res.ok) {
      const resBody = await res.json();
      setApiError(isErrorResponse(resBody) ? resBody.error.message : 'Terjadi kesalahan');
      return;
    }
    invalidate();
    reset();
    toast({
      description: 'Berhasil update metode pembayaran.',
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
              <DialogTitle>Update Metode Pembayaran</DialogTitle>
            </DialogHeader>
            <div className="mb-4 space-y-2">
              <div className="grid grid-cols-2 gap-4">
                {/* Input nama metode */}
                <FormField
                  control={control}
                  name="metode"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel htmlFor="metode">Nama Metode</FormLabel>
                      <FormControl>
                        <Input id="metode" placeholder="Masukkan metode pembayaran" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                {/* Input nama metode */}

                {/* Select Input tipe pembayaran */}
                <FormField
                  control={control}
                  name="tipePembayaran"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel htmlFor="tipePembayaran">Tipe Pembayaran</FormLabel>
                      <Select onValueChange={field.onChange} name={field.name}>
                        <FormControl id="tipePembayaran">
                          <SelectTrigger>
                            <SelectValue placeholder="Tipe Pembayaran" />
                          </SelectTrigger>
                        </FormControl>
                        <SelectContent>
                          {tipe_pembayaran_opts.map((opt, i) => (
                            <SelectItem key={i} value={opt}>
                              {opt}
                            </SelectItem>
                          ))}
                        </SelectContent>
                      </Select>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                {/* Select Input tipe pembayaran */}
              </div>

              <div>
                {/* Input deskripsi */}
                <FormField
                  control={control}
                  name="deskripsi"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel htmlFor="deskripsi">Deskripsi</FormLabel>
                      <FormControl>
                        <Textarea id="deskripsi" placeholder="Masukkan deskripsi metode pembayaran" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                {/* Input deskripsi */}
              </div>

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
