### A Simple Currency Converter

(Will tidy this up soon!)

This is a simple currency converter written in Go.

### Justifying Dependencies Used

- `github.com/gin-gonic/gin v1.7.2` => this was used to allow for the use of Go Gin web framework  
- `github.com/joho/godotenv v1.3.0` => to load environment variables from a .env file  
- `github.com/kylelemons/go-gypsy v1.0.0` => was used to parse yml file to seed currencies model into the database   
- `gorm.io/driver/postgres v1.1.0` => postgres db driver to allow for connection to postgres database   
- `gorm.io/gorm v1.21.11` => allow for usage of a relational database  


### Build/Compile the application
- change .env.sample to .env
- fill .env file with right parameters
- run `go build -o=./out/main`


### Run the application
- run `./out/main`


### Endpoint
- http://localhost:{PORT}/v1/convert/currency


### Request Type
- POST

### Sample Request
![Sample Request](/imgs/sample_request.png?raw=true) 


### Sample Response
![Sample Response](/imgs/sample_response.png?raw=true) 


### Currencies Table Structure
![Currencies Table Structure](/imgs/currencies_table_structure.png?raw=true) 

### Exchange Rates Table Structure
![Exchange Rates Structure](/imgs/exchange_rates_table_structure.png?raw=true) 

