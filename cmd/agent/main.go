package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/bxcodec/faker/v3"
)

type Message struct {
	UUID      string   `faker:"uuid_digit"`
	FirstName string   `faker:"first_name"`
	LastName  string   `faker:"last_name"`
	Email     string   `faker:"email"`
	UnixTime  int64    `faker:"unix_time"`
	Latitude  float32  `faker:"lat"`
	Longitude float32  `faker:"long"`
	Steps     []string `faker:"-"`
}

const (
	MAX_BATCH_SIZE = 5
	TARGET_URL     = "http://localhost:8000"
)

func postMessages(messages []Message, timeout bool) error {
	fmt.Printf("Posting %d, Timeout: %v \n", len(messages), timeout)
	postData, _ := json.Marshal(messages)

	request, err := http.NewRequest("POST", TARGET_URL, bytes.NewBuffer(postData))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Running agent!")

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	transportChannel := make(chan Message, 10)
	doneChannel := make(chan bool, 1)

	//producer
	go func() {
		defer waitGroup.Done()
	terminate:
		for {
			select {
			case <-doneChannel:
				break terminate
			default:
				message := Message{}
				err := faker.FakeData(&message)
				if err != nil {
					fmt.Println(err)
				}
				transportChannel <- message
				time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			}
		}
	}()

	//batch sender (consumer)
	go func() {
		defer waitGroup.Done()
		messages := make([]Message, 0)
		timer := time.After(5 * time.Second)
	terminate:
		for {
			select {
			case message := <-transportChannel:
				messages = append(messages, message)
				if len(messages) == 5 {
					err := postMessages(messages, false)
					if err != nil {
						break terminate
					}
					messages = make([]Message, 0)
					timer = time.After(5 * time.Second)
				}
			case <-timer:
				err := postMessages(messages, true)
				if err != nil {
					break terminate
				}
				messages = make([]Message, 0)
				timer = time.After(5 * time.Second)
			}
		}
		doneChannel <- true
	}()

	waitGroup.Wait()
}
