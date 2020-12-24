import React, { useEffect, useState } from 'react'
import { useParams, useHistory } from 'react-router-dom'
import styled from 'styled-components'
import requests from "../../requests"
import ReactMde from 'react-mde'
import Showdown from 'showdown'
import { toastNotification } from '../../utils/toasters'
import "react-mde/lib/styles/css/react-mde-all.css";
import { Container } from '../ArticlesList/ArticlesList';
import { Button, Title } from '../ArticleDetail/ArticleDetail'

const ButtonGroup = styled.div`
  display: flex;
  justify-content: center;
  margin-top: 1rem;
`

const converter = new Showdown.Converter({
  tables: true,
  simplifiedAutoLink: true,
  strikethrough: true,
  tasklists: true
});

const EditArticle: React.FC = () => {
  const { name } = useParams<{ name: string }>()
  const history = useHistory()
  const [isLoading, setIsLoading] = useState(false)
  const [content, setContent] = useState('')
  const [selectedTab, setSelectedTab] = React.useState<"write" | "preview">("write");

  async function getSingleArticle() {
    setIsLoading(true)
    const { data } = await requests.getSingleArticle(name)
    setContent(data)
    setIsLoading(false)
  }

  async function handlePublish(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault()
    requests.editArticle(name, content)
    .then(() => {
      toastNotification.success("Edit Successfully 😁")
      history.push("/")
    })
    .catch(() => {
      toastNotification.error("Failed to edit ❗")
    })
  }

  useEffect(() => {
    getSingleArticle()
  }, [])

  if (isLoading) {
    return <div>Loading....</div>
  }

  return (
    <Container>
      <form onSubmit={handlePublish}>
        <Title>{name}</Title>
        <ReactMde
          value={content}
          onChange={setContent}
          selectedTab={selectedTab}
          onTabChange={setSelectedTab}
          generateMarkdownPreview={markdown =>
            Promise.resolve(converter.makeHtml(markdown))
          }
          childProps={{
            writeButton: {
              tabIndex: -1
            }
          }}
        />
        <ButtonGroup>
          <Button type="submit" color="#38a169">Save</Button>
          <Button onClick={() =>  history.push("/" + name)} color="#ff4757">Cancel</Button>
        </ButtonGroup>
      </form>
    </Container>

  );
}

export default EditArticle