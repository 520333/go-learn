DROP TABLE IF EXISTS `stree_path`;
CREATE TABLE `stree_path` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `level` tinyint(4) NOT NULL COMMENT '服务层级: 1=G(业务线) 2=P(项目) 3=A(应用)',
    `path` varchar(200) DEFAULT NULL COMMENT '物化路径 (例如: 0, /1/, /1/2/)',
    `node_name` varchar(200) DEFAULT NULL COMMENT '节点名称 (例如: inf, monitor, kafka)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_key` (`level`,`path`,`node_name`) USING BTREE COMMENT '联合唯一索引: 保证同层级同父节点下不重名'
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='服务树节点路径表';