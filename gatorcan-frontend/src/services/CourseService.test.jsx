import axios from "axios";
import CourseService from "./CourseService";

jest.mock("axios");

describe("CourseService", () => {
  afterEach(() => {
    jest.clearAllMocks();
  });

  test("fetchAllCourses - successfully fetches all courses", async () => {
    const mockCourses = { courses: [{ id: 1, name: "Data Science" }] };
    axios.get.mockResolvedValue({ data: mockCourses });

    const result = await CourseService.fetchAllCourses();
    expect(result).toEqual(mockCourses);
  });

  test("fetchEnrolledCourses - successfully fetches enrolled courses", async () => {
    const mockCourses = [
      {
        ID: 1,
        Name: "Math",
        Description: "Math Course",
        StartDate: "2025-01-01",
        EndDate: "2025-05-01",
        InstructorName: "John Doe",
        InstructorEmail: "john@example.com",
      },
    ];
    axios.get.mockResolvedValue({ data: mockCourses });

    const result = await CourseService.fetchEnrolledCourses();
    expect(result).toEqual([
      {
        id: 1,
        name: "Math",
        description: "Math Course",
        created_at: "2025-01-01",
        updated_at: "2025-05-01",
        instructorName: "John Doe",
        instructorEmail: "john@example.com",
      },
    ]);
  });

  test("enrollInCourse - successfully enrolls in a course", async () => {
    axios.post.mockResolvedValue({});

    const result = await CourseService.enrollInCourse(1);
    expect(result).toEqual({ success: true });
  });

  test("fetchCourse - successfully fetches a course", async () => {
    const mockCourse = { id: 1, name: "Data Science" };
    axios.get.mockResolvedValue({ data: mockCourse });

    const result = await CourseService.fetchCourse({ id: 1 });
    expect(result).toEqual(mockCourse);
  });

  test("fetchAssignments - successfully fetches course assignments", async () => {
    const mockAssignments = [{ id: 1, name: "Assignment 1" }];
    axios.get.mockResolvedValue({ data: mockAssignments });

    const result = await CourseService.fetchAssignments({ id: 1 });
    expect(result).toEqual(mockAssignments);
  });

  test("fetchAssignmentDetails - successfully fetches assignment details", async () => {
    const mockAssignment = { id: 1, name: "Assignment 1" };
    axios.get.mockResolvedValue({ data: mockAssignment });

    const result = await CourseService.fetchAssignmentDetails({
      id: 1,
      assignment_id: 1,
    });
    expect(result).toEqual(mockAssignment);
  });
});
