import React from "react";
import { render, screen, waitFor } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import CourseAssignments from "./CourseAssignments";
import { fetchAssignments } from "../../../services/CourseService";
import "@testing-library/jest-dom";

// Mock the fetchAssignments function
jest.mock("../../../services/CourseService");

describe("CourseAssignments Component", () => {
  // Mocking the fetchAssignments function to return mock data
  beforeEach(() => {
    fetchAssignments.mockResolvedValue({
      assignments: [
        {
          id: "1",
          title: "Upcoming Assignment",
          deadline: "2025-04-15T12:00:00Z",
          max_points: 100,
        },
        {
          id: "2",
          title: "Past Assignment",
          deadline: "2025-03-15T12:00:00Z",
          max_points: 80,
        },
      ],
    });
  });

  test("should display 'Upcoming Assignments' and 'Past Assignments' text", async () => {
    render(
      <MemoryRouter>
        <CourseAssignments />
      </MemoryRouter>
    );

    // Check if 'Upcoming Assignments' text is present
    await waitFor(() =>
      expect(screen.getByText("Upcoming Assignments")).toBeInTheDocument()
    );

    // Check if 'Past Assignments' text is present
    await waitFor(() =>
      expect(screen.getByText("Past Assignments")).toBeInTheDocument()
    );
  });

  test("should fetch assignments data", async () => {
    render(
      <MemoryRouter>
        <CourseAssignments />
      </MemoryRouter>
    );
    
    // Check if the API call fetches the correct data
    await waitFor(() =>
      expect(screen.getByText("Upcoming Assignment")).toBeInTheDocument()
    );

    await waitFor(() =>
      expect(screen.getByText("Past Assignment")).toBeInTheDocument()
    );
  });
});