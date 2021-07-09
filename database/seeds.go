package database

import (
	"fmt"
	"os"
	"strings"
	"log"

	"github.com/kylelemons/go-gypsy/yaml"
	"github.com/mhope-2/go_book_api/database/models"
	"gorm.io/gorm"
)

type SeedFn func(db *gorm.DB, path string)

func RunSeeds(db *gorm.DB, seeds []SeedFn) {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for _, seed := range seeds {
		seed(db, path)
	}
}




















