package cmd

import (
	"os"
	"os/signal"
	"parcel-service/internal/app/parcel"
	"parcel-service/internal/app/server"
	"parcel-service/internal/pkg/postgres"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server",
	Long:  `Start server`,
	RunE: func(cmd *cobra.Command, args []string) error {

		os.Setenv("APP_PORT", ":8080")
		os.Setenv("DB_HOST", "localhost")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("DB_USER", "bongo")
		os.Setenv("DB_PASSWORD", "")
		os.Setenv("DB_NAME", "parcelsrv")

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
			parcel.NewService(parcel.NewRepository(db)),
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
