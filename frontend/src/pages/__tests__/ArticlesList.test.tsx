import React from 'react'
import axios from 'axios'
// import react-testing methods
import { render, act } from '@testing-library/react'
import ArticlesList from "../ArticlesList"
import { BrowserRouter as Router } from "react-router-dom";

jest.mock('axios')

const mockedAxios = axios as jest.Mocked<typeof axios>

const allArticles = [
  {
    Name: 'wiki',
    Content: 'A wiki is a knowledge base website'
  },
  {
    Name: 'rest_api',
    Content: ''
  }
]

const RouterWrappedComponent : React.FC = () => (
  <Router>
    <ArticlesList />
  </Router>
)

describe("ArticlesList Test", () => {
  it("should show 'No results' text, when articles don't exist", async () =>  {
    const promise = Promise.resolve()
    const { getByTestId, findByText } = render(<RouterWrappedComponent />);

    expect(getByTestId('loading')).toHaveTextContent("Loading....")

    await findByText(/no results/i)
    await act(() => promise)
  })

  it('should renders <ArticlesList /> when the articles exist', async () =>  {
    const promise = Promise.resolve()
    mockedAxios.get.mockResolvedValueOnce({ data: allArticles })
    const { getByTestId, findAllByRole } = render(<RouterWrappedComponent />);

    expect(getByTestId('loading')).toHaveTextContent("Loading....")

    const articleList = await findAllByRole('link')

    expect(articleList).toHaveLength(2)
    expect(articleList[0]).toHaveTextContent('wiki')
    expect(articleList[0]).toHaveTextContent('A wiki is a knowledge base website')
    await act(() => promise)
  })

  it('should display "No content exists" when there is no content for an article', async () => {
    const promise = Promise.resolve()
    mockedAxios.get.mockResolvedValueOnce({ data: allArticles })
    const { findAllByRole } = render(<RouterWrappedComponent />);

    const articleList = await findAllByRole('link')

    expect(articleList[1]).toHaveTextContent('rest_api')
    expect(articleList[1]).toHaveTextContent('No content exists')

    await act(() => promise)
  })
})
