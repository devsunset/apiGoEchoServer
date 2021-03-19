package models

type Region struct {
	Region string `json:"region"`
}

type Regions struct {
	Result_Code    string   `json:"RESULT_CODE"`
	Result_Message string   `json:"RESULT_MESSAGE"`
	Regions        []Region `json:"DATA"`
}

type Country struct {
	Country string `json:"country"`
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
	Sl_translate             string `json:"sl_translate"`
	Address_line_1           string `json:"address_line_1"`
	Address_line_2           string `json:"address_line_2"`
	Address_notes            string `json:"address_notes"`
	Address                  string `json:"address"`
	Amenities                string `json:"amenities"`
	Baidu_lat                string `json:"baidu_lat"`
	Baidu_lng                string `json:"baidu_lng"`
	Chargers                 string `json:"chargers"`
	City                     string `json:"city"`
	Common_name              string `json:"common_name"`
	Country                  string `json:"country"`
	Destination_charger_logo string `json:"destination_charger_logo"`
	Destination_website      string `json:"destination_website"`
	Directions_link          string `json:"directions_link"`
	Emails                   string `json:"emails"`
	Geocode                  string `json:"geocode"`
	Hours                    string `json:"hours"`
	Is_gallery               string `json:"is_gallery"`
	Kiosk_pin_x              string `json:"kiosk_pin_x"`
	Kiosk_pin_y              string `json:"kiosk_pin_y"`
	Kiosk_zoom_pin_x         string `json:"kiosk_zoom_pin_x"`
	Kiosk_zoom_pin_y         string `json:"kiosk_zoom_pin_y"`
	Latitude                 string `json:"latitude"`
	Location_id              string `json:"location_id"`
	Location_type            string `json:"location_type"`
	Longitude                string `json:"longitude"`
	Nid                      string `json:"nid"`
	Open_soon                string `json:"open_soon"`
	Path                     string `json:"path"`
	Postal_code              string `json:"postal_code"`
	Province_state           string `json:"province_state"`
	Region                   string `json:"region"`
	Sales_phone              string `json:"sales_phone"`
	Sales_representative     string `json:"sales_representative"`
	Sub_region               string `json:"sub_region"`
	Title                    string `json:"title"`
	Trt_id                   string `json:"trt_id"`
	Created_dttm             string `json:"created_dttm"`
}

type Locations struct {
	Result_Code    string     `json:"RESULT_CODE"`
	Result_Message string     `json:"RESULT_MESSAGE"`
	Locations      []Location `json:"DATA"`
}
