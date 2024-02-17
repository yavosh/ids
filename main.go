package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/snowflake"
	"github.com/gofrs/uuid/v5"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type generator func() (string, string, []byte)

func uuid4Generator() (string, string, []byte) {
	u, err := uuid.NewV4()
	if err != nil {
		return "uuid4", "error:" + err.Error(), nil
	}

	return "uuid4", u.String(), u.Bytes()
}

func uuid6Generator() (string, string, []byte) {
	u, err := uuid.NewV6()
	if err != nil {
		return "uuid6", "error:" + err.Error(), nil
	}

	return "uuid6", u.String(), u.Bytes()
}

func uuid7Generator() (string, string, []byte) {
	u, err := uuid.NewV7()
	if err != nil {
		return "uuid7", "error:" + err.Error(), nil
	}

	return "uuid7", u.String(), u.Bytes()
}

func xidGenerator() (string, string, []byte) {
	x := xid.New()
	return "xid", x.String(), x.Bytes()
}

func objectIdGenerator() (string, string, []byte) {
	id := primitive.NewObjectID()
	h, _ := hex.DecodeString(id.Hex())
	return "ObjectId", id.Hex(), h
}

func snowflakeGenerator() (string, string, []byte) {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(int64(rand.Int() % 10))
	if err != nil {
		return "snowflak", "error:" + err.Error(), nil
	}

	// Generate a snowflake ID.
	id := node.Generate()

	b := id.IntBytes()
	return "snowflak", id.String(), b[:]
}

func main() {
	fmt.Printf("%8s %-36s %-31s\n", "name", "string", "hex")
	fmt.Println(strings.Repeat("-", 80))

	for _, g := range []generator{
		xidGenerator,
		uuid4Generator,
		uuid6Generator,
		uuid7Generator,
		objectIdGenerator,
		snowflakeGenerator,
	} {
		name, id, bb := g()
		fmt.Printf("%8s %-36s 0x%-32x\n", name, id, bb)
	}
}
