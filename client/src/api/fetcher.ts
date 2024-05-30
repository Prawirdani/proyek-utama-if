import { toast } from '@/components/ui/use-toast';

export function Fetch<T>(fn: (...args: any[]) => Promise<T>, toastSuccessMessage?: string) {
  return async (...args: any[]): Promise<T> => {
    try {
      const result = await fn(...args);
      if (toastSuccessMessage) {
        toast({ description: toastSuccessMessage });
      }
      return result;
    } catch (error) {
      toast({ description: (error as Error).message, variant: 'destructive' });
      throw error; // Rethrow the error to maintain the original behavior
    }
  };
}

export function isErrorResponse(body: any): body is ErrorResponse {
  return (
    typeof body === 'object' &&
    body !== null &&
    'error' in body &&
    typeof body.error === 'object' &&
    'code' in body.error &&
    typeof body.error.code === 'number' &&
    'message' in body.error &&
    typeof body.error.message === 'string'
  );
}
