package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/rulanugrh/lysithea/internal/config"
	"github.com/rulanugrh/lysithea/internal/util"
	"gorm.io/gorm"
)

func help() {
	helpContent := [][]string{
		{"help", "show help message"},
		{"migrate", "command for running migration model"},
		{"seeder", "command for seeder to db"},
		{"serve", "command for serve http API"},
	}

	maxLen := len(helpContent[0][0])
	for _, part := range helpContent {
		length := len(part[0])
		if length > maxLen {
			maxLen = length
		}
	}

	var builder strings.Builder
	const space = 4
	for _, part := range helpContent {
		builder.WriteString(part[0])
		spacer := (maxLen - len(part[0])) + space
		for spacer > 0 {
			builder.WriteByte(' ')
			spacer--
		}
		builder.WriteString(part[1])
		builder.WriteByte('\n')
	}

	println(builder.String()[:builder.Len()-1])
}

func migrate(db *gorm.DB) {
	err := util.Migrate(db)
	if err != nil {
		log.Printf("[%s %d:%d:%d] Database Failed Migration", time.DateOnly, time.Hour, time.Minute, time.Second)
	}

	log.Printf("[%s %d:%d:%d] Database Success Migration", time.DateOnly, time.Hour, time.Minute, time.Second)

}

func seeder(db *gorm.DB) {
	err := util.Seeder(db)
	if err != nil {
		log.Printf("[%s %d:%d:%d] Failed Seeder", time.DateOnly, time.Hour, time.Minute, time.Second)
	}

	log.Printf("[%s %d:%d:%d] Seeder Finished", time.DateOnly, time.Hour, time.Minute, time.Second)
}

func main() {
	db, err := config.NewConnection()
	if err != nil {
		log.Printf("error to connect database %v", err)
	}

	args := os.Args[1]

	switch args {
	case "migrate":
		migrate(db)
	case "seeder":
		seeder(db)
	case "help":
		help()
	default:
		println("Use args help to show message")
	}
}
