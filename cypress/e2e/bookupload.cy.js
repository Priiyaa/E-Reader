import 'cypress-file-upload';

describe("PDF Upload to S3", () => {
  beforeEach(() => {
    // Visit the Dashboard page before each test
    cy.visit("http://localhost:3000/dashboard");
  });

  it("should upload a PDF file to S3 and verify 200 status on PUT request", () => {
    // Intercept S3 PUT request and validate status
    cy.intercept("PUT", "https://books-uploaded.s3.us-east-2.amazonaws.com/**").as("s3Upload");

    // Attach a sample PDF file for upload
    cy.get("#pdf-upload").attachFile("AStudyinDrowningAvaReid.pdf");

    // Wait for the S3 upload request and validate 200 status
    cy.wait("@s3Upload", { timeout: 20000 }).then((interception) => {
      expect(interception).to.exist; // Ensure request was intercepted
      expect(interception.response.statusCode).to.eq(200); // Confirm 200 response
      cy.log("✅ S3 Upload request completed successfully");
    });

    // Verify UI shows upload success message
    cy.contains("File uploaded successfully").should("be.visible");
  });
});
