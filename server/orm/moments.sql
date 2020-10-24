--aid
--send_time # 二级索引，时间戳，用来组建朋友圈timeline
--content # 文章内容
--photo_list # 图片url列表
--privacy # 3天可见等

create table if not exisit Aticle(
    aid int AUTO_INCREMENT  primary key ,
    send_time int not null,
    content text not null,
    photo_list text default null,
    privacy int default 0,
    is_deleted int default 0
)ENGINE=InnoDB DEFAULT CHARSET=utf8;