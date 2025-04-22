import React from "react";
import StudentNavbar from "../StudentNavbar";
import { useState, useEffect } from "react";
import { Box, Typography } from "@mui/material";
import { useParams } from "react-router-dom";
import CourseNavbar from "./CourseNavbar";
import { fetchGrades } from "../../../services/CourseService";
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper } from "@mui/material";

function CourseGrades() {
  const [loading, setLoading] = useState(false);
  const [errMessage, setErrMessage] = useState("");
  const [grades, setGrades] = useState([]);

  let { id } = useParams();

  useEffect(() => {
    const fetchData = async () => {
      const assignmentsData = await fetchGrades({ id });
      if (assignmentsData != null) {
        console.log(assignmentsData);
        setGrades(assignmentsData);
      } else {
        setErrMessage("Unable to fetch grades, retry");
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
                  <Grades grades={grades}></Grades>
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

import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
} from "@mui/material";

function Grades({ grades }) {
  const [feedbackDialogOpen, setFeedbackDialogOpen] = useState(false);
  const [selectedFeedback, setSelectedFeedback] = useState("");

  const handleOpenFeedback = (assignment) => {
    const feedback = `Feedback for "${assignment.title}": ${(!assignment.feedback || assignment.feedback == "") ? "N/A" : assignment.feedback}`;
    setSelectedFeedback(feedback);
    setFeedbackDialogOpen(true);
  };

  const handleClose = () => {
    setFeedbackDialogOpen(false);
    setSelectedFeedback("");
  };

  const renderGrades = (grades) => {
    return grades
      .slice()
      .reverse()
      .map((grade, index) => (
        <TableRow key={index}>
          <TableCell>{grade.title}</TableCell>
          <TableCell>{new Date(grade.updated_at).toLocaleString()}</TableCell>
          <TableCell>{new Date(grade.deadline).toLocaleString()}</TableCell>
          <TableCell>{grade.marks}</TableCell>
          <TableCell>
            <Button
              variant="outlined"
              size="small"
              onClick={() => handleOpenFeedback(grade)}
            >
              View Feedback
            </Button>
          </TableCell>
          <TableCell>{grade.max_points}</TableCell>
        </TableRow>
      ));
  };

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
                <b>Submitted</b>
              </TableCell>
              <TableCell>
                <b>Due</b>
              </TableCell>
              <TableCell>
                <b>Score</b>
              </TableCell>
              <TableCell>
                <b>Feedback</b>
              </TableCell>
              <TableCell>
                <b>MaxPoints</b>
              </TableCell>
            </TableRow>
          </TableHead>
          <TableBody>{renderGrades(grades)}</TableBody>
          <TableRow>
            <TableCell colSpan={3}>Total</TableCell>
            <TableCell>0</TableCell>
            <TableCell>0.0%</TableCell>
          </TableRow>
        </Table>
      </TableContainer>

      {/* Feedback Dialog */}
      <Dialog open={feedbackDialogOpen} onClose={handleClose}>
        <DialogTitle>Feedback</DialogTitle>
        <DialogContent>
          <Typography>{selectedFeedback}</Typography>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Close</Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}
