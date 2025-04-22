import React, { useEffect, useState } from "react";
import { Container, Typography, TextField, Button, Box } from "@mui/material";
import { useParams } from "react-router-dom";
import InstructorNavbar from "../../InstructorNavbar";
import InstructorCourseNavbar from "../InstructorCourseNavbar";
import CourseService from "../../../../services/CourseService";
import InstructorService from "../../../../services/InstructorService";

const InstructorCourseAssignmentEdit = () => {
  const { id, assignment_id } = useParams();

  const isEdit = !!assignment_id;

  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [deadline, setDeadline] = useState("");
  const [maxPoints, setMaxPoints] = useState("");
  const [notFound, setNotFound] = useState(false);

  useEffect(() => {
    const loadAssignment = async () => {
      if (isEdit) {
        const data = await CourseService.fetchAssignmentDetails({
          id,
          assignment_id,
        });

        const assignment = data?.assignments;
        console.log("Assignment data:", assignment);

        if (assignment) {
          setTitle(assignment.title);
          setDescription(assignment.description);
          setDeadline(new Date(assignment.deadline).toISOString().slice(0, 16));
          setMaxPoints(assignment.max_points.toString());
        } else {
          setNotFound(true);
        }
      }
    };

    loadAssignment();
  }, [isEdit, id, assignment_id]);

  const handleSubmit = async (e) => {
    e.preventDefault();

    const payload = {
      id: isEdit ? parseInt(assignment_id, 10) : 0,
      title,
      description,
      deadline: new Date(deadline).toISOString(),
      max_points: parseInt(maxPoints, 10),
    };

    const result = await InstructorService.upsertAssignment({
      courseID: id,
      assignment: payload,
    });

    if (result.success) {
      console.log("Assignment saved.");
      // TODO: add toast or redirect
      window.alert("Assignment saved successfully.");
      window.location.href = `/instructor-courses/${id}/assignments`;
    } else {
      window.alert(result.error);
    }
  };

  if (isEdit && notFound) {
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

        <Box component="form" onSubmit={handleSubmit} mt={4} maxWidth="sm">
          {isEdit && (
            <TextField
              label="Assignment ID"
              fullWidth
              disabled
              value={assignment_id}
              margin="normal"
            />
          )}

          <TextField
            label="Title"
            fullWidth
            required
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            margin="normal"
          />

          <TextField
            label="Description"
            multiline
            fullWidth
            minRows={3}
            required
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            margin="normal"
          />

          <TextField
            label="Deadline"
            type="datetime-local"
            fullWidth
            required
            value={deadline}
            onChange={(e) => setDeadline(e.target.value)}
            InputLabelProps={{ shrink: true }}
            margin="normal"
          />

          <TextField
            label="Max Points"
            type="number"
            fullWidth
            required
            value={maxPoints}
            onChange={(e) => setMaxPoints(e.target.value)}
            margin="normal"
          />

          <Box mt={3}>
            <Button type="submit" variant="contained">
              {isEdit ? "Save Changes" : "Submit Assignment"}
            </Button>
          </Box>
        </Box>
      </Container>
    </>
  );
};

export default InstructorCourseAssignmentEdit;
