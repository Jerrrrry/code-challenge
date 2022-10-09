package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/babolivier/go-doh-client"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/weppos/dnscaa"
)

type Query struct {
	Domain string `json:"domain"`
	Record string `json:"record"`
}

func queryAAAA(results []string, resolver doh.Resolver, newQuery Query) []string {
	results = append(results, "####################AAAA RECORD START###############################")
	records, ttls, err := resolver.LookupAAAA(newQuery.Domain)
	if err == nil {
		for i, record := range records {
			results = append(results, fmt.Sprintf("TTL : %d seconds", ttls[i]))
			results = append(results, "Data : ")
			results = append(results, record.IP6)
			results = append(results, "-------------------------------------------")
		}
	}
	results = append(results, "####################AAAA RECORD END###############################")
	return results

}

func queryA(results []string, resolver doh.Resolver, newQuery Query) []string {
	records, ttls, err := resolver.LookupA(newQuery.Domain)
	results = append(results, "####################A RECORD START###############################")
	if err == nil {
		for i, record := range records {
			results = append(results, fmt.Sprintf("TTL : %d seconds", ttls[i]))
			results = append(results, "Data : ")
			results = append(results, record.IP4)
			results = append(results, "-------------------------------------------")
		}
	}
	results = append(results, "####################A RECORD END###############################")
	return results
}

func querySOA(results []string, resolver doh.Resolver, newQuery Query) []string {
	// Perform a SOA lookup on newQuery.Domain
	records, ttls, err := resolver.LookupSOA(newQuery.Domain)
	results = append(results, "####################SOA RECORD START###############################")
	if err == nil {
		for i, record := range records {
			results = append(results, fmt.Sprintf("TTL : %d seconds", ttls[i]))
			fmt.Println(record.PrimaryNS)
			results = append(results, fmt.Sprintf("---Primary NS :%s", record.PrimaryNS))
			fmt.Println(record.RespMailbox)
			results = append(results, fmt.Sprintf("---RespMailbox :%s", record.RespMailbox))
			fmt.Println(record.Serial)
			results = append(results, fmt.Sprintf("---Serial :%d", record.Serial))
			fmt.Println(record.Refresh)
			results = append(results, fmt.Sprintf("---Refresh :%d", record.Refresh))
			fmt.Println(record.Retry)
			results = append(results, fmt.Sprintf("---Retry :%d", record.Retry))
			fmt.Println(record.Expire)
			results = append(results, fmt.Sprintf("---Expire :%d", record.Expire))
			fmt.Println(record.Minimum)
			results = append(results, fmt.Sprintf("---TTL :%d seconds", record.Minimum))

			results = append(results, "-------------------------------------------")
		}
	}
	results = append(results, "####################SOA RECORD END###############################")
	return results
}
func querySRV(results []string, resolver doh.Resolver, newQuery Query) []string {
	//_xmpp-server._tcp.google.com.
	records, ttls, err := resolver.LookupSRV(newQuery.Domain)
	results = append(results, "####################SRV RECORD START###############################")
	if err == nil {
		for i, record := range records {
			results = append(results, fmt.Sprintf("TTL : %d seconds", ttls[i]))
			results = append(results, fmt.Sprintf("Target : %s   Port:%d   Priority:%d   Weight:%d", record.Target, record.Port, record.Priority, record.Weight))
			results = append(results, "-------------------------------------------")
		}
	}
	results = append(results, "####################SRV RECORD END###############################")
	return results
}

func queryTXT(results []string, resolver doh.Resolver, newQuery Query) []string {
	records, ttls, err := resolver.LookupTXT(newQuery.Domain)
	results = append(results, "####################TXT RECORD START###############################")
	if err == nil {
		for i, record := range records {
			results = append(results, fmt.Sprintf("TTL : %d seconds", ttls[i]))
			results = append(results, "Value : ")
			results = append(results, record.TXT)
			results = append(results, "-------------------------------------------")
		}
	}
	results = append(results, "####################TXT RECORD END###############################")
	return results
}

