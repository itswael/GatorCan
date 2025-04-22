# Backend API Documentation

Login - This endpoint allows users to log into the system by providing their username and password in the request body. If the credentials are valid, the system responds with a success message and an authentication token, which is required for accessing protected resources. If the credentials are incorrect, an error message is returned.

Create User (Admin Only) - This endpoint allows an administrator to create a new user by providing a username, password, and role in the request body. The request must include a valid admin authentication token in the headers. If the input data is valid, the user is created successfully; otherwise, an error is returned. Unauthorized users attempting to access this endpoint will receive a forbidden response.

Delete User (Admin Only) - Admins can delete a user from the system by specifying their username in the URL. A valid admin authentication token is required in the request headers. If the specified user exists, they are removed from the system. If the user does not exist, a "User not found" error is returned. Unauthorized attempts will be rejected with a forbidden response.

Get User Details - This endpoint retrieves information about a specific user, including their username, role, and email. The request requires an authentication token in the headers. Both regular users and admins can access this endpoint, but they may only retrieve details of authorized users. If the specified user does not exist, a "User not found" error is returned.

Enroll in a Course - A user can enroll in a course by sending a request containing the course ID in the body. An authentication token is required in the headers. If the course exists and the user is authorized to enroll, the system responds with a success message. If the course ID is invalid, a bad request error is returned. Unauthorized users will receive a forbidden response.

Get Assignments for a Course - This endpoint allows authenticated users to retrieve all assignments for a specific course. The course ID is specified in the URL, and the response includes assignment details such as ID, title, description, and due date. If the course does not exist, an error message is returned. Accessing this endpoint without proper authentication is not allowed.

Update User Role (Admin Only)
This endpoint allows an admin to change a user's role. The request must include an admin authentication token in the headers. The request body should contain the username and the new role. If successful, the response confirms the update. If the input is invalid, a bad request error is returned. Unauthorized attempts will be rejected.

Update User Details
This endpoint allows a user to update their own details, such as email, phone number, and address. The request must include a valid user authentication token. If the provided data is valid, the system updates the details and returns a success message. Invalid input results in a bad request error, and unauthorized access attempts are rejected.

Get Enrolled Courses
Users can retrieve a list of courses they are enrolled in using this endpoint. The request requires an authentication token. If authorized, the system responds with a list of enrolled courses. If the user is not authenticated, an unauthorized error is returned.

Get All Available Courses
Users can retrieve a list of all courses offered by the system. The request must include an authentication token. If authorized, the response contains course details. Unauthorized access attempts are denied.

Get Course Details
Users can view details of a specific course by providing the course ID in the URL. A valid authentication token is required. If the course exists, details such as course name, instructor, and description are returned. If the course is not found, an error message is returned.

Get Course Assignments
This endpoint allows users to retrieve all assignments for a particular course. The request requires authentication, and the course ID must be included in the URL. If the course exists, the response includes assignment details such as title, description, deadline, and max points. If the course is not found, an error message is returned.

Get Assignment Details
Users can retrieve details of a specific assignment by specifying the course ID and assignment ID in the URL. The response includes information like title, description, deadline, and max points. If the assignment does not exist, an error message is returned.

Upload Assignment File
Users can upload a file for a specific assignment by sending the assignment ID, course ID, file name, and file URL in the request body. A valid authentication token is required. If successful, the system confirms the upload and returns the file URL. If the assignment or course is not found, an error is returned. If an error occurs during upload, a server error response is given.


## 1. Login

### Brief Description
Allows users to log in to the system by providing their credentials.

### URL
POST /login

### Headers
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Content-Type  | application/json | Specifies the request body format. |

### Request Body
json
{
  "username": "string",
  "password": "string"
}


### Response Body
#### Success (200 OK)
json
{
  "message": "Login successful",
  "token": "string"
}


#### Failure (401 Unauthorized)
json
{
  "error": "Invalid credentials"
}


### Status Codes
- 200 OK: Login successful.
- 401 Unauthorized: Invalid credentials.

---

## 2. Create User (Admin Only)

### Brief Description
Allows an admin to create a new user.

### URL
POST /admin/add_user

### Headers
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Content-Type  | application/json | Specifies the request body format. |
| Authorization | Bearer <token> | Admin authentication token. |

### Request Body
json
{
  "username": "string",
  "password": "string",
  "role": "string"
}


