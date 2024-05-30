import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { Button } from '@/components/ui/button';
import { Loader2, Plus } from 'lucide-react';
import { Input } from '@/components/ui/input';
import { toast } from '@/components/ui/use-toast';
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { useEffect, useState } from 'react';
import { useUsers } from '@/context/UserProvider';
import { UserRegisterSchema, userRegisterSchema } from '@/lib/schemas/user';

export default function FormRegister() {
  const { invalidate, registerUser } = useUsers();
  const [open, setOpen] = useState(false);
  const [apiError, setApiError] = useState<string | null>(null);

  const form = useForm<UserRegisterSchema>({
    resolver: zodResolver(userRegisterSchema),
    defaultValues: {
      nama: '',
      username: '',
      password: '',
      repeatPassword: '',
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

  const onSubmit = async (data: UserRegisterSchema) => {
    const res = await registerUser(data);
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
        <span>Akun Kasir</span>
      </Button>
      {/* Dialog Trigger Button */}

      <DialogContent className="sm:max-w-[525px]">
        <Form {...form}>
          <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
            <DialogHeader className="mb-4">
              <DialogTitle>Register Pengguna Kasir</DialogTitle>
            </DialogHeader>
            <div className="mb-4 space-y-2">
              <FormField
                control={control}
                name="nama"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel htmlFor="nama">Nama Kasir</FormLabel>
                    <FormControl>
                      <Input autoComplete="off" id="nama" placeholder="Masukkan nama kasir" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={control}
                name="username"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel htmlFor="username">Username</FormLabel>
                    <FormControl>
                      <Input autoComplete="user-off" id="username" placeholder="Masukkan username kasir" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={control}
                name="password"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel htmlFor="password">Password</FormLabel>
                    <FormControl>
                      <Input
                        autoComplete="pass-off"
                        id="password"
                        type="password"
                        placeholder="Masukkan password kasir"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={control}
                name="repeatPassword"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel htmlFor="repeatPassword">Ulangi Password</FormLabel>
                    <FormControl>
                      <Input
                        autoComplete="off"
                        id="repeatPassword"
                        type="password"
                        placeholder="Ulangi password"
                        {...field}
                      />
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
