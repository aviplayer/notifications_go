import axios from 'axios';
import applyCaseMiddleware from 'axios-case-converter';

//export const httpClient = applyCaseMiddleware(axios.create());

export const httpClient = applyCaseMiddleware(
  axios.create({baseURL: 'http://localhost:8989'})
);
