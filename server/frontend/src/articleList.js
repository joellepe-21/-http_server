import React, {useEffect, useState} from "react";
import { getArticles } from "./api";

const ArticleList = () => {
    const [article, setArticle] = useState([])

    useEffect(() => {
        getArticles().then(setArticle);
    }, [])
    return (
        <div>
            <h2>Список статей</h2>
            <ul>
                {ArticleList.map((article, index) => (
                    <li key={index}>{article.name}:{article.articles}</li>
                ))}    
            </ul>
        </div>
    );
};

export default ArticleList