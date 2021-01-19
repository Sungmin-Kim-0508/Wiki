import React, { useEffect, useState } from 'react'
import styled from 'styled-components'
import MarkdownEditor from 'react-markdown';
import { requests } from "../../requests"
import { v4 } from "uuid"
import { Link } from 'react-router-dom'

const Box = styled.li`
  border-bottom: 1px solid #d8d8d8;
`

const Anchor = styled(Link)`
  text-decoration: none;
`

const ArticlesList: React.FC = () => {
  const [isLoading, setIsLoading] = useState(false)
  const [articles, setArticles] = useState<any[]>([])

  useEffect(() => {
    let ignore = false;
    async function getAllArticles() {
      try {
        setIsLoading(true)
        const { data } = await requests.getAllArticles()
        if (!ignore) setArticles(data)
      } catch (error) {
      } finally {
        setIsLoading(false)
      }
    }
    getAllArticles()
    return (() => { ignore = true; });
  }, [])

  if (isLoading && articles.length === 0) {
    return <div data-testid="loading">Loading....</div>
  }

  if (!isLoading && articles.length === 0) {
    return <div>No results</div>
  }
  return (
    <ul>
      {articles.map(article => (
        <Article key={v4()} {...article} />
      ))}
    </ul>
  );
}

const Article : React.FC<{ Name: string; Content: string }> = ({ Name, Content }) => {
  return(
    <Box>
      <Anchor to={'/' + Name}>
        <h1>{Name}</h1>
        <MarkdownEditor>{Content.length > 0 ? Content : "No content exists"}</MarkdownEditor>
      </Anchor>
    </Box>
  )
}

export default ArticlesList