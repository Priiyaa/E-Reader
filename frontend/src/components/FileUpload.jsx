import React, { useState } from 'react';
import { useCurrentUser } from './useCurrentUser'; // ✅ adjust path if needed

function FileUpload() {
  const [file, setFile] = useState(null);
  const [isUploading, setIsUploading] = useState(false);
  const [uploadError, setUploadError] = useState('');
  const [uploadSuccess, setUploadSuccess] = useState('');

  const { currentUser } = useCurrentUser();

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
    setUploadError('');
    setUploadSuccess('');
  };

  const handleUpload = async () => {
    if (!file) {
      setUploadError('Please select a file to upload.');
      return;
    }

    setIsUploading(true);
    // ✅ Fetch the current user before uploading

    console.log(currentUser);


    if (!currentUser) {

      setUploadError('User S3 folder not found.');
      setIsUploading(false);
      return;
    }

    const formData = new FormData();
    formData.append('pdf', file);
    formData.append('s3_path', currentUser.user.s3Folder); // ✅ Use the S3 folder from user data
    console.log(currentUser.user.s3Folder+'Manoj');


    try {
      const response = await fetch('http://localhost:5003/upload', {
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
      setUploadError('An error occurred while uploading the file: ' + error.message);
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
