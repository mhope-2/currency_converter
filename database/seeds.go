package database

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kylelemons/go-gypsy/yaml"
	"github.com/mhope-2/currency_converter/database/models"
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


func SeedCurrencies (DB *gorm.DB, path string){
	config, err := yaml.ReadFile(path + "/database/currencies.yml")

	if err != nil {
		panic(err)
	}

	currencyList, ok := config.Root.(yaml.List)
	if !ok {
		panic("failed to parse currencies.yml")
	}

	for i := 0; i < currencyList.Len(); i++ {	
		currency_str := strings.ToLower(fmt.Sprintf("%s", currencyList.Item(i)))
		name := strings.Split(currency_str, "|")[0]
		code := strings.Split(currency_str, "|")[1]
		symbol := strings.Split(currency_str, "|")[2]

		var currency models.Currency

		if err := DB.Where("name = ? AND code = ? AND symbol = ?", name, code, symbol).First(&currency).Error; err != nil {
			
			if err == gorm.ErrRecordNotFound{
				DB.Create(&models.Currency{Name: name, Code: code, Symbol: symbol})
			}else{
				log.Printf("Currency [ %s ] could not be found", name)
				log.Println(err)
			}
		}
	}

}


















