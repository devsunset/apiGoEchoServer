-- public.tb_tesla_location definition

-- Drop table

-- DROP TABLE public.tb_tesla_location;

CREATE TABLE public.tb_tesla_location (
        sl_translate varchar(2048) NULL,
        address_line_1 varchar(2048) NULL,
        address_line_2 varchar(2048) NULL,
        address_notes varchar(2048) NULL,
        address varchar(2048) NULL,
        amenities varchar(2048) NULL,
        baidu_lat varchar(2048) NULL,
        baidu_lng varchar(2048) NULL,
        chargers varchar(2048) NULL,
        city varchar(2048) NULL,
        common_name varchar(2048) NULL,
        country varchar(2048) NULL,
        destination_charger_logo varchar(2048) NULL,
        destination_website varchar(2048) NULL,
        directions_link varchar(2048) NULL,
        emails varchar(2048) NULL,
        geocode varchar(2048) NULL,
        hours varchar(2048) NULL,
        is_gallery varchar(2048) NULL,
        kiosk_pin_x varchar(2048) NULL,
        kiosk_pin_y varchar(2048) NULL,
        kiosk_zoom_pin_x varchar(2048) NULL,
        kiosk_zoom_pin_y varchar(2048) NULL,
        latitude varchar(2048) NULL,
        location_id varchar(2048) NULL,
        location_type varchar(2048) NULL,
        longitude varchar(2048) NULL,
        nid varchar(2048) NOT NULL,
        open_soon varchar(2048) NULL,
        "path" varchar(2048) NULL,
        postal_code varchar(2048) NULL,
        province_state varchar(2048) NULL,
        region varchar(2048) NULL,
        sales_phone varchar(2048) NULL,
        sales_representative varchar(2048) NULL,
        sub_region varchar(2048) NULL,
        title varchar(2048) NULL,
        trt_id varchar(2048) NULL,
        created_dttm timestamp NULL,
        CONSTRAINT tb_tesla_location_pkey PRIMARY KEY (nid)
);
	
