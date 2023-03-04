require("dotenv").config();

const express = require("express");
const app = express();
const port = 4000;

const { Configuration, OpenAIApi } = require("openai");
const configuration = new Configuration({
  apiKey: process.env.REACT_APP_OPENAI_API_KEY,
});
const openai = new OpenAIApi(configuration);

// adding body-parser and cors
const bodyParser = require("body-parser");
const cors = require("cors");

app.use(bodyParser.json());
app.use(cors());

app.post("/", async (req, res) => {
  const { message } = req.body;
  //console.log("message", message);
  const response = await openai.createChatCompletion({
    model: "gpt-3.5-turbo",
    messages: [{role: "user", content: message}],
    max_tokens: 3000,
    temperature: 0.3,
  });
  console.log("response", response.data);
  res.json({ botResponse: response.data.choices[0].message.content });
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
