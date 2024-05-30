import PaymentMethodsProvider from '@/context/PaymentMethodsProvider';
import Page from '@/components/payment/Page';

export default function PaymentPage() {
  return (
    <PaymentMethodsProvider>
      <Page />
    </PaymentMethodsProvider>
  );
}