### Response Body
#### Success (201 Created)
json
{
  "message": "User created successfully"
}


#### Failure (400 Bad Request)
json
{
  "error": "Invalid input data"
}


#### Failure (403 Forbidden)
json
{
  "error": "Unauthorized access"
}


### Status Codes
- 201 Created: User created successfully.
- 400 Bad Request: Invalid input data.
- 403 Forbidden: Unauthorized access.

---

## 3. Delete User (Admin Only)

### Brief Description
Allows an admin to delete a user by username.

### URL
DELETE /admin/:username

### Headers
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Authorization | Bearer <token> | Admin authentication token. |

### Request Body
None.

### Response Body
#### Success (200 OK)
json
{
  "message": "User deleted successfully"
}


#### Failure (404 Not Found)
json
{
  "error": "User not found"
}


#### Failure (403 Forbidden)
json
{
  "error": "Unauthorized access"
}


### Status Codes
- 200 OK: User deleted successfully.
- 404 Not Found: User not found.
- 403 Forbidden: Unauthorized access.

---

## 4. Get User Details

### Brief Description
Allows a user or admin to retrieve user details by username.

### URL
GET /user/:username

### Headers
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Authorization | Bearer <token> | Authentication token. |

### Request Body
None.

### Response Body
#### Success (200 OK)
json
{
  "username": "string",
  "role": "string",
  "email": "string"
}


#### Failure (404 Not Found)
json
{
  "error": "User not found"
}


### Status Codes
- 200 OK: User details retrieved successfully.
- 404 Not Found: User not found.

---

## 5. Enroll in a Course

### Brief Description
Allows a user to enroll in a course.

### URL
POST /courses/enroll

### Headers
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Content-Type  | application/json | Specifies the request body format. |
| Authorization | Bearer <token> | Authentication token. |

### Request Body
json
{
  "course_id": "string"
}


### Response Body
#### Success (200 OK)
json
{
  "message": "Enrolled in course successfully"
}


#### Failure (400 Bad Request)
json
{
  "error": "Invalid course ID"
}


#### Failure (403 Forbidden)
json
{
  "error": "Unauthorized access"
}


### Status Codes
- 200 OK: Enrolled in course successfully.
- 400 Bad Request: Invalid course ID.
- 403 Forbidden: Unauthorized access.

---

## 6. Get Assignments for a Course

### Brief Description
Allows users to retrieve all assignments for a specific course.

### URL
GET /courses/:cid/assignments

### Headers
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Authorization | Bearer <token> | Authentication token. |

### Request Body
None.

### Response Body
#### Success (200 OK)
json
[
  {
    "assignment_id": "string",
    "title": "string",
    "description": "string",
    "due_date": "string"
  }
]


#### Failure (404 Not Found)
json
{
  "error": "Course not found"
}


### Status Codes
- 200 OK: Assignments retrieved successfully.
- 404 Not Found: Course not found.

---

## *7. PUT /admin/update_role*

### *Brief Description*
Allows an admin to update a user's role in the system.

### *URL*
PUT /admin/update_role

### *Headers*
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Content-Type  | application/json | Specifies the request body format. |
| Authorization | Bearer <token> | Admin authentication token. |

### *Request Body*
json
{
  "username": "string",
  "new_role": "string"
}


### *Response Body*
#### *Success (200 OK)*
json
{
  "message": "User role updated successfully"
}


#### *Failure (400 Bad Request)*
json
{
  "error": "Invalid input data"
}


#### *Failure (403 Forbidden)*
json
{
  "error": "Unauthorized access"
}


### *Status Codes*
- 200 OK: User role updated successfully.
- 400 Bad Request: Invalid input data.
- 403 Forbidden: Unauthorized access.

---

## *8. PUT /user/update*

### *Brief Description*
Allows a user to update their own details.

### *URL*
PUT /user/update

### *Headers*
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Content-Type  | application/json | Specifies the request body format. |
| Authorization | Bearer <token> | User authentication token. |

### *Request Body*
json
{
  "email": "string",
  "phone": "string",
  "address": "string"
}


### *Response Body*
#### *Success (200 OK)*
json
{
  "message": "User details updated successfully"
}


#### *Failure (400 Bad Request)*
json
{
  "error": "Invalid input data"
}


