import { createGlobalStyle } from 'styled-components';
export const GlobalStyle = createGlobalStyle`
  html, body {
    height: 100%;
    margin: 0;
    padding: 0;
  }
  
  #root {
    height: 100%;
    overflow-y: auto;
  }
  
  * {
    box-sizing: border-box;
    font-family: Roboto, serif;
  }
`;
