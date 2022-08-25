import * as React from 'react';
import TBar from './components/Toolbar'
import SideBar from './components/SideBar'
import Courses from './components/Courses'
import './App.css';
import Box from '@mui/material/Box';

function App() {
  const [open, setOpen] = React.useState(false);
  const handleDrawer = () => {
    setOpen(!open)
  }
  return (
    <div className="App">
      <TBar handleDrawer={handleDrawer} />
      <Box sx={{ display: 'flex' }}>
        <SideBar
          open={open}
          handleDrawer={handleDrawer}
        />
        <Courses />
      </Box>
    </div>
  );
}

export default App;
