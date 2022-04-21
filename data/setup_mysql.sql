
DROP TABLE IF EXISTS `go_repository`.`tenant`;
DROP TABLE IF EXISTS `go_repository`.`property`;
DROP TABLE IF EXISTS `go_repository`.`setup`;

DROP TABLE IF EXISTS `go_repository`.`post`;
DROP TABLE IF EXISTS `go_repository`.`thread`;
DROP TABLE IF EXISTS `go_repository`.`session`;
DROP TABLE IF EXISTS `go_repository`.`user`;
DROP TABLE IF EXISTS `go_repository`.`sys_adminstrator`;

CREATE TABLE `go_repository`.`sys_adminstrator` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(50) NOT NULL unique DEFAULT '' COMMENT '管理员名',
  `password` varchar(256) NOT NULL DEFAULT '' COMMENT '密码',
  `email` varchar(50) DEFAULT NULL unique COMMENT '邮箱',
  `phone` varchar(11) DEFAULT NULL COMMENT '座机号',
  `mobile_phone` varchar(11) DEFAULT NULL unique COMMENT 'mobile_phone手机号',
  `role` varchar(50) NOT NULL DEFAULT '' COMMENT '角色名',
  `status` int(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created_at`     int(11)  null COMMENT 'created_at',
  `updated_at`     int(11)  null COMMENT 'updated_at',
  
  PRIMARY KEY (`id`),
   UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE)
   ENGINE = InnoDB
  AUTO_INCREMENT=0
  DEFAULT CHARACTER SET = utf8mb4 COMMENT='管理员表';

INSERT INTO `go_repository`.`sys_adminstrator`
(`id`,`user_name`,`email`,`role`,`password`)
VALUES
(0,'root','13808013567@163.com','systemAdmin',
'7c4a8d09ca3762af61e59520943dc26494f8941b');


-- 应用承租户表
CREATE TABLE `go_repository`.`tenant` (
  `id`               int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id`        VARCHAR(64) NOT NULL unique COMMENT 'tenant_id', -- uuid
  `tenant_name`      varchar(255) not null unique COMMENT 'tenant_name',
  `tenant_des`       varchar(255) COMMENT 'tenant_des',
  `disabled_is`      tinyint(1) not null DEFAULT 0 COMMENT 'disabled_is', -- 1:禁用当前记录
  `cctasp`           varchar(255) not null DEFAULT ''  COMMENT 'cctasp', -- 云计算技术和服务提供商
  `start_at`         varchar(255)  COMMENT 'start_at', -- 承租起始日期
  `status`           varchar(10) not null DEFAULT 'Primary'  COMMENT 'status', -- 状态
  `due_at`           varchar(255)  COMMENT 'due_at',   -- 承租到期日期
  `created_at`       int(11) not null COMMENT 'created_at',
  `updated_at`       int(11) not null COMMENT 'updated_at',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `tenant_id_UNIQUE` (`tenant_id` ASC) VISIBLE)
  ENGINE = InnoDB
  AUTO_INCREMENT=0
  DEFAULT CHARACTER SET = utf8mb4 COMMENT='tenant';

-- property 承租所有物    
CREATE TABLE `go_repository`.`property` (
  `id`               VARCHAR(64) NOT NULL unique COMMENT 'id', -- uuid
  `parent_id`        VARCHAR(64) NULL  COMMENT 'parent_id', -- uuid
  `type`             VARCHAR(64) NOT NULL DEFAULT 'Property' COMMENT 'type', -- Property(房屋及院落\庄园_分公司办事处)
  `region`           varchar(10) not null  DEFAULT 'cn'  COMMENT 'region',  -- 国别的域名缩写 cn
  `name`             varchar(255) not null unique COMMENT 'name',
  `code`             varchar(10) not null unique COMMENT 'code', -- name的缩写大写英文字母无空格 2到10个字符
  `des`              varchar(255) null COMMENT 'des', -- property详细描述
  `sequence`         int(11)  not null  DEFAULT 1 COMMENT 'sequence', -- name排序
  `published_is`     tinyint(1) not null DEFAULT 1 COMMENT 'disable_identity', -- 1:代表发布
  `disabled_is`      tinyint(1) not null DEFAULT 0 COMMENT 'disable_identity', -- 1:禁用当前记录
  `created_at`       int(11) not null COMMENT 'created_at',
  `updated_at`       int(11) not null COMMENT 'updated_at',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE)
  ENGINE = InnoDB
  AUTO_INCREMENT=0
  DEFAULT CHARACTER SET = utf8mb4 COMMENT='property';


-- 系统自定义设置表
CREATE TABLE `go_repository`.`setup` (
  `id`               VARCHAR(64) NOT NULL COMMENT 'id', -- uuid
  `parent_id`        VARCHAR(64) NOT NULL DEFAULT '' COMMENT 'parent_id', -- uuid
  `property_id`      VARCHAR(64) NOT NULL  DEFAULT ''  COMMENT 'property_id',  -- property_id property中type=='Property'的id
  `setup_genre`      VARCHAR(64) NOT NULL  COMMENT 'setup_genre', -- setup的 体裁 类别
  `code`             varchar(16) not null unique COMMENT 'code',  -- 大写英文字母无空格 2到16个字符  
  `key`              varchar(64) not null COMMENT 'key', -- 大小写英文字母无空格 code的别名
  `value`            varchar(255) not null COMMENT 'value', -- 存key的值
  `value_type`       VARCHAR(64) NOT NULL  COMMENT 'value_type', -- value的数据类型:string,int,bool
  `titel`            VARCHAR(255) NOT NULL  COMMENT 'titel', -- value的输入在页显示的标题
  `label`            VARCHAR(255) NOT NULL  COMMENT 'label', -- value的输入在页显示的输入框label
  `des`              VARCHAR(255) NOT NULL DEFAULT 'null'  COMMENT 'label', -- value的输入在页显示的输入框label
  `sequence`         int(11) not null DEFAULT '1'  COMMENT  'sequence', -- value排序
  `extension`        VARCHAR(255)  NULL COMMENT 'extension', -- 用于扩展备用
  `disabled_is`      tinyint(1) not null DEFAULT 0 COMMENT 'disable_identity', -- 1:禁用当前记录
  `created_at`       int(11) not null COMMENT 'created_at',
  `updated_at`       int(11) not null COMMENT 'updated_at',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE)
  ENGINE = InnoDB
  AUTO_INCREMENT=0
  DEFAULT CHARACTER SET = utf8mb4 COMMENT='setup';




-- 用户表
CREATE TABLE `go_repository`.`user` (
  `id`              int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id`       VARCHAR(64) NOT NULL  COMMENT 'tenant_id',
  `uuid`            VARCHAR(64) NOT NULL unique COMMENT 'uuid',
  `user_name`        varchar(255) COMMENT 'user_name',
  `first_name` 		  varchar(255) COMMENT 'first_name',
	`last_name`        varchar(255) COMMENT 'last_name',
  `gender`           varchar(16) COMMENT 'gender',
  `email`           varchar(255) not null unique COMMENT 'email',
  `mobile_phone`   varchar(255) not null unique COMMENT 'mobile_phone',
  `password`        varchar(255) not null COMMENT 'password', 
  `disabled_is`     tinyint(1) not null DEFAULT 0 COMMENT 'disable_identity', -- 1:禁用当前记录 
  `created_at`      int(11) not null COMMENT 'created_at',
  `updated_at`      int(11) not null COMMENT 'updated_at',
  --  `created_at`  timestamp not null DEFAULT CURRENT_TIMESTAMP COMMENT 'created_at',
  -- `updated_at`    timestamp not null DEFAULT CURRENT_TIMESTAMP COMMENT 'updated_at',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uuid_UNIQUE` (`uuid` ASC) VISIBLE)
  ENGINE = InnoDB
  AUTO_INCREMENT=0
  DEFAULT CHARACTER SET = utf8mb4 COMMENT='user';

-- session 
CREATE TABLE `go_repository`.`session` (
  `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uuid`        VARCHAR(64) NOT NULL,
  `email`      varchar(255),
  `user_id`    integer references users(id),
  `created_at` timestamp not null,   
PRIMARY KEY (`id`),
  UNIQUE INDEX `uuid_UNIQUE` (`uuid` ASC) VISIBLE)
  ENGINE = InnoDB
  AUTO_INCREMENT=0
  DEFAULT CHARACTER SET = utf8mb4 COMMENT='session';

CREATE TABLE `go_repository`.`thread` (
  `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uuid`        VARCHAR(64) NOT NULL,
  `topic`      text,
  `user_id`    integer references users(id),
  `created_at` timestamp not null ,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uuid_UNIQUE` (`uuid` ASC) VISIBLE)
  ENGINE = InnoDB
  AUTO_INCREMENT=0
  DEFAULT CHARACTER SET = utf8mb4 COMMENT='thread';

CREATE TABLE `go_repository`.`post` (
  `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uuid`        VARCHAR(64) NOT NULL,
  `body`       text,
  `user_id`    integer references users(id),
  `thread_id`  integer references threads(id),
  `created_at` timestamp not null ,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uuid_UNIQUE` (`uuid` ASC) VISIBLE)
  ENGINE = InnoDB
  AUTO_INCREMENT=0
  DEFAULT CHARACTER SET = utf8mb4 COMMENT='post';