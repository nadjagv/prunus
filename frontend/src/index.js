import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import {BrowserRouter} from "react-router-dom"
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { green, purple } from '@mui/material/colors';
import AuthServis from './servisi/AuthServis';
import axios from "axios";

const theme = createTheme({
  palette: {
    primary: {
      light: "#e1a273",
      main: "#e88433",
      dark: "#8a3f05",
      contrastText: "#fff"
    },
    secondary: {
      main: "#f0f0f0",
    },
    typography: {
      fontFamily: [
        "Roboto Slab",
      ].join(",")
    },

    
  },
});

axios.interceptors.request.use((request) => {
  let korisnik = AuthServis.preuzmiKorisnika();
  if (korisnik) {
    request.headers.Authorization = "Bearer " + korisnik.Token;
  }
  return request;
});

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>

  <ThemeProvider theme={theme}>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </ThemeProvider>
  </React.StrictMode>

);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
