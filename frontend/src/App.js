import React from "react";
import { BrowserRouter as Router } from "react-router-dom";

import AppNav from "./components/navs";
import AppRouter from "./Router";

const App = () => (
  <Router>
    <div className="d-flex">
      <AppNav />
      <AppRouter />
    </div>
  </Router>
);

export default App;
