import TableProvider from '@/context/TableProvider';
import Page from '@/components/table/Page';

export default function TablePage() {
  return (
    <TableProvider>
      <Page />
    </TableProvider>
  );
}
