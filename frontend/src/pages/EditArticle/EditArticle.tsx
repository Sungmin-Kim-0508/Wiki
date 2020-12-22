import React from 'react'
import { useParams } from 'react-router-dom'

type EditArticleProps = {
}

const EditArticle: React.FC<EditArticleProps> = ({}) => {
  const params = useParams()
  console.log(params)
  return (
    <div>
      <p>EditArticle</p>
    </div>
  );
}

export default EditArticle