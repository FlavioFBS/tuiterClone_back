package bd

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN es objeto de conexion a BD
var MongoCN = ConectarBD()

// Options to mongoAtlas
var userMongo = os.Getenv("USER_MONGO")
var passwordMongo = os.Getenv("PASSWORD_MONGO")
var clusterMongo = os.Getenv("CLUSTER_MONGO")
var dbMongo = os.Getenv("DB_MONGO")
var goEnv = os.Getenv("GO_ENV")

// Options to localhost connect
var host string = "localhost"
var port = 27017

var uriConexion_mongoAtlas = "mongodb+srv://" + userMongo + ":" + passwordMongo + "@" + clusterMongo + ".bhind.mongodb.net/" + dbMongo + "?retryWrites=true&w=majority"
var clientOptions *options.ClientOptions

// ConectarBD ...
func ConectarBD() *mongo.Client {
	if goEnv == "dev" {
		// fmt.Println(goEnv)
		fmt.Println("development")
		clientOptions = options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port))
	} else {
		fmt.Println(goEnv)
		fmt.Println("production")
		clientOptions = options.Client().ApplyURI(uriConexion_mongoAtlas)
	}
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Error............")
		log.Fatal(err.Error())
		return client
	}
	// llamar a la bd para comprobar si está activa
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		// error en la BD
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion-BD establecida")
	return client
}

// ChequeoConnection es el ping a la BD
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
