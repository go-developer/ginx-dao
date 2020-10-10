--------------------------------------
-- 支持的协议列表: http/https/ws/wss等
-------------------------------------
CREATE TABLE `ginx_scheme` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `scheme` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '请求协议',
  `status` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态 0 - 待启用 1 - 可用中 2 - 禁用',
  `create_user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人ID',
  `modify_user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP() COMMENT '创建时间',
  `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP() COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 100000 DEFAULT CHARSET = utf8mb4 COMMENT = 'scheme信息表';
-----------------------------
--- 请求方法信息表 GET/POST 等
-----------------------------
CREATE TABLE `ginx_method`(
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMME方法 '主键ID',
  `method` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '请求协议',
  `status` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态 0 - 待启用 1 - 可用中 2 - 禁用',
  `create_user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人ID',
  `modify_user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP() COMMENT '创建时间',
  `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP() COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 100000 DEFAULT CHARSET = utf8mb4 COMMENT = '请求方法信息表';
------------
-- 用户信息表
------------
CREATE TABLE `ginx_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `mail` varchar(128) NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` varchar(16) NOT NULL DEFAULT '' COMMENT '手机',
  `status` tinyint(4) unsigned NOT NULL DEFAULT 0 COMMENT '状态 0 - 新加 1 - 正常 2 - 禁用 3 - 删除',
  `role` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '角色, 0 - 超级管理员',
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT '姓名',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(32) NOT NULL DEFAULT '' COMMENT '计算密码的盐值',
  `create_user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人ID',
  `modify_user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP() COMMENT '创建时间',
  `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP() COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unqi_mail` (`mail`),
  UNIQUE KEY `unqi_phone` (`phone`)
) ENGINE = InnoDB AUTO_INCREMENT = 100000 DEFAULT CHARSET = utf8mb4 COMMENT = '用户信息表';
------------
-- 项目信息表
------------
CREATE TABLE `ginx_project` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `flag` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '项目标识,全局唯一',
  `name` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '项目名称',
  `description` VARCHAR(4096) NOT NULL DEFAULT '' COMMENT '项目详细描述',
  `status` INT(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '项目状态: 0 - 待激活 1 - 正常 2 - 停用',
  `domain` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '项目域名',
  `port` INT(10) NOT NULL DEFAULT 80 COMMENT '服务端口',
  `scheme_id` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '协议ID',
  `create_user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人ID',
  `modify_user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP() COMMENT '创建时间',
  `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP() COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unqi_flag` (`flag`),
  UNIQUE KEY `unqi_domain` (`domain`),
  index `idx_create_id` (`create_user_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 100000 DEFAULT CHARSET = utf8mb4 COMMENT = '项目信息表';
-------------------
--- API 信息表
-------------------
CREATE TABLE `ginx_project_api` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `project_id` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属项目ID',
  `name` VARCHAR(128) NOT NULL DEFAULT '' COMMENT 'api名称',
  `status` INT(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '项目状态: 0 - 待激活 1 - 正常 2 - 停用',
  `uri` VARCHAR(512) NOT NULL DEFAULT '' COMMENT '请求的uri(网关对外暴露的uri)',
  `map_uri` VARCHAR(512) NOT NULL DEFAULT '' COMMENT '请求的uri(真实服务的uri,不配置默认与api相同)',
  `scheme_id` BIGINT(20) NOT NULL DEFAULT '' COMMENT '请求协议',
  `method_id` BIGINT(20) NOT NULL DEFAULT 0 COMMENT '请求方法',
  `description` VARCHAR(4096) NOT NULL DEFAULT '' COMMENT '功能的详细描述',
  `create_user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人ID',
  `modify_user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP() COMMENT '创建时间',
  `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP() COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unqi_project_api` (`project_id`, `api`)
) ENGINE = InnoDB AUTO_INCREMENT = 100000 DEFAULT CHARSET = utf8mb4 COMMENT = 'API信息表';
----------------------
-- api参数规则表
----------------------
CREATE TABLE `ginx_project_api_param` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `project_id` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属项目ID',
  `api_id` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属的API的ID',
  `param_name` VARCHAR(63) NOT NULL DEFAULT 0 COMMENT '参数名称',
  `status` INT(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '项目状态: 0 - 待激活 1 - 正常 2 - 停用',
  `param_type` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '参数类型: 0 - string 1 - int64 2 - float64 3 - list(指定分隔符) 4 - list(json序列化后的list) 5 - json(map[string]) 6 - 其他',
  `is_required` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否必传',
  `example_value` VARCHAR(2048) NOT NULL DEFAULT '' COMMENT '示例的值',
  `create_user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人ID',
  `modify_user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP() COMMENT '创建时间',
  `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP() COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unqi_api_param` (`api_id`, `param_name`)
) ENGINE = InnoDB AUTO_INCREMENT = 100000 DEFAULT CHARSET = utf8mb4 COMMENT = 'API参数信息表';
