import React from "react"
import { Routes, Route, BrowserRouter } from "react-router-dom"
import Home from "../Home"
import Login from "../Auth/Login"
import Profile from "../Profile"
import NavBar from "../shared/components/NavBar"
import PrivateRoute from "../shared/components/PrivateRoute"

function App() {
  return (
    <React.Fragment>
      <BrowserRouter>
        <Routes>
          {/* App */}
          <Route element={<NavBar />}>
            <Route path="/" element={<Home />}/>
            <Route 
              path="/profile"
              element={
                <PrivateRoute>
                  <Profile />
                </PrivateRoute>
              }
            />
          </Route>

          {/* Auth */}
          <Route path="/login" element={<Login />} />
          <Route path="*" element={<div>No page found!</div>} />
        </Routes>
      </BrowserRouter>
    </React.Fragment>
  )
}

export default App
