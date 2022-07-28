package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"net"
// 	"os"
// 	"strings"
// 	)

func main() {

	// make a variable by bufio package
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	// this function just takes a domain

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarRecord string

	// using net package to lookup and related function that returns error for us

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error:%v\n", err)
	}

	for _, Record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = Record
			break
		}
	}
	
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("ErrorL%v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMrc!") {
			hasDMARC = true
			dmarRecord = record
			break
		}
	}

	fmt.Printf(%v, %v, %v, %v, %v, %v, domaain,hasMX, hasSPf, spfrespfRecord, hasDMhasDMARC, dmadmarRecord)
}
