# OneBlogService


## SQL

- CREATE DATABASE
```sql
CREATE DATABASE
IF NOT EXISTS blog_service DEFAULT CHARACTER 
SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;
```
- 公共字段
```sql
 `crated_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `crated_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMNET '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 未删 ， 1 已删',
```
- blog_tag
```sql
CREATE TABLE `blog_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) DEFAULT '' COMMENT '标签名',
    `crated_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `crated_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 未删 ， 1 已删', 
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0 禁用，1 启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';
```
- blog_article 
```sql
CREATE TABLE `blog_article` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(100) default '' COMMENT '文章标题',
    `desc` varchar(255) default '' COMMENT '文章简述',
    `cover_image_url` varchar(255) default '' COMMENT '封面图片地址',
    `content` longtext COMMENT '文章内容',
    `crated_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `crated_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 未删 ， 1 已删',
     `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0 禁用，1 启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理'; 
```

- blog_article_tag
```sql 
   CREATE TABLE `blog_article_tag` (
     `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
      `article_id` int(11)   NOT NULL COMMENT '文章ID',
       `tag_id` int(11)   NOT NULL COMMENT '文章ID',
        PRIMARY KEY (`id`)
   ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章关联标签'; 
```
