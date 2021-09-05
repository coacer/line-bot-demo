const express = require('express');
const router = express.Router();
module.exports = router;

const { ProtoLoader } = require('../../plugins/grpc');

router.get('/', (req, res) => {
  return res.status(200).json({ status: 200 });
});

router.get('/test', async (req, res) => {
  const loader = new ProtoLoader('Trigger');
  const client = loader.getClient();

  const data = client.test({}, (err, response) => {
    console.log('err: ', err);
    console.log(response);
    return res.json(response);
  });
});
