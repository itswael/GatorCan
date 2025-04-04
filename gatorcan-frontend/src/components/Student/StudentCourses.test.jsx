import React from "react";
import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import StudentCourses from "./StudentCourses";
import CourseService from "../../services/CourseService";
import "@testing-library/jest-dom";

jest.mock("../../services/CourseService"); // Mock API calls

describe("StudentCourses Component", () => {

  test("triggers enroll function when enroll button is clicked", async () => {
    // Mock API responses
    CourseService.fetchAllCourses.mockResolvedValue([
      {
        id: 1,
        name: "Test Course",
        description: "Test Description",
        created_at: "2025-01-01",
        updated_at: "2025-02-01",
      },
    ]);
    CourseService.fetchEnrolledCourses.mockResolvedValue([]);
    CourseService.enrollInCourse.mockResolvedValue({ success: true });

    render(
      <MemoryRouter>
        <StudentCourses />
      </MemoryRouter>
    );

    // Wait for courses to load
    await waitFor(() =>
      expect(screen.getByText("Test Course")).toBeInTheDocument()
    );

    // Click enroll button
    fireEvent.click(screen.getByRole("button", { name: /enroll/i }));

    // Verify API call for enrollment
    await waitFor(() => {
      expect(CourseService.enrollInCourse).toHaveBeenCalledWith(1);
    });
  });
});
