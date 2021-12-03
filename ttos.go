package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/yinghuocho/gotun2socks"
	"github.com/yinghuocho/gotun2socks/tun"
)

func main() {

	config := Config()

	var tunDNS string
	flag.StringVar(&tunDNS, "tun-dns", "8.8.8.8,8.8.4.4", "tun dns servers")
	dnsServers := strings.Split(tunDNS, ",")

	f, err := tun.OpenTunDevice(config.Tunnel.Name, config.Tunnel.Address.Interface, config.Tunnel.Address.Gateway, config.Tunnel.Address.Mask, dnsServers)
	if err != nil {
		fmt.Println(err)
	}

	tun := gotun2socks.New(f, fmt.Sprintf("%s:%d", config.Proxy.Address, config.Proxy.Port), dnsServers, false, false)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		s := <-ch
		switch s {
		default:
			tun.Stop()
		}
	}()

	tun.Run()

}
