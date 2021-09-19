'use strict';

const express = require('express');
const bodyParser = require('body-parser');
require('express-async-errors');

// Constants
const PORT = process.env.API_PORT;
const HOST = '0.0.0.0';

// App
const app = express();

app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());
app.use(require('./src/app'));
app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);
