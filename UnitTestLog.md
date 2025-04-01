=== RUN   TestGetAssignmentsByCourseID
=== RUN   TestGetAssignmentsByCourseID/Success
=== RUN   TestGetAssignmentsByCourseID/No_Assignments_Found
--- PASS: TestGetAssignmentsByCourseID (0.00s)
    --- PASS: TestGetAssignmentsByCourseID/Success (0.00s)
    --- PASS: TestGetAssignmentsByCourseID/No_Assignments_Found (0.00s)
=== RUN   TestGetAssignmentByIDAndCourseID
=== RUN   TestGetAssignmentByIDAndCourseID/Success
=== RUN   TestGetAssignmentByIDAndCourseID/Assignment_Not_Found
=== RUN   TestGetAssignmentByIDAndCourseID/Course_Not_Found
--- PASS: TestGetAssignmentByIDAndCourseID (0.00s)
    --- PASS: TestGetAssignmentByIDAndCourseID/Success (0.00s)
    --- PASS: TestGetAssignmentByIDAndCourseID/Assignment_Not_Found (0.00s)
    --- PASS: TestGetAssignmentByIDAndCourseID/Course_Not_Found (0.00s)
=== RUN   TestGetAssignments
=== RUN   TestGetAssignments/Success
=== RUN   TestGetAssignments/Invalid_Course_ID
=== RUN   TestGetAssignments/No_Assignments_Found
=== RUN   TestGetAssignments/Server_Error
--- PASS: TestGetAssignments (0.00s)
    --- PASS: TestGetAssignments/Success (0.00s)
    --- PASS: TestGetAssignments/Invalid_Course_ID (0.00s)
    --- PASS: TestGetAssignments/No_Assignments_Found (0.00s)
    --- PASS: TestGetAssignments/Server_Error (0.00s)
=== RUN   TestGetAssignment
=== RUN   TestGetAssignment/Success
=== RUN   TestGetAssignment/Invalid_Course_ID
=== RUN   TestGetAssignment/Assignment_Not_Found
=== RUN   TestGetAssignment/Server_Error
--- PASS: TestGetAssignment (0.00s)
    --- PASS: TestGetAssignment/Success (0.00s)
    --- PASS: TestGetAssignment/Invalid_Course_ID (0.00s)
    --- PASS: TestGetAssignment/Assignment_Not_Found (0.00s)
    --- PASS: TestGetAssignment/Server_Error (0.00s)
=== RUN   TestUploadFileToAssignment
=== RUN   TestUploadFileToAssignment/Success
=== RUN   TestUploadFileToAssignment/Unauthorized
=== RUN   TestUploadFileToAssignment/Invalid_Course_ID
=== RUN   TestUploadFileToAssignment/User_Not_Found
--- PASS: TestUploadFileToAssignment (0.00s)
    --- PASS: TestUploadFileToAssignment/Success (0.00s)
    --- PASS: TestUploadFileToAssignment/Unauthorized (0.00s)
    --- PASS: TestUploadFileToAssignment/Invalid_Course_ID (0.00s)
    --- PASS: TestUploadFileToAssignment/User_Not_Found (0.00s)
=== RUN   TestGetEnrolledCourses
=== RUN   TestGetEnrolledCourses/Success_with_courses
=== RUN   TestGetEnrolledCourses/Success_with_empty_courses
=== RUN   TestGetEnrolledCourses/User_not_found
=== RUN   TestGetEnrolledCourses/Database_error
--- PASS: TestGetEnrolledCourses (0.00s)
    --- PASS: TestGetEnrolledCourses/Success_with_courses (0.00s)
    --- PASS: TestGetEnrolledCourses/Success_with_empty_courses (0.00s)
    --- PASS: TestGetEnrolledCourses/User_not_found (0.00s)
    --- PASS: TestGetEnrolledCourses/Database_error (0.00s)
=== RUN   TestGetCourses
=== RUN   TestGetCourses/Success_with_default_pagination
=== RUN   TestGetCourses/Success_with_custom_pagination
=== RUN   TestGetCourses/Invalid_page_parameter
=== RUN   TestGetCourses/Service_error
--- PASS: TestGetCourses (0.00s)
    --- PASS: TestGetCourses/Success_with_default_pagination (0.00s)
    --- PASS: TestGetCourses/Success_with_custom_pagination (0.00s)
    --- PASS: TestGetCourses/Invalid_page_parameter (0.00s)
    --- PASS: TestGetCourses/Service_error (0.00s)
