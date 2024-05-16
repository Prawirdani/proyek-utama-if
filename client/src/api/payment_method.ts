export const fetchPaymentMethods = async () => {
  const res = await fetch('/api/v1/payments/methods', {
    method: 'GET',
    credentials: 'include',
  });

  if (!res.ok) {
    const errorBody = await res.json();
    throw new Error((errorBody as ErrorResponse).error.message);
  }

  const resBody = (await res.json()) as ApiResponse<MetodePembayaran[]>;
  return resBody.data;
};
