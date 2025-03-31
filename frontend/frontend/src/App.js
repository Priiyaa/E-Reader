// src/App.js
import React from 'react';
import { BrowserRouter as Router, Routes, Route, useLocation } from 'react-router-dom';
import Header from './components/Header';
import HomePage from './components/HomePage';
import AboutPage from './components/AboutPage';
import Footer from './components/Footer';
import Dashboard from "./components/Dashboard";

const AppContent = () => {
    const location = useLocation(); // Correct placement of useLocation()

    return (
        <div>
            {/* Conditionally render Header */}
            {location.pathname !== "/dashboard" && <Header />}

            {/* Define routes */}
            <Routes>
                <Route path="/" element={<HomePage />} />
                <Route path="/about" element={<AboutPage />} />
                <Route path="/dashboard" element={<Dashboard />} />
            </Routes>

            <Footer />
        </div>
    );
};

function App() {
    return (
        <Router>
            <AppContent /> {/* Wrapped content to use useLocation() correctly */}
        </Router>
    );
}

export default App;
