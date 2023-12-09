package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) ([]byte, error) {
	marshaler := protojson.MarshalOptions{
		AllowPartial:   false,
		UseEnumNumbers: false,
		Indent:         "  ",
		UseProtoNames:  true,
	}

	return marshaler.Marshal(message)
}
