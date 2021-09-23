const express = require('express');
const router = express.Router();
module.exports = router;

router.get('/', (req, res) => {
  return res.status(200).json({ status: 200 });
});

router.use('/channel', require('./channel'));
router.use('/line_webhook', require('./line_webhook'));
