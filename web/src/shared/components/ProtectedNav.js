import React from 'react'
import { Link } from 'react-router-dom'

const ProtectedNav = () => {
  return (
    <div>
      <h1>This is the protected navbar</h1>

      <nav>
        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
        </ul>
      </nav>
      
    </div>
  )
}

export default ProtectedNav
