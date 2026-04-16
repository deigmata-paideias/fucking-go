package main

import (
	"encoding/json"
	"fmt"
	"log"

	hellopb "github.com/deigmata-paideias/fucking-go/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

// gRPC Proto 的 oneof 语法在使用 Go SDK JSON 库序列化时会丢失，需要用 protojson 库来序列化和反序列化，保证值不丢失。
func main() {

	req := &hellopb.HelloRequest{
		Name: "Alice",
		Payload: &hellopb.HelloRequest_Text{
			Text: "hello from oneof",
		},
	}

	log.Printf("原始 oneof 值: %s", payloadValue(req))

	stdJSON, err := json.Marshal(req)
	must(err)
	log.Printf("encoding/json 序列化结果: %s", stdJSON)

	protoJSON, err := protojson.Marshal(req)
	must(err)
	log.Printf("protojson 序列化结果: %s", protoJSON)

	// encoding/json 只认识 Go 结构体字段，不认识 protobuf descriptor。
	// 读取合法的 Protobuf JSON 时，oneof 字段 text 会被忽略，反序列化后变成空值。
	var stdDecoded hellopb.HelloRequest
	must(json.Unmarshal(protoJSON, &stdDecoded))
	log.Printf("encoding/json 反序列化 读取 protojson 后的 oneof 值: %s", payloadValue(&stdDecoded))

	var protoDecoded hellopb.HelloRequest
	must(protojson.Unmarshal(protoJSON, &protoDecoded))
	log.Printf("protojson 反序列化 读取 protojson 后的 oneof 值: %s", payloadValue(&protoDecoded))
}

func payloadValue(req *hellopb.HelloRequest) string {

	switch payload := req.Payload.(type) {
	case *hellopb.HelloRequest_Text:
		return fmt.Sprintf("text=%q", payload.Text)
	case *hellopb.HelloRequest_Code:
		return fmt.Sprintf("code=%d", payload.Code)
	default:
		return "<nil>"
	}
}

func must(err error) {

	if err != nil {
		log.Fatal(err)
	}
}
