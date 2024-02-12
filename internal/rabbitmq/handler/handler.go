package handler

import amqp "github.com/rabbitmq/amqp091-go"

// OnIncomingMessage type define callback function when receiving a message,
// all handlers MUST return this type
// message MUST be manually acked by the function either positively or negatively
// example: `msg.Ack(false)`.
type OnIncomingMessage func(msg amqp.Delivery)
