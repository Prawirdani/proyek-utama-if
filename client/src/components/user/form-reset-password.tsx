import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { Button } from '@/components/ui/button';
import { Loader2 } from 'lucide-react';
import { Input } from '@/components/ui/input';
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { useEffect, useState } from 'react';
import { toast } from '@/components/ui/use-toast';
import { useUser } from '@/context/UserProvider';
import { UserResetPasswordSchema, userResetPasswordSchema } from '@/lib/schemas/user';
import { isErrorResponse } from '@/api/fetcher';

interface Props {
  open: boolean;
  setOpen: (open: boolean) => void;
  updateTarget: User;
}
export default function FormResetPassword({ open, setOpen, updateTarget }: Props) {
  const [apiError, setApiError] = useState<string | null>(null);

  const { invalidate, resetPassword } = useUser();

  const form = useForm<UserResetPasswordSchema>({
    resolver: zodResolver(userResetPasswordSchema),
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
      newPassword: '',
      repeatPassword: '',
    });
    setApiError(null);
  }, [open, updateTarget]);

  const onSubmit = async (data: UserResetPasswordSchema) => {
    const res = await resetPassword(data);
    if (!res.ok) {
      const resBody = await res.json();
      setApiError(isErrorResponse(resBody) ? resBody.error.message : 'Terjadi kesalahan');
      return;
    }
    invalidate();
    reset();
    toast({
      description: 'Berhasil reset password pengguna.',
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
              <DialogTitle>Reset Password Pengguna</DialogTitle>
            </DialogHeader>
            <div className="mb-4 space-y-2">
              {/* Input New Password */}
              <FormField
                control={control}
                name="newPassword"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel htmlFor="newPassword">Password Baru</FormLabel>
                    <FormControl>
                      <Input id="newPassword" type="password" placeholder="Masukkan password baru" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              {/* Input New Password */}

              {/* Input Repeat Password */}
              <FormField
                control={control}
                name="repeatPassword"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel htmlFor="repeatPassword">Ulangi password</FormLabel>
                    <FormControl>
                      <Input
                        id="repeatPassword"
                        type="password"
                        placeholder="Masukkan ulang password baru"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              {/* Input Repeat Password */}

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
