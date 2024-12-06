package main

import (
	"testing"

	"github.com/raiich/protobuf-test-go/generated/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestName(t *testing.T) {
	// VARINT field (using only 1 byte)
	{
		m := &pb.MyU32{U32: 3}
		bs, err := proto.Marshal(m)
		assert.NoError(t, err)

		hex := []byte{
			0x08,
			0x03,
		}
		assert.Equal(t, hex, bs)

		readable := []byte{
			// field=1, WireType=0b000 (←VARINT)
			1<<3 | 0b000,
			// value=3
			3,
		}
		assert.Equal(t, hex, readable)

		m2 := &pb.MyU32{}
		assert.NoError(t, proto.Unmarshal(readable, m2))
		assert.EqualExportedValues(t, m, m2)
	}

	// VARINT field (using more than 1 byte)
	{
		m := &pb.MyU32{U32: 1 << 7}
		bs, err := proto.Marshal(m)
		assert.NoError(t, err)

		hex := []byte{
			0x08,
			0x80,
			0x01,
		}
		assert.Equal(t, hex, bs)

		readable := []byte{
			// field=1, WireType=0b000 (←VARINT)
			1<<3 | 0b000,

			// 0b_0000001, 0b_0000000
			//   ↓ reverse byte array
			// 0b_0000000, 0b_0000001
			//   ↓ if there is trailing byte(s), set continuation bit to 1.
			// 0b10000000, 0b00000001
			0b10000000 | 0b_0000000, 0b00000000 | 0b_0000001,
		}
		assert.Equal(t, hex, readable)

		m2 := &pb.MyU32{}
		assert.NoError(t, proto.Unmarshal(readable, m2))
		assert.EqualExportedValues(t, m, m2)
	}

	// string field
	{
		m := &pb.MyStr{
			Str: "aaa",
		}
		bs, err := proto.Marshal(m)
		assert.NoError(t, err)

		hex := []byte{
			0x0a,
			0x03,
			0x61, 0x61, 0x61,
		}
		assert.Equal(t, hex, bs)

		readable := []byte{
			// field=1, WireType=0b010 (←variable length byte array)
			1<<3 | 0b010,
			3, // length=3
			'a', 'a', 'a',
		}
		assert.Equal(t, hex, readable)

		m2 := &pb.MyStr{}
		assert.NoError(t, proto.Unmarshal(readable, m2))
		assert.EqualExportedValues(t, m, m2)
	}

	// Sub messages
	{
		m := &pb.MySub{
			Sub: &pb.MyStr{
				Str: "aaa",
			},
		}
		bs, err := proto.Marshal(m)
		assert.NoError(t, err)

		hex := []byte{
			0x0a, 0x05,
			0x0a, 0x03, 0x61, 0x61, 0x61,
		}
		assert.Equal(t, hex, bs)

		readable := []byte{
			1<<3 | 0b010,  // field=1, WireType=0b010 (←variable length byte array)
			5,             // length=5
			1<<3 | 0b010,  // field=1, WireType=0b010 (←variable length byte array)
			3,             // length=3
			'a', 'a', 'a', // value="aaa"
		}
		assert.Equal(t, hex, readable)

		m2 := &pb.MySub{}
		assert.NoError(t, proto.Unmarshal(readable, m2))
		assert.EqualExportedValues(t, m, m2)
	}

	// repeated
	{
		m := &pb.MyRepeated{
			U32: []uint32{3, 4},
			Str: []string{"a", "b"},
		}
		bs, err := proto.Marshal(m)
		assert.NoError(t, err)
		assert.Equal(t, []byte{0xa, 0x02, 0x03, 0x04, 0x12, 0x01, 0x61, 0x12, 0x01, 0x62}, bs)

		hex := []byte{
			0x08, 0x03,
			0x08, 0x04,
			0x12, 0x1, 0x61,
			0x12, 0x1, 0x62,
		}
		assert.NotEqual(t, hex, bs)

		readable := []byte{
			1<<3 | 0b000, // field=1, WireType=0b000 (←VARINT)
			3,            // value=3
			1<<3 | 0b000, // field=1, WireType=0b000 (←VARINT)
			4,            // value=4
			2<<3 | 0b010, // field=2, WireType=0b010 (←variable length byte array)
			1,            // length=1
			'a',          // value="a"
			2<<3 | 0b010, // field=2, WireType=0b010 (←variable length byte array)
			1,            // length=1
			'b',          // value="b"
		}
		assert.Equal(t, hex, readable)

		m2 := &pb.MyRepeated{}
		assert.NoError(t, proto.Unmarshal(readable, m2))
		assert.EqualExportedValues(t, m, m2)
	}

	// repeated fields with random order
	{
		m := &pb.MyRepeated{
			U32: []uint32{3, 4},
			Str: []string{"aa", "bb"},
		}
		bs, err := proto.Marshal(m)
		assert.NoError(t, err)
		assert.Equal(t, []byte{0x0a, 0x02, 0x03, 0x04, 0x12, 0x02, 0x61, 0x61, 0x12, 0x02, 0x62, 0x62}, bs)

		hex := []byte{
			0x12, 0x02, 0x61, 0x61,
			0x08, 0x03,
			0x12, 0x02, 0x62, 0x62,
			0x08, 0x04,
		}
		assert.NotEqual(t, hex, bs)

		readable := []byte{
			2<<3 | 0b010, // field=2, WireType=0b010 (←variable length byte array)
			2,            // length=2
			'a', 'a',     // value="aa"
			1<<3 | 0b000, // field=1, WireType=0b000 (←VARINT)
			3,            // value=3
			2<<3 | 0b010, // field=2, WireType=0b010 (←variable length byte array)
			2,            // length=2
			'b', 'b',     // value="bb"
			1<<3 | 0b000, // field=1, WireType=0b000 (←VARINT)
			4,            // value=4
		}
		assert.Equal(t, hex, readable)

		m2 := &pb.MyRepeated{}
		assert.NoError(t, proto.Unmarshal(readable, m2))
		assert.EqualExportedValues(t, m, m2)
	}

	// non repeated but there are multiple fields
	{
		m := &pb.MyU32{
			U32: 4,
		}
		hex := []byte{
			0x08,
			0x03,
			0x08,
			0x04,
		}
		readable := []byte{
			// field=1, WireType=0b000 (←VARINT)
			1<<3 | 0b000,
			3, // value=3
			// field=1, WireType=0b000 (←VARINT)
			1<<3 | 0b000,
			4, // value=4
		}
		assert.Equal(t, hex, readable)

		m2 := &pb.MyU32{}
		assert.NoError(t, proto.Unmarshal(readable, m2))
		assert.EqualExportedValues(t, m, m2)
	}
}
