=== RUN   TestGetAssignmentsByCourseID
=== RUN   TestGetAssignmentsByCourseID/Success
--- PASS: TestGetAssignmentsByCourseID/Success (0.00s)
=== RUN   TestGetAssignmentsByCourseID/No_Assignments_Found
--- PASS: TestGetAssignmentsByCourseID/No_Assignments_Found (0.00s)
--- PASS: TestGetAssignmentsByCourseID (0.00s)
=== RUN   TestGetAssignmentByIDAndCourseID
=== RUN   TestGetAssignmentByIDAndCourseID/Success
--- PASS: TestGetAssignmentByIDAndCourseID/Success (0.00s)
=== RUN   TestGetAssignmentByIDAndCourseID/Assignment_Not_Found
--- PASS: TestGetAssignmentByIDAndCourseID/Assignment_Not_Found (0.00s)
=== RUN   TestGetAssignmentByIDAndCourseID/Course_Not_Found
--- PASS: TestGetAssignmentByIDAndCourseID/Course_Not_Found (0.00s)
--- PASS: TestGetAssignmentByIDAndCourseID (0.00s)
=== RUN   TestUploadFileToAssignment_service
=== RUN   TestUploadFileToAssignment_service/Success
--- PASS: TestUploadFileToAssignment_service/Success (0.00s)
--- PASS: TestUploadFileToAssignment_service (0.00s)
=== RUN   TestCreateOrUpdateAssignment
=== RUN   TestCreateOrUpdateAssignment/Create_Assignment_Success
--- PASS: TestCreateOrUpdateAssignment/Create_Assignment_Success (0.00s)
=== RUN   TestCreateOrUpdateAssignment/Course_Not_Found
--- PASS: TestCreateOrUpdateAssignment/Course_Not_Found (0.00s)
=== RUN   TestCreateOrUpdateAssignment/Upsert_Error
--- PASS: TestCreateOrUpdateAssignment/Upsert_Error (0.00s)
--- PASS: TestCreateOrUpdateAssignment (0.00s)
=== RUN   TestGetAssignments
=== RUN   TestGetAssignments/Success
--- PASS: TestGetAssignments/Success (0.00s)
=== RUN   TestGetAssignments/Invalid_Course_ID
--- PASS: TestGetAssignments/Invalid_Course_ID (0.00s)
=== RUN   TestGetAssignments/No_Assignments_Found
--- PASS: TestGetAssignments/No_Assignments_Found (0.00s)
=== RUN   TestGetAssignments/Server_Error
--- PASS: TestGetAssignments/Server_Error (0.00s)
--- PASS: TestGetAssignments (0.00s)
=== RUN   TestGetAssignment
=== RUN   TestGetAssignment/Success
--- PASS: TestGetAssignment/Success (0.00s)
=== RUN   TestGetAssignment/Invalid_Course_ID
--- PASS: TestGetAssignment/Invalid_Course_ID (0.00s)
=== RUN   TestGetAssignment/Assignment_Not_Found
--- PASS: TestGetAssignment/Assignment_Not_Found (0.00s)
=== RUN   TestGetAssignment/Server_Error
--- PASS: TestGetAssignment/Server_Error (0.00s)
--- PASS: TestGetAssignment (0.00s)
=== RUN   TestUploadFileToAssignment
=== RUN   TestUploadFileToAssignment/Success
--- PASS: TestUploadFileToAssignment/Success (0.00s)
=== RUN   TestUploadFileToAssignment/Unauthorized
--- PASS: TestUploadFileToAssignment/Unauthorized (0.00s)
=== RUN   TestUploadFileToAssignment/Invalid_Course_ID
--- PASS: TestUploadFileToAssignment/Invalid_Course_ID (0.00s)
=== RUN   TestUploadFileToAssignment/User_Not_Found
--- PASS: TestUploadFileToAssignment/User_Not_Found (0.00s)
--- PASS: TestUploadFileToAssignment (0.00s)
=== RUN   TestCreateAssignment
=== RUN   TestCreateAssignment/Success
--- PASS: TestCreateAssignment/Success (0.00s)
=== RUN   TestCreateAssignment/Invalid_Course_ID
--- PASS: TestCreateAssignment/Invalid_Course_ID (0.00s)
=== RUN   TestCreateAssignment/Missing_Title
--- PASS: TestCreateAssignment/Missing_Title (0.00s)
=== RUN   TestCreateAssignment/Service_Error
--- PASS: TestCreateAssignment/Service_Error (0.00s)
--- PASS: TestCreateAssignment (0.00s)
=== RUN   TestGetEnrolledCourses
=== RUN   TestGetEnrolledCourses/Success_with_courses
--- PASS: TestGetEnrolledCourses/Success_with_courses (0.00s)
=== RUN   TestGetEnrolledCourses/Success_with_empty_courses
--- PASS: TestGetEnrolledCourses/Success_with_empty_courses (0.00s)
=== RUN   TestGetEnrolledCourses/User_not_found
--- PASS: TestGetEnrolledCourses/User_not_found (0.00s)
=== RUN   TestGetEnrolledCourses/Database_error
--- PASS: TestGetEnrolledCourses/Database_error (0.00s)
--- PASS: TestGetEnrolledCourses (0.00s)
=== RUN   TestGetCourses
=== RUN   TestGetCourses/Success_with_default_pagination
--- PASS: TestGetCourses/Success_with_default_pagination (0.00s)
=== RUN   TestGetCourses/Success_with_custom_pagination
--- PASS: TestGetCourses/Success_with_custom_pagination (0.00s)
=== RUN   TestGetCourses/Invalid_page_parameter
--- PASS: TestGetCourses/Invalid_page_parameter (0.00s)
=== RUN   TestGetCourses/Service_error
--- PASS: TestGetCourses/Service_error (0.00s)
--- PASS: TestGetCourses (0.00s)
=== RUN   TestEnrollInCourse
=== RUN   TestEnrollInCourse/Success
--- PASS: TestEnrollInCourse/Success (0.00s)
=== RUN   TestEnrollInCourse/Invalid_course_ID
--- PASS: TestEnrollInCourse/Invalid_course_ID (0.00s)
=== RUN   TestEnrollInCourse/Missing_course_ID
--- PASS: TestEnrollInCourse/Missing_course_ID (0.00s)
=== RUN   TestEnrollInCourse/User_not_found
--- PASS: TestEnrollInCourse/User_not_found (0.00s)
=== RUN   TestEnrollInCourse/Course_not_found
--- PASS: TestEnrollInCourse/Course_not_found (0.00s)
=== RUN   TestEnrollInCourse/Already_enrolled
--- PASS: TestEnrollInCourse/Already_enrolled (0.00s)
=== RUN   TestEnrollInCourse/Course_full
--- PASS: TestEnrollInCourse/Course_full (0.00s)
--- PASS: TestEnrollInCourse (0.00s)
=== RUN   TestGetCourse
=== RUN   TestGetCourse/Success
--- PASS: TestGetCourse/Success (0.00s)
=== RUN   TestGetCourse/Invalid_Course_ID
--- PASS: TestGetCourse/Invalid_Course_ID (0.00s)
=== RUN   TestGetCourse/Course_Not_Found
--- PASS: TestGetCourse/Course_Not_Found (0.00s)
=== RUN   TestGetCourse/Database_Error
--- PASS: TestGetCourse/Database_Error (0.00s)
--- PASS: TestGetCourse (0.00s)
=== RUN   TestGetCoursesService
=== RUN   TestGetCoursesService/Success_-_Full_Page
--- PASS: TestGetCoursesService/Success_-_Full_Page (0.00s)
=== RUN   TestGetCoursesService/Success_-_Empty_Page
--- PASS: TestGetCoursesService/Success_-_Empty_Page (0.00s)
=== RUN   TestGetCoursesService/User_Not_Found
--- PASS: TestGetCoursesService/User_Not_Found (0.00s)
=== RUN   TestGetCoursesService/Database_Error
--- PASS: TestGetCoursesService/Database_Error (0.00s)
=== RUN   TestGetCoursesService/Success_-_Instructor_Courses
--- PASS: TestGetCoursesService/Success_-_Instructor_Courses (0.00s)
=== RUN   TestGetCoursesService/Instructor_Database_Error
--- PASS: TestGetCoursesService/Instructor_Database_Error (0.00s)
--- PASS: TestGetCoursesService (0.00s)
=== RUN   TestGetEnrolledCourses_service
=== RUN   TestGetEnrolledCourses_service/Success_With_Enrolled_Courses
--- PASS: TestGetEnrolledCourses_service/Success_With_Enrolled_Courses (0.00s)
=== RUN   TestGetEnrolledCourses_service/Success_With_No_Courses
--- PASS: TestGetEnrolledCourses_service/Success_With_No_Courses (0.00s)
=== RUN   TestGetEnrolledCourses_service/User_Not_Found
--- PASS: TestGetEnrolledCourses_service/User_Not_Found (0.00s)
=== RUN   TestGetEnrolledCourses_service/Error_Fetching_Enrollments
--- PASS: TestGetEnrolledCourses_service/Error_Fetching_Enrollments (0.00s)
--- PASS: TestGetEnrolledCourses_service (0.00s)
=== RUN   TestEnrollUser
=== RUN   TestEnrollUser/Success
--- PASS: TestEnrollUser/Success (0.21s)
=== RUN   TestEnrollUser/User_Not_Found
--- PASS: TestEnrollUser/User_Not_Found (0.00s)
=== RUN   TestEnrollUser/Course_Not_Found
--- PASS: TestEnrollUser/Course_Not_Found (0.00s)
=== RUN   TestEnrollUser/Course_Full
--- PASS: TestEnrollUser/Course_Full (0.00s)
=== RUN   TestEnrollUser/Course_Not_Active_Yet
--- PASS: TestEnrollUser/Course_Not_Active_Yet (0.00s)
--- PASS: TestEnrollUser (0.21s)
=== RUN   TestCourseResponseDTOConversion
--- PASS: TestCourseResponseDTOConversion (0.00s)
=== RUN   TestGetCourseByID
=== RUN   TestGetCourseByID/Success
--- PASS: TestGetCourseByID/Success (0.00s)
=== RUN   TestGetCourseByID/Active_Course_Not_Found
--- PASS: TestGetCourseByID/Active_Course_Not_Found (0.00s)
=== RUN   TestGetCourseByID/Course_Details_Not_Found
--- PASS: TestGetCourseByID/Course_Details_Not_Found (0.00s)
=== RUN   TestGetCourseByID/Instructor_Not_Found
--- PASS: TestGetCourseByID/Instructor_Not_Found (0.00s)
--- PASS: TestGetCourseByID (0.00s)
=== RUN   TestGetSubmission
=== RUN   TestGetSubmission/Success
--- PASS: TestGetSubmission/Success (0.00s)
=== RUN   TestGetSubmission/Unauthorized
--- PASS: TestGetSubmission/Unauthorized (0.00s)
=== RUN   TestGetSubmission/User_Not_Found
--- PASS: TestGetSubmission/User_Not_Found (0.00s)
=== RUN   TestGetSubmission/Invalid_Course_ID
--- PASS: TestGetSubmission/Invalid_Course_ID (0.00s)
=== RUN   TestGetSubmission/Invalid_Assignment_ID
--- PASS: TestGetSubmission/Invalid_Assignment_ID (0.00s)
=== RUN   TestGetSubmission/Submission_Not_Found
--- PASS: TestGetSubmission/Submission_Not_Found (0.00s)
=== RUN   TestGetSubmission/Server_Error
--- PASS: TestGetSubmission/Server_Error (0.00s)
--- PASS: TestGetSubmission (0.00s)
=== RUN   TestGradeSubmission
=== RUN   TestGradeSubmission/Success
--- PASS: TestGradeSubmission/Success (0.00s)
=== RUN   TestGradeSubmission/Unauthorized
--- PASS: TestGradeSubmission/Unauthorized (0.00s)
=== RUN   TestGradeSubmission/Invalid_Course_ID
--- PASS: TestGradeSubmission/Invalid_Course_ID (0.00s)
=== RUN   TestGradeSubmission/Invalid_Request_Body
--- PASS: TestGradeSubmission/Invalid_Request_Body (0.00s)
=== RUN   TestGradeSubmission/Submission_Not_Found
--- PASS: TestGradeSubmission/Submission_Not_Found (0.00s)
=== RUN   TestGradeSubmission/Server_Error
--- PASS: TestGradeSubmission/Server_Error (0.00s)
=== RUN   TestGradeSubmission/SNS_Notification_Failed
--- PASS: TestGradeSubmission/SNS_Notification_Failed (0.00s)
--- PASS: TestGradeSubmission (0.00s)
=== RUN   TestGradeSubmission_service
=== RUN   TestGradeSubmission_service/Valid_submission_grading
--- PASS: TestGradeSubmission_service/Valid_submission_grading (0.00s)
=== RUN   TestGradeSubmission_service/User_not_found
TEST: 2025/04/22 00:37:03 User not found: nonexistentuser
--- PASS: TestGradeSubmission_service/User_not_found (0.00s)
=== RUN   TestGradeSubmission_service/Course_not_found
TEST: 2025/04/22 00:37:03 Course not found: 1
--- PASS: TestGradeSubmission_service/Course_not_found (0.00s)
=== RUN   TestGradeSubmission_service/Grading_submission_failed
TEST: 2025/04/22 00:37:03 Error grading submission: grading failed
--- PASS: TestGradeSubmission_service/Grading_submission_failed (0.00s)
--- PASS: TestGradeSubmission_service (0.00s)
=== RUN   TestGetSubmissionService
=== RUN   TestGetSubmissionService/Success
--- PASS: TestGetSubmissionService/Success (0.00s)
=== RUN   TestGetSubmissionService/Submission_Not_Found
--- PASS: TestGetSubmissionService/Submission_Not_Found (0.00s)
--- PASS: TestGetSubmissionService (0.00s)
=== RUN   TestCreateUser
=== RUN   TestCreateUser/Success
--- PASS: TestCreateUser/Success (0.00s)
=== RUN   TestCreateUser/Invalid_Email
--- PASS: TestCreateUser/Invalid_Email (0.00s)
=== RUN   TestCreateUser/User_Already_Exists
--- PASS: TestCreateUser/User_Already_Exists (0.00s)
--- PASS: TestCreateUser (0.00s)
=== RUN   TestLogin
=== RUN   TestLogin/Success
--- PASS: TestLogin/Success (0.00s)
=== RUN   TestLogin/Invalid_Credentials
--- PASS: TestLogin/Invalid_Credentials (0.00s)
=== RUN   TestLogin/User_Not_Found
--- PASS: TestLogin/User_Not_Found (0.00s)
--- PASS: TestLogin (0.00s)
=== RUN   TestGetUserDetails
=== RUN   TestGetUserDetails/Success
--- PASS: TestGetUserDetails/Success (0.00s)
=== RUN   TestGetUserDetails/User_Not_Found
--- PASS: TestGetUserDetails/User_Not_Found (0.00s)
--- PASS: TestGetUserDetails (0.00s)
=== RUN   TestDeleteUser
=== RUN   TestDeleteUser/Success
--- PASS: TestDeleteUser/Success (0.00s)
=== RUN   TestDeleteUser/User_Not_Found
--- PASS: TestDeleteUser/User_Not_Found (0.00s)
--- PASS: TestDeleteUser (0.00s)
=== RUN   TestUpdateUser
=== RUN   TestUpdateUser/Success
--- PASS: TestUpdateUser/Success (0.00s)
=== RUN   TestUpdateUser/Invalid_Password
--- PASS: TestUpdateUser/Invalid_Password (0.00s)
--- PASS: TestUpdateUser (0.00s)
=== RUN   TestUpdateRoles
=== RUN   TestUpdateRoles/Success
--- PASS: TestUpdateRoles/Success (0.00s)
=== RUN   TestUpdateRoles/User_Not_Found
--- PASS: TestUpdateRoles/User_Not_Found (0.00s)
--- PASS: TestUpdateRoles (0.00s)
=== RUN   TestLogin_service
=== RUN   TestLogin_service/Successful_Login
--- PASS: TestLogin_service/Successful_Login (0.32s)
=== RUN   TestLogin_service/User_Not_Found
--- PASS: TestLogin_service/User_Not_Found (0.00s)
=== RUN   TestLogin_service/Invalid_Password
--- PASS: TestLogin_service/Invalid_Password (0.27s)
--- PASS: TestLogin_service (0.59s)
=== RUN   TestCreateUser_service
=== RUN   TestCreateUser_service/Successful_User_Creation
--- PASS: TestCreateUser_service/Successful_User_Creation (0.07s)
=== RUN   TestCreateUser_service/User_Already_Exists
--- PASS: TestCreateUser_service/User_Already_Exists (0.00s)
=== RUN   TestCreateUser_service/Role_Not_Found
--- PASS: TestCreateUser_service/Role_Not_Found (0.07s)
--- PASS: TestCreateUser_service (0.14s)
=== RUN   TestGetUserDetails_service
=== RUN   TestGetUserDetails_service/Success
--- PASS: TestGetUserDetails_service/Success (0.00s)
=== RUN   TestGetUserDetails_service/User_Not_Found
--- PASS: TestGetUserDetails_service/User_Not_Found (0.00s)
--- PASS: TestGetUserDetails_service (0.00s)
=== RUN   TestDeleteUser_service
=== RUN   TestDeleteUser_service/Success
--- PASS: TestDeleteUser_service/Success (0.00s)
=== RUN   TestDeleteUser_service/User_Not_Found
--- PASS: TestDeleteUser_service/User_Not_Found (0.00s)
=== RUN   TestDeleteUser_service/Delete_Error
--- PASS: TestDeleteUser_service/Delete_Error (0.00s)
--- PASS: TestDeleteUser_service (0.00s)
=== RUN   TestUpdateUser_service
=== RUN   TestUpdateUser_service/Success
--- PASS: TestUpdateUser_service/Success (0.20s)
=== RUN   TestUpdateUser_service/User_Not_Found
--- PASS: TestUpdateUser_service/User_Not_Found (0.00s)
=== RUN   TestUpdateUser_service/Incorrect_Old_Password
--- PASS: TestUpdateUser_service/Incorrect_Old_Password (0.13s)
=== RUN   TestUpdateUser_service/Update_Error
--- PASS: TestUpdateUser_service/Update_Error (0.14s)
--- PASS: TestUpdateUser_service (0.47s)
=== RUN   TestUpdateRoles_service
=== RUN   TestUpdateRoles_service/Success
--- PASS: TestUpdateRoles_service/Success (0.00s)
=== RUN   TestUpdateRoles_service/User_Not_Found
--- PASS: TestUpdateRoles_service/User_Not_Found (0.00s)
=== RUN   TestUpdateRoles_service/Role_Not_Found
--- PASS: TestUpdateRoles_service/Role_Not_Found (0.00s)
=== RUN   TestUpdateRoles_service/Update_Error
--- PASS: TestUpdateRoles_service/Update_Error (0.00s)
--- PASS: TestUpdateRoles_service (0.00s)
PASS
ok      gatorcan-backend/unit_tests