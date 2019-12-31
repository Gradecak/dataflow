package types

type WireTransportable interface {
	WireEncode() []byte
	WireDecode([]byte) WireTransportable
}
