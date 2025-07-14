package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var rootCmd = &cobra.Command{
	Use:   "blog-tech",
	Short: "Blog Tech",
	Long:  "A tech blog application",
	Run: func(cmd *cobra.Command, args []string) {
		if err := godotenv.Load(); err != nil {
			log.Printf("Warning: No .env file found: %v", err)
		}

		dsn := os.Getenv("MYSQL_URI")
		if dsn == "" {
			log.Fatal("MYSQL_URI environment variable is not set")
		}

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

		db = db.Debug()

		log.Printf("Server starting on port :8080...")
		http.ListenAndServe(":8080", nil)
	},
}

func Execute() {
	rootCmd.AddCommand(outenvCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
