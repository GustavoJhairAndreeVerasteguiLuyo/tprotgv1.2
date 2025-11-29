import React from 'react';
import { CardElement, useStripe, useElements } from '@stripe/react-stripe-js';

interface Props {
  service: { id: string; name: string; price: number };
}

export default function CheckoutForm({ service }: Props) {
  const stripe = useStripe();
  const elements = useElements();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!stripe || !elements) return;

    alert(`Simulación de pago con tarjeta: ${service.name} - $${service.price}`);
  };

  return (
    <form onSubmit={handleSubmit}>
      <CardElement />
      <button type="submit" style={{ marginTop:10 }}>Pagar ${service.price}</button>
    </form>
  );
}
