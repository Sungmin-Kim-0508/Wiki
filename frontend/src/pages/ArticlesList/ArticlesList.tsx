import React, { useEffect, useState } from 'react'
import requests from "../../requests"

type ArticlesListProps = {
}


const ArticlesList: React.FC<ArticlesListProps> = ({}) => {
  const [isLoading, setIsLoading] = useState(false)
  const [articles, setArticles] = useState(new Array())

  async function getAllArticles() {
    setIsLoading(true)
    const { data } = await requests.getAllArticles()
    setArticles(data)
    setIsLoading(false)
  }
  
  useEffect(() => {
    getAllArticles()
  }, [])

  if (isLoading) {
    return <div>Loading....</div>
  }

  if (!isLoading && articles.length === 0) {
    return <div>No results</div>
  }
  
  return (
    <div>
      <p>ArticlesList</p>
    </div>
  );
}

export default ArticlesList