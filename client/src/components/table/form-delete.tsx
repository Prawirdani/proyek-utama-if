import { isErrorResponse } from '@/api/fetcher';
import { Button } from '@/components/ui/button';
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { toast } from '@/components/ui/use-toast';
import { useTable } from '@/context/TableProvider';
import { useEffect, useState } from 'react';

interface Props {
  id: number;
  open: boolean;
  setOpen: (open: boolean) => void;
}

export default function FormDelete({ id, open, setOpen }: Props) {
  const [apiError, setApiError] = useState<string | null>(null);

  const { deleteMeja, invalidate } = useTable();

  useEffect(() => {
    setApiError(null);
  }, [open]);

  const handleDelete = async () => {
    const res = await deleteMeja(id);
    if (!res.ok) {
      const resBody = await res.json();
      setApiError(isErrorResponse(resBody) ? resBody.error.message : 'Terjadi kesalahan');
      return;
    }
    invalidate();
    toast({ description: 'Berhasil menghapus meja.' });
    setOpen(false);
    setApiError(null);
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent className="sm:max-w-[525px]">
        <DialogHeader className="mb-4">
          <DialogTitle>Hapus Meja</DialogTitle>
        </DialogHeader>
        <p>Apakah Anda yakin ingin menghapus meja ini?</p>
        <p className="text-sm text-destructive">{apiError}</p>
        <div className="flex justify-end [&>button]:w-24 gap-2">
          <Button type="button" variant="secondary" onClick={() => setOpen(false)}>
            <span>Batal</span>
          </Button>
          <Button type="button" onClick={handleDelete}>
            <span>Ya</span>
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  );
}
