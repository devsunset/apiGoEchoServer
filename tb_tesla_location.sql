-- public.tb_tesla_location definition

-- Drop table

-- DROP TABLE public.tb_tesla_location;

CREATE TABLE public.tb_tesla_location (
	sl_translate varchar(10) NULL,
	address_line_1 varchar(500) NULL,
	address_line_2 varchar(500) NULL,
	address_notes varchar(500) NULL,
	address varchar(500) NULL,
	amenities varchar(1000) NULL,
	baidu_lat varchar(20) NULL,
	baidu_lng varchar(20) NULL,
	chargers varchar(500) NULL,
	city varchar(100) NULL,
	common_name varchar(100) NULL,
	country varchar(50) NULL,
	destination_charger_logo varchar(1000) NULL,
	destination_website varchar(255) NULL,
	directions_link varchar(500) NULL,
	emails varchar(500) NULL,
	geocode varchar(50) NULL,
	hours varchar(1000) NULL,
	is_gallery varchar(50) NULL,
	kiosk_pin_x varchar(10) NULL,
	kiosk_pin_y varchar(10) NULL,
	kiosk_zoom_pin_x varchar(50) NULL,
	kiosk_zoom_pin_y varchar(50) NULL,
	latitude varchar(20) NULL,
	location_id varchar(100) NULL,
	location_type varchar(100) NULL,
	longitude varchar(20) NULL,
	nid varchar(10) NOT NULL,
	open_soon varchar(10) NULL,
	"path" varchar(100) NULL,
	postal_code varchar(20) NULL,
	province_state varchar(50) NULL,
	region varchar(50) NULL,
	sales_phone varchar(500) NULL,
	sales_representative varchar(50) NULL,
	sub_region varchar(50) NULL,
	title varchar(255) NULL,
	trt_id varchar(10) NULL,
	created_dttm timestamp NULL,
	CONSTRAINT tb_tesla_location_pkey PRIMARY KEY (nid)
);