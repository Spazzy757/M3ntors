import React from "react";
import { useAuth0 } from "@auth0/auth0-react";
import Avatar from '@mui/material/Avatar';

const Profile = () => {
  const { user, isAuthenticated } = useAuth0();

  return (
    isAuthenticated && (
        <Avatar src={user.picture} alt={user.name} />
    )
  );
};

export default Profile;
