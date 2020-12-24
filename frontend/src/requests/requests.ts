import axios, { AxiosResponse } from 'axios'

const baseURL = process.env.REACT_APP_DEV_ENV === "development" ?  "http://localhost:9090" : "http://localhost:8080/api"

const requestArticleURL = `${baseURL}/articles`
const requests = {
  getAllArticles: async () => {
    return await axios.get<any, AxiosResponse<[{}]>>(baseURL + "/articles")
  },
  getSingleArticle: async (name: string) => {
    const requestURL = `${requestArticleURL}/${name}`
    return await axios.get<any, AxiosResponse>(requestURL)
  },
  editArticle: async (name: string, content: string) => {
    const requestURL = `${requestArticleURL}/${name}`
    return await axios.put<any, AxiosResponse>(requestURL, content, {
      headers: {
        "Content-Type": "text/html"
      }
    })
  }
}

export default requests