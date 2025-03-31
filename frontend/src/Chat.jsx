// import { useEffect, useRef, useState } from "react";
// if (!localStorage.getItem("token")) {
//   localStorage.setItem("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI2N2FhODhjM2Q0MzUyODc1MzVjNzcwYjciLCJleHAiOjE3NTExODE4NjB9.i7FT-FP2qHTnpqP6A11a9gWp5JpbrmvD_8hBEMYEQgU");
//   localStorage.setItem("userId", "67aa88c3d435287535c770b7");
// }


// const Chat = () => {
  
  
//   const [token] = useState(localStorage.getItem("token"));
//   const [currentUserId] = useState(localStorage.getItem("userId"));
//   const [onlineUsers, setOnlineUsers] = useState([]);
//   const [users, setUsers] = useState([]);
//   const [selectedUser, setSelectedUser] = useState(null);
//   const [messages, setMessages] = useState([]);
//   const [text, setText] = useState("");

//   const socketRef = useRef(null);
//   if (!token || !currentUserId) {
//     return <div style={{ padding: "2rem" }}>Please login or set token manually in localStorage.</div>;
//   }

//   // ✅ WebSocket setup
//   useEffect(() => {
//     if (!currentUserId) return;

//     const socket = new WebSocket(`ws://localhost:5002/ws?userId=${currentUserId}`);
//     socketRef.current = socket;

//     socket.onmessage = (event) => {
//       const data = JSON.parse(event.data);

//       if (data.type === "getOnlineUsers") {
//         setOnlineUsers(data.users.filter((u) => u !== currentUserId));
//       }

//       if (data.type === "chatMessage") {
//         const incoming = data.data;
//         if (
//           [incoming.senderId, incoming.receiverId].includes(selectedUser?._id)
//         ) {
//           setMessages((prev) => [...prev, incoming]);
//         }
//       }
//     };

//     return () => socket.close();
//   }, [currentUserId, selectedUser]);

//   // ✅ Fetch sidebar users
//   useEffect(() => {
//     const fetchUsers = async () => {
//       const res = await fetch("http://localhost:5002/api/messages/users", {
//         headers: { Authorization: `Bearer ${token}` },
//       });

//       if (!res.ok) {
//         const text = await res.text();
//         console.error("Failed to fetch users:", text);
//         return;
//       }

//       const data = await res.json();
//       setUsers(data);
//     };

//     if (token) fetchUsers();
//   }, [token]);

//   // ✅ Fetch chat history
//   useEffect(() => {
//     if (!selectedUser) return;

//     const fetchMessages = async () => {
//       const res = await fetch(`http://localhost:5002/api/messages/${selectedUser.id}`, {
//         headers: { Authorization: `Bearer ${token}` },
//       });

//       if (!res.ok) {
//         const text = await res.text();
//         console.error("Failed to fetch messages:", text);
//         return;
//       }

//       const data = await res.json();
//       setMessages(data);
//     };

//     fetchMessages();
//   }, [selectedUser]);

//   // ✅ Send new message
//   const sendMessage = async () => {
//     if (!text.trim() || !selectedUser) return;
//     if (!selectedUser || !selectedUser.id) {
//       console.warn("❌ Cannot send message: No user selected.");
//       return;
//     }
    

//     const body = { text, image: "" };

//     const res = await fetch(`http://localhost:5002/api/messages/send/${selectedUser.id}`, {
//       method: "POST",
//       headers: {
//         "Content-Type": "application/json",
//         Authorization: `Bearer ${token}`,
//       },
//       body: JSON.stringify(body),
//     });

//     if (!res.ok) {
//       const textErr = await res.text();
//       console.error("Failed to send message:", textErr);
//       return;
//     }

//     const newMsg = await res.json();
//     setMessages((prev = []) => [...prev, newMsg]);

//     setText("");

//     // 🔄 Send via WebSocket
//     // socketRef.current?.send(JSON.stringify({ type: "chatMessage", data: newMsg }));
//     console.log("Sending WS message to:", selectedUser.id);
//     console.log("currentUserId:", currentUserId);
//   console.log("selectedUser:", selectedUser);
//   console.log("selectedUser._id:", selectedUser?.id);


//     socketRef.current?.send(
//       JSON.stringify({
//         type: "chatMessage",
//         data: {
//           senderId: currentUserId,
//           receiverId: selectedUser.id, // ✅ include this!
//           text,
//           image: "",
//         },
//       })
//     );
    
//   };
//   return (
//     <div style={{ display: "flex", height: "100vh", backgroundColor: "#1e1e1e", color: "#fff" }}>
//       {/* Sidebar */}
//       <div style={{
//         width: "250px",
//         borderRight: "1px solid #444",
//         padding: "1rem",
//         overflowY: "auto",
//         backgroundColor: "#2b2b2b"
//       }}>
//         <h2 style={{ fontSize: "1.2rem", marginBottom: "1rem" }}>Users</h2>
//         <ul style={{ listStyle: "none", padding: 0 }}>
//           {users.map((user) => (
//             <li
//               key={user.id || user._id || user.email}
//               onClick={() => setSelectedUser(user)}
//               style={{
//                 padding: "0.5rem",
//                 marginBottom: "0.5rem",
//                 borderRadius: "8px",
//                 cursor: "pointer",
//                 backgroundColor: selectedUser?.id === user.id ? "#444" : "transparent",
//                 display: "flex",
//                 alignItems: "center",
//                 justifyContent: "space-between",
//                 fontWeight: "500",
//               }}
//             >
//               <span>{user.fullName || "Unknown"}</span>
//               {onlineUsers.includes(user.id || user._id) && (
//                 <span style={{ color: "#7CFC00", fontSize: "0.8rem" }}>●</span>
//               )}
//             </li>
//           ))}
//         </ul>
//       </div>
  
