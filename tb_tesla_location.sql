-- public.tb_tesla_location definition

-- Drop table

-- DROP TABLE public.tb_tesla_location;

CREATE TABLE public.tb_tesla_location (
        sl_translate varchar(500) NULL,
        address_line_1 varchar(500) NULL,
        address_line_2 varchar(500) NULL,
        address_notes varchar(500) NULL,
        address varchar(500) NULL,
        amenities varchar(1000) NULL,
        baidu_lat varchar(500) NULL,
        baidu_lng varchar(500) NULL,
        chargers varchar(500) NULL,
        city varchar(500) NULL,
        common_name varchar(500) NULL,
        country varchar(500) NULL,
        destination_charger_logo varchar(1000) NULL,
        destination_website varchar(500) NULL,
        directions_link varchar(500) NULL,
        emails varchar(500) NULL,
        geocode varchar(500) NULL,
        hours varchar(1000) NULL,
        is_gallery varchar(500) NULL,
        kiosk_pin_x varchar(500) NULL,
        kiosk_pin_y varchar(500) NULL,
        kiosk_zoom_pin_x varchar(500) NULL,
        kiosk_zoom_pin_y varchar(500) NULL,
        latitude varchar(500) NULL,
        location_id varchar(500) NULL,
        location_type varchar(500) NULL,
        longitude varchar(500) NULL,
        nid varchar(500) NOT NULL,
        open_soon varchar(500) NULL,
        "path" varchar(500) NULL,
        postal_code varchar(500) NULL,
        province_state varchar(500) NULL,
        region varchar(500) NULL,
        sales_phone varchar(500) NULL,
        sales_representative varchar(500) NULL,
        sub_region varchar(500) NULL,
        title varchar(500) NULL,
        trt_id varchar(500) NULL,
        created_dttm timestamp NULL,
        CONSTRAINT tb_tesla_location_pkey PRIMARY KEY (nid)
);
	
