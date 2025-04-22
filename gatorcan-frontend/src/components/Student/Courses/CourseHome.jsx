import React from "react";
import StudentNavbar from "../StudentNavbar";
import { useState, useEffect } from "react";
import {
  Box,
  Typography,
} from "@mui/material";
import { useParams } from "react-router-dom";
import { fetchCourse } from "../../../services/CourseService";
import CourseNavbar from "./CourseNavbar";
import "react-chat-widget/lib/styles.css";
import ChatBox from "./ChatBox";


function CourseHome() {
  
  const [course_name, setCourseName] = useState("");
  const [course_id, setCourseId] = useState("");
  const [course_description, setCourseDescription] = useState("");
  const [instructor_name, setInstructorName] = useState("");
  const [instructor_email, setInstructorEmail] = useState("");
  const [loading, setLoading] = useState(true);
  const [errMessage, setErrMessage] = useState("");

  let { id } = useParams();
  console.log("id: " + id);

  const loremText =
    `Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. \nDonec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo.

Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi. Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus.

Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus.

Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem. Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc,`;

  useEffect(() => {
    const fetchData = async () => {
      const courseData = await fetchCourse({ id });
      console.log(courseData);
      if (courseData != null) {
        setCourseName(courseData.name);
        setCourseId(courseData.id);
        setCourseDescription(courseData.description);
        setInstructorName(courseData.instructorName);
        setInstructorEmail(courseData.instructorEmail);
      } else {
        setErrMessage("Unable to fetch course details, retry");
      }
      setLoading(false);
    };
    fetchData();
  }, []);

  return (
    <>
    <ChatBox course_id={id} />
      <StudentNavbar />
      <div style={{ marginLeft: "100px" }}>
        <Box sx={{ display: "flex", height: "100vh" }}>
          <CourseNavbar />
          <Box sx={{ flex: 1, padding: 2 }}>
            <div sx={{ display: "flex" }}>
              {errMessage != "" ? (
                <h4>{errMessage}</h4>
              ) : loading ? (
                <h4>Loading...</h4>
              ) : (
                <div>
                  <Typography variant="h3">Home</Typography>
                  <Typography variant="h4" sx={{ textAlign: "left" }}>
                    {"#" + course_id + "-" + course_name}
                  </Typography>
                  <Typography variant="body1" sx={{ textAlign: "left" }}>
                    <strong>Instructor:</strong> {instructor_name}
                  </Typography>
                  <Typography variant="body1" sx={{ textAlign: "left" }}>
                    <strong>Instructor email:</strong> {instructor_email}
                  </Typography>
                  <Typography
                    variant="body1"
                    sx={{ textAlign: "left", marginTop: 2 }}
                  >
                    {loremText}
                  </Typography>
                </div>
              )}
            </div>
          </Box>
        </Box>
      </div>
    </>
  );
}

export default CourseHome;