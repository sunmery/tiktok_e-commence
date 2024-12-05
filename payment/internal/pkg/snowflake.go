package pkg

import (
	"github.com/bwmarrin/snowflake"
	"log"
)

func SnowflakeID() snowflake.ID {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Panic(err)
	}

	// Generate a snowflake ID.
	id := node.Generate()

	return id
}
