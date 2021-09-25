const express = require('express');
const router = express.Router();
module.exports = router;

const LineWebhook = require('../../libs/line_webhook');

router.get('/health', async (req, res) => {
  const webhook = new LineWebhook();

  const result = await webhook.health();
  return res.status(200).json({ status: 200 });
});

router.use('/', async (req, res) => {
  const webhook = new LineWebhook();
  const { events } = req.body;

  for (const event of events) {
    switch (event.type) {
      case 'message': {
        await webhook.message(event);
        break;
      }
      default: {
        const err = new Error('Not found');
        err.status = 404;
        throw err;
      }
    }
  }

  return res.status(200).json({ status: 200 });
});