#### *Failure (401 Unauthorized)*
json
{
  "error": "Unauthorized access"
}


### *Status Codes*
- 200 OK: User details updated successfully.
- 400 Bad Request: Invalid input data.
- 401 Unauthorized: Unauthorized access.

---

## *9. GET /courses/enrolled*

### *Brief Description*
Allows a user to retrieve a list of courses they are enrolled in.

### *URL*
GET /courses/enrolled

### *Headers*
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Authorization | Bearer <token> | User authentication token. |

### *Request Body*
None.

### *Response Body*
#### *Success (200 OK)*
json
[
  {
    "course_id": "string",
    "course_name": "string",
    "instructor": "string"
  }
]


#### *Failure (401 Unauthorized)*
json
{
  "error": "Unauthorized access"
}


### *Status Codes*
- 200 OK: List of enrolled courses retrieved successfully.
- 401 Unauthorized: Unauthorized access.

---

## *10. GET /courses/*

### *Brief Description*
Allows a user to retrieve a list of all available courses.

### *URL*
GET /courses/

### *Headers*
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Authorization | Bearer <token> | User authentication token. |

### *Request Body*
None.

### *Response Body*
#### *Success (200 OK)*
json
[
  {
    "course_id": "string",
    "course_name": "string",
    "instructor": "string"
  }
]


#### *Failure (401 Unauthorized)*
json
{
  "error": "Unauthorized access"
}


### *Status Codes*
- 200 OK: List of courses retrieved successfully.
- 401 Unauthorized: Unauthorized access.

---

## *11. GET /courses/:cid*

### *Brief Description*
Allows a user to retrieve details of a specific course by course ID.

### *URL*
GET /courses/:cid

### *Headers*
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Authorization | Bearer <token> | User authentication token. |

### *Request Body*
None.

### *Response Body*
#### *Success (200 OK)*
json
{
  "course_id": "string",
  "course_name": "string",
  "instructor": "string",
  "description": "string"
}


#### *Failure (404 Not Found)*
json
{
  "error": "Course not found"
}


#### *Failure (401 Unauthorized)*
json
{
  "error": "Unauthorized access"
}


### *Status Codes*
- 200 OK: Course details retrieved successfully.
- 404 Not Found: Course not found.
- 401 Unauthorized: Unauthorized access.

---

## *12. GET /courses/:cid/assignments*

### *Brief Description*
Allows a user to retrieve all assignments for a specific course.

### *URL*
GET /courses/:cid/assignments

### *Headers*
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Authorization | Bearer <token> | User authentication token. |

### *Request Body*
None.

### *Response Body*
#### *Success (200 OK)*
json
[
  {
    "assignment_id": "string",
    "title": "string",
    "description": "string",
    "deadline": "string",
    "max_points": "number"
  }
]


#### *Failure (404 Not Found)*
json
{
  "error": "Course not found"
}


### *Status Codes*
- 200 OK: Assignments retrieved successfully.
- 404 Not Found: Course not found.

---

## *13. GET /courses/:cid/assignments/:aid*

### *Brief Description*
Allows a user to retrieve details of a specific assignment by assignment ID and course ID.

### *URL*
GET /courses/:cid/assignments/:aid

### *Headers*
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Authorization | Bearer <token> | User authentication token. |

### *Request Body*
None.

### *Response Body*
#### *Success (200 OK)*
json
{
  "assignment_id": "string",
  "title": "string",
  "description": "string",
  "deadline": "string",
  "max_points": "number"
}


#### *Failure (404 Not Found)*
json
{
  "error": "Assignment not found"
}


### *Status Codes*
- 200 OK: Assignment details retrieved successfully.
- 404 Not Found: Assignment not found.

---

## *14. POST /courses/:cid/assignments/:aid/upload*

### *Brief Description*
Allows a user to upload a file to a specific assignment.

### *URL*
POST /courses/:cid/assignments/:aid/upload

### *Headers*
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Content-Type  | application/json | Specifies the request body format. |
| Authorization | Bearer <token> | User authentication token. |

### *Request Body*
json
{
  "assignment_id": "number",
  "course_id": "number",
  "file_name": "string",
  "file_url": "string",
  "file_type": "string"
}


### *Response Body*
#### *Success (200 OK)*
json
{
  "message": "File uploaded successfully",
  "file_url": "string"
}


#### *Failure (404 Not Found)*
json
{
  "error": "Assignment or course not found"
}


