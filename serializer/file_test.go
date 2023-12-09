package serializer_test

import (
	"fmt"
	"testing"

	"github.com/Aaketk17/go+gRPC/sample"
	"github.com/Aaketk17/go+gRPC/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../temp/laptop.bin"
	jsonFile := "../temp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	fmt.Println(err)
	require.NoError(t, err)

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)

	laptop2 := sample.NewLaptop()
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop1, laptop2))
}
