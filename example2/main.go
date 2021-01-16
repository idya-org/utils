// package main

// import (
// 	"io/ioutil"
// 	"log"
// 	"stash.bms.bz/merchandise/utils"
// )

// // Aerospike ...
// type Aerospike struct {
// 	HostList []string
// 	Port     int
// 	UserName string
// 	Password string
// }

// var envConfig *Aerospike

// func main() {
// 	// Creates config path
// 	configFilePath := "conf.json"

// 	// Reads config json file
// 	bytes, err := ioutil.ReadFile(configFilePath)
// 	if err != nil {
// 		log.Println("ERRR - ", err)
// 	}

// 	getCcmsValue := func(key string) (string, error) {
// 		return key, nil
// 	}

// 	// Binds config based on ccms and present values
// 	err = utils.BindConfig(bytes, &envConfig, "ccms", getCcmsValue)
// 	if err != nil {
// 		log.Println("ERRRRR - ", err)
// 	}

// 	log.Println(envConfig)
// }
