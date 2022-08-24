import React from "react";
import { useAuth0 } from "@auth0/auth0-react";
import Button from '@mui/material/Button';

const LoginButton = () => {
  const { isAuthenticated, loginWithRedirect } = useAuth0();

  return (
    !isAuthenticated &&
    (
      <Button color="inherit" onClick={() => loginWithRedirect()}>
        Log In
      </Button>
    )
  );
};

export default LoginButton;

