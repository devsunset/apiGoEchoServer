package controllers

import (
	"apiServer/db"
	"apiServer/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"
)

/////////////////////////////////////////////////////////////////
//
//	Chargers
//
/////////////////////////////////////////////////////////////////

// getLocationRegion godoc
// @Summary Tesla Chargers Location Region.
// @Description get tesla Chargers Location Region.
// @Tags API Tesla App
// @ID getLocationRegion
// @Accept */*
// @Produce json
// @Success 200 {object} models.Regions
// @Router /getLocationRegion [get]
func GetLocationRegion(c echo.Context) error {

	db := db.DbManager()

	sqlStatement := "SELECT DISTINCT REGION FROM TB_TESLA_LOCATION"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		//	return c.JSON(http.StatusCreated, u)
	}

	defer rows.Close()
	defer db.Close()

	result := models.Regions{Result_Code: "S", Result_Message: "Success"}

	for rows.Next() {
		region := models.Region{}
		err2 := rows.Scan(&region.Region)
		if err2 != nil {
			return err2
		}
		result.Regions = append(result.Regions, region)
	}
	return c.JSON(http.StatusCreated, result)
}

func GetLocationRegionBackUp(c echo.Context) error {

	db := db.DbManager()

	sqlStatement := "SELECT DISTINCT REGION FROM TB_TESLA_LOCATION"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		//	return c.JSON(http.StatusCreated, u)
	}

	defer rows.Close()
	defer db.Close()

	result := []map[string]interface{}{}
	cols, _ := rows.Columns()
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))

		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			return err
		}
		m := make(map[string]interface{})

		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		fmt.Println(m)
		result = append(result, m)

	}
	return c.JSON(http.StatusCreated, result)
}

// getLocationCountry godoc
// @Summary Tesla Chargers Location Country.
// @Description get tesla Chargers Location Country.
// @Tags API Tesla App
// @ID getLocationCountry
// @Accept */*
// @Produce json
// @Param region query string false "asia_pacific"
// @Success 200 {object} models.Countrys
// @Router /getLocationCountry [get]
func GetLocationCountry(c echo.Context) error {

	db := db.DbManager()

	region := c.QueryParam("region")

	sqlStatement := "SELECT DISTINCT COUNTRY FROM TB_TESLA_LOCATION WHERE 1=1 "

	if len(region) > 0 {
		sqlStatement += " AND REGION =  '" + region + "'"
	}

	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		//	return c.JSON(http.StatusCreated, u)
	}

	defer rows.Close()
	defer db.Close()

	result := models.Countrys{Result_Code: "S", Result_Message: "Success"}

	for rows.Next() {
		country := models.Country{}
		err2 := rows.Scan(&country.Country)
		if err2 != nil {
			return err2
		}
		result.Countrys = append(result.Countrys, country)
	}
	return c.JSON(http.StatusCreated, result)
}

// getLocationCity godoc
// @Summary Tesla Chargers Location City.
// @Description get tesla Chargers Location City.
// @Tags API Tesla App
// @ID getLocationCity
// @Accept */*
// @Produce json
// @Param country query string false "South Korea"
// @Success 200 {object} models.Citys
// @Router /getLocationCity [get]
func GetLocationCity(c echo.Context) error {

	db := db.DbManager()

	country := c.QueryParam("country")

	sqlStatement := "SELECT DISTINCT CITY FROM TB_TESLA_LOCATION WHERE 1=1 "

	if len(country) > 0 {
		sqlStatement += " AND COUNTRY =  '" + country + "'"
	}

	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		//	return c.JSON(http.StatusCreated, u)
	}

	defer rows.Close()
	defer db.Close()

	result := models.Citys{Result_Code: "S", Result_Message: "Success"}

	for rows.Next() {
		City := models.City{}
		err2 := rows.Scan(&City.City)
		if err2 != nil {
			return err2
		}
		result.Citys = append(result.Citys, City)
	}
	return c.JSON(http.StatusCreated, result)
}

