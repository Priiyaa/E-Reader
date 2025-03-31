// src/index.js
import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import { ThemeProvider, createTheme } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';

// Create a custom theme
const theme = createTheme({
  palette: {
    primary: {
      main: '#6d4c41', // A rich brown tone reminiscent of old books
    },
    secondary: {
      main: '#8d6e63', // A complementary, lighter brown shade
    },
    background: {
      default: '#fafafa',
    },
  },
  typography: {
    fontFamily: 'Georgia, serif', // A classic serif font for a literary feel
  },
});

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <CssBaseline /> {/* Provides a consistent baseline style */}
      <App />
    </ThemeProvider>
  </React.StrictMode>
);