=== RUN   TestEnrollInCourse
=== RUN   TestEnrollInCourse/Success
=== RUN   TestEnrollInCourse/Invalid_course_ID
=== RUN   TestEnrollInCourse/Missing_course_ID
=== RUN   TestEnrollInCourse/User_not_found
=== RUN   TestEnrollInCourse/Course_not_found
=== RUN   TestEnrollInCourse/Already_enrolled
=== RUN   TestEnrollInCourse/Course_full
--- PASS: TestEnrollInCourse (0.00s)
    --- PASS: TestEnrollInCourse/Success (0.00s)
    --- PASS: TestEnrollInCourse/Invalid_course_ID (0.00s)
    --- PASS: TestEnrollInCourse/Missing_course_ID (0.00s)
    --- PASS: TestEnrollInCourse/User_not_found (0.00s)
    --- PASS: TestEnrollInCourse/Course_not_found (0.00s)
    --- PASS: TestEnrollInCourse/Already_enrolled (0.00s)
    --- PASS: TestEnrollInCourse/Course_full (0.00s)
=== RUN   TestGetCourse
=== RUN   TestGetCourse/Success
=== RUN   TestGetCourse/Invalid_Course_ID
=== RUN   TestGetCourse/Course_Not_Found
=== RUN   TestGetCourse/Database_Error
--- PASS: TestGetCourse (0.00s)
    --- PASS: TestGetCourse/Success (0.00s)
    --- PASS: TestGetCourse/Invalid_Course_ID (0.00s)
    --- PASS: TestGetCourse/Course_Not_Found (0.00s)
    --- PASS: TestGetCourse/Database_Error (0.00s)
=== RUN   TestGetCoursesService
=== RUN   TestGetCoursesService/Success_-_Full_Page
=== RUN   TestGetCoursesService/Success_-_Empty_Page
=== RUN   TestGetCoursesService/User_Not_Found
=== RUN   TestGetCoursesService/Database_Error
--- PASS: TestGetCoursesService (0.00s)
    --- PASS: TestGetCoursesService/Success_-_Full_Page (0.00s)
    --- PASS: TestGetCoursesService/Success_-_Empty_Page (0.00s)
    --- PASS: TestGetCoursesService/User_Not_Found (0.00s)
    --- PASS: TestGetCoursesService/Database_Error (0.00s)
=== RUN   TestGetEnrolledCourses_service
=== RUN   TestGetEnrolledCourses_service/Success_With_Enrolled_Courses
=== RUN   TestGetEnrolledCourses_service/Success_With_No_Courses
=== RUN   TestGetEnrolledCourses_service/User_Not_Found
=== RUN   TestGetEnrolledCourses_service/Error_Fetching_Enrollments
--- PASS: TestGetEnrolledCourses_service (0.00s)
    --- PASS: TestGetEnrolledCourses_service/Success_With_Enrolled_Courses (0.00s)
    --- PASS: TestGetEnrolledCourses_service/Success_With_No_Courses (0.00s)
    --- PASS: TestGetEnrolledCourses_service/User_Not_Found (0.00s)
    --- PASS: TestGetEnrolledCourses_service/Error_Fetching_Enrollments (0.00s)
=== RUN   TestEnrollUser
=== RUN   TestEnrollUser/Success
=== RUN   TestEnrollUser/User_Not_Found
=== RUN   TestEnrollUser/Course_Not_Found
=== RUN   TestEnrollUser/Course_Full
=== RUN   TestEnrollUser/Course_Not_Active_Yet
--- PASS: TestEnrollUser (0.21s)
    --- PASS: TestEnrollUser/Success (0.21s)
    --- PASS: TestEnrollUser/User_Not_Found (0.00s)
    --- PASS: TestEnrollUser/Course_Not_Found (0.00s)
    --- PASS: TestEnrollUser/Course_Full (0.00s)
    --- PASS: TestEnrollUser/Course_Not_Active_Yet (0.00s)
=== RUN   TestCourseResponseDTOConversion
--- PASS: TestCourseResponseDTOConversion (0.00s)
=== RUN   TestGetCourseByID
=== RUN   TestGetCourseByID/Success
=== RUN   TestGetCourseByID/Active_Course_Not_Found
=== RUN   TestGetCourseByID/Course_Details_Not_Found
=== RUN   TestGetCourseByID/Instructor_Not_Found
--- PASS: TestGetCourseByID (0.00s)
    --- PASS: TestGetCourseByID/Success (0.00s)
    --- PASS: TestGetCourseByID/Active_Course_Not_Found (0.00s)
    --- PASS: TestGetCourseByID/Course_Details_Not_Found (0.00s)
    --- PASS: TestGetCourseByID/Instructor_Not_Found (0.00s)
