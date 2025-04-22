import React, { useEffect, useRef, useState } from "react";
import {
  Grid,
  Paper,
  Input,
  Button,
  Box,
  Typography,
  MenuItem,
  TextField,
} from "@mui/material";

function ActivateCourse({ setCurrPage }) {
  const courseRef = useRef();

  const [selectedCourseId, setSelectedCourseId] = useState("");
  const [selectedInstructor, setSelectedInstructor] = useState("");
  const [description, setDescription] = useState("");
  const [startDate, setStartDate] = useState("");
  const [endDate, setEndDate] = useState("");
  const [capacity, setCapacity] = useState("");
  const [displaySuccess, setDisplaySuccess] = useState("");

  // Static data simulating API response
  const courses = [
    {
      id: "1",
      name: "React Fundamentals",
      description: "Learn React from scratch.",
    },
    { id: "2", name: "Advanced JS", description: "Deep dive into JavaScript." },
  ];

  const instructors = [
    { id: "I001", name: "Dr. Smith" },
    { id: "I002", name: "Prof. Jane" },
  ];

  useEffect(() => {
    courseRef.current?.focus();
  }, []);

  useEffect(() => {
    const course = courses.find((c) => c.id === selectedCourseId);
    setDescription(course ? course.description : "");
  }, [selectedCourseId]);

  const handleSubmit = (e) => {
    e.preventDefault();
    setDisplaySuccess(`Course "${selectedCourseId}" activated successfully!`);

    setSelectedCourseId("");
    setSelectedInstructor("");
    setStartDate("");
    setEndDate("");
    setCapacity("");
    setDescription("");
  };

  const paperStyle = {
    padding: 20,
    width: 400,
    margin: "19px auto",
  };
  const btnstyle = { backgroundColor: "#1B6DA1", margin: "20px 0" };
  const inputStyle = { margin: "12px auto" };

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
                  onClick={() => setDisplaySuccess("")}
                >
                  DONE
                </Button>
              </Paper>
            ) : (
              <form onSubmit={handleSubmit}>
                <Paper elevation={12} style={paperStyle}>
                  <Grid align="center">
                    <h2>Activate Course</h2>
                  </Grid>

                  <TextField
                    select
                    label="Choose Course"
                    fullWidth
                    value={selectedCourseId}
                    onChange={(e) => setSelectedCourseId(e.target.value)}
                    required
                    inputRef={courseRef}
                    style={inputStyle}
                  >
                    {courses.map((course) => (
                      <MenuItem key={course.id} value={course.id}>
                        {course.name}
                      </MenuItem>
                    ))}
                  </TextField>

                  <Input
                    type="text"
                    value={description}
                    placeholder="Course Description"
                    disabled
                    fullWidth
                    style={inputStyle}
                  />

                  <TextField
                    select
                    label="Choose Instructor"
                    fullWidth
                    value={selectedInstructor}
                    onChange={(e) => setSelectedInstructor(e.target.value)}
                    required
                    disabled={!selectedCourseId}
                    style={inputStyle}
                  >
                    {instructors.map((inst) => (
                      <MenuItem key={inst.id} value={inst.id}>
                        {inst.name}
                      </MenuItem>
                    ))}
                  </TextField>

                  <TextField
                    type="date"
                    label="Start Date"
                    fullWidth
                    value={startDate}
                    onChange={(e) => setStartDate(e.target.value)}
                    InputLabelProps={{ shrink: true }}
                    required
                    disabled={!selectedCourseId}
                    style={inputStyle}
                  />

                  <TextField
                    type="date"
                    label="End Date"
                    fullWidth
                    value={endDate}
                    onChange={(e) => setEndDate(e.target.value)}
                    InputLabelProps={{ shrink: true }}
                    required
                    disabled={!selectedCourseId}
                    style={inputStyle}
                  />

                  <TextField
                    type="number"
                    label="Capacity"
                    fullWidth
                    value={capacity}
                    onChange={(e) => setCapacity(e.target.value)}
                    required
                    disabled={!selectedCourseId}
                    style={inputStyle}
                  />

                  <Button
                    type="submit"
                    variant="contained"
                    color="primary"
                    fullWidth
                    style={btnstyle}
                  >
                    Submit
                  </Button>
                </Paper>
              </form>
            )}
          </Grid>
        </Box>
      </Grid>
    </>
  );
}

export default ActivateCourse;
