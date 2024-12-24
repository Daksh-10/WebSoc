"use client";
import { useState, useEffect } from "react";
import axios from "axios";

export default function HomePage() {
  const [counter, setCounter] = useState(0);
  const [randomString, setRandomString] = useState("");

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");

    ws.onmessage = async (event) => {
      const data = JSON.parse(event.data);

      // Update the counter
      setCounter(data.counter);

      // Fetch the random string using API
      const response = await axios.get("http://localhost:8080/data");
      setRandomString(response.data.randomString);
    };

    ws.onclose = () => console.log("WebSocket closed");

    return () => ws.close();
  }, []);

  return (
    <div className="flex flex-col items-center justify-top h-screen">
      <div className="text-5xl mt-20 mb-10">
        Real-time Counter and Random String
      </div>
      <div className="bg-blue-400 mb-5 text-black font-bold text-xl w-60 h-20 flex items-center justify-center rounded">
        Counter : {counter}
      </div>
      <div className="bg-blue-400 mt-2 text-black font-bold text-xl w-96 h-20 flex items-center justify-center rounded">
        Random String: {randomString}
      </div>
    </div>
  );
}
