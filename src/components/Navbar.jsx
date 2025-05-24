import React from 'react';
import { Link } from 'react-router-dom';

function Navbar() {
  const linkStyle = {
    color: '#2c3e50',
    textDecoration: 'none',
    fontWeight: '500',
  };

  return (
    <nav style={{
      display: 'flex',
      fontSize: '18px',
      justifyContent: 'space-between',
      alignItems: 'center',
      padding: '1rem',
      background: '#f5f5f5',
      borderBottom: '2px solid #ccc'
    }}>
      <div style={{ flex: 1 }}></div> {/* Пустой блок для выравнивания */}
      
      <div style={{ display: 'flex', gap: '1rem' }}>
        <Link to="/home" style={linkStyle}>Home</Link>
        <Link to="/appointments" style={linkStyle}>My appointments</Link>
        <Link to="/profile" style={linkStyle}>My profile</Link>
        <Link to="/specialists" style={linkStyle}>Specialists</Link>
        <Link to="/address" style={linkStyle}>Address</Link>
      </div>

      <div style={{ flex: 1, display: 'flex', justifyContent: 'flex-end' }}>
        <Link to="/" style={linkStyle}>Log out</Link>
      </div>
    </nav>
  );
}

export default Navbar;
