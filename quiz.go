package main

import ()

func quiz(chatroom chan Message) {
	<-chatroom
	qa := oneQA()
	for _, message := range qa.question {
		channel <- message
	}

	text := <-chatroom
	if isCorrectAnswer(text, qa) {
		chatroom <- "なんで知ってるの...?"
	} else {
		chatroom <- "やーいやーーいwwwwwwwwwwwwwwwwwww"
		chatroom <- Message("せぃかぃゎ" + qa.answer)
	}
}

func isCorrectAnswer(text Message, qa QA) bool {
	return text == Message(qa.answer)
}
