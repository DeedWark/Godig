//@Kenji DURIEZ - [DeedWark] - 2020
//Resolve DNS w/ Domain (A, MX, TXT, DMARC, DKIM)
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	bold   = "\033[1m"
	end    = "\033[00m"
	blue   = "\033[34m"
	red    = "\033[91m"
	yellow = "\033[93m"
	green  = "\033[32m"
)

func afinder(domain string) {
	a, _ := net.LookupIP(domain)
	fmt.Println(bold + "DNS (A):" + end)
	if len(a) == 0 {
		fmt.Println("No DNS found")
	} else {
		for _, ip := range a {
			fmt.Println(ip.String())
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
			mxRaw := strings.TrimRight(mx.Host, ".")
			fmt.Println(mx.Pref, mxRaw)
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
	if flag.Arg(1) == "" {
		fmt.Println("Add a selector (ex: domain.com selector)")
		fmt.Println("Try with " + blue + "G" + red + "o" + yellow + "o" + blue + "g" + green + "l" + red + "e" + end + " as selector:" + "\n")
	}
	if len(dkim) == 0 {
		fmt.Println("No DKIM key found" + "\n")
	} else {
		for _, dkimk := range dkim {
			fmt.Println(dkimk)
		}
	}
}

func main() {
	help :=
		`
MUDIG - Most Useful DIG commands in same script
        Usage:   digo [domain] [selector]

        Example: digo domain.com
                 digo domain.com protonmail

Use [digo help] to show this message
`
	flag.Parse()
	domain := flag.Arg(0)
	if domain == "" {
		fmt.Println(help)
		os.Exit(1)
	}

	selector := flag.Arg(1)
	if selector == "" {
		selector = "google"
	}

	afinder(domain)
	mxfinder(domain)
	txtfinder(domain)
	dmarcfinder(domain)
	dkimfinder(domain, selector)
}
