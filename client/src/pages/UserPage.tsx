import UsersProvider from '@/context/UserProvider';
import Page from '@/components/user/Page';

export default function UserPage() {
  return (
    <UsersProvider>
      <Page />
    </UsersProvider>
  );
}
