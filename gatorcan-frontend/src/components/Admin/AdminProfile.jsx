import Avatar from "@mui/material/Avatar";
import { Typography, Container, Box, Button } from "@mui/material";
import AdminNavbar from "./AdminNavbar";
import { getUserDetails, resetPassword } from "../../services/UserService";
import { useState, useEffect, useRef } from "react";
import CircularProgress from "@mui/material/CircularProgress";

import Grid from "@mui/material/Grid2";
import Paper from "@mui/material/Paper";
import Input from "@mui/material/Input";

const AdminProfile = () => {
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

  

  return (
    <>
      <AdminNavbar />
      <h1>Profile</h1>
      <hr />
      <Container maxWidth="sm" sx={{ mt: 4 }}>
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
                  <Typography variant="body1">Username: {username}</Typography>
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
    </>
  );
};

export default AdminProfile;



