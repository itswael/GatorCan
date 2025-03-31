import React from "react";
import { Paper } from "@mui/material";
import ListItem from "@mui/material/ListItem";
import ListItemText from "@mui/material/ListItemText";
import List from "@mui/material/List";
import { useNavigate, useParams } from "react-router-dom";

function CourseNavbar() {
    const navigate = useNavigate();

    const handleNavigation = (path) => {
      navigate(path, { replace: false });
    };

    let { id } = useParams();

  return (
    <>
      <Paper
        elevation={3}
        sx={{
          width: "130px",
          paddingLeft: 2,
          borderRight: "0px solid black",
        }}
      >
        <List>
          <ListItem
            disableGutters
            sx={{ paddingBottom: 0.5 }}
            onClick={() => handleNavigation("/student-course/" + id)}
          >
            <ListItemText
              primary="Home"
              sx={{
                color: "rgb(29, 74, 124)",
                cursor: "pointer",
                "&:hover": { textDecoration: "underline" },
              }}
            />
          </ListItem>
          <ListItem
            disableGutters
            sx={{ paddingBottom: 0.5 }}
            onClick={() =>
              handleNavigation("/student-course/" + id + "/announcements")
            }
          >
            <ListItemText
              primary="Announcements"
              sx={{
                color: "rgb(29, 74, 124)",
                cursor: "pointer",
                "&:hover": { textDecoration: "underline" },
              }}
            />
          </ListItem>
          <ListItem
            disableGutters
            sx={{ paddingBottom: 0.5 }}
            onClick={() =>
              handleNavigation("/student-course/" + id + "/assignments")
            }
          >
            <ListItemText
              primary="Assignments"
              sx={{
                color: "rgb(29, 74, 124)",
                cursor: "pointer",
                "&:hover": { textDecoration: "underline" },
              }}
            />
          </ListItem>
          <ListItem
            disableGutters
            sx={{ paddingBottom: 0.5 }}
            onClick={() =>
              handleNavigation("/student-course/" + id + "/grades")
            }
          >
            <ListItemText
              primary="Grades"
              sx={{
                color: "rgb(29, 74, 124)",
                cursor: "pointer",
                "&:hover": { textDecoration: "underline" },
              }}
            />
          </ListItem>
          <ListItem
            disableGutters
            sx={{ paddingBottom: 0.5 }}
            onClick={() =>
              handleNavigation("/student-course/" + id + "/syllabus")
            }
          >
            <ListItemText
              primary="Syllabus"
              sx={{
                color: "rgb(29, 74, 124)",
                cursor: "pointer",
                "&:hover": { textDecoration: "underline" },
              }}
            />
          </ListItem>
          <ListItem
            disableGutters
            sx={{ paddingBottom: 0.5 }}
            onClick={() =>
              handleNavigation("/student-course/" + id + "/discussions")
            }
          >
            <ListItemText
              primary="Discussions"
              sx={{
                color: "rgb(29, 74, 124)",
                cursor: "pointer",
                "&:hover": { textDecoration: "underline" },
              }}
            />
          </ListItem>
          <ListItem
            disableGutters
            sx={{ paddingBottom: 0.5 }}
            onClick={() => handleNavigation("/student-course/" + id + "/chat")}
          >
            <ListItemText
              primary="Chat"
              sx={{
                color: "rgb(29, 74, 124)",
                cursor: "pointer",
                "&:hover": { textDecoration: "underline" },
              }}
            />
          </ListItem>
        </List>
      </Paper>
    </>
  );
}

export default CourseNavbar;
