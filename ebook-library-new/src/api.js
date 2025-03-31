const API_BASE = "http://localhost:5000";  // Change if hosted

export const uploadFile = async (file) => {
    const formData = new FormData();
    formData.append("pdf", file);

    const response = await fetch(`${API_BASE}/`, {
        method: "POST",
        body: formData,
    });

    return response.json();
};

export const fetchLibrary = async () => {
    const response = await fetch(`${API_BASE}/library`);
    return response.json();
};
