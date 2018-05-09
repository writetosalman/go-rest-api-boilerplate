CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `google2fa_enabled` tinyint(1) NOT NULL DEFAULT '0',
  `google2fa_object` blob DEFAULT NULL,

  `email_verified` tinyint(1) NOT NULL DEFAULT '0',
  `email_verification_code` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `last_login` timestamp NULL DEFAULT NULL,
  `reset_password_code` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,


  `first_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT '',
  `last_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT '',
  `address` varchar(255) COLLATE utf8_unicode_ci DEFAULT '',
  `city` varchar(100) COLLATE utf8_unicode_ci DEFAULT '',
  `state` varchar(100) COLLATE utf8_unicode_ci DEFAULT '',
  `country` varchar(100) COLLATE utf8_unicode_ci DEFAULT '',
  `post_code` varchar(20) COLLATE utf8_unicode_ci DEFAULT '',
  `phone` varchar(100) COLLATE utf8_unicode_ci DEFAULT '',

  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,

  PRIMARY KEY (`id`),
  UNIQUE KEY `users_email_unique` (`email`),
  KEY `users_email_verified_index` (`email_verified`),
  KEY `users_email_verification_code_index` (`email_verification_code`),
  KEY `users_email_index` (`email`)

) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci