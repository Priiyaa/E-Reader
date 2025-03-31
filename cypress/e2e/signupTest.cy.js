describe('Login Page Test', () => {
    it('should allow a user to log in', () => {
      // Visit the login page
      cy.visit('http://localhost:3000/signup');
  
      // Fill in the email and password fields
      cy.get('input[name="FullName"]').type('user3');
      cy.get('input[name="email"]').type('user3@gmail.com');
      cy.get('input[name="password"]').type('password123');
      cy.get('input[name="confirmPassword"]').type('password123');
  
      // Click the login button
      cy.get('button[type="submit"]').click();
  
      // Verify that the user is redirected to the dashboard (or another page)
      cy.url().should('include', '/dashboard');
    });
  });