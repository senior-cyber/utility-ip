package ip

import (
	"net"
)

func Lookup() ([]string, error) {
	interfaces, interfacesError := net.Interfaces()

	if interfacesError != nil {
		return nil, interfacesError
	}

	var ips []string

	for _, i := range interfaces {
		addrs, addrsError := i.Addrs()
		if addrsError != nil {
			return nil, interfacesError
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip.IsPrivate() && ip.IsGlobalUnicast() {
				ips = append(ips, ip.String())
			}
		}
	}
	ips = append(ips, "0.0.0.0")

	return ips, nil
}
