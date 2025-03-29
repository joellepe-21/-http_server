import React, { useState } from "react";
import { Link } from "react-router-dom";
import "../css/main.css";

const token = sessionStorage.getItem("token")

function UpdateArticle() {
  const [name, setName] = useState("");
  const [article, setText] = useState("");
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch("http://localhost:8000/api/update", {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify({ name, article }),
      });

      if (response.ok) {
        const data = await response.json();
        console.log("Статья успешно обновлена:", data);
        setSuccess("Статья успешно обновлена!");
        setName("");
        setText("");
      } else {
        setError("Ошибка при обновлении статьи. Пожалуйста, попробуйте снова.");
      }
    } catch (error) {
      setError("Ошибка при обновлении статьи. Пожалуйста, попробуйте снова.");
      console.error("Ошибка при обновлении статьи:", error);
    }
  };

  return (
    <div className="center-container">
      <h1>Update article</h1>
      {error && <div className="error">{error}</div>}
      {success && <div className="success">{success}</div>}
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="name">Article title:</label>
          <input
            type="text"
            id="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="text">Article text:</label>
          <textarea
            id="text"
            value={article}
            onChange={(e) => setText(e.target.value)}
            required
          />
        </div>
        <button type="submit" className="button">
          Update article
        </button>
      </form>
      <Link to="/admine" className="button">
        Return to admin panel
      </Link>
    </div>
  );
}

export default UpdateArticle;