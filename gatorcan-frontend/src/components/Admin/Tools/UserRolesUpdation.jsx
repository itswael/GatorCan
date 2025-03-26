import Grid from "@mui/material/Grid2";
import Paper from "@mui/material/Paper";
import Input from "@mui/material/Input";
import {
  Button,
  FormControl,
  FormGroup,
  FormControlLabel,
  Checkbox,
} from "@mui/material";

import { useRef, useState, useEffect } from "react";
import { Box } from "@mui/material";
import {updateUser} from "../../../services/AdminService";

const UserRolesUpdation = ({ setCurrPage }) => {
  const userRef = useRef();
  const errRef = useRef();

  const [user, setUser] = useState("");
  const [errMsg, setErrMsg] = useState("");
  const [displaySuccess, setDisplaySuccess] = useState("");

  // State for roles
  const [selectedRoles, setSelectedRoles] = useState({
    student: false,
    admin: false,
    instructor: false,
  });

  useEffect(() => {
    userRef.current.focus();
  }, []);

  useEffect(() => {
    setErrMsg("");
  }, [user, selectedRoles]);

  // Handle checkbox change for roles
  const handleRoleChange = (event) => {
    setSelectedRoles({
      ...selectedRoles,
      [event.target.name]: event.target.checked,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Get selected roles
    const roles = Object.keys(selectedRoles).filter(
      (role) => selectedRoles[role]
    );

    if (roles.length === 0) {
      setErrMsg("Please select at least one role");
      return;
    }

    try {
      // Simulate successful user creation
      const response = await updateUser(user, roles);
      console.log("Update user API Successful:", response);
      let success = response["success"];
      if (!success) {
        console.log(response["message"]);
        setErrMsg(response["message"]);
      } else {
        setDisplaySuccess(
          "User with name: " + user + " updated successfully!!"
        );
        setUser("");
        setSelectedRoles({
          student: false,
          admin: false,
          instructor: false,
        });
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
            style={{ marginTop: "80px", alignSelf: "flex-start" }}
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
                      <h2>Update user</h2>
                    </Grid>
                    <FormControl component="fieldset">
                      <FormGroup>
                        <FormControlLabel
                          control={
                            <Checkbox
                              checked={selectedRoles.student}
                              onChange={handleRoleChange}
                              name="student"
                            />
                          }
                          label="Student"
                        />
                        <FormControlLabel
                          control={
                            <Checkbox
                              checked={selectedRoles.admin}
                              onChange={handleRoleChange}
                              name="admin"
                            />
                          }
                          label="Admin"
                        />
                        <FormControlLabel
                          control={
                            <Checkbox
                              checked={selectedRoles.instructor}
                              onChange={handleRoleChange}
                              name="instructor"
                            />
                          }
                          label="Instructor"
                        />
                      </FormGroup>
                    </FormControl>
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

export default UserRolesUpdation;
