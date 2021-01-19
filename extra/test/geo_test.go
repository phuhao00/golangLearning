package test
//
//import (
//	"fmt"
//	"github.com/oschwald/geoip2-golang"
//	"github.com/oschwald/maxminddb-golang"
//	"log"
//	"net"
//	"testing"
//)
//
//func TestGeo(t *testing.T) {
//	db, err := geoip2.Open("./GeoLite2-City.mmdb")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//	// If you are using strings that may be invalid, check that ip is not nil
//	ip := net.ParseIP("81.2.69.142")
//	record, err := db.City(ip)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["pt-BR"])
//	if len(record.Subdivisions) > 0 {
//		fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
//	}
//	fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
//	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
//	fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
//	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
//	// Output:
//	// Portuguese (BR) city name: Londres
//	// English subdivision name: England
//	// Russian country name: Великобритания
//	// ISO country code: GB
//	// Time zone: Europe/London
//	// Coordinates: 51.5142, -0.0931
//}
//func ExampleReader_Lookup_struct() {
//	db, err := maxminddb.Open("test-data/test-data/GeoIP2-City-Test.mmdb")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	ip := net.ParseIP("81.2.69.142")
//
//	var record struct {
//		Country struct {
//			ISOCode string `maxminddb:"iso_code"`
//		} `maxminddb:"country"`
//	} // Or any appropriate struct
//
//	err = db.Lookup(ip, &record)
//	if err != nil {
//		log.Panic(err)
//	}
//	fmt.Print(record.Country.ISOCode)
//	// Output:
//	// GB
//}