=== RUN   TestGetSubmission
=== RUN   TestGetSubmission/Success
=== RUN   TestGetSubmission/Unauthorized
=== RUN   TestGetSubmission/User_Not_Found
=== RUN   TestGetSubmission/Invalid_Course_ID
=== RUN   TestGetSubmission/Invalid_Assignment_ID
=== RUN   TestGetSubmission/Submission_Not_Found
=== RUN   TestGetSubmission/Server_Error
--- PASS: TestGetSubmission (0.00s)
    --- PASS: TestGetSubmission/Success (0.00s)
    --- PASS: TestGetSubmission/Unauthorized (0.00s)
    --- PASS: TestGetSubmission/User_Not_Found (0.00s)
    --- PASS: TestGetSubmission/Invalid_Course_ID (0.00s)
    --- PASS: TestGetSubmission/Invalid_Assignment_ID (0.00s)
    --- PASS: TestGetSubmission/Submission_Not_Found (0.00s)
    --- PASS: TestGetSubmission/Server_Error (0.00s)
=== RUN   TestGradeSubmission
=== RUN   TestGradeSubmission/Success
=== RUN   TestGradeSubmission/Unauthorized
=== RUN   TestGradeSubmission/Invalid_Course_ID
=== RUN   TestGradeSubmission/Invalid_Request_Body
=== RUN   TestGradeSubmission/Submission_Not_Found
=== RUN   TestGradeSubmission/Server_Error
--- PASS: TestGradeSubmission (0.00s)
    --- PASS: TestGradeSubmission/Success (0.00s)
    --- PASS: TestGradeSubmission/Unauthorized (0.00s)
    --- PASS: TestGradeSubmission/Invalid_Course_ID (0.00s)
    --- PASS: TestGradeSubmission/Invalid_Request_Body (0.00s)
    --- PASS: TestGradeSubmission/Submission_Not_Found (0.00s)
    --- PASS: TestGradeSubmission/Server_Error (0.00s)
=== RUN   TestGradeSubmission_service
=== RUN   TestGradeSubmission_service/Valid_submission_grading
=== RUN   TestGradeSubmission_service/User_not_found
TEST: 2025/03/30 22:45:03 User not found: nonexistentuser
=== RUN   TestGradeSubmission_service/Course_not_found
TEST: 2025/03/30 22:45:03 Course not found: 1
=== RUN   TestGradeSubmission_service/Grading_submission_failed
TEST: 2025/03/30 22:45:03 Error grading submission: grading failed
--- PASS: TestGradeSubmission_service (0.00s)
    --- PASS: TestGradeSubmission_service/Valid_submission_grading (0.00s)
    --- PASS: TestGradeSubmission_service/User_not_found (0.00s)
    --- PASS: TestGradeSubmission_service/Course_not_found (0.00s)
    --- PASS: TestGradeSubmission_service/Grading_submission_failed (0.00s)
=== RUN   TestGetSubmissionService
=== RUN   TestGetSubmissionService/Success
=== RUN   TestGetSubmissionService/Submission_Not_Found
--- PASS: TestGetSubmissionService (0.00s)
    --- PASS: TestGetSubmissionService/Success (0.00s)
    --- PASS: TestGetSubmissionService/Submission_Not_Found (0.00s)
=== RUN   TestCreateUser
=== RUN   TestCreateUser/Success
=== RUN   TestCreateUser/Invalid_Email
=== RUN   TestCreateUser/User_Already_Exists
--- PASS: TestCreateUser (0.00s)
    --- PASS: TestCreateUser/Success (0.00s)
    --- PASS: TestCreateUser/Invalid_Email (0.00s)
    --- PASS: TestCreateUser/User_Already_Exists (0.00s)
=== RUN   TestLogin
=== RUN   TestLogin/Success
=== RUN   TestLogin/Invalid_Credentials
=== RUN   TestLogin/User_Not_Found
--- PASS: TestLogin (0.00s)
    --- PASS: TestLogin/Success (0.00s)
    --- PASS: TestLogin/Invalid_Credentials (0.00s)
    --- PASS: TestLogin/User_Not_Found (0.00s)
