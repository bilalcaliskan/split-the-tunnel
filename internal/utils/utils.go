package utils

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func ResolveDomain(domain string) ([]string, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}

	var ipStrings []string
	for _, ip := range ips {
		ipStrings = append(ipStrings, ip.String())
	}

	return ipStrings, nil
}

func GetDefaultNonVPNGateway() (string, error) {
	file, err := os.Open("/proc/net/route")
	if err != nil {
		return "", errors.Wrap(err, "failed to open routing info file")
	}
	defer file.Close()

	var bestGateway string
	highestMetric := -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) >= 8 && fields[1] == "00000000" {
			metric, err := strconv.Atoi(fields[6])
			if err != nil {
				continue // Ignore lines with invalid metric
			}

			// Looking for the highest metric, assuming it's non-VPN
			if metric > highestMetric {
				highestMetric = metric

				bestGateway, err = parseHexIP(fields[2])
				if err != nil {
					continue // Ignore lines with invalid gateway IPs
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", errors.Wrap(err, "error reading file")
	}

	if bestGateway == "" {
		return "", fmt.Errorf("non-VPN gateway not found")
	}

	return bestGateway, nil
}

func parseHexIP(hexStr string) (string, error) {
	ipBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", errors.Wrap(err, "failed to decode hex string")
	}

	if len(ipBytes) != 4 {
		return "", fmt.Errorf("invalid IP length: %d", len(ipBytes))
	}

	// Reverse the byte order (little endian)
	for i, j := 0, len(ipBytes)-1; i < j; i, j = i+1, j-1 {
		ipBytes[i], ipBytes[j] = ipBytes[j], ipBytes[i]
	}

	return fmt.Sprintf("%d.%d.%d.%d", ipBytes[0], ipBytes[1], ipBytes[2], ipBytes[3]), nil
}

/*func addRoute(ip, gateway string) error {
	cmd := exec.Command("sudo", "ip", "route", "add", ip, "via", gateway)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to add route for %s: %w", ip, err)
	}
	return nil
}*/
