import React, { useState } from 'react';

function FileUpload() {
  const [file, setFile] = useState(null);
  const [isUploading, setIsUploading] = useState(false);
  const [uploadError, setUploadError] = useState('');
  const [uploadSuccess, setUploadSuccess] = useState('');

  // Handle file selection
  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
    setUploadError('');
    setUploadSuccess('');
  };

  // Handle file upload
  const handleUpload = async () => {
    if (!file) {
      setUploadError('Please select a file to upload.');
      return;
    }

    setIsUploading(true);

    const formData = new FormData();
    formData.append('pdf', file);

    try {
      const response = await fetch('http://localhost:5000/upload', {
        method: 'POST',
        body: formData,
      });

      const data = await response.json();

      if (response.ok) {
        setUploadSuccess(`File uploaded successfully: ${data.pdf_name}`);
      } else {
        setUploadError(data.error || 'Upload failed');
      }
    } catch (error) {
      setUploadError('An error occurred while uploading the file.');
    } finally {
      setIsUploading(false);
    }
  };

  return (
    <div>
      <h2>Upload PDF</h2>
      <input type="file" onChange={handleFileChange} />
      <button onClick={handleUpload} disabled={isUploading}>
        {isUploading ? 'Uploading...' : 'Upload'}
      </button>

      {uploadError && <div style={{ color: 'red' }}>{uploadError}</div>}
      {uploadSuccess && <div style={{ color: 'green' }}>{uploadSuccess}</div>}
    </div>
  );
}

export default FileUpload;
