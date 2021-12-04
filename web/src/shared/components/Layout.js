import React from 'react'
import { Link } from 'react-router-dom'

import AuthStatus from './AuthStatus'

const Layout = () => {
  return (
    <div>
      <AuthStatus />

      <ul>
        <li>
          <Link to="/">Public pages</Link>
        </li>
        <li>
          <Link to="/profile">Protected pages</Link>
        </li>
      </ul>
      
    </div>
  )
}

export default Layout
