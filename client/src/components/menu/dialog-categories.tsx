import { Input } from '@/components/ui/input';
import { useForm } from 'react-hook-form';
import { toast } from '@/components/ui/use-toast';
import { Button } from '@/components/ui/button';
import { Plus, Loader2, List, Edit2, Trash2, X, Check } from 'lucide-react';
import { useEffect, useState } from 'react';
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { MenuCategorySchema, menuCategorySchema } from '@/lib/schemas/menu';
import { useMenu } from '@/context/MenuProvider';

export default function DialogCategories() {
  const [open, setOpen] = useState(false);
  const [addMode, setAddMode] = useState(false);
  const { categories } = useMenu();

  useEffect(() => {
    setAddMode(false);
  }, [open]);

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      {/* Dialog Trigger Button */}
      <Button variant="outline" className="space-x-1" onClick={() => setOpen(true)}>
        <List />
        <span>Kategori</span>
      </Button>
      {/* Dialog Trigger Button */}

      <DialogContent className="sm:max-w-[400px]">
        <DialogHeader className="mb-4">
          <DialogTitle>Kategori Menu</DialogTitle>
        </DialogHeader>
        <div>
          <div className="space-y-2 mb-4">
            {categories?.map((category) => <CategoryTile key={category.id} category={category} />)}
          </div>
          {addMode ? (
            <AddForm setAddMode={setAddMode} />
          ) : (
            <div className="flex justify-end">
              <Button size="sm" variant="outline" onClick={() => setAddMode(true)}>
                <Plus className="mr-1" />
                Kategori
              </Button>
            </div>
          )}
        </div>
      </DialogContent>
    </Dialog>
  );
}

type MenuCategoryProps = {
  category: Kategori;
};

function CategoryTile({ category }: MenuCategoryProps) {
  const [deleteLoading, setDeleteLoading] = useState(false);
  const [editMode, setEditMode] = useState(false);

  const { updateMenuCategory, deleteMenuCategory, invalidate } = useMenu();
  const editForm = useForm<MenuCategorySchema>({
    resolver: zodResolver(menuCategorySchema),
    defaultValues: {
      nama: category.nama,
    },
  });

  const {
    handleSubmit,
    control,
    formState: { isSubmitting },
  } = editForm;

  async function onSubmit(data: MenuCategorySchema) {
    const res = await updateMenuCategory(category.id, data);
    if (!res.ok) {
      toast({
        description: 'Gagal update kategori menu!',
        variant: 'destructive',
      });
      return;
    }
    await invalidate();
    setEditMode(false);
    toast({
      description: 'Kategori Menu berhasil diupdate!',
    });
  }

  async function handleDelete() {
    setDeleteLoading(true);
    const res = await deleteMenuCategory(category.id);
    if (!res.ok) {
      toast({
        description: 'Gagal menghapus kategori menu!',
        variant: 'destructive',
      });
      setDeleteLoading(false);
      return;
    }
    await invalidate();
    toast({
      description: 'Kategori Menu berhasil dihapus!',
    });

    setDeleteLoading(false);
  }

  return (
    <Form {...editForm}>
      <form onSubmit={handleSubmit(onSubmit)} className="flex w-full justify-between gap-2 items-center">
        {editMode ? (
          <FormField
            control={control}
            name="nama"
            render={({ field }) => (
              <FormItem>
                <FormLabel hidden htmlFor="nama">
                  Nama
                </FormLabel>
                <FormControl>
                  <Input placeholder="Nama kategori" id="nama" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        ) : (
          <p>{category.nama}</p>
        )}

        {editMode && (
          <div className="flex gap-2">
            <Button size="icon" variant="outline" onClick={() => setEditMode(false)} disabled={isSubmitting}>
              <X size={18} />
            </Button>
            <Button size="icon" type="submit" disabled={isSubmitting}>
              {isSubmitting ? <Loader2 className="animate-spin" /> : <Check size={18} />}
            </Button>
          </div>
        )}

        {!editMode && (
          <div className="flex gap-2">
            <Button size="icon" onClick={() => setEditMode(true)} disabled={deleteLoading}>
              <Edit2 size={18} />
            </Button>
            <Button size="icon" variant="destructive" type="button" onClick={handleDelete} disabled={deleteLoading}>
              {deleteLoading ? <Loader2 className="animate-spin" /> : <Trash2 size={18} />}
            </Button>
          </div>
        )}
      </form>
    </Form>
  );
}

type AddFormProps = {
  setAddMode: (mode: boolean) => void;
};
function AddForm({ setAddMode }: AddFormProps) {
  const form = useForm<MenuCategorySchema>({
    resolver: zodResolver(menuCategorySchema),
    defaultValues: {
      nama: '',
    },
  });

  const { createMenuCategory, invalidate } = useMenu();

  const {
    handleSubmit,
    control,
    reset,
    formState: { isSubmitting },
  } = form;

  const onSubmit = async (data: MenuCategorySchema) => {
    const res = await createMenuCategory(data);
    if (!res.ok) {
      toast({
        description: 'Gagal menambahkan kategori menu',
        variant: 'destructive',
      });
      return;
    }
    reset();
    await invalidate();
    setAddMode(false);
    toast({
      description: 'Kategori Menu berhasil ditambahkan',
    });
  };

  return (
    <Form {...form}>
      <form onSubmit={handleSubmit(onSubmit)} className="flex justify-between gap-4">
        <FormField
          control={control}
          name="nama"
          render={({ field }) => (
            <FormItem>
              <FormLabel hidden htmlFor="nama">
                Nama
              </FormLabel>
              <FormControl>
                <Input placeholder="Nama kategori" id="nama" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <div className="flex gap-2">
          <Button
            size="icon"
            variant="outline"
            onClick={() => {
              reset();
              setAddMode(false);
            }}
            disabled={isSubmitting}
          >
            <X size={18} />
          </Button>
          <Button size="icon" type="submit" disabled={isSubmitting}>
            {isSubmitting ? <Loader2 className="animate-spin" /> : <Check size={18} />}
          </Button>
        </div>
      </form>
    </Form>
  );
}
