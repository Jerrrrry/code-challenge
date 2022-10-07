package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/weppos/dnscaa"
)

type Query struct {
	Domain string `json:"domain"`
	Record string `json:"record"`
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/query", func(c *gin.Context) {
		var newQuery Query
		if err := c.BindJSON(&newQuery); err != nil {
			return
		}
		results := make([]string, 0)
		switch newQuery.Record {
		case "AAAA":
			records, _ := net.LookupIP(newQuery.Domain)
			for _, ip := range records {
				fmt.Println(ip)
				results = append(results, "Data : ")
				results = append(results, ip.String())
			}
		case "A":
			records, _ := net.LookupIP(newQuery.Domain)
			for _, ip := range records {
				fmt.Println(ip)
				results = append(results, "Data : ")
				results = append(results, ip.String())
			}

		case "TXT":
			records, _ := net.LookupTXT(newQuery.Domain)

			for _, txt := range records {
				fmt.Println(txt)
				results = append(results, "Value : ")
				results = append(results, txt)
			}

		case "MX":
			records, _ := net.LookupMX(newQuery.Domain)
			for _, mx := range records {
				fmt.Println(mx.Host, mx.Pref)
				results = append(results, "Exchange : ")
				results = append(results, fmt.Sprintf("%s", mx.Host))
				results = append(results, "Preference : ")
				results = append(results, fmt.Sprintf("%d", mx.Pref))
			}

		case "NS":
			records, _ := net.LookupNS(newQuery.Domain)
			for _, ns := range records {
				fmt.Println(ns)
				results = append(results, "Target : ")
				results = append(results, ns.Host)
			}
		case "CNAME":
			record, _ := net.LookupCNAME(newQuery.Domain)
			results = append(results, record)
			fmt.Println(record)
		case "SRV":
			cname, srvs, err := net.LookupSRV("xmpp-server", "tcp", newQuery.Domain)
			if err != nil {
				break
			}
			fmt.Printf("\ncname: %s \n\n", cname)
			results = append(results, fmt.Sprintf("\ncname: %s \n\n", cname))
			for _, srv := range srvs {
				fmt.Printf("%v:%v:%d:%d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
				results = append(results, fmt.Sprintf("%v:%v:%d:%d\n", srv.Target, srv.Port, srv.Priority, srv.Weight))
			}
		case "PTR":
			records, _ := net.LookupAddr(newQuery.Domain)
			for _, record := range records {
				fmt.Println(record)
				results = append(results, record)
			}
		// case "TLSA":
		// 	resolver, err := dane.GetResolver()
		// 	tlsa, err := dane.GetTLSA(resolver, newQuery.Domain, 443)
		// 	iplist, err := dane.GetAddresses(resolver, newQuery.Domain, true)
		// 	for _, ip := range iplist {
		// 		daneconfig := dane.NewConfig(newQuery.Domain, ip, 443)
		// 		daneconfig.SetTLSA(tlsa)
		// 		conn, err := dane.DialTLS(daneconfig)
		// 		if err != nil {
		// 			results = append(results, fmt.Sprintf("Result: FAILED: %s\n", err.Error()))
		// 			continue
		// 		}
		// 		if daneconfig.Okdane {
		// 			fmt.Printf("Result: DANE OK\n")
		// 			results = append(results, "Result: DANE OK")
		// 		} else if daneconfig.Okpkix {
		// 			fmt.Printf("Result: PKIX OK\n")
		// 			results = append(results, "Result: PKIX OK")
		// 		} else {
		// 			fmt.Printf("Result: FAILED\n")
		// 			results = append(results, "Result: FAILED")
		// 		}
		// 		//
		// 		// do some stuff with the obtained TLS connection here
		// 		//
		// 		conn.Close()
		// 	}
		case "CAA":
			records, err := dnscaa.Lookup(newQuery.Domain)

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("%d records found\n", len(records))

			for _, record := range records {
				fmt.Println(record)
				results = append(results, "Data:")
				results = append(results, fmt.Sprintf("%s", record))
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"results": results,
		})

	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