=== RUN   TestGetUserDetails
=== RUN   TestGetUserDetails/Success
=== RUN   TestGetUserDetails/User_Not_Found
--- PASS: TestGetUserDetails (0.00s)
    --- PASS: TestGetUserDetails/Success (0.00s)
    --- PASS: TestGetUserDetails/User_Not_Found (0.00s)
=== RUN   TestDeleteUser
=== RUN   TestDeleteUser/Success
=== RUN   TestDeleteUser/User_Not_Found
--- PASS: TestDeleteUser (0.00s)
    --- PASS: TestDeleteUser/Success (0.00s)
    --- PASS: TestDeleteUser/User_Not_Found (0.00s)
=== RUN   TestUpdateUser
=== RUN   TestUpdateUser/Success
=== RUN   TestUpdateUser/Invalid_Password
--- PASS: TestUpdateUser (0.00s)
    --- PASS: TestUpdateUser/Success (0.00s)
    --- PASS: TestUpdateUser/Invalid_Password (0.00s)
=== RUN   TestUpdateRoles
=== RUN   TestUpdateRoles/Success
=== RUN   TestUpdateRoles/User_Not_Found
--- PASS: TestUpdateRoles (0.00s)
    --- PASS: TestUpdateRoles/Success (0.00s)
    --- PASS: TestUpdateRoles/User_Not_Found (0.00s)
=== RUN   TestLogin_service
=== RUN   TestLogin_service/Successful_Login
=== RUN   TestLogin_service/User_Not_Found
=== RUN   TestLogin_service/Invalid_Password
--- PASS: TestLogin_service (0.33s)
    --- PASS: TestLogin_service/Successful_Login (0.17s)
    --- PASS: TestLogin_service/User_Not_Found (0.00s)
    --- PASS: TestLogin_service/Invalid_Password (0.16s)
=== RUN   TestCreateUser_service
=== RUN   TestCreateUser_service/Successful_User_Creation
=== RUN   TestCreateUser_service/User_Already_Exists
=== RUN   TestCreateUser_service/Role_Not_Found
--- PASS: TestCreateUser_service (0.08s)
    --- PASS: TestCreateUser_service/Successful_User_Creation (0.04s)
    --- PASS: TestCreateUser_service/User_Already_Exists (0.00s)
    --- PASS: TestCreateUser_service/Role_Not_Found (0.04s)
=== RUN   TestGetUserDetails_service
=== RUN   TestGetUserDetails_service/Success
=== RUN   TestGetUserDetails_service/User_Not_Found
--- PASS: TestGetUserDetails_service (0.00s)
    --- PASS: TestGetUserDetails_service/Success (0.00s)
    --- PASS: TestGetUserDetails_service/User_Not_Found (0.00s)
=== RUN   TestDeleteUser_service
=== RUN   TestDeleteUser_service/Success
=== RUN   TestDeleteUser_service/User_Not_Found
=== RUN   TestDeleteUser_service/Delete_Error
--- PASS: TestDeleteUser_service (0.00s)
    --- PASS: TestDeleteUser_service/Success (0.00s)
    --- PASS: TestDeleteUser_service/User_Not_Found (0.00s)
    --- PASS: TestDeleteUser_service/Delete_Error (0.00s)
=== RUN   TestUpdateUser_service
=== RUN   TestUpdateUser_service/Success
=== RUN   TestUpdateUser_service/User_Not_Found
=== RUN   TestUpdateUser_service/Incorrect_Old_Password
=== RUN   TestUpdateUser_service/Update_Error
--- PASS: TestUpdateUser_service (0.29s)
    --- PASS: TestUpdateUser_service/Success (0.12s)
    --- PASS: TestUpdateUser_service/User_Not_Found (0.00s)
    --- PASS: TestUpdateUser_service/Incorrect_Old_Password (0.08s)
    --- PASS: TestUpdateUser_service/Update_Error (0.09s)
=== RUN   TestUpdateRoles_service
=== RUN   TestUpdateRoles_service/Success
=== RUN   TestUpdateRoles_service/User_Not_Found
=== RUN   TestUpdateRoles_service/Role_Not_Found
=== RUN   TestUpdateRoles_service/Update_Error
--- PASS: TestUpdateRoles_service (0.00s)
    --- PASS: TestUpdateRoles_service/Success (0.00s)
    --- PASS: TestUpdateRoles_service/User_Not_Found (0.00s)
    --- PASS: TestUpdateRoles_service/Role_Not_Found (0.00s)
    --- PASS: TestUpdateRoles_service/Update_Error (0.00s)
PASS
ok      gatorcan-backend/unit_tests