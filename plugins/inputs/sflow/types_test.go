package sflow

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRawPacketHeaderFlowData(t *testing.T) {
	h := rawPacketHeaderFlowData{
		HeaderProtocol: headerProtocolTypeEthernetISO88023,
		FrameLength:    64,
		Bytes:          64,
		StrippedOctets: 0,
		HeaderLength:   0,
		Header:         nil,
	}
	tags := h.getTags()
	fields := h.getFields()

	require.NotNil(t, fields)
	require.NotNil(t, tags)
	require.Contains(t, tags, "header_protocol")
	require.Len(t, tags, 1)
}

// process a raw ethernet packet without any encapsulated protocol
func TestEthHeader(t *testing.T) {
	h := ethHeader{
		DestinationMAC:        [6]byte{0xca, 0xff, 0xee, 0xff, 0xe, 0x0},
		SourceMAC:             [6]byte{0xde, 0xad, 0xbe, 0xef, 0x0, 0x0},
		TagProtocolIdentifier: 0x88B5, // IEEE Std 802 - Local Experimental Ethertype
		TagControlInformation: 0,
		EtherTypeCode:         0,
		EtherType:             "",
		IPHeader:              nil,
	}
	tags := h.getTags()
	fields := h.getFields()

	require.NotNil(t, fields)
	require.NotNil(t, tags)
}
