import React from 'react'
import ArticleDetail from '../pages/ArticleDetail'
import EditArticle from '../pages/EditArticle'
import ArticlesList from '../pages/ArticlesList'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'

function App() {
  return (
    <Router>
      <Switch>
        <Route path="/:name">
          <ArticleDetail />
        </Route>
        <Route path="/edit/:id">
          <EditArticle />
        </Route>
        <Route path="/">
          <ArticlesList />
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
