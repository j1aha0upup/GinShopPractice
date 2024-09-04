/*
 Navicat Premium Dump SQL

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80400 (8.4.0)
 Source Host           : localhost:3306
 Source Schema         : practice_gin_store

 Target Server Type    : MySQL
 Target Server Version : 80400 (8.4.0)
 File Encoding         : 65001

 Date: 04/09/2024 10:44:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for access
-- ----------------------------
DROP TABLE IF EXISTS `access`;
CREATE TABLE `access`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `module_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `action_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `type` tinyint NOT NULL,
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `status` tinyint NOT NULL,
  `module_id` int NOT NULL,
  `sort` int NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `add_time` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 36 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of access
-- ----------------------------
INSERT INTO `access` VALUES (15, '管理员管理', '模块', 1, '', 1, 0, 100, '', 1722309362);
INSERT INTO `access` VALUES (16, '管理员列表', '菜单', 2, '/admin/getAdminUser', 1, 15, 100, '', 1722309413);
INSERT INTO `access` VALUES (17, '增加管理员', '菜单', 2, '/admin/addAdminUser', 1, 15, 100, '', 1722309432);
INSERT INTO `access` VALUES (18, '角色管理', '模块', 1, '', 1, 0, 100, '', 1722309452);
INSERT INTO `access` VALUES (19, '角色列表', '菜单', 2, '/admin/roleList', 1, 18, 100, '', 1722309643);
INSERT INTO `access` VALUES (20, '增加角色', '菜单', 2, '/admin/addRole', 1, 18, 100, '', 1722309673);
INSERT INTO `access` VALUES (21, '权限管理', '模块', 1, '', 1, 0, 100, '', 1722309686);
INSERT INTO `access` VALUES (23, '权限列表', '菜单', 2, '/admin/accessList', 1, 21, 100, '', 1722309872);
INSERT INTO `access` VALUES (24, '增加权限', '菜单', 2, '/admin/addAccess', 1, 21, 100, '', 1722309897);
INSERT INTO `access` VALUES (25, '轮播图管理', '模块', 1, '', 1, 0, 100, '', 1722473966);
INSERT INTO `access` VALUES (26, '增加轮播图', '菜单', 2, '/admin/focusAdd', 1, 25, 101, '', 1722473991);
INSERT INTO `access` VALUES (27, '轮播图列表', '菜单', 2, '/admin/focusList', 1, 25, 102, '', 1722474012);
INSERT INTO `access` VALUES (28, '商品管理', '模块', 1, '', 1, 0, 100, '', 1722836920);
INSERT INTO `access` VALUES (29, '商品分类列表', '菜单', 2, '/admin/goodsCateList', 1, 28, 100, '', 1722836953);
INSERT INTO `access` VALUES (30, '商品分类添加', '动作', 3, '/admin/goodsCateAdd', 1, 28, 100, '', 1722836967);
INSERT INTO `access` VALUES (31, '商品类型列表', '菜单', 2, '/admin/goodsTypeList', 1, 28, 100, '', 1722906228);
INSERT INTO `access` VALUES (33, '商品列表', '菜单', 2, '/admin/goodsList', 1, 28, 100, '', 1722938520);
INSERT INTO `access` VALUES (34, '导航管理', '模块', 1, '', 1, 0, 100, '', 1723918527);
INSERT INTO `access` VALUES (35, '导航列表', '菜单', 2, '/admin/navList', 1, 34, 100, '', 1723918583);

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `default_address` tinyint NOT NULL,
  `add_time` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of address
-- ----------------------------
INSERT INTO `address` VALUES (5, 5, 'zzz', '15936030516', 'cccc', 0, 1725402438);
INSERT INTO `address` VALUES (6, 5, 'zzz', '15936030516', 'cccc', 0, 1725402438);
INSERT INTO `address` VALUES (7, 5, 'zzz', '15936030516', 'cccccccc', 0, 1725402448);
INSERT INTO `address` VALUES (8, 5, 'zzz', '15936030516', 'cccccccc', 0, 1725402448);
INSERT INTO `address` VALUES (9, 5, 'dddd', '15936030516', '22', 1, 1725402517);

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `mobile` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `status` int NOT NULL,
  `role_id` int NOT NULL,
  `add_time` int NOT NULL,
  `is_super` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_user
-- ----------------------------
INSERT INTO `admin_user` VALUES (1, 'admin', '21232f297a57a5a743894a0e4a801fc3', '15936030516', '3839003861@qq.com', 0, 4, 0, 1);

-- ----------------------------
-- Table structure for cart
-- ----------------------------
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `price` float NOT NULL,
  `goods_version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `uid` int NOT NULL,
  `num` int NOT NULL,
  `goods_gift` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `goods_fitting` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `goods_color` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `goods_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `goods_attr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `checked` tinyint NOT NULL,
  `goods_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of cart
-- ----------------------------

-- ----------------------------
-- Table structure for focus
-- ----------------------------
DROP TABLE IF EXISTS `focus`;
CREATE TABLE `focus`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `type` int NOT NULL,
  `focus_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `sort` int NOT NULL,
  `status` int NOT NULL,
  `add_time` int NOT NULL,
  `focus_img_small` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of focus
-- ----------------------------
INSERT INTO `focus` VALUES (10, '1', 1, 'static/uploaded/uploadedFocusImages/2024/9/3/20240903233339-1635757964474199700.jpg', 'https://www.baidu.com', 50, 1, 1725377619, 'static/uploaded/uploadedFocusImages/2024/9/3/20240903233339-1635757964474199700small.png');
INSERT INTO `focus` VALUES (11, '2', 1, 'static/uploaded/uploadedFocusImages/2024/9/3/20240903233347-1635757971905268100.jpg', 'https://www.baidu.com', 50, 1, 1725377627, 'static/uploaded/uploadedFocusImages/2024/9/3/20240903233347-1635757971905268100small.png');
INSERT INTO `focus` VALUES (12, '3', 1, 'static/uploaded/uploadedFocusImages/2024/9/3/20240903233356-1635757979944161500.jpg', 'https://www.baidu.com', 50, 1, 1725377636, 'static/uploaded/uploadedFocusImages/2024/9/3/20240903233356-1635757979944161500small.png');
INSERT INTO `focus` VALUES (13, '4', 1, 'static/uploaded/uploadedFocusImages/2024/9/3/20240903233403-1635758011028757700.jpg', 'https://www.baidu.com', 50, 1, 1725377643, 'static/uploaded/uploadedFocusImages/2024/9/3/20240903233403-1635758011028757700small.png');
INSERT INTO `focus` VALUES (14, '5', 1, 'static/uploaded/uploadedFocusImages/2024/9/3/20240903233412-1635758018523031700.jpg', 'https://www.baidu.com', 50, 1, 1725377652, 'static/uploaded/uploadedFocusImages/2024/9/3/20240903233412-1635758018523031700small.png');

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `sub_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `goods_sn` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `cate_id` int NOT NULL,
  `click_count` int NOT NULL,
  `goods_number` int NOT NULL,
  `shop_price` float NOT NULL,
  `market_price` float NOT NULL,
  `goods_relation` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `goods_attrs` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `goods_color` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `goods_version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `goods_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `goods_gift` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `goods_fitting` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `goods_keywords` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `goods_description` varchar(5000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `goods_content` varchar(5000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `is_delete` tinyint NOT NULL,
  `is_hot` tinyint NOT NULL,
  `is_best` tinyint NOT NULL,
  `is_new` tinyint NOT NULL,
  `goods_type_id` int NOT NULL,
  `sort` int NOT NULL,
  `status` int NOT NULL,
  `add_time` int NOT NULL,
  `more_choice` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `goods_images` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 62 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of goods
-- ----------------------------
INSERT INTO `goods` VALUES (19, '小米9-8GB+256GB', '火爆热卖中，6GB+64GB/6GB+128GB闪降100元，到手价1299元起', '', 1, 100, 1000, 999, 999, '19 53', '662 663 664 665 666 667 668 669 670', '', '8GB+256GB', 'static/uploaded/uploadedGoodsImages/2024/9/4/1592392307796676500.jpg', '1,2', '1,2', '', '<p>火爆热卖中，6GB+64GB/6GB+128GB闪降100元，到手价1299元起</p><p><br></p><p><br></p><p><br></p><p><br></p>', '', 0, 0, 1, 1, 1, 100, 1, 1592392307, '', ' 66');
INSERT INTO `goods` VALUES (20, 'Redmi Note 11 5G手机 1亿像素 55W有线闪充 50W无线闪充  6G+128GB 手机', '购机赠价值897元善诊3人体检套餐；赠Keep会员7日体验卡；光大信用卡分期支付满1000元减80元，数量有限', '', 2, 100, 100, 3699, 4599, '20 54', '671 672 673 674 675 676 677 678 679', '', '6G+128GB', 'static/uploaded/uploadedGoodsImages/2024/9/4/1637139107685884400.jpg', '', '', '', '<p>火爆热卖中，6GB+64GB/6GB+128GB闪降100元，到手价1299元起</p>', '', 0, 1, 1, 0, 1, 0, 1, 1592392495, '', ' 67');
INSERT INTO `goods` VALUES (21, '小米8年度旗舰', '火爆热卖中，6GB+64GB/6GB+128GB闪降100元，到手价1299元起', '', 36, 100, 1000, 1112, 1113, '1 2', '680 681 682 683 684 685 686 687 688', '', '3GB+32GB', 'static/uploaded/uploadedGoodsImages/2024/9/4/1635849810407008900.png', '1,2', '1,2', '1,2', '<p>火爆热卖中，6GB+64GB/6GB+128GB闪降100元，到手价1299元起</p><p><br></p><p><img src=\"http://bee.apiying.comstatic/upload/20211101/1635736323217965200.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p>', '', 0, 0, 1, 1, 1, 11, 1, 1592392825, '', ' 68');
INSERT INTO `goods` VALUES (22, 'Redmi 7A', '「3GB+32GB到手价仅549元」4000mAh超长续航 / 骁龙8核处理器 / 标配10W快充 / AI人脸解锁 / 大字体，大音量，无线收音机 / 整机生活防泼溅 / 极简模式，亲情守护', '', 2, 100, 1000, 549, 799, '', '698 699 700 701 702 703 704 705 706', '', '3GB+32GB', 'static/uploaded/uploadedGoodsImages/2024/9/4/1592820040.jpg', '', '', '', '<p><span style=\"color: rgb(51, 51, 51); font-family: F9ab65; font-size: 10.4922px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: left; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-style: initial; text-decoration-color: initial; display: inline !important; float: none;\">小巧机身蕴藏4000mAh大电量，配合MIUI系统级省电优化，精细调控，从此告别电量焦虑，尽情尽欢！</span></p>', '', 0, 0, 1, 0, 1, 100, 1, 1592820016, '', ' 70');
INSERT INTO `goods` VALUES (23, 'Redmi 智能电视 X65', '全金属边框/4K超高清/MEMC运动补偿/8单元重低音音响系统', '', 5, 100, 1000, 2999, 3299, '', '689 690 691 692 693 694 695 696 697', '4', '56寸', 'static/uploaded/uploadedGoodsImages/2024/9/4/1592820111.jpg', '', '', '', '<p><span style=\'color: rgb(176, 176, 176); font-family: \"Helvetica Neue\", Helvetica, Arial, \"Microsoft Yahei\", \"Hiragino Sans GB\", \"Heiti SC\", \"WenQuanYi Micro Hei\", sans-serif; font-size: 14px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-style: initial; text-decoration-color: initial; display: inline !important; float: none;\'>全金属边框/4K超高清/MEMC运动补偿/8单元重低音音响系统</span></p>', '', 0, 0, 1, 0, 1, 100, 0, 1592820111, '', ' 69');
INSERT INTO `goods` VALUES (24, 'RedmiBook 13 全面屏', '四窄边全面屏 / 全新十代酷睿™处理器 / 全金属超轻机身 / MX250 高性能独显 / 小米互传 / 专业「飓风」散热系统 / 11小时长续航', '', 39, 100, 1000, 4499, 4799, '', '707 708 709 710 711 712 713 714 715', '', '8G+128G', 'static/uploaded/uploadedGoodsImages/2024/9/4/1592820244.jpg', '', '', '', '<p><span style=\'color: rgb(176, 176, 176); font-family: \"Helvetica Neue\", Helvetica, Arial, \"Microsoft Yahei\", \"Hiragino Sans GB\", \"Heiti SC\", \"WenQuanYi Micro Hei\", sans-serif; font-size: 14px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-style: initial; text-decoration-color: initial; display: inline !important; float: none;\'>四窄边全面屏 / 全新十代酷睿&trade;处理器 / 全金属超轻机身 / MX250 高性能独显 / 小米互传 / 专业「飓风」散热系统 / 11小时长续航</span></p>', '', 0, 0, 1, 0, 1, 100, 1, 1592820244, '', ' 71');
INSERT INTO `goods` VALUES (25, '米家电磁炉', '99挡微调控火 / 支持低温烹饪 / 100+烹饪模式', '', 1, 100, 1000, 299, 399, '', '716 717 718 719 720 721 722 723 724', '', '', 'static/uploaded/uploadedGoodsImages/2024/9/4/1592820331.jpg', '', '', '', '<p>米家电磁炉</p>', '', 0, 1, 1, 0, 1, 100, 1, 1592820331, '', ' 72');
INSERT INTO `goods` VALUES (26, '黑鲨双向快充移动电源', '18W双向快充 / 铠甲机身 / 一入三出 / 炫酷灯效', '', 20, 100, 1000, 0, 0, '', '725 726 727 728 729 730 731 732 733', '', '', 'static/uploaded/uploadedGoodsImages/2024/9/4/1592820494.jpg', '', '', '', '', '', 0, 0, 1, 0, 1, 100, 1, 1592820494, '', ' 73');
INSERT INTO `goods` VALUES (36, '小米手机2222', '1111111111', '', 1, 100, 0, 4444, 5555, '', '734 735 736 737 738 739 740 741 742', '', '33332222222222', 'static/uploaded/uploadedGoodsImages/2024/9/4/1637218489739052500.jpg', '8888', '999999999', '', '<p>小米10</p><p><img src=\"static/upload/20211011/1633944295927657900.png\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p><p><img src=\"static/upload/20211011/1633944470896453500.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p>', '', 0, 0, 0, 0, 1, 12, 1, 1633755416, '', ' 74');
INSERT INTO `goods` VALUES (37, '小米电视测试', '222222222222222222', '', 1, 100, 0, 444444, 5555, '', '743 744 745 746 747 748 749 750 751', '', '333333333333333', 'static/uploaded/uploadedGoodsImages/2024/9/4/1633755741820253400.png', '8888888888', '999999', '111 1111111 111111111111111111', '<p>666666666666666</p>', '', 0, 1, 1, 1, 1, 0, 0, 1633755741, '', ' 75');
INSERT INTO `goods` VALUES (38, '小米手机测试111', '124214214214', '', 1, 100, 0, 0, 0, '', '788 789 790 791 792 793 794 795 796', '', '', 'static/uploaded/uploadedGoodsImages/2024/9/4/1637218459887075300.jpg', '', '', '', '', '', 0, 1, 1, 1, 1, 0, 0, 1633755959, '', ' 80');
INSERT INTO `goods` VALUES (39, 'Redmi k30', '6.53\"水滴大屏 | 5020mAh超长续航 | G80高性能处理器 | 全场景 AI 四摄 | 大功率扬声器 | 指纹识别 | 人脸解锁 | 红外遥控', '', 38, 100, 100, 899, 899, '', '779 780 781 782 783 784 785 786 787', '', '', 'static/uploaded/uploadedGoodsImages/2024/9/4/1637026344085801400.jpg', '', '', '', '', '', 0, 0, 0, 0, 1, 100, 1, 1635502706, '', ' 79');
INSERT INTO `goods` VALUES (40, 'Xiaomi MIX 4', 'CUP全面屏 | 真彩原色 + 120Hz | 一体化轻量陶瓷机身 | 高通骁龙™888+ | WiFi 6 增强版 | 石墨烯「冰封」散热系统', '', 37, 100, 100, 0, 0, '', '770 771 772 773 774 775 776 777 778', '', '', 'static/uploaded/uploadedGoodsImages/2024/9/4/1637026171480899500.jpg', '', '', '', '', '', 0, 0, 0, 0, 1, 100, 1, 1635503000, '', ' 78');
INSERT INTO `goods` VALUES (41, 'Xiaomi Civi', '轻薄潮流设计 | 丝绒AG工艺 | 原生美肌人像 | 像素级肌肤焕新技术 | 3200万高清质感自拍 | 双柔光灯+自动对焦 | 3D曲面OLED柔性屏 | 120Hz+Dolby Vision | 4500mAh 大电量 | 55W有线闪充 | 立体声双扬声器', '', 36, 100, 100, 1200, 1400, '', '761 762 763 764 765 766 767 768 769', '', '', 'static/uploaded/uploadedGoodsImages/2024/9/4/1637026086634961500.jpg', '', '', '', '', '', 0, 0, 0, 0, 1, 100, 1, 1635503077, '', ' 77');
INSERT INTO `goods` VALUES (42, 'Redmi Note 10 5G', '5G小金刚｜旗舰长续航｜双5G待机｜5000mAh充电宝级大容量｜4800万高清相机｜天玑700八核高性能处理器', '', 35, 100, 100, 0, 0, '', '752 753 754 755 756 757 758 759 760', '', '', 'static/uploaded/uploadedGoodsImages/2024/9/4/1637025991576339600.jpg', '', '', '', '', '', 0, 0, 0, 0, 1, 100, 1, 1635503644, '', ' 76');
INSERT INTO `goods` VALUES (43, 'Xiaomi 10S', '骁龙870 | 对称式双扬立体声 | 1亿像素 8K电影相机 | 33W有线快充 | 30W无线快充 | 10W反向充电 | 4780mAh超大电池 | LPDDR5+UFS3.0+Wi-Fi 6 | VC液冷散热 | 双模5G', '', 35, 100, 100, 2699, 3699, '', '', '1,2,3', '8GB+128GB', 'static/uploaded/uploadedGoodsImages/2024/9/4/1635841579767962200.jpg', '', '', '', '<p id=\"isPasted\"><br></p><p>高通骁龙&trade;870</p><p>哈曼卡顿｜对称式双扬立体声</p><p>4780mAh 大电量</p><p>三重快充 33W有线+30W无线+10W反向充电</p><p>小至尊经典外观</p><p>LPDDR5+UFS3.0+WiFi6</p><p>1 亿像素电影相机</p><p>8K 电影模式</p><p><br></p>', '', 0, 0, 1, 0, 1, 100, 1, 1635841578, '', '');
INSERT INTO `goods` VALUES (44, 'Xiaomi 11 Pro', '至高享24期免息，赠蓝牙耳机Air2 SE，+1元得30W立式无线充', '', 2, 100, 100, 0, 0, '', '', '2,3,4', '', 'static/uploaded/uploadedGoodsImages/2024/9/4/1635841908156579200.jpg', '', '', '', '<p><br></p><p id=\"isPasted\" style=\"text-align: center;\"><span style=\"font-size: 24px;\">联合研发18个月</span></p><p style=\"text-align: center;\"><span style=\"font-size: 24px;\">2亿影像投入，打造超强规格主摄</span></p><p style=\"text-align: center;\"><span style=\"font-size: 24px;\">这是颗&ldquo;巨型大底&rdquo;的面积，甚至可以媲美专业便携式相机，超大的进光量，</span></p><p style=\"text-align: center;\"><span style=\"font-size: 24px;\">带来了前所未有丰富的细节，&ldquo;夜视&rdquo;能力因此远超人眼，更能&ldquo;看懂&rdquo;夜色。</span></p><p><img src=\"/static/uploaded/uploadedGoodsImages/2024/9/4/1635841855622147000.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p><p><br></p>', '', 0, 0, 1, 0, 0, 100, 1, 1635841907, '', '');
INSERT INTO `goods` VALUES (45, '小米移动电源3 20000mAh USB-C双向快充版', '', '', 20, 100, 100, 100, 100, '', '', '', '', 'static/uploaded/uploadedGoodsImages/2024/9/4/1635844763742258900.jpg', '', '', '', '', '', 0, 0, 0, 0, 0, 100, 1, 1635844763, '', '');
INSERT INTO `goods` VALUES (46, '小米移动电源3 10000mAh 超级闪充版 （50W）', '', '', 20, 100, 100, 125, 155, '', '', '', '', 'static/uploaded/uploadedGoodsImages/2024/9/4/1635844808324401400.jpg', '', '', '', '', '', 0, 0, 0, 0, 0, 100, 1, 1635844808, '', '');
INSERT INTO `goods` VALUES (47, '小米6A Type-C快充数据线', '快速充电 | 快速传输 | 更强兼容 | 安全可靠', '', 9, 100, 100, 29, 29, '', '', '', '', 'static/uploaded/uploadedGoodsImages/2024/9/4/1637218054509467500.jpg', '', '', '', '', '', 0, 0, 0, 0, 0, 100, 1, 1635845354, '', '');
INSERT INTO `goods` VALUES (48, '小米USB-C数据线 编织线版 100cm', '', '', 9, 100, 100, 0, 0, '', '', '', '', 'static/uploaded/uploadedGoodsImages/2024/9/4/1635845426055325800.jpg', '', '', '', '<p><img src=\"/static/uploaded/uploadedGoodsImages/2024/9/4/1635845418913722200.png\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p>', '', 0, 0, 0, 0, 0, 100, 1, 1635845425, '', '');
INSERT INTO `goods` VALUES (49, 'Redmi Note 11 Pro系列', '三星AMOLED高刷屏 | JBL 对称式立体声 | 一亿像素超清影像 | 天玑920液冷芯 | VC液冷立体散热', '', 2, 100, 100, 0, 0, '', '', '', '', 'static/uploaded/uploadedGoodsImages/2024/9/4/1637025826328576500.jpg', '', '', '', '', '', 0, 0, 0, 0, 0, 100, 1, 1637025826, '', '');
INSERT INTO `goods` VALUES (50, '小米全面屏电视 55英寸PRO E55S', 'Amlogic T972超强悍处理器 / 4K超高清画质 细腻如真 / 支持8K视频解码 / 2G+32G超大存储 / 内置小爱同学 语音控制更方便 / 智能Patchwall汇聚海量好内容', '', 19, 100, 100, 2399, 2499, '', '', '', '', 'static/uploaded/uploadedGoodsImages/2024/9/4/1637049463471284100.jpg', '', '', '', '', '', 0, 0, 0, 0, 0, 100, 1, 1637049463, '', '');
INSERT INTO `goods` VALUES (51, '米家互联网对开门冰箱 540L', '风冷无霜/环绕出风/纤薄箱体/电脑控温,持久保鲜/智能互联', '', 13, 100, 100, 2899, 2999, '23 24 39', '', '1,2,3,4', '', 'static/uploaded/uploadedGoodsImages/2024/9/4/1637049592911969300.jpg', '23,24,39', '23,24,39', '', '', '', 0, 0, 0, 0, 0, 100, 1, 1637049592, '', '');
INSERT INTO `goods` VALUES (52, '米家微波炉', '智能APP操控 / 平板式加热 / 专项分类解冻 / 20L容量 / 30+精选食谱', '', 1, 100, 100, 399, 499, '', '', '', '', 'static/uploaded/uploadedGoodsImages/2024/9/4/1637049679925704800.jpg', '', '', '', '', '', 0, 0, 0, 0, 0, 100, 1, 1637049679, '', '');
INSERT INTO `goods` VALUES (53, '小米9 6GB+128GB', '', '', 38, 100, 100, 1113, 1167, '19 53', '', '2,6,7', '6GB+128GB', 'static/uploaded/uploadedGoodsImages/2024/9/4/1637063708413624300.jpg', '', '', '', '<p id=\"isPasted\">火爆热卖中，6GB+64GB/6GB+128GB闪降100元，到手价1299元起</p><p><br></p><p><img src=\"http://bee.apiying.comstatic/upload/20211101/1635739607166546900.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dii\"></p><p><img src=\"http://bee.apiying.comstatic/upload/20211101/1635740680831942900.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dii\"></p>', '', 0, 0, 0, 0, 0, 100, 1, 1637063708, '', '');
INSERT INTO `goods` VALUES (54, 'Redmi Note 11 5G手机 1亿像素 55W有线闪充 50W无线闪充  8G+256GB 手机', '双卡双5G | X轴线性马达 | 5000mAh 大电量 | 33W快充 | 立体声双扬声器 | 天玑810处理器 | 90Hz变速高刷屏', '', 2, 100, 100, 4199, 4599, '20 54', '颜色:红色,白色,黄色 | 尺寸:41,42,43', '2,8', ' 8G+256GB', 'static/uploaded/uploadedGoodsImages/2024/9/4/1637219213943897000.jpg', '', '', '', '<p id=\"isPasted\" style=\'margin: 0px; font-weight: bolder; font-size: 18px; color: rgb(0, 0, 0); font-family: \"Helvetica Neue\", Helvetica, Arial, \"Microsoft Yahei\", \"Hiragino Sans GB\", \"Heiti SC\", \"WenQuanYi Micro Hei\", sans-serif; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; letter-spacing: normal; orphans: 2; text-align: justify; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;\'><br>8GB + 256GB 最高可选</p><p style=\'margin: 0px; color: rgb(0, 0, 0); font-family: \"Helvetica Neue\", Helvetica, Arial, \"Microsoft Yahei\", \"Hiragino Sans GB\", \"Heiti SC\", \"WenQuanYi Micro Hei\", sans-serif; font-size: 17px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: justify; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;\'>4GB + 128GB</p><p style=\'margin: 0px; color: rgb(0, 0, 0); font-family: \"Helvetica Neue\", Helvetica, Arial, \"Microsoft Yahei\", \"Hiragino Sans GB\", \"Heiti SC\", \"WenQuanYi Micro Hei\", sans-serif; font-size: 17px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: justify; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;\'>6GB + 128GB</p><p style=\'margin: 0px; color: rgb(0, 0, 0); font-family: \"Helvetica Neue\", Helvetica, Arial, \"Microsoft Yahei\", \"Hiragino Sans GB\", \"Heiti SC\", \"WenQuanYi Micro Hei\", sans-serif; font-size: 17px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: justify; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;\'>8GB + 128GB</p><p style=\'margin: 0px; color: rgb(0, 0, 0); font-family: \"Helvetica Neue\", Helvetica, Arial, \"Microsoft Yahei\", \"Hiragino Sans GB\", \"Heiti SC\", \"WenQuanYi Micro Hei\", sans-serif; font-size: 17px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: justify; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;\'>8GB + 256GB</p><p style=\'margin: 0px; color: rgb(0, 0, 0); font-family: \"Helvetica Neue\", Helvetica, Arial, \"Microsoft Yahei\", \"Hiragino Sans GB\", \"Heiti SC\", \"WenQuanYi Micro Hei\", sans-serif; font-size: 17px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: justify; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;\'>LPDDR4X 内存 +UFS2.2 闪存</p>', '', 0, 0, 0, 0, 1, 100, 1, 1637139500, '', '');

-- ----------------------------
-- Table structure for goods_attribute
-- ----------------------------
DROP TABLE IF EXISTS `goods_attribute`;
CREATE TABLE `goods_attribute`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `goods_id` int NOT NULL,
  `type_id` int NOT NULL,
  `attribute_id` int NOT NULL,
  `attribute_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `attribute_type` int NOT NULL,
  `attribute_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `sort` int NOT NULL,
  `add_time` int NOT NULL,
  `status` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 797 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of goods_attribute
-- ----------------------------
INSERT INTO `goods_attribute` VALUES (108, 37, 2, 2, '主体', 3, '111\r\n', 10, 1634722866, 1);
INSERT INTO `goods_attribute` VALUES (109, 37, 2, 3, '内存', 1, '内存', 10, 1634722866, 1);
INSERT INTO `goods_attribute` VALUES (110, 37, 2, 4, '硬盘', 1, '硬盘', 10, 1634722866, 1);
INSERT INTO `goods_attribute` VALUES (111, 37, 2, 5, '显示器', 1, '显示器:', 10, 1634722866, 1);
INSERT INTO `goods_attribute` VALUES (112, 37, 2, 6, '支持蓝牙', 3, '否', 10, 1634722866, 1);
INSERT INTO `goods_attribute` VALUES (117, 43, 1, 1, '基本信息', 1, '高通骁龙™870', 10, 1637025929, 1);
INSERT INTO `goods_attribute` VALUES (118, 43, 1, 7, '性能	', 2, '4780mAh 大电量', 10, 1637025929, 1);
INSERT INTO `goods_attribute` VALUES (119, 43, 1, 8, '相机', 2, '4780mAh 大电量', 10, 1637025929, 1);
INSERT INTO `goods_attribute` VALUES (120, 43, 1, 9, '支持蓝牙', 3, '是\r\n', 10, 1637025929, 1);
INSERT INTO `goods_attribute` VALUES (157, 20, 1, 1, '基本信息', 1, '', 10, 1637139540, 1);
INSERT INTO `goods_attribute` VALUES (158, 20, 1, 7, '性能	', 2, '', 10, 1637139540, 1);
INSERT INTO `goods_attribute` VALUES (159, 20, 1, 8, '相机', 2, '', 10, 1637139541, 1);
INSERT INTO `goods_attribute` VALUES (160, 20, 1, 9, '支持蓝牙', 3, '是\r\n', 10, 1637139541, 1);
INSERT INTO `goods_attribute` VALUES (161, 38, 1, 1, '基本信息', 1, '124214', 10, 1637218460, 1);
INSERT INTO `goods_attribute` VALUES (162, 38, 1, 7, '性能	', 2, '214214', 10, 1637218460, 1);
INSERT INTO `goods_attribute` VALUES (163, 38, 1, 8, '相机', 2, '214214', 10, 1637218460, 1);
INSERT INTO `goods_attribute` VALUES (164, 38, 1, 9, '支持蓝牙', 3, '是\r\n', 10, 1637218460, 1);
INSERT INTO `goods_attribute` VALUES (165, 36, 2, 2, '主体', 3, '1111', 10, 1637218489, 1);
INSERT INTO `goods_attribute` VALUES (166, 36, 2, 3, '内存', 1, 'aaaaaaa', 10, 1637218489, 1);
INSERT INTO `goods_attribute` VALUES (167, 36, 2, 4, '硬盘', 1, 'bbbb', 10, 1637218489, 1);
INSERT INTO `goods_attribute` VALUES (168, 36, 2, 5, '显示器', 1, 'cccc', 10, 1637218490, 1);
INSERT INTO `goods_attribute` VALUES (169, 36, 2, 6, '支持蓝牙', 3, '是\r\n', 10, 1637218490, 1);
INSERT INTO `goods_attribute` VALUES (536, 54, 1, 1, '基本信息', 1, '', 10, 1637294588, 1);
INSERT INTO `goods_attribute` VALUES (537, 54, 1, 7, '性能	', 2, '### 天玑810\r\n\r\n* 6nm工艺制程\r\n\r\nCPU：八核处理器，最高主频2.4GHz\r\n\r\nGPU：Mali-G57', 10, 1637294588, 1);
INSERT INTO `goods_attribute` VALUES (538, 54, 1, 8, '相机', 2, '5000万超清主摄\r\n\r\n\r\n50MP，f/1.8 ，4in1\r\n\r\n\r\n800万超广角相机\r\n\r\n\r\n8MP，IMX355，119°', 10, 1637294588, 1);
INSERT INTO `goods_attribute` VALUES (539, 54, 1, 9, '支持蓝牙', 3, '是\r\n', 10, 1637294588, 1);
INSERT INTO `goods_attribute` VALUES (540, 54, 1, 14, '内存与容量', 2, '8GB + 256GB 最高可选\r\n\r\n4GB + 128GB\r\n\r\n6GB + 128GB\r\n\r\n8GB + 128GB\r\n\r\n8GB + 256GB\r\n\r\nLPDDR4X 内存 +UFS2.2 闪存', 10, 1637294588, 1);
INSERT INTO `goods_attribute` VALUES (541, 54, 1, 15, '外观尺寸', 2, '高度：163.56mm\r\n\r\n宽度：75.78mm\r\n\r\n厚度：8.75mm\r\n\r\n重量：195g', 10, 1637294588, 1);
INSERT INTO `goods_attribute` VALUES (542, 54, 1, 16, '充电与电池', 2, '5000mAh电池（典型值）\r\n\r\n标配33W充电器\r\n\r\n充电接口：Type-C充电接口', 10, 1637294588, 1);
INSERT INTO `goods_attribute` VALUES (543, 54, 1, 17, '影像系统', 1, '', 10, 1637294588, 1);
INSERT INTO `goods_attribute` VALUES (544, 54, 1, 18, '传感器', 1, '', 10, 1637294588, 1);
INSERT INTO `goods_attribute` VALUES (545, 19, 1, 1, '基本信息', 1, '', 10, 1637296719, 1);
INSERT INTO `goods_attribute` VALUES (546, 19, 1, 7, '性能	', 2, '### 高通骁龙™888\r\n\r\n**CPU 架构工艺：Kryo 680 架构，5nm 工艺制程**\r\n\r\nCPU 主频：八核处理器，最高主频可达：2.84GHz\r\n\r\nGPU ：Adreno 660 图形处理器，最高频率可达 840MHz \r\n\r\nAI：第六代 AI 引擎\r\n\r\n![尺寸](https://cdn.cnbj1.fds.api.mi-img.com/product-images/mix4/specs_m.png)\r\n\r\n\r\n', 10, 1637296719, 1);
INSERT INTO `goods_attribute` VALUES (547, 19, 1, 8, '相机', 2, '', 10, 1637296719, 1);
INSERT INTO `goods_attribute` VALUES (548, 19, 1, 9, '支持蓝牙', 3, '是\r\n', 10, 1637296719, 1);
INSERT INTO `goods_attribute` VALUES (549, 19, 1, 14, '内存与容量', 2, '12GB + 256GB 最高可选\r\n\r\n* 运行内存：8GB / 12GB LPDDR5 高速内存（6400Mbps）\r\n\r\n* 机身存储：128GB / 256GB UFS 3.1 高速存储', 10, 1637296719, 1);
INSERT INTO `goods_attribute` VALUES (550, 19, 1, 15, '外观尺寸', 2, '', 10, 1637296719, 1);
INSERT INTO `goods_attribute` VALUES (551, 19, 1, 16, '充电与电池', 2, '5000mAh（typ） / 4900mAh（min）\r\n\r\n内置锂离子聚合物电池，不可拆卸\r\n\r\nUSB Type-C 双面充电接口\r\n\r\n手机支持 QC4 / QC3+ / PD3.0 快充协议\r\n\r\n67W有线闪充/67W无线闪充/10W无线反充', 10, 1637296719, 1);
INSERT INTO `goods_attribute` VALUES (552, 19, 1, 17, '影像系统', 1, '', 10, 1637296719, 1);
INSERT INTO `goods_attribute` VALUES (553, 19, 1, 18, '传感器', 1, '', 10, 1637296719, 1);
INSERT INTO `goods_attribute` VALUES (554, 52, 1, 1, '基本信息', 1, '', 10, 1725384758, 1);
INSERT INTO `goods_attribute` VALUES (555, 52, 1, 7, '性能	', 2, '', 111, 1725384758, 1);
INSERT INTO `goods_attribute` VALUES (556, 52, 1, 8, '相机', 2, '', 0, 1725384758, 1);
INSERT INTO `goods_attribute` VALUES (557, 52, 1, 9, '支持蓝牙', 3, '', 0, 1725384758, 1);
INSERT INTO `goods_attribute` VALUES (558, 52, 1, 14, '内存与容量', 2, '', 10, 1725384758, 1);
INSERT INTO `goods_attribute` VALUES (559, 52, 1, 15, '外观尺寸', 2, '', 10, 1725384758, 1);
INSERT INTO `goods_attribute` VALUES (560, 52, 1, 16, '充电与电池', 2, '', 10, 1725384758, 1);
INSERT INTO `goods_attribute` VALUES (561, 52, 1, 17, '影像系统', 1, '', 10, 1725384758, 1);
INSERT INTO `goods_attribute` VALUES (562, 52, 1, 18, '传感器', 1, '', 10, 1725384758, 1);
INSERT INTO `goods_attribute` VALUES (581, 54, 1, 1, '基本信息', 1, '', 10, 1725384874, 1);
INSERT INTO `goods_attribute` VALUES (582, 54, 1, 7, '性能	', 2, '### 天玑810\r\n\r\n* 6nm工艺制程\r\n\r\nCPU：八核处理器，最高主频2.4GHz\r\n\r\nGPU：Mali-G57', 111, 1725384874, 1);
INSERT INTO `goods_attribute` VALUES (583, 54, 1, 8, '相机', 2, '5000万超清主摄\r\n\r\n\r\n50MP，f/1.8 ，4in1\r\n\r\n\r\n800万超广角相机\r\n\r\n\r\n8MP，IMX355，119°', 0, 1725384874, 1);
INSERT INTO `goods_attribute` VALUES (584, 54, 1, 9, '支持蓝牙', 3, '是\r\n', 0, 1725384874, 1);
INSERT INTO `goods_attribute` VALUES (585, 54, 1, 14, '内存与容量', 2, '8GB + 256GB 最高可选\r\n\r\n4GB + 128GB\r\n\r\n6GB + 128GB\r\n\r\n8GB + 128GB\r\n\r\n8GB + 256GB\r\n\r\nLPDDR4X 内存 +UFS2.2 闪存', 10, 1725384874, 1);
INSERT INTO `goods_attribute` VALUES (586, 54, 1, 15, '外观尺寸', 2, '高度：163.56mm\r\n\r\n宽度：75.78mm\r\n\r\n厚度：8.75mm\r\n\r\n重量：195g', 10, 1725384874, 1);
INSERT INTO `goods_attribute` VALUES (587, 54, 1, 16, '充电与电池', 2, '5000mAh电池（典型值）\r\n\r\n标配33W充电器\r\n\r\n充电接口：Type-C充电接口', 10, 1725384874, 1);
INSERT INTO `goods_attribute` VALUES (588, 54, 1, 17, '影像系统', 1, '', 10, 1725384874, 1);
INSERT INTO `goods_attribute` VALUES (589, 54, 1, 18, '传感器', 1, '', 10, 1725384874, 1);
INSERT INTO `goods_attribute` VALUES (608, 58, 1, 1, '基本信息', 1, '', 10, 1725387325, 1);
INSERT INTO `goods_attribute` VALUES (609, 58, 1, 7, '性能	', 2, '', 111, 1725387325, 1);
INSERT INTO `goods_attribute` VALUES (610, 58, 1, 8, '相机', 2, '', 0, 1725387325, 1);
INSERT INTO `goods_attribute` VALUES (611, 58, 1, 9, '支持蓝牙', 3, '是\r\n', 0, 1725387325, 1);
INSERT INTO `goods_attribute` VALUES (612, 58, 1, 14, '内存与容量', 2, '', 10, 1725387325, 1);
INSERT INTO `goods_attribute` VALUES (613, 58, 1, 15, '外观尺寸', 2, '', 10, 1725387325, 1);
INSERT INTO `goods_attribute` VALUES (614, 58, 1, 16, '充电与电池', 2, '', 10, 1725387325, 1);
INSERT INTO `goods_attribute` VALUES (615, 58, 1, 17, '影像系统', 1, '', 10, 1725387325, 1);
INSERT INTO `goods_attribute` VALUES (616, 58, 1, 18, '传感器', 1, '', 10, 1725387325, 1);
INSERT INTO `goods_attribute` VALUES (617, 60, 1, 1, '基本信息', 1, '', 10, 1725387409, 1);
INSERT INTO `goods_attribute` VALUES (618, 60, 1, 7, '性能	', 2, '', 111, 1725387409, 1);
INSERT INTO `goods_attribute` VALUES (619, 60, 1, 8, '相机', 2, '', 0, 1725387409, 1);
INSERT INTO `goods_attribute` VALUES (620, 60, 1, 9, '支持蓝牙', 3, '是\r\n', 0, 1725387409, 1);
INSERT INTO `goods_attribute` VALUES (621, 60, 1, 14, '内存与容量', 2, '', 10, 1725387409, 1);
INSERT INTO `goods_attribute` VALUES (622, 60, 1, 15, '外观尺寸', 2, '', 10, 1725387409, 1);
INSERT INTO `goods_attribute` VALUES (623, 60, 1, 16, '充电与电池', 2, '', 10, 1725387409, 1);
INSERT INTO `goods_attribute` VALUES (624, 60, 1, 17, '影像系统', 1, '', 10, 1725387409, 1);
INSERT INTO `goods_attribute` VALUES (625, 60, 1, 18, '传感器', 1, '', 10, 1725387409, 1);
INSERT INTO `goods_attribute` VALUES (626, 61, 1, 1, '基本信息', 1, '', 10, 1725387465, 1);
INSERT INTO `goods_attribute` VALUES (627, 61, 1, 7, '性能	', 2, '', 111, 1725387465, 1);
INSERT INTO `goods_attribute` VALUES (628, 61, 1, 8, '相机', 2, '', 0, 1725387465, 1);
INSERT INTO `goods_attribute` VALUES (629, 61, 1, 9, '支持蓝牙', 3, '是\r\n', 0, 1725387465, 1);
INSERT INTO `goods_attribute` VALUES (630, 61, 1, 14, '内存与容量', 2, '', 10, 1725387465, 1);
INSERT INTO `goods_attribute` VALUES (631, 61, 1, 15, '外观尺寸', 2, '', 10, 1725387465, 1);
INSERT INTO `goods_attribute` VALUES (632, 61, 1, 16, '充电与电池', 2, '', 10, 1725387465, 1);
INSERT INTO `goods_attribute` VALUES (633, 61, 1, 17, '影像系统', 1, '', 10, 1725387465, 1);
INSERT INTO `goods_attribute` VALUES (634, 61, 1, 18, '传感器', 1, '', 10, 1725387465, 1);
INSERT INTO `goods_attribute` VALUES (662, 19, 1, 1, '基本信息', 1, '', 10, 1725417607, 1);
INSERT INTO `goods_attribute` VALUES (663, 19, 1, 7, '性能	', 2, '', 111, 1725417607, 1);
INSERT INTO `goods_attribute` VALUES (664, 19, 1, 8, '相机', 2, '', 0, 1725417607, 1);
INSERT INTO `goods_attribute` VALUES (665, 19, 1, 9, '支持蓝牙', 3, '是\r\n', 0, 1725417607, 1);
INSERT INTO `goods_attribute` VALUES (666, 19, 1, 14, '内存与容量', 2, '', 10, 1725417607, 1);
INSERT INTO `goods_attribute` VALUES (667, 19, 1, 15, '外观尺寸', 2, '', 10, 1725417607, 1);
INSERT INTO `goods_attribute` VALUES (668, 19, 1, 16, '充电与电池', 2, '', 10, 1725417607, 1);
INSERT INTO `goods_attribute` VALUES (669, 19, 1, 17, '影像系统', 1, '', 10, 1725417607, 1);
INSERT INTO `goods_attribute` VALUES (670, 19, 1, 18, '传感器', 1, '', 10, 1725417607, 1);
INSERT INTO `goods_attribute` VALUES (671, 20, 1, 1, '基本信息', 1, '', 10, 1725417614, 1);
INSERT INTO `goods_attribute` VALUES (672, 20, 1, 7, '性能	', 2, '', 111, 1725417614, 1);
INSERT INTO `goods_attribute` VALUES (673, 20, 1, 8, '相机', 2, '', 0, 1725417614, 1);
INSERT INTO `goods_attribute` VALUES (674, 20, 1, 9, '支持蓝牙', 3, '是\r\n', 0, 1725417614, 1);
INSERT INTO `goods_attribute` VALUES (675, 20, 1, 14, '内存与容量', 2, '', 10, 1725417614, 1);
INSERT INTO `goods_attribute` VALUES (676, 20, 1, 15, '外观尺寸', 2, '', 10, 1725417614, 1);
INSERT INTO `goods_attribute` VALUES (677, 20, 1, 16, '充电与电池', 2, '', 10, 1725417614, 1);
INSERT INTO `goods_attribute` VALUES (678, 20, 1, 17, '影像系统', 1, '', 10, 1725417614, 1);
INSERT INTO `goods_attribute` VALUES (679, 20, 1, 18, '传感器', 1, '', 10, 1725417614, 1);
INSERT INTO `goods_attribute` VALUES (680, 21, 1, 1, '基本信息', 1, '', 10, 1725417624, 1);
INSERT INTO `goods_attribute` VALUES (681, 21, 1, 7, '性能	', 2, '', 111, 1725417624, 1);
INSERT INTO `goods_attribute` VALUES (682, 21, 1, 8, '相机', 2, '', 0, 1725417624, 1);
INSERT INTO `goods_attribute` VALUES (683, 21, 1, 9, '支持蓝牙', 3, '', 0, 1725417624, 1);
INSERT INTO `goods_attribute` VALUES (684, 21, 1, 14, '内存与容量', 2, '', 10, 1725417624, 1);
INSERT INTO `goods_attribute` VALUES (685, 21, 1, 15, '外观尺寸', 2, '', 10, 1725417624, 1);
INSERT INTO `goods_attribute` VALUES (686, 21, 1, 16, '充电与电池', 2, '', 10, 1725417624, 1);
INSERT INTO `goods_attribute` VALUES (687, 21, 1, 17, '影像系统', 1, '', 10, 1725417624, 1);
INSERT INTO `goods_attribute` VALUES (688, 21, 1, 18, '传感器', 1, '', 10, 1725417624, 1);
INSERT INTO `goods_attribute` VALUES (689, 23, 1, 1, '基本信息', 1, '', 10, 1725417632, 1);
INSERT INTO `goods_attribute` VALUES (690, 23, 1, 7, '性能	', 2, '', 111, 1725417632, 1);
INSERT INTO `goods_attribute` VALUES (691, 23, 1, 8, '相机', 2, '', 0, 1725417632, 1);
INSERT INTO `goods_attribute` VALUES (692, 23, 1, 9, '支持蓝牙', 3, '', 0, 1725417632, 1);
INSERT INTO `goods_attribute` VALUES (693, 23, 1, 14, '内存与容量', 2, '', 10, 1725417632, 1);
INSERT INTO `goods_attribute` VALUES (694, 23, 1, 15, '外观尺寸', 2, '', 10, 1725417632, 1);
INSERT INTO `goods_attribute` VALUES (695, 23, 1, 16, '充电与电池', 2, '', 10, 1725417632, 1);
INSERT INTO `goods_attribute` VALUES (696, 23, 1, 17, '影像系统', 1, '', 10, 1725417632, 1);
INSERT INTO `goods_attribute` VALUES (697, 23, 1, 18, '传感器', 1, '', 10, 1725417632, 1);
INSERT INTO `goods_attribute` VALUES (698, 22, 1, 1, '基本信息', 1, '', 10, 1725417640, 1);
INSERT INTO `goods_attribute` VALUES (699, 22, 1, 7, '性能	', 2, '', 111, 1725417640, 1);
INSERT INTO `goods_attribute` VALUES (700, 22, 1, 8, '相机', 2, '', 0, 1725417640, 1);
INSERT INTO `goods_attribute` VALUES (701, 22, 1, 9, '支持蓝牙', 3, '', 0, 1725417640, 1);
INSERT INTO `goods_attribute` VALUES (702, 22, 1, 14, '内存与容量', 2, '', 10, 1725417640, 1);
INSERT INTO `goods_attribute` VALUES (703, 22, 1, 15, '外观尺寸', 2, '', 10, 1725417640, 1);
INSERT INTO `goods_attribute` VALUES (704, 22, 1, 16, '充电与电池', 2, '', 10, 1725417640, 1);
INSERT INTO `goods_attribute` VALUES (705, 22, 1, 17, '影像系统', 1, '', 10, 1725417640, 1);
INSERT INTO `goods_attribute` VALUES (706, 22, 1, 18, '传感器', 1, '', 10, 1725417640, 1);
INSERT INTO `goods_attribute` VALUES (707, 24, 1, 1, '基本信息', 1, '', 10, 1725417671, 1);
INSERT INTO `goods_attribute` VALUES (708, 24, 1, 7, '性能	', 2, '', 111, 1725417671, 1);
INSERT INTO `goods_attribute` VALUES (709, 24, 1, 8, '相机', 2, '', 0, 1725417671, 1);
INSERT INTO `goods_attribute` VALUES (710, 24, 1, 9, '支持蓝牙', 3, '', 0, 1725417671, 1);
INSERT INTO `goods_attribute` VALUES (711, 24, 1, 14, '内存与容量', 2, '', 10, 1725417671, 1);
INSERT INTO `goods_attribute` VALUES (712, 24, 1, 15, '外观尺寸', 2, '', 10, 1725417671, 1);
INSERT INTO `goods_attribute` VALUES (713, 24, 1, 16, '充电与电池', 2, '', 10, 1725417671, 1);
INSERT INTO `goods_attribute` VALUES (714, 24, 1, 17, '影像系统', 1, '', 10, 1725417671, 1);
INSERT INTO `goods_attribute` VALUES (715, 24, 1, 18, '传感器', 1, '', 10, 1725417671, 1);
INSERT INTO `goods_attribute` VALUES (716, 25, 1, 1, '基本信息', 1, '', 10, 1725417677, 1);
INSERT INTO `goods_attribute` VALUES (717, 25, 1, 7, '性能	', 2, '', 111, 1725417677, 1);
INSERT INTO `goods_attribute` VALUES (718, 25, 1, 8, '相机', 2, '', 0, 1725417677, 1);
INSERT INTO `goods_attribute` VALUES (719, 25, 1, 9, '支持蓝牙', 3, '', 0, 1725417677, 1);
INSERT INTO `goods_attribute` VALUES (720, 25, 1, 14, '内存与容量', 2, '', 10, 1725417677, 1);
INSERT INTO `goods_attribute` VALUES (721, 25, 1, 15, '外观尺寸', 2, '', 10, 1725417677, 1);
INSERT INTO `goods_attribute` VALUES (722, 25, 1, 16, '充电与电池', 2, '', 10, 1725417677, 1);
INSERT INTO `goods_attribute` VALUES (723, 25, 1, 17, '影像系统', 1, '', 10, 1725417677, 1);
INSERT INTO `goods_attribute` VALUES (724, 25, 1, 18, '传感器', 1, '', 10, 1725417677, 1);
INSERT INTO `goods_attribute` VALUES (725, 26, 1, 1, '基本信息', 1, '', 10, 1725417683, 1);
INSERT INTO `goods_attribute` VALUES (726, 26, 1, 7, '性能	', 2, '', 111, 1725417683, 1);
INSERT INTO `goods_attribute` VALUES (727, 26, 1, 8, '相机', 2, '', 0, 1725417683, 1);
INSERT INTO `goods_attribute` VALUES (728, 26, 1, 9, '支持蓝牙', 3, '', 0, 1725417683, 1);
INSERT INTO `goods_attribute` VALUES (729, 26, 1, 14, '内存与容量', 2, '', 10, 1725417683, 1);
INSERT INTO `goods_attribute` VALUES (730, 26, 1, 15, '外观尺寸', 2, '', 10, 1725417683, 1);
INSERT INTO `goods_attribute` VALUES (731, 26, 1, 16, '充电与电池', 2, '', 10, 1725417683, 1);
INSERT INTO `goods_attribute` VALUES (732, 26, 1, 17, '影像系统', 1, '', 10, 1725417683, 1);
INSERT INTO `goods_attribute` VALUES (733, 26, 1, 18, '传感器', 1, '', 10, 1725417683, 1);
INSERT INTO `goods_attribute` VALUES (734, 36, 1, 1, '基本信息', 1, '', 10, 1725417689, 1);
INSERT INTO `goods_attribute` VALUES (735, 36, 1, 7, '性能	', 2, '', 111, 1725417689, 1);
INSERT INTO `goods_attribute` VALUES (736, 36, 1, 8, '相机', 2, '', 0, 1725417689, 1);
INSERT INTO `goods_attribute` VALUES (737, 36, 1, 9, '支持蓝牙', 3, '', 0, 1725417689, 1);
INSERT INTO `goods_attribute` VALUES (738, 36, 1, 14, '内存与容量', 2, '', 10, 1725417689, 1);
INSERT INTO `goods_attribute` VALUES (739, 36, 1, 15, '外观尺寸', 2, '', 10, 1725417689, 1);
INSERT INTO `goods_attribute` VALUES (740, 36, 1, 16, '充电与电池', 2, '', 10, 1725417689, 1);
INSERT INTO `goods_attribute` VALUES (741, 36, 1, 17, '影像系统', 1, '', 10, 1725417689, 1);
INSERT INTO `goods_attribute` VALUES (742, 36, 1, 18, '传感器', 1, '', 10, 1725417689, 1);
INSERT INTO `goods_attribute` VALUES (743, 37, 1, 1, '基本信息', 1, '', 10, 1725417696, 1);
INSERT INTO `goods_attribute` VALUES (744, 37, 1, 7, '性能	', 2, '', 111, 1725417696, 1);
INSERT INTO `goods_attribute` VALUES (745, 37, 1, 8, '相机', 2, '', 0, 1725417696, 1);
INSERT INTO `goods_attribute` VALUES (746, 37, 1, 9, '支持蓝牙', 3, '', 0, 1725417696, 1);
INSERT INTO `goods_attribute` VALUES (747, 37, 1, 14, '内存与容量', 2, '', 10, 1725417696, 1);
INSERT INTO `goods_attribute` VALUES (748, 37, 1, 15, '外观尺寸', 2, '', 10, 1725417696, 1);
INSERT INTO `goods_attribute` VALUES (749, 37, 1, 16, '充电与电池', 2, '', 10, 1725417696, 1);
INSERT INTO `goods_attribute` VALUES (750, 37, 1, 17, '影像系统', 1, '', 10, 1725417696, 1);
INSERT INTO `goods_attribute` VALUES (751, 37, 1, 18, '传感器', 1, '', 10, 1725417696, 1);
INSERT INTO `goods_attribute` VALUES (752, 42, 1, 1, '基本信息', 1, '', 10, 1725417701, 1);
INSERT INTO `goods_attribute` VALUES (753, 42, 1, 7, '性能	', 2, '', 111, 1725417701, 1);
INSERT INTO `goods_attribute` VALUES (754, 42, 1, 8, '相机', 2, '', 0, 1725417701, 1);
INSERT INTO `goods_attribute` VALUES (755, 42, 1, 9, '支持蓝牙', 3, '', 0, 1725417701, 1);
INSERT INTO `goods_attribute` VALUES (756, 42, 1, 14, '内存与容量', 2, '', 10, 1725417701, 1);
INSERT INTO `goods_attribute` VALUES (757, 42, 1, 15, '外观尺寸', 2, '', 10, 1725417701, 1);
INSERT INTO `goods_attribute` VALUES (758, 42, 1, 16, '充电与电池', 2, '', 10, 1725417701, 1);
INSERT INTO `goods_attribute` VALUES (759, 42, 1, 17, '影像系统', 1, '', 10, 1725417701, 1);
INSERT INTO `goods_attribute` VALUES (760, 42, 1, 18, '传感器', 1, '', 10, 1725417701, 1);
INSERT INTO `goods_attribute` VALUES (761, 41, 1, 1, '基本信息', 1, '', 10, 1725417708, 1);
INSERT INTO `goods_attribute` VALUES (762, 41, 1, 7, '性能	', 2, '', 111, 1725417708, 1);
INSERT INTO `goods_attribute` VALUES (763, 41, 1, 8, '相机', 2, '', 0, 1725417708, 1);
INSERT INTO `goods_attribute` VALUES (764, 41, 1, 9, '支持蓝牙', 3, '', 0, 1725417708, 1);
INSERT INTO `goods_attribute` VALUES (765, 41, 1, 14, '内存与容量', 2, '', 10, 1725417708, 1);
INSERT INTO `goods_attribute` VALUES (766, 41, 1, 15, '外观尺寸', 2, '', 10, 1725417708, 1);
INSERT INTO `goods_attribute` VALUES (767, 41, 1, 16, '充电与电池', 2, '', 10, 1725417708, 1);
INSERT INTO `goods_attribute` VALUES (768, 41, 1, 17, '影像系统', 1, '', 10, 1725417708, 1);
INSERT INTO `goods_attribute` VALUES (769, 41, 1, 18, '传感器', 1, '', 10, 1725417708, 1);
INSERT INTO `goods_attribute` VALUES (770, 40, 1, 1, '基本信息', 1, '', 10, 1725417714, 1);
INSERT INTO `goods_attribute` VALUES (771, 40, 1, 7, '性能	', 2, '', 111, 1725417714, 1);
INSERT INTO `goods_attribute` VALUES (772, 40, 1, 8, '相机', 2, '', 0, 1725417714, 1);
INSERT INTO `goods_attribute` VALUES (773, 40, 1, 9, '支持蓝牙', 3, '', 0, 1725417714, 1);
INSERT INTO `goods_attribute` VALUES (774, 40, 1, 14, '内存与容量', 2, '', 10, 1725417714, 1);
INSERT INTO `goods_attribute` VALUES (775, 40, 1, 15, '外观尺寸', 2, '', 10, 1725417714, 1);
INSERT INTO `goods_attribute` VALUES (776, 40, 1, 16, '充电与电池', 2, '', 10, 1725417714, 1);
INSERT INTO `goods_attribute` VALUES (777, 40, 1, 17, '影像系统', 1, '', 10, 1725417714, 1);
INSERT INTO `goods_attribute` VALUES (778, 40, 1, 18, '传感器', 1, '', 10, 1725417714, 1);
INSERT INTO `goods_attribute` VALUES (779, 39, 1, 1, '基本信息', 1, '', 10, 1725417719, 1);
INSERT INTO `goods_attribute` VALUES (780, 39, 1, 7, '性能	', 2, '', 111, 1725417719, 1);
INSERT INTO `goods_attribute` VALUES (781, 39, 1, 8, '相机', 2, '', 0, 1725417719, 1);
INSERT INTO `goods_attribute` VALUES (782, 39, 1, 9, '支持蓝牙', 3, '', 0, 1725417719, 1);
INSERT INTO `goods_attribute` VALUES (783, 39, 1, 14, '内存与容量', 2, '', 10, 1725417719, 1);
INSERT INTO `goods_attribute` VALUES (784, 39, 1, 15, '外观尺寸', 2, '', 10, 1725417719, 1);
INSERT INTO `goods_attribute` VALUES (785, 39, 1, 16, '充电与电池', 2, '', 10, 1725417719, 1);
INSERT INTO `goods_attribute` VALUES (786, 39, 1, 17, '影像系统', 1, '', 10, 1725417719, 1);
INSERT INTO `goods_attribute` VALUES (787, 39, 1, 18, '传感器', 1, '', 10, 1725417719, 1);
INSERT INTO `goods_attribute` VALUES (788, 38, 1, 1, '基本信息', 1, '124214', 10, 1725417726, 1);
INSERT INTO `goods_attribute` VALUES (789, 38, 1, 7, '性能	', 2, '214214', 111, 1725417726, 1);
INSERT INTO `goods_attribute` VALUES (790, 38, 1, 8, '相机', 2, '214214', 0, 1725417726, 1);
INSERT INTO `goods_attribute` VALUES (791, 38, 1, 9, '支持蓝牙', 3, '是\r\n', 0, 1725417726, 1);
INSERT INTO `goods_attribute` VALUES (792, 38, 1, 14, '内存与容量', 2, '', 10, 1725417726, 1);
INSERT INTO `goods_attribute` VALUES (793, 38, 1, 15, '外观尺寸', 2, '', 10, 1725417726, 1);
INSERT INTO `goods_attribute` VALUES (794, 38, 1, 16, '充电与电池', 2, '', 10, 1725417726, 1);
INSERT INTO `goods_attribute` VALUES (795, 38, 1, 17, '影像系统', 1, '', 10, 1725417726, 1);
INSERT INTO `goods_attribute` VALUES (796, 38, 1, 18, '传感器', 1, '', 10, 1725417726, 1);

-- ----------------------------
-- Table structure for goods_cate
-- ----------------------------
DROP TABLE IF EXISTS `goods_cate`;
CREATE TABLE `goods_cate`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `cate_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `template` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `pid` int NOT NULL,
  `sub_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `keywords` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` int NOT NULL,
  `sort` int NOT NULL,
  `add_time` int NOT NULL,
  `filter_attribute` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 40 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of goods_cate
-- ----------------------------
INSERT INTO `goods_cate` VALUES (1, '手机', '', '', '', 0, '手机', '手机', '手机', 1, 10, 1582461745, '');
INSERT INTO `goods_cate` VALUES (2, '小米11 Pro', 'static/uploaded/uploadedGoodsCate/2024/9/4/1582463294.png', '', '', 1, '小米10 Pro', '小米10 Pro', '小米10 Pro', 1, 0, 1582463294, '');
INSERT INTO `goods_cate` VALUES (3, 'Redmi 8', 'static/uploaded/uploadedGoodsCate/2024/9/4/1582463357.png', 'https://www.baidu.com', '11', 1, 'Redmi 8 11', 'Redmi 8 111', 'Redmi 8 111', 1, 11, 1582463357, '');
INSERT INTO `goods_cate` VALUES (4, '电视 笔记本', '', '', '', 0, '电视 盒子', '电视 盒子', '电视 盒子', 1, 0, 1582463515, '');
INSERT INTO `goods_cate` VALUES (5, '小米电视5 55英寸', 'static/uploaded/uploadedGoodsCate/2024/9/4/1582464603.png', '', '', 4, '小米电视5 55英寸', '小米电视5 55英寸', '小米电视5 55英寸', 1, 0, 1582464603, '');
INSERT INTO `goods_cate` VALUES (6, '家电', '', '', 'https://www.baidu.com', 0, '', '', '', 1, 0, 1582513219, '');
INSERT INTO `goods_cate` VALUES (7, '出行 穿戴', '', '', '', 0, '', '', '', 1, 0, 1582513235, '');
INSERT INTO `goods_cate` VALUES (8, '智能 路由器', '', '', '', 0, '', '', '', 1, 0, 1582513270, '');
INSERT INTO `goods_cate` VALUES (9, '电源 配件', '', '', '', 0, '', '', '', 1, 0, 1582513285, '');
INSERT INTO `goods_cate` VALUES (10, '健康 儿童', 'static/uploaded/uploadedGoodsCate/2024/9/4/1635413604527197900.jpg', '', '', 0, '', '', '', 1, 0, 1582513300, '');
INSERT INTO `goods_cate` VALUES (11, '耳机 音响', '', '', '', 0, '', '', '', 1, 0, 1582513338, '');
INSERT INTO `goods_cate` VALUES (12, '生活 箱包', '', '', '', 0, '', '', '', 1, 0, 1582513349, '');
INSERT INTO `goods_cate` VALUES (13, '冰箱', 'static/uploaded/uploadedGoodsCate/2024/9/4/1582513945.jpg', '', '', 6, '冰箱', '冰箱', '冰箱', 1, 0, 1582513945, '');
INSERT INTO `goods_cate` VALUES (14, '微波炉', 'static/uploaded/uploadedGoodsCate/2024/9/4/1582514001.jpg', '', '', 6, '', '', '', 1, 0, 1582513960, '');
INSERT INTO `goods_cate` VALUES (15, '小米手表', 'static/uploaded/uploadedGoodsCate/2024/9/4/1582514113.png', '', '', 7, '小米手表', '小米手表', '小米手表', 1, 0, 1582514113, '');
INSERT INTO `goods_cate` VALUES (16, '平衡车', 'static/uploaded/uploadedGoodsCate/2024/9/4/1582514151.jpg', '', '', 7, '平衡车', '平衡车', '平衡车', 1, 0, 1582514151, '');
INSERT INTO `goods_cate` VALUES (17, '路由器', 'static/uploaded/uploadedGoodsCate/2024/9/4/1582514289.png', '', '', 8, '路由器', '路由器', '路由器', 1, 0, 1582514289, '');
INSERT INTO `goods_cate` VALUES (18, '摄像机', 'static/uploaded/uploadedGoodsCate/2024/9/4/1582514318.jpg', '', '', 8, '摄像机', '摄像机', '摄像机', 1, 0, 1582514318, '');
INSERT INTO `goods_cate` VALUES (19, '全屏电视55寸', 'static/uploaded/uploadedGoodsCate/2024/9/4/1582514664.jpg', '', '', 4, '', '', '', 1, 0, 1582514664, '');
INSERT INTO `goods_cate` VALUES (20, '移动电源', 'static/uploaded/uploadedGoodsCate/2024/9/4/1582514810.png', '', '', 9, '移动电源', '移动电源', '移动电源', 1, 0, 1582514810, '');
INSERT INTO `goods_cate` VALUES (35, 'Xiaomi 10S', 'static/uploaded/uploadedGoodsCate/2024/9/4/1635841294026066400.png', '', '', 1, '', '', '', 1, 10, 1635817714, '');
INSERT INTO `goods_cate` VALUES (36, 'Xiaomi Civi', 'static/uploaded/uploadedGoodsCate/2024/9/4/1635841252665099500.png', '', '', 1, '', '', '', 1, 10, 1635841252, '');
INSERT INTO `goods_cate` VALUES (37, 'Xiaomi MIX 4', 'static/uploaded/uploadedGoodsCate/2024/9/4/1635841362004932300.png', '', '', 1, '', '', '', 1, 10, 1635841362, '');
INSERT INTO `goods_cate` VALUES (38, 'Redmi K30S 至尊纪念版', 'static/uploaded/uploadedGoodsCate/2024/9/4/1635841411131518300.png', '', '', 1, '', '', '', 1, 10, 1635841411, '');
INSERT INTO `goods_cate` VALUES (39, '笔记本', 'static/uploaded/uploadedGoodsCate/2024/9/4/1637218647770107500.jpg', '', '', 4, '', '', '', 1, 10, 1637218647, '');

-- ----------------------------
-- Table structure for goods_color
-- ----------------------------
DROP TABLE IF EXISTS `goods_color`;
CREATE TABLE `goods_color`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `color_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `color_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of goods_color
-- ----------------------------
INSERT INTO `goods_color` VALUES (1, '红色', 'red', 1);
INSERT INTO `goods_color` VALUES (2, '黑色', '#000', 1);
INSERT INTO `goods_color` VALUES (3, '黄色', 'yellow', 1);
INSERT INTO `goods_color` VALUES (4, '金色', '#ebf10f', 1);
INSERT INTO `goods_color` VALUES (5, '灰色', '#eee', 1);
INSERT INTO `goods_color` VALUES (6, '紫色', '#9932CD ', 1);
INSERT INTO `goods_color` VALUES (7, '淡绿色', '#90EE90', 1);
INSERT INTO `goods_color` VALUES (8, '蓝色', 'blue', 1);

-- ----------------------------
-- Table structure for goods_image
-- ----------------------------
DROP TABLE IF EXISTS `goods_image`;
CREATE TABLE `goods_image`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `goods_id` int NOT NULL,
  `image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `color_id` int NOT NULL,
  `sort` int NOT NULL,
  `add_time` int NOT NULL,
  `status` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 81 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of goods_image
-- ----------------------------
INSERT INTO `goods_image` VALUES (5, 37, 'static/uploaded/uploadedGoodsImages/2024/9/4/1633755740718752300.png', 2, 10, 1633755741, 1);
INSERT INTO `goods_image` VALUES (6, 37, 'static/uploaded/uploadedGoodsImages/2024/9/4/1633755740714630100.jpg', 4, 10, 1633755741, 1);
INSERT INTO `goods_image` VALUES (19, 43, 'static/uploaded/uploadedGoodsImages/2024/9/4/1635841578192734800.png', 0, 10, 1635841580, 1);
INSERT INTO `goods_image` VALUES (20, 44, 'static/uploaded/uploadedGoodsImages/2024/9/4/1635841907018281600.jpg', 0, 10, 1635841908, 1);
INSERT INTO `goods_image` VALUES (21, 19, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637052716813265400.jpg', 0, 10, 1637052718, 1);
INSERT INTO `goods_image` VALUES (22, 19, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637052716850583600.jpg', 0, 10, 1637052718, 1);
INSERT INTO `goods_image` VALUES (23, 19, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637052716884345200.jpg', 0, 10, 1637052718, 1);
INSERT INTO `goods_image` VALUES (27, 20, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637063586126852600.jpg', 8, 10, 1637063587, 1);
INSERT INTO `goods_image` VALUES (28, 53, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637063716307468700.jpg', 2, 10, 1637063716, 1);
INSERT INTO `goods_image` VALUES (29, 53, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637063716342727800.jpg', 2, 10, 1637063716, 1);
INSERT INTO `goods_image` VALUES (30, 53, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637138323678153500.jpg', 6, 10, 1637138326, 1);
INSERT INTO `goods_image` VALUES (31, 53, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637138323703994500.jpg', 0, 10, 1637138326, 1);
INSERT INTO `goods_image` VALUES (32, 53, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637138323728025100.jpg', 7, 10, 1637138326, 1);
INSERT INTO `goods_image` VALUES (33, 53, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637138324148951200.jpg', 7, 10, 1637138326, 1);
INSERT INTO `goods_image` VALUES (34, 53, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637138324173917100.jpg', 7, 10, 1637138326, 1);
INSERT INTO `goods_image` VALUES (35, 20, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637139106080100100.jpg', 5, 10, 1637139108, 1);
INSERT INTO `goods_image` VALUES (36, 20, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637139106095482200.jpg', 5, 10, 1637139108, 1);
INSERT INTO `goods_image` VALUES (37, 20, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637139106121304700.jpg', 7, 10, 1637139108, 1);
INSERT INTO `goods_image` VALUES (38, 20, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637139106142051100.jpg', 2, 10, 1637139108, 1);
INSERT INTO `goods_image` VALUES (39, 20, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637139106176296000.jpg', 8, 10, 1637139108, 1);
INSERT INTO `goods_image` VALUES (40, 20, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637139106213056300.jpg', 7, 10, 1637139108, 1);
INSERT INTO `goods_image` VALUES (41, 54, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637139499477843800.jpg', 2, 10, 1637139501, 1);
INSERT INTO `goods_image` VALUES (42, 54, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637139499510094900.jpg', 2, 10, 1637139501, 1);
INSERT INTO `goods_image` VALUES (43, 54, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637139499528738400.jpg', 8, 10, 1637139501, 1);
INSERT INTO `goods_image` VALUES (44, 54, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637139499755592600.jpg', 8, 10, 1637139501, 1);
INSERT INTO `goods_image` VALUES (47, 52, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637217857219083500.jpg', 0, 10, 1637217858, 1);
INSERT INTO `goods_image` VALUES (48, 51, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637217872435156200.jpg', 0, 10, 1637217873, 1);
INSERT INTO `goods_image` VALUES (49, 50, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637217890454460200.jpg', 0, 10, 1637217895, 1);
INSERT INTO `goods_image` VALUES (50, 49, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637217907397100400.jpg', 0, 10, 1637217908, 1);
INSERT INTO `goods_image` VALUES (51, 48, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637217953241815700.jpg', 0, 10, 1637217954, 1);
INSERT INTO `goods_image` VALUES (52, 46, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637218111418352100.jpg', 0, 10, 1637218112, 1);
INSERT INTO `goods_image` VALUES (53, 45, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637218361128479400.jpg', 0, 10, 1637218361, 1);
INSERT INTO `goods_image` VALUES (54, 42, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637218381814981600.jpg', 0, 10, 1637218382, 1);
INSERT INTO `goods_image` VALUES (55, 41, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637218395623952700.jpg', 0, 10, 1637218396, 1);
INSERT INTO `goods_image` VALUES (56, 40, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637218411828092500.png', 0, 10, 1637218412, 1);
INSERT INTO `goods_image` VALUES (57, 39, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637218432993290700.jpg', 0, 10, 1637218433, 1);
INSERT INTO `goods_image` VALUES (58, 38, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637218459159740500.jpg', 0, 10, 1637218459, 1);
INSERT INTO `goods_image` VALUES (59, 36, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637218489222313500.jpg', 0, 10, 1637218489, 1);
INSERT INTO `goods_image` VALUES (60, 26, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637218536145683500.jpg', 0, 10, 1637218536, 1);
INSERT INTO `goods_image` VALUES (61, 25, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637218579047173100.jpg', 0, 10, 1637218579, 1);
INSERT INTO `goods_image` VALUES (62, 24, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637218693021098600.jpg', 0, 10, 1637218693, 1);
INSERT INTO `goods_image` VALUES (63, 23, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637218736074295700.jpg', 0, 10, 1637218736, 1);
INSERT INTO `goods_image` VALUES (64, 22, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637218748569814100.jpg', 0, 10, 1637218749, 1);
INSERT INTO `goods_image` VALUES (65, 21, 'static/uploaded/uploadedGoodsImages/2024/9/4/1637218768855028300.jpg', 0, 10, 1637218769, 1);
INSERT INTO `goods_image` VALUES (66, 19, 'static/uploaded/uploadedGoodsImages/2024/9/4/20240904104005-d009b3de9c82d15891ad27a0800a19d8bc3e42f7.jpg', 0, 100, 1725417607, 1);
INSERT INTO `goods_image` VALUES (67, 20, 'static/uploaded/uploadedGoodsImages/2024/9/4/20240904104013-d009b3de9c82d15891ad27a0800a19d8bc3e42f7.jpg', 0, 100, 1725417614, 1);
INSERT INTO `goods_image` VALUES (68, 21, 'static/uploaded/uploadedGoodsImages/2024/9/4/20240904104022-d009b3de9c82d15891ad27a0800a19d8bc3e42f7.jpg', 0, 100, 1725417624, 1);
INSERT INTO `goods_image` VALUES (69, 23, 'static/uploaded/uploadedGoodsImages/2024/9/4/20240904104031-d009b3de9c82d15891ad27a0800a19d8bc3e42f7.jpg', 0, 100, 1725417632, 1);
INSERT INTO `goods_image` VALUES (70, 22, 'static/uploaded/uploadedGoodsImages/2024/9/4/20240904104038-d009b3de9c82d15891ad27a0800a19d8bc3e42f7.jpg', 0, 100, 1725417640, 1);
INSERT INTO `goods_image` VALUES (71, 24, 'static/uploaded/uploadedGoodsImages/2024/9/4/20240904104109-d009b3de9c82d15891ad27a0800a19d8bc3e42f7.jpg', 0, 100, 1725417671, 1);
INSERT INTO `goods_image` VALUES (72, 25, 'static/uploaded/uploadedGoodsImages/2024/9/4/20240904104116-d009b3de9c82d15891ad27a0800a19d8bc3e42f7.jpg', 0, 100, 1725417677, 1);
INSERT INTO `goods_image` VALUES (73, 26, 'static/uploaded/uploadedGoodsImages/2024/9/4/20240904104122-d009b3de9c82d15891ad27a0800a19d8bc3e42f7.jpg', 0, 100, 1725417683, 1);
INSERT INTO `goods_image` VALUES (74, 36, 'static/uploaded/uploadedGoodsImages/2024/9/4/20240904104127-d009b3de9c82d15891ad27a0800a19d8bc3e42f7.jpg', 0, 100, 1725417689, 1);
INSERT INTO `goods_image` VALUES (75, 37, 'static/uploaded/uploadedGoodsImages/2024/9/4/20240904104134-d009b3de9c82d15891ad27a0800a19d8bc3e42f7.jpg', 0, 100, 1725417696, 1);
INSERT INTO `goods_image` VALUES (76, 42, 'static/uploaded/uploadedGoodsImages/2024/9/4/20240904104140-d009b3de9c82d15891ad27a0800a19d8bc3e42f7.jpg', 0, 100, 1725417701, 1);
INSERT INTO `goods_image` VALUES (77, 41, 'static/uploaded/uploadedGoodsImages/2024/9/4/20240904104146-d009b3de9c82d15891ad27a0800a19d8bc3e42f7.jpg', 0, 100, 1725417708, 1);
INSERT INTO `goods_image` VALUES (78, 40, 'static/uploaded/uploadedGoodsImages/2024/9/4/20240904104152-d009b3de9c82d15891ad27a0800a19d8bc3e42f7.jpg', 0, 100, 1725417714, 1);
INSERT INTO `goods_image` VALUES (79, 39, 'static/uploaded/uploadedGoodsImages/2024/9/4/20240904104158-d009b3de9c82d15891ad27a0800a19d8bc3e42f7.jpg', 0, 100, 1725417719, 1);
INSERT INTO `goods_image` VALUES (80, 38, 'static/uploaded/uploadedGoodsImages/2024/9/4/20240904104205-d009b3de9c82d15891ad27a0800a19d8bc3e42f7.jpg', 0, 100, 1725417726, 1);

-- ----------------------------
-- Table structure for goods_type
-- ----------------------------
DROP TABLE IF EXISTS `goods_type`;
CREATE TABLE `goods_type`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` int NOT NULL,
  `add_time` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of goods_type
-- ----------------------------
INSERT INTO `goods_type` VALUES (1, '手机', '手机22', 1, 1725377619);
INSERT INTO `goods_type` VALUES (2, '电脑', '电脑', 1, 1725377619);
INSERT INTO `goods_type` VALUES (3, '笔记本', '笔记本', 1, 1725377619);
INSERT INTO `goods_type` VALUES (4, '路由器', '路由器', 1, 1725377619);
INSERT INTO `goods_type` VALUES (9, '衣服', '衣服', 1, 1725377619);

-- ----------------------------
-- Table structure for goods_type_attribute
-- ----------------------------
DROP TABLE IF EXISTS `goods_type_attribute`;
CREATE TABLE `goods_type_attribute`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `type_id` int NOT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `attribute_type` int NOT NULL,
  `attribute_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` int NOT NULL,
  `sort` int NOT NULL,
  `add_time` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of goods_type_attribute
-- ----------------------------
INSERT INTO `goods_type_attribute` VALUES (1, 1, '基本信息', 1, '', 1, 10, 1725377619);
INSERT INTO `goods_type_attribute` VALUES (2, 2, '主体', 3, '111\r\n1111', 1, 19, 1725377619);
INSERT INTO `goods_type_attribute` VALUES (3, 2, '内存', 1, '', 1, 10, 1725377619);
INSERT INTO `goods_type_attribute` VALUES (4, 2, '硬盘', 1, '', 1, 10, 1725377619);
INSERT INTO `goods_type_attribute` VALUES (5, 2, '显示器', 1, '', 1, 111, 1725377619);
INSERT INTO `goods_type_attribute` VALUES (6, 2, '支持蓝牙', 3, '是\r\n否', 1, 1011, 1725377619);
INSERT INTO `goods_type_attribute` VALUES (7, 1, '性能	', 2, '', 1, 111, 1725377619);
INSERT INTO `goods_type_attribute` VALUES (8, 1, '相机', 2, '', 1, 0, 1725377619);
INSERT INTO `goods_type_attribute` VALUES (9, 1, '支持蓝牙', 3, '是\r\n否', 1, 0, 1725377619);
INSERT INTO `goods_type_attribute` VALUES (10, 4, '是否支持蓝牙', 3, '是\r\n否', 1, 1022, 1725377619);
INSERT INTO `goods_type_attribute` VALUES (12, 3, '尺寸1', 1, '', 1, 10, 1725377619);
INSERT INTO `goods_type_attribute` VALUES (14, 1, '内存与容量', 2, '', 1, 10, 1725377619);
INSERT INTO `goods_type_attribute` VALUES (15, 1, '外观尺寸', 2, '', 1, 10, 1725377619);
INSERT INTO `goods_type_attribute` VALUES (16, 1, '充电与电池', 2, '', 1, 10, 1725377619);
INSERT INTO `goods_type_attribute` VALUES (17, 1, '影像系统', 1, '', 1, 10, 1725377619);
INSERT INTO `goods_type_attribute` VALUES (18, 1, '传感器', 1, '', 1, 10, 1725377619);
INSERT INTO `goods_type_attribute` VALUES (19, 3, '支持蓝牙', 3, '是\r\n否', 1, 10, 1725377619);

-- ----------------------------
-- Table structure for nav
-- ----------------------------
DROP TABLE IF EXISTS `nav`;
CREATE TABLE `nav`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `position` tinyint NOT NULL,
  `is_opennew` tinyint NOT NULL,
  `relation` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `sort` int NOT NULL,
  `status` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of nav
-- ----------------------------
INSERT INTO `nav` VALUES (1, '小米商城', 'https://www.baidu.com', 2, 0, '19 20 22 23 24 53 36 38', 10, 1);
INSERT INTO `nav` VALUES (2, 'MIUI', 'https://www.baidu.com', 1, 1, '1', 10, 1);
INSERT INTO `nav` VALUES (3, '小米手机', 'https://www.baidu.com', 2, 0, '19 20 53 54', 10, 1);
INSERT INTO `nav` VALUES (4, '小米电视', 'https://www.baidu.com', 2, 0, '23 24', 10, 1);
INSERT INTO `nav` VALUES (5, '路由器', 'https://www.baidu.com', 2, 1, '25', 10, 1);
INSERT INTO `nav` VALUES (8, '云服务', 'https://www.baidu.com', 1, 0, '2', 10, 1);
INSERT INTO `nav` VALUES (9, '金融', 'https://www.baidu.com', 1, 0, '1', 10, 1);
INSERT INTO `nav` VALUES (10, '有品', 'https://www.baidu.com', 1, 0, '1', 10, 1);
INSERT INTO `nav` VALUES (11, '家电', 'https://www.baidu.com', 2, 1, '1', 10, 1);
INSERT INTO `nav` VALUES (12, '智能电视', 'https://www.baidu.com', 2, 1, '23', 10, 1);
INSERT INTO `nav` VALUES (14, '小米帮助中心2', 'https://www.baidu.com', 3, 0, '12 13 14', 10, 1);

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL,
  `order_sn` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `total_price` float NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `pay_status` tinyint NOT NULL,
  `pay_type` tinyint NOT NULL,
  `order_status` tinyint NOT NULL,
  `add_time` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO `order` VALUES (12, 5, '202409040628395', 4499, 'dddd', '15936030516', '22', 0, 0, 1, 1725402519);

-- ----------------------------
-- Table structure for order_item
-- ----------------------------
DROP TABLE IF EXISTS `order_item`;
CREATE TABLE `order_item`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` int NOT NULL,
  `goods_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `goods_id` int NOT NULL,
  `goods_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `goods_price` float NOT NULL,
  `add_time` int NOT NULL,
  `goods_num` int NOT NULL,
  `goods_version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `goods_color` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `uid` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of order_item
-- ----------------------------
INSERT INTO `order_item` VALUES (7, 12, 'RedmiBook 13 全面屏', 24, 'static/uploaded/uploadedGoodsImages/2024/9/4/1592820244.jpg', 4499, 1725402519, 1, '8G+128G', '金色', 5);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `status` tinyint NOT NULL,
  `add_time` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (9, '超级管理员', 'aaaaaaaa超级管理员', 1, 1725377619);
INSERT INTO `role` VALUES (14, '软件部门', '软件部门', 1, 1725377619);
INSERT INTO `role` VALUES (16, '销售部门', '销售部门', 1, 1725377619);

-- ----------------------------
-- Table structure for role_access
-- ----------------------------
DROP TABLE IF EXISTS `role_access`;
CREATE TABLE `role_access`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `access_id` int NOT NULL,
  `role_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 119 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of role_access
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `last_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` tinyint NOT NULL,
  `add_time` int NOT NULL,
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (5, 'e10adc3949ba59abbe56e057f20f883e', '15936030516', '127.0.0.1', '3839003861@qq.com', 1, 1725402425, 'mztkn');

SET FOREIGN_KEY_CHECKS = 1;
