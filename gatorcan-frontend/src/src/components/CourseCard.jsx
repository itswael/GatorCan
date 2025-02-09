import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import CampaignIcon from '@mui/icons-material/Campaign';
import Typography from "@mui/material/Typography";
import EditNoteIcon from "@mui/icons-material/EditNote";
import MarkUnreadChatAltIcon from "@mui/icons-material/MarkUnreadChatAlt";
import FolderCopyIcon from "@mui/icons-material/FolderCopy";
import { Container } from "@mui/material";

export default function MediaCard({text1, text2, color}) {
  return (
    <Card sx={{ maxWidth: 250, margin: "20px" }} elevation={10}>
      <Container
        sx={{ backgroundColor: color, height: "150px" }}
      ></Container>
      <CardContent>
        <Typography gutterBottom variant="h8" component="div" color={color}>
          {text1}
        </Typography>
        <Typography variant="body2" sx={{ color: "text.secondary" }}>
          {text2}
        </Typography>
      </CardContent>
      <CardActions style={{ margin: "2 px", justifyContent: "space-between" }}>
        <CampaignIcon />
        <EditNoteIcon />
        <MarkUnreadChatAltIcon />
        <FolderCopyIcon />
      </CardActions>
    </Card>
  );
}
