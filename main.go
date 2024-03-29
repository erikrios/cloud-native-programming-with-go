package main

import (
	"context"
	"fmt"
	"log"

	"github.com/erikrios/cloud-native-programming-with-go/config"
	_ "github.com/erikrios/cloud-native-programming-with-go/config"
	"github.com/erikrios/cloud-native-programming-with-go/controller"
	"github.com/erikrios/cloud-native-programming-with-go/lib/persistence/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	connection := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/%s/?authSource=admin",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connection))
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalln(err)
		}
	}()

	mongoDBHandler := mongodb.NewMongoDBLayer(client.Database(config.DBName))
	log.Printf("Successfully connected into database with address %p\n", client)

	log.Println(fmt.Sprintf("Server started on port %d...\n", config.Port))
	if err := controller.ServeAPI(mongoDBHandler); err != nil {
		log.Fatalln(err)
	}
}
