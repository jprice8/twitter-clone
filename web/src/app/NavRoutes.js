import React from 'react'
import { BrowserRouter as Router, Routes, Route } from "react-router-dom"

import Home from '../Home'
import Login from "../Auth/Login"
import ProtectedRoute from '../shared/components/ProtectedRoute'
import Profile from "../Profile"

const NavRoutes = () => {
  return (
    <Router>
      <Routes>
        <Route exact path="/" component={Home} />
        <Route exact path="/login" component={Login} />
        <ProtectedRoute exact path="/profile" component={Profile} />
      </Routes>
    </Router>
  )
}

export default NavRoutes
