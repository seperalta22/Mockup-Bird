package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()

var clientOptions = options.Client().ApplyURI("mongodb+srv://SergioPeralta:ss3L3vhFAKdiEnEE@monguito.9xy6g.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

//*? ConectarBD es la funcion que me permite conectar la base de datos */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa con la BD")
	return client
}

//* CheckConnection sirve para chequear que se conecto con la BD
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
