import Page from '@/components/menu/Page';
import MenuProvider from '@/context/MenuProvider';

export default function MenuPage() {
  return (
    <MenuProvider>
      <Page />
    </MenuProvider>
  );
}
