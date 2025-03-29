import React from 'react'
import { useParams } from "react-router-dom";
import StudentNavbar from "../StudentNavbar";

function CourseDetails() {

  let { id } = useParams();
  console.log(id);
  return (<>
  <StudentNavbar />
  CourseDetails {id}</>);
}

export default CourseDetails