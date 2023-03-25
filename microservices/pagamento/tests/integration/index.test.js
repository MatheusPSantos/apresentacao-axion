const { createClient } = require("../setup");

describe('Testar endpoint de pagamento', () => {
  let client;
  beforeAll(async () => {
    client = await createClient();
  });

  it('Deve retornar o status no body', async () => {
    const res = await client.post('/pagar');
    expect(res.status).toBe(200);
    expect(res.body).toHaveProperty("status", "OK");
  })
})