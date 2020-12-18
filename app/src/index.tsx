import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import reportWebVitals from './reportWebVitals';
import {GlobalStyle} from "./components";
import {createStore} from "./state/store";
import {Provider} from 'react-redux';
import {ChakraProvider} from "@chakra-ui/react";


const store = createStore();
ReactDOM.render(
  <Provider store={store}>
    <React.StrictMode>
      <GlobalStyle/>
      <ChakraProvider>
        <App/>
      </ChakraProvider>
    </React.StrictMode>
  </Provider>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
