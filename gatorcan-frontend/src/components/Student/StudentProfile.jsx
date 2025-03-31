import Avatar from "@mui/material/Avatar";
import { Typography, Container, Box, Button } from "@mui/material";
import StudentNavbar from "./StudentNavbar";
import { getUserDetails, resetPassword } from "../../services/UserService";
import { useState, useEffect, useRef } from "react";
import CircularProgress from "@mui/material/CircularProgress";

import Grid from "@mui/material/Grid2";
import Paper from "@mui/material/Paper";
import Input from "@mui/material/Input";
import React from "react";

const StudentProfile = () => {

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
        if (response["message"] != undefined) {
          setErrMsg(response.message);
        } else {
          setErrMsg("Unknow error occured");
        }
      } 
    } catch (e) {
      setErrMsg(e.message);
    }
    setloading(false);
  }

  useEffect( () => {
    const fetchData = async () => {
      await fetchUserDetails();
    }
    fetchData();
  }, []);

  useEffect(() => {
    setErrMsg("");
  }, [loading]);
  
  return (
    <>
      <StudentNavbar />
      <div style={{ marginLeft: "120px" }}>
        <h1>Profile</h1>
        <hr />
        <br></br>
        <Container maxWidth="sm" sx={{ mt: 4, textAlign: "center" }}>
          {resetPwd ? (
            <ResetPassword setResetPwd={setResetPwd} />
          ) : (
            <div>
              {loading ? (
                <Loading />
              ) : (
                <div>
                  <Avatar
                    alt="Profile Picture"
                    src={photoUrl}
                    sx={{ width: 120, height: 120, mx: "auto", mb: 2 }}
                  />

                  <Box sx={{ textAlign: "center", mt: 3 }}>
                    <Typography
                      variant="h6"
                      sx={{ fontWeight: "bold", color: "gray" }}
                    >
                      Account Details
                    </Typography>
                    <Typography variant="body1">
                      Username: {username}
                    </Typography>
                    <Typography variant="body1">Email: {email}</Typography>
                  </Box>

                  <Box sx={{ textAlign: "center", mt: 3 }}>
                    <Typography
                      variant="h6"
                      sx={{ fontWeight: "bold", color: "gray" }}
                    >
                      Personal Details
                    </Typography>
                    <Typography variant="body1">First Name: {first}</Typography>
                    <Typography variant="body1">Last Name: {last}</Typography>
                    <Typography variant="body1">Phone: {phone}</Typography>
                    <Typography variant="body1">
                      Address Line 1: {line}
                    </Typography>
                    <Typography variant="body1">
                      Address Line 2: {line2}
                    </Typography>
                  </Box>

                  <Box sx={{ textAlign: "center", mt: 3 }}>
                    <Button onClick={() => setResetPwd(true)}>
                      Reset Password
                    </Button>
                  </Box>
                </div>
              )}
            </div>
          )}
        </Container>
      </div>
    </>
  );
};

export default StudentProfile;

const Loading = () => {
  return (
    <div>
      <CircularProgress></CircularProgress>
    </div>
  );
};

const ResetPassword = ({setResetPwd}) => {
  const username = localStorage.getItem("username");
  const errRef = useRef();

  const [pwd, setPwd] = useState("");
  const [rePwd, setRePwd] = useState("");
  const [oldPwd, setOldPwd] = useState("")
  const [errMsg, setErrMsg] = useState("");

  const [displaySuccess, setDisplaySuccess] = useState("");

  const btnstyle = { backgroundColor: "#1B6DA1", margin: "20px 0" };
  const paperStyle = {
    padding: 20,
    width: 400,
    margin: "19px auto",
  };
  const inputStyle = { margin: "12px auto" };
  const errorStyle = { color: "red" };

  useEffect(() => {
    setErrMsg("");
  }, [pwd, rePwd]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    // validations
    if (pwd != rePwd) {
      setErrMsg("Passwords do not match");
      return;
    }

    try {
      const response = await resetPassword(oldPwd, pwd);
      console.log("Login API Successful:", response);
      let success = response["success"];
      if (!success) {
        console.log(response["message"]);
        setErrMsg(response["message"]);
      } else {
        setDisplaySuccess("Password updated successfully");
      }
    } catch (error) {
      console.log(error);
      setErrMsg(error.response?.data?.message || "Unknown error");
    }
  };

  return displaySuccess !== "" ? (
    <Paper elevation={12} style={paperStyle}>
      <Grid align="center">
        <p>{displaySuccess}</p>
      </Grid>
      <Button
        style={btnstyle}
        color="primary"
        variant="contained"
        fullWidth
        onClick={() => {
          setResetPwd(false);
        }}
      >
        DONE
      </Button>
    </Paper>
  ) : (
    <div>
      <Grid>
        <form onSubmit={handleSubmit}>
          <Paper elevation={12} style={paperStyle}>
            <Grid align="center">
              <h2 data-testid="cypress-title">Reset Password</h2>
            </Grid>
            <Input
              type="text"
              id="username"
              autoComplete="off"
              value={username}
              required
              style={inputStyle}
              placeholder="Username"
              fullWidth
              disabled
            />
            <Input
              type="password"
              id="oldpassword"
              onChange={(e) => setOldPwd(e.target.value)}
              value={oldPwd}
              required
              placeholder="Old Password"
              fullWidth
              style={inputStyle}
            />
            <Input
              type="password"
              id="password"
              onChange={(e) => setPwd(e.target.value)}
              value={pwd}
              required
              placeholder="Password"
              fullWidth
              style={inputStyle}
            />
            <Input
              type="password"
              id="repassword"
              onChange={(e) => setRePwd(e.target.value)}
              value={rePwd}
              required
              placeholder="Repeat Password"
              fullWidth
              style={inputStyle}
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
            <Box sx={{ textAlign: "center", mt: 3 }}>
              <Button onClick={() => setResetPwd(false)}>Back</Button>
            </Box>
          </Paper>
        </form>
        <p
          ref={errRef}
          role="alert"
          className={errMsg ? "errmsg" : "offscreen"}
          aria-live="assertive"
          style={errorStyle}
        >
          {errMsg}
        </p>
      </Grid>
    </div>
  );
};
