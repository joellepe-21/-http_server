import React, { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import "../css/main.css";

function Login() {
  const [login, setLogin] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch("http://localhost:8000/authorization", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ login, password }),
      });

      // Проверяем статус ответа
      if (response.ok) {
        const data = await response.json();
        const token = data.token;

        sessionStorage.setItem("token", token)

        console.log("Успешная авторизация:", data);
        navigate("/admine"); // Перенаправляем на главную страницу
      } else {
        setError("Ошибка при авторизации. Пожалуйста, проверьте логин и пароль.");
      }
    } catch (error) {
      setError("Ошибка при авторизации. Пожалуйста, попробуйте снова.");
      console.error("Ошибка при авторизации:", error);
    }
  };

  return (
    <div className="login-container">
      <h1>Authorization</h1>
      {error && <div className="error">{error}</div>}
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="login">Login:</label>
          <input
            type="text"
            id="login"
            value={login}
            onChange={(e) => setLogin(e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="password">Password:</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <button type="submit" className="button">
          Sign in
        </button>
      </form>
      <Link to="/" className="button">
      Return to home page
      </Link>
    </div>
  );
}

export default Login;