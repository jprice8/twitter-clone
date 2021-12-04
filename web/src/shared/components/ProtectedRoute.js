import React from "react"
import { useSelector } from "react-redux"
import { Navigate, Route } from "react-router-dom"

const ProtectedRoute = ({ component: Component, ...restOfProps }) => {
  const isAuthed = useSelector((state) => state.users.isAuthenticated)
  console.log(isAuthed)

  // if (!isAuthed) {
    // return <Navigate to="/login" />
  // }

  return (
    <Route {...restOfProps} 
      render={( props ) => 
        isAuthed ? <Component {...props} /> : <Navigate to="/login" />
      }
    />
  )
}

export default ProtectedRoute
