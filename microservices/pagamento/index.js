const express = require('express');
const app = express();
const cors = require('cors');

app.use(cors());

app.post('/pagar', async (request, response) => {
  return response.status(200).json({
    data: "OK",
  });
});

app.listen(3002, () => console.info('server listing por 3002'));

module.exports = app;
