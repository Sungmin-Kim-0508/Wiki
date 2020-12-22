import { render, screen } from '@testing-library/react';
import ArticlesList from "../ArticlesList"
import React from 'react'

describe("ArticlesList Test", () => {
  it('renders <ArticlesList />', () =>  {
    const { getByText } = render(<ArticlesList />);
    getByText('ArticlesList')
  })
})
