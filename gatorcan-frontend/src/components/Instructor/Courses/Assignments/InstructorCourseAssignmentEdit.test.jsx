import React from "react";
import { render, screen, waitFor, fireEvent } from "@testing-library/react";
import InstructorCourseAssignmentEdit from "./InstructorCourseAssignmentEdit";
import CourseService from "../../../../services/CourseService";
import InstructorService from "../../../../services/InstructorService";
import { MemoryRouter, Route, Routes } from "react-router-dom";
import "@testing-library/jest-dom";

jest.mock("../../../../services/CourseService");
jest.mock("../../../../services/InstructorService");

const mockAssignment = {
  title: "Test Assignment",
  description: "This is a test",
  deadline: "2025-05-01T12:00:00Z",
  max_points: 100,
};

describe("InstructorCourseAssignmentEdit Component", () => {
  test("renders form fields with correct values", async () => {
    CourseService.fetchAssignmentDetails.mockResolvedValue({
      assignments: mockAssignment,
    });

    render(
      <MemoryRouter
        initialEntries={["/instructor-courses/1/assignments/1/edit-assignment"]}
      >
        <Routes>
          <Route
            path="/instructor-courses/:id/assignments/:assignment_id/edit-assignment"
            element={<InstructorCourseAssignmentEdit />}
          />
        </Routes>
      </MemoryRouter>
    );

    await waitFor(() =>
      expect(screen.getByDisplayValue("Test Assignment")).toBeInTheDocument()
    );
    expect(screen.getByDisplayValue("This is a test")).toBeInTheDocument();
    expect(screen.getByDisplayValue("100")).toBeInTheDocument(); // maxPoints
  });

  test("shows not found message for missing assignment", async () => {
    CourseService.fetchAssignmentDetails.mockResolvedValue({
      assignments: null,
    });

    render(
      <MemoryRouter
        initialEntries={[
          "/instructor-courses/1/assignments/999/edit-assignment",
        ]}
      >
        <Routes>
          <Route
            path="/instructor-courses/:id/assignments/:assignment_id/edit-assignment"
            element={<InstructorCourseAssignmentEdit />}
          />
        </Routes>
      </MemoryRouter>
    );

    await waitFor(() =>
      expect(screen.getByText("Assignment not found.")).toBeInTheDocument()
    );
  });
});
