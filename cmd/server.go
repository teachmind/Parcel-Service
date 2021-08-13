package cmd

import (
	"os"
	"os/signal"
	"parcel-service/internal/app/server"
	"parcel-service/internal/pkg/postgres"
	"parcel-service/internal/app/parcel_carrier"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server",
	Long:  `Start server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := postgres.New(&postgres.Config{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		})
		println(db)

		if err != nil {
			panic(err)
		}
		s := server.NewServer(os.Getenv("APP_PORT"),
			parcel_carrier.NewService(parcel_carrier.NewRepository(db)),
		)

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			<-sig
			if err := s.Shutdown(); err != nil {
				log.Error().Err(err).Msg("error during server shutdown")
			}
		}()

		return s.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
