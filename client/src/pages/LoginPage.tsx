import TitleSetter from '@/components/pageTitle';
import { Button } from '@/components/ui/button';
import { Card } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { useAuth } from '@/context/useAuth';
import { Loader2 } from 'lucide-react';
import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { H2 } from '@/components/typography';

const formSchema = z.object({
  username: z.string().min(1, { message: 'Mohon isi kolom username' }),
  password: z.string().min(1, { message: 'Mohon isi kolom password' }),
});

export default function LoginPage() {
  const navigate = useNavigate();
  const [apiError, setApiError] = useState<string | null>(null);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      username: '',
      password: '',
    },
  });

  const { login } = useAuth();

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    const res = await login(values.username, values.password);
    if (!res.ok) {
      setApiError(res.status === 401 ? 'Username atau password salah!' : 'Terjadi kesalahan');
      return;
    }
    navigate('/', { replace: true });
  };
  return (
    <>
      <TitleSetter title="Login" />
      <div className="h-screen flex place-items-center bg-secondary overflow-hidden">
        <Card className="mx-auto w-[calc(100%-10%)] xs:w-[350px] shadow-lg">
          <Form {...form}>
            <form autoComplete="on" onSubmit={form.handleSubmit(onSubmit)} className="space-y-8 p-8">
              <H2 className="text-center">Login</H2>
              <div className="space-y-2">
                <FormField
                  control={form.control}
                  name="username"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Username</FormLabel>
                      <FormControl>
                        <Input autoComplete="on" placeholder="Masukkan username anda" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <FormField
                  control={form.control}
                  name="password"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Password</FormLabel>
                      <FormControl>
                        <Input autoComplete="on" type="password" placeholder="Masukkan password anda" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
              <div>
                {apiError && <p className="text-destructive font-medium text-sm">{apiError}</p>}
                <Button type="submit" className="w-full" disabled={form.formState.isSubmitting}>
                  {form.formState.isSubmitting ? (
                    <>
                      <Loader2 className="animate-spin mr-2" />
                      <span>Mohon tunggu</span>
                    </>
                  ) : (
                    <span>Login</span>
                  )}
                </Button>
              </div>
            </form>
          </Form>
        </Card>
      </div>
    </>
  );
}
