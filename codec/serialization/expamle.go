package main

import (
	"encoding/json"
	"log"

	"github.com/golang/protobuf/proto"

	"github.com/Softwarekang/expamle/codec/serialization/pb"
)

func jsonCodec(student *pb.Student) {
	bytes, err := json.Marshal(student)
	if err != nil {
		log.Fatalf("marshal err:%v", err)
	}

	var target pb.Student
	if err := json.Unmarshal(bytes, &target); err != nil {
		log.Fatalf("unmarshal err:%v", err)
	}
}

func protobufCodec(student *pb.Student) {
	bytes, err := proto.Marshal(student)
	if err != nil {
		log.Fatalf("marshal err:%v", err)
	}

	var target pb.Student
	if err := proto.Unmarshal(bytes, &target); err != nil {
		log.Fatalf("unmarshal err:%v", err)
	}
}
