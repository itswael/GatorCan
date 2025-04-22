// InstructorNavbar.jsx
import React from "react";
import { AppBar, Toolbar, Typography, Button, Box } from "@mui/material";
import { useNavigate } from "react-router-dom";

const InstructorNavbar = () => {
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.clear();
    navigate("/login", { replace: true });
  };

  return (
    <AppBar position="static" color="primary">
      <Toolbar sx={{ justifyContent: "space-between" }}>
        <Button color="inherit" onClick={() => navigate("/instructor-dashboard")}>
          Instructor Dashboard
        </Button>
        <Box>
          <Button
            color="inherit"
            onClick={() => navigate("/instructor-profile")}
          >
            Profile
          </Button>
          <Button color="inherit" onClick={() => navigate("/instructor-chat")}>
            Chat
          </Button>
          <Button color="inherit" onClick={handleLogout}>
            Logout
          </Button>
        </Box>
      </Toolbar>
    </AppBar>
  );
};

export default InstructorNavbar;
