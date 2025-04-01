# 🏆 Sprint 3 - GatorCan

## 📅 Duration: [03/04/2025] - [03/31/2025]

## Visual Demo Links
- [Sprint 3 Integrated Demo](https://drive.google.com/file/d/1D9-meydP8ja-mxD-ICXWkcSgtUs_fDTe/view?usp=drivesdk)
## API Documentation
- [Backend API Documentation](https://github.com/itswael/GatorCan/blob/main/BackendAPIDocumentation.md)

## 🎯 Goal
Implement the assignment submission and grading system, integrate AWS S3 for file storage, and develop a real-time course messaging feature. Ensure seamless backend-frontend integration with robust testing.

---

## 📌 User Stories & Assignments

### **🔹 Backend (Mohammad & Muthu)**

#### **1️⃣ Define Database Schema for Assignments & Submissions (Mohammad)**
- **Who:** Backend Developers
- **Why:** To store and manage assignment details and student submissions.
- **What:** Implement tables for assignments and submissions with necessary relationships.

#### **2️⃣ REST APIs for Fetching and Viewing Assignments (Mohammad)**
- **Who:** Students, Instructors
- **Why:** To allow students to access assignments and instructors to view submissions.
- **What:** Implement `GET /assignments` (available assignments) and `GET /assignments/submitted` (submitted assignments).

#### **3️⃣ File Upload for Assignment Submissions with AWS S3 (Muthu)**
- **Who:** Students
- **Why:** To enable students to submit assignments securely.
- **What:** Integrate AWS S3 for secure file uploads and implement `POST /assignments/upload` API.

#### **4️⃣ Grading API for Instructors (Muthu)**
- **Who:** Instructors
- **Why:** To allow instructors to grade submissions.
- **What:** Implement `POST /assignments/grade` API to update submission grades and feedback.

#### **✅ Unit Tests and Functional Tests for Backend (Mohammad & Muthu)**
##### **Unit Tests:**
###### Sprint 2 Unit Tests
- courseController -> Mocking courseService
    tests the controller functions like getCourse, getEnrolledCourse etc while integrated with the mocked service.
- userController -> Mocking userService
    tests the controller functions like getUserDetails, login, updateUserDetails etc while integrated with the mocked service.
- userService -> Mocking userRepository, courseRepository
    tests the user service functions for implemented business logic for the functions like getUserDetails, login, updateUserDetails etc while integrated with the mocked reporsitories.
- courseService -> Mocking userRepository, courseRepository
    tests the course service functions for implemented business logic for the functions like getCourse, getEnrolledCourse etc while integrated with the mocked reporsitories.
###### Sprint 3 Unit Tests
- AssignmentController -> Mocking assignmentService
  tests the controller functions like getAssignments, getAssignment etc while integrated with the mocked service.
- SubmissionController -> Mocking submissionService
  tests the controller functions like getsubmission, getSubmittedfiles etc while integrated with the mocked service.
- AssignmentService -> Mocking assignmentRepository, userRepository
  tests the user service functions for implemented business logic for the functions like getAssignments, getAssignment etc while integrated with the mocked reporsitories.
- SubmissionService -> Mocking submissionRepository, userRepository
  tests the user service functions for implemented business logic for the functions like getsubmission, getSubmittedfiles etc while integrated with the mocked reporsitories.

##### **Functional Tests:**
- Fetch assignments
- Submit assignments
- Grade assignments
- Role-based access testing
- Authentication & authorization testing

**Tested using positive, negative, and edge test cases.**
[Full Backend Test Log](https://github.com/itswael/GatorCan/blob/main/UnitTestLog.md)

---

### **🔹 Frontend (Navnit & Harsh)**

#### **5️⃣ Assignment Listing and Submission UI (Navnit)**
- **Who:** Students
- **Why:** To allow students to view and submit assignments.
- **What:** Implement a UI fetching data from `GET /assignments` and allowing file uploads.

#### **6️⃣ Real-Time Messaging UI for Each Course (Navnit)**
- **Who:** Students, Instructors
- **Why:** To enable real-time discussions within courses.
- **What:** Design a messaging UI for real-time chat within courses.

#### **7️⃣ WebSocket Integration for Real-Time Messaging (Harsh)**
- **Who:** Students, Instructors
- **Why:** To support real-time communication.
- **What:** Implement WebSocket connection for live messaging.

#### **8️⃣ Display Graded Assignments and Feedback (Harsh)**
- **Who:** Students
- **Why:** To allow students to see graded submissions.
- **What:** Implement UI to display assignment grades and instructor feedback.

#### **✅ Unit Tests and Cypress Test for Frontend (Navnit & Harsh)**
##### **Cypress Tests:**
- Fetch and validate assignments page
- Submit assignments and check file upload
- Fetch and validate real-time messages
- Test role-based UI components

##### **Unit Tests:**
- **Sprint 2 Tests:**
  - Login:
    - Check if username and password are rendered correctly
    - Check if we are able to change username and password correctly
    - Check if we get an error message on passing invalid credentials

  - AdminDashboard:
    - Check if add user tool renders correctly

  - StudentCourses:
    - Check if enrolled courses and all courses heading is rendered properly
    - Check if after fetch all courses API runs (mock), it loads the courses onto the courses tab
    - Check if "No enrolled courses" text renders if there are no courses enrolled by the student

  - StudentNavbar:
    - Check if all Navbar components are rendered such as Profile, Calendar, etc.

  - AdminDashboard:
    - Check if mock add user API gives correct success or failure responses

  - AuthService:
    - Check if mock login API gives correct success or failure response, and if local storage is updated with refresh token correctly

  - CourseService:
    - Check if fetch all courses API gives correct success or failure responses

  - UserNavigation:
    - Check if all elements are rendered correctly
    - Check if mock add user API gives correct display message on success or failure

---

- **Sprint 3 Tests:**
  - AdminDashboard:
    - Successfully adds a user
    - Successfully deletes a user
    - Successfully updates a user

  - StudentCourses:
    - Successfully fetches a course
    - Successfully fetches course assignments
    - Successfully fetches assignment details

  - AuthService:
    - Successfully gets user details
    - Successfully resets password

  - CourseService:
    - FetchCourse - Successfully fetches a course
    - FetchAssignments - Successfully fetches course assignments
    - FetchAssignmentDetails - Successfully fetches assignment details

  - UserNavigation:
    - Successfully adds a user
    - Successfully deletes a user
    - Successfully updates a user

  - ChatBox:
    - Should toggle chat window, send message, and scroll to the latest message

  - UserService:
    - Successfully gets user details
    - Successfully resets password

  - StudentProfile:
    - Displays user details after fetching
    - Shows loading indicator when fetching user details

  - CourseAssignments:
    - Should display 'Upcoming Assignments' and 'Past Assignments' text
    - Should fetch assignments data

  - CourseGrades:
    - Fetches assignments and displays total row

  - CourseHome:
    - Displays 'Home' text in the document

---

## ⚙️ **Sprint 3 - Issues & Completion Status**
### **Planned Issues:**
- Define and implement assignment database schema
- Develop REST APIs for assignments and submissions
- Implement AWS S3 file storage integration
- Build UI for assignment submissions and grading
- Implement real-time messaging with WebSocket(Localized, early stage)
- Conduct unit and functional testing

### **Successfully Completed:** ✅ All planned issues were completed.

---

## 🚀 Outcome
By the end of Sprint 3, we have:
- ✅ Database schema for assignments and submissions
- ✅ APIs for fetching, submitting, and grading assignments
- ✅ AWS S3 integration for file storage
- ✅ UI for assignment submissions and grading
- ✅ Real-time messaging via WebSocket (early stage)
- ✅ Comprehensive unit and functional tests for both backend and frontend

---

## 💚 Notes & Discussions
- [ ] Optimize WebSocket performance for high-traffic messaging
- [ ] Improve UI error handling for file uploads
- [ ] Enhance grading UI with filtering and sorting options
- [ ] Plan for next sprint (User Notifications & Advanced Course Analytics)

---

### 🔥 Sprint 3 Successfully Completed! 🚀

