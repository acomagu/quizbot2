package main

import ()

func talk(chatroom chan Message) {
	for {
		quiz(chatroom)
	}
}
