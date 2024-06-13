package test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestJaon(t *testing.T) {
	// 测试 Add 函数的正确性
	jsonStr := `[  
        {  
          "id": "44444",  
          "name": "张三",  
          "age": 20  
        },  
        {  
          "id": 2022,  
          "name": "李四",  
          "age": 21  
        },  
        {  
          "id": "5555",  
          "name": "王五",  
          "age": 19  
        }  
    ]`

	var people []Person
	err := json.Unmarshal([]byte(jsonStr), &people)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(people[0].ID))

	// Post-process ID field to handle mixed types
	for _, person := range people {
		var idInt int
		var idStr string
		if err = json.Unmarshal(person.ID, &idInt); err == nil {
			fmt.Printf("ID (int): %d, Name: %s, Age: %d\n", idInt, person.Name, person.Age)
		} else if err := json.Unmarshal(person.ID, &idStr); err == nil {
			fmt.Printf("ID (string): %s, Name: %s, Age: %d\n", idStr, person.Name, person.Age)
		} else {
			fmt.Println("ID is neither int nor string")
		}
	}
}

type Person struct {
	ID   json.RawMessage `json:"id"`
	Name string          `json:"name"`
	Age  int             `json:"age"`
}
