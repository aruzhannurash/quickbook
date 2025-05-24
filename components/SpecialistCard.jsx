import React, { useState } from 'react';
import axios from 'axios';

const styles = {
  card: {
    border: '1px solid #ccc',
    padding: '20px',
    marginBottom: '15px',
    borderRadius: '8px',
    backgroundColor: '#f9f9f9',
  },
  name: {
    fontSize: '20px',
    fontWeight: 'bold',
    marginBottom: '10px',
  },
  input: {
    display: 'block',
    marginTop: '5px',
    marginBottom: '10px',
    padding: '5px',
    fontSize: '14px',
  },
  button: {
    padding: '8px 16px',
    fontSize: '14px',
    backgroundColor: '#007BFF',
    color: '#fff',
    border: 'none',
    borderRadius: '4px',
    cursor: 'pointer',
  },
  success: {
    color: 'green',
    marginTop: '10px',
  },
  error: {
    color: 'red',
    marginTop: '10px',
  },
};

function SpecialistCard({ specialist }) {
  const [appointmentTime, setAppointmentTime] = useState('');
  const [successMessage, setSuccessMessage] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  if (!specialist) {
    return (
      <div style={styles.card}>
        <p>Specialist data not loaded.</p>
      </div>
    );
  }

  const handleAppointment = async () => {
    setIsLoading(true);
    setErrorMessage('');
    setSuccessMessage('');

    try {
      if (!appointmentTime) {
        setErrorMessage('Choose time');
        return;
      }

      const date = new Date(appointmentTime);
      if (isNaN(date.getTime())) {
        setErrorMessage('Invalid date format');
        return;
      }

      if (date < new Date()) {
        setErrorMessage('Select a date and time in the future');
        return;
      }

      const token = localStorage.getItem('token');
      if (!token) {
        setErrorMessage('Authorization required!');
        return;
      }

      await axios.post(
        'http://localhost:8083/appointments',
        {
          specialist_id: specialist.id,
          datetime: date.toISOString(),
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
            'Content-Type': 'application/json',
          },
        }
      );

      setSuccessMessage('You have successfully booked!');
      setAppointmentTime('');
    } catch (error) {
      console.error('Error while booking:', error);
      setErrorMessage(
        error.response?.data?.message ||
        error.message ||
        'Error booking an appointment'
      );
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div style={styles.card}>
      <h3 style={styles.name}>{specialist.name || 'Name not specified'}</h3>
      <p>Specialty: {specialist.specialty || '-'}</p>
      <p>Experience: {specialist.experience || '0'} years</p>

      <label>
        Choose time:
        <input
          type="datetime-local"
          value={appointmentTime}
          onChange={(e) => setAppointmentTime(e.target.value)}
          style={styles.input}
          min={new Date().toISOString().slice(0, 16)}
        />
      </label>

      <button
        onClick={handleAppointment}
        style={styles.button}
        disabled={isLoading}
      >
        {isLoading ? 'Loading...' : 'Book'}
      </button>

      {successMessage && <p style={styles.success}>{successMessage}</p>}
      {errorMessage && <p style={styles.error}>{errorMessage}</p>}
    </div>
  );
}

export default SpecialistCard;