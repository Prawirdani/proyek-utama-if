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
import { useAuth } from '@/hooks/useAuth';
import { FieldValues, useForm } from 'react-hook-form';

export default function LoginPage() {
  const { login } = useAuth();
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
    reset,
  } = useForm();

  const onSubmit = async (data: FieldValues) => {
    console.log(data);
    await login(data.username, data.password);
    reset();
  };
  return (
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
          </CardContent>
          <CardFooter className="mb-1">
            <Button type="submit" className="w-full" disabled={isSubmitting}>
              Login
            </Button>
          </CardFooter>
        </form>
      </Card>
    </div>
  );
}
