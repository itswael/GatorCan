import axios from "axios";

const base_url = "http://gatorcan-backend.us-east-2.elasticbeanstalk.com/courses";

const getAuthHeader = () => {
  const refreshToken = localStorage.getItem("refreshToken");
  return {
    Authorization: `Bearer ${refreshToken}`,
    "Content-Type": "application/json",
  };
};

// Fetch all available courses
export const fetchAllCourses = async () => {
  try {
    const response = await axios.get(`${base_url}/?page=1&pageSize=10`, {
      headers: getAuthHeader(),
    });
    return response.data;
  } catch (error) {
    console.error("Error fetching all courses:", error);
    return null;
  }
};

// Fetch enrolled courses
export const fetchEnrolledCourses = async () => {
  try {
    const response = await axios.get(`${base_url}/enrolled`, {
      headers: getAuthHeader(),
    });

    if (response.data === null) return [];

    return response.data.map((course) => ({
      id: course.ID,
      name: course.Name,
      description: course.Description,
      created_at: course.StartDate,
      updated_at: course.EndDate,
      instructorName: course.InstructorName,
      instructorEmail: course.InstructorEmail,
    }));
  } catch (error) {
    console.error("Error fetching enrolled courses:", error);
    return [];
  }
};

// Enroll in a course
export const enrollInCourse = async (courseID) => {
  try {
    const response = await axios.post(
      `${base_url}/enroll`,
      { courseID },
      { headers: getAuthHeader() }
    );

    //alert(`Successfully enrolled in course ID: ${courseID}`);
    return { success: true };
  } catch (error) {
    console.error("Error enrolling in course:", error);
    //alert("Enrollment failed!");
    return { success: false };
  }
};

// fetch current course
export const fetchCourse = async ({id}) => {
  try {
    const response = await axios.get(`${base_url}/${id}`, {
      headers: getAuthHeader(),
    });

    if (response.data === null) return [];
    return response.data;
  } catch (error) {
    console.error("Error fetching enrolled courses:", error);
    return [];
  }
};

// fetch current course assignments
export const fetchAssignments = async ({id}) => {
  try {
    const response = await axios.get(`${base_url}/${id}/assignments`, {
      headers: getAuthHeader(),
    });

    if (response.data === null) return [];
    return response.data;
  } catch (error) {
    console.error("Error fetching course assignments:", error);
    return [];
  }
};

// fetch current course assignment details
export const fetchAssignmentDetails = async ({ id, assignment_id }) => {
  try {
    const response = await axios.get(
      `${base_url}/${id}/assignments/${assignment_id}`,
      {
        headers: getAuthHeader(),
      }
    );

    if (response.data === null) return [];
    return response.data;
  } catch (error) {
    console.error("Error fetching course assignment details:", error);
    return [];
  }
};

// fetch current course
export const fetchCourseRecommentations = async () => {
  try {
    const response = await axios.get(`${base_url}/recommendations`, {
      headers: getAuthHeader(),
    });

    if (response.data === null) return [];
    return response.data;
  } catch (error) {
    console.error("Error fetching course recommendations:", error);
    return [];
  }
};

// submit assignment file
export const submitAssignmentFile = async ({
  course_id,
  assignment_id,
  data,
}) => {
  try {
    const response = await axios.post(
      `${base_url}/${course_id}/assignments/${assignment_id}/upload`,
      data,
      { headers: getAuthHeader() }
    );
    return { success: true, data: response.data };
  } catch (error) {
    console.error("Error submitting assignment file:", error);
    return { success: false };
  }
};

// submit assignment
export const submitAssignment = async ({
  course_id,
  assignment_id,
  data,
}) => {
  try {
    const response = await axios.post(
      `${base_url}/${course_id}/assignments/${assignment_id}/submit`,
      data,
      { headers: getAuthHeader() }
    );
    return { success: true, data: response.data };
  } catch (error) {
    console.error("Error submitting assignment file:", error);
    return { success: false };
  }
};

// fetch current course assignment submission details
export const fetchAssignmentSubmissionDetails = async ({ cid, aid}) => {
  try {
    const response = await axios.get(
      `${base_url}/${cid}/assignments/${aid}/submission`,
      {
        headers: getAuthHeader(),
      }
    );

    if (response.data === null) return {success: false};
    return {
      success: true,
      data: response.data
    }
  } catch (error) {
    console.error("Error fetching course assignment submission details:", error);
    return [];
  }
};

// fetch current course assignment submissions
export const fetchAssignmentSubmissions = async ({ cid, aid }) => {
  try {
    const response = await axios.get(
      `${base_url}/${cid}/assignments/${aid}/submissions`,
      {
        headers: getAuthHeader(),
      }
    );

    if (response.data === null) return { success: false };
    return {
      success: true,
      data: response.data,
    };
  } catch (error) {
    console.error(
      "Error fetching course assignment submission details:",
      error
    );
    return [];
  }
};

// fetch current course grades
export const fetchGrades = async ({id}) => {
  try {
    const response = await axios.get(`${base_url}/${id}/grades`, {
      headers: getAuthHeader(),
    });

    if (response.data === null) return [];
    return response.data;
  } catch (error) {
    console.error("Error fetching course assignments:", error);
    return [];
  }
};

export default {
  fetchAllCourses,
  fetchEnrolledCourses,
  enrollInCourse,
  fetchCourse,
  fetchAssignments,
  fetchAssignmentDetails,
  fetchCourseRecommentations,
  submitAssignmentFile,
  submitAssignment,
  fetchAssignmentSubmissionDetails,
  fetchAssignmentSubmissions,
  fetchGrades,
};
