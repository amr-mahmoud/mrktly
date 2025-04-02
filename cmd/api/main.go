package main

import (
	"fmt"
	"net/http"
	"github.com/amr-mahmoud/mrktly/internal/interface/controllers/user"
	"github.com/amr-mahmoud/mrktly/pkg/logger"
	"github.com/amr-mahmoud/mrktly/cmd/api/server"
)

//configuration
type Config struct {
	ServerAddress string
	Environment  string
}	
func main() {
	config := Config{
		ServerAddress: ":8080",
		Environment:  "development",
	}
	fmt.Println("Hello, This is my Technical Project Mrktly ! ")
	fmt.Println("This is a test for the new branch.")

		userCtrl := user.NewUserController()

		logger := logger.NewLogger("development")

		server := server.NewServer(userCtrl)
		logger.Info().Msgf("Starting server on %s", config.ServerAddress)
		logger.Info().Msgf("Server is running in %s mode", config.Environment)
		
		if err := http.ListenAndServe(config.ServerAddress, server); err != nil {
			logger.Fatal().Err(err).Msg("Failed to start server")
		}


}