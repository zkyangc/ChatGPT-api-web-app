import React, { useEffect } from "react";
import { useState } from "react";
import ReactMarkdown from "react-markdown";

const BotResponse = ({ response }) => {
  const [botResoponse, setBotResponse] = useState("");

  useEffect(() => {
    let index = 1;
    let msg = setInterval(() => {
      setBotResponse(response.slice(0, index));
      if (index >= response.length) {
        clearInterval(msg);
      }
      index++;
    }, 0.01);
  }, [response]);

  return (
    <pre>
      <ReactMarkdown>
        {botResoponse}
      </ReactMarkdown>
      {botResoponse === response ? "" : "|"}
    </pre>
  );
};

export default BotResponse;
