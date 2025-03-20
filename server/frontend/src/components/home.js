import React from "react";
import { Link, useNavigate } from "react-router-dom"; // Добавляем useNavigate
import myImage from "../images/WalterHartwellWhite.png";

function Home() {
  const navigate = useNavigate(); // Хук для навигации

  return (
    <div className="home-container">
      {/* Кнопки в правом верхнем углу */}
      <div className="top-right-buttons">
        <button
          className="auth-button"
          onClick={() => navigate("/register")} // Переход на /register
        >
          Log up
        </button>
        <button
          className="auth-button"
          onClick={() => navigate("/authorization")} 
        >
          Log in
        </button>
      </div>

      <div className="welcome-section">
        <h1 className="welcome-text">Welcome to my lab</h1>
        <img src={myImage} alt="Walter White" className="home-image" />
      </div>
      <p className="home-description">Here I share my research and articles.</p>
      <Link to="/articles" className="button">
        Go to articles
      </Link>
    </div>
  );
}

export default Home;