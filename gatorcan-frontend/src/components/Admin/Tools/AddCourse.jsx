import React, { useRef, useState, useEffect } from "react";
import Grid from "@mui/material/Grid2";
import Paper from "@mui/material/Paper";
import Input from "@mui/material/Input";
import { Box, Typography, Button } from "@mui/material";

function AddCourse({ setCurrPage }) {
  const nameRef = useRef();
  const errRef = useRef();

  const [courseName, setCourseName] = useState("");
  const [courseDescription, setCourseDescription] = useState("");
  const [errMsg, setErrMsg] = useState("");
  const [displaySuccess, setDisplaySuccess] = useState("");

  useEffect(() => {
    nameRef.current.focus();
  }, []);

  useEffect(() => {
    setErrMsg("");
  }, [courseName, courseDescription]);

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Placeholder: simulate success
    setDisplaySuccess(`Course "${courseName}" created successfully!`);
    setCourseName("");
    setCourseDescription("");
  };

  const paperStyle = {
    padding: 20,
    width: 400,
    margin: "19px auto",
  };
  const btnstyle = { backgroundColor: "#1B6DA1", margin: "20px 0" };
  const inputStyle = { margin: "12px auto" };
  const errorStyle = { color: "red" };

  return (
    <>
      <Grid container spacing={5} maxWidth={600}>
        <Box
          sx={{
            minHeight: "100vh",
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
          }}
        >
          <Button
            style={{ marginTop: "80px", alignSelf: "flex-start" }}
            onClick={() => setCurrPage(-1)}
          >
            Back
          </Button>
          <Grid>
            {displaySuccess !== "" ? (
              <Paper elevation={12} style={paperStyle}>
                <Grid align="center">
                  <Typography>{displaySuccess}</Typography>
                </Grid>
                <Button
                  style={btnstyle}
                  color="primary"
                  variant="contained"
                  fullWidth
                  onClick={() => {
                    setDisplaySuccess("");
                  }}
                >
                  DONE
                </Button>
              </Paper>
            ) : (
              <div>
                <form onSubmit={handleSubmit}>
                  <Paper elevation={12} style={paperStyle}>
                    <Grid align="center">
                      <h2>Create Course</h2>
                    </Grid>
                    <Input
                      type="text"
                      id="courseName"
                      ref={nameRef}
                      autoComplete="off"
                      onChange={(e) => setCourseName(e.target.value)}
                      value={courseName}
                      required
                      style={inputStyle}
                      placeholder="Course Name"
                      fullWidth
                    />
                    <Input
                      type="text"
                      id="courseDescription"
                      autoComplete="off"
                      onChange={(e) => setCourseDescription(e.target.value)}
                      value={courseDescription}
                      required
                      style={inputStyle}
                      placeholder="Course Description"
                      fullWidth
                    />
                    <Button
                      style={btnstyle}
                      type="submit"
                      color="primary"
                      variant="contained"
                      fullWidth
                    >
                      Submit
                    </Button>
                  </Paper>
                </form>
                <p
                  ref={errRef}
                  className={errMsg ? "errmsg" : "offscreen"}
                  aria-live="assertive"
                  style={errorStyle}
                >
                  {errMsg}
                </p>
              </div>
            )}
          </Grid>
        </Box>
      </Grid>
    </>
  );
}

export default AddCourse;
