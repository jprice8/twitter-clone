import { createSlice, createAsyncThunk } from "@reduxjs/toolkit"

import api from "../utils/api"

const initialState = {
  status: "idle",
  error: null,
  isAuthenticated: false,
  id: "",
  email: "",
  password: "",
  name: "",
  created_at: "",
}

// fetch JWT from login
export const fetchJWT = createAsyncThunk(
  "users/fetchJWT",
  async ({ identity, password }) => {
    const formData = {
      identity: identity,
      password: password,
    }
    const response = await api({
      method: "POST",
      url: `/auth/login`,
      data: formData,
    })
    // Set token in localStorage
    const access_token = response.data.data
    localStorage.setItem("access_token", access_token)
    return response.data
  }
)

// Fetch user from JWT
export const fetchUserFromJWT = createAsyncThunk(
  "users/fetchUserFromJWT",
  async ({ userId }) => {
    const token = localStorage.getItem("access_token")

    const headers = {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    }

    const response = await api({
      method: "GET",
      url: `/user/${userId}`,
      headers,
    })
    return response.data
  }
)

export const usersSlice = createSlice({
  name: "users",
  initialState: initialState,
  reducers: {
    logUserOut(state, action) {
      const token = localStorage.getItem("access_token")
      if (token) {
        localStorage.removeItem("access_token")
        state.id = ""
        state.email = ""
        state.password = ""
        state.name = ""
        state.created_at = ""
      }
    },
    naiveLogin(state, action) {
      state.isAuthenticated = true
    }
  },
  extraReducers: {
    // Fetch JWT from form
    [fetchJWT.pending]: (state, action) => {
      state.status = "loading"
    },
    [fetchJWT.fulfilled]: (state, action) => {
      state.status = "succeeded"
      state.isAuthenticated = true
    },
    [fetchJWT.rejected]: (state, action) => {
      state.status = "failed"
      state.error = action.payload
      state.isAuthenticated = false
    },

    // Fetch user from access token
    [fetchUserFromJWT.pending]: (state, action) => {
      state.status = "loading"
    },
    [fetchUserFromJWT.fulfilled]: (state, action) => {
      state.status = "succeeded"
      state.id = action.payload.id
      state.email = action.payload.email
      state.password = action.payload.password
      state.name = action.payload.name
      state.created_at = action.payload.created_at

      state.isAuthenticated = true
    },
    [fetchUserFromJWT.rejected]: (state, action) => {
      state.status = "failed"
      state.error = action.payload
      state.isAuthenticated = false
    },
  },
})

export const { logUserOut, naiveLogin } = usersSlice.actions

export default usersSlice.reducer