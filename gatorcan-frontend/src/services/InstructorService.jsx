import axios from "axios";

const base_url = "http://gatorcan-backend.us-east-2.elasticbeanstalk.com/instructor";

const getAuthHeader = () => {
  const refreshToken = localStorage.getItem("refreshToken");
  return {
    Authorization: `Bearer ${refreshToken}`,
    "Content-Type": "application/json",
  };
};

// Fetch instructor courses
export const fetchInstructorCourses = async () => {
  try {
    const response = await axios.get(
      `http://gatorcan-backend.us-east-2.elasticbeanstalk.com/courses/?page=1&pageSize=10`,
      {
        headers: getAuthHeader(),
      }
    );
    console.log(response.data);
    if (!response.data) return [];
    return response.data.map((course) => ({
      id: course.id,
      name: course.name,
      description: course.description,
    }));
  } catch (error) {
    console.error("Error fetching instructor courses:", error);
    return [];
  }
};

// This function is used to create or update an assignment
export const upsertAssignment = async ({ courseID, assignment }) => {
  try {
    console.log(assignment);
    const response = await axios.post(
      `${base_url}/courses/${courseID}/upsertassignment`,
      assignment,
      { headers: getAuthHeader() }
    );
    if (!response || response.error)
      return { success: false, error: response.error ?? "Unknown error" };
    else return { success: true };
  } catch (error) {
    console.error("Error upserting assignment:", error);
    return { success: false };
  }
};

// This function is used to grade an assignment
export const gradeAssignment = async ({ cid, aid, grades }) => {
  try {
    console.log(grades);
    const response = await axios.post(
      `${base_url}/courses/${cid}/assignments/${aid}/grade`,
      grades,
      { headers: getAuthHeader() }
    );
    if (!response || response.error)
      return { success: false, error: response.error ?? "Unknown error" };
    else return { success: true };
  } catch (error) {
    console.error("Error upsertigradingng assignment:", error);
    return { success: false };
  }
};

export default {
  fetchInstructorCourses,
  upsertAssignment,
  gradeAssignment,
};
