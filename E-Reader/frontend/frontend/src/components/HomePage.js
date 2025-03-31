// src/components/HomePage.js
import React from 'react';
import { Container, Typography, Button, Box } from '@mui/material';

const HomePage = () => {
  return (
    <Box
      sx={{
        position: 'relative',
        // Adjust the minHeight to account for header and footer (example: 128px total)
        minHeight: 'calc(100vh - 128px)',
        // High-resolution background image from Unsplash (feel free to replace with your preferred image)
        backgroundImage: 'url(https://images.unsplash.com/photo-1524995997946-a1c2e315a42f?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80)',
        backgroundSize: 'cover',
        backgroundPosition: 'center',
      }}
    >
      {/* Dark overlay for improved text readability */}
      <Box
        sx={{
          position: 'absolute',
          top: 0,
          left: 0,
          right: 0,
          bottom: 0,
          backgroundColor: 'rgba(0, 0, 0, 0.6)',
        }}
      />

      <Container
        sx={{
          position: 'relative',
          zIndex: 1,
          height: '100%',
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'center',
          alignItems: 'center',
          textAlign: 'center',
          color: '#fff',
          px: 2,
        }}
      >
        <Typography variant="h2" component="h1" gutterBottom sx={{ fontWeight: 'bold' }}>
          Welcome to E-Reader
        </Typography>
        <Typography variant="h5" component="p" gutterBottom>
          Immerse yourself in the world of literature. Discover new stories, revisit classics, and share your passion for reading.
        </Typography>
        <Button variant="contained" color="secondary" size="large" sx={{ mt: 3 }}>
          Explore Now
        </Button>
      </Container>
    </Box>
  );
};

export default HomePage;
