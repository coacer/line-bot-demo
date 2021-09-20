const express = require('express');
const router = express.Router();
module.exports = router;
const errorHandler404 = require('../../middleware/errorHandler404');
const errorHandler = require('../../middleware/errorHandler');

const LineWebhook = require('../../libs/line_webhook');

router.get('/', (req, res) => {
  return res.status(200).json({ status: 200 });
});

router.use('/channel', require('./channel'));
router.use('/line_webhook', require('./line_webhook'));

router.use(errorHandler404);
router.use(errorHandler);
