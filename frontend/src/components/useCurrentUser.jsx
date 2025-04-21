import React, { useState, useEffect } from 'react';
export function useCurrentUser() {

  const [currentUser, setCurrentUser] = useState(null);

  useEffect(() => {
    const fetchCurrentUser = async () => {
      const userId = localStorage.getItem("userId");
      const token = localStorage.getItem("token");
      console.log(token + 'token');
      if (!userId || !token) return;

      try {
        const res = await fetch(`http://localhost:5002/api/auth/user/${userId}`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        const data = await res.json();

        setCurrentUser(data);
      } catch (err) {
        console.error("Failed to fetch current user:", err);
      }
    };

    fetchCurrentUser();
  }, []);

  return { currentUser };
}
