import React from 'react'
import { useDispatch } from 'react-redux'
import { useNavigate } from "react-router-dom"
import { naiveLogin } from '../../shared/redux/usersSlice'

const Login = () => {
  const dispatch = useDispatch()
  const navigate = useNavigate()

  const onSubmit = (event) => {
    event.preventDefault()
    dispatch(naiveLogin())
    navigate("/")
  }

  return (
    <div>
      <h1>Login page</h1>
      <form onSubmit={onSubmit}>
        <label>Username</label>
        <input type="text" />
        <button className="bg-blue-200 py-2 px-4 rounded-md" type="submit">Login</button>
      </form>
    </div>
  )
}

export default Login
