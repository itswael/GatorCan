import React from "react";
import { Container, Typography, Box } from "@mui/material";
import { useParams } from "react-router-dom";
import InstructorNavbar from "../InstructorNavbar";
import InstructorCourseNavbar from "./InstructorCourseNavbar";

const InstructorCourseAnnouncements = () => {
  const { id } = useParams();

  return (
    <>
      <InstructorNavbar />
      <Container sx={{ mt: 4 }}>

        <InstructorCourseNavbar activeTab="announcements" />

        <Box mt={4}>
          <Typography variant="body2" color="text.secondary">
            Announcement content goes here.
          </Typography>
        </Box>
      </Container>
    </>
  );
};

export default InstructorCourseAnnouncements;
