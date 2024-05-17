import UsersProvider from '@/context/UserProvider';
import Page from './user/Page';

export default function UserPage() {
  return (
    <UsersProvider>
      <Page />
    </UsersProvider>
  );
}
