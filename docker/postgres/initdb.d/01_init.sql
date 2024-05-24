CREATE DATABASE ur_v3;

\c ur_v3;

CREATE SEQUENCE IF NOT EXISTS prefs_id_seq;

CREATE TABLE "public"."prefs" (
    "id" int8 NOT NULL DEFAULT nextval('prefs_id_seq'::regclass),
    "code" varchar(10) NOT NULL,
    "region" varchar(15) NOT NULL,
    "name" varchar(5) NOT NULL,
    "is_crawl" bool NOT NULL DEFAULT false,
    "created_at" timestamp(0),
    "updated_at" timestamp(0),
    PRIMARY KEY ("id"),
    UNIQUE ("code")
);


CREATE SEQUENCE IF NOT EXISTS houses_id_seq;

CREATE TABLE "public"."houses" (
    "id" int8 NOT NULL DEFAULT nextval('houses_id_seq'::regclass),
    "code" varchar(7) NOT NULL,
    "pref_code" varchar(10) NOT NULL,
    "name" varchar(30) NOT NULL,
    "rooms_got_at" timestamp(0) NOT NULL DEFAULT '1000-01-01 00:00:00'::timestamp without time zone,
    "created_at" timestamp(0),
    "updated_at" timestamp(0),
    PRIMARY KEY ("id"),
    UNIQUE ("code")
);


CREATE SEQUENCE IF NOT EXISTS rooms_id_seq;

CREATE TABLE "public"."rooms" (
    "id" int8 NOT NULL DEFAULT nextval('rooms_id_seq'::regclass),
    "house_code" varchar(7) NOT NULL,
    "room_code" varchar(15) NOT NULL,
    "status" varchar(255) NOT NULL DEFAULT 'ready'::character varying CHECK ((status)::text = ANY ((ARRAY['ready'::character varying, 'closed'::character varying])::text[])),
    "got_at" timestamp(0) NOT NULL DEFAULT '1000-01-01 00:00:00'::timestamp without time zone,
    "data" json,
    "price" int4 NOT NULL DEFAULT 0,
    "fee" int4 NOT NULL DEFAULT 0,
    "type" varchar(20),
    "space" int4 NOT NULL DEFAULT 0,
    "floor" int4 NOT NULL DEFAULT 0,
    "layout_url" varchar(150),
    "url" varchar(100),
    "created_at" timestamp(0),
    "updated_at" timestamp(0),
    PRIMARY KEY ("id"),
    UNIQUE ("house_code", "room_code")
);

INSERT INTO prefs (code, region, name, is_crawl, created_at, updated_at) VALUES
('hokkaido', 'hokkaitohoku', '北海道', false, NOW(), NOW()),
('miyagi', 'hokkaitohoku', '宮城県', false, NOW(), NOW()),
('tokyo', 'kanto', '東京都', true, NOW(), NOW()),
('kanagawa', 'kanto', '神奈川県', true, NOW(), NOW()),
('chiba', 'kanto', '千葉県', true, NOW(), NOW()),
('saitama', 'kanto', '埼玉県', true, NOW(), NOW()),
('ibaraki', 'kanto', '茨城県', false, NOW(), NOW()),
('aichi', 'tokai', '愛知県', false, NOW(), NOW()),
('mie', 'tokai', '三重県', false, NOW(), NOW()),
('gifu', 'tokai', '岐阜県', false, NOW(), NOW()),
('shizuoka', 'tokai', '静岡県', false, NOW(), NOW()),
('osaka', 'kansai', '大阪府', true, NOW(), NOW()),
('hyogo', 'kansai', '兵庫県', false, NOW(), NOW()),
('kyoto', 'kansai', '京都府', false, NOW(), NOW()),
('shiga', 'kansai', '滋賀県', false, NOW(), NOW()),
('nara', 'kansai', '奈良県', false, NOW(), NOW()),
('wakayama', 'kansai', '和歌山県', false, NOW(), NOW()),
('okayama', 'chugoku', '岡山県', false, NOW(), NOW()),
('hiroshima', 'chugoku', '広島県', false, NOW(), NOW()),
('yamaguchi', 'chugoku', '山口県', false, NOW(), NOW()),
('fukuoka', 'kyushu', '福岡県', true, NOW(), NOW());