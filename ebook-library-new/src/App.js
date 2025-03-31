import React, { useState } from "react";
import FileUpload from "./components/FileUpload";
import Library from "./components/Library";

function App() {
    const [uploadedFile, setUploadedFile] = useState(null);

    return (
        <div>
            <h1>eBook Library</h1>
            <FileUpload onUploadSuccess={setUploadedFile} />
            <hr />
            <Library />
        </div>
    );
}

export default App;
