import React, {useEffect, useState} from 'react'

export default function App(){
  const [events, setEvents] = useState<any[]>([])
  useEffect(()=>{
    fetch('/api/events').then(r=>r.json()).then(setEvents).catch(()=>{})
  },[])
  return (
    <div style={{padding:20}}>
      <h1>Perimeter Secure - Dashboard</h1>
      <ul>
        {events.map(e=> <li key={e.id}>{e.site} - {e.type} - {e.created_at}</li>)}
      </ul>
      <p>Frontend placeholder. Connect with backend at /api.</p>
    </div>
  )
}
