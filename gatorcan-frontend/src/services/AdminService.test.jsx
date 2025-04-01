import axios from "axios";
import { addUser, deleteUser, updateUser } from "./AdminService";

jest.mock("axios");

describe("AdminService", () => {
  test("successfully adds a user", async () => {
    axios.post.mockResolvedValue({
      data: { message: "User added successfully" },
    });

    const result = await addUser(
      "testuser",
      "password123",
      "test@example.com",
      ["admin"]
    );

    expect(result).toEqual({ success: true });
    expect(axios.post).toHaveBeenCalledWith(
      "http://localhost:8080/admin/add_user",
      {
        username: "testuser",
        password: "password123",
        email: "test@example.com",
        roles: ["admin"],
      },
      expect.objectContaining({
        headers: expect.objectContaining({
          Authorization: expect.stringContaining("Bearer "),
        }),
      })
    );
  });

  test("successfully deletes a user", async () => {
    axios.delete.mockResolvedValue({
      data: { message: "User deleted successfully" },
    });

    const result = await deleteUser("testuser");

    expect(result).toEqual({ success: true });
    expect(axios.delete).toHaveBeenCalledWith(
      "http://localhost:8080/admin/testuser",
      expect.objectContaining({
        headers: expect.objectContaining({
          Authorization: expect.stringContaining("Bearer "),
        }),
        data: {},
      })
    );
  });

  test("successfully updates a user", async () => {
    axios.put.mockResolvedValue({
      data: { message: "User updated successfully" },
    });

    const result = await updateUser("testuser", ["admin", "user"]);

    expect(result).toEqual({ success: true });
    expect(axios.put).toHaveBeenCalledWith(
      "http://localhost:8080/admin/update_role",
      {
        username: "testuser",
        roles: ["admin", "user"],
      },
      expect.objectContaining({
        headers: expect.objectContaining({
          Authorization: expect.stringContaining("Bearer "),
        }),
      })
    );
  });
});
