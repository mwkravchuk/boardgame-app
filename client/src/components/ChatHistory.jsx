import { Component } from "react";

import styles from "./ChatHistory.module.css";

const ChatHistory = ({ chatHistory }) => {
  return (
    <div>
      <h2>Chat History</h2>
      <ul className={styles.chatHistory}>
        {chatHistory.map((msg, index) => (
          <li key={index}>{msg}</li>
        ))}
      </ul>
    </div>
  );
};

export default ChatHistory;