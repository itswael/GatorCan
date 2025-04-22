import axios from "axios";

const base_url = "http://gatorcan-backend.us-east-2.elasticbeanstalk.com/";

export const getUserDetails = async (username) => {
  const get_user_url = base_url + "user/" + username;

  try {
    const refreshToken = localStorage.getItem("refreshToken");
    const response = await axios.get(get_user_url, {
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + refreshToken,
      },
    });
    console.log("RESPONSE RECEIVED: ", response.data);

    return { success: true, data: response.data };
  } catch (e) {
    if (e.response) {
      console.error("Get user failed:", e.response.data);
      return {
        success: false,
        message: "Get user failed: " + (e.response.data?.error || ""),
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

export const resetPassword = async (oldPassword, password) => {
  const reset_pwd_url = base_url + "user/update/";

  try {
    const refreshToken = localStorage.getItem("refreshToken");
    const response = await axios.put(
      reset_pwd_url,
      {
        old_password: oldPassword,
        new_password: password,
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
      console.error("Reset password failed:", e.response.data);
      return {
        success: false,
        message: "Reset password failed: " + (e.response.data?.error || ""),
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

export default getUserDetails;