package cmd

import (
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
)

var client *kafka.Client

func kafkaCommand(cmd *cobra.Command, args []string) {
	brokers, _ := cmd.Flags().GetString("brokers")
	fmt.Printf("Kafka Cluster Broker Addr: %s\n", brokers)

	client = &kafka.Client{
		Addr:    kafka.TCP(brokers),
		Timeout: 10 * time.Second,
	}

	// for _, broker := range strings.Split(brokers, ",") {
	// 	conn, err := kafka.DialLeader(context.Background(), "tcp", broker, "", 0)
	// 	if err != nil {
	// 		fmt.Printf("Failed to dial leader for broker %s: %v", broker, err)
	// 		continue
	// 	}

	// 	break
	// }

	// if conn == nil {
	// 	fmt.Println("Failed to connect to any broker")
	// }

}

// Check Kafka cluster status
// func clusterStatus(brokers []string) {

// }

// Check Replicas on Cluster Mode
// func replicasReady(brokers []string) {

// }

// func producer(brokers []string) {

// }

// func consumer(brokers []string) {
// }

// func msgStressTest(brokers []string) {
// }
