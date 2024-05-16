import TablesProvider from '@/context/TableProvider';
import Page from './table/Page';

export default function TablePage() {
  return (
    <TablesProvider>
      <Page />
    </TablesProvider>
  );
}
