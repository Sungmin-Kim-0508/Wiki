import React from 'react'
import styled from 'styled-components'
import ArticleDetail from '../pages/ArticleDetail'
import EditArticle from '../pages/EditArticle'
import ArticlesList from '../pages/ArticlesList'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'

const Container = styled.div`
  padding: 5rem 12rem;
`

function App() {
  return (
    <Container>
      <Router>
        <Switch>
          <Route path="/edit/:name">
            <EditArticle />
          </Route>
          <Route path="/:name">
            <ArticleDetail />
          </Route>
          <Route path="/">
            <ArticlesList />
          </Route>
        </Switch>
      </Router>
    </Container>
  );
}

export default App;
