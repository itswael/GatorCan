import Grid from "@mui/material/Grid2";
import Paper from "@mui/material/Paper";
import Input from "@mui/material/Input";
import {Button} from "@mui/material";

import { useRef, useState, useEffect } from "react";
import { Box } from "@mui/material";
import {deleteUser} from "../../../services/AdminService";

const UserRegistration = ({ setCurrPage }) => {
  const userRef = useRef();
  const errRef = useRef();

  const [user, setUser] = useState("");
  const [errMsg, setErrMsg] = useState("");
  const [displaySuccess, setDisplaySuccess] = useState("");

  useEffect(() => {
    userRef.current.focus();
  }, []);

  useEffect(() => {
    setErrMsg("");
  }, [user]);

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await deleteUser(user);
      console.log("Delete user API Successful:", response);
      let success = response["success"];
      if (!success) {
        console.log(response["message"]);
        setErrMsg(response["message"]);
      } else {
        setDisplaySuccess(
          "Username: " + user + " deleted successfully!!"
        );
        setUser("");
      }
    } catch (error) {
        console.log(error);
      setErrMsg(error.response?.data?.message || "Unknown error");
    }
  };

  const paperStyle = {
    padding: 20,
    width: 400,
    margin: "19px auto",
  };
  const btnstyle = { backgroundColor: "#1B6DA1", margin: "20px 0" };
  const inputStyle = { margin: "12px auto" };
  const errorStyle = { color: "red" };

  return (
    <>
      <Grid container spacing={5} maxWidth={600}>
        <Box
          sx={{
            minHeight: "100vh",
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
          }}
        >
          <Button
            style={{ marginTop:"80px", alignSelf: "flex-start" }}
            onClick={() => setCurrPage(-1)}
          >
            Back
          </Button>
          <Grid>
            {displaySuccess !== "" ? (
              <Paper elevation={12} style={paperStyle}>
                <Grid align="center">
                  <p>{displaySuccess}</p>
                </Grid>
                <Button
                  style={btnstyle}
                  color="primary"
                  variant="contained"
                  fullWidth
                  onClick={() => {
                    setDisplaySuccess("");
                  }}
                >
                  DONE
                </Button>
              </Paper>
            ) : (
              <div>
                <form onSubmit={handleSubmit}>
                  <Paper elevation={12} style={paperStyle}>
                    <Grid align="center">
                      <h2>Delete user</h2>
                    </Grid>
                    <Input
                      type="text"
                      id="username"
                      ref={userRef}
                      autoComplete="off"
                      onChange={(e) => setUser(e.target.value)}
                      value={user}
                      required
                      style={inputStyle}
                      placeholder="Username"
                      fullWidth
                    />
                    <Button
                      style={btnstyle}
                      type="submit"
                      color="primary"
                      variant="contained"
                      fullWidth
                    >
                      Submit
                    </Button>
                  </Paper>
                </form>
                <p
                  ref={errRef}
                  className={errMsg ? "errmsg" : "offscreen"}
                  aria-live="assertive"
                  style={errorStyle}
                >
                  {errMsg}
                </p>
              </div>
            )}
          </Grid>
        </Box>
      </Grid>
    </>
  );
};

export default UserRegistration;
