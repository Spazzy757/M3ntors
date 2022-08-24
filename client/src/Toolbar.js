import React from "react";
import LoginButton from './auth0/LoginButton'
import LogoutButton from './auth0/LogoutButton'
import Profile from './auth0/Profile'
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';

const TBar = () => {
  return (
    <Toolbar> 
      <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
          M3NTORS
      </Typography>
      <Profile />
      <LoginButton />
      <LogoutButton />
    </ Toolbar>
  );
};

export default TBar;
