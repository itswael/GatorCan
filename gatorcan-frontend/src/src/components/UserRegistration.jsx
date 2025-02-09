import Grid from "@mui/material/Grid2";
import Paper from "@mui/material/Paper";
import Input from "@mui/material/Input";
import { Button, MenuItem, Select, InputLabel } from "@mui/material";

import { useRef, useState, useEffect } from "react";
import { Box } from "@mui/material";

import axios from "../api/axios";
const REGISTER_URL = "/register";

const UserRegistration = () => {

  const userRef = useRef();
  const errRef = useRef();

  const [role, setRole] = useState(1001);
  const [user, setUser] = useState("");
  const [email, setEmail] = useState("");
  const [pwd, setPwd] = useState("");
  const [rePwd, setRePwd] = useState("");
  const [name, setName] = useState("");
  const [errMsg, setErrMsg] = useState("");
  const [displaySuccess, setDisplaySuccess] = useState("");

  useEffect(() => {
  if (!displaySuccess) userRef.current.focus();
  }, []);

  useEffect(() => {
    setErrMsg("");
  }, [user, pwd, rePwd, name, role]);

  const validateFormFields = () => {
    if(user == "" || pwd == "" || rePwd == "" || name == "") {
      setErrMsg("Please fille in the mandatory fields");
      return false;
    }
    if(pwd != rePwd) {
      setErrMsg("Passwords not matching")
      return false
    }
    return true;
  }

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      if (!validateFormFields()) {
        return;
      }
      // const response = await axios.post(
      //   REGISTER_URL,
      //   JSON.stringify({ role, user, pwd }),
      //   {
      //     headers: { "Content-Type": "application/json" },
      //     withCredentials: true,
      //   }
      // );
      // console.log(JSON.stringify(response?.data));
      // const success = response?.data?.success;
      const success = true;
      if (success) {
        setDisplaySuccess("User with name: " + name + " created successfully!!");
        setRole(1001)
        setUser("");
        setEmail("");
        setPwd("");
        setRePwd("");
        setName("");
      } else {
        setErrMsg(response?.data?.errMsg);
      }
    } catch (e) {
      if (!e?.response) {
        setErrMsg("No Server Response");
      } else if (e.response?.status === 400) {
        setErrMsg("Server validation failed");
      } else if (e.response?.status === 401) {
        setErrMsg("User unauthorized to register");
      } else {
        setErrMsg("User Registration Failed");
      }
    }
  };

  const paperStyle = {
    padding: 20,
    width: 280,
    margin: "19px auto",
  };
  const btnstyle = { backgroundColor: "#1B6DA1", margin: "20px 0" };
  const inputStyle = { margin: "10px auto" };
  const errorStyle = { color: "red" };
  const selectStyle = {}

  return (
    <Box
      sx={{
        backgroundImage: `url('')`,
        backgroundSize: "cover",
        backgroundPosition: "center",
        minHeight: "100vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Grid>
        {displaySuccess != "" ? (
          <Paper elavation={12} style={paperStyle}>
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
              Create new user
            </Button>
          </Paper>
        ) : (
          <div>
            <form onSubmit={handleSubmit}>
              <Paper elavation={12} style={paperStyle}>
                <Grid align="center">
                  <h2>User Registration</h2>
                </Grid>
                <InputLabel id="role-select-label">Choose Role</InputLabel>
                <Select
                  labelId="role-select-label"
                  id="role"
                  value={role}
                  label="Role"
                  onChange={(e) => {
                    setRole(e.target.value);
                  }}
                  style={selectStyle}
                  fullWidth
                >
                  <MenuItem value={1001}>Student</MenuItem>
                  <MenuItem value={5001}>Admin</MenuItem>
                  <MenuItem value={9001}>Instructor</MenuItem>
                </Select>
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
                <Input
                  type="email"
                  id="email"
                  autoComplete="off"
                  onChange={(e) => setEmail(e.target.value)}
                  value={email}
                  required
                  style={inputStyle}
                  placeholder="Email"
                  fullWidth
                />
                <Input
                  type="password"
                  id="password"
                  onChange={(e) => setPwd(e.target.value)}
                  value={pwd}
                  required
                  placeholder="Password"
                  fullWidth
                  style={inputStyle}
                />
                <Input
                  type="password"
                  id="repassword"
                  onChange={(e) => setRePwd(e.target.value)}
                  value={rePwd}
                  required
                  placeholder="Re-enter Password"
                  fullWidth
                  style={inputStyle}
                />
                <Input
                  type="text"
                  id="name"
                  autoComplete="off"
                  onChange={(e) => setName(e.target.value)}
                  value={name}
                  required
                  style={inputStyle}
                  placeholder="Name"
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
  );
};

export default UserRegistration;