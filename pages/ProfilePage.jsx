import React from 'react'
import Navbar from '../components/Navbar'
import Footer from '../components/Footer'

function ProfilePage() {
  return (
    <>
      <Navbar />
      <main style={{ padding: '2rem' }}>
        <h2>Мой профиль</h2>
        {/* Здесь можно будет показать имя, email и кнопку редактирования */}
      </main>
      <Footer />
    </>
  )
}

export default ProfilePage