//       {/* Chat Panel */}
//       <div style={{ flex: 1, display: "flex", flexDirection: "column", padding: "1rem" }}>
//         {selectedUser ? (
//           <>
//             <div style={{
//               marginBottom: "1rem",
//               fontSize: "1.2rem",
//               fontWeight: "bold",
//               borderBottom: "1px solid #444",
//               paddingBottom: "0.5rem"
//             }}>
//               Chat with {selectedUser.fullName}
//             </div>
  
//             <div style={{
//               flex: 1,
//               overflowY: "auto",
//               paddingRight: "1rem",
//               marginBottom: "1rem",
//               display: "flex",
//               flexDirection: "column",
//               gap: "0.5rem"
//             }}>
//               {messages.map((msg, i) => (
//                 <div
//                   key={i}
//                   style={{
//                     alignSelf: msg.senderId === currentUserId ? "flex-end" : "flex-start",
//                     backgroundColor: msg.senderId === currentUserId ? "#a2f1a2" : "#f0f0f0",
//                     color: "#000",
//                     padding: "0.6rem 1rem",
//                     borderRadius: "20px",
//                     maxWidth: "70%",
//                     wordBreak: "break-word"
//                   }}
//                 >
//                   {msg.text}
//                 </div>
//               ))}
//             </div>
  
//             {/* Input */}
//             <div style={{ display: "flex", gap: "0.5rem" }}>
//               <input
//                 type="text"
//                 value={text}
//                 onChange={(e) => setText(e.target.value)}
//                 placeholder="Type a message..."
//                 style={{
//                   flex: 1,
//                   padding: "0.6rem 1rem",
//                   borderRadius: "20px",
//                   border: "1px solid #666",
//                   outline: "none",
//                   backgroundColor: "#333",
//                   color: "#fff"
//                 }}
//               />
//               <button
//                 onClick={sendMessage}
//                 disabled={!text.trim()}
//                 style={{
//                   padding: "0.6rem 1.2rem",
//                   borderRadius: "20px",
//                   backgroundColor: "#4caf50",
//                   color: "white",
//                   border: "none",
//                   cursor: "pointer",
//                   opacity: text.trim() ? 1 : 0.5
//                 }}
//               >
//                 Send
//               </button>
//             </div>
//           </>
//         ) : (
//           <div style={{ fontSize: "1rem", color: "#aaa" }}>
//             Select a user to chat with.
//           </div>
//         )}
//       </div>
//     </div>
//   );
  
// };

// export default Chat;



import { useEffect, useRef, useState } from "react";

