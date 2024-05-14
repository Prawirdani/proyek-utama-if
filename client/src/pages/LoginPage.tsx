import TitleSetter from '@/components/pageTitle';
import { Button } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { useAuth } from '@/context/useAuth';
import { Loader2 } from 'lucide-react';
import { useState } from 'react';
import { FieldValues, useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';

export default function LoginPage() {
  const navigate = useNavigate();
  const [apiError, setApiError] = useState<string | null>(null);
  const {
    register,
    handleSubmit,
    formState: { isSubmitting },
  } = useForm();

  const { login } = useAuth();

  const onSubmit = async (data: FieldValues) => {
    const res = await login(data.username, data.password);
    if (!res.ok) {
      setApiError(
        res.status === 401
          ? 'Username atau password salah'
          : 'Terjadi kesalahan',
      );
      return;
    }
    navigate('/', { replace: true });
  };
  return (
    <>
      <TitleSetter title="Login" />
      <div className="h-screen flex place-items-center bg-secondary overflow-hidden">
        <Card className="mx-auto w-[calc(100%-5%)] sm:w-[400px] space-y-4 shadow-lg">
          <form autoComplete="on" onSubmit={handleSubmit(onSubmit)}>
            <CardHeader>
              <CardTitle className="text-center">Login</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="grid w-full items-center gap-4">
                <div className="flex flex-col space-y-1.5">
                  <Label htmlFor="username">Username</Label>
                  <Input
                    {...register('username', { required: true })}
                    id="username"
                    autoComplete="on"
                    placeholder="Masukkan username anda"
                  />
                </div>
                <div className="flex flex-col space-y-1.5">
                  <Label htmlFor="password">Password</Label>
                  <Input
                    {...register('password', { required: true })}
                    id="password"
                    autoComplete="on"
                    type="password"
                    placeholder="Masukkan password anda"
                  />
                </div>
              </div>
              {apiError && (
                <div className="text-destructive text-sm mt-2">{apiError}</div>
              )}
            </CardContent>
            <CardFooter>
              <Button type="submit" className="w-full" disabled={isSubmitting}>
                {isSubmitting ? (
                  <>
                    <Loader2 className="animate-spin mr-2" />
                    Mohon tunggu
                  </>
                ) : (
                  'Login'
                )}
              </Button>
            </CardFooter>
          </form>
        </Card>
      </div>
    </>
  );
}
