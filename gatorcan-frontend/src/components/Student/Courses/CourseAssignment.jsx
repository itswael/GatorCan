import React, { useEffect } from "react";
import { useLocation } from "react-router-dom";
import StudentNavbar from "../StudentNavbar";
import { useState } from "react";
import { Box, Typography, Button } from "@mui/material";
import { useParams } from "react-router-dom";
import CourseNavbar from "./CourseNavbar";
import {
  fetchAssignmentDetails,
} from "../../../services/CourseService";
import s3Client from "../../../awsConfig";
import { PutObjectCommand, ListObjectsV2Command } from "@aws-sdk/client-s3";
import CheckCircleIcon from "@mui/icons-material/CheckCircle";

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
  const [submitting, setSubmitting] = useState(false);
  const [submittedData, setSubmittedData] = useState(null);
  const [fileInfo, setFileInfo] = useState(null);
  const [assignmentDetails, setAssignmentDetails] = useState([]);
  const [loading, setLoading] = useState(false);
  const [errMessage, setErrMessage] = useState("");

  useEffect(() => {
    const fetchData = async () => {
      const assignmentData = await fetchAssignmentDetails({
        id,
        assignment_id: assignment.id,
      });
      if (assignmentData != null) {
        console.log(assignmentData.assignments);
        setAssignmentDetails(assignmentData);
      } else {
        setErrMessage("Unable to fetch assignments, retry");
      }

      setLoading(false);
    };
    fetchData();
  }, []);

  const uploadFile = async (file_param) => {
    setUploading(true);

    try {
      const arrayBuffer = await file_param.arrayBuffer();
      const fileBuffer = new Uint8Array(arrayBuffer);

      const params = {
        Bucket: import.meta.env.VITE_S3_BUCKET_NAME,
        Key: file_param.name,
        Body: fileBuffer,
        ContentType: file_param.type,
      };

      const command = new PutObjectCommand(params);
      await s3Client.send(command);

      const fileUrl = `https://d8t4c0gsca730.cloudfront.net/${encodeURIComponent(
        params.Key
      )}`;
      setFile(file_param);
      setFileInfo({ url: fileUrl, name: file_param.name });

    } catch (error) {
      console.error("Upload failed:", error);
    } finally {
      setUploading(false);
    }
  };

  const handleFileChange = (event) => {
    const selectedFile = event.target.files[0];
    if (selectedFile) {
      uploadFile(selectedFile);
    }
  };

  const handleSubmit = async () => {
    if (!fileInfo) return;

    setSubmitting(true);

    // No re-fetch — update locally
    setSubmittedData({
      file_url: fileInfo.url,
      filename: fileInfo.name,
      submitted_at: new Date().toISOString(),
      grade: 0,
      feedback: "",
    });

    setSubmitting(false);
  };

  if (submittedData != null) {
    const submittedDate = new Date(submittedData.submitted_at).toLocaleString();
    const fileName = submittedData.file_url.split("/").pop();

    return (
      <div style={{ padding: "20px", fontFamily: "Arial, sans-serif" }}>
        <Typography variant="h4" fontWeight="bold">
          Submission
        </Typography>
        <Typography variant="body1" sx={{ mt: 2 }}>
          Submitted on: {submittedDate}
        </Typography>
        <Typography variant="body1" sx={{ mt: 1 }}>
          Grade: {submittedData.grade == 0 ? "N/A" : submittedData.grade}
        </Typography>
        <Typography variant="body1" sx={{ mt: 1 }}>
          Feedback: {submittedData.feedback || "N/A"}
        </Typography>
        {submittedData.file_url != null ? (
          <>
            <Button
              href={submittedData.file_url}
              target="_blank"
              rel="noopener noreferrer"
              sx={{ mt: 2 }}
              variant="outlined"
            >
              {fileName}
            </Button>
          </>
        ) : (
          <></>
        )}
        <Box sx={{ mt: 2 }}>
          <CheckCircleIcon sx={{ color: "green", fontSize: 32 }} />
        </Box>
      </div>
    );
  }

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
      <Typography variant="h4" fontWeight="bold">
        Submission
      </Typography>
      <Typography variant="body1" sx={{ mt: 2 }}>
        Guidelines: No additional details were added for this assignment.
      </Typography>

      <Box sx={{ display: "flex", alignItems: "center", marginTop: "20px" }}>
        <Button
          variant="contained"
          component="label"
          disabled={uploading}
          style={{
            backgroundColor: "rgb(29, 74, 124)",
            color: "white",
            marginRight: "10px",
          }}
        >
          {uploading ? "Uploading..." : "Upload File"}
          <input type="file" accept=".pdf" hidden onChange={handleFileChange} />
        </Button>

        {fileInfo && (
          <Button
            href={fileInfo.url}
            target="_blank"
            rel="noopener noreferrer"
            variant="text"
          >
            {fileInfo.name}
          </Button>
        )}
      </Box>

      <Typography variant="body2" sx={{ color: "gray", marginTop: "2px" }}>
        .pdf only
      </Typography>

      <Box sx={{ textAlign: "center", marginTop: "auto" }}>
        <Button
          onClick={handleSubmit}
          disabled={!fileInfo || submitting}
          style={{
            backgroundColor: "rgb(29, 74, 124)",
            color: "white",
            padding: "8px 16px",
            cursor: "pointer",
          }}
        >
          {submitting ? "Submitting..." : "Submit"}
        </Button>
      </Box>

      <Box sx={{ marginTop: "20px", textAlign: "left" }}>
        <Typography variant="body2" sx={{ color: "gray" }}>
          Note: Ensure your file is named appropriately before uploading.
        </Typography>
      </Box>
    </div>
  );
}
