import React, { useEffect, useState} from "react";
import { useNavigate, useParams } from "react-router-dom";
import { Tabs, Tab, Paper, Typography } from "@mui/material";
import { fetchCourse } from "../../../services/CourseService";

const tabLabels = [
  "home",
  "syllabus",
  "assignments",
  "announcements",
  "discussions",
  "statistics",
];

const InstructorCourseNavbar = ({ activeTab }) => {
  const navigate = useNavigate();
  const { id } = useParams();
  const [courseTitle, setCourseTitle] = useState("");

  const handleChange = (event, newValue) => {
    navigate(`/instructor-courses/${id}/${tabLabels[newValue]}`);
  };

  useEffect(() => {
    const fetchData = async () => {
      const result = await fetchCourse({ id });
      if (result) {
        setCourseTitle(result.name);
      }
    };

    fetchData();
  
    return () => {}
  }, []);

  return (
    <>
      <Typography variant="h5" gutterBottom>
        #{id} - {courseTitle}
      </Typography>
      <Paper sx={{ mt: 4 }}>
        <Tabs
          value={tabLabels.indexOf(activeTab)}
          onChange={handleChange}
          variant="scrollable"
          scrollButtons="auto"
        >
          <Tab label="Home" />
          <Tab label="Syllabus" />
          <Tab label="Assignments" />
          <Tab label="Announcements" />
          <Tab label="Discussions" />
          <Tab label="Statistics" />
        </Tabs>
      </Paper>
    </>
  );
};

export default InstructorCourseNavbar;
