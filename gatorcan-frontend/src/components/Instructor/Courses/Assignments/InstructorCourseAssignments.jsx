import React, { useState, useEffect } from "react";
import {
  Container,
  Typography,
  Box,
  Button,
  Card,
  CardContent,
  Grid,
  CardActionArea,
} from "@mui/material";
import { useNavigate, useParams } from "react-router-dom";
import InstructorNavbar from "../../InstructorNavbar";
import InstructorCourseNavbar from "../InstructorCourseNavbar";
import dayjs from "dayjs";
import {fetchAssignments} from "../../../../services/CourseService";

const assignments = [
  {
    assignment_id: 1,
    title: "React Basics",
    description: "Build a simple React component with props and state.",
    deadline: "2025-04-25T23:59:00",
    course_id: 1,
    max_points: 100,
  },
  {
    assignment_id: 2,
    title: "React Router",
    description: "Implement routing with React Router DOM.",
    deadline: "2025-04-10T23:59:00",
    course_id: 1,
    max_points: 100,
  },
  {
    assignment_id: 3,
    title: "React Forms",
    description: "Create a form using controlled components.",
    deadline: "2025-04-18T23:59:00",
    course_id: 1,
    max_points: 50,
  },
];


const InstructorCourseAssignments = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const now = dayjs();

  const [assignments, setAssignments] = useState([]);

  const sortedAssignments = [...assignments]
    .filter((a) => a.course_id.toString() === id)
    .sort((a, b) => dayjs(b.deadline).diff(dayjs(a.deadline)));

  const isPastDeadline = (deadline) => dayjs(deadline).isBefore(now);

  useEffect(() => {
    const fetchData = async () => {
      const result = await fetchAssignments({ id });
      console.log(result);
      if (result) {
        const formattedAssignments = result.assignments.map((assignment) => ({
          assignment_id: assignment.id,
          title: assignment.title,
          description: assignment.description,
          deadline: assignment.deadline,
          course_id: assignment.course_id,
          max_points: assignment.max_points,
        }));
        setAssignments(formattedAssignments);
      }
    };
    fetchData();
  }, [id]);


  return (
    <>
      <InstructorNavbar />
      <Container sx={{ mt: 4 }}>

        <InstructorCourseNavbar activeTab="assignments" />

        <Box mt={3} display="flex" justifyContent="flex-end">
          <Button
            variant="contained"
            onClick={() =>
              navigate(`/instructor-courses/${id}/assignments/edit-assignment`)
            }
          >
            Add New Assignment
          </Button>
        </Box>

        <Grid container spacing={3} mt={1}>
          {sortedAssignments.map((assignment) => (
            <Grid item xs={12} md={6} key={assignment.assignment_id}>
              <Card
                sx={{
                  border: isPastDeadline(assignment.deadline)
                    ? "2px solid #d32f2f"
                    : "1px solid #1976d2",
                  backgroundColor: isPastDeadline(assignment.deadline)
                    ? "#fbe9e7"
                    : "#e3f2fd",
                }}
              >
                <CardActionArea
                  onClick={() =>
                    navigate(
                      `/instructor-courses/${id}/assignments/${assignment.assignment_id}`
                    )
                  }
                >
                  <CardContent>
                    <Typography variant="h6">{assignment.title}</Typography>
                    <Typography
                      variant="body2"
                      color="text.secondary"
                      gutterBottom
                    >
                      {assignment.description}
                    </Typography>
                    <Typography variant="caption" display="block">
                      Deadline:{" "}
                      {dayjs(assignment.deadline).format("MMM D, YYYY h:mm A")}
                    </Typography>
                    <Typography variant="caption" display="block">
                      Max Points: {assignment.max_points}
                    </Typography>
                  </CardContent>
                </CardActionArea>
              </Card>
            </Grid>
          ))}
        </Grid>
      </Container>
    </>
  );
};

export default InstructorCourseAssignments;