#### *Failure (500 Internal Server Error)*
json
{
  "error": "Error uploading file"
}


### *Status Codes*
- 200 OK: File uploaded successfully.
- 404 Not Found: Assignment or course not found.
- 500 Internal Server Error: Error uploading file.

---
## *15. POST /courses/:cid/assignments/:aid/grade*

### *Brief Description*
Allows an instructor to grade a specific assignment submission.

### *URL*
POST /courses/:cid/assignments/:aid/grade

### *Headers*
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Content-Type  | application/json | Specifies the request body format. |
| Authorization | Bearer <token>   | Instructor authentication token. |

### *Request Body*
json
{
  "submission_id": "string",
  "grade": "number",
  "feedback": "string"
}

### *Response Body*
#### *Success (200 OK)*
json
{
  "message": "Submission graded successfully"
}

#### *Failure (404 Not Found)*
json
{
  "error": "Submission or assignment not found"
}

#### *Failure (403 Forbidden)*
json
{
  "error": "Unauthorized access"
}

### *Status Codes*
- 200 OK: Submission graded successfully.
- 404 Not Found: Submission or assignment not found.
- 403 Forbidden: Unauthorized access.

---

## *16. POST /courses/:cid/upsertassignment*

### *Brief Description*
Allows an instructor to create or update an assignment for a specific course.

### *URL*
POST /courses/:cid/upsertassignment

### *Headers*
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Content-Type  | application/json | Specifies the request body format. |
| Authorization | Bearer <token>   | Instructor authentication token. |

### *Request Body*
json
{
  "assignment_id": "string (optional)",
  "title": "string",
  "description": "string",
  "deadline": "string",
  "max_points": "number"
}

### *Response Body*
#### *Success (200 OK)*
json
{
  "message": "Assignment created/updated successfully"
}

#### *Failure (400 Bad Request)*
json
{
  "error": "Invalid input data"
}

#### *Failure (403 Forbidden)*
json
{
  "error": "Unauthorized access"
}

### *Status Codes*
- 200 OK: Assignment created/updated successfully.
- 400 Bad Request: Invalid input data.
- 403 Forbidden: Unauthorized access.

---

## *17. GET /assignments/:aid/submissions*

### *Brief Description*
Allows an instructor to retrieve all submissions for a specific assignment.

### *URL*
GET /assignments/:aid/submissions

### *Headers*
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Authorization | Bearer <token>   | Instructor authentication token. |

### *Request Body*
None.

### *Response Body*
#### *Success (200 OK)*
json
[
  {
    "submission_id": "string",
    "student_id": "string",
    "file_url": "string",
    "submitted_at": "string",
    "grade": "number (optional)"
  }
]

#### *Failure (404 Not Found)*
json
{
  "error": "Assignment not found"
}

### *Status Codes*
- 200 OK: Submissions retrieved successfully.
- 404 Not Found: Assignment not found.

---

## *18. GET /courses/:cid/grades*

### *Brief Description*
Allows an instructor to retrieve all grades for a specific course.

### *URL*
GET /courses/:cid/grades

### *Headers*
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Authorization | Bearer <token>   | Instructor authentication token. |

### *Request Body*
None.

### *Response Body*
#### *Success (200 OK)*
json
[
  {
    "student_id": "string",
    "assignment_id": "string",
    "grade": "number",
    "feedback": "string"
  }
]

#### *Failure (404 Not Found)*
json
{
  "error": "Course not found"
}

### *Status Codes*
- 200 OK: Grades retrieved successfully.
- 404 Not Found: Course not found.

---

## *19. GET /courses/*

### *Brief Description*
Allows an instructor to retrieve a list of all courses they manage.

### *URL*
GET /courses/

### *Headers*
| Key           | Value            | Description              |
|---------------|------------------|--------------------------|
| Authorization | Bearer <token>   | Instructor authentication token. |

### *Request Body*
None.

### *Response Body*
#### *Success (200 OK)*
json
[
  {
    "course_id": "string",
    "course_name": "string",
    "description": "string",
    "instructor": "string"
  }
]

#### *Failure (401 Unauthorized)*
json
{
  "error": "Unauthorized access"
}

### *Status Codes*
- 200 OK: List of courses retrieved successfully.
- 401 Unauthorized: Unauthorized access.

---
