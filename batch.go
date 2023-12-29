/*
	https://www.sohamkamani.com/golang/parsing-json/
	https://www.sohamkamani.com/golang/sql-transactions/
	https://golangdocs.com/golang-postgresql-example
	go get github.com/lib/pq
        go get github.com/tkanos/gonfig
        go get github.com/sirupsen/logrus
*/

package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "crypto/tls"
	// "net/http"
	"os"
	"strings"
	
	_ "github.com/lib/pq"

	"apiServer/config"
	"github.com/sirupsen/logrus"
)

type Location struct {
	Sl_translate             string
	Address_line_1           string
	Address_line_2           string
	Address_notes            string
	Address                  string
	Amenities                string
	Baidu_lat                string
	Baidu_lng                string
	Chargers                 string
	City                     string
	Common_name              string
	Country                  string
	Destination_charger_logo string
	Destination_website      string
	Directions_link          string
	Emails                   []Emails_Data
	Geocode                  string
	Hours                    string
	Is_gallery               string
	Kiosk_pin_x              string
	Kiosk_pin_y              string
	Kiosk_zoom_pin_x         string
	Kiosk_zoom_pin_y         string
	Latitude                 string
	Location_id              string
	Location_type            []string
	Longitude                string
	Nid                      string
	Open_soon                string
	Path                     string
	Postal_code              string
	Province_state           string
	Region                   string
	Sales_phone              []Sales_Phone_Data
	Sales_representative     string
	Sub_region               string
	Title                    string
	Trt_id                   string
}

type Emails_Data struct {
	Label string
	Email string
}

type Sales_Phone_Data struct {
	Label      string
	Number     string
	Line_below string
}

func main() {

	var log = logrus.New()

	log.Out = os.Stdout

	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true

	file, err := os.OpenFile("logs/batch.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	log.Info("##############################")

	configuration := config.GetConfig()

	log.Info("0. ---------- : batch process start")
	//------------------------------------------------------------
	// https request
	//------------------------------------------------------------
	// resp, err := http.Get(configuration.TESLA_REST_API_URL)
	// if err != nil {
	// 	log.Fatal(err)
	// 	panic(err)
	// }

	// defer resp.Body.Close()

	// data, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// 	panic(err)
	// }

	//------------------------------------------------------------
	// https request - TLS 설정에서 인증서 검증 무시
	//------------------------------------------------------------
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }

	// // 사용자 정의 Transport를 사용하여 HTTP 클라이언트 생성
	// client := &http.Client{Transport: tr}

	// // HTTPS URL에 GET 요청 보내기
	// resp, err := client.Get(configuration.TESLA_REST_API_URL)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// defer resp.Body.Close()

	// // 응답 본문 읽기
	// data, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	//------------------------------------------------------------
	// read json file (all-locations.json)
	//------------------------------------------------------------
	// JSON 파일 읽기
	filePath := "all-locations.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Failed to read JSON file:", err)
		return
	}

	//log.Printf("%s\n", string(data))

	log.Println("1. ---------- : tesla rest api call end")
	var location []Location
	json.Unmarshal([]byte(string(data)), &location)

	log.Println("2. ---------- : rest api data string convert to json object")
	if len(location) > 0 {
		log.Println("3. ---------- : db process start")
		// Create a new connection to our database
		connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", configuration.DB_HOST, configuration.DB_PORT, configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_NAME)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()

		// Create a new context, and begin a transaction
		ctx := context.Background()
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		// `tx` is an instance of `*sql.Tx` through which we can execute our queries
		_, err = tx.ExecContext(ctx, "DELETE FROM TB_TESLA_LOCATION")
		if err != nil {
			log.Fatal(err)
			tx.Rollback()
			return
		}

		sqltext := `INSERT INTO public.tb_tesla_location
					(sl_translate, address_line_1, address_line_2, address_notes, address, amenities, baidu_lat, baidu_lng, chargers, city
					, common_name, country, destination_charger_logo, destination_website, directions_link, emails, geocode, hours, is_gallery
					, kiosk_pin_x, kiosk_pin_y, kiosk_zoom_pin_x, kiosk_zoom_pin_y, latitude, location_id, location_type, longitude, nid, open_soon
					, path, postal_code, province_state, region, sales_phone, sales_representative, sub_region, title, trt_id, created_dttm)
					VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25,$26,$27,$28, $29, $30, $31, $32, $33, $34, $35,$36, $37, $38, now())`

		for _, loc := range location {

			emailData, _ := json.Marshal(loc.Emails)
			salesPhoneData, _ := json.Marshal(loc.Sales_phone)

			// Here, the query is executed on the transaction instance, and not applied to the database yet
			_, err = tx.ExecContext(ctx, sqltext, loc.Sl_translate, loc.Address_line_1, loc.Address_line_2, loc.Address_notes, loc.Address, loc.Amenities, loc.Baidu_lat, loc.Baidu_lng, loc.Chargers, loc.City, loc.Common_name, loc.Country, loc.Destination_charger_logo, loc.Destination_website, loc.Directions_link, string(emailData), loc.Geocode, loc.Hours, loc.Is_gallery, loc.Kiosk_pin_x, loc.Kiosk_pin_y, loc.Kiosk_zoom_pin_x, loc.Kiosk_zoom_pin_y, loc.Latitude, loc.Location_id, strings.Join(loc.Location_type, "|"), loc.Longitude, loc.Nid, loc.Open_soon, loc.Path, loc.Postal_code, loc.Province_state, loc.Region, string(salesPhoneData), loc.Sales_representative, loc.Sub_region, loc.Title, loc.Trt_id)
			if err != nil {
				log.Println(loc)
				log.Fatal(err)
				tx.Rollback()
				return
			}
		}

		log.Println("4. ---------- : db process end")
		// Finally, if no errors are recieved from the queries, commit the transaction
		// this applies the above changes to our database
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("5. ---------- : batch process end")
	}else{
	 log.Println("No Data")
	}
}
