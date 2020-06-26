package main

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func main() {
	driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "123456", ""), func(c *neo4j.Config) {
		c.Encrypted = false
	})

	if err != nil {
		fmt.Println(err)
	}
	defer driver.Close()

	session, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	greeting, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		fmt.Println("0")
		result, err := transaction.Run(
			"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
			map[string]interface{}{"message": "Hello World!"})

		if err != nil {
			fmt.Println(err)
		}

		if result.Next() {
			return result.Record().GetByIndex(0), nil
		}

		return nil, result.Err()
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(greeting.(string))
}
