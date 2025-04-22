import React from "react";
import { Container, Typography, Box, Button, TextField } from "@mui/material";
import { useParams } from "react-router-dom";
import InstructorNavbar from "../InstructorNavbar";
import InstructorCourseNavbar from "./InstructorCourseNavbar";
import { useState } from "react";

const InstructorCourseSyllabus = () => {
  const { id } = useParams();

  const [editMode, setEditMode] = useState(false);
    const [text, setText] = useState(
      `Syllabus: Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. \nDonec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo.
  
  Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi. Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus.
  
  Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus.
  
  Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem. Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc,`
    );
    const [savedText, setSavedText] = useState(text);
  
    const handleSave = () => {
      setSavedText(text);
      setEditMode(false);
    };
  
    const handleCancel = () => {
      setText(savedText);
      setEditMode(false);
    };

  return (
    <>
      <InstructorNavbar />
      <Container sx={{ mt: 4 }}>

        <InstructorCourseNavbar activeTab="syllabus" />

        <Box mt={4}>
          {!editMode ? (
            <>
              <Typography
                variant="body1"
                color="text.secondary"
                textAlign="justify"
              >
                {savedText}
              </Typography>
              <Box mt={2}>
                <Button variant="outlined" onClick={() => setEditMode(true)}>
                  Edit
                </Button>
              </Box>
            </>
          ) : (
            <>
              <TextField
                multiline
                fullWidth
                minRows={8}
                value={text}
                onChange={(e) => setText(e.target.value)}
              />
              <Box mt={2} display="flex" gap={2}>
                <Button variant="contained" onClick={handleSave}>
                  Save
                </Button>
                <Button variant="outlined" onClick={handleCancel}>
                  Cancel
                </Button>
              </Box>
            </>
          )}
        </Box>
      </Container>
    </>
  );
};

export default InstructorCourseSyllabus;
