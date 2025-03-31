import React, { useState, useEffect } from 'react';

import { fetchLibrary } from '../api';
function Library() {
    const [files, setFiles] = useState([]);

 

    useEffect(() => {
        fetchLibrary()
            .then(data => {
                console.log("API response:", data);
                setFiles(data.files || []);
            })
            .catch(error => console.error('Error fetching library:', error));
    }, []);

    
    
    return (
        <div>
            <h1>My Library</h1>
            <ul>
                {files.map((file, index) => (
                    <li key={index}>
                        <a href={file} target="_blank" rel="noopener noreferrer">{file}</a>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default Library;
