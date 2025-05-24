import React, { useState } from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'

function RegisterPage() {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [error, setError] = useState('')
  const navigate = useNavigate()

  const handleSubmit = async (e) => {
    e.preventDefault()
    console.log("Отправка на /register:", { username, password })

    try {
      await axios.post('http://localhost:8081/register', {
        username,
        password,
        name: name || null,
        email: email || null,
      })
      navigate('/') 
    } catch (err) {
      setError(err.response?.data?.message || 'Ошибка регистрации')
    }
  }

  return (
    <div>
      <h1>Регистрация</h1>
      <form onSubmit={handleSubmit}>
        <input 
          type="text" 
          placeholder="Username" 
          value={username}
          onChange={e => setUsername(e.target.value)}
          required
        />
        <input 
          type="password" 
          placeholder="Password" 
          value={password}
          onChange={e => setPassword(e.target.value)}
          required
        />
        <input 
          type="text" 
          placeholder="Name" 
          value={name}
          onChange={e => setName(e.target.value)}
        />
        <input 
          type="email" 
          placeholder="Email" 
          value={email}
          onChange={e => setEmail(e.target.value)}
        />
        <button type="submit">Зарегистрироваться</button>
      </form>
      {error && <p style={{color: 'red'}}>{error}</p>}
    </div>
  )
}

export default RegisterPage
