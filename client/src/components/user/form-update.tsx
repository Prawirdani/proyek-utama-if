import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { Button } from '@/components/ui/button';
import { Loader2 } from 'lucide-react';
import { Input } from '@/components/ui/input';
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { useEffect, useState } from 'react';
import { toast } from '@/components/ui/use-toast';
import { useUsers } from '@/context/UserProvider';
import { UserUpdateSchema, userUpdateSchema } from '@/lib/schemas/user';
import { isErrorResponse } from '@/api/fetcher';

interface Props {
  open: boolean;
  setOpen: (open: boolean) => void;
  updateTarget: User;
}
export default function FormUpdate({ open, setOpen, updateTarget }: Props) {
  const [apiError, setApiError] = useState<string | null>(null);

  const { invalidate, updateUser } = useUsers();

  const form = useForm<UserUpdateSchema>({
    resolver: zodResolver(userUpdateSchema),
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
      nama: updateTarget.nama,
      username: updateTarget.username,
    });
    setApiError(null);
  }, [open, updateTarget]);

  const onSubmit = async (data: UserUpdateSchema) => {
    const res = await updateUser(data);
    if (!res.ok) {
      const resBody = await res.json();
      setApiError(isErrorResponse(resBody) ? resBody.error.message : 'Terjadi kesalahan');
      return;
    }
    invalidate();
    reset();
    toast({
      description: 'Berhasil update akun pengguna.',
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
              <DialogTitle>Update Data Pengguna</DialogTitle>
            </DialogHeader>
            <div className="mb-4 space-y-2">
              {/* Input Nama Kasir */}
              <FormField
                control={control}
                name="nama"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel htmlFor="nama">Nama Kasir</FormLabel>
                    <FormControl>
                      <Input id="nama" placeholder="Masukkan nama kasir" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              {/* Input Nama Kasir */}

              {/* Input Username  */}
              <FormField
                control={control}
                name="username"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel htmlFor="username">Username</FormLabel>
                    <FormControl>
                      <Input id="username" placeholder="Masukkan username pengguna" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              {/* Input Username  */}
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
