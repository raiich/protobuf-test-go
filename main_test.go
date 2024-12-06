package main

import (
	"testing"

	"github.com/raiich/protobuf-test-go/generated/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

// repeatedでないのに複数回出現したとき

func TestName(t *testing.T) {
	// VARINTのフィールド (1バイトに収まる場合)
	{
		m := &pb.MyU32{
			U32: 1,
		}
		bs, err := proto.Marshal(m)
		assert.NoError(t, err)

		hex := []byte{
			0x08,
			0x01,
		}
		assert.Equal(t, hex, bs)

		readable := []byte{
			// フィールド番号=1, WireType=0 (←VARINT)
			1<<3 | 0,
			// 値=1
			1,
		}
		assert.Equal(t, hex, readable)

		m2 := &pb.MyU32{}
		assert.NoError(t, proto.Unmarshal(readable, m2))
		assert.EqualExportedValues(t, m, m2)
	}

	// VARINTのフィールド (2バイト以上使う場合)
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
			// フィールド番号=1, WireType=0 (←VARINT)
			1<<3 | 0,

			// 0b_0000001, 0b_0000000
			//   ↓ バイト列を反転
			// 0b_0000000, 0b_0000001
			//   ↓ 後続があれば最上位ビットを立てる
			// 0b10000000, 0b00000001
			0b10000000 | 0b_0000000,
			0b00000000 | 0b_0000001,
		}
		assert.Equal(t, hex, readable)

		m2 := &pb.MyU32{}
		assert.NoError(t, proto.Unmarshal(readable, m2))
		assert.EqualExportedValues(t, m, m2)
	}

	// 文字列のフィールド
	{
		m := pb.MyStr{
			Str: "aa",
		}
		bs, err := proto.Marshal(&m)
		assert.NoError(t, err)

		hex := []byte{
			0x0a,
			0x02,
			0x61, 0x61,
		}
		assert.Equal(t, hex, bs)

		readable := []byte{
			// フィールド番号=1, WireType=2 (←可変長バイト列)
			1<<3 | 2,
			2, // 長さ=2
			'a', 'a',
		}
		assert.Equal(t, hex, readable)

		m2 := &pb.MyStr{}
		assert.NoError(t, proto.Unmarshal(readable, m2))
		assert.EqualExportedValues(t, &m, m2)
	}

	// ネスト
	{
		m := &pb.MyNest{
			Str: &pb.MyStr{
				Str: "aa",
			},
		}
		bs, err := proto.Marshal(m)
		assert.NoError(t, err)

		hex := []byte{
			0x0a, 0x04,
			0x0a, 0x02, 0x61, 0x61,
		}
		assert.Equal(t, hex, bs)

		readable := []byte{
			1<<3 | 2, // フィールド番号=1, WireType=2 (←可変長バイト列)
			4,        // 長さ=2
			1<<3 | 2, // フィールド番号=1, WireType=2 (←可変長バイト列)
			2,        // 長さ=2
			'a', 'a', // 値="aa"
		}
		assert.Equal(t, hex, readable)

		m2 := &pb.MyNest{}
		assert.NoError(t, proto.Unmarshal(readable, m2))
		assert.EqualExportedValues(t, m, m2)
	}

	// repeated
	{
		m := &pb.MyRepeated{
			U32: []uint32{1, 2},
			Str: []string{"aa", "bb"},
		}
		bs, err := proto.Marshal(m)
		assert.NoError(t, err)
		assert.Equal(t, []byte{0xa, 0x2, 0x1, 0x2, 0x12, 0x2, 0x61, 0x61, 0x12, 0x2, 0x62, 0x62}, bs)

		hex := []byte{
			0x08, 0x01,
			0x08, 0x02,
			0x12, 0x2, 0x61, 0x61,
			0x12, 0x2, 0x62, 0x62,
		}
		assert.NotEqual(t, hex, bs)

		readable := []byte{
			1<<3 | 0, // フィールド番号=1, WireType=0 (←VARINT)
			1,        // 値=1
			1<<3 | 0, // フィールド番号=1, WireType=0 (←VARINT)
			2,        // 値=2
			2<<3 | 2, // フィールド番号=2, WireType=2 (←可変長バイト列)
			2,        // 長さ=2
			'a', 'a', // 値="aa"
			2<<3 | 2, // フィールド番号=2, WireType=2 (←可変長バイト列)
			2,        // 長さ=2
			'b', 'b', // 値="bb"
		}
		assert.Equal(t, hex, readable)

		m2 := &pb.MyRepeated{}
		assert.NoError(t, proto.Unmarshal(readable, m2))
		assert.EqualExportedValues(t, m, m2)
	}

	// 順不同なrepeated
	{
		m := &pb.MyRepeated{
			U32: []uint32{1, 2},
			Str: []string{"aa", "bb"},
		}
		bs, err := proto.Marshal(m)
		assert.NoError(t, err)
		assert.Equal(t, []byte{0xa, 0x2, 0x1, 0x2, 0x12, 0x2, 0x61, 0x61, 0x12, 0x2, 0x62, 0x62}, bs)

		hex := []byte{
			0x12, 0x2, 0x61, 0x61,
			0x08, 0x01,
			0x12, 0x2, 0x62, 0x62,
			0x08, 0x02,
		}
		assert.NotEqual(t, hex, bs)

		readable := []byte{
			2<<3 | 2, // フィールド番号=2, WireType=2 (←可変長バイト列)
			2,        // 長さ=2
			'a', 'a', // 値="aa"
			1<<3 | 0, // フィールド番号=1, WireType=0 (←VARINT)
			1,        // 値=1
			2<<3 | 2, // フィールド番号=2, WireType=2 (←可変長バイト列)
			2,        // 長さ=2
			'b', 'b', // 値="bb"
			1<<3 | 0, // フィールド番号=1, WireType=0 (←VARINT)
			2,        // 値=2
		}
		assert.Equal(t, hex, readable)

		m2 := &pb.MyRepeated{}
		assert.NoError(t, proto.Unmarshal(readable, m2))
		assert.EqualExportedValues(t, m, m2)
	}

	{
		m := &pb.MyU32{
			U32: 2,
		}
		hex := []byte{
			0x08,
			0x01,
			0x08,
			0x02,
		}
		readable := []byte{
			// フィールド番号=1, WireType=0 (←VARINT)
			1<<3 | 0,
			1, // 値=1
			// フィールド番号=1, WireType=0 (←VARINT)
			1<<3 | 0,
			2, // 値=2
		}
		assert.Equal(t, hex, readable)

		m2 := &pb.MyU32{}
		assert.NoError(t, proto.Unmarshal(readable, m2))
		assert.EqualExportedValues(t, m, m2)
	}
}
