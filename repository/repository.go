package repository

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/rs/xid"
)

var db *dynamo.DB

func init() {
	isLambda := strings.ToUpper(os.Getenv("LAMBDA"))

	if isLambda == "TRUE" {
		db = dynamo.New(session.New(), &aws.Config{Region: aws.String(os.Getenv("REGION"))})
	} else {
		s := session.Must(session.NewSessionWithOptions(session.Options{}))
		config := aws.Config{
			Endpoint: aws.String("http://localhost:8000"),
			Region:   aws.String("us-east-1"),
		}
		db = dynamo.New(s, &config)
	}

	log.Println(Agent{}.TableName())
}

// DB returns connection to db
func DB() *dynamo.DB {
	return db
}

// Table creates connection to table
func Table(name string) dynamo.Table {
	return db.Table(name)
}

// UUID generates unique identifier
func UUID() string {
	guid := xid.New()
	return guid.String()
}

// Token generates token
func Token() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	token := fmt.Sprintf("%x", b[0:])
	return token
}
