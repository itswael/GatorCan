// InstructorProfile.jsx
import React, { useState, useEffect, useRef } from "react";
import {
  Avatar,
  Typography,
  Container,
  Box,
  Button,
  CircularProgress,
  Paper,
  Input,
} from "@mui/material";
import Grid from "@mui/material/Grid2";
import InstructorNavbar from "./InstructorNavbar";
import { getUserDetails, resetPassword } from "../../services/UserService";

const InstructorProfile = () => {
  const username = localStorage.getItem("username");
  const [email, setEmail] = useState("");
  const [first, setFirst] = useState("John");
  const [last, setLast] = useState("Doe");
  const [phone, setphone] = useState("+1 355-442-4452");
  const [line, setline] = useState("123 Main St");
  const [line2, setline2] = useState("Apt 456");
  const [photoUrl, setphotoUrl] = useState(
    "https://microbiology.ucr.edu/sites/default/files/styles/form_preview/public/blank-profile-pic.png?itok=4teBBoet"
  );
  const [loading, setloading] = useState(true);
  const [resetPwd, setResetPwd] = useState(false);
  const [errMsg, setErrMsg] = useState("");

  const fetchUserDetails = async () => {
    setloading(true);
    try {
      const response = await getUserDetails(username);
      if (response["success"]) {
        setEmail(response["data"]?.email);
      } else {
        setErrMsg(response["message"] || "Unknown error occurred");
      }
    } catch (e) {
      setErrMsg(e.message);
    }
    setloading(false);
  };

  useEffect(() => {
    fetchUserDetails();
  }, []);

  useEffect(() => {
    setErrMsg("");
  }, [loading]);

  return (
    <>
      <InstructorNavbar />
        <Container maxWidth="sm" sx={{ mt: 4, textAlign: "center" }}>
        <Typography variant="h5" gutterBottom>
            Profile
        </Typography>
        {resetPwd ? (
            <ResetPassword setResetPwd={setResetPwd} />
        ) : loading ? (
            <CircularProgress />
        ) : (
            <>
            <Avatar
                alt="Profile Picture"
                src={photoUrl}
                sx={{ width: 120, height: 120, mx: "auto", mb: 2 }}
            />

            <Box sx={{ mt: 3 }}>
                <Typography variant="h6" color="gray">
                Account Details
                </Typography>
                <Typography>Username: {username}</Typography>
                <Typography>Email: {email}</Typography>
            </Box>

            <Box sx={{ mt: 3 }}>
                <Typography variant="h6" color="gray">
                Personal Details
                </Typography>
                <Typography>First Name: {first}</Typography>
                <Typography>Last Name: {last}</Typography>
                <Typography>Phone: {phone}</Typography>
                <Typography>Address Line 1: {line}</Typography>
                <Typography>Address Line 2: {line2}</Typography>
            </Box>

            <Box sx={{ mt: 3 }}>
                <Button onClick={() => setResetPwd(true)}>
                Reset Password
                </Button>
            </Box>
            </>
        )}
        </Container>
    </>
  );
};

export default InstructorProfile;

const ResetPassword = ({ setResetPwd }) => {
  const username = localStorage.getItem("username");
  const errRef = useRef();
  const [pwd, setPwd] = useState("");
  const [rePwd, setRePwd] = useState("");
  const [oldPwd, setOldPwd] = useState("");
  const [errMsg, setErrMsg] = useState("");
  const [displaySuccess, setDisplaySuccess] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (pwd !== rePwd) {
      setErrMsg("Passwords do not match");
      return;
    }
    try {
      const response = await resetPassword(oldPwd, pwd);
      if (!response.success) {
        setErrMsg(response.message);
      } else {
        setDisplaySuccess("Password updated successfully");
      }
    } catch (error) {
      setErrMsg(error.response?.data?.message || "Unknown error");
    }
  };

  const inputStyle = { margin: "12px auto" };
  const paperStyle = { padding: 20, width: 400, margin: "19px auto" };

  return displaySuccess ? (
    <Paper elevation={12} style={paperStyle}>
      <Grid align="center">
        <p>{displaySuccess}</p>
        <Button
          variant="contained"
          fullWidth
          onClick={() => setResetPwd(false)}
          sx={{ mt: 2 }}
        >
          DONE
        </Button>
      </Grid>
    </Paper>
  ) : (
    <form onSubmit={handleSubmit}>
      <Paper elevation={12} style={paperStyle}>
        <Grid align="center">
          <h2>Reset Password</h2>
        </Grid>
        <Input
          fullWidth
          disabled
          value={username}
          placeholder="Username"
          style={inputStyle}
        />
        <Input
          type="password"
          placeholder="Old Password"
          fullWidth
          required
          value={oldPwd}
          onChange={(e) => setOldPwd(e.target.value)}
          style={inputStyle}
        />
        <Input
          type="password"
          placeholder="New Password"
          fullWidth
          required
          value={pwd}
          onChange={(e) => setPwd(e.target.value)}
          style={inputStyle}
        />
        <Input
          type="password"
          placeholder="Repeat Password"
          fullWidth
          required
          value={rePwd}
          onChange={(e) => setRePwd(e.target.value)}
          style={inputStyle}
        />
        <Button type="submit" variant="contained" fullWidth sx={{ mt: 2 }}>
          Submit
        </Button>
        <Box sx={{ textAlign: "center", mt: 2 }}>
          <Button onClick={() => setResetPwd(false)}>Back</Button>
        </Box>
        <Typography ref={errRef} color="error">
          {errMsg}
        </Typography>
      </Paper>
    </form>
  );
};
