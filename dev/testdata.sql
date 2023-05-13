/* 插入初始化数据 */
/*
 新建用户
 */
INSERT INTO `users` (
        `id`,
        `created_at`,
        `updated_at`,
        `deleted_at`,
        `username`,
        `password`,
        `email`,
        `phone`,
        `avatar`,
        `role`,
        `is_super_admin`,
        `deleted`,
        `age`,
        `gender`
    )
VALUES (
        1,
        '2023-05-11 07:19:11.683',
        '2023-05-11 07:19:11.683',
        NULL,
        'werido',
        'e10adc3949ba59abbe56e057f20f883e',
        '359066432@qq.com',
        '',
        '',
        0,
        0,
        0,
        0,
        0
    )
ON DUPLICATE KEY UPDATE 

    ;
/*
 插入内置的bilibili运行计划
 */
INSERT INTO `simple_notification`.`scheduler` (
        `id`,
        `created_at`,
        `updated_at`,
        `deleted_at`,
        `period`,
        `user_id`,
        `type`,
        `active`,
        `platform`,
        `name`,
        `description`
    )
VALUES (
        1,
        '2023-05-11 07:21:59.132',
        '2023-05-11 07:21:59.132',
        NULL,
        '*/1 * * * * ',
        1,
        'custom',
        1,
        'bilibili',
        'bilibiliUp内容发布新同志获取',
        'test scheduler'
    );
/*
 插入一个 bilibili up 作品订阅通知 task
 
 */
INSERT INTO `task` (
        `id`,
        `created_at`,
        `updated_at`,
        `deleted_at`,
        `platform`,
        `ups`,
        `user_id`,
        `active`,
        `name`,
        `description`
    )
VALUES (
        1,
        '2023-05-11 07:22:14.428',
        '2023-05-11 07:22:14.428',
        NULL,
        'bilibili',
        '{\"敬汉卿\": 9824766, \"盗月社食遇记\": 99157282}',
        1,
        1,
        'bibibiliTaskTest',
        ''
    );
/*
 绑定上述插入的task到scheduler
 */
INSERT INTO `scheduler_task` (
        `id`,
        `created_at`,
        `updated_at`,
        `deleted_at`,
        `scheduler_id`,
        `task_id`
    )
VALUES (
        1,
        '2023-05-11 07:24:30.163',
        '2023-05-11 07:24:30.163',
        NULL,
        1,
        1
    );
/*
 创建一个email通知
 */
INSERT INTO `email_notifier` (
        `id`,
        `created_at`,
        `updated_at`,
        `deleted_at`,
        `user_id`,
        `sender`,
        `pwd`,
        `receiver`,
        `content`
    )
VALUES (
        1,
        '2023-05-11 08:54:59.575',
        '2023-05-11 08:54:59.575',
        NULL,
        1,
        '359066432@qq.com',
        'wjcyfwozmmerbghi',
        '[\"weridolin@qq.com\"]',
        '测试内容'
    );

    
/*
 绑定一个email notify 到task
 */
INSERT INTO `email_notifier_tasks` (
        `id`,
        `created_at`,
        `updated_at`,
        `deleted_at`,
        `email_notifier_id`,
        `task_id`
    )
VALUES (
        1,
        '2023-05-11 08:55:10.494',
        '2023-05-11 08:55:10.494',
        NULL,
        1,
        1
    );