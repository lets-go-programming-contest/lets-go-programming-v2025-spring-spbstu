package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	handler "phonebook/internal/handler/http"
	"phonebook/internal/phonebook"
	"phonebook/internal/service"
)

func main() {
	var cfgFile string

	rootCmd := &cobra.Command{
		Use:   "phonebook-server",
		Short: "Phonebook REST API Server",
		Run: func(cmd *cobra.Command, args []string) {
			viper.SetConfigFile(cfgFile)
			if err := viper.ReadInConfig(); err != nil {
				log.Fatalf("Error reading config: %v", err)
			}

			dbPath := viper.GetString("database.path")
			port := viper.GetString("server.port")

			repo, err := phonebook.NewSQLiteRepo(dbPath)
			if err != nil {
				log.Fatalf("Error initializing DB: %v", err)
			}

			svc := service.New(repo)
			h := handler.NewHandler(svc)

			r := mux.NewRouter()
			h.RegisterRoutes(r)

			log.Printf("Listening on port %s\n", port)
			log.Fatal(http.ListenAndServe(":"+port, r))
		},
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "config file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
