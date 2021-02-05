CREATE TABLE zip (
    id BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'ID', 
    local_government_code BIGINT(20) NOT NULL COMMENT '全国地方公共団体コード',
    prefecture_code BIGINT(20) NOT NULL COMMENT '都道府県コード',
    zip_code BIGINT(20) NOT NULL COMMENT '郵便コード',
    prefecture_kana VARCHAR(255) NOT NULL COMMENT '都道府県カナ',
    city_kana VARCHAR(255) NOT NULL COMMENT '市区町村カナ',
    town_kana VARCHAR(255) COMMENT '町域カナ',
    prefecture VARCHAR(255) NOT NULL COMMENT '都道府県',
    city VARCHAR(255) NOT NULL COMMENT '市区町村',
    town VARCHAR(255) NOT NULL COMMENT '町域',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4