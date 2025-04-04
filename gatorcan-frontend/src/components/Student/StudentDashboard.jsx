import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import ButtonBase from "@mui/material/CardActionArea";
import CampaignIcon from '@mui/icons-material/Campaign';
import Typography from "@mui/material/Typography";
import EditNoteIcon from "@mui/icons-material/EditNote";
import MarkUnreadChatAltIcon from "@mui/icons-material/MarkUnreadChatAlt";
import FolderCopyIcon from "@mui/icons-material/FolderCopy";
import { Container } from "@mui/material";
import StudentNavbar from "./StudentNavbar";
import CourseService from "../../services/CourseService";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

import React from 'react'

function MediaCard({ text1, text2, color, course_id }) {
  const navigate = useNavigate();
  const handleNavigate = (path) => {
    navigate(path, { replace: false });
  };
  return (
    <Card sx={{ maxWidth: 250, margin: "20px" }} elevation={10}>
      <ButtonBase
        onClick={() => handleNavigate(`/student-course/${course_id}`)}
      >
        <Container sx={{ backgroundColor: color, height: "150px" }}></Container>
        <CardContent>
          <Typography gutterBottom variant="h8" component="div" color={color}>
            {text1}
          </Typography>
          <Typography variant="body2" sx={{ color: "text.secondary" }}>
            {text2}
          </Typography>
        </CardContent>
      </ButtonBase>
    </Card>
  );
}

function StudentDashboard() {

  const [enrolledCourses, setEnrolledCourses] = useState([]);
  const [loadingEnrolledCourses, setLoadingEnrolledCourses] = useState(true);

  const loadEnrolledCourses = async () => {
    setLoadingEnrolledCourses(true);
    const courses = await CourseService.fetchEnrolledCourses();
    setEnrolledCourses(courses);
    setLoadingEnrolledCourses(false);
  };

  useEffect(() => {
    loadEnrolledCourses();
  }, []);

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
      <div style={{ marginLeft: "120px" }}>
        <h1>Dashboard</h1>
        <hr />
        <br></br>
        <h3>My Courses</h3>
        <div
          style={{
            display: "flex",
            flexWrap: "wrap",
            justifyContent: "left",
            flexDirection: "row",
            width: "80%",
          }}
        >
          {loadingEnrolledCourses ? (
            <p>Loading enrolled courses...</p>
          ) : enrolledCourses.length == 0 ? (
            <p>No enrolled courses</p>
          ) : (
            enrolledCourses.map((course, index) => {
              return (
                <MediaCard
                  key={course.id}
                  text1={course.name}
                  text2={course.description}
                  color={colors[index % colors.length]}
                  course_id={course.id}
                ></MediaCard>
              );
            })
          )}
        </div>
      </div>
    </>
  );
}

export default StudentDashboard