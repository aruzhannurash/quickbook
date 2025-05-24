import React from 'react'
import Navbar from '../components/Navbar'
import Footer from '../components/Footer'

function HomePage() {
  return (
    <>
      <Navbar />
      <main style={{ padding: '2rem', textAlign: 'center' }}>
        <h1
          style={{
            fontSize: '2.5rem',
            color: '#2c3e50',
            marginBottom: '2rem',
          }}
        >
          Quickbook – book a doctor in just a few clicks.
        </h1>

        <div style={{ display: 'flex', justifyContent: 'center', gap: '2rem', flexWrap: 'wrap' }}>
          <div
            style={{
              width: '220px',
              border: '1px solid #ccc',
              borderRadius: '10px',
              padding: '1rem',
              textAlign: 'center',
              boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
            }}
          >
            <img
              src="/images/doctor1.jpg"
              style={{
                width: '100%',
                height: '180px',
                objectFit: 'cover',
                borderRadius: '8px',
                marginBottom: '0.8rem',
              }}
            />
            <h3 style={{ fontSize: '1.2rem', marginBottom: '0.5rem' }}>Aidana Nurlanova</h3>
            <button
              style={{
                padding: '0.5rem 1rem',
                backgroundColor: '#007bff',
                color: '#fff',
                border: 'none',
                borderRadius: '5px',
                cursor: 'pointer',
              }}
            >
              Book
            </button>
          </div>

          <div
            style={{
              width: '220px',
              border: '1px solid #ccc',
              borderRadius: '10px',
              padding: '1rem',
              textAlign: 'center',
              boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
            }}
          >
            <img
              src="/images/doctor2.jpg"
              style={{
                width: '100%',
                height: '180px',
                objectFit: 'cover',
                borderRadius: '8px',
                marginBottom: '0.8rem',
              }}
            />
            <h3 style={{ fontSize: '1.2rem', marginBottom: '0.5rem' }}>Yerasyl Tulegenov</h3>
            <button
              style={{
                padding: '0.5rem 1rem',
                backgroundColor: '#007bff',
                color: '#fff',
                border: 'none',
                borderRadius: '5px',
                cursor: 'pointer',
              }}
            >
              Book
            </button>
          </div>

          <div
            style={{
              width: '220px',
              border: '1px solid #ccc',
              borderRadius: '10px',
              padding: '1rem',
              textAlign: 'center',
              boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
            }}
          >
            <img
              src="/images/doctor3.jpg"
              style={{
                width: '100%',
                height: '180px',
                objectFit: 'cover',
                borderRadius: '8px',
                marginBottom: '0.8rem',
              }}
            />
            <h3 style={{ fontSize: '1.2rem', marginBottom: '0.5rem' }}>Zhanel Ospanova</h3>
            <button
              style={{
                padding: '0.5rem 1rem',
                backgroundColor: '#007bff',
                color: '#fff',
                border: 'none',
                borderRadius: '5px',
                cursor: 'pointer',
              }}
            >
              Book
            </button>
          </div>
        </div>
      </main>
      <Footer />
    </>
  )
}

export default HomePage
