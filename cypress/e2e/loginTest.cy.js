describe('Login Test', () => {
    it('should allow a user to log in', () => {
      // Visit the login page
      cy.visit('http://localhost:3000/login');
  
      // Fill in the email and password fields
      cy.get('input[name="email"]').type('priyaMittal@example.com');
      cy.get('input[name="password"]').type('password123');
  
      // Click the login button
      cy.get('button[type="submit"]').click();
  
      // Verify that the user is redirected to the dashboard
      cy.url().should('include', '/dashboard');
    });
  });
  