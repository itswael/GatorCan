import React, { useState, useEffect, useRef } from "react";
import {
  Container,
  Typography,
  Tabs,
  Tab,
  Box,
  Grid,
  List,
  ListItem,
  ListItemText,
  Paper,
  TextField,
  Button,
} from "@mui/material";
import InstructorNavbar from "./InstructorNavbar";
import dayjs from "dayjs";
import {
  fetchChatIds,
  getChatsByChatId,
  sendChat,
  subscribeToChats,
} from "../../graphql/queries";
import { generateClient } from "aws-amplify/api";

const InstructorChat = () => {
  const client = generateClient();
  const username = localStorage.getItem("username");
  const instructor_id = username;

  const courses = [1, 2];
  const [selectedCourse, setSelectedCourse] = useState(courses[0]);
  const [chats, setChats] = useState({});
  const [selectedChat, setSelectedChat] = useState(null);
  const [messagesMap, setMessagesMap] = useState({});
  const [newMessage, setNewMessage] = useState("");
  const bottomRef = useRef();

  // Subscription ref
  const chatSubscriptionRef = useRef();
  const courseSubscriptionRef = useRef();

  // Scroll to bottom of chat
  useEffect(() => {
    scrollToBottom();
  }, [selectedChat, messagesMap]);

  const scrollToBottom = () => {
    if (bottomRef.current) {
      bottomRef.current.scrollIntoView({ behavior: "smooth" });
    }
  };

  // Fetch chat list for selected course
  useEffect(() => {
    const loadChatList = async () => {
      try {
        const res = await client.graphql({
          query: fetchChatIds,
          variables: { input: { course_id: selectedCourse.toString() } },
        });

        const chatList = res.data.fetchChatIds;

        console.log("chatList:" + chatList);

        // Keep only the latest message per chat_id
        const latestPerChatId = {};
        chatList.forEach((msg) => {
          if (msg.sender_id == username) {
            return;
          }
          if (
            !latestPerChatId[msg.chat_id] ||
            new Date(msg.timestamp) >
              new Date(latestPerChatId[msg.chat_id].timestamp)
          ) {
            latestPerChatId[msg.chat_id] = msg;
          }
        });

        const uniqueChats = Object.values(latestPerChatId);

        const newChatMap = uniqueChats.map((chat) => ({
          chat_id: chat.chat_id,
          student_id: chat.sender_id,
          student_name: chat.sender_id,
        }));

        setChats((prev) => ({
          ...prev,
          [selectedCourse]: newChatMap,
        }));
      } catch (err) {
        console.error("Failed to fetch chats:", err);
      }
    };

    loadChatList();

    // Subscribe to new chats in this course
    if (courseSubscriptionRef.current) {
      courseSubscriptionRef.current.unsubscribe();
    }

    courseSubscriptionRef.current = client
      .graphql({
        query: subscribeToChats,
        variables: { chat_id: "*" },
      })
      .subscribe({
        next: ({ data }) => {
          const chat = data.subscribeToChats;
          if (chat.course_id === selectedCourse.toString()) {
            setChats((prev) => {
              const list = prev[selectedCourse] || [];
              const exists = list.some((c) => c.chat_id === chat.chat_id);
              if (!exists) {
                return {
                  ...prev,
                  [selectedCourse]: [
                    {
                      chat_id: chat.chat_id,
                      student_id: chat.sender_id,
                      student_name: chat.sender_id,
                    },
                    ...list,
                  ],
                };
              }
              return prev;
            });
          }
        },
        error: (err) => console.error("Course chat subscription error:", err),
      });

    return () => {
      courseSubscriptionRef.current?.unsubscribe();
    };
  }, [selectedCourse]);

  // Fetch messages when a chat is selected
  useEffect(() => {
    if (!selectedChat) return;

    const loadMessages = async () => {
      try {
        const res = await client.graphql({
          query: getChatsByChatId,
          variables: { chat_id: selectedChat },
        });

        const msgs = res.data.getChatsByChatId;
        const formatted = msgs.map((msg) => ({
          text: msg.message,
          datetime: msg.timestamp,
          sender: msg.sender_id === instructor_id ? "instructor" : "student",
        }));

        setMessagesMap((prev) => ({
          ...prev,
          [selectedChat]: formatted,
        }));
      } catch (err) {
        console.error("Failed to load chat messages:", err);
      }
    };

    loadMessages();

    if (chatSubscriptionRef.current) {
      chatSubscriptionRef.current.unsubscribe();
    }

    chatSubscriptionRef.current = client
      .graphql({
        query: subscribeToChats,
        variables: { chat_id: selectedChat },
      })
      .subscribe({
        next: ({ data }) => {
          const msg = data.subscribeToChats;
          setMessagesMap((prev) => ({
            ...prev,
            [msg.chat_id]: [
              ...(prev[msg.chat_id] || []),
              {
                text: msg.message,
                datetime: msg.timestamp,
                sender:
                  msg.sender_id === instructor_id ? "instructor" : "student",
              },
            ],
          }));
        },
        error: (err) => console.error("Chat subscription error:", err),
      });

    return () => {
      chatSubscriptionRef.current?.unsubscribe();
    };
  }, [selectedChat]);

  const handleSend = async () => {
    if (!newMessage.trim() || !selectedChat) return;

    const input = {
      chat_id: selectedChat,
      sender_id: instructor_id,
      receiver_id: null,
      course_id: selectedCourse.toString(),
      message: newMessage,
      timestamp: new Date().toISOString(),
    };

    try {
      await client.graphql({
        query: sendChat,
        variables: { input },
      });

      setNewMessage("");
    } catch (err) {
      console.error("Send error:", err);
    }
  };

  return (
    <>
      <InstructorNavbar />
      <Container
        sx={{ mt: 3, height: "80vh", display: "flex", flexDirection: "column" }}
      >
        <Typography variant="h5" gutterBottom>
          Chat
        </Typography>

        <Tabs
          value={selectedCourse}
          onChange={(e, val) => {
            setSelectedCourse(val);
            setSelectedChat(null);
          }}
        >
          {courses.map((courseId) => (
            <Tab
              key={courseId}
              label={`Course #${courseId}`}
              value={courseId}
            />
          ))}
        </Tabs>

        <Grid container spacing={2} sx={{ flexGrow: 1, minHeight: 0, mt: 2 }}>
          <Grid item xs={4}>
            <Paper sx={{ height: "100%", overflowY: "auto" }}>
              <List>
                {chats[selectedCourse]?.map((chat) => (
                  <ListItem
                    key={chat.chat_id}
                    button
                    selected={selectedChat === chat.chat_id}
                    onClick={() => setSelectedChat(chat.chat_id)}
                  >
                    <ListItemText
                      primary={chat.student_name}
                      secondary={chat.student_id}
                    />
                  </ListItem>
                ))}
              </List>
            </Paper>
          </Grid>

          <Grid item xs={8}>
            <Paper
              sx={{ height: "100%", display: "flex", flexDirection: "column" }}
            >
              {selectedChat ? (
                <>
                  <Box
                    sx={{
                      flexGrow: 1,
                      overflowY: "auto",
                      p: 2,
                      backgroundColor: "#f5f5f5",
                      minHeight: 0,
                      height: "0px",
                    }}
                  >
                    {messagesMap[selectedChat]?.map((msg, index) => (
                      <Box
                        key={index}
                        sx={{
                          display: "flex",
                          justifyContent:
                            msg.sender === "instructor"
                              ? "flex-end"
                              : "flex-start",
                          mb: 1,
                        }}
                      >
                        <Box
                          sx={{
                            p: 1.5,
                            borderRadius: 2,
                            maxWidth: "75%",
                            backgroundColor:
                              msg.sender === "instructor"
                                ? "#1976d2"
                                : "#e0e0e0",
                            color:
                              msg.sender === "instructor" ? "#fff" : "#000",
                          }}
                        >
                          <Typography variant="body2">{msg.text}</Typography>
                          <Typography
                            variant="caption"
                            sx={{
                              display: "block",
                              mt: 0.5,
                              textAlign: "right",
                            }}
                          >
                            {dayjs(msg.datetime).format("MMM D, h:mm A")}
                          </Typography>
                        </Box>
                      </Box>
                    ))}
                    <div ref={bottomRef}></div>
                  </Box>

                  <Box
                    sx={{ display: "flex", p: 2, borderTop: "1px solid #ccc" }}
                  >
                    <TextField
                      fullWidth
                      value={newMessage}
                      onChange={(e) => setNewMessage(e.target.value)}
                      placeholder="Type a message..."
                      size="small"
                      onKeyDown={(e) => {
                        if (e.key === "Enter" && !e.shiftKey) {
                          e.preventDefault();
                          handleSend();
                        }
                      }}
                    />
                    <Button
                      variant="contained"
                      onClick={handleSend}
                      sx={{ ml: 2 }}
                    >
                      Send
                    </Button>
                  </Box>
                </>
              ) : (
                <Box
                  sx={{
                    height: "100%",
                    display: "flex",
                    alignItems: "center",
                    justifyContent: "center",
                    color: "#999",
                  }}
                >
                  <Typography>Select a chat to view messages</Typography>
                </Box>
              )}
            </Paper>
          </Grid>
        </Grid>
      </Container>
    </>
  );
};

export default InstructorChat;
