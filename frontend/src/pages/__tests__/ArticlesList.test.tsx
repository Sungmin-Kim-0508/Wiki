import { render, screen } from '@testing-library/react';
import ArticlesList from "../ArticlesList"
import React from 'react'

describe("ArticlesList Test", () => {
  it('renders <ArticlesList />', () =>  {
    render(<ArticlesList />);
    // getByText('ArticlesList')
  })
})
