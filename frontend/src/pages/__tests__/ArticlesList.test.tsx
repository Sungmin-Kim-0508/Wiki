import React from 'react'
import { render, waitFor } from '@testing-library/react';
import ArticlesList from "../ArticlesList"
import '@testing-library/jest-dom/extend-expect'


describe("ArticlesList Test", () => {

  it('renders <ArticlesList />', async () =>  {
    const { getByText } = render(<ArticlesList />);

    await waitFor(() => getByText('Loading....'))
  })
})
