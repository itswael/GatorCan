import { Drawer, List, ListItem, ListItemIcon, ListItemText } from "@mui/material";
import AccountCircleIcon from "@mui/icons-material/AccountCircle";
import DashboardCustomizeRoundedIcon from "@mui/icons-material/DashboardCustomizeRounded";
import CollectionsBookmarkRoundedIcon from "@mui/icons-material/CollectionsBookmarkRounded";
import CalendarMonthRoundedIcon from "@mui/icons-material/CalendarMonthRounded";
import MailOutlineRoundedIcon from "@mui/icons-material/MailOutlineRounded";

function MyListItem({icon, name}) {
  return (
    <>
      <ListItem button sx={{ flexDirection: "column", alignItems: "center" }}>
        <ListItemIcon
          sx={{
            minWidth: "unset",
            display: "flex",
            justifyContent: "center",
          }}
        >
          {icon}
        </ListItemIcon>
        <ListItemText primary={name} />
      </ListItem>
    </>
  );
}

function Navbar() {
  return (
    <div>
      <Drawer
        variant="permanent"
        anchor="left"
        sx={{
          width: 120,
          flexShrink: 0,
          "& .MuiDrawer-paper": {
            width: 100,
            boxSizing: "border-box",
          },
        }}
        PaperProps={{
          sx: {
            backgroundColor: "rgb(29, 74, 124)",
            color: "white",
          },
        }}
      >
        <List>
          <ListItem button>
            <ListItemText primary="GatorCan" />
          </ListItem>
          <MyListItem
            icon={<AccountCircleIcon sx={{ fontSize: 40, color: "white" }} />}
            name="Profile"
          />
          <MyListItem
            icon={
              <DashboardCustomizeRoundedIcon
                sx={{ fontSize: 40, color: "white" }}
              />
            }
            name="Dashboard"
          />
          <MyListItem
            icon={
              <CollectionsBookmarkRoundedIcon
                sx={{ fontSize: 40, color: "white" }}
              />
            }
            name="Courses"
          />
          <MyListItem
            icon={
              <CalendarMonthRoundedIcon sx={{ fontSize: 40, color: "white" }} />
            }
            name="Calendar"
          />
          <MyListItem
            icon={
              <MailOutlineRoundedIcon sx={{ fontSize: 40, color: "white" }} />
            }
            name="Inbox"
          />
        </List>
      </Drawer>
    </div>
  );
}

export default Navbar;
