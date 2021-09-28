package net

import (
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
)

func LookupIP(domain string) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		log.Fatal().Err(err)
	}
	fmt.Println(domain)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func LookupAddr(addr string) {
	names, err := net.LookupAddr(addr)
	if err != nil {
		log.Fatal().Err(err)
	}
	fmt.Println(addr)
	for _, n := range names {
		fmt.Println(n)
	}
}
