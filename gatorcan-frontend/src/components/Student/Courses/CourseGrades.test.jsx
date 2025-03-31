import React from "react";
import { render, screen, waitFor } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import CourseGrades from "./CourseGrades";
import { fetchAssignments } from "../../../services/CourseService";
import "@testing-library/jest-dom";

jest.mock("../../../services/CourseService");

describe("CourseGrades Component", () => {
  // Mocking the fetchAssignments function to return mock data
  beforeEach(() => {
    fetchAssignments.mockResolvedValue({
      assignments: [
        {
          title: "Assignment 1",
          deadline: "2025-04-15T12:00:00Z",
          max_points: 100,
        },
        {
          title: "Assignment 2",
          deadline: "2025-03-15T12:00:00Z",
          max_points: 80,
        },
      ],
    });
  });

  test("fetches assignments and displays total row", async () => {
    render(
      <MemoryRouter>
        <CourseGrades />
      </MemoryRouter>
    );

    // Wait for the assignments to be fetched
    await waitFor(() => expect(fetchAssignments).toHaveBeenCalledTimes(1));

    // Wait for the total row to be present in the document
    await waitFor(() => expect(screen.getByText("Total")).toBeInTheDocument());
  });
});
