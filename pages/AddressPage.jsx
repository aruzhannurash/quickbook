import React from 'react'
import Navbar from '../components/Navbar'
import Footer from '../components/Footer'

function AddressPage() {
  return (
    <>
      <Navbar />
      <main style={{ padding: '2rem' }}>
        <h2>Адреса и филиалы</h2>
        {/* Тут можно показать список адресов клиник */}
      </main>
      <Footer />
    </>
  )
}

export default AddressPage
