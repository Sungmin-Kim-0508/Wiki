import axios, { AxiosResponse } from 'axios'

const baseURL = "http://localhost:8080"

const requests = {
  getAllArticles: async () => {
    return await axios.get<any, AxiosResponse<[{}]>>(baseURL + "/api/articles")
  },
  getSingleArticle: async () => {
    fetch(baseURL + "/wiki", {
      method: 'GET',
      mode: 'cors'
    })
    .then(res => {
      console.log(res)
    })
  }
}

export default requests