import React from "react";
import { Container, Typography, Box } from "@mui/material";
import { useParams } from "react-router-dom";
import InstructorNavbar from "../InstructorNavbar";
import InstructorCourseNavbar from "./InstructorCourseNavbar";

const InstructorCourseDiscussions = () => {
  const { id } = useParams();

  return (
    <>
      <InstructorNavbar />
      <Container sx={{ mt: 4 }}>

        <InstructorCourseNavbar activeTab="discussions" />

        <Box mt={4}>
          <Typography variant="body2" color="text.secondary">
            Discussion content goes here.
          </Typography>
        </Box>
      </Container>
    </>
  );
};

export default InstructorCourseDiscussions;
