const express = require('express');
const router = express.Router();
module.exports = router;

const Trigger = require('../../libs/trigger');

router.get('/', (req, res) => {
  return res.status(200).json({ status: 200 });
});

router.get('/test', async (req, res) => {
  const trigger = new Trigger();

  const result = await trigger.test();
  console.log('result: ', result);
  return res.json(result);
});
