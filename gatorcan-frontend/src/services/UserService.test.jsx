import axios from "axios";
import { getUserDetails, resetPassword } from "./UserService";

jest.mock("axios");

describe("UserService", () => {
  test("successfully gets user details", async () => {
    axios.get.mockResolvedValue({
      data: { username: "testuser", email: "test@example.com" },
    });

    const result = await getUserDetails("testuser");

    expect(result).toEqual({
      success: true,
      data: { username: "testuser", email: "test@example.com" },
    });
    expect(axios.get).toHaveBeenCalledWith(
      "http://localhost:8080/user/testuser",
      expect.objectContaining({
        headers: expect.objectContaining({
          Authorization: expect.stringContaining("Bearer "),
        }),
      })
    );
  });

  test("successfully resets password", async () => {
    axios.put.mockResolvedValue({
      data: { message: "Password reset successfully" },
    });

    const result = await resetPassword("oldPass", "newPass");

    expect(result).toEqual({ success: true });
    expect(axios.put).toHaveBeenCalledWith(
      "http://localhost:8080/user/update/",
      {
        old_password: "oldPass",
        new_password: "newPass",
      },
      expect.objectContaining({
        headers: expect.objectContaining({
          Authorization: expect.stringContaining("Bearer "),
        }),
      })
    );
  });
});