const Chat = () => {
  const [token, setToken] = useState(localStorage.getItem("token"));
  const [currentUserId, setCurrentUserId] = useState(localStorage.getItem("userId"));
  const [users, setUsers] = useState([]);
  const [onlineUsers, setOnlineUsers] = useState([]);
  const [selectedUser, setSelectedUser] = useState(null);
  const [messages, setMessages] = useState([]);
  const [text, setText] = useState("");
  const socketRef = useRef(null);

  // 🔄 Watch for login via polling every 500ms
  useEffect(() => {
    const interval = setInterval(() => {
      const newToken = localStorage.getItem("token");
      const newUserId = localStorage.getItem("userId");

      if (newToken !== token || newUserId !== currentUserId) {
        setToken(newToken);
        setCurrentUserId(newUserId);
      }
    }, 500);

    return () => clearInterval(interval);
  }, [token, currentUserId]);

  // ✅ Connect WebSocket when user logs in
  useEffect(() => {
    if (!currentUserId) return;

    const socket = new WebSocket(`ws://localhost:5002/ws?userId=${currentUserId}`);
    socketRef.current = socket;

    socket.onmessage = (event) => {
      const data = JSON.parse(event.data);
      if (data.type === "getOnlineUsers") {
        const online = data.users.map((id) => String(id));
        setOnlineUsers(online.filter((id) => id !== currentUserId));
      }

      if (data.type === "chatMessage") {
        const incoming = data.data;
        if (
          incoming.senderId === selectedUser?.id ||
          incoming.receiverId === selectedUser?.id
        ) {
          setMessages((prev = []) => [...prev, incoming]);
        }
      }
    };

    return () => socket.close();
  }, [currentUserId, selectedUser]);

  // ✅ Fetch users when token becomes available
  useEffect(() => {
    if (!token) return;

    const fetchUsers = async () => {
      const res = await fetch("http://localhost:5002/api/messages/users", {
        headers: { Authorization: `Bearer ${token}` },
      });

      if (!res.ok) return;
      const data = await res.json();
      const normalized = data.map((u) => ({
        ...u,
        id: u.id || u._id,
      }));
      setUsers(normalized);
    };

    fetchUsers();
  }, [token]);

  // ✅ Fetch messages when user is selected
  useEffect(() => {
    if (!selectedUser || !token) return;

    const fetchMessages = async () => {
      const res = await fetch(`http://localhost:5002/api/messages/${selectedUser.id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });

      if (!res.ok) return;
      const data = await res.json();
      setMessages(data);
    };

    fetchMessages();
  }, [selectedUser, token]);

  
  // ✅ Send message
  const sendMessage = (e) => {
    e?.preventDefault();
    if (!text.trim() || !selectedUser) return;

    const messageData = {
      senderId: currentUserId,
      receiverId: selectedUser.id,
      text,
      image: "",
    };

    fetch(`http://localhost:5002/api/messages/send/${selectedUser.id}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(messageData),
    })
      .then((res) => res.json())
      .then((newMsg) => {
        setMessages((prev = []) => [...prev, newMsg]);
        setText("");

        socketRef.current?.send(
          JSON.stringify({ type: "chatMessage", data: newMsg })
        );
      });
  };

  if (!token || !currentUserId) {
    return <div style={{ padding: "2rem", color: "#fff" }}>Please log in first.</div>;
  }





  const filteredUsers = users.filter((user) => user.id !== currentUserId);
  console.log("currentUserId:", currentUserId);
  console.log("users:", users);
  console.log("filtered users:",filteredUsers)





  return (
 

    <div style={{ display: "flex", height: "100vh", backgroundColor: "#1e1e1e", color: "#fff" }}>
      {/* Sidebar */}
      <div style={{
        width: "250px",
        borderRight: "1px solid #444",
        padding: "1rem",
        overflowY: "auto",
        backgroundColor: "#2b2b2b"
      }}>
        <h2 style={{ fontSize: "1.2rem", marginBottom: "1rem" }}>Users</h2>
        <ul style={{ listStyle: "none", padding: 0 }}>
          
          {filteredUsers.map((user) => (
            <li
              key={user.id}
              onClick={() => setSelectedUser(user)}
              style={{
                padding: "0.5rem",
                marginBottom: "0.5rem",
                borderRadius: "8px",
                cursor: "pointer",
                backgroundColor: selectedUser?.id === user.id ? "#444" : "transparent",
                display: "flex",
                alignItems: "center",
                justifyContent: "space-between",
                fontWeight: "500",
              }}
            >
              <span>{user.fullName}</span>
              {onlineUsers.includes(user.id) && (
                <span style={{ color: "#7CFC00", fontSize: "0.8rem" }}>●</span>
              )}
            </li>
          ))}
        </ul>
      </div>

      {/* Chat Panel */}
      <div style={{ flex: 1, display: "flex", flexDirection: "column", padding: "1rem" }}>
        {selectedUser ? (
          <>
            <div style={{
              marginBottom: "1rem",
              fontSize: "1.2rem",
              fontWeight: "bold",
              borderBottom: "1px solid #444",
              paddingBottom: "0.5rem"
            }}>
              Chat with {selectedUser.fullName}
            </div>

            <div style={{
              flex: 1,
              overflowY: "auto",
              paddingRight: "1rem",
              marginBottom: "1rem",
              display: "flex",
              flexDirection: "column",
              gap: "0.5rem"
            }}>
              {Array.isArray(messages) && messages.map((msg, i) => {
                const isSentByMe = msg.senderId === currentUserId;

                return (
                  <div
                    key={i}
                    style={{
                      alignSelf: isSentByMe ? "flex-end" : "flex-start",
                      backgroundColor: isSentByMe ? "#a2f1a2" : "#f0f0f0",
                      color: "#000",
                      padding: "0.6rem 1rem",
                      borderRadius: "20px",
                      maxWidth: "70%",
                      wordBreak: "break-word"
                    }}
                  >
                    {msg.text}
                  </div>
                );
              })}
            </div>

            <form onSubmit={sendMessage} style={{ display: "flex", gap: "0.5rem" }}>
              <input
                type="text"
                value={text}
                onChange={(e) => setText(e.target.value)}
                placeholder="Type a message..."
                style={{
                  flex: 1,
                  padding: "0.6rem 1rem",
                  borderRadius: "20px",
                  border: "1px solid #666",
                  outline: "none",
                  backgroundColor: "#333",
                  color: "#fff"
                }}
              />
              <button
                type="submit"
                disabled={!text.trim()}
                style={{
                  padding: "0.6rem 1.2rem",
                  borderRadius: "20px",
                  backgroundColor: "#4caf50",
                  color: "white",
                  border: "none",
                  cursor: "pointer",
                  opacity: text.trim() ? 1 : 0.5
                }}
              >
                Send
              </button>
            </form>
          </>
        ) : (
          <div style={{ fontSize: "1rem", color: "#aaa" }}>
            Select a user to chat with.
          </div>
        )}
      </div>
    </div>
  );
};

export default Chat;
