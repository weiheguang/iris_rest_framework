/* 
 用户表 
 一个用户属于一个组织, 不按照手机号进行判断用户, 按照org_id和phong一起来判断
 */
CREATE TABLE `auth_user` (
	`id` varchar(40) COMMENT "用户ID",
	`username` varchar(30) NOT NULL COMMENT "用户名",
	`password` varchar(128) NULL COMMENT "用户密码"",
	`is_superuser` tinyint(1) DEFAULT 0 COMMENT " 是否超级用户 ",
	`phone` varchar(11) NULL COMMENT " 用户手机号 ",
	`is_active` tinyint(1) DEFAULT 1 COMMENT " 用户状态 ",
	`is_del` tinyint(1) DEFAULT 1 COMMENT '软删除',
	`created_at` datetime(6) NULL COMMENT '创建时间',
	PRIMARY KEY (`id`),
	Unique KEY `auth_user_username_jmjhg`(`username`) USING BTREE,
	Unique KEY `auth_user_phone_dljfe`(`phone`) USING BTREE
);

insert into
	`auth_user`(
		`id`,
		`password`,
		`is_superuser`,
		`username`,
		`phone`,
		`is_active`,
		`is_del`,
		`created_at`
	)
values
	(
		" 459894584599458934984958 ",
		'bjc_md5$10000$37821e94625d763c2ff217d089523262',
		1,
		'admin',
		'13888888888',
		1,
		0,
		'2022-09-27'
	);