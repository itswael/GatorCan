import { Routes, Route } from "react-router-dom";

import Login from "./components/Login";
import ProtectedRoute from "./components/ProtectedRoute";
import ProtectedDashboard from "./components/ProtectedDashboard";
import AdminDashboard from "./components/Admin/AdminDashboard";
import StudentDashboard from "./components/Student/StudentDashboard";
import InstructorDashboard from "./components/Instructor/InstructorDashboard";
import InstructorProfile from "./components/Instructor/InstructorProfile";
import InstructorChat from "./components/Instructor/InstructorChat";
import InstructorCourseHome from "./components/Instructor/Courses/InstructorCourseHome";
import InstructorCourseSyllabus from "./components/Instructor/Courses/InstructorCourseSyllabus";
import InstructorCourseAssignments from "./components/Instructor/Courses/Assignments/InstructorCourseAssignments";
import InstructorCourseAssignmentEdit from "./components/Instructor/Courses/Assignments/InstructorCourseAssignmentEdit";
import InstructorCourseAssignmentHome from "./components/Instructor/Courses/Assignments/InstructorCourseAssignmentHome";
import InstructorCourseAssignmentViewSubmissions from "./components/Instructor/Courses/Assignments/InstructorCourseAssignmentViewSubmissions";
import InstructorCourseAnnouncements from "./components/Instructor/Courses/InstructorCourseAnnouncements";
import InstructorCourseDiscussions from "./components/Instructor/Courses/InstructorCourseDiscussions";
import InstructorCourseStatistics from "./components/Instructor/Courses/InstructorCourseStatistics";
import StudentCalendar from "./components/Student/StudentCalendar";
import StudentInbox from "./components/Student/StudentInbox";
import StudentProfile from "./components/Student/StudentProfile";
import StudentCourses from "./components/Student/StudentCourses";
import AdminProfile from "./components/Admin/AdminProfile"
import CourseHome from "./components/Student/Courses/CourseHome";
import CourseAssignments from "./components/Student/Courses/CourseAssignments";
import CourseAnnouncements from "./components/Student/Courses/CourseAnnouncements";
import CourseDiscussions from "./components/Student/Courses/CourseDiscussions";
import CourseGrades from "./components/Student/Courses/CourseGrades";
import CourseChat from "./components/Student/Courses/CourseChat";
import CourseSyllabus from "./components/Student/Courses/CourseSyllabus";
import CourseAssignment from "./components/Student/Courses/CourseAssignment";
import Dummy from "./components/Dummy";
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
        path="instructor-profile"
        element={
          <ProtectedDashboard allowedRoles={["instructor"]}>
            <InstructorProfile />
          </ProtectedDashboard>
        }
      />

      <Route
        path="instructor-chat"
        element={
          <ProtectedDashboard allowedRoles={["instructor"]}>
            <InstructorChat />
          </ProtectedDashboard>
        }
      />

      <Route
        path="instructor-courses/:id/home"
        element={
          <ProtectedDashboard allowedRoles={["instructor"]}>
            <InstructorCourseHome />
          </ProtectedDashboard>
        }
      />

      <Route
        path="instructor-courses/:id/syllabus/"
        element={
          <ProtectedDashboard allowedRoles={["instructor"]}>
            <InstructorCourseSyllabus />
          </ProtectedDashboard>
        }
      />

      <Route
        path="instructor-courses/:id/assignments/"
        element={
          <ProtectedDashboard allowedRoles={["instructor"]}>
            <InstructorCourseAssignments />
          </ProtectedDashboard>
        }
      />

      <Route
        path="instructor-courses/:id/assignments/edit-assignment"
        element={
          <ProtectedDashboard allowedRoles={["instructor"]}>
            <InstructorCourseAssignmentEdit />
          </ProtectedDashboard>
        }
      />

      <Route
        path="instructor-courses/:id/assignments/:assignment_id/edit-assignment"
        element={
          <ProtectedDashboard allowedRoles={["instructor"]}>
            <InstructorCourseAssignmentEdit />
          </ProtectedDashboard>
        }
      />

      <Route
        path="instructor-courses/:id/assignments/:assignment_id"
        element={
          <ProtectedDashboard allowedRoles={["instructor"]}>
            <InstructorCourseAssignmentHome />
          </ProtectedDashboard>
        }
      />

      <Route
        path="instructor-courses/:id/assignments/:assignment_id/view-submissions"
        element={
          <ProtectedDashboard allowedRoles={["instructor"]}>
            <InstructorCourseAssignmentViewSubmissions />
          </ProtectedDashboard>
        }
      />

      <Route
        path="instructor-courses/:id/announcements/"
        element={
          <ProtectedDashboard allowedRoles={["instructor"]}>
            <InstructorCourseAnnouncements />
          </ProtectedDashboard>
        }
      />

      <Route
        path="instructor-courses/:id/discussions/"
        element={
          <ProtectedDashboard allowedRoles={["instructor"]}>
            <InstructorCourseDiscussions />
          </ProtectedDashboard>
        }
      />

      <Route
        path="instructor-courses/:id/statistics/"
        element={
          <ProtectedDashboard allowedRoles={["instructor"]}>
            <InstructorCourseStatistics />
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
            <CourseHome />
          </ProtectedDashboard>
        }
      />

      <Route
        path="student-course/:id/assignments/"
        element={
          <ProtectedDashboard allowedRoles={["student"]}>
            <CourseAssignments />
          </ProtectedDashboard>
        }
      />

      <Route
        path="student-courses/:id/assignments/:assignment_id"
        element={
          <ProtectedDashboard allowedRoles={["student"]}>
            <CourseAssignment />
          </ProtectedDashboard>
        }
      />

      <Route
        path="student-course/:id/announcements/"
        element={
          <ProtectedDashboard allowedRoles={["student"]}>
            <CourseAnnouncements />
          </ProtectedDashboard>
        }
      />

      <Route
        path="student-course/:id/grades/"
        element={
          <ProtectedDashboard allowedRoles={["student"]}>
            <CourseGrades />
          </ProtectedDashboard>
        }
      />

      <Route
        path="student-course/:id/syllabus/"
        element={
          <ProtectedDashboard allowedRoles={["student"]}>
            <CourseSyllabus />
          </ProtectedDashboard>
        }
      />

      <Route
        path="student-course/:id/discussions/"
        element={
          <ProtectedDashboard allowedRoles={["student"]}>
            <CourseDiscussions />
          </ProtectedDashboard>
        }
      />

      <Route
        path="student-course/:id/chat/"
        element={
          <ProtectedDashboard allowedRoles={["student"]}>
            <CourseChat />
          </ProtectedDashboard>
        }
      />

      <Route
        path="dummy"
        element={
          <ProtectedDashboard allowedRoles={["student", "admin", "instructor"]}>
            <Dummy />
          </ProtectedDashboard>
        }
      ></Route>

      <Route path="*" element={<Login />} />
    </Routes>
  );
}

export default App;
