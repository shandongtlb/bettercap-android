package network

import (
	"bufio"
	"encoding/binary"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/bettercap/bettercap/v2/core"

	"github.com/evilsocket/islazy/str"
)

// Hi, i'm Android and my mum said I'm special.
//
// The legacy implementation guessed the gateway from `getprop net.dns1`,
// which assumes DNS == gateway. That assumption breaks on modern Android
// (9+): the global net.dns1 property is deprecated/empty (per-network
// LinkProperties + Private DNS), so gateway detection silently fails and
// bettercap falls back to treating the interface itself as the gateway,
// making ARP spoofing a no-op. We now read the real routing table from
// /proc/net/route (always present, no external command), and only fall
// back to getprop for old devices where the parse yields nothing.
func FindGateway(iface *Endpoint) (*Endpoint, error) {
	if gw := defaultGatewayFromProcRoute(iface.Name()); gw != "" {
		mac, err := ArpLookup(iface.Name(), gw, false)
		if err != nil {
			return nil, err
		}
		return NewEndpoint(gw, mac), nil
	}

	// Fallback: legacy getprop heuristic for older devices.
	if output, err := core.Exec("getprop", []string{"net.dns1"}); err == nil {
		gw := str.Trim(output)
		if IPv4Validator.MatchString(gw) {
			mac, err := ArpLookup(iface.Name(), gw, false)
			if err != nil {
				return nil, err
			}
			return NewEndpoint(gw, mac), nil
		}
	}

	return nil, ErrNoGateway
}

// defaultGatewayFromProcRoute parses /proc/net/route and returns the IPv4
// default gateway (destination 0.0.0.0 with the RTF_GATEWAY flag) for the
// given interface, or "" if none is found.
//
// /proc/net/route columns (tab separated):
//
//	Iface Destination Gateway Flags RefCnt Use Metric Mask MTU Window IRTT
//
// Destination/Gateway/Mask are little-endian hex IPv4 addresses.
func defaultGatewayFromProcRoute(ifaceName string) string {
	file, err := os.Open("/proc/net/route")
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Skip header line.
	if scanner.Scan() {
		_ = scanner.Text()
	}

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 4 {
			continue
		}

		iface := fields[0]
		destination := fields[1]
		gateway := fields[2]
		flags, err := strconv.ParseInt(fields[3], 16, 64)
		if err != nil {
			continue
		}

		const (
			rtfUp      = 0x0001
			rtfGateway = 0x0002
		)

		// Default route: destination 0.0.0.0, up, and an actual gateway.
		if destination != "00000000" {
			continue
		}
		if flags&rtfUp == 0 || flags&rtfGateway == 0 {
			continue
		}
		if ifaceName != "" && iface != ifaceName {
			continue
		}

		if ip := hexToIPv4(gateway); ip != "" {
			return ip
		}
	}

	return ""
}

// hexToIPv4 converts a little-endian 8-char hex string (as found in
// /proc/net/route) into a dotted-quad IPv4 string. Returns "" on error
// or for the all-zero address.
func hexToIPv4(hexAddr string) string {
	value, err := strconv.ParseUint(hexAddr, 16, 32)
	if err != nil || value == 0 {
		return ""
	}
	buf := make([]byte, 4)
	// /proc/net/route stores addresses in host (little-endian) byte order.
	binary.LittleEndian.PutUint32(buf, uint32(value))
	return net.IPv4(buf[0], buf[1], buf[2], buf[3]).String()
}
