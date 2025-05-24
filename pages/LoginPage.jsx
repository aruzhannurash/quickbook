import React, { useState } from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'
import { Link } from 'react-router-dom'


function LoginPage() {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const navigate = useNavigate()

  const handleSubmit = async (e) => {
    e.preventDefault()
    console.log("Отправка данных:", { username, password })
  
    try {
      const response = await axios.post('http://localhost:8081/login', { username, password })


      
  
      console.log("Успешный ответ:", response.data)
  
      localStorage.setItem('token', response.data.token)
      navigate('/home') 
  
    } catch (err) {
      console.error("Ошибка входа:", err.response ? err.response.data : err.message)
      setError('Invalid username or password')
    }
  }
  

  return (
    <div>
      <h1>Login</h1>
      <form onSubmit={handleSubmit}>
        <input 
          type="text" 
          placeholder="Username" 
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <input 
          type="password" 
          placeholder="Password" 
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <button type="submit">Login</button>
      </form>
      {error && <p style={{color: 'red'}}>{error}</p>}
      <p>
  Нет аккаунта? <Link to="/register">Зарегистрироваться</Link>
</p>
    </div>
    
  )
}

export default LoginPage
