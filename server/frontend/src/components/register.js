import React, { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import "../css/main.css";

function Register() {
  const [login, setLogin] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      // Отправляем POST-запрос на сервер
      const response = await fetch("http://localhost:8000/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ login, password }),
      });

      // Проверяем статус ответа
      if (response.ok) {
        const data = await response.json();
        console.log("Успешная регистрация:", data);
        navigate("/"); // Перенаправляем на главную страницу
      } else {
        setError("Ошибка при регистрации. Пожалуйста, попробуйте снова.");
      }
    } catch (error) {
      setError("Ошибка при регистрации. Пожалуйста, попробуйте снова.");
      console.error("Ошибка при регистрации:", error);
    }
  };

  return (
    <div className="register-container">
      <h1>Регистрация</h1>
      {error && <div className="error">{error}</div>}
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="login">Логин:</label>
          <input
            type="text"
            id="login"
            value={login}
            onChange={(e) => setLogin(e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="password">Пароль:</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <button type="submit" className="button">
          Зарегистрироваться
        </button>
      </form>
      <Link to="/" className="button">
        Вернуться на главную
      </Link>
    </div>
  );
}

export default Register;