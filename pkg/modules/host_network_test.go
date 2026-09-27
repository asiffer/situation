package modules

import (
	"net"
	"testing"

	"github.com/asiffer/situation/pkg/models"
)

func TestFindHostNIC(t *testing.T) {
	byMAC := &models.NetworkInterface{ID: 1, Name: "eth0", MAC: "02:00:00:00:00:01"}
	withoutMAC := &models.NetworkInterface{ID: 2, Name: "lo"}
	hostNICMap := buildHostNICMap([]*models.NetworkInterface{byMAC, withoutMAC})

	tests := []struct {
		name  string
		iface net.Interface
		want  *models.NetworkInterface
	}{
		{
			name: "same MAC after interface rename",
			iface: net.Interface{
				Name:         "renamed0",
				HardwareAddr: net.HardwareAddr{0x02, 0x00, 0x00, 0x00, 0x00, 0x01},
			},
			want: byMAC,
		},
		{
			name: "changed MAC with same interface name",
			iface: net.Interface{
				Name:         "eth0",
				HardwareAddr: net.HardwareAddr{0x02, 0x00, 0x00, 0x00, 0x00, 0x02},
			},
			want: byMAC,
		},
		{
			name:  "interface without MAC",
			iface: net.Interface{Name: "lo"},
			want:  withoutMAC,
		},
		{
			name:  "new interface",
			iface: net.Interface{Name: "eth1", HardwareAddr: net.HardwareAddr{0x02, 0x00, 0x00, 0x00, 0x00, 0x03}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := findHostNIC(test.iface, hostNICMap); got != test.want {
				t.Errorf("findHostNIC() = %p, want %p", got, test.want)
			}
		})
	}
}

func TestRefreshHostNICIPs(t *testing.T) {
	nic := &models.NetworkInterface{IP: []string{"192.168.122.1"}}
	seen := resetNICIPs(nic)

	if len(nic.IP) != 0 {
		t.Fatalf("resetNICIPs() left stale addresses: %v", nic.IP)
	}
	if !appendNICIP(nic, seen, "172.17.0.1") {
		t.Fatal("appendNICIP() rejected a new address")
	}
	if appendNICIP(nic, seen, "172.17.0.1") {
		t.Fatal("appendNICIP() accepted a duplicate address")
	}
	if len(nic.IP) != 1 || nic.IP[0] != "172.17.0.1" {
		t.Fatalf("refreshed addresses = %v, want [172.17.0.1]", nic.IP)
	}
}
