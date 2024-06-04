import PaymentMethodProvider from '@/context/PaymentMethodProvider';
import Page from '@/components/payment/Page';

export default function PaymentPage() {
  return (
    <PaymentMethodProvider>
      <Page />
    </PaymentMethodProvider>
  );
}
