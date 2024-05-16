import { toast } from '@/components/ui/use-toast';

export function Fetch(fn: Function, toastSuccessMessage?: string) {
  return async (...args: any[]) => {
    try {
      const result = await fn(...args);
      if (toastSuccessMessage) {
        toast({ description: toastSuccessMessage });
      }
      return result;
    } catch (error) {
      toast({ description: (error as Error).message, variant: 'destructive' });
    }
  };
}
