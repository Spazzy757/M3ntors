import React from "react";
import { styled } from '@mui/material/styles';
import Drawer from '@mui/material/Drawer';
import ChevronLeft from '@mui/icons-material/ChevronLeft';

const drawerWidth = 240;
const DrawerHeader = styled('div')(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  padding: theme.spacing(0, 1),
  // necessary for content to be below app bar
  ...theme.mixins.toolbar,
  justifyContent: 'flex-end',
}));

const SideBar = ({open, handleDrawer}) => {

  return (
    <Drawer
      sx={{
          width: drawerWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: drawerWidth,
            boxSizing: 'border-box',
          },
      }}
      open={open}
      anchor="left"
      variant="persistent"
    >
    <DrawerHeader>
      <ChevronLeft color="primary" onClick={() => { handleDrawer() }}/>
    </DrawerHeader>
    </Drawer>
  );
};

export default SideBar;
