// src/components/AboutPage.js
import React from 'react';
import { Container, Typography, Box, Paper } from '@mui/material';

const AboutPage = () => {
  return (
    <Container maxWidth="lg" sx={{ py: 5 }}>
      {/* Hero Section with Gradient Background */}
      <Box
        sx={{
          position: 'relative',
          height: { xs: 250, md: 350 },
          background: 'linear-gradient(135deg,rgb(3, 3, 3) 0%,rgba(0, 0, 0, 0.13) 100%)',
          borderRadius: 2,
          mb: 5,
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          color: 'white',
          textAlign: 'center',
          px: 2,
        }}
      >
        <Box>
          <Typography variant="h3" component="h1" sx={{ fontWeight: 'bold', mb: 2 }}>
            About E-Reader
          </Typography>
          <Typography variant="h6">
            Your Digital Gateway to a World of Books
          </Typography>
        </Box>
      </Box>

      {/* Our Story Section */}
      <Paper elevation={2} sx={{ p: 4, mb: 5 }}>
        <Typography variant="h4" sx={{ fontWeight: 'bold', mb: 2, textAlign: 'center' }}>
          Our Story
        </Typography>
        <Typography variant="body1" paragraph>
          E-Reader was created to transform the way you experience literature in the digital age. By blending
          innovative technology with a deep passion for reading, we have built a platform that is both functional
          and inspiring.
        </Typography>
        <Typography variant="body1" paragraph>
          Our platform empowers you to manage your personal library, customize your reading experience, and engage
          with a vibrant community of fellow book enthusiasts—all at your fingertips.
        </Typography>
      </Paper>

      {/* Our Mission Section */}
      <Paper elevation={2} sx={{ p: 4 }}>
        <Typography variant="h4" sx={{ fontWeight: 'bold', mb: 2, textAlign: 'center' }}>
          Our Mission
        </Typography>
        <Typography variant="body1" paragraph>
          At E-Reader, we believe in the transformative power of storytelling. Our mission is to make digital
          reading accessible, engaging, and inspiring for everyone.
        </Typography>
        <Typography variant="body1">
          Join us on this journey to unlock a universe of knowledge, creativity, and endless adventures—one page at
          a time.
        </Typography>
      </Paper>
    </Container>
  );
};

export default AboutPage;
