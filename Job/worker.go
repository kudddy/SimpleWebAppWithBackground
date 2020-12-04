package Job

import (
	"awesomeProject/MessageTypes"
	"fmt"
	"time"
)

func Workers(in chan MessageTypes.Profile, out chan MessageTypes.Profile) {
	for item := range in {
		fmt.Println("1. Item", item)
		time.Sleep(10 * time.Second)
		item.Name = "Kirill"
		out <- item
	}
}
