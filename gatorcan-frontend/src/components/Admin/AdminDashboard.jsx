import * as React from "react";

import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

import Grid from "@mui/material/Grid";
import PersonAddOutlinedIcon from '@mui/icons-material/PersonAddOutlined';
import PersonRemoveOutlinedIcon from "@mui/icons-material/PersonRemoveOutlined";
import CollectionsBookmarkOutlinedIcon from "@mui/icons-material/CollectionsBookmarkOutlined";
import EditNoteOutlinedIcon from "@mui/icons-material/EditNoteOutlined";
import BorderColorOutlinedIcon from "@mui/icons-material/BorderColorOutlined";
import BackspaceOutlinedIcon from "@mui/icons-material/BackspaceOutlined";
import PersonSearchOutlinedIcon from "@mui/icons-material/PersonSearchOutlined";
import ContentPasteSearchOutlinedIcon from "@mui/icons-material/ContentPasteSearchOutlined";
import QueryStatsOutlinedIcon from "@mui/icons-material/QueryStatsOutlined";

import UserRegistration from "./Tools/UserRegistration";
import UserDeletion from "./Tools/UserDeletion";
import UserRolesUpdation from "./Tools/UserRolesUpdation";
import AddCourse from "./Tools/AddCourse";
import ActivateCourse from "./Tools/ActivateCourse";

import { useState } from "react";

import AdminNavbar from "./AdminNavbar";

import {
  Card,
  CardActionArea,
  CardContent,
} from "@mui/material";

function AdminDashboard() {

  const [currPage, setCurrPage] = useState(-1);

  const tools = [
    [<PersonAddOutlinedIcon />, "LightSalmon", "Add User"],
    [<PersonRemoveOutlinedIcon />, "LightPink", "Delete User"],
    [<BorderColorOutlinedIcon />, "Salmon", "Edit User"],
    [<CollectionsBookmarkOutlinedIcon />, "Coral", "Add Course"],
    [<BackspaceOutlinedIcon />, "Gold", "Activate Course"],
    [<EditNoteOutlinedIcon />, "Tomato", "Delete Course"],
    [<PersonSearchOutlinedIcon />, "PapayaWhip", "View Users"],
    [<ContentPasteSearchOutlinedIcon />, "Khaki", "View Courses"],
    [<QueryStatsOutlinedIcon />, "Crimson", "Statistics"],
  ];

  return (
    <div>
      <AdminNavbar />
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        minHeight="100vh"
      >
        {currPage == 1 ? (
          <UserRegistration setCurrPage={setCurrPage} />
        ) : currPage == 2 ? (
          <UserDeletion setCurrPage={setCurrPage} />
        ) : currPage == 3 ? (
          <UserRolesUpdation setCurrPage={setCurrPage} />
        ) : currPage == 4 ? (
          <AddCourse setCurrPage={setCurrPage} />
        ) : currPage == 5 ? (
          <ActivateCourse setCurrPage={setCurrPage} />
        ) : (
          <Grid container spacing={5} maxWidth={600}>
            {tools.map((_, index) => (
              <Grid
                item
                key={index}
                xs={4}
                display="flex"
                justifyContent="center"
              >
                {" "}
                <Card
                  sx={{
                    width: 120,
                    height: 120,
                    backgroundColor: tools[index][1],
                  }}
                  onClick={() => setCurrPage(index + 1)}
                >
                  <CardActionArea sx={{ height: "100%" }}>
                    <CardContent
                      sx={{
                        display: "flex",
                        flexDirection: "column",
                        alignItems: "center",
                        justifyContent: "center",
                        height: "100%",
                      }}
                    >
                      {tools[index][0]}
                      <Typography variant="caption">
                        {tools[index][2]}
                      </Typography>
                    </CardContent>
                  </CardActionArea>
                </Card>
              </Grid>
            ))}
          </Grid>
        )}
      </Box>
    </div>
  );
}

export default AdminDashboard;
