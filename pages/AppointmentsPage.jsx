import React, { useEffect, useState } from 'react'
import Navbar from '../components/Navbar'
import Footer from '../components/Footer'
import axios from 'axios'

function AppointmentsPage() {
  const [appointments, setAppointments] = useState([]) 
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    axios.get('http://localhost:8082/appointments') 
      .then(res => {
        setAppointments(res.data)
        setLoading(false)
      })
      .catch(err => {
        console.error('Ошибка при загрузке записей:', err)
        setLoading(false)
      })
  }, [])

  return (
    <>
      <Navbar />
      <main style={{ padding: '2rem' }}>
        <h2>Мои записи</h2>

        {loading ? (
          <p>Загрузка...</p>
        ) : appointments.length === 0 ? (
          <p>Нет записей.</p>
        ) : (
          <ul>
            {appointments.map((appointment) => (
              <li key={appointment.id}>
                Запись к специалисту ID: {appointment.specialist_id} <br />
                Дата и время: {appointment.datetime}
              </li>
            ))}
          </ul>
        )}
      </main>
      <Footer />
    </>
  )
}

export default AppointmentsPage
