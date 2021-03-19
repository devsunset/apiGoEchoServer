package models

type Region struct {
	Region string `json:"REGION"`
}

type Regions struct {
	Result_Code    string   `json:"RESULT_CODE"`
	Result_Message string   `json:"RESULT_MESSAGE"`
	Regions        []Region `json:"DATA"`
}

type Country struct {
	Country string `json:"COUNTRY"`
}

type Countrys struct {
	Result_Code    string    `json:"RESULT_CODE"`
	Result_Message string    `json:"RESULT_MESSAGE"`
	Countrys       []Country `json:"DATA"`
}

type City struct {
	City string `json:"CITY"`
}

type Citys struct {
	Result_Code    string `json:"RESULT_CODE"`
	Result_Message string `json:"RESULT_MESSAGE"`
	Citys          []City `json:"DATA"`
}

type Location_type struct {
	Location_type string `json:"LOCATION_TYPE"`
}

type Location_types struct {
	Result_Code    string          `json:"RESULT_CODE"`
	Result_Message string          `json:"RESULT_MESSAGE"`
	Location_types []Location_type `json:"DATA"`
}

type Location struct {
	Sl_translate             string `json:"SL_TRANSLATE"`
	Address_line_1           string `json:"ADDRESS_LINE_1"`
	Address_line_2           string `json:"ADDRESS_LINE_2"`
	Address_notes            string `json:"ADDRESS_NOTES"`
	Address                  string `json:"ADDRESS"`
	Amenities                string `json:"AMENITIES"`
	Baidu_lat                string `json:"BAIDU_LAT"`
	Baidu_lng                string `json:"BAIDU_LNG"`
	Chargers                 string `json:"CHARGERS"`
	City                     string `json:"CITY"`
	Common_name              string `json:"COMMON_NAME"`
	Country                  string `json:"COUNTRY"`
	Destination_charger_logo string `json:"DESTINATION_CHARGER_LOGO"`
	Destination_website      string `json:"DESTINATION_WEBSITE"`
	Directions_link          string `json:"DIRECTIONS_LINK"`
	Emails                   string `json:"EMAILS"`
	Geocode                  string `json:"GEOCODE"`
	Hours                    string `json:"HOURS"`
	Is_gallery               string `json:"IS_GALLERY"`
	Kiosk_pin_x              string `json:"KIOSK_PIN_X"`
	Kiosk_pin_y              string `json:"KIOSK_PIN_Y"`
	Kiosk_zoom_pin_x         string `json:"KIOSK_ZOOM_PIN_X"`
	Kiosk_zoom_pin_y         string `json:"KIOSK_ZOOM_PIN_Y"`
	Latitude                 string `json:"LATITUDE"`
	Location_id              string `json:"LOCATION_ID"`
	Location_type            string `json:"LOCATION_TYPE"`
	Longitude                string `json:"LONGITUDE"`
	Nid                      string `json:"NID"`
	Open_soon                string `json:"OPEN_SOON"`
	Path                     string `json:"PATH"`
	Postal_code              string `json:"POSTAL_CODE"`
	Province_state           string `json:"PROVINCE_STATE"`
	Region                   string `json:"REGION"`
	Sales_phone              string `json:"SALES_PHONE"`
	Sales_representative     string `json:"SALES_REPRESENTATIVE"`
	Sub_region               string `json:"SUB_REGION"`
	Title                    string `json:"TITLE"`
	Trt_id                   string `json:"TRT_ID"`
	Created_dttm             string `json:"CREATED_DTTM"`
}

type Locations struct {
	Result_Code    string     `json:"RESULT_CODE"`
	Result_Message string     `json:"RESULT_MESSAGE"`
	Locations      []Location `json:"DATA"`
}
