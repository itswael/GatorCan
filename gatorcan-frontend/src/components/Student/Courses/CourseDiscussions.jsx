import React from "react";
import StudentNavbar from "../StudentNavbar";
import { useState } from "react";
import { Box, Typography } from "@mui/material";
import { useParams } from "react-router-dom";
import CourseNavbar from "./CourseNavbar";

function CourseDiscussions() {
  const [course_name, setCourseName] = useState("");
  const [loading, setLoading] = useState(false);
  const [errMessage, setErrMessage] = useState("");

  let { id } = useParams();
  console.log(id);

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
                  <Typography variant="h3">Discussions</Typography>
                </div>
              )}
            </div>
          </Box>
        </Box>
      </div>
    </>
  );
}

export default CourseDiscussions;
