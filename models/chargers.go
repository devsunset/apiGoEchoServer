package models

type Region struct {
	Region string `json:"Region"`
}

type Regions struct {
	Result_Code    string   `json:"RESULT_CODE"`
	Result_Message string   `json:"RESULT_MESSAGE"`
	Regions        []Region `json:"DATA"`
}

type Country struct {
	Country string `json:"Country"`
}

type Countrys struct {
	Result_Code    string    `json:"RESULT_CODE"`
	Result_Message string    `json:"RESULT_MESSAGE"`
	Countrys       []Country `json:"DATA"`
}

type City struct {
	City string `json:"City"`
}

type Citys struct {
	Result_Code    string `json:"RESULT_CODE"`
	Result_Message string `json:"RESULT_MESSAGE"`
	Citys          []City `json:"DATA"`
}

type Location_type struct {
	Location_type string `json:"Location_type"`
}

type Location_types struct {
	Result_Code    string          `json:"RESULT_CODE"`
	Result_Message string          `json:"RESULT_MESSAGE"`
	Location_types []Location_type `json:"DATA"`
}

type Location struct {
	Sl_translate             string `json:"Sl_translate"`
	Address_line_1           string `json:"Address_line_1"`
	Address_line_2           string `json:"Address_line_2"`
	Address_notes            string `json:"Address_notes"`
	Address                  string `json:"Address"`
	Amenities                string `json:"Amenities"`
	Baidu_lat                string `json:"Baidu_lat"`
	Baidu_lng                string `json:"Baidu_lng"`
	Chargers                 string `json:"Chargers"`
	City                     string `json:"City"`
	Common_name              string `json:"Common_name"`
	Country                  string `json:"Country"`
	Destination_charger_logo string `json:"Destination_charger_logo"`
	Destination_website      string `json:"Destination_website"`
	Directions_link          string `json:"Directions_link"`
	Emails                   string `json:"Emails"`
	Geocode                  string `json:"Geocode"`
	Hours                    string `json:"Hours"`
	Is_gallery               string `json:"Is_gallery"`
	Kiosk_pin_x              string `json:"Kiosk_pin_x"`
	Kiosk_pin_y              string `json:"Kiosk_pin_y"`
	Kiosk_zoom_pin_x         string `json:"Kiosk_zoom_pin_x"`
	Kiosk_zoom_pin_y         string `json:"Kiosk_zoom_pin_y"`
	Latitude                 string `json:"Latitude"`
	Location_id              string `json:"Location_id"`
	Location_type            string `json:"Location_type"`
	Longitude                string `json:"Longitude"`
	Nid                      string `json:"Nid"`
	Open_soon                string `json:"Open_soon"`
	Path                     string `json:"Path"`
	Postal_code              string `json:"Postal_code"`
	Province_state           string `json:"Province_state"`
	Region                   string `json:"Region"`
	Sales_phone              string `json:"Sales_phone"`
	Sales_representative     string `json:"Sales_representative"`
	Sub_region               string `json:"Sub_region"`
	Title                    string `json:"Title"`
	Trt_id                   string `json:"Trt_id"`
	Created_dttm             string `json:"Created_dttm"`
}

type Locations struct {
	Result_Code    string     `json:"RESULT_CODE"`
	Result_Message string     `json:"RESULT_MESSAGE"`
	Locations      []Location `json:"DATA"`
}
