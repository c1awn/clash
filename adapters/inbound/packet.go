package inbound

import (
	"github.com/whojave/clash/component/socks5"
	C "github.com/whojave/clash/constant"
)

// PacketAdapter is a UDP Packet adapter for socks/redir/tun
type PacketAdapter struct {
	C.UDPPacket
	metadata *C.Metadata
}

// Metadata returns destination metadata
func (s *PacketAdapter) Metadata() *C.Metadata {
	return s.metadata
}

// NewPacket is PacketAdapter generator
func NewPacket(target socks5.Addr, packet C.UDPPacket, source C.Type, netType C.NetWork) *PacketAdapter {
	metadata := parseSocksAddr(target)
	metadata.NetWork = netType
	metadata.Type = source
	if ip, port, err := parseAddr(packet.LocalAddr().String()); err == nil {
		metadata.SrcIP = ip
		metadata.SrcPort = port
	}

	return &PacketAdapter{
		UDPPacket: packet,
		metadata:  metadata,
	}
}
