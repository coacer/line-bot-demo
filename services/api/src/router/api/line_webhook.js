const express = require('express');
const router = express.Router();
module.exports = router;

const LineWebhook = require('../../libs/line_webhook');

router.get('/health', async (req, res) => {
  const webhook = new LineWebhook();

  const result = await webhook.health();
  return res.status(200).json({ status: 200 });
});
