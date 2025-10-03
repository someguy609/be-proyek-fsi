package config

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/someguy609/be-proyek-fsi/constants"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

// func RunExtension(db *gorm.DB) {
// 	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
// }

func SetUpDatabaseConnection() *mongo.Database {
	if os.Getenv("APP_ENV") != constants.ENUM_RUN_PRODUCTION {
		err := godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// todo: fix this later
	dbUri := fmt.Sprintf("mongodb+srv://%v:%v@%v/%v", dbUser, dbPass, dbHost, dbName)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dbUri).SetServerAPIOptions(serverAPI).SetRetryWrites(true).SetAppName(dbName)
	client, err := mongo.Connect(opts)

	if err != nil {
		panic(err)
	}

	db := client.Database(dbName)

	// dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=require", dbHost, dbUser, dbPass, dbName, dbPort)

	// db, err := gorm.Open(postgres.New(postgres.Config{
	// 	DSN:                  dsn,
	// 	PreferSimpleProtocol: true,
	// }), &gorm.Config{
	// 	Logger: SetupLogger(),
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// RunExtension(db)

	return db
}

func CloseDatabaseConnection(db *mongo.Database) {
	db.Client().Disconnect(context.TODO())
	// dbSQL, err := db.DB()
	// if err != nil {
	// 	panic(err)
	// }
	// dbSQL.Close()
}
