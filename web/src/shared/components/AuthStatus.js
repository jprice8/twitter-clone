import React from 'react'
import { useSelector } from 'react-redux'
import { useNavigate } from 'react-router-dom'

const AuthStatus = () => {
  let user = useSelector(state => state.users)
  let navigate = useNavigate()

  if (!user.isAuthenticated) {
    return <p>You are not logged in.</p>
  }

  return (
    <div>
      <h1>Welcome! Your are logged in.</h1>
    </div>
  )
}

export default AuthStatus
