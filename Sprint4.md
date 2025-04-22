# 🏆 Sprint 4 - GatorCan

## 📅 Duration: [01/04/2025] - [04/21/2025]

## Visual Demo Links
- [Sprint 4 Integrated Demo](https://drive.google.com/file/d/1nz1pBu1L1eCL7OnUPgYRo8zwy-cdHaDi/view?usp=sharing)
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

#### **5️⃣ Submit Assignment PDF File (Navnit)**
- **Who:** Students  
- **Why:** To allow students to submit their assignment work digitally and securely.  
- **What:** Implement a UI that enables students to upload PDF files for assignments, which are stored in an AWS S3 bucket via the backend API.

#### **6️⃣ Integrate AI-Based Course Recommendations (Navnit)**
- **Who:** Students  
- **Why:** To help students discover relevant courses that match their interests or performance trends.  
- **What:** Fetch personalized course recommendations from the backend via `GET /recommendations` and display them as interactive course cards along with the reasoning or tags behind the suggestion.

#### **7️⃣ Build Real-Time Chat (Harsh)**
- **Who:** Students, Instructors  
- **Why:** To allow students and instructors to communicate instantly within a course context.  
- **What:** Design and implement a real-time messaging feature using WebSocket or subscription-based APIs to enable instant course-related discussions.

#### **8️⃣ Build Instructor Dashboard (Harsh)**
- **Who:** Instructors  
- **Why:** To provide instructors with tools to manage course assignments and student submissions.  
- **What:** Create a dashboard that lets instructors create and edit assignments, view student submissions, and grade them with feedback.

### **✅ Unit Tests and Cypress Test for Frontend (Navnit & Harsh)**

##### **Cypress Tests:**
- Fetch and validate assignments page  
- Submit assignments and check file upload  
- Fetch and validate real-time messages  
- Test role-based UI components  

##### **Unit Tests:**

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

- **Sprint 4 Tests:**
  - InstructorDashboard.test.jsx:
    - Renders message when no courses are allocated
    - Renders course cards when courses are fetched
  - InstructorCourseAssignmentEdit.test.jsx:
    - Renders form fields with correct values
    - Shows not found message for missing assignment
  - InstructorService.test.js:
    - `fetchInstructorCourses` returns course list
    - `upsertAssignment` returns success true on success
    - `gradeAssignment` returns success false on error

---

## ⚙️ **Sprint 4 - Issues & Completion Status**
### **Planned Issues:**
- Submit Assignment PDF File  
- Integrate AI-Based Course Recommendations  
- Build Real-Time Chat  
- Build Instructor Dashboard  
- Created Instructor Dashboard for managing course content  
- Implemented real-time instructor-student Chat feature  
- Built file upload system for assignment submissions 

### **Successfully Completed:** ✅ All planned issues were completed.

---

## 🚀 Outcome
By the end of Sprint 4, we have:
- ✅ Instructor Dashboard to create, edit, and grade assignments
- ✅ Real-time instructor-student chat integrated with WebSocket
- ✅ File upload system for student assignment submissions using AWS S3
- ✅ AI-powered course recommendation engine with UI integration
- ✅ Optimized frontend with lazy loading and environment configs
- ✅ CI/CD pipeline with AWS CloudWatch monitoring
- ✅ Full coverage of unit and functional tests for new features

---

### 🔥 Sprint 4 Successfully Completed! 🚀

