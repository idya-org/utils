// package main

// import (
// 	"log"
// 	"time"

// 	"stash.bms.bz/merchandise/utils"
// )

// func main() {
// 	log.Print(utils.GetID())
// 	log.Print(utils.GetID("SUB"))

// 	// test collision

// 	m := make(map[string]interface{})

// 	start := time.Now()

// 	for i := 0; i < 10000000; i++ {
// 		id := utils.GetID()

// 		if m[id] != nil {
// 			log.Println("Collision detected ", i, id)
// 			break
// 		}
// 		m[id] = id
// 	}

// 	elapsed := time.Since(start)

// 	log.Println("No collisions, time taken: ", elapsed)
// }
