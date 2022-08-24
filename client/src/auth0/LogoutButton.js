import React from "react";
import { useAuth0 } from "@auth0/auth0-react";
import Button from '@mui/material/Button';

const LogoutButton = () => {
  const { isAuthenticated, logout } = useAuth0();

  return (
    isAuthenticated &&
    (
      <Button color="inherit" onClick={() => logout({ returnTo: window.location.origin })}>
        Log Out
      </Button>
    )
  );
};

export default LogoutButton;
