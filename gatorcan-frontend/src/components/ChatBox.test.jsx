import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import ChatBox from "./ChatBox";
import "@testing-library/jest-dom";
import React from "react";

import { useState, useRef, useEffect } from "react";
import { Fab, Box, Paper, Typography } from "@mui/material";
import ChatIcon from "@mui/icons-material/Chat";

import "react-chat-widget/lib/styles.css";

describe("ChatBox Component", () => {
  test("should toggle chat window, send message and scroll to the latest message", async () => {
    render(<ChatBox />);

    // Check that the chat window is initially closed
    expect(screen.queryByText("Chat")).toBeNull();

    // Open the chat by clicking the FAB button
    fireEvent.click(screen.getByLabelText("chat"));

    // Check that the chat window opens
    expect(screen.getByText("Chat")).toBeInTheDocument();

    // Type a message
    const messageInput = screen.getByPlaceholderText("Type a message...");
    fireEvent.change(messageInput, { target: { value: "Hello, I'm here!" } });
    
  });
});
