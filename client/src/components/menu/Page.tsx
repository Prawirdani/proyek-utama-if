import TitleSetter from '@/components/pageTitle';
import MenuCard from '@/components/menu/card';
import Loader from '@/components/ui/loader';
import { H2 } from '@/components/typography';
import FormAdd from './form-add';
import { useMenu } from '@/context/MenuProvider';
import { useState } from 'react';
import FormUpdate from './form-update';
import DialogCategories from './dialog-categories';

export default function Page() {
  const { loading, menus } = useMenu();
  const [updateTarget, setUpdateTarget] = useState<Menu>({} as Menu);
  const [openUpdateDialog, setOpenUpdateDialog] = useState(false);

  const triggerUpdateDialog = (menu: Menu) => {
    setUpdateTarget(menu);
    setOpenUpdateDialog(true);
  };

  return loading ? (
    <Loader />
  ) : (
    <section className="relative">
      <TitleSetter title="Menu" />
      <div className="-space-y-1 mb-4">
        <H2>Menu</H2>
        <p>Manajemen Menu dan Kategori Menu</p>
      </div>
      <div className="flex justify-end mb-4 gap-4">
        <DialogCategories />
        <FormAdd />
        <FormUpdate open={openUpdateDialog} setOpen={setOpenUpdateDialog} updateTarget={updateTarget} />
      </div>

      <div className="grid gap-6 sm:grid-cols-2 md:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5">
        {menus?.map((menu) => <MenuCard key={menu.id} menu={menu} onClick={() => triggerUpdateDialog(menu)} />)}
      </div>
    </section>
  );
}
