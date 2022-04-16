package main

import (
	"testing"

	"github.com/Softwarekang/expamle/codec/serialization/pb"
)

// BenchmarkJsonCodec benchmark test for jsonCodec
func BenchmarkJsonCodec(t *testing.B) {
	student := &pb.Student{
		Name:   "codec",
		Male:   true,
		Scores: []int32{99, 100, 99},
	}
	for i := 0; i < t.N; i++ {
		jsonCodec(student)
	}
}

// BenchmarkProtobufCodec benchmark test for protobufCodec
func BenchmarkProtobufCodec(t *testing.B) {
	student := &pb.Student{
		Name:   "codec",
		Male:   true,
		Scores: []int32{99, 100, 99},
	}
	for i := 0; i < t.N; i++ {
		protobufCodec(student)
	}
}
