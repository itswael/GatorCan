import { render, screen, waitFor } from "@testing-library/react";
import InstructorDashboard from "./InstructorDashboard";
import InstructorService from "../../services/InstructorService";
import { MemoryRouter } from "react-router-dom";
import "@testing-library/jest-dom";
import React from "react";

jest.mock("../../services/InstructorService");

describe("InstructorDashboard Component", () => {
  test("renders message when no courses are allocated", async () => {
    InstructorService.fetchInstructorCourses.mockResolvedValue([]);
    render(
      <MemoryRouter>
        <InstructorDashboard />
      </MemoryRouter>
    );

    await waitFor(() =>
      expect(
        screen.getByText("No courses allocated to you.")
      ).toBeInTheDocument()
    );
  });

  test("renders course cards when courses are fetched", async () => {
    InstructorService.fetchInstructorCourses.mockResolvedValue([
      { id: 1, name: "Math", description: "Algebra and Geometry" },
      { id: 2, name: "Physics", description: "Mechanics and Thermo" },
    ]);

    render(
      <MemoryRouter>
        <InstructorDashboard />
      </MemoryRouter>
    );

    await waitFor(() =>
      expect(screen.getByText("#1 - Math")).toBeInTheDocument()
    );
    expect(screen.getByText("#2 - Physics")).toBeInTheDocument();
  });
});
