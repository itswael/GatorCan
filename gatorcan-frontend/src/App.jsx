import { Routes, Route } from "react-router-dom";

import Login from "./components/Login";
import ProtectedRoute from "./components/ProtectedRoute";
import ProtectedDashboard from "./components/ProtectedDashboard";
import AdminDashboard from "./components/Admin/AdminDashboard";
import StudentDashboard from "./components/Student/StudentDashboard";
import InstructorDashboard from "./components/Instructor/InstructorDashboard";
import StudentCalendar from "./components/Student/StudentCalendar";
import StudentInbox from "./components/Student/StudentInbox";
import StudentProfile from "./components/Student/StudentProfile";
import StudentCourses from "./components/Student/StudentCourses";
import AdminProfile from "./components/Admin/AdminProfile"
import CourseDetails from "./components/Student/Courses/CourseDetails";

import "./App.css";

function App() {
  return (
    <Routes>
      <Route path="login" element={<Login />} />
      <Route path="dashboard" element={<ProtectedRoute />} />

      {/* Protecting dashboard routes */}
      <Route
        path="admin-dashboard"
        element={
          <ProtectedDashboard allowedRoles={["admin"]}>
            <AdminDashboard />
          </ProtectedDashboard>
        }
      />
      <Route
        path="admin-profile"
        element={
          <ProtectedDashboard allowedRoles={["admin"]}>
            <AdminProfile />
          </ProtectedDashboard>
        }
      />
      <Route
        path="student-dashboard"
        element={
          <ProtectedDashboard allowedRoles={["student"]}>
            <StudentDashboard />
          </ProtectedDashboard>
        }
      />
      <Route
        path="instructor-dashboard"
        element={
          <ProtectedDashboard allowedRoles={["instructor"]}>
            <InstructorDashboard />
          </ProtectedDashboard>
        }
      />

      <Route
        path="student-profile"
        element={
          <ProtectedDashboard allowedRoles={["student"]}>
            <StudentProfile />
          </ProtectedDashboard>
        }
      />

      <Route
        path="student-calendar"
        element={
          <ProtectedDashboard allowedRoles={["student"]}>
            <StudentCalendar />
          </ProtectedDashboard>
        }
      />

      <Route
        path="student-inbox"
        element={
          <ProtectedDashboard allowedRoles={["student"]}>
            <StudentInbox />
          </ProtectedDashboard>
        }
      />

      <Route
        path="student-courses"
        element={
          <ProtectedDashboard allowedRoles={["student"]}>
            <StudentCourses />
          </ProtectedDashboard>
        }
      />

      <Route
        path="student-course/:id"
        element={
          <ProtectedDashboard allowedRoles={["student"]}>
            <CourseDetails />
          </ProtectedDashboard>
        }
      />

      <Route path="*" element={<Login />} />
    </Routes>
  );
}

export default App;
