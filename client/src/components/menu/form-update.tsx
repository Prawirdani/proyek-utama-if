import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { useForm } from 'react-hook-form';
import { toast } from '@/components/ui/use-toast';
import { Button } from '@/components/ui/button';
import { Loader2 } from 'lucide-react';
import { useEffect, useState } from 'react';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { UpdateMenuSchema, updateMenuSchema } from '@/lib/schemas/menu';
import { useMenu } from '@/context/MenuProvider';

interface Props {
  open: boolean;
  setOpen: (open: boolean) => void;
  updateTarget: Menu;
}

export default function FormUpdate({ open, setOpen, updateTarget }: Props) {
  const [deleteLoading, setDeleteLoading] = useState(false);

  const form = useForm<UpdateMenuSchema>({
    resolver: zodResolver(updateMenuSchema),
  });

  const {
    handleSubmit,
    control,
    reset,
    register,
    formState: { isSubmitting, errors },
  } = form;

  const { categories, invalidate, updateMenu, deleteMenu } = useMenu();

  async function onSubmit(data: UpdateMenuSchema) {
    const formData = new FormData();
    formData.append(
      'data',
      JSON.stringify({
        nama: data.nama,
        deskripsi: data.deskripsi,
        harga: Number(data.harga),
        kategoriId: Number(data.kategori_id),
      }),
    );
    if (data.image) {
      formData.append('image', data.image[0]);
    }

    const res = await updateMenu(updateTarget.id, formData);
    if (!res.ok) {
      toast({ variant: 'destructive', description: 'Gagal mengupdate menu!' });
      return;
    }

    toast({
      description: 'Menu berhasil diupdate!',
      duration: 2000,
    });
    await invalidate();
    setOpen(false);
  }

  async function handleDelete() {
    setDeleteLoading(true);
    const res = await deleteMenu(updateTarget.id);
    if (!res.ok) {
      toast({ variant: 'destructive', description: 'Gagal menghapus menu!' });
      setDeleteLoading(false);
      return;
    }
    toast({
      description: 'Menu berhasil dihapus!',
      duration: 2000,
    });

    await invalidate();
    setOpen(false);
    setDeleteLoading(false);
  }

  const [imagePreview, setImagePreview] = useState<string | null>(null);
  let imageInputRef: HTMLInputElement;
  const { ref, ...imageRegister } = register('image');

  const handleImageClick = () => {
    imageInputRef.click();
  };

  const imageOnChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files;
    form.setValue('image', e.target.files);
    if (files && files[0]) {
      const imageUrl = URL.createObjectURL(files[0]);
      setImagePreview(imageUrl);
    }
  };

  useEffect(() => {
    reset({
      nama: updateTarget.nama ?? '',
      harga: updateTarget.harga ?? 0,
      kategori_id: String(updateTarget.kategori?.id) ?? '',
      deskripsi: updateTarget.deskripsi ?? '',
    });
    setImagePreview(null);
  }, [open]);

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent className="sm:max-w-[800px]">
        <Form {...form}>
          <form onSubmit={handleSubmit(onSubmit)}>
            <DialogHeader className="mb-4">
              <DialogTitle>Update data menu</DialogTitle>
            </DialogHeader>
            <div className="grid grid-cols-5 gap-4 mb-4">
              <div
                className="col-span-2 border rounded-md justify-center items-center flex flex-col hover:cursor-pointer"
                onClick={handleImageClick}
              >
                <input
                  type="file"
                  id="image"
                  hidden
                  {...imageRegister}
                  ref={(e) => {
                    imageInputRef = e!;
                  }}
                  onChange={imageOnChange}
                />
                {imagePreview ? (
                  <img src={imagePreview} alt="Preview" className="h-full w-full object-cover" />
                ) : (
                  <img src={`/api/images/${updateTarget.url}`} className="h-full w-full object-cover" />
                )}
                {errors.image && (
                  <span className="mt-2 text-sm text-destructive text-center">{String(errors.image.message)}</span>
                )}
              </div>
              <div className="space-y-2 col-span-3">
                <FormField
                  control={control}
                  name="nama"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel htmlFor="nama">Nama Menu</FormLabel>
                      <FormControl>
                        <Input id="nama" placeholder="Masukkan nama menu" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <div className="grid grid-cols-2 gap-4">
                  <FormField
                    control={control}
                    name="harga"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel htmlFor="harga">Harga</FormLabel>
                        <FormControl>
                          <Input
                            placeholder="Masukkan harga menu"
                            id="harga"
                            type="number"
                            {...field}
                            onChange={(e) => field.onChange(e.target.valueAsNumber)}
                          />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={control}
                    name="kategori_id"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel htmlFor="kategori_id">Kategori Menu</FormLabel>
                        <Select onValueChange={field.onChange} name={field.name}>
                          <FormControl id="kategori_id">
                            <SelectTrigger>
                              <SelectValue placeholder={updateTarget.kategori.nama} defaultValue={field.value} />
                            </SelectTrigger>
                          </FormControl>
                          <SelectContent>
                            {categories.map((category) => (
                              <SelectItem key={category.id} value={String(category.id)}>
                                {category.nama}
                              </SelectItem>
                            ))}
                          </SelectContent>
                        </Select>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                </div>

                <div>
                  <FormField
                    control={control}
                    name="deskripsi"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel htmlFor="deskripsi">Deskripsi</FormLabel>
                        <FormControl>
                          <Textarea id="deskripsi" placeholder="Masukkan deskripsi menu" {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                </div>
              </div>
            </div>

            <div className="flex justify-end gap-2">
              <Button
                type="button"
                onClick={handleDelete}
                variant="destructive"
                disabled={deleteLoading || isSubmitting}
              >
                <span>Hapus menu</span>
                {deleteLoading && <Loader2 className="animate-spin ml-2" />}
              </Button>
              <Button type="submit" disabled={isSubmitting || deleteLoading}>
                {isSubmitting && <Loader2 className="animate-spin mr-2" />}
                <span>Simpan</span>
              </Button>
            </div>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
