-- 创建POI表结构
CREATE TABLE IF NOT EXISTS `poi_items` (
  `id` VARCHAR(20) NOT NULL COMMENT 'POI唯一标识',
  `name` VARCHAR(100) NOT NULL COMMENT 'POI名称',
  `type` VARCHAR(100) COMMENT '类型描述',
  `location` JSON COMMENT '位置坐标[经度,纬度]',
  `address` VARCHAR(255) COMMENT '详细地址',
  `tel` VARCHAR(20) COMMENT '联系电话',
  `distance` INT COMMENT '距离(米)',
  `shopinfo` VARCHAR(50) COMMENT '商铺信息',
  `website` VARCHAR(255) COMMENT '网站',
  `pcode` VARCHAR(10) COMMENT '省份编码',
  `citycode` VARCHAR(10) COMMENT '城市编码',
  `adcode` VARCHAR(10) COMMENT '区域编码',
  `postcode` VARCHAR(10) COMMENT '邮政编码',
  `pname` VARCHAR(50) COMMENT '省份名称',
  `cityname` VARCHAR(50) COMMENT '城市名称',
  `adname` VARCHAR(50) COMMENT '区域名称',
  `email` VARCHAR(100) COMMENT '电子邮箱',
  `photos` JSON COMMENT '照片信息',
  `entr_location` JSON NULL COMMENT '入口位置',
  `exit_location` JSON NULL COMMENT '出口位置',
  `groupbuy` TINYINT(1) DEFAULT 0 COMMENT '是否支持团购',
  `discount` TINYINT(1) DEFAULT 0 COMMENT '是否有优惠',
  `indoor_map` TINYINT(1) DEFAULT 0 COMMENT '是否有室内地图',
  `rating` DECIMAL(2,1) COMMENT '评分',
  `cost` DECIMAL(10,2) COMMENT '消费金额',
  `poiType` VARCHAR(10) COMMENT 'POI类型编码',
  `BigCategory` VARCHAR(50) COMMENT '大类别',
  `MidCategory` VARCHAR(50) COMMENT '中类别',
  `SubCategory` VARCHAR(50) COMMENT '子类别',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='兴趣点(POI)信息表';

-- 插入示例数据
INSERT INTO `poi_items` VALUES (
  'B0FFM8NQKE',
  '演员的秘密·皮肤管理(合生·麒麟社店)',
  '生活服务;美容美发店;美容美发店',
  '[116.477546, 39.997325]',
  '合生麒麟社2号楼',
  '',
  11,
  '1',
  '',
  '110000',
  '010',
  '110105',
  '',
  '北京市',
  '北京市',
  '朝阳区',
  '',
  '[
    {"title": "", "url": "http://store.is.autonavi.com/showpic/07ec9228235da549c5b3864be93d7f94"},
    {"title": "", "url": "http://store.is.autonavi.com/showpic/1dad4ab4fdf464ee95407bf049d994d7"},
    {"title": "", "url": "http://store.is.autonavi.com/showpic/ebeb319d5d039a4309775f9a9b7354fd"}
  ]',
  NULL,
  NULL,
  0,
  0,
  0,
  3.9,
  220.00,
  0,
  0,
  '071100',
  '生活服务',
  '美容美发店',
  '美容美发店'
);
