import { Routes, Route } from "react-router-dom";
import Dashboard from "./components/Dashboard";
import Login from "./components/Login";
import Courses from "./components/Courses"
import Calendar from "./components/Calendar"
import Inbox from "./components/Inbox"
import ResetPassword from "./components/ResetPassword"
import Layout from "./components/Layout";
import Missing from "./components/Missing";
import LinkPage from "./components/LinkPage";
import Unauthorized from "./components/Unauthorized";
import PersistLogin from "./components/PersistLogin";
import RequireAuth from "./components/RequireAuth";
import Admin from "./components/Admin";
import Instructor from "./components/Instructor";
import UserRegistration from "./components/UserRegistration";

import "./App.css";

const ROLES = {
  Student: 1001,
  Admin: 9001,
  Instructor: 5001,
};

function App() {
  return (
    <Routes>
      <Route path="/" element={<Layout />}>
        {/* public routes */}
        <Route path="login" element={<Login />} />
        <Route path="linkpage" element={<LinkPage />} />
        <Route path="unauthorized" element={<Unauthorized />} />
        <Route path="/resetPassword" element={<ResetPassword />} />

        {/* we want to protect these routes */}
        <Route element={<PersistLogin />}>
          <Route element={<RequireAuth allowedRoles={[ROLES.Student]} />}>
            <Route path="/" element={<Dashboard />} />
            <Route path="/courses" element={<Courses />} />
            <Route path="/calendar" element={<Calendar />} />
            <Route path="/inbox" element={<Inbox />} />
          </Route>

          <Route element={<RequireAuth allowedRoles={[ROLES.Admin]} />}>
            <Route path="admin" element={<Admin />} />
            <Route path="userRegistration" element={<UserRegistration />} />
          </Route>

          <Route element={<RequireAuth allowedRoles={[ROLES.Instructor]} />}>
            <Route path="instructor" element={<Instructor />} />
          </Route>
        </Route>
      </Route>

      {/* catch all */}
      <Route path="*" element={<Missing />} />
    </Routes>
  );
}

export default App;
