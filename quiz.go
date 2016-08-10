package main

import ()

func quiz(chatroom chan Message) {
	<-chatroom
	qa := oneQA()
	sendTexts(chatroom, qa.question)

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

func sendTexts(chatroom chan Message, texts []Message) error {
	for _, text := range texts {
		chatroom <- text
	}
	return nil
}
