import React from "react";
import { render, screen, waitFor } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import CourseHome from "./CourseHome";
import { fetchCourse } from "../../../services/CourseService";
import "@testing-library/jest-dom";

jest.mock("../../../services/CourseService");

describe("CourseHome Component", () => {
  // Mocking the fetchCourse function to return mock data
  beforeEach(() => {
    fetchCourse.mockResolvedValue({
      id: "101",
      name: "Course Name",
      description: "Course Description",
      instructorName: "Instructor Name",
      instructorEmail: "instructor@example.com",
    });
  });

  test("displays 'Home' text in the document", async () => {
    render(
      <MemoryRouter>
        <CourseHome />
      </MemoryRouter>
    );

    // Wait for the 'Home' text to be rendered in the document
    await waitFor(() => expect(screen.getByText("Home")).toBeInTheDocument());
  });

});
