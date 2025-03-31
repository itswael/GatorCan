import React from "react";
import StudentNavbar from "../StudentNavbar";
import { useState, useEffect } from "react";
import { Box, Typography } from "@mui/material";
import { useParams } from "react-router-dom";
import CourseNavbar from "./CourseNavbar";
import { fetchAssignments } from "../../../services/CourseService";
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper } from "@mui/material";

function CourseGrades() {
  const [loading, setLoading] = useState(false);
  const [errMessage, setErrMessage] = useState("");

  let { id } = useParams();

  const [assignments, setAssignments] = useState([]);
  const [upcomingAssignments, setUpcomingAssignments] = useState([]);
  const [pastAssignments, setPastAssignments] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const assignmentsData = await fetchAssignments({ id });
      if (assignmentsData != null) {
        console.log(assignmentsData);
        setAssignments(assignmentsData.assignments);
        const currentDate = new Date();
        const upcomingAssignments = assignmentsData.assignments.filter(
          (assignment) => new Date(assignment.deadline) > currentDate
        );
        const pastAssignments = assignmentsData.assignments.filter(
          (assignment) => new Date(assignment.deadline) <= currentDate
        );
        setUpcomingAssignments(upcomingAssignments);
        setPastAssignments(pastAssignments);
      } else {
        setErrMessage("Unable to fetch assignments, retry");
      }
      setLoading(false);
    };
    fetchData();
  }, []);

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
                  <Typography variant="h3">Grades</Typography>
                  <Grades
                    upcomingAssignments={upcomingAssignments}
                    pastAssignments={pastAssignments}
                  ></Grades>
                </div>
              )}
            </div>
          </Box>
        </Box>
      </div>
    </>
  );
}

export default CourseGrades;

function Grades({ upcomingAssignments, pastAssignments }) {
  const renderAssignments = (assignments) => {
    return assignments
      .slice()
      .reverse()
      .map((assignment, index) => (
        <TableRow key={index}>
          <TableCell>{assignment.title}</TableCell>
          <TableCell>{new Date(assignment.deadline).toLocaleString()}</TableCell>
          <TableCell></TableCell>
          <TableCell>0</TableCell>
          <TableCell>{assignment.max_points}</TableCell>
        </TableRow>
      ));
  };

  const [feedback, setFeedback] = useState("");

  return (
    <div>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>
                <b>Name</b>
              </TableCell>
              <TableCell>
                <b>Due</b>
              </TableCell>
              <TableCell>
                <b>Submitted</b>
              </TableCell>
              <TableCell>
                <b>Score</b>
              </TableCell>
              <TableCell>
                <b>MaxPoints</b>
              </TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {renderAssignments(pastAssignments)}
            {renderAssignments(upcomingAssignments)}
          </TableBody>
          <TableRow>
            <TableCell colSpan={3}>Total</TableCell>
            <TableCell>0</TableCell>
            <TableCell>0.0%</TableCell>
          </TableRow>
        </Table>
      </TableContainer>
      <Box
        sx={{
          marginTop: "20px",
          textAlign: "left",
        }}
      >
        <Typography variant="h6">
          <strong>Feedback</strong>: {feedback == "" ? "N/A" : feedback}
        </Typography>
      </Box>
    </div>
  );
}