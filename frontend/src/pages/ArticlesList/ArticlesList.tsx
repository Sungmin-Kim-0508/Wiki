import React, { useEffect, useState } from 'react'
import styled from 'styled-components'
import MarkdownEditor from 'react-markdown';
import requests from "../../requests"
import { v4 } from "uuid"
import { Link } from 'react-router-dom'

export const Container = styled.div`
  padding: 5rem 12rem;
`

const Box = styled.div`
  border-bottom: 1px solid #d8d8d8;
`

const Anchor = styled(Link)`
  text-decoration: none;
`

const ArticlesList: React.FC = () => {
  const [isLoading, setIsLoading] = useState(false)
  const [articles, setArticles] = useState<any[]>([])

  async function getAllArticles() {
    setIsLoading(true)
    const { data } = await requests.getAllArticles()
    setArticles(data)
    setIsLoading(false)
  }
  
  useEffect(() => {
    getAllArticles()
  }, [])

  // if (isLoading) {
  //   return <div>Loading....</div>
  // }

  if (!isLoading && articles.length === 0) {
    return <div>No results</div>
  }
  return (
    <Container>
      {articles.map(article => (
        <Article key={v4()} {...article} />
      ))}
    </Container>
  );
}

const Article : React.FC<{ Name: string; Content: string }> = ({ Name, Content }) => {
  return(
    <Box data-testid="resolved">
      <Anchor to={'/' + Name}>
        <h1>{Name}</h1>
        <MarkdownEditor>{Content.length > 0 ? Content : "No content exists"}</MarkdownEditor>
      </Anchor>
    </Box>
  )
}

export default ArticlesList