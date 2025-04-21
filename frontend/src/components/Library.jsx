// library code
import React, { useState, useEffect } from 'react';
import { useNavigate, Link } from "react-router-dom";
import FileUpload from './FileUpload';

function Library() {
  const navigate = useNavigate();
  const [books, setBooks] = useState([]);
  const DEFAULT_COVER = "/default-thumbnail.jpeg";

  useEffect(() => {
      const userId = localStorage.getItem("userId");
    fetch("http://localhost:5003/library?userId="+userId.toString()) // Calling aws backend.
      .then(response => response.json())
      .then(data => {
        console.log("API response:", data);
        if (data.books) {
          setBooks(data.books);
        }
      })
      .catch(error => console.error('Error fetching library:', error));
  }, []);

  return (
    <div style={{ fontFamily: "sans-serif", backgroundColor: "#f9f9f9", minHeight: "100vh" }}>
      {/* ✅ Navbar */}
      <nav
        style={{
          backgroundColor: "#6d4c41",
          padding: "1rem 2rem",
          color: "#fff",
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
        }}
      >
        <h1 style={{ margin: 0 }}>E-Reader Library</h1>
        <div style={{ display: "flex", gap: "1.5rem" }}>
          <Link to="/" style={navLinkStyle}>Home</Link>
          <Link to="/chat" style={navLinkStyle}>Messages</Link>
          <Link to="/about" style={navLinkStyle}>About</Link>
        </div>
      </nav>

      {/* ✅ Main Content */}
      <div style={{ padding: "2rem" }}>
        <FileUpload />

        <h2 style={{ margin: "2rem 0 1rem" }}>My Library</h2>
        <div style={{
          display: "flex",
          flexWrap: "wrap",
          gap: "20px",
          justifyContent: "start"
        }}>
          {books.map((book, index) => (
            <div
              key={index}
              style={{
                textAlign: "center",
                width: "160px",
                backgroundColor: "#fff",
                padding: "1rem",
                borderRadius: "8px",
                boxShadow: "0 4px 12px rgba(0, 0, 0, 0.1)",
              }}
            >
              <img
                src={DEFAULT_COVER}
                alt={book.name}
                style={{
                  width: "150px",
                  height: "200px",
                  objectFit: "cover",
                  borderRadius: "5px",
                  marginBottom: "0.75rem",
                }}
              />
              <p style={{ fontSize: "14px", fontWeight: "bold", marginBottom: "0.5rem" }}>
                {book.name}
              </p>
              <button
                onClick={() => navigate("/reader", { state: { url: book.url } })}
                style={{
                  padding: "6px 12px",
                  fontSize: "14px",
                  backgroundColor: "#007bff",
                  color: "white",
                  border: "none",
                  borderRadius: "4px",
                  cursor: "pointer",
                }}
              >
                View
              </button>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

const navLinkStyle = {
  color: "#fff",
  textDecoration: "none",
  fontWeight: "bold",
};

export default Library;