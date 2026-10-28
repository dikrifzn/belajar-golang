package main

import "fmt"

type Message interface{
	SendMessage()
}

type User struct {
	Name string
}

func (user User) SendMessage(){
	fmt.Println("Sending message from User", user.Name)
}

type Admin struct {
	Name string
}

func (admin Admin) SendMessage(){
	fmt.Println("Sending message from Admin", admin.Name)
}

type Bot struct {
	Name string
}

func (bot Bot) SendMessage(){
	fmt.Println("Sending message from Bot", bot.Name)
}

type Guest struct{
	Name string
}

func (guest Guest)	SendMessage(){
	fmt.Println("Sending message from Guest", guest.Name)
}

func SendNotification(message Message){
	message.SendMessage()
}

func main(){
	user := User{Name: "John"}
	admin := Admin{Name: "Alice"}
	bot := Bot{Name: "Bot123"}
	gues := Guest{Name: "GuestUser"}

	SendNotification(user)
	SendNotification(admin)
	SendNotification(bot)
	SendNotification(gues)
}