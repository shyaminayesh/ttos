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

	// tun.Open()

	var tunDNS string
	flag.StringVar(&tunDNS, "tun-dns", "8.8.8.8,8.8.4.4", "tun dns servers")
	dnsServers := strings.Split(tunDNS, ",")

	f, err := tun.OpenTunDevice("tun0", "172.16.10.100", "172.16.10.200", "255.255.255.0", dnsServers)
	if err != nil {
		fmt.Println(err)
	}

	tun := gotun2socks.New(f, "127.0.0.1:1337", dnsServers, false, false)

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
