import React, { useEffect } from "react";
import { useLocation } from "react-router-dom";
import StudentNavbar from "../StudentNavbar";
import { useState } from "react";
import { Box, Typography, Button } from "@mui/material";
import { useParams } from "react-router-dom";
import CourseNavbar from "./CourseNavbar";
import { fetchAssignmentDetails } from "../../../services/CourseService";
import s3Client from "../../../awsConfig";
import { PutObjectCommand, ListObjectsV2Command } from "@aws-sdk/client-s3";

function CourseAssignment() {

    const location = useLocation();
    const { assignment } = location.state || {};

    const [loading, setLoading] = useState(false);
    const [errMessage, setErrMessage] = useState("");

    let { id } = useParams();

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
                    <Typography variant="h3">Assignments</Typography>
                    <AssignmentDetails id={id} assignment={assignment} />
                    <Submission id={id} assignment={assignment} />
                  </div>
                )}
              </div>
            </Box>
          </Box>
        </div>
      </>
    );
}

export default CourseAssignment;

function AssignmentDetails({ id, assignment }) {

    const [loading, setLoading] = useState(false);
    const [errMessage, setErrMessage] = useState("");
    const [assignmentDetails, setAssignmentDetails] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
          const assignmentData = await fetchAssignmentDetails({ id, assignment_id: assignment.id });
          if (assignmentData != null) {
            console.log(assignmentData.assignments);
            setAssignmentDetails(assignmentData);
          } else {
            setErrMessage("Unable to fetch assignments, retry");
          }
          setLoading(false);
        };
        // fetchData();
      }, []);

    const formattedDeadline = new Date(assignment.deadline).toLocaleString();
    const createdDate = new Date(assignment.created_at).toLocaleDateString();
    const updatedDate = new Date(assignment.updated_at).toLocaleDateString();
    
    return (
      <div style={{ padding: "20px", fontFamily: "Arial, sans-serif" }}>
        <Typography variant="h4" sx={{ fontWeight: "bold", textAlign: "left" }}>
          Details
        </Typography>
        <Box
          sx={{
            display: "flex",
            justifyContent: "space-between",
            alignItems: "center",
          }}
        >
          <Typography
            variant="h5"
            sx={{ fontWeight: "bold", textAlign: "left" }}
          >
            {assignment.title}
          </Typography>
          <Typography
            variant="body1"
            sx={{ fontWeight: "bold", textAlign: "right" }}
          >
            Max Points: {assignment.max_points}
          </Typography>
        </Box>

        <Box sx={{ marginTop: "10px", textAlign: "left" }}>
          <Typography variant="body2" sx={{ color: "gray" }}>
            Created at: {createdDate}
          </Typography>
          <Typography variant="body2" sx={{ color: "gray" }}>
            Updated at: {updatedDate}
          </Typography>
        </Box>

        <Typography
          variant="body1"
          sx={{ fontWeight: "bold", marginTop: "10px", textAlign: "left" }}
        >
          Deadline: {formattedDeadline}
        </Typography>

        <Typography
          variant="body1"
          sx={{ marginTop: "10px", textAlign: "left" }}
        >
          <strong>Description:</strong> {assignment.description}
        </Typography>
      </div>
    );
}

function Submission({ id, assignment }) {

  const [file, setFile] = useState(null);
  const [uploading, setUploading] = useState(false);
  const [message, setMessage] = useState("");

  const uploadFile = async (file_param) => {
  
      setUploading(true);
      setMessage("");
  
      // Convert file to an ArrayBuffer
      const arrayBuffer = await file_param.arrayBuffer();
      const fileBuffer = new Uint8Array(arrayBuffer); // Convert to Uint8Array
  
      const params = {
        Bucket: import.meta.env.VITE_S3_BUCKET_NAME,
        Key: file_param.name,
        Body: fileBuffer,
        ContentType: file_param.type,
      };
  
      try {
        const command = new PutObjectCommand(params);
        console.log("Sending upload request...");
        const response = await s3Client.send(command);
        console.log("Upload completed:", response);
        setMessage("Upload successful!");
      } catch (error) {
        setMessage(`Upload failed: ${error.message}..retry`);
      } finally {
        setUploading(false);
      }
    };

  const handleSubmit = async () => {
    // TODO: handle assignment submission
  }

  const handleFileChange = (event) => {
    setFile(event.target.files[0]);
    uploadFile(event.target.files[0]);
    console.log("Assignment submitted");
  };

  return (
    <div
      style={{
        padding: "20px",
        fontFamily: "Arial, sans-serif",
        minHeight: "50vh",
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
      }}
    >
      <Box
        sx={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
        }}
      >
        <Typography variant="h4" sx={{ fontWeight: "bold", textAlign: "left" }}>
          Submission
        </Typography>
      </Box>

      <Box sx={{ marginTop: "10px", textAlign: "left" }}>
        <Typography variant="body1">
          Guidelines: No additional details were added for this assignment.
        </Typography>
      </Box>
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          marginTop: "20px",
        }}
      >
        <Button
          variant="contained"
          component="label"
          style={{
            backgroundColor: "rgb(29, 74, 124)",
            color: "white",
            marginRight: "10px",
          }}
          disabled={uploading}
        >
          {uploading ? "Uploading..." : "Upload File"}
          <input
            type="file"
            accept=".pdf"
            hidden
            onChange={(event) => {
              setMessage("");
              handleFileChange(event);
            }}
          />
        </Button>
        {file && (
          <Typography variant="body2" sx={{ marginLeft: "10px" }}>
            {file.name}
          </Typography>
        )}
      </Box>
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          marginTop: "2px",
        }}
      >
        <Typography variant="body2" sx={{ color: "gray", marginLeft: "0px" }}>
          .pdf only
        </Typography>
      </Box>

      <Box sx={{ marginTop: "10px", textAlign: "left" }}>
        {message && (
          <Typography
            variant="body2"
            sx={{
              color: "red",
              marginTop: "0px",
            }}
          >
            {message}
          </Typography>
        )}
      </Box>

      <Box
        sx={{
          textAlign: "center",
          marginTop: "auto",
        }}
      >
        <Button
          onClick={handleSubmit}
          style={{
            backgroundColor: "rgb(29, 74, 124)",
            color: "white",
            border: "none",
            padding: "8px 16px",
            cursor: "pointer",
          }}
        >
          Submit
        </Button>
      </Box>

      <Box
        sx={{
          marginTop: "20px",
          textAlign: "left",
        }}
      >
        <Typography variant="body2" sx={{ color: "gray" }}>
          Note: Ensure your file is named appropriately before uploading.
        </Typography>
      </Box>
    </div>
  );
}