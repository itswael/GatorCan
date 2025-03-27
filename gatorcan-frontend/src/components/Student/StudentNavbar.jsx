import React from "react";

import { useState, useEffect } from "react";

import {
  Drawer,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Box,
  ClickAwayListener,
  Paper,
} from "@mui/material";
import { useNavigate } from "react-router-dom";
import AccountCircleIcon from "@mui/icons-material/AccountCircle";
import DashboardCustomizeRoundedIcon from "@mui/icons-material/DashboardCustomizeRounded";
import CollectionsBookmarkRoundedIcon from "@mui/icons-material/CollectionsBookmarkRounded";
import CalendarMonthRoundedIcon from "@mui/icons-material/CalendarMonthRounded";
import MailOutlineRoundedIcon from "@mui/icons-material/MailOutlineRounded";
import LogoutRoundedIcon from "@mui/icons-material/Logout";
import CourseService from "../../services/CourseService";

function MyListItem({ icon, name, path, handleNavigate }) {
  const [open, setOpen] = useState(false);
  const [enrolledCourses, setEnrolledCourses] = useState([]);
  const [loadingEnrolledCourses, setLoadingEnrolledCourses] = useState(false);

  const handleToggle = () => {
    setOpen((prev) => !prev);
  };

  const handleClose = () => {
    if (open) setOpen(false);
  };

  const loadEnrolledCourses = async () => {
    setLoadingEnrolledCourses(true);
    try {
      const courses = await CourseService.fetchEnrolledCourses();
      setEnrolledCourses(courses);
    } catch (error) {
      console.error("Error fetching courses:", error);
    } finally {
      setLoadingEnrolledCourses(false);
    }
  };

  useEffect(() => {
    if (open) {
      loadEnrolledCourses();
    }
  }, [open]);

  if (name === "Courses") {
    return (
      <ClickAwayListener onClickAway={handleClose}>
        <Box sx={{ position: "relative", width: "100%", overflow: "visible" }}>
          <ListItem
            button
            sx={{
              flexDirection: "column",
              alignItems: "center",
              position: "relative",
              zIndex: 20,
            }}
            onClick={handleToggle}
          >
            {icon}
            <ListItemText
              primary={name}
              sx={{ "& .MuiTypography-root": { fontSize: "0.7em" } }}
            />
          </ListItem>

          {open && (
            <Paper
              elevation={6}
              sx={{
                position: "absolute",
                left: "100%",
                top: "0%",
                minWidth: "180px",
                backgroundColor: "rgb(29, 74, 124)",
                color: "orange",
                border: "1px solid white",
                zIndex: 30,
                padding: "10px",
                borderRadius: "8px",
                boxShadow: "0px 4px 10px rgba(0, 0, 0, 0.2)",
              }}
            >
              <List>
                {loadingEnrolledCourses ? (
                  <ListItem>
                    <ListItemText
                      primary="Loading..."
                      sx={{ padding: "5px", color: "white" }}
                    />
                  </ListItem>
                ) : enrolledCourses.length === 0 ? (
                  <ListItem>
                    <ListItemText
                      primary="No courses enrolled"
                      sx={{ padding: "5px", color: "white" }}
                    />
                  </ListItem>
                ) : (
                  enrolledCourses.map((course, index) => (
                    <ListItem
                      button
                      key={index}
                      onClick={() => handleNavigate(`/course/${course.id}`)}
                    >
                      <ListItemText
                        primary={course.name}
                        sx={{ padding: "5px", color: "white" }}
                      />
                    </ListItem>
                  ))
                )}
                <ListItem
                  button
                  onClick={() => handleNavigate("/student-courses")}
                >
                  <ListItemText
                    primary="View all courses"
                    sx={{ padding: "5px", color: "white" }}
                  />
                </ListItem>
              </List>
            </Paper>
          )}
        </Box>
      </ClickAwayListener>
    );
  } else {
    return (
      <ListItem
        button
        sx={{ flexDirection: "column", alignItems: "center" }}
        onClick={() => handleNavigate(path)}
      >
        <ListItemIcon
          sx={{
            minWidth: "unset",
            display: "flex",
            justifyContent: "center",
          }}
        >
          {icon}
        </ListItemIcon>
        <ListItemText
          primary={name}
          sx={{ "& .MuiTypography-root": { fontSize: "0.7em" } }}
        />
      </ListItem>
    );
  }
}

function StudentNavbar() {

    const handleLogout = () => {
      localStorage.clear();
      navigate("/login", { replace: true });
    };

    const navigate = useNavigate();

    const handleNavigate = (path) => {
      navigate(path, { replace: false });
    };

  return (
    <Drawer
      variant="permanent"
      anchor="left"
      sx={{
        width: 120,
        flexShrink: 0,
        overflow: "visible",
        "& .MuiDrawer-paper": {
          width: 100,
          overflow: "visible",
          boxSizing: "border-box",
          display: "flex",
          flexDirection: "column",
          justifyContent: "space-between", // Pushes logout to the bottom
        },
      }}
      PaperProps={{
        sx: {
          backgroundColor: "rgb(29, 74, 124)",
          color: "white",
        },
      }}
    >
      {/* Top Menu Items */}
      <Box sx={{ flexGrow: 1 }}>
        <List>
          <ListItem button>
            <ListItemText primary="GatorCan" />
          </ListItem>
          <MyListItem
            icon={<AccountCircleIcon sx={{ fontSize: 40, color: "white" }} />}
            name="Profile"
            path="/student-profile"
            handleNavigate={handleNavigate}
          />
          <MyListItem
            icon={
              <DashboardCustomizeRoundedIcon
                sx={{ fontSize: 40, color: "white" }}
              />
            }
            name="Dashboard"
            path="/student-dashboard"
            handleNavigate={handleNavigate}
          />
          <MyListItem
            icon={
              <CollectionsBookmarkRoundedIcon
                sx={{ fontSize: 40, color: "white" }}
              />
            }
            name="Courses"
            path="/student-courses"
            handleNavigate={handleNavigate}
          />
          <MyListItem
            icon={
              <CalendarMonthRoundedIcon sx={{ fontSize: 40, color: "white" }} />
            }
            name="Calendar"
            path="/student-calendar"
            handleNavigate={handleNavigate}
          />
          <MyListItem
            icon={
              <MailOutlineRoundedIcon sx={{ fontSize: 40, color: "white" }} />
            }
            name="Inbox"
            path="/student-inbox"
            handleNavigate={handleNavigate}
          />
        </List>
      </Box>

      {/* Logout Button at the Bottom */}
      <List>
        <ListItem
          button
          sx={{ flexDirection: "column", alignItems: "center" }}
          onClick={handleLogout} // Handle click event
        >
          <ListItemIcon
            sx={{
              minWidth: "unset",
              display: "flex",
              justifyContent: "center",
            }}
          >
            <LogoutRoundedIcon sx={{ fontSize: 40, color: "white" }} />
          </ListItemIcon>
          <ListItemText primary="Logout" />
        </ListItem>
      </List>
    </Drawer>
  );
}

export default StudentNavbar;
