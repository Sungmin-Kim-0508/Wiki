import React from 'react'
import { useParams } from 'react-router-dom'
import ReactMarkdown from 'react-markdown'

type ArticleDetailProps = {
}

const ArticleDetail: React.FC<ArticleDetailProps> = ({}) => {
  const params = useParams()
  console.log(params)
  return (
    <div>
      <ReactMarkdown># ArticleDetail</ReactMarkdown>
    </div>
  );
}

export default ArticleDetail