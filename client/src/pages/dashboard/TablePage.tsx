import TablesProvider from '@/context/TableProvider';
import Page from '@/components/table/Page';

export default function TablePage() {
  return (
    <TablesProvider>
      <Page />
    </TablesProvider>
  );
}
