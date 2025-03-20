import { useEffect, useState } from "react";
import {Link} from "react-router-dom";
import "../css/main.css"; // Подключаем стили

function Articles() {
  const [articles, setArticles] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [pagination, setPagination] = useState({
    page: 1,
    limit: 3,
    totalRows: 0,
    totalPages: 1,
  });

  useEffect(() => {
    const fetchArticles = async () => {
      try {
        // Отправляем POST-запрос с параметрами page и limit
        const response = await fetch("http://localhost:8000/article", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            page: pagination.page,
            limit: pagination.limit,
          }),
        });

        if (!response.ok) {
          throw new Error(`Server responded with status ${response.status}`);
        }

        const data = await response.json();

        if (data && Array.isArray(data.Data)) {
          setArticles(data.Data); // Используйте data.Data вместо data.data
          setPagination((prev) => ({
              ...prev,
              totalRows: data.TotalRows,
              totalPages: data.TotalPage,
          }));
      } else {
          setError("Статьи не найдены в ответе сервера");
      }
      } catch (error) {
        console.error("Ошибка при загрузке статей:", error);
        setError("Ошибка при загрузке статей. Пожалуйста, попробуйте позже.");
      } finally {
        setLoading(false);
      }
    };

    fetchArticles();
  }, [pagination.page, pagination.limit]);

  const handlePageChange = (newPage) => {
    setPagination((prev) => ({ ...prev, page: newPage }));
  };

  if (loading) {
    return <div className="loading">Loading...</div>;
  }

  if (error) {
    return <div className="error">{error}</div>;
  }

  return (
    <div className="articles-page">
      <h1>Articles</h1>
      <div className="articles-container">
        {articles.length > 0 ? (
          articles.map((article, index) => (
            <div key={index} className="article-card">
              <h2>{article.name}</h2>
              <p>{article.article}</p>
            </div>
          ))
        ) : (
          <div>Статьи не найдены</div>
        )}
      </div>

      {/* Элементы управления пагинацией */}
      <div className="pagination-controls">
        <button
          onClick={() => handlePageChange(pagination.page - 1)}
          disabled={pagination.page === 1}
        >
          Previous
        </button>

        {Array.from({ length: pagination.totalPages }, (_, i) => i + 1).map((page) => (
          <button
            key={page}
            onClick={() => handlePageChange(page)}
            className={pagination.page === page ? "active" : ""}
          >
            {page}
          </button>
        ))}

        <button
          onClick={() => handlePageChange(pagination.page + 1)}
          disabled={pagination.page === pagination.totalPages}
        >
          Next
        </button>
      </div>

      <Link to="/" className="button">
        Вернуться на главную
      </Link>
    </div>
  );
}

export default Articles;