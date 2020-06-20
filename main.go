package main

import (
	"os"
	"strings"

	"github.com/amagimedia/seshat/pkg"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	server := pkg.Server{}
	server.Mount(e)

	isLambda := strings.ToUpper(os.Getenv("LAMBDA"))

	if isLambda == "TRUE" {
		lambdaAdapter := &pkg.LambdaAdapter{Echo: e}
		lambda.Start(lambdaAdapter.Handler)
	} else {
		e.Logger.Fatal(e.Start(":1234"))
	}
}
