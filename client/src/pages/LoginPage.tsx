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

export default function LoginPage() {
  return (
    <div className="h-screen flex place-items-center">
      <Card className="mx-auto w-[calc(100%-5%)] sm:w-[400px] space-y-4 shadow-lg">
        <CardHeader>
          <CardTitle className="text-center">Login</CardTitle>
        </CardHeader>
        <CardContent>
          <form autoComplete="on">
            <div className="grid w-full items-center gap-4">
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="username">Username</Label>
                <Input
                  id="username"
                  autoComplete="on"
                  placeholder="Masukkan username anda"
                />
              </div>
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="password">Password</Label>
                <Input
                  id="password"
                  autoComplete="on"
                  type="password"
                  placeholder="Masukkan password anda"
                />
              </div>
            </div>
          </form>
        </CardContent>
        <CardFooter className="mb-1">
          <Button className="w-full">Login</Button>
        </CardFooter>
      </Card>
    </div>
  );
}