func queryMX(results []string, newQuery Query) []string {
	records, _ := net.LookupMX(newQuery.Domain)
	results = append(results, "####################MX RECORD START###############################")
	for _, mx := range records {
		fmt.Println(mx.Host, mx.Pref)
		results = append(results, "Exchange : ")
		results = append(results, fmt.Sprintf("%s", mx.Host))
		results = append(results, "Preference : ")
		results = append(results, fmt.Sprintf("%d", mx.Pref))
	}
	results = append(results, "####################MX RECORD END###############################")
	return results
}
func queryCNAME(results []string, newQuery Query) []string {
	results = append(results, "####################CNAME RECORD START###############################")
	record, _ := net.LookupCNAME(newQuery.Domain)
	results = append(results, record)
	results = append(results, "####################CNAME RECORD END###############################")
	return results
}

func queryPTR(results []string, newQuery Query) []string {
	results = append(results, "####################PTR RECORD START###############################")
	records, _ := net.LookupAddr(newQuery.Domain)
	for _, record := range records {
		fmt.Println(record)
		results = append(results, record)
	}
	results = append(results, "####################PTR RECORD START###############################")
	return results
}

func queryCAA(results []string, newQuery Query) []string {
	results = append(results, "####################CAA RECORD START###############################")
	records, err := dnscaa.Lookup(newQuery.Domain)

	if err == nil {
		fmt.Printf("%d records found\n", len(records))

		for _, record := range records {
			fmt.Println(record)
			results = append(results, "Data:")
			results = append(results, fmt.Sprintf("%s", record))
		}
	}
	results = append(results, "####################CAA RECORD END###############################")
	return results
}

func queryNS(results []string, resolver doh.Resolver, newQuery Query) []string {
	results = append(results, "####################NS RECORD START###############################")
	records, ttls, err := resolver.LookupNS(newQuery.Domain)
	if err == nil {
		for i, record := range records {
			results = append(results, fmt.Sprintf("TTL : %d seconds", ttls[i]))
			results = append(results, "Target : ")
			results = append(results, record.Host)
			results = append(results, "-------------------------------------------")
		}
	}
	results = append(results, "####################NS RECORD END###############################")
	return results
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	resolver := doh.Resolver{
		Host:  "8.8.8.8", // Change this with your favourite DoH-compliant resolver.
		Class: doh.IN,
	}
	r.POST("/query", func(c *gin.Context) {
		var newQuery Query
		if err := c.BindJSON(&newQuery); err != nil {
			return
		}
		results := make([]string, 0)
		switch newQuery.Record {
		case "AAAA":
			results = queryAAAA(results, resolver, newQuery)
		case "A":
			results = queryA(results, resolver, newQuery)
		case "TXT":
			results = queryTXT(results, resolver, newQuery)
		case "MX":
			results = queryMX(results, newQuery)
		case "NS":
			results = queryNS(results, resolver, newQuery)
		case "CNAME":
			results = queryCNAME(results, newQuery)
		case "SRV":
			results = querySRV(results, resolver, newQuery)
		case "PTR":
			results = queryPTR(results, newQuery)
		case "CAA":
			results = queryCAA(results, newQuery)
		case "SOA":
			results = querySOA(results, resolver, newQuery)
		case "ANY":
			results = queryA(results, resolver, newQuery)
			results = queryAAAA(results, resolver, newQuery)
			results = queryTXT(results, resolver, newQuery)
			results = queryMX(results, newQuery)
			results = queryNS(results, resolver, newQuery)
			results = queryCNAME(results, newQuery)
			results = querySRV(results, resolver, newQuery)
			results = queryPTR(results, newQuery)
			results = queryCAA(results, newQuery)
			results = querySOA(results, resolver, newQuery)

		}
		c.JSON(http.StatusOK, gin.H{
			"results": results,
		})

	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
