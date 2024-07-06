package main

import (
	"fmt"
	"math/rand"
)

func main() {
  massOfUser()
}

func massOfUser(){
  users := []string{"yanis", "jamari", "dan", "gov", "nick", "oliver", "zaahir"}
  randomUser := users[rand.Intn(len(users))]

  if randomUser == "gov" {
      mass := 18375663
      fmt.Println(randomUser, "has a mass of", mass, "kg")
    } else {
      mass := rand.Intn(100)
      fmt.Println(randomUser, "has a mass of", mass, "kg")}
}