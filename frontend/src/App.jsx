// import Chat from './Chat'

// function App() {
//   return (
//     <div>
//       <h1>Go Chat App</h1>
//       <Chat />
//     </div>
//   );
// }

// export default App;

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Chat from "./Chat";
import Login from "./login";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/chat" element={<Chat />} />
        {/* You can add more routes here like <Route path="/signup" .../> */}
      </Routes>
    </Router>
  );
}

export default App;
