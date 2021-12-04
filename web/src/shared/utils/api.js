import axios from "axios"

const host = (process.env.NODE_ENV === "production"
  ? "amazinclone.com"
  : "localhost")

export default axios.create({
  baseURL: `http://${host}/api`,
  timeout: 10000,
})
