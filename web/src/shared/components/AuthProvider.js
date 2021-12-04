import React from "react"
import { useSelector } from "react-redux"

let AuthContext = React.createContext()

export const AuthProvider = ({ children }) => {
  const user = useSelector((state) => state.users)
  console.log(`provider user: ${user.isAuthenticated}`)

  let signin = () => console.log("Signed in")

  let signout = () => console.log("Signed out")

  let value = { user, signin, signout }

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
}

export const useAuth = () => {
  return React.useContext(AuthContext)
}
