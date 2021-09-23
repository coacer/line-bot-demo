const express = require('express');
const router = express.Router();
module.exports = router;

const errorHandler404 = require('./middleware/errorHandler404');
const errorHandler = require('./middleware/errorHandler');

router.use('/api', require('./router/api'));

router.use(errorHandler404);
router.use(errorHandler);
