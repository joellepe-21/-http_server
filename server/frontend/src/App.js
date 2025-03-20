import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Home from "./components/home";
import Articles from "./components/articles";
import Register from "./components/register";
import Authorization from "./components/authorization";
import DeleteArticle from "./components/deleteArticle";
import AdminPanel from "./components/adminPanel"; 
import AddArticle from "./components/addArticle"; 
import UpdateArticle from "./components/updateArticle";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/articles" element={<Articles />} />
        <Route path="/register" element={<Register />} /> 
        <Route path="/authorization" element={<Authorization />} />
        <Route path="/admine" element={<AdminPanel />} />
        <Route path="/add" element={<AddArticle />} /> 
        <Route path="/delete" element={<DeleteArticle />} /> 
        <Route path="/update" element={<UpdateArticle />} /> 
      </Routes>
    </Router>
  );
}

export default App;