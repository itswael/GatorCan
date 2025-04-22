import React from "react";
import axios from "axios";
import InstructorService, {
  fetchInstructorCourses,
  upsertAssignment,
  gradeAssignment,
} from "./InstructorService";

jest.mock("axios");

beforeEach(() => {
  localStorage.setItem("refreshToken", "mocked-token");
});

describe("InstructorService", () => {
  test("fetchInstructorCourses returns course list", async () => {
    const mockCourses = [
      { id: 1, name: "Math", description: "Algebra" },
      { id: 2, name: "Physics", description: "Mechanics" },
    ];

    axios.get.mockResolvedValue({ data: mockCourses });

    const result = await fetchInstructorCourses();

    expect(result).toEqual(mockCourses);
    expect(axios.get).toHaveBeenCalledWith(
      "http://gatorcan-backend.us-east-2.elasticbeanstalk.com/courses/?page=1&pageSize=10",
      expect.objectContaining({
        headers: expect.objectContaining({
          Authorization: "Bearer mocked-token",
        }),
      })
    );
  });

  test("upsertAssignment returns success true on success", async () => {
    axios.post.mockResolvedValue({});

    const result = await upsertAssignment({
      courseID: 1,
      assignment: { id: 1, title: "Assignment" },
    });

    expect(result).toEqual({ success: true });
  });

  test("gradeAssignment returns success false on error", async () => {
    axios.post.mockRejectedValue(new Error("Fail"));

    const result = await gradeAssignment({
      cid: 1,
      aid: 2,
      grades: { grade: 100 },
    });

    expect(result).toEqual({ success: false });
  });
});
