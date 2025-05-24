import React from 'react'
import Navbar from '../components/Navbar'
import Footer from '../components/Footer'

function DoctorsPage() {
  return (
    <>
      <Navbar />
      <main style={{ padding: '2rem' }}>
        <h2>Список врачей</h2>
        {/* Здесь можно позже добавить компонент DoctorCard и фильтры */}
      </main>
      <Footer />
    </>
  )
}

export default DoctorsPage
