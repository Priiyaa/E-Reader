// src/components/Footer.js
import React from 'react';
import { Box, Typography } from '@mui/material';

const Footer = () => {
  return (
    <Box
      component="footer"
      sx={{
        backgroundColor: '#6d4c41', // matches our theme's primary color (a rich brown)
        color: '#fff',
        py: 2,
        textAlign: 'center',
      }}
    >
      <Typography variant="body2">
        © {new Date().getFullYear()} E-Reader. All rights reserved.
      </Typography>
    </Box>
  );
};

export default Footer;
