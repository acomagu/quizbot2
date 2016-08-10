package main

import ()

func quiz(chatroom chan Message) {
	<-chatroom
	qa := oneQA()
	sendTexts(chatroom, qa.question)

	text := <-chatroom
	if isCorrectAnswer(text, qa) {
		sendTexts(chatroom, []Message{"なんで知ってるの...?"})
	} else {
		sendTexts(chatroom, []Message{
			"やーいやーーいwwwwwwwwwwwwwwwwwww",
			Message("せぃかぃゎ" + qa.answer),
		})
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
