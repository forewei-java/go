CREATE DATABASE
IF
	NOT EXISTS blog_service DEFAULT CHARACTER 
	SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS blog_service.`blog_tag` (
    `id` INT ( 10 ) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR ( 100 ) DEFAULT '' COMMENT '标签名称',
    `created_on` INT ( 10 ) UNSIGNED DEFAULT '0' COMMENT '创建时间',
    `created_by` VARCHAR ( 100 ) DEFAULT '' COMMENT '创建人',
    `modified_on` INT ( 10 ) UNSIGNED DEFAULT '0' COMMENT '修改时间',
    `modified_by` VARCHAR ( 100 ) DEFAULT '' COMMENT '修改人',
    `deleted_on` INT ( 10 ) UNSIGNED DEFAULT '0' COMMENT '删除时间',
    `is_del` TINYINT ( 3 ) UNSIGNED DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
    `state` TINYINT ( 3 ) UNSIGNED DEFAULT '1' COMMENT '状态 0 为禁用、1 为启用',
    PRIMARY KEY ( `id` )
    ) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '标签管理';

CREATE TABLE IF NOT EXISTS blog_service.`blog_article` (
    `id` INT ( 10 ) UNSIGNED NOT NULL AUTO_INCREMENT,
    `title` VARCHAR ( 100 ) DEFAULT '' COMMENT '文章标题',
    `desc` VARCHAR ( 255 ) DEFAULT '' COMMENT '文章简述',
    `cover_image_url` VARCHAR ( 255 ) DEFAULT '' COMMENT '封面图片地址',
    `content` LONGTEXT COMMENT '文章内容',
    `created_on` INT ( 10 ) UNSIGNED DEFAULT '0' COMMENT '创建时间',
    `created_by` VARCHAR ( 100 ) DEFAULT '' COMMENT '创建人',
    `modified_on` INT ( 10 ) UNSIGNED DEFAULT '0' COMMENT '修改时间',
    `modified_by` VARCHAR ( 100 ) DEFAULT '' COMMENT '修改人',
    `deleted_on` INT ( 10 ) UNSIGNED DEFAULT '0' COMMENT '删除时间',
    `is_del` TINYINT ( 3 ) UNSIGNED DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
    `state` TINYINT ( 3 ) UNSIGNED DEFAULT '1' COMMENT '状态 0 为禁用、1 为启用',
    PRIMARY KEY ( `id` )
    ) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '文章管理';


CREATE TABLE IF NOT EXISTS blog_service.`blog_article_tag` (
    `id` INT ( 10 ) UNSIGNED NOT NULL AUTO_INCREMENT,
    `article_id` INT ( 11 ) NOT NULL COMMENT '文章 ID',
    `tag_id` INT ( 10 ) UNSIGNED NOT NULL DEFAULT '0' COMMENT '标签 ID',
    `created_on` INT ( 10 ) UNSIGNED DEFAULT '0' COMMENT '创建时间',
    `created_by` VARCHAR ( 100 ) DEFAULT '' COMMENT '创建人',
    `modified_on` INT ( 10 ) UNSIGNED DEFAULT '0' COMMENT '修改时间',
    `modified_by` VARCHAR ( 100 ) DEFAULT '' COMMENT '修改人',
    `deleted_on` INT ( 10 ) UNSIGNED DEFAULT '0' COMMENT '删除时间',
    `is_del` TINYINT ( 3 ) UNSIGNED DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
    PRIMARY KEY ( `id` )
    ) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '文章标签关联';