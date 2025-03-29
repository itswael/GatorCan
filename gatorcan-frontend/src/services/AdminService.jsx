import axios from "axios";

const base_url = "http://localhost:8080/";

export const addUser = async (username, password, email, roles) => {
  const add_user_url = base_url + "admin/add_user";
  console.log(roles);
  var roles_string = "";
  roles.forEach((role) => {
    roles_string = roles_string + '"' + role + '",';
  });
  roles_string = roles_string.slice(0, -1);
  console.log(roles_string);

  try {
    const refreshToken = localStorage.getItem("refreshToken");
    const response = await axios.post(
      add_user_url,
      { username, password, email, roles },
      {
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + refreshToken
        },
      }
    );
    console.log("RESPONSE RECEIVED: ", response.data);
    return { success: true };
  } catch (err) {
    if (err.response) {
      console.error("Add user failed:", err.response.data);
      return {
        success: false,
        message: "Add user failed: " + (err.response.data?.error || ""),
      };
    } else if (err.request) {
      console.error("No response received from server: ", err.request);
      return { success: false, message: "No response received from server" };
    } else {
      console.error("AXIOS ERROR:", err.message);
      return { success: false, message: "Unknown error" };
    }
  }
};

export const deleteUser = async (user) => {
  const delete_user_url = base_url + "admin/" + user;
  console.log(user);

  try {
    const refreshToken = localStorage.getItem("refreshToken");
    const response = await axios.delete(delete_user_url, {
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + refreshToken,
      },
      data: {},
    });
    console.log("RESPONSE RECEIVED: ", response);
    return { success: true };
  } catch (e) {
    if (e.response) {
      console.error("Delete user failed:", e.response.data);
      return {
        success: false,
        message: "Delete user failed: " + (e.response.data?.error || ""),
      };
    } else if (e.request) {
      console.error("No response received from server: ", e.request);
      return { success: false, message: "No response received from server" };
    } else {
      console.error("AXIOS ERROR:", e.message);
      return { success: false, message: "Unknown error" };
    }
  }
};

export const updateUser = async (user, roles) => {
  const update_user_url = base_url + "admin/update_role";
  console.log(user);

  var roles_string = "";
  roles.forEach((role) => {
    roles_string = roles_string + '"' + role + '",';
  });
  roles_string = roles_string.slice(0, -1);

  try {
    const refreshToken = localStorage.getItem("refreshToken");
    const response = await axios.put(
      update_user_url,
      {
        username: user, // Send the username key
        roles: roles, // Send roles as an array
      },
      {
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + refreshToken,
        },
      }
    );
    console.log("RESPONSE RECEIVED: ", response);
    return { success: true };
  } catch (e) {
    if (e.response) {
      console.error("Update user failed:", e.response.data);
      return {
        success: false,
        message: "Update user failed: " + (e.response.data?.error || ""),
      };
    } else if (e.request) {
      console.error("No response received from server: ", e.request);
      return { success: false, message: "No response received from server" };
    } else {
      console.error("AXIOS ERROR:", e.message);
      return { success: false, message: "Unknown error" };
    }
  }
};

export default addUser;
