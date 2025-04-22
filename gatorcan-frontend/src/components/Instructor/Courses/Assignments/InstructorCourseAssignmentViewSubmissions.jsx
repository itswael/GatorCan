import React, { useState, useEffect } from "react";
import {
  Container,
  Typography,
  Box,
  Paper,
  Accordion,
  AccordionSummary,
  AccordionDetails,
  TextField,
  Button,
  Link,
} from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import { useParams } from "react-router-dom";
import InstructorNavbar from "../../InstructorNavbar";
import InstructorCourseNavbar from "../InstructorCourseNavbar";
import dayjs from "dayjs";
import { fetchAssignmentSubmissions } from "../../../../services/CourseService";
import InstructorService from "../../../../services/InstructorService";

const InstructorCourseAssignmentViewSubmissions = () => {
  const { id, assignment_id } = useParams();
  const [submissions, setSubmissions] = useState([]);
  const [expandedIndex, setExpandedIndex] = useState(null);
  const maxPoints = 100;

  const handleExpand = (index) => {
    setExpandedIndex((prev) => (prev === index ? null : index));
  };

  const handleChange = (index, field, value) => {
    const updated = [...submissions];
    updated[index][field] = value;
    setSubmissions(updated);
  };

  const handleSubmit = async (index) => {
    const submission = submissions[index];
    const numericMarks = parseFloat(submission.marks);

    if (isNaN(numericMarks) || numericMarks <= 0 || numericMarks > maxPoints) {
      alert(
        `Marks must be a number between 1 and ${maxPoints}. Please correct it.`
      );
      return;
    }

    const payload = {
      assignment_id: parseInt(assignment_id),
      grade: numericMarks,
      feedback: submission.feedback,
      course_id: parseInt(id),
      user_id: submission.student_id,
    };
    
    const result = await InstructorService.gradeAssignment({
      cid: id,
      aid: assignment_id,
      grades: payload,
    });
    
    if (result.success) {
      const data = submissions;
      data[index].isGraded = true;
      setSubmissions(data);

      setExpandedIndex(null);

      console.log("Submitted grading:", {
        student_id: submission.student_id,
        marks: numericMarks,
        feedback: submission.feedback,
      });
    } else {
      window.alert(result.error);
    }
  };

  const sortedSubmissions = [...submissions].sort(
    (a, b) => dayjs(a.submitted_at).valueOf() - dayjs(b.submitted_at).valueOf()
  );

  useEffect(() => {
    
    const fetchData = async () => {
      try {
        const response = await fetchAssignmentSubmissions({cid: id, aid: assignment_id});
        if (response.success) {
          const formattedData = response.data.map((submission) => ({
            student_id: submission.user_id,
            student_name: submission.username,
            file_name: submission.file_name,
            file_url: submission.file_url,
            submitted_at: new Date(submission.submitted_at).toISOString().split(".")[0],
            marks: submission.grade == 0 ? "" : submission.grade,
            feedback: submission.feedback,
            isGraded: submission.grade > 0,
          }));
          setSubmissions(formattedData);
        }
      }
      catch (error) {
        console.error("Error fetching submissions:", error);
      }
    }
  
    fetchData();
  }, []);
  


  return (
    <>
      <InstructorNavbar />
      <Container sx={{ mt: 4 }}>

        <InstructorCourseNavbar activeTab="assignments" />

        <Box mt={4}>
          {sortedSubmissions.map((submission, index) => {
            const isGraded = submission.isGraded;

            return (
              <Accordion
                key={index}
                expanded={expandedIndex === index}
                onChange={() => handleExpand(index)}
                sx={{
                  backgroundColor: isGraded ? "#e8f5e9" : "inherit",
                }}
              >
                <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                  <Typography>
                    {submission.student_name} (#{submission.student_id})
                    {isGraded && (
                      <Typography
                        component="span"
                        sx={{ fontSize: "0.85rem", color: "green", ml: 1 }}
                      >
                        Graded
                      </Typography>
                    )}
                  </Typography>
                </AccordionSummary>
                <AccordionDetails>
                  <Paper elevation={2} sx={{ p: 2 }}>
                    <Typography variant="body2">
                      <strong>File:</strong>{" "}
                      <Link
                        href={submission.file_url}
                        target="_blank"
                        rel="noopener"
                      >
                        {submission.file_name}
                      </Link>
                    </Typography>
                    <Typography variant="body2">
                      <strong>Submitted On:</strong>{" "}
                      {dayjs(submission.submitted_at).format(
                        "MMM D, YYYY h:mm A"
                      )}
                    </Typography>

                    <Box display="flex" alignItems="center" gap={1} mt={2}>
                      <Typography>Marks:</Typography>
                      <TextField
                        type="number"
                        size="small"
                        value={submission.marks}
                        onChange={(e) =>
                          handleChange(index, "marks", e.target.value)
                        }
                        sx={{ width: 80 }}
                        inputProps={{ min: 0, max: maxPoints }}
                        error={
                          submission.marks !== "" &&
                          (isNaN(submission.marks) ||
                            submission.marks <= 0 ||
                            submission.marks > maxPoints)
                        }
                        helperText={
                          submission.marks !== "" &&
                          (isNaN(submission.marks) ||
                            submission.marks <= 0 ||
                            submission.marks > maxPoints)
                            ? `Marks must be 1–${maxPoints}`
                            : ""
                        }
                      />
                      <Typography variant="body2">/ {maxPoints}</Typography>
                    </Box>

                    <TextField
                      label="Feedback"
                      multiline
                      fullWidth
                      minRows={2}
                      sx={{ mt: 2 }}
                      value={submission.feedback}
                      onChange={(e) =>
                        handleChange(index, "feedback", e.target.value)
                      }
                    />

                    <Box mt={2}>
                      <Button
                        variant="contained"
                        onClick={() => handleSubmit(index)}
                      >
                        Submit
                      </Button>
                    </Box>
                  </Paper>
                </AccordionDetails>
              </Accordion>
            );
          })}
        </Box>
      </Container>
    </>
  );
};

export default InstructorCourseAssignmentViewSubmissions;
