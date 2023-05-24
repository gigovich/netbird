package uspfilter

import (
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"

	fw "github.com/netbirdio/netbird/client/firewall"
)

// Rule to handle management of rules
type Rule struct {
	id         string
	ip         net.IP
	ipLayer    gopacket.LayerType
	protoLayer gopacket.LayerType
	direction  fw.RuleDirection
	sPort      uint16
	dPort      uint16
	drop       bool
	comment    string

	udpHook func(*layers.UDP) bool
}

// GetRuleID returns the rule id
func (r *Rule) GetRuleID() string {
	return r.id
}