// getLocationType godoc
// @Summary Tesla Chargers Location Type.
// @Description get tesla Chargers Location Type.
// @Tags API Tesla App
// @ID getLocationType
// @Accept */*
// @Produce json
// @Param country query string false "South Korea"
// @Param city query string false "서울특별시"
// @Success 200 {object} models.Location_types
// @Router /getLocationType [get]
func GetLocationType(c echo.Context) error {

	db := db.DbManager()

	country := c.QueryParam("country")

	city := c.QueryParam("city")

	sqlStatement := "SELECT DISTINCT (unnest(string_to_array(LOCATION_TYPE, '|'))) AS LOCATION_TYPE FROM TB_TESLA_LOCATION where  1=1 "

	if len(country) > 0 {
		sqlStatement += " AND COUNTRY =  '" + country + "'"
	}

	if len(city) > 0 {
		sqlStatement += " AND CITY =  '" + city + "'"
	}

	sqlStatement += " ORDER BY 1 DESC"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		//	return c.JSON(http.StatusCreated, u)
	}

	defer rows.Close()
	defer db.Close()

	result := models.Location_types{Result_Code: "S", Result_Message: "Success"}

	for rows.Next() {
		Location_type := models.Location_type{}
		err2 := rows.Scan(&Location_type.Location_type)
		if err2 != nil {
			return err2
		}
		result.Location_types = append(result.Location_types, Location_type)
	}
	return c.JSON(http.StatusCreated, result)
}

// getLocation godoc
// @Summary Tesla Chargers Location.
// @Description get tesla Chargers Location.
// @Tags API Tesla App
// @ID getLocation
// @Accept */*
// @Produce json
// @Param region query string false "asia_pacific"
// @Param country query string true "South Korea"
// @Param city query string false "서울특별시"
// @Param location_type query string false "supercharger"
// @Param nid query string false "28865"
// @Param title query string false "여의도"
// @Success 200 {object} models.Locations
// @Router /getLocation [get]
func GetLocation(c echo.Context) error {

	db := db.DbManager()

	sqlStatement := `SELECT 
			  sl_translate, address_line_1, address_line_2, address_notes, address, amenities, baidu_lat, baidu_lng, chargers, city
			, common_name, country, destination_charger_logo, destination_website, directions_link, emails, geocode, hours, is_gallery
			, kiosk_pin_x, kiosk_pin_y, kiosk_zoom_pin_x, kiosk_zoom_pin_y, latitude, location_id, location_type, longitude, nid, open_soon
			, path, postal_code, province_state, region, sales_phone, sales_representative, sub_region, title, trt_id, created_dttm FROM tb_tesla_location WHERE 1=1 `

	region := c.QueryParam("region")

	country := c.QueryParam("country")

	city := c.QueryParam("city")

	location_type := c.QueryParam("location_type")

	nid := c.QueryParam("nid")

	title := c.QueryParam("title")

	if len(region) > 0 {
		sqlStatement += " AND REGION =  '" + region + "'"
	}

	if len(country) > 0 {
		sqlStatement += " AND COUNTRY =  '" + country + "'"
	}

	if len(city) > 0 {
		sqlStatement += " AND CITY =  '" + city + "'"
	}

	if len(location_type) > 0 {
		sqlStatement += " AND LOCATION_TYPE LIKE  '%" + location_type + "%'"
	}

	if len(nid) > 0 {
		sqlStatement += " AND NID =  '" + nid + "'"
	}

	if len(title) > 0 {
		sqlStatement += " AND TITLE LIKE '%" + title + "%'"
	}
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		//	return c.JSON(http.StatusCreated, u)
	}

	defer rows.Close()
	defer db.Close()

	result := models.Locations{Result_Code: "S", Result_Message: "Success"}

	for rows.Next() {
		location := models.Location{}
		err2 := rows.Scan(&location.Sl_translate, &location.Address_line_1, &location.Address_line_2, &location.Address_notes, &location.Address, &location.Amenities, &location.Baidu_lat, &location.Baidu_lng, &location.Chargers, &location.City, &location.Common_name, &location.Country, &location.Destination_charger_logo, &location.Destination_website, &location.Directions_link, &location.Emails, &location.Geocode, &location.Hours, &location.Is_gallery, &location.Kiosk_pin_x, &location.Kiosk_pin_y, &location.Kiosk_zoom_pin_x, &location.Kiosk_zoom_pin_y, &location.Latitude, &location.Location_id, &location.Location_type, &location.Longitude, &location.Nid, &location.Open_soon, &location.Path, &location.Postal_code, &location.Province_state, &location.Region, &location.Sales_phone, &location.Sales_representative, &location.Sub_region, &location.Title, &location.Trt_id, &location.Created_dttm)
		if err2 != nil {
			return err2
		}
		result.Locations = append(result.Locations, location)
	}
	return c.JSON(http.StatusCreated, result)
}
