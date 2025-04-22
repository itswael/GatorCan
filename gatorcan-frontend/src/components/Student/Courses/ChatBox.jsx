import React, { useState, useRef, useEffect } from "react";
import { Fab, Box, Paper, Typography } from "@mui/material";
import ChatIcon from "@mui/icons-material/Chat";

import { generateClient } from "aws-amplify/api";
import {
  getChatsByChatId,
  fetchChatId,
  sendChat,
  subscribeToChats,
} from "../../../graphql/queries";

import "react-chat-widget/lib/styles.css";
import { generateUUID } from "../../../utils/Utils";
import dayjs from "dayjs";

function ChatBox({ course_id }) {
  const sender_id = localStorage.getItem("username");

  const client = generateClient();

  const [chatId, setChatId] = useState(null);

  const [openChat, setOpenChat] = useState(false);
  const [messages, setMessages] = useState([]);
  const [newMessage, setNewMessage] = useState("");
  const chatEndRef = useRef(null);
  const chatBoxRef = useRef(null);

  const toggleChat = () => {
    setOpenChat(!openChat);
  };

  const handleSendMessage = async () => {
    if (newMessage.trim() !== "") {
      const now = new Date();
      const isoTimestamp = now.toISOString();

      const input = {
        chat_id: chatId,
        sender_id: sender_id,
        receiver_id: null,
        course_id: course_id,
        message: newMessage,
        timestamp: isoTimestamp,
      };

      try {
        await client.graphql({
          query: sendChat,
          variables: { input },
        }).then(() => {
          console.log("hello");
          return;
        });

        // setMessages((prev) => [
        //   ...prev,
        //   {
        //     message: newMessage,
        //     dateTime: dayjs(isoTimestamp).format("MM-DD-YYYY HH:mm"),
        //     sender: true,
        //   },
        // ]);
        setNewMessage("");
      } catch (err) {
        console.error("Error sending chat:", err);
      }
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

    // input: course_id, user_id (jwtToken) -> fetchChatID(course_id, user_id)
    // if data then setChatID else new UUID

    const fetchChatMessages = async (chat_id) => {
      try {
        const res = await client.graphql({
          query: getChatsByChatId,
          variables: { chat_id },
        });

        const rawChats = res.data.getChatsByChatId;

        const formatted = rawChats.map((msg) => ({
          message: msg.message,
          dateTime: dayjs(msg.timestamp).format("MM-DD-YYYY HH:mm"),
          sender: msg.sender_id === sender_id,
        }));

        setMessages(formatted);
      } catch (err) {
        console.error("Error fetching chat messages:", err);
      }
    };

    const checkChatExists = async (studentId, courseId) => {
      try {
        const input = {
          sender_id: studentId,
          course_id: courseId,
        };

        const res = await client.graphql({
          query: fetchChatId,
          variables: { input },
        });

        const chatList = res.data.fetchChatId;

        if (chatList.length > 0) {
          const existingChatId = chatList[0].chat_id;
          setChatId(existingChatId);
          fetchChatMessages(existingChatId);
        } else {
          const uuid = generateUUID();
          setChatId(uuid);
          setMessages([]); // new chat, so no messages yet
        }
      } catch (err) {
        console.error("GraphQL error:", err);
      }
    };
    
    checkChatExists(sender_id, course_id);

    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  useEffect(() => {
    if (!chatId) return;

    const sub = client
      .graphql({
        query: subscribeToChats,
        variables: { chat_id: chatId },
      })
      .subscribe({
        next: ({ data }) => {
          const msg = data?.subscribeToChats;
          if (!msg) return;

          setMessages((prev) => [
            ...prev,
            {
              message: msg.message,
              dateTime: dayjs(msg.timestamp).format("MM-DD-YYYY HH:mm"),
              sender: msg.sender_id === sender_id,
            },
          ]);
        },
        error: (err) => console.error("Subscription error:", err),
      });

    return () => sub.unsubscribe();
  }, [chatId]);

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
