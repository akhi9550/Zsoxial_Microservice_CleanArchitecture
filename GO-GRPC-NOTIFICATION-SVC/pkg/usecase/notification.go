package usecase

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	authclientinterfaces "github.com/akhi9550/notification-svc/pkg/client/interface"
	"github.com/akhi9550/notification-svc/pkg/config"
	interfaces "github.com/akhi9550/notification-svc/pkg/repository/interface"
	services "github.com/akhi9550/notification-svc/pkg/usecase/interface"
	"github.com/akhi9550/notification-svc/pkg/utils/models"
)

type NotificationUseCase struct {
	notificationRepository interfaces.NotificationRepository
	authClient             authclientinterfaces.NewauthClient
}

func NewUserUserUseCase(repository interfaces.NotificationRepository, authClient authclientinterfaces.NewauthClient) services.NotificationUsecaseInterface {
	return &NotificationUseCase{
		notificationRepository: repository,
		authClient:             authClient,
	}
}

func (c *NotificationUseCase) ConsumeLikeMessage() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	configs := sarama.NewConfig()
	configs.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer([]string{cfg.KafkaPort}, configs)
	if err != nil {
		fmt.Println("Error creating Kafka consumer:", err)
		return
	}

	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition(cfg.KafkaLikeTopic, 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("Error creating partition consumer:", err)
		return
	}
	defer partitionConsumer.Close()
	fmt.Println("Kafka Likeconsumer started")
	for {
		select {
		case message := <-partitionConsumer.Messages():
			msg, err := c.UnmarshelNotification(message.Value)
			if err != nil {
				fmt.Println("Error unmarshalling message:", err)
				continue
			}
			fmt.Println("Received message:", msg)
			err = c.notificationRepository.StoreNotification(*msg)
			if err != nil {
				fmt.Println("Error storing message in repository:", err)
				continue
			}
		case err := <-partitionConsumer.Errors():
			fmt.Println("Kafka consumer error:", err)
		}
	}
}

func (c *NotificationUseCase) UnmarshelNotification(data []byte) (*models.NotificationReq, error) {
	var message models.NotificationReq
	err := json.Unmarshal(data, &message)
	if err != nil {
		return nil, err
	}
	message.CreatedAt = time.Now()

	return &message, nil
}

func (c *NotificationUseCase) ConsumeCommentMessage() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	configs := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{cfg.KafkaPort}, configs)
	if err != nil {
		fmt.Println("Error creating Kafka consumer:", err)
		return
	}
	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition(cfg.KafkaCommentTopic, 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("Error creating partition consumer for topic:-", cfg.KafkaCommentTopic, ":", err)
		return
	}
	defer partitionConsumer.Close()
	fmt.Println("Kafka Commentconsumer started")
	for {
		select {
		case message := <-partitionConsumer.Messages():
			msg, err := c.UnmarshelNotification(message.Value)
			if err != nil {
				fmt.Println("Error unmarshalling message:", err)
				continue
			}
			fmt.Println("Received message:", msg)
			err = c.notificationRepository.StoreNotification(*msg)
			if err != nil {
				fmt.Println("Error storing message in repository:", err)
				continue
			}
		case err := <-partitionConsumer.Errors():
			fmt.Println("Kafka consumer error:", err)
		}
	}
}

func (c *NotificationUseCase) ConsumeFollowReqMessage() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	configs := sarama.NewConfig()
	configs.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer([]string{cfg.KafkaPort}, configs)
	if err != nil {
		fmt.Println("Error creating Kafka consumer:", err)
		return
	}

	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition(cfg.KafkaFollowReqTopic, 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("Error creating partition consumer:", err)
		return
	}
	defer partitionConsumer.Close()
	fmt.Println("Kafka FollowReqconsumer started")
	for {
		select {
		case message := <-partitionConsumer.Messages():
			msg, err := c.UnmarshelNotification(message.Value)
			if err != nil {
				fmt.Println("Error unmarshalling message:", err)
				continue
			}
			fmt.Println("Received message:", msg)
			err = c.notificationRepository.StoreNotification(*msg)
			if err != nil {
				fmt.Println("Error storing message in repository:", err)
				continue
			}
		case err := <-partitionConsumer.Errors():
			fmt.Println("Kafka consumer error:", err)
		}
	}
}

func (c *NotificationUseCase) ConsumeAcceptFollowReqMessage() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	configs := sarama.NewConfig()
	configs.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer([]string{cfg.KafkaPort}, configs)
	if err != nil {
		fmt.Println("Error creating Kafka consumer:", err)
		return
	}

	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition(cfg.KafkaAcceptReqTopic, 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("Error creating partition consumer:", err)
		return
	}
	defer partitionConsumer.Close()
	fmt.Println("Kafka AcceptReqconsumer started")
	for {
		select {
		case message := <-partitionConsumer.Messages():
			msg, err := c.UnmarshelNotification(message.Value)
			if err != nil {
				fmt.Println("Error unmarshalling message:", err)
				continue
			}
			fmt.Println("Received message:", msg)
			err = c.notificationRepository.StoreNotification(*msg)
			if err != nil {
				fmt.Println("Error storing message in repository:", err)
				continue
			}
		case err := <-partitionConsumer.Errors():
			fmt.Println("Kafka consumer error:", err)
		}
	}
}

func (n *NotificationUseCase) GetNotification(userID int, pagination models.Pagination) ([]models.NotificationResponse, error) {
	var err error
	data, err := n.notificationRepository.GetNotification(userID, pagination)
	if err != nil {
		return nil, err
	}
	var response []models.NotificationResponse
	for _, res := range data {
		userData, err := n.authClient.UserData(res.UserID)
		if err != nil {
			return nil, err
		}
		response = append(response, models.NotificationResponse{
			UserID:    int(userData.UserId),
			Username:  userData.Username,
			Profile:   userData.Profile,
			PostID:    res.PostID,
			Message:   res.Message,
			CreatedAt: res.CreatedAt,
		})
	}
	return response, nil
}
