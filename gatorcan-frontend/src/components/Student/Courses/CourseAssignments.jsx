import React from "react";
import StudentNavbar from "../StudentNavbar";
import { useState, useEffect } from "react";
import { Box, Typography } from "@mui/material";
import { useParams, useNavigate } from "react-router-dom";
import CourseNavbar from "./CourseNavbar";
import { fetchAssignments } from "../../../services/CourseService";
import SchoolIcon from "@mui/icons-material/School";

function CourseAssignments() {
  const [loading, setLoading] = useState(false);
  const [errMessage, setErrMessage] = useState("");
  const [assignments, setAssignments] = useState([]);
  const [upcomingAssignments, setUpcomingAssignments] = useState([]);
  const [pastAssignments, setPastAssignments] = useState([]);

  let { id } = useParams();

  useEffect(() => {
    const fetchData = async () => {
    const assignmentsData = await fetchAssignments({ id });
    if (assignmentsData != null) {
      console.log(assignmentsData);
      setAssignments(assignmentsData.assignments);
      const currentDate = new Date();
      const upcomingAssignments = assignmentsData.assignments.filter(
        (assignment) => new Date(assignment.deadline) > currentDate
      );
      const pastAssignments = assignmentsData.assignments.filter(
        (assignment) => new Date(assignment.deadline) <= currentDate
      );
      setUpcomingAssignments(upcomingAssignments);
      setPastAssignments(pastAssignments);
    } else {
      setErrMessage("Unable to fetch assignments, retry");
    }
    setLoading(false);
  };
  fetchData();
  }, []);

  return (
    <>
      <StudentNavbar />
      <div style={{ marginLeft: "100px" }}>
        <Box sx={{ display: "flex", height: "100vh" }}>
          <CourseNavbar />
          <Box sx={{ flex: 1, padding: 2 }}>
            <div sx={{ display: "flex" }}>
              {errMessage != "" ? (
                <h4>{errMessage}</h4>
              ) : loading ? (
                <h4>Loading...</h4>
              ) : (
                <div>
                  <Typography variant="h3">Assignments</Typography>
                  <br></br>
                  <Typography variant="h4" sx={{ marginTop: "20px" }}>
                    Upcoming Assignments
                  </Typography>
                  {upcomingAssignments.length == 0 ? (
                    <Typography variant="body1">
                      No upcoming assignments
                    </Typography>
                  ) : (
                    upcomingAssignments.map((assignment, index) => (
                      <AssignmentCard
                        key={assignment.id}
                        id={assignment.id}
                        assignment={assignment}
                        upcoming={true}
                        course_id={id}
                      />
                    ))
                  )}
                  <Typography variant="h4" sx={{ marginTop: "20px" }}>
                    Past Assignments
                  </Typography>
                  {pastAssignments.length == 0 ? (
                    <Typography variant="body1">No past assignments</Typography>
                  ) : (
                    pastAssignments.map((assignment, index) => (
                      <AssignmentCard
                        key={index}
                        id={assignment.id}
                        assignment={assignment}
                        upcoming={false}
                        course_id={id}
                      />
                    ))
                  )}
                </div>
              )}
            </div>
          </Box>
        </Box>
      </div>
    </>
  );
}

function AssignmentCard({ id, assignment, upcoming, course_id }) {
  const navigate = useNavigate();
  const handleNavigation = () => {
    navigate("/student-courses/" + course_id + "/assignments/" + id, {
      state: { assignment },
      replace: false,
    });
  };

  return (
    <Box
      sx={{
        border: "1px solid #ccc",
        borderRadius: "8px",
        padding: "10px",
        marginBottom: "10px",
        backgroundColor: upcoming ? "#e8f5e9" : "#ffebee",
        display: "flex",
        alignItems: "center",
        cursor: upcoming ? "pointer" : "default",
      }}
      onClick={upcoming ? handleNavigation : undefined}
    >
      <Box sx={{ marginRight: "16px" }}>
        <SchoolIcon
          sx={{
            fontSize: "30px",
            color: upcoming ? "#4caf50" : "#f44336",
          }}
        />
      </Box>
      <Box sx={{ textAlign: "left" }}>
        <Typography variant="h6" sx={{ fontWeight: "bold", fontSize: "15px" }}>
          {assignment.title}
        </Typography>
        <Typography variant="body2" sx={{ color: "#757575" }}>
          Deadline: {new Date(assignment.deadline).toLocaleString()}
        </Typography>
        <Typography variant="body2" sx={{ color: "#757575" }}>
          Max Points: {assignment.max_points}
        </Typography>
      </Box>
    </Box>
  );
}

export default CourseAssignments;
