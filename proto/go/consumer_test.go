package _go

import (
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestSubscribeInitAutoAckWireCompatibility(t *testing.T) {
	for _, autoAck := range []bool{false, true} {
		original := &SubscribeInit{Topic: "orders", GroupId: "workers", AutoAck: autoAck}
		data, err := proto.Marshal(original)
		if err != nil {
			t.Fatal(err)
		}
		var decoded SubscribeInit
		if err := proto.Unmarshal(data, &decoded); err != nil {
			t.Fatal(err)
		}
		if !proto.Equal(original, &decoded) {
			t.Fatalf("round trip: got %v, want %v", &decoded, original)
		}
	}
	// Field 4 is a varint bool; older handshakes omit it and stay manual.
	var decoded SubscribeInit
	if err := proto.Unmarshal([]byte{0x20, 0x01}, &decoded); err != nil {
		t.Fatal(err)
	}
	if !decoded.GetAutoAck() {
		t.Fatal("field 4 did not enable auto-ack")
	}
	decoded.Reset()
	if decoded.GetAutoAck() || (*SubscribeInit)(nil).GetAutoAck() {
		t.Fatal("auto-ack must default to false")
	}
}
