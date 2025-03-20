import React from "react";
import { Link } from "react-router-dom";
import "../css/main.css";

function AdminPanel() {
  return (
    <div className="admin-panel-container">
      <h1>Admin panel</h1>
      <div className="button-group">
        <Link to="/add" className="button">
          Add article
        </Link>
        <Link to="/update" className="button">
          Update article
        </Link>
        <Link to="/delete" className="button">
          Delete article
        </Link>
      </div>
      <Link to="/" className="button"> {/* Кнопка для возврата на главную */}
        Go to home page
      </Link>
    </div>
  );
}

export default AdminPanel;