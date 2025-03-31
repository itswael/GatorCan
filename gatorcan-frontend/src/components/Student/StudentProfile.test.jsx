import React from "react";
import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import StudentProfile from "./StudentProfile";
import { getUserDetails, resetPassword } from "../../services/UserService";
import "@testing-library/jest-dom";

jest.mock("../../services/UserService");

describe("StudentProfile Component", () => {
  test("displays user details after fetching", async () => {
    getUserDetails.mockResolvedValue({
      success: true,
      data: { email: "test@example.com" },
    });

    render(
      <MemoryRouter>
        <StudentProfile />
      </MemoryRouter>
    );

    await waitFor(() =>
      expect(screen.getByText("Email: test@example.com")).toBeInTheDocument()
    );
  });

  test("shows loading indicator when fetching user details", () => {
    getUserDetails.mockResolvedValueOnce({
      success: true,
      data: { email: "test@example.com" },
    });

    render(
      <MemoryRouter>
        <StudentProfile />
      </MemoryRouter>
    );

    expect(screen.getByRole("progressbar")).toBeInTheDocument();
  });
});
