package main

import (
	"testing"
	"fmt"
	"time"
	"sync"
)

func TestCreateChannel(t *testing.T) {
	// membuat channel
	channel := make(chan string)
	defer close(channel) // menutup channel

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Belajar Golang Channel" // mengirim data ke channel

		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel // menerima data dari channel
	fmt.Println(data)



	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) { //Channel sebagai parameter
	time.Sleep(2 * time.Second)
	channel <- "Belajar Golang Channel"
}


func TestChannelAsParameter(t *testing.T) { //Channel sebagai parameter
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) { //Channel IN
	time.Sleep(2 * time.Second)
	channel <- "Belajar Golang Channel"
}

func OnlyOut(channel <-chan string) { //Channel OUT
	data := <-channel
	fmt.Println(data)
	
}

func TestInOutChannel(t *testing.T) { //Channel IN dan OUT
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) { //Buffered Channel
	channel := make (chan string, 3)
	defer close(channel)
	go func(){
		channel <- "Belajar Golang Channel 1"
		channel <- "Belajar Golang Channel 2"
		channel <- "Belajar Golang Channel 3"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("selesai mengirim data ke channel")
}

func TestRangeChannel(t *testing.T) { //Range Channel
	channel := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- fmt.Sprintf("Data ke-%d", i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Selesai menerima data dari channel")
}

func TestSelectChannel(t *testing.T) { //Select Channel
	channel1 := make(chan string)
	channel2 := make(chan string)
	
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	defer close(channel1)
	defer close(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2:", data)
			counter++
		default:
			fmt.Println("Menunggu data...", counter)
		}

		if counter == 2 {
			break
		}
	}
}

func TestRaceCondition(t *testing.T) { //Race Condition
	x := 0

	var mutex sync.Mutex // membuat mutex untuk menghindari race condition
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Nilai ", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) { //RWMutex
	account := BankAccount{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Balance akhir:", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}