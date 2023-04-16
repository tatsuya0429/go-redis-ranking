import axios from "axios"


export const getTimeline = async (limit: number, offset: number) => {
  const res = await axios.get(`http://localhost:8080/timeline?limit=${limit}&offset=${offset}`)
  return res.data;
}