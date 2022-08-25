import React from "react";
import Card from '@mui/material/Card';
import CardMedia from '@mui/material/CardMedia';
import test from '../test.png'
import FavoriteIcon from '@mui/icons-material/Favorite';
import ShareIcon from '@mui/icons-material/Share';
import IconButton from '@mui/material/IconButton';
import Grid2 from '@mui/material/Unstable_Grid2'; 
const courses = [
    {id: 1},
    {id: 2},
    {id: 3},
    {id: 4},
    {id: 5},
  ];
const Courses = () => {
  return (
    <Grid2 container spacing={2}>
      {courses.map(course => {
        return (
          <Grid2 xs={3} key={course.id}>
            <Card sx={{ maxWidth: 345 }}>
              <CardMedia
                  component="img"
                  image={test}
                  alt="Paella dish"
                >
              </CardMedia>
              <IconButton aria-label="add to favorites">
                    <FavoriteIcon />
                  </IconButton>
                  <IconButton aria-label="share">
                    <ShareIcon />
              </IconButton>
            </Card>
          </Grid2>
        )
      })}
    </Grid2>
  );
};

export default Courses;
