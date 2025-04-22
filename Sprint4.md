# 🏆 Sprint 4 - GatorCan

## 📅 Duration: [04/01/2025] - [04/21/2025]

## Visual Demo Links
- [Sprint 4 Integrated Demo](https://drive.google.com/file/d/1LywboJhJ_LxYwHm6a4FTGKUGW-ZBsTMz/view?usp=sharing)

## API Documentation
- [Backend API Documentation](https://github.com/itswael/GatorCan/blob/main/BackendAPIDocumentation.md)

## 🎯 Goal
Implement notifications using AWS SNS, add AI-powered course recommendations and NLP-based lecture summarization (optional), and prepare the application for production deployment with CI/CD and AWS monitoring.

---

## 📌 User Stories & Assignments

### **🔹 Backend (Muthu & Mohammad)**

#### **1️⃣ SNS for Assignments & Deadlines Notifications (Muthu)**
- **Who:** Backend Developers
- **Why:** To inform students instantly about new assignments and upcoming deadlines.
- **What:** Integrate AWS SNS to send notifications when assignments are created or nearing deadline.

#### **2️⃣ SNS for Grading & Announcements Notifications (Muthu)**
- **Who:** Backend Developers
- **Why:** To notify students about assignment grading and important course announcements.
- **What:** Trigger SNS notifications on assignment grading and new course announcements.

#### **3️⃣ AI-Based Course Recommendations (Mohammad)**
- **Who:** Backend Developers
- **Why:** To enhance user experience through personalized course suggestions.
- **What:** Implement recommendation engine based on previous enrollments and user interests. Provide `GET /recommendations` API.

#### **4️⃣ NLP-Based Lecture Summarization (Mohammad)**
- **Who:** Backend Developers
- **Why:** To help students quickly review lecture content.
- **What:** Accept lecture transcript and return summarized version via `POST /summarize` API using basic NLP techniques.

#### **✅ Unit Tests and Functional Tests for Backend (Muthu & Mohammad)**
##### **Unit Tests:**
###### Sprint 3 Unit Tests
- courseController -> Mocking courseService
    tests the controller functions like getCourse, getEnrolledCourse etc while integrated with the mocked service.
- userController -> Mocking userService
    tests the controller functions like getUserDetails, login, updateUserDetails etc while integrated with the mocked service.
- userService -> Mocking userRepository, courseRepository
    tests the user service functions for implemented business logic for the functions like getUserDetails, login, updateUserDetails etc while integrated with the mocked reporsitories.
- courseService -> Mocking userRepository, courseRepository
    tests the course service functions for implemented business logic for the functions like getCourse, getEnrolledCourse etc while integrated with the mocked reporsitories.
- AssignmentController -> Mocking assignmentService
    tests the controller functions like getAssignments, getAssignment etc while integrated with the mocked service.
- SubmissionController -> Mocking submissionService
    tests the controller functions like getsubmission, getSubmittedfiles etc while integrated with the mocked service.
- AssignmentService -> Mocking assignmentRepository, userRepository
    tests the user service functions for implemented business logic for the functions like getAssignments, getAssignment etc while integrated with the mocked reporsitories.
- SubmissionService -> Mocking submissionRepository, userRepository
    tests the user service functions for implemented business logic for the functions like getsubmission, getSubmittedfiles etc while integrated with the mocked reporsitories.

###### Sprint 4 Unit Tests
- NotificationService -> Mocking SNS Client
- RecommendationService -> Mocking Course Repository, tests the recommendation functionality to get the recommended course  based on different user course history.
- RecommendationService Controller -> Mocking RecommendationService, tests the recommendation functionality to get the recommended course based on different user course history.
- SummarizationService -> Mocking NLP summarization logic, to generate the summary of the text content passed in the body parameter with one third length of the original content.

##### **Functional Tests:**
- Fetch assignments
- Submit assignments
- Grade assignments
- Role-based access testing
- Authentication & authorization testing
- New assignment notification trigger
- Assignment graded notification trigger
- Course recommendation API response
- Lecture summarization API response

**Tested using positive, negative, and edge test cases.**
[Full Backend Test Log](https://github.com/itswael/GatorCan/blob/main/UnitTestLog.md)

---

### **🔹 Frontend (Navnit & Harsh)**

#### **5️⃣ Notification Center UI (Navnit)**
- **Who:** Students
- **Why:** To display real-time notifications for better user awareness.
- **What:** Build a Notification Center dashboard showing categorized SNS notifications.

#### **6️⃣ AI-Based Course Recommendations Integration (Navnit)**
- **Who:** Students
- **Why:** To suggest personalized courses.
- **What:** Fetch and display course recommendations from `GET /recommendations` API.

#### **7️⃣ Frontend Optimization for Production Deployment (Harsh)**
- **Who:** Frontend Developers
- **Why:** To ensure optimal app performance.
- **What:** Implement lazy loading, tree-shaking, and reduce bundle size for production readiness.

#### **8️⃣ CI/CD Pipeline & AWS Monitoring Integration (Harsh)**
- **Who:** Frontend & DevOps Developer
- **Why:** To automate deployment and monitor application performance.
- **What:** Set up CI/CD using GitHub Actions/AWS CodePipeline and integrate AWS CloudWatch for monitoring.

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
- Implement AWS SNS notifications
- Build AI-based course recommendations and summarization
- Develop Notification Center and integrate AI features in frontend
- Optimize frontend for production
- Set up CI/CD and monitoring
- Submit Assignment PDF File  
- Build Real-Time Chat  
- Build Instructor Dashboard  
- Created Instructor Dashboard for managing course content  
- Implemented real-time instructor-student Chat feature  
- Built file upload system for assignment submissions

### **Successfully Completed:** ✅ All planned issues were completed.

---

## 🚀 Outcome
By the end of Sprint 4, we have:
- ✅ AWS SNS-based notifications for assignments, grades, and announcements
- ✅ AI-powered course recommendation engine
- ✅ NLP-based lecture summarization API
- ✅ Notification Center UI in frontend
- ✅ Integration of AI features in frontend
- ✅ Optimized frontend for production deployment
- ✅ CI/CD pipeline setup with AWS CloudWatch monitoring
- ✅ Comprehensive backend and frontend testing

---

## 💚 Notes & Discussions
- [ ] Improve AI recommendation algorithm (future enhancement)
- [ ] Integrate push notifications in mobile view (future enhancement)
- [ ] Plan final project wrap-up presentation and documentation

---

### 🔥 Sprint 4 Successfully Completed! 🚀

