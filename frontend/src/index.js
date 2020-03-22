import React from 'react';
import ReactDOM from 'react-dom';
import App from './router';

import './styles/index.css';

require('dotenv').config();

ReactDOM.render(<App />, document.getElementById('root'));
