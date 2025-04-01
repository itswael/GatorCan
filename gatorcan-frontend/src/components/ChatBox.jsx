import React, { useState, useRef, useEffect } from "react";
import { Fab, Box, Paper, Typography } from "@mui/material";
import ChatIcon from "@mui/icons-material/Chat";

import "react-chat-widget/lib/styles.css";

function ChatBox() {
  const [openChat, setOpenChat] = useState(false);
  const [messages, setMessages] = useState([
    { message: "Hello!", dateTime: "2023-10-01 10:00 AM", sender: true },
    {
      message: "Hi there! How can I help you?",
      dateTime: "2023-10-01 10:01 AM",
      sender: false,
    },
    {
      message: "I need some assistance.",
      dateTime: "2023-10-01 10:02 AM",
      sender: true,
    },
    {
      message: "Sure! Let me know your query.",
      dateTime: "2023-10-01 10:03 AM",
      sender: false,
    },
  ]);
  const [newMessage, setNewMessage] = useState("");
  const chatEndRef = useRef(null);
  const chatBoxRef = useRef(null);

  const toggleChat = () => {
    setOpenChat(!openChat);
  };

  const handleSendMessage = () => {
    if (newMessage.trim() !== "") {
      const now = new Date();
      const formattedDateTime = now.toLocaleString("en-US", {
        dateStyle: "short",
        timeStyle: "short",
      });
      setMessages([
        ...messages,
        { message: newMessage, dateTime: formattedDateTime, sender: true },
      ]);
      setNewMessage("");
    }
  };

  useEffect(() => {
    if (chatEndRef.current) {
      chatEndRef.current.scrollIntoView({ behavior: "smooth" });
    }
  }, [messages]);

  useEffect(() => {
    const handleClickOutside = (event) => {
      if (
        chatBoxRef.current &&
        !chatBoxRef.current.contains(event.target) &&
        !event.target.closest('[aria-label="chat"]')
      ) {
        setOpenChat(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  return (
    <div className="App">
      <Fab
        color="secondary"
        aria-label="chat"
        style={{
          position: "fixed",
          bottom: 16,
          right: 16,
          backgroundColor: "rgb(29, 74, 124)",
        }}
        onClick={toggleChat}
      >
        <ChatIcon style={{ color: "white" }} />
      </Fab>

      {openChat && (
        <Box
          ref={chatBoxRef}
          sx={{
            position: "fixed",
            bottom: 80,
            right: 16,
            width: 300,
            height: 400,
            backgroundColor: "rgb(29, 74, 124)",
            boxShadow: 3,
            borderRadius: 2,
            overflow: "hidden",
            display: "flex",
            flexDirection: "column",
          }}
        >
          <Paper
            elevation={3}
            sx={{
              height: "100%",
              display: "flex",
              flexDirection: "column",
            }}
          >
            <Box
              sx={{
                backgroundColor: "rgb(29, 74, 124)",
                color: "white",
                padding: 1,
                textAlign: "center",
              }}
            >
              <Typography variant="h6">Chat</Typography>
            </Box>
            <Box
              sx={{
                flex: 1,
                padding: 2,
                overflowY: "auto",
                display: "flex",
                flexDirection: "column",
                gap: 1,
              }}
            >
              {messages.map((msg, index) => (
                <Box
                  key={index}
                  sx={{
                    alignSelf: msg.sender ? "flex-end" : "flex-start",
                    backgroundColor: msg.sender
                      ? "rgb(29, 74, 124)"
                      : "#e0e0e0",
                    color: msg.sender ? "white" : "black",
                    padding: 1,
                    borderRadius: 2,
                    maxWidth: "70%",
                  }}
                >
                  <Typography variant="body1">{msg.message}</Typography>
                  <Typography
                    variant="caption"
                    sx={{
                      display: "block",
                      textAlign: "right",
                      marginTop: 0.5,
                    }}
                  >
                    {msg.dateTime}
                  </Typography>
                </Box>
              ))}
              <div ref={chatEndRef} />
            </Box>
            <Box
              sx={{
                display: "flex",
                padding: 1,
                borderTop: "1px solid #e0e0e0",
              }}
            >
              <input
                type="text"
                value={newMessage}
                onChange={(e) => setNewMessage(e.target.value)}
                onKeyDown={(e) => {
                  if (e.key === "Enter") {
                    handleSendMessage();
                    e.preventDefault();
                  }
                }}
                placeholder="Type a message..."
                style={{
                  flex: 1,
                  border: "none",
                  outline: "none",
                  padding: "8px",
                  fontSize: "14px",
                }}
              />
              <button
                onClick={handleSendMessage}
                style={{
                  backgroundColor: "rgb(29, 74, 124)",
                  color: "white",
                  border: "none",
                  padding: "8px 16px",
                  cursor: "pointer",
                }}
              >
                Send
              </button>
            </Box>
          </Paper>
        </Box>
      )}
    </div>
  );
}

export default ChatBox;
