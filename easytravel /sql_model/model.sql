CREATE TABLE "country"(
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "title" VARCHAR NOT NULL,
    "code" VARCHAR NOT NULL,
    "continent" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "city" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "title" VARCHAR NOT NULL,
    "country_id" UUID REFERENCES "country"("id"),
    "city_code" VARCHAR NOT NULL,
    "latitude" VARCHAR NOT NULL,
    "longitude" VARCHAR NOT NULL,
    "offset" VARCHAR NOT NULL,
    "country_name" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "aeroport" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "title" VARCHAR NOT NULL,
    "country_id" UUID REFERENCES "country"("id"),
    "city_id" UUID REFERENCES "city"("id"),
    "latitude" VARCHAR NOT NULL,
    "longitude" VARCHAR NOT NULL,
    "radius" VARCHAR DEFAULT 'nul',
    "image" VARCHAR DEFAULT '',
    "address" VARCHAR DEFAULT '',
    "country" VARCHAR DEFAULT '',
    "city" VARCHAR DEFAULT '',
    "search_text" VARCHAR DEFAULT '',
    "code" VARCHAR DEFAULT '',
    "product_count" NUMERIC DEFAULT 0,
    "gmt" VARCHAR DEFAULT '00:00',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);




-- jq -c '.[]' country.json >> your_data.json
-- cat country.json | jq -cr '.[]' | sed 's/\\[tn]//g' > countryjson.json
-- cat aeroport.json | jq -cr '.[]' | sed 's/\\[tn]//g' > aeroportjson.json


-- \copy temp(data) from './mock/countryjson.json';
-- \copy temp(data) from './mock_data/your_data.json';
-- \copy aeroportjson(data) from 'aeroportjson.json';

-- DELETE FROM temp WHERE length(data ->> 'country_id') = 2;

-- INSERT INTO country (
--     "title",
--     "code",
--     "continent"
-- )
-- SELECT 
--     JSON_BUILD_OBJECT(
--         countrydata ->>'title'
--         countrydata ->>'code'
--         countrydata ->>'continent'
--     )
-- FROM 
--     countryjson;


-- INSERT INTO city(
--     "title",
--     "country_id",
--     "city_code",
--     "latitude",
--     "longitude",
--     "offset",
--     "country_name",
--     "updated_at"
-- ) 
-- SELECT 
--     JSON_BUILD_OBJECT(
--         data -> 'title',
--         data -> 'country_id',
--         data -> 'city_code',
--         data -> 'latitude',
--         data -> 'longitude',
--         data -> 'offset',
--         data -> 'country_name',
--         NOW()
--     )
-- FROM 
--     temp;
