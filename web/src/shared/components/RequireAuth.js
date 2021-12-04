import React from 'react'
import { useSelector } from 'react-redux'
import { useLocation, Navigate } from 'react-router-dom'

const RequireAuth = ({ children }) => {
  const user = useSelector(state => state.users)
  let location = useLocation()

  if (!user.isAuthenticated) {
    return <Navigate to="/login" state={{ from: location }} />
  }

  return children
}

export default RequireAuth
