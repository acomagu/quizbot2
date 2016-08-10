package main

import ()

type Message string
type UserID string

var chatrooms map[UserID]chan Message = make(map[UserID]chan Message)

func reply(text Message, userID UserID) error {
	chatroom, ok := chatrooms[userID]
	if !ok {
		chatroom = make(chan Message)
		go sendMessageFromChatroom(chatroom, userID)
		go talk(chatroom)
	}
	chatroom <- text
	return nil
}

func sendMessageFromChatroom(chatroom chan Message, userID UserID) {
	for {
		text := <-chatroom
		bot.SendText([]string{string(userID)}, string(text))
	}
}
