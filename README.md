# ZSOXIAL SOCIAL MEADIA BACKEND REST API
Zsoxial Social Media Backend Rest API is a sophisticated backend solution tailored for Social
Media applications. Built using Golang and the Gin web framework, it employs a Microservice
architecture with gRPC for seamless communication. This API excels in routing and processing
HTTP requests, meticulously adhering to industry-leading practices in code architecture and
dependency management. It offers a robust foundation for scalable and efficient Social Media
platforms.

## Key Features
- **Microservice Architecture**: Modular architecture promoting scalability and resilience
through loosely coupled services communicating via lightweight protocols.
- **Clean Code Architecture**: Structured, readable codebase following SOLID principles,
enhancing maintainability, and facilitating future development and updates.
Dependency Injection: Decoupling components by injecting dependencies, promoting
flexibility, testability, and easier code maintenance.
- **Database (Postgres and MongoDB)**: Utilizing Postgres for structured data and MongoDB for
flexibility in handling unstructured data, catering to diverse data requirements.
- **AWS Integration (S3)**: Seamlessly integrates with AWS S3 for scalable, secure, and reliable
cloud-based storage solutions, ensuring robust data management capabilities.
-**Social Media Features**: Incorporates essential functionalities like messaging, video calls,
notifications, posts, stories, facilitating engaging user interactions and content sharing.
-**Real-time Processing (Kafka, Websocket)**: Enables real-time data processing using Kafka for
distributed event streaming and Websocket for bidirectional communication, ensuring
timely updates and interactions.
-**Real-time Communication (WebRTC)**: Employs WebRTC for efficient peer-to-peer
communication, enabling high-quality, real-time audio/video calls and data sharing.
-**Caching (Redis)**: Utilizes Redis for fast, in-memory caching, enhancing performance by
storing frequently accessed data and reducing database load.
-**Containerization and Orchestration**: Implements containerization using Docker for
consistent deployment environments and orchestration with Kubernetes for efficient
management and scaling of containerized applications.


## API Documentation

For interactive API documentation, Swagger is implemented. You can explore and test the API endpoints in real-time.

## Security

Security is a top priority for the project:

- **OTP Verification**: Twilio API is integrated for OTP verification.
- **Refresh Tokens**: Enhances security and extends user sessions using refresh tokens.

## Getting Started

To run the project locally, you can follow these steps:

1. Clone the repository.
2. Set up your environment with the required dependencies, including Golang, PostgreSQL, Docker, and Wire.
3. Configure your environment variables (e.g., database credentials, AWS keys, Twilio credentials).
4. Build and run the project.


## Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Wire for Dependency Injection](https://github.com/google/wire)
- [PostgreSQL](https://www.postgresql.org/)
- [MongoDB](https://www.mongodb.com/)
- [AWS S3](https://aws.amazon.com/s3/)
- [Swagger API Documentation](https://swagger.io/)
- [Twilio](https://www.twilio.com/)
- [Kafka](https://kafka.apache.org/)
- [Redis](https://redis.io/)


## Using `go-gin-clean-arch` project

To use `go-gin-clean-arch` project, follow these steps:

```bash
# Navigate into the project
cd ./go-gin-clean-arch

# Generate wire_gen.go for dependency injection
make wire

# Run the project in Development Mode
make run
```

Additional commands:
```bash
➔ make help
build                          Compile the code, build Executable File
run                            Start application
wire                           Generate wire_gen.go
swag                           Generate swagger docs
```

# Environment Variables

Before running the project, you need to set the following environment variables with your corresponding values:

## PostgreSQL

- `BASE_URL`: Base url
- `DB_HOST`: Database host
- `DB_NAME`: Database name
- `DB_USER`: Database user
- `DB_PORT`: Database port
- `DB_PASSWORD`: Database password

## MongoDB

- `BASE_URL`=:Base url
- `DB_NAME`=Database name
- `DB_URL`="DatabaseURL url"

## Kafka

- `KAFKA_PORT`=Port
- `KAFKA_TOPIC`=Topic

## Redis
- `REDIS_PORT`=Port
- `REDIS_URL`=Url
- `REDIS_PASSWORD`=Password

## Twilio

- `DB_AUTHTOKEN`: Twilio authentication token
- `DB_ACCOUNTSID`: Twilio account SID
- `DB_SERVICESID`: Twilio services ID

## AWS

- `AWS_REGION`: AWS region
- `AWS_ACCESS_KEY_ID`: AWS access key ID
- `AWS_SECRET_ACCESS_KEY`: AWS secret access key


Make sure to provide the appropriate values for these environment variables to configure the project correctly.

# Live Demo

[SocialMediaBackedRestApi](https://zsoxial.zhooze.shop/swagger/index.html)