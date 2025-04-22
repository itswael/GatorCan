export const fetchChatId = `
  query FetchChatId($input: FetchChatIdInput!) {
    fetchChatId(input: $input) {
      chat_id
      sender_id
      receiver_id
      course_id
      message
      timestamp
    }
  }
`;

export const getChatsByChatId = `
  query GetChatsByChatId($chat_id: ID!) {
    getChatsByChatId(chat_id: $chat_id) {
      chat_id
      message
      sender_id
      timestamp
    }
  }
`;

export const sendChat = `
  mutation SendChat($input: SendChatInput!) {
    sendChat(input: $input) {
      chat_id
      sender_id
      course_id
      message
      timestamp
    }
  }
`;

export const subscribeToChats = `
  subscription SubscribeToChats($chat_id: ID!) {
    subscribeToChats(chat_id: $chat_id) {
      chat_id
      sender_id
      message
      timestamp
    }
  }
`;

export const fetchChatIds = `
  query FetchChatIds($input: FetchChatIdsByCourseInput!) {
    fetchChatIds(input: $input) {
      chat_id
      sender_id
      timestamp
    }
  }
`;
