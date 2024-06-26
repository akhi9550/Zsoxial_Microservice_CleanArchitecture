package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/IBM/sarama"
	"github.com/akhi9550/api-gateway/pkg/config"
	"github.com/akhi9550/api-gateway/pkg/utils/models"
	"github.com/gorilla/websocket"
)

type Helper struct {
	config *config.Config
}

func NewHelper(config *config.Config) *Helper {
	return &Helper{
		config: config,
	}
}

func (r *Helper) SendMessageToUser(User map[string]*websocket.Conn, msg []byte, userID string) {
	var message models.Message
	if err := json.Unmarshal([]byte(msg), &message); err != nil {
		fmt.Println("error while unmarshel ", err)
	}

	message.SenderID = userID
	recipientConn, ok := User[message.RecipientID]
	if ok {
		recipientConn.WriteMessage(websocket.TextMessage, msg)
	}
	err := KafkaProducer(message)
	fmt.Println("==send succesfully==", err)
}

func KafkaProducer(message models.Message) error {
	fmt.Println("from kafka ", message)
	cfg, _ := config.LoadConfig()
	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = true
	configs.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{cfg.KafkaPort}, configs)
	if err != nil {
		return err
	}

	result, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{Topic: cfg.KafkaTopic, Key: sarama.StringEncoder("Friend message"), Value: sarama.StringEncoder(result)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("err send message in kafka ", err)
	}
	log.Printf("[producer] partition id: %d; offset:%d, value: %v\n", partition, offset, msg)
	return nil
}

func (h *Helper) SendMessageToGroup(User map[string]*websocket.Conn, msg []byte, groupID, senderID string) {
	for key, conn := range User {
		if isUserInGroup(key, groupID) && key != senderID {
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				fmt.Printf("Error sending message to user %s: %s\n", key, err.Error())
				continue
			}
			fmt.Printf("Message sent to user %s\n", key)
		}
	}
}

func isUserInGroup(userID, groupID string) bool {
	return strings.HasPrefix(userID, groupID+"_")
}
