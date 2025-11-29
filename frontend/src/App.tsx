import React, { useState } from 'react';
import { loadStripe } from '@stripe/stripe-js'; 
import { Elements } from '@stripe/react-stripe-js';
import CheckoutForm from './components/CheckoutForm';


interface Servicio {
  id: string;
  name: string;
  price: number;
}

const serviciosDisponibles: Servicio[] = [
  { id: '1', name: 'Cámara de seguridad', price: 50 },
  { id: '2', name: 'Sensor de movimiento', price: 30 },
  { id: '3', name: 'Patrulla periódica', price: 100 },
];

const stripePromise = loadStripe('pk_test_TU_CLAVE_PUBLICA'); // reemplaza con tu clave pública Stripe

export default function App() {
  const [selectedService, setSelectedService] = useState<Servicio | null>(null);

  return (
    <div style={{ maxWidth: 800, margin: '0 auto' }}>
      <h1>Perimeter Secure - Pagos</h1>

      <section>
        <h2>Servicios disponibles</h2>
        <div style={{ display:'grid', gridTemplateColumns:'repeat(auto-fit,minmax(250px,1fr))', gap:20 }}>
          {serviciosDisponibles.map(s => (
            <div className="card" key={s.id}>
              <h3>{s.name}</h3>
              <p>Precio: ${s.price}</p>
              <button onClick={() => setSelectedService(s)}>Seleccionar</button>
            </div>
          ))}
        </div>
      </section>

      {selectedService && (
        <section style={{ marginTop:30 }}>
          <h2>Pagar por: {selectedService.name} (${selectedService.price})</h2>

          <div style={{ display:'flex', gap:10, flexWrap:'wrap' }}>
            <button onClick={() => alert(`Generando QR Yape para ${selectedService.name}`)}>Yape</button>
            <button onClick={() => alert(`Generando QR Plin para ${selectedService.name}`)}>Plin</button>
            <button onClick={() => alert(`Mostrando QR genérico para ${selectedService.name}`)}>QR</button>
          </div>

          <div style={{ marginTop:20, padding:20, background:'#e6f7ff', borderRadius:10 }}>
            <h3>Pagar con Tarjeta</h3>
            <Elements stripe={stripePromise}>
              <CheckoutForm service={selectedService} />
            </Elements>
          </div>
        </section>
      )}
    </div>
  );
}
