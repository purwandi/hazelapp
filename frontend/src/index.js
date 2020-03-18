import React from 'react';
import ReactDOM from 'react-dom';
import App from './app';

import './styles/index.css';

require('dotenv').config();

ReactDOM.render(<App />, document.getElementById('root'));
