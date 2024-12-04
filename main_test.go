package main

import (
	"fmt"
	"github.com/raiich/protobuf-test-go/generated/pb"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestName(t *testing.T) {
	m := pb.MyMessage{
		MDouble: 1,
		MUint32: 1,
		MString: "a",
	}
	bs, err := proto.Marshal(&m)
	assert.NoError(t, err)
	fmt.Printf("%#v\n", bs)
}
