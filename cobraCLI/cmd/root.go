package cmd

import (
	"database/sql"
	"github.com/LucasMacedoViana/posfullcycle/cobraCLI/internal/database"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"os"
)

func GetDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	return db
}

func GetCategoryDb(db *sql.DB) database.Category {
	return *database.NewCategory(db)
}

var rootCmd = &cobra.Command{
	Use:   "cobraCLI",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
