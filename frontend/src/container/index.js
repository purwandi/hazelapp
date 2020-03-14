import React from "react";

import Navigation from "./nav";
import AppContainerRouter from "./router";

const AppContainer = props => {
  return (
    <div className="h-screen flex">
      <div className="bg-gray-100 w-56 p-4 border-r border-gray-400 overflow-auto">
        <Navigation {...props} />
      </div>
      <div className="flex-1 min-w-0 bg-white">
        <AppContainerRouter />
      </div>
    </div>
  );
};

export default AppContainer;
