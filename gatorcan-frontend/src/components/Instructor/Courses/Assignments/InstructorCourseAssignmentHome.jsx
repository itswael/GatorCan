import React, { useEffect, useState } from "react";
import { Container, Typography, Box, Button, Paper } from "@mui/material";
import { useParams, useNavigate } from "react-router-dom";
import InstructorNavbar from "../../InstructorNavbar";
import InstructorCourseNavbar from "../InstructorCourseNavbar";
import CourseService from "../../../../services/CourseService";
import dayjs from "dayjs";

const InstructorCourseAssignmentHome = () => {
  const { id, assignment_id } = useParams();
  const navigate = useNavigate();

  const [assignment, setAssignment] = useState(null);

  useEffect(() => {
    const loadAssignment = async () => {
      const result = await CourseService.fetchAssignmentDetails({
        id,
        assignment_id,
      });
      console.log(result.assignments);
      setAssignment(result.assignments);
    };

    loadAssignment();
  }, [id, assignment_id]);

  if (!assignment) {
    return (
      <>
        <InstructorNavbar />
        <Container sx={{ mt: 4 }}>
          <Typography variant="h6">Assignment not found.</Typography>
        </Container>
      </>
    );
  }

  return (
    <>
      <InstructorNavbar />
      <Container sx={{ mt: 4 }}>
        <InstructorCourseNavbar activeTab="assignments" />

        <Paper elevation={3} sx={{ p: 3, mt: 4 }}>
          <Typography variant="body1" gutterBottom>
            <strong>Description:</strong> {assignment.description}
          </Typography>
          <Typography variant="body1" gutterBottom>
            <strong>Deadline:</strong>{" "}
            {dayjs(assignment.deadline).format("MMM D, YYYY h:mm A")}
          </Typography>
          <Typography variant="body1">
            <strong>Max Points:</strong> {assignment.max_points}
          </Typography>

          <Box mt={3} display="flex" gap={2}>
            <Button
              variant="outlined"
              onClick={() =>
                navigate(
                  `/instructor-courses/${id}/assignments/${assignment_id}/edit-assignment`
                )
              }
            >
              Edit
            </Button>
            <Button
              variant="contained"
              onClick={() =>
                navigate(
                  `/instructor-courses/${id}/assignments/${assignment_id}/view-submissions`
                )
              }
            >
              View Submissions
            </Button>
          </Box>
        </Paper>
      </Container>
    </>
  );
};

export default InstructorCourseAssignmentHome;
