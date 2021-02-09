//@Kenji DURIEZ - [DeedWark] - 2020
//Resolve DNS w/ Domain (A, MX, TXT, DMARC, DKIM)
package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

const (
	bold   = "\033[1m"
	end    = "\033[00m"
	blue   = "\033[34m"
	red    = "\033[91m"
	yellow = "\033[93m"
	green  = "\033[32m"
)

func afinderRes(domain string, resolver string) {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, "udp", resolver+":53")
		},
	}
	ip, _ := r.LookupHost(context.Background(), domain)

	fmt.Println(bold + "A field:" + end)
	if len(ip) == 0 {
		fmt.Println("No DNS found")
	} else {
		for _, ip := range ip {
			fmt.Println(ip)
		}
	}
}

func afinder(domain string) {
	a, _ := net.LookupIP(domain)
	fmt.Println(bold + "A field:" + end)
	if len(a) == 0 {
		fmt.Println("No DNS found")
	} else {
		for _, ip := range a {
			fmt.Println(ip.String())
		}
	}
}

func mxfinderRes(domain string, resolver string) {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, "udp", resolver+":53")
		},
	}
	mxs, _ := r.LookupMX(context.Background(), domain)

	fmt.Println("")
	fmt.Println(bold + "MX fields:" + end)
	if len(mxs) == 0 {
		fmt.Println("No MX found")
	} else {
		for _, mx := range mxs {
			fmt.Println(mx.Pref, mx.Host)
		}
	}
}

func mxfinder(domain string) {
	mxs, _ := net.LookupMX(domain)
	fmt.Println("")
	fmt.Println(bold + "MX fields:" + end)
	if len(mxs) == 0 {
		fmt.Println("No MX found")
	} else {
		for _, mx := range mxs {
			fmt.Println(mx.Pref, mx.Host)
		}
	}
}

func txtfinder(domain string) {
	txts, _ := net.LookupTXT(domain)
	fmt.Println("")
	fmt.Println(bold + "TXT records:" + end)
	if len(txts) == 0 {
		fmt.Println("No TXT found")
	} else {
		for _, txt := range txts {
			fmt.Println(txt)
		}
	}
}

func dmarcfinder(domain string) {
	dmarc, _ := net.LookupTXT("_dmarc." + domain)
	fmt.Println("")
	fmt.Println(bold + "DMARC key:" + end)
	if len(dmarc) == 0 {
		fmt.Println("No DMARC key found")
	} else {
		for _, dmkey := range dmarc {
			fmt.Println(dmkey)
		}
	}
}

func dkimfinder(domain string, selector string) {
	dkim, _ := net.LookupTXT(selector + "._domainkey." + domain)
	fmt.Println("")
	fmt.Println(bold + "DKIM key:" + end)

	if selector == "" || selector == "google" {
		fmt.Println("Add a selector (ex: domain.com selector)")
		fmt.Println("Try with " + blue + "G" + red + "o" + yellow + "o" + blue + "g" + green + "l" + red + "e" + end + " as selector:" + "\n")
	}

	if len(dkim) == 0 && selector != "google" {
		fmt.Println("No DKIM key found with " + bold + selector + end + " as selector" + "\n")
	} else if len(dkim) == 0 && selector == "google" {
		fmt.Println("No DKIM key found" + "\n")
	} else {
		for _, dkimk := range dkim {
			fmt.Println(dkimk)
		}
	}
}

func main() {

	name := os.Args[0]

	help := "GODIG - Domain DNS Resolver in Golang" + "\r\n" +
		"        Usage:   " + name + " [domain] [selector | @IPresolver]" + "\r\n\r\n" +
		"        Example: " + name + " domain.com" + "\r\n" +
		"                 " + name + " domain.com mailjet" + "\r\n" +
		"                 " + name + " domain.com @8.8.8.8" + "\r\n\r\n" +
		"Use [" + name + " help] to show this message"

	flag.Parse()

	domain := flag.Arg(0)
	if domain == "" || domain == "help" {
		fmt.Println(help)
		os.Exit(0)
	}

	///////////////

	var resolver string
	var selector string
	var both string

	// If no selector or resolver
	both = flag.Arg(1)
	if both == "" {
		resolver = ""
		selector = "google"
	} else if both[0] == '@' {
		resolver = both[1:]
		selector = "google"
	} else {
		resolver = ""
		selector = both
	}

	if resolver != "" {
		afinderRes(domain, resolver)
	} else {
		afinder(domain)
	}

	if resolver != "" {
		mxfinderRes(domain, resolver)
	} else {
		mxfinder(domain)
	}

	txtfinder(domain)
	dmarcfinder(domain)
	dkimfinder(domain, selector)
}
