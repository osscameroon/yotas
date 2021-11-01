import React from "react";
import ReactDOM from "react-dom";
import "./index.scss";
import "@fontsource/poppins";
import App from "./App";
import reportWebVitals from "./reportWebVitals";

import theme from "./themes";
import { ChakraProvider } from "@chakra-ui/react";
import LocaleProvider from "./locale/local-provider";

ReactDOM.render(
  <React.StrictMode>
    <ChakraProvider theme={theme}>
      <LocaleProvider>
        <App />
      </LocaleProvider>
    </ChakraProvider>
  </React.StrictMode>,
  document.getElementById("root")
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
