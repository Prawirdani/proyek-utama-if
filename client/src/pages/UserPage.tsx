import UserProvider from '@/context/UserProvider';
import Page from '@/components/user/Page';

export default function UserPage() {
  return (
    <UserProvider>
      <Page />
    </UserProvider>
  );
}
