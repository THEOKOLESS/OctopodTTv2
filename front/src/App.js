import React, { useState, useEffect } from 'react';
import './App.css';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import TextField from '@mui/material/TextField';
import Grid from '@mui/material/Grid';
import Select from '@mui/material/Select';
import MenuItem from '@mui/material/MenuItem';
import InputLabel from '@mui/material/InputLabel';
import FormControl from '@mui/material/FormControl';

function BasicCard({country}) {
  return (
    <Grid item xs={12} sm={6} md={4} lg={3}>
      <Card sx={{ minWidth: 275 }}>
        <CardContent>
          <Typography variant="h5" component="div">
            {country.name} {country.flag}
          </Typography>
        </CardContent>
      </Card>
    </Grid>
  );
}

function App() {
  const [countries, setCountries] = useState([]);
  const [searchTerm, setSearchTerm] = useState("");
  const [sortOrder, setSortOrder] = useState("name");

  useEffect(() => {
    const fetchData = () => {
      fetch('http://127.0.0.1:8000/api/v1/countries')
        .then(response => response.json())
        .then(data => {
          setCountries(data);
        })
        .catch(error => {
          console.error('There was an error!', error);
        });
    }

    fetchData();

    const intervalId = setInterval(fetchData, 30000); // Refresh every 30 seconds

    return () => {
      clearInterval(intervalId); // Clear the interval when the component is unmounted
    }
  }, []);

  const sortedCountries = [...countries].sort((a, b) => {
    if (sortOrder === "name") {
      return a.name.localeCompare(b.name);
    } else {
      return a.population - b.population;
    }
  });

  return (
    <div>
      <div className="search-sort-container">
      <TextField
        label="Search"
        variant="outlined"
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
        className="search-input"
      />
      <FormControl>
        <InputLabel>Sort by</InputLabel>
        <Select
          value={sortOrder}
          onChange={(e) => setSortOrder(e.target.value)}
          className="sort-select"
        >
          <MenuItem value="name">Name</MenuItem>
          <MenuItem value="population">Population</MenuItem>
        </Select>
      </FormControl>
      </div>
      <Grid container spacing={3}>
        {sortedCountries.filter(country => country.name.toLowerCase().includes(searchTerm.toLowerCase())).map((country) => (
          <BasicCard key={country.name} country={country} />
        ))}
      </Grid>
    </div>
  );
}

export default App;
