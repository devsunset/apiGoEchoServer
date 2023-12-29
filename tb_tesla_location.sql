-- public.tb_tesla_location definition

-- Drop table

-- DROP TABLE public.tb_tesla_location;

CREATE TABLE public.tb_tesla_location (
        sl_translate varchar(1024) NULL,
        address_line_1 varchar(1024) NULL,
        address_line_2 varchar(1024) NULL,
        address_notes varchar(1024) NULL,
        address varchar(1024) NULL,
        amenities varchar(1024) NULL,
        baidu_lat varchar(1024) NULL,
        baidu_lng varchar(1024) NULL,
        chargers varchar(1024) NULL,
        city varchar(1024) NULL,
        common_name varchar(1024) NULL,
        country varchar(1024) NULL,
        destination_charger_logo varchar(1024) NULL,
        destination_website varchar(1024) NULL,
        directions_link varchar(1024) NULL,
        emails varchar(1024) NULL,
        geocode varchar(1024) NULL,
        hours varchar(1024) NULL,
        is_gallery varchar(1024) NULL,
        kiosk_pin_x varchar(1024) NULL,
        kiosk_pin_y varchar(1024) NULL,
        kiosk_zoom_pin_x varchar(1024) NULL,
        kiosk_zoom_pin_y varchar(1024) NULL,
        latitude varchar(1024) NULL,
        location_id varchar(1024) NULL,
        location_type varchar(1024) NULL,
        longitude varchar(1024) NULL,
        nid varchar(1024) NOT NULL,
        open_soon varchar(1024) NULL,
        "path" varchar(1024) NULL,
        postal_code varchar(1024) NULL,
        province_state varchar(1024) NULL,
        region varchar(1024) NULL,
        sales_phone varchar(1024) NULL,
        sales_representative varchar(1024) NULL,
        sub_region varchar(1024) NULL,
        title varchar(1024) NULL,
        trt_id varchar(1024) NULL,
        created_dttm timestamp NULL,
        CONSTRAINT tb_tesla_location_pkey PRIMARY KEY (nid)
);
	
