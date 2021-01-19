import React, { useEffect, useState } from 'react'
import MarkdownEditor from 'react-markdown';
import styled from 'styled-components'
import { requests } from "../../requests"
import { useParams, useHistory } from 'react-router-dom'

const Header = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: space-between;
`

export const Title = styled.span`
  font-size: 3rem;
  font-weight: 700;
`

export const Button = styled.button<{ color: string; }>`
  color: ${props => props.color};
  border-width: 1px;
  border-color: ${props => props.color};
  background-color: white;
  border-radius: 5px;
  padding: 0.5rem 1rem;
  outline: none;

  margin-right: 1rem;

  :hover {
    background-color: ${props => props.color};
    color: white;
  }

  :focus {
    outline: none;
  }
`

const ArticleDetail: React.FC = () => {
  const { name } = useParams<{ name: string }>()
  const history = useHistory()
  const [isLoading, setIsLoading] = useState(false)
  const [content, setContent] = useState('')

  useEffect(() => {
    async function getSingleArticle() {
      setIsLoading(true)
      const { data } = await requests.getSingleArticle(name)
      setContent(data)
      setIsLoading(false)
    }

    getSingleArticle()
  }, [name])

  const MovePageButton : React.FC<{ path: string, label: string, btnColor: string }> = ({ path, label, btnColor }) => (
    <Button onClick={() => history.push(path)} color={btnColor}>{label}</Button>
  )

  if (isLoading) {
    return <div>Loading....</div>
  }
  
  return (
    <>
      <Header>
        <Title>{name}</Title>
        <div>
          <MovePageButton label="Edit" path={"/edit/" + name} btnColor="#38a169" />
          <MovePageButton label="Go back to the list" path="/" btnColor="#747d8c" />
        </div>
      </Header>
      <MarkdownEditor>{content.length > 0 ? content : "No article with this exact name found. Use Edit Button in the header to add it."}</MarkdownEditor>
    </>
  );
}

export default ArticleDetail