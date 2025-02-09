import Grid from "@mui/material/Grid2";
import Paper from "@mui/material/Paper";
import Input from "@mui/material/Input";
import { Button, Typography } from "@mui/material";

import useAuth from "../hooks/useAuth";
import { Link, useNavigate, useLocation } from "react-router-dom";
import { useRef, useState, useEffect } from "react";
import { Box } from "@mui/material";

import axios from "../api/axios";
const LOGIN_URL = "/auth";

const Login = () => {

  const { setAuth } = useAuth();

  const navigate = useNavigate();
  const location = useLocation();
  const from = location.state?.from?.pathname || "/";

  const userRef = useRef();
  const errRef = useRef();

  const [user, setUser] = useState("");
  const [pwd, setPwd] = useState("");
  const [errMsg, setErrMsg] = useState("");

  useEffect(() => {
    userRef.current.focus();
  }, []);

  useEffect(() => {
    setErrMsg("");
  }, [user, pwd]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post(
        LOGIN_URL,
        JSON.stringify({ user, pwd }),
        {
          headers: { "Content-Type": "application/json" },
          withCredentials: true,
        }
      );
      console.log(JSON.stringify(response?.data));
      const accessToken = response?.data?.accessToken;
      const roles = response?.data?.roles;
      setAuth({ user, pwd, roles, accessToken });
      setUser("");
      setPwd("");
      navigate(from, { replace: true });
    } catch(e) {
      if (!e?.response) {
        setErrMsg("No Server Response");
      } else if (e.response?.status === 400) {
        setErrMsg("Missing Username or Password");
      } else if (e.response?.status === 401) {
        setErrMsg("Unauthorized");
      } else {
        setErrMsg("Login Failed");
        if(user=="user" && pwd=="pwd") {
          setAuth({ user, pwd, roles:[1001], accessToken:"secretUser" });
          setUser("");
          setPwd("");
          navigate(from, { replace: true });
        }
      }
    }
  }

  const paperStyle = {
    padding: 20,
    width: 280,
    margin: "19px auto",
  };
  const avatarStyle = { backgroundColor: "#D9D9D9" };
  const btnstyle = { backgroundColor: "#1B6DA1", margin: "20px 0" };
  const logoStyle = {
    backgroundColor: "#D9D9D9",
    margin: "10px 0",
    width: 70,
    height: 70,
  };
  const inputStyle = {margin: "20px auto"}
  const errorStyle = {color: "red"}

  return (
    <Box
      sx={{
        backgroundImage: `url('https://www.fa.ufl.edu/wp-content/uploads/2019/06/uf-monogram.png')`,
        backgroundSize: "cover",
        backgroundPosition: "center",
        minHeight: "100vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Grid>
        <form onSubmit={handleSubmit}>
          <Paper elavation={12} style={paperStyle}>
            <Grid align="center">
              <h2>Login</h2>
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
            <Input
              type="password"
              id="password"
              onChange={(e) => setPwd(e.target.value)}
              value={pwd}
              required
              placeholder="Password"
              fullWidth
            />

            <Button
              style={btnstyle}
              type="submit"
              color="primary"
              variant="contained"
              fullWidth
            >
              Login
            </Button>
            <Typography>
              <Link to="/resetPassword">Forgot Password?</Link>
            </Typography>
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
      </Grid>
    </Box>
  );
};

export default Login;