import React, { useEffect, useState } from "react";
import StudentNavbar from "./StudentNavbar";
import CourseService from "../../services/CourseService";

function StudentCourses() {
  const [allCourses, setAllCourses] = useState([]);
  const [enrolledCourses, setEnrolledCourses] = useState([]);
  const [recommendedCourses, setRecommendedCourses] = useState([]);
  const [loadingAllCourses, setLoadingAllCourses] = useState(true);
  const [loadingEnrolledCourses, setLoadingEnrolledCourses] = useState(true);
  const [loadingCourseRecommendations, setLoadingCourseRecommendations] = useState(true);

  useEffect(() => {
    loadCourses();
    loadEnrolledCourses();
    loadRecommendations();
  }, []);

  const loadRecommendations = async () => {
    try {
      setLoadingCourseRecommendations(true);

      // Get list of recommended courses
      const recommendations = await CourseService.fetchCourseRecommentations();

      if (!recommendations || recommendations.length === 0) {
        setRecommendedCourses([]);
        return;
      }

      // For each recommendation, fetch full course details in parallel
      const fullCourses = await Promise.all(
        recommendations.map((rec) => CourseService.fetchCourse({ id: rec.id }))
      );

      // Store the full course objects
      setRecommendedCourses(fullCourses);
    } catch (error) {
      console.error("Failed to load recommended courses:", error);
    } finally {
      setLoadingCourseRecommendations(false);
    }
  };

  const loadCourses = async () => {
    setLoadingAllCourses(true);
    const courses = await CourseService.fetchAllCourses();
    console.log("all courses");
    console.log(courses);
    setAllCourses(courses || []);
    setLoadingAllCourses(false);
  };

  const loadEnrolledCourses = async () => {
    setLoadingEnrolledCourses(true);
    const courses = await CourseService.fetchEnrolledCourses();
    console.log("enrolled courses");
    console.log(courses);
    setEnrolledCourses(courses);
    setLoadingEnrolledCourses(false);
  };

  const handleEnroll = async (courseID) => {
    const result = await CourseService.enrollInCourse(courseID);
    if (result.success) {
      loadEnrolledCourses();
      loadRecommendations();
    }
  };

  const colors = [
    "forestgreen",
    "darkorchid",
    "MediumVioletRed",
    "crimson",
    "teal",
    "royalblue",
    "darkslateblue",
    "indigo",
  ];

  return (
    <>
      <StudentNavbar />
      <div style={{ marginLeft: "120px", padding: "20px" }}>
        <h1>Courses</h1>
        <hr />

        <h3>Course Recommendations</h3>
        <hr />
        {loadingCourseRecommendations ? (
          <p>Loading recommendations...</p>
        ) : enrolledCourses.length === 0 ? (
          <p>Choose more courses to see recommendations</p>
        ) : (
          <div className="grid-container">
            {recommendedCourses != null && recommendedCourses.length > 0 ? (
              <>
                {recommendedCourses.map((course, index) => (
                  <CourseCard
                    key={course.id}
                    course={course}
                    enrollInCourse={handleEnroll}
                    color={colors[index + (42 % colors.length)]}
                  />
                ))}
              </>
            ) : (
              <p>Enroll in courses to see recommendations</p>
            )}
          </div>
        )}
        <h3>All Courses</h3>
        <hr />
        {loadingAllCourses ? (
          <p>Loading courses...</p>
        ) : allCourses.length === 0 ? (
          <p>No course available</p>
        ) : (
          <div className="grid-container">
            {allCourses.map((course, index) => (
              <CourseCard
                key={course.id}
                course={course}
                enrollInCourse={handleEnroll}
                color={colors[index % colors.length]}
              />
            ))}
          </div>
        )}
        <h3>Enrolled Courses</h3>
        <hr />
        {loadingEnrolledCourses ? (
          <p>Loading enrolled courses...</p>
        ) : enrolledCourses.length === 0 ? (
          <p>No enrolled courses</p>
        ) : (
          <div className="grid-container">
            {enrolledCourses.map((course, index) => (
              <CourseCard
                key={course.id}
                course={course}
                isEnrolled
                color={colors[index % colors.length]}
              />
            ))}
          </div>
        )}

        <br />
      </div>
      <style>
        {`
    .grid-container {
      display: grid;
      grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
      gap: 20px;
      padding: 20px;
    }

    @media (max-width: 900px) {
      .grid-container {
        grid-template-columns: repeat(2, minmax(300px, 1fr));
      }
    }

    @media (max-width: 600px) {
      .grid-container {
        grid-template-columns: repeat(1, minmax(300px, 1fr));
      }
    }
  `}
      </style>
    </>
  );
}

const CourseCard = ({ course, enrollInCourse, isEnrolled = false, color }) => {
  return (
    <div
      style={{
        border: "1px solid #ddd",
        borderRadius: "8px",
        padding: "15px",
        boxShadow: "2px 2px 10px rgba(0,0,0,0.1)",
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
        minHeight: "200px",
        position: "relative",
      }}
    >
      <h3 style={{ textAlign: "center", color: color }}>{course.name}</h3>
      <p style={{ textAlign: "left", marginTop: "5px" }}>
        {course.description}
      </p>

      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          fontSize: "14px",
          marginTop: "10px",
          color: "#555",
        }}
      >
        {course.created_at == null ? (
          <></>
        ) : (
          <>
            <span>
              <strong>Created:</strong>{" "}
              {new Date(course.created_at).toLocaleDateString()}
            </span>
            <span>
              <strong>Updated:</strong>{" "}
              {new Date(course.updated_at).toLocaleDateString()}
            </span>
          </>
        )}
      </div>

      {isEnrolled ? (
        <>
          <p style={{ textAlign: "left", marginTop: "10px" }}>
            <strong>Instructor:</strong> {course.instructorName} (
            {course.instructorEmail})
          </p>
          <p
            style={{
              color: "green",
              fontWeight: "bold",
              textAlign: "center",
              marginTop: "10px",
            }}
          >
            Enrolled
          </p>
        </>
      ) : (
        <button
          style={{
            alignSelf: "flex-end",
            padding: "8px 12px",
            backgroundColor: "#ff8c00", // Orange color
            color: "#fff",
            border: "none",
            cursor: "pointer",
            borderRadius: "4px",
            marginTop: "15px",
            position: "relative",
            top: "-10px",
          }}
          onClick={() => enrollInCourse(course.id)}
        >
          Enroll
        </button>
      )}
    </div>
  );
};

export default StudentCourses;
