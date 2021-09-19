const express = require('express');
const router = express.Router();
module.exports = router;

const Channel = require('../../libs/channel');

/**
 * ヘルスチェック
 */
router.get('/health', async (req, res) => {
  const channel = new Channel();
  const result = await channel.test();
  return res.status(200).json(result);
});

/**
 * LineBotChannel作成
 */
router.post('/', async (req, res) => {
  const channel = new Channel();

  const result = await channel.invoke('create', req.body);
  return res.status(201).json(result);
});
