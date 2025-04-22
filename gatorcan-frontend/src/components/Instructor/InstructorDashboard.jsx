import React from "react";
import {
  Container,
  Box,
  Grid,
  Card,
  CardContent,
  CardActionArea,
  Typography,
} from "@mui/material";
import { useNavigate } from "react-router-dom";
import InstructorNavbar from "./InstructorNavbar";

import { useEffect, useState } from "react";
import InstructorService from "../../services/InstructorService";


function InstructorDashboard() {
  const navigate = useNavigate();

  const [courses, setCourses] = useState([]);

  useEffect(() => {
    const loadCourses = async () => {
      const result = await InstructorService.fetchInstructorCourses();
      console.log(result);
      if (!result) {
        setCourses([]);
        return;
      }
      setCourses(result);
    };

    loadCourses();
  }, []);

  return (
    <>
      <InstructorNavbar />
      <Container sx={{ mt: 4 }}>
        {courses.length === 0 ? (
          <Typography variant="h5" align="center" color="textSecondary">
            No courses allocated to you.
          </Typography>
        ) : (
          <>
            <Typography variant="h5" gutterBottom>
              My Courses
            </Typography>
            <Grid container spacing={3}>
              {courses.map((course) => (
                <Grid item xs={12} sm={6} md={4} key={course.id}>
                  <Card>
                    <CardActionArea
                      onClick={() =>
                        navigate(`/instructor-courses/${course.id}/home`)
                      }
                    >
                      <CardContent>
                        <Typography variant="h6" gutterBottom>
                          #{course.id} - {course.name}
                        </Typography>
                        <Typography variant="body2" color="text.secondary">
                          {course.description}
                        </Typography>
                      </CardContent>
                    </CardActionArea>
                  </Card>
                </Grid>
              ))}
            </Grid>
          </>
        )}
      </Container>
    </>
  );
}

export default InstructorDashboard;
