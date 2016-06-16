drop database if exists test;
create database test;
use test;


CREATE TABLE IF NOT EXISTS `picture_table` (
 `pic_id` INT(4) UNSIGNED NOT NULL COMMENT 'picture_number',
 `pic_title` VARCHAR(128) NOT NULL COMMENT 'picture_title',
 `pic_url` VARCHAR(256) NOT NULL COMMENT 'image_url',
 PRIMARY KEY (`pic_id`))
 ENGINE = MyISAM
 DEFAULT CHARACTER SET = utf8
 COLLATE = utf8_unicode_ci
 COMMENT = 'pictures'
