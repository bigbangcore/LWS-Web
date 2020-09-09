/*
 Navicat Premium Data Transfer

 Source Server         : lws
 Source Server Type    : MySQL
 Source Server Version : 80020
 Source Host           : 192.168.0.101:3306
 Source Schema         : lws

 Target Server Type    : MySQL
 Target Server Version : 80020
 File Encoding         : 65001

 Date: 09/09/2020 17:00:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for block
-- ----------------------------
DROP TABLE IF EXISTS `block`;
CREATE TABLE `block`  (
  `id` int unsigned NOT NULL,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `hash` varbinary(255) NOT NULL,
  `version` int unsigned NULL,
  `block_type` int unsigned NULL,
  `prev` varbinary(255) NULL DEFAULT NULL,
  `tstamp` int unsigned NULL,
  `merkle` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `height` int unsigned NULL,
  `mint_tx_id` varbinary(255) NULL DEFAULT NULL,
  `sig` blob NULL,
  PRIMARY KEY (`id`, `hash`) USING BTREE,
  INDEX `idx_block_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for device
-- ----------------------------
DROP TABLE IF EXISTS `device`;
CREATE TABLE `device`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `address_id` int(0) NULL DEFAULT NULL,
  `device_id` varbinary(12) NULL DEFAULT NULL,
  `smec_id` varbinary(12) NULL DEFAULT NULL,
  `guard_id` varbinary(8) NULL DEFAULT NULL,
  `created_at` timestamp(4) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for migrations
-- ----------------------------
DROP TABLE IF EXISTS `migrations`;
CREATE TABLE `migrations`  (
  `id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for test
-- ----------------------------
DROP TABLE IF EXISTS `test`;
CREATE TABLE `test`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tx
-- ----------------------------
DROP TABLE IF EXISTS `tx`;
CREATE TABLE `tx`  (
  `id` int unsigned NOT NULL,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `hash` varbinary(32) NULL DEFAULT NULL,
  `version` int unsigned NULL,
  `tx_type` int unsigned NULL,
  `block_id` int(0) NULL DEFAULT NULL,
  `block_hash` varbinary(32) NULL DEFAULT NULL,
  `block_height` int unsigned NULL,
  `lock_until` int unsigned NULL,
  `inputs` mediumblob NULL,
  `amount` bigint(0) NULL DEFAULT NULL,
  `change` bigint(0) NULL DEFAULT NULL,
  `fee` bigint(0) NULL DEFAULT NULL,
  `send_to` varbinary(33) NULL DEFAULT NULL,
  `sender` varbinary(33) NULL DEFAULT NULL,
  `data` mediumblob NULL,
  `sig` mediumblob NULL,
  `forkid` varbinary(32) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `hash`(`hash`) USING BTREE,
  INDEX `idx_tx_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `forkid`(`forkid`) USING BTREE,
  INDEX `send_to`(`send_to`) USING BTREE,
  INDEX `sender`(`sender`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tx_pool
-- ----------------------------
DROP TABLE IF EXISTS `tx_pool`;
CREATE TABLE `tx_pool`  (
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `hash` varbinary(32) NOT NULL,
  `version` int unsigned NULL,
  `tx_type` int unsigned NULL,
  `block_height` int unsigned NULL,
  `lock_until` int unsigned NULL,
  `inputs` mediumblob NULL,
  `amount` bigint(0) NULL DEFAULT NULL,
  `change` bigint(0) NULL DEFAULT NULL,
  `fee` bigint(0) NULL DEFAULT NULL,
  `send_to` varbinary(33) NULL DEFAULT NULL,
  `sender` varbinary(33) NULL DEFAULT NULL,
  `data` mediumblob NULL,
  `sig` mediumblob NULL,
  `forkid` varbinary(32) NULL DEFAULT NULL,
  PRIMARY KEY (`hash`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `address_id` int unsigned NOT NULL,
  `address` varbinary(33) NULL DEFAULT NULL,
  `api_key` varbinary(32) NULL DEFAULT NULL,
  `topic_prefix` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `fork_num` tinyint unsigned NULL,
  `fork_list` blob NULL,
  `reply_utxon` int unsigned NULL,
  `time_stamp` int unsigned NULL,
  `nonce` int unsigned NULL,
  PRIMARY KEY (`address_id`) USING BTREE,
  INDEX `idx_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for utxo
-- ----------------------------
DROP TABLE IF EXISTS `utxo`;
CREATE TABLE `utxo`  (
  `id` int unsigned NOT NULL,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `tx_hash` varbinary(32) NOT NULL,
  `destination` varbinary(33) NULL DEFAULT NULL,
  `amount` bigint(0) NULL DEFAULT NULL,
  `forkid` varbinary(32) NULL DEFAULT NULL,
  `block_height` int unsigned NULL,
  `out` tinyint unsigned NULL,
  `idx` varbinary(33) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_utxo_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `utxo_idx_idx`(`idx`) USING BTREE,
  INDEX `forkid`(`forkid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for utxo_pool
-- ----------------------------
DROP TABLE IF EXISTS `utxo_pool`;
CREATE TABLE `utxo_pool`  (
  `id` int unsigned NOT NULL,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `tx_hash` varbinary(32) NOT NULL,
  `destination` varbinary(33) NULL DEFAULT NULL,
  `amount` bigint(0) NULL DEFAULT NULL,
  `forkid` varbinary(32) NULL DEFAULT NULL,
  `out` tinyint unsigned NULL,
  `idx` varbinary(33) NOT NULL,
  `is_delete` tinyint(1) NOT NULL,
  `tx_owner` varbinary(32) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_utxo_pool_idx`(`idx`) USING BTREE,
  INDEX `tx_owner`(`tx_owner`) USING BTREE,
  CONSTRAINT `utxo_pool_ibfk_1` FOREIGN KEY (`tx_owner`) REFERENCES `tx_pool` (`hash`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for web_auth
-- ----------------------------
DROP TABLE IF EXISTS `web_auth`;
CREATE TABLE `web_auth`  (
  `ID` int(0) NOT NULL AUTO_INCREMENT,
  `AuthName` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `AuthURL` varchar(800) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Icon` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `IsShow` int(0) NULL DEFAULT NULL,
  `Sort` int(0) NULL DEFAULT NULL,
  `State` int(0) NOT NULL DEFAULT 0,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for web_device
-- ----------------------------
DROP TABLE IF EXISTS `web_device`;
CREATE TABLE `web_device`  (
  `ID` bigint(0) NOT NULL AUTO_INCREMENT,
  `DeviceID` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `DeviceName` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `DeviceDes` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `DeviceRegTime` bigint(0) NULL DEFAULT NULL,
  `DeviceUpdateTime` bigint(0) NULL DEFAULT NULL,
  `DeviceStatus` int(0) NULL DEFAULT NULL,
  `KeyID` bigint(0) NOT NULL DEFAULT 0,
  `VersionID` bigint(0) NOT NULL DEFAULT 0,
  `UserID` int(0) NOT NULL DEFAULT 0,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of web_device
-- ----------------------------
INSERT INTO `web_device` VALUES (2, 'did yyy 123', 'name xxx zhongjin', '1234232', 1595582135, 1596619047, 0, 0, 5, 0);
INSERT INTO `web_device` VALUES (3, 'ID yyy 123', 'name zzz', 'des   hahahah  content', 1595582236, 1596619047, 0, 0, 5, 0);
INSERT INTO `web_device` VALUES (5, 'idxxxxxxxxxxxxx4444444', '33333333333', 'des xxxxxxx5555555555', 1595831307, 1596619047, 0, 0, 5, 0);
INSERT INTO `web_device` VALUES (7, 'test Device ID 003', 'test xxx 002', 'test Describe xxx content', 1595835001, 1596619047, 0, 58, 5, 0);
INSERT INTO `web_device` VALUES (8, 'ID 00002', 'name 000001', 'describe xxxx', 1595835223, 1596619000, 0, 63, 5, 0);
INSERT INTO `web_device` VALUES (10, 'ID  xxxx 00002', 'NNNNAME 0001', 'Descccccc x  hahahha', 1595837629, 1596676744, 1, 0, 0, 0);

-- ----------------------------
-- Table structure for web_devicekey
-- ----------------------------
DROP TABLE IF EXISTS `web_devicekey`;
CREATE TABLE `web_devicekey`  (
  `ID` bigint(0) NOT NULL AUTO_INCREMENT,
  `Pubkey` varchar(800) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Prikey` varchar(800) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `CreateTime` bigint(0) NULL DEFAULT NULL,
  `State` int(0) NULL DEFAULT NULL,
  `Des` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `UserID` int(0) NOT NULL DEFAULT 0,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 67 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of web_devicekey
-- ----------------------------
INSERT INTO `web_devicekey` VALUES (12, 'a8c0f6ae118b6287fc3bc7cccb5210af0635765d25780926134ef9b336367107', 'eff50e94f3f4541fa27c91cdfe2d434420aff7a8a62ba5d3b7659639e2da368b', 1596531015, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (13, 'a9c11a36b177738ce8b49dbbeedb67bf6d3a0b0fd7fc0f05faf11c15d3a14db8', '422a51a69ce737fd9973c22dd65ab1df8eae00bddb063cff0c133ef089da0cb7', 1596531015, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (14, '06837bedc3dc25dba61a7a98d8f4f4750771f41c024e7b82d174a720db68d8fc', '605180e3effb12e6f5a40d5305b1caae1b283d46b10b8f2b1affaf22226197bf', 1596531015, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (15, '812866ff81c649c9b694fe2c3f5d18e2d7ad7988581cf71d8927b74b82a6785a', 'e12f407ba565531f71af3586fa6f1eb412e229ddb5e985eb1abb96d274435da1', 1596531015, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (16, 'b7b2e81b096dc1b3d16a06ba1a68962b00c807f4f3d3f12b1016cb95efd7c63b', '2085604456a8c4802624ec7578a33e9994b6fe8abca022c98db1a095778a4a10', 1596531015, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (17, 'cdd4b30513fc0b54433ec9aa213eeeac3c159959a93a54c5fbfa0c0ad2d21dc9', 'b8ebb3c9c88a776496f729fd9404d7d44e2849fb219b5da2ddf7f2c43ffa15fc', 1596531039, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (18, 'd8ae79e6763905c9c2b4338436ff7d6c8751d8a8f117678dca1e55ccd3eace9c', '04595d7a747c8fcd2f9c95139adbb4870219bc3f7dc822b345834195a9b0fd67', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (19, '934314b73957e29433ce5e8a997170d9ea108ea04ef138cf9d3e3e82415b9cd8', '5db17c655ad788a2761d1bee21ecbd0fccef6efd25fe732bafa5ff25ce70a595', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (20, '7899a32ad40e2d0f9f99a54373ba6caef1229fafcd3c9c73cbd714498d6d1175', 'acff864c826127749f28c278adb29c717e6bf1d875469adf9e664412a361500c', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (21, '5ba2b71fd512fa70b2e261248ec554a9a977b37137b2542d4ebc8cd773bc6546', 'cc3ea2e013287088bfa52765a0e402bcf445bdf7644addd19dd6aa1d684a8fa0', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (22, 'dfb0e93dcee80d24246fe471fe9fa2f10531808a339c70454131ed63e395079b', '49c8032f4d70be8c7d05cd36cbafa4ffe4f8b61a78c47d9d6512f45543588fe7', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (23, '4933977aa8cfe5ccd03d0c967db51d156aaed841185cb681eb3ba9e1dae84857', 'abda88d04cd638f33e212bf76b37e455e43aec7c72d99056e9894b3b09308438', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (24, '7274ef3a8963dfdb6d8647d84309e96d1b1d55d0916ebccf39291cb8ca1fdd2b', '5a9f5fdb0f0ea600177968f8115c8337ead52bee3becec623cc8260e80360ed4', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (25, '2116ae1093dccb4d53c95d068310db41fdcc23672d6e676b9f471691443375f8', '5da0a8fde861426b9d981cc34e203868072a8cee6ca429195473e49227a1bc80', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (26, 'a4519d20af63a972b29bec30ae8ce8cef6be7273219dd727de9ea773be076778', '74949b63746f436c20a759ada6e5a2055f996241c04ee324caabc171550c4fee', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (27, '57afe03df0355d991ec7da23285105aad1730bcbded93dc5b6cba8b2dea3d869', 'e576b67d2dc80c31e8ff191873e5f026a8079ddd4cf0ec76f9664cfd2bbe9b3b', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (28, '40c8a23fac751ee41a88b97d83c7fdf8e6fc1d237e5a418ba53dca392c7b717d', 'bc0cb30d77a3f13169e7e7b08ae855cb8c077c54a253047541f4e1095c2dfc0e', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (29, 'c803ac801ee6dc2866e16a6eaeaac5b9cbf3fdff333b4a31adc24ce00c3c7f26', '28dc6ef1cc7991f6449dbea21a227826add210067f80c4fb3e02d233ba7d9681', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (30, '569a4993ffc17dc70a4b913571d28cbb5b818ab2a3967798135531ab416a1f72', 'deeecdbc26a06d6109c4e3acc957c077dd3f675ee418d461036cbc0a34950617', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (31, '777027e7bfa5c9f9b359da38903a1e737d2fde526b495e6b38fd0c5a0835c954', 'e9db31b6fb1204574921395f1eb3abdf49acb75da22486a92ed50f347480002d', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (32, 'b6d3cc7975dc3dbe4a9eb50505739922f1a8adbc1103f863d402bd76866e3bf2', '9bb50405a8b3100a27a5163ccd6b99443940f64f8bf6bf5c34fce75e354ba7c8', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (33, '8e161c223ee332ec4841a35f115f99fe87cddd937fe72431b3fcdf041686ca9e', '6a6f776b7895f8d7fd606e2e298567f8326e5a358af091d4b269c24ae53eed96', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (34, 'eab8f021aebacdf0062488bfc815b9c532f5630250777a5c9d78a67a1203f026', '5513d508ee291e9d70151a30b5b9e60b07bb333c02e2667a84d7ee7847c96649', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (35, '7c694083b72e13a13b5971f4bd550e7bd5c933814fb38b091ad48e0981bc7ea3', 'af03b09b2778c3c3f83efa6f2a44b8ed12ae6b51510c585a1846000a5e68ce60', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (36, 'd4957861d6eb94523f154a8e6ed9c1abe411262cfeb6772ee6ccb574742ff264', '69d7f99a47edd5bf4da8825a3ea407479439b68482295ba88f5286ce910671cd', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (37, '5f23ad6933b29869b960d42e1c2bd633bf05b78a1fbd17e3029f6f86c0679eaf', '67724987b27c0cf7d7e267ce8201ba1a8e71f85a63a3047ba2d389caae3afef2', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (38, '00fc34a688e475d28ca83fd39ba550b95c13d5dd972837a973ae7d0a9ccd9ccb', 'd539c5826565d0d1272d8a3b4b553775e3a58c2045d43c93e3934d9c9c56780d', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (39, 'b5e1f9cf8a38c5c96da9e96cfd466bc22cb3e0568dfc9ba2921cf36493b90645', '4c66dd2307a43e4c928903b48d33afe8c651e8ea44cebe8f7da3f2e373190b08', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (40, '085b364454a9c73776706afd13ee8c47a30818a4d6d4349a0206264cfab6a4ac', 'd09eed55aa34bd150a2a9e10b0b9255ec40005320aeb1ead461be5c378f3d25a', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (41, '3d38836863b37fca2960046c0bf78dff34dc04131de615ebfc4b290491f5f3fe', 'a8dc104a3031cfbbfbfd99434f1c6ad54fde6cdcf1b0f000948a509ab3187b0f', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (42, '62388301b7e06fcb7dab22ba7e2132693430aae449de0a7fe27c8076690a6648', '9b36ddda22befa27c44ff79473c44071d367458525a1a6ae4425da6d6250aa01', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (43, 'f778aae5acc6f1c34179cdb9022a442b5e3e7c112869fae72dc27e1dac7681b2', '9ddea950618a513cde5a248589af89531ac6116c6abe9263d839e37ec1b61647', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (44, 'a810b769d1cef4ef0dc8965ccd3a52ba8cb35b38976b5db6b07d2e183a268bad', '1a7e61be5ac58047f1bc019a2051dc1bed3f4957f14403ba2a19e55055b5ca71', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (45, 'f0559c4d7519fcef9dead8deb4d8b1fd1ec90de3e4156dd302f41979e9ebdee7', 'f09c9e21e991046c976fd78989fd65716ae5853d11df9f98e4933bbc2492a0b9', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (46, 'bf68d9f8a9ecc70df7413d0aa699025ad8da9115f5629885953a18b3740d1ecf', '01eef00a8c5e353ebd24039ce7dc632b060e4197ebab32fab7fd02f5a2abd989', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (47, '6b3bb6756e16710402446059601ff94e1500638eb12ac55497ff9a74617e2f9d', '0c3d0bb0490907ac48af278cb3e973b84c2a2847c15ef24c1545f89e6f4cd316', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (48, '8c89783988030dd5250f7193d0ef0a3a55f83199cf5f10cc12e0f678658d8b38', '4d407ecd53fcd817d74ffbc94500f489f445e125ecbd21f0e2a1d906ef2ae2df', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (49, 'e7f8b255830ff7302f916f8a488eb58f8cdaadad5fcd694dd8dbf1e8346f3ff2', 'cd5858c575145aa0a824e2788477fc6b34873f98b32487b4aea1b98e4ec83d2b', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (50, '2e3da49778534e5e795c69eef9fa3c11a950a3fb489058f77469018da44e505f', '11782bb9c012af15f76038df599c0d97bcfb4b1ef74dbc67fe73e4e4bcb3db66', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (51, '27dfd70d9313eddc5221a5a06f6892507026701c8bb5cc2b5136a826ee2ba583', '8e83a3ac416d102a9329cb41a7e0f13ce60d877598b2f1615234e48105ba2228', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (52, '7e3211fb558c67a2395d2548054ad709e6320d5e59dfba91816652f428b5f715', 'c9dca0969be0bdf46fe0326b55a17faa52af1c83d5bd0cfc2e2445abc3793872', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (53, 'add230bfcd08884e7be60f687e230542c9afbe490066d277f253055827f864f4', '5b17d1dbcd07a43594a35bd15f09234b14d4430f7fd15408964cf4b932421ccc', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (54, '969362131216d067b7c3c97f9bb93d2daee89fdc662642807e92fe22f6fc6a07', 'f439d2d6262f399c6a9673db8335dd67cde8e46f03a9e1fa0a8f75f346dcb30c', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (55, '84854e11f0aec3478b24409c73b824812f8c89da7a5891f465e9514e1cfd57d9', '3b461743f6c8077376039c6b770981c25b5ad1d7285d63de651b4d4ac35f6a45', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (56, '1d11de2ca1ec39a898042f922378a61859796e9146958c5081fc744e96602049', 'bf84762131f72f0ff730ba3854381a4aace3914cadfe5cd9e7c244d3213605be', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (57, '3c3db3b7f823c208806a65d348a5812d74dd0729b59f54f3f40416db8fce78fd', '198dc47472da7f0cfa68008055df8962e0e84d9475068ce7836f4d207296c2a8', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (58, '8e0f6ee76c586f1f93605088b6f0afe4890600a3589157db3948d67ee02eb2c6', '9fb1f9fd5a6b421a1622aff5a46d833d39a72f755900802d96436c62687dc51f', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (59, '377664a8b96f16ccaea28e8feb8a454eefb160cbc687ee4cbefab27f3d7e4438', '8d8e7ccef4d93118172138fdcbf5cb6c4e91a4c1213ea778c5cc22e475fd9e70', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (60, '11e586e0c04293f3928038e628375ecc0402d035207023d220fe8f3cc3bf27dd', '3757336b9d8fe0b7af7d9d3e3606a957111d965c7a9990f424996e2074a06eb8', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (61, '4faa3a046d947b0b3c5f9a915f424e7c27ebf29907413c8df44b4d1612c57ddd', 'e14eb1a1663395ff75c739c8ce302e13bd33a4949d345c7dbea439d6c5c5759b', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (62, '0e4243be97b4c3f5c8dd9929331ead7a05505dc57bd775087f8bf18339c4bc27', 'e791130992a2cec0fc815d7c2a2c70cac0899b50a6398c4616237c6f872299ec', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (63, 'd5ac80bd2d7a8a8f8489178900115b4f91321e1faeaef7cc5de4184c3343655d', '9688ce10a241afc8aa65e1a98c933bc26e8df887453f514bb80e38159dd93a9e', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (64, 'b9f50ae6a42e4753eed22ad05e0b9d099e7f346c70530fedcf411ad16441967b', '2c636c0e0b08b0433f2af6274071e3beb6b57c4005a6f9b98ce35fb78c927c2d', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (65, 'fe8634fa0eb3deb34ef35e5f477e450245765a3bce3327e0613a22229a154606', 'f66855e287c926f491f4de12b22f7a295a2d74ebf1896e70b1707fab9560677a', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (66, 'c09c4ba53ca5621982a7fe725c4fc37c63dfbd6e01491cc182bc89b3068f6a36', 'bc6834a96a77eb3434b0a85a31b0d5ab5d516d74c6d16faacd59773317ba0e1a', 1596531056, 1, 'create', 1);
INSERT INTO `web_devicekey` VALUES (67, 'd44a6c5a84dd1740fb33e750d57d2c2b79fe38762751f1006461a17a345dce2b', 'b4c2d1d08d49964628d80d4340ea4013a2482c5cc42473264152600dbba7606e', 1596531056, 1, 'create', 1);

-- ----------------------------
-- Table structure for web_group
-- ----------------------------
DROP TABLE IF EXISTS `web_group`;
CREATE TABLE `web_group`  (
  `ID` int(0) NOT NULL AUTO_INCREMENT,
  `PID` int(0) NOT NULL DEFAULT 0,
  `GroupName` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `GroupType` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `GroupDes` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `CreateTime` bigint(0) NOT NULL DEFAULT 0,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for web_grouprole
-- ----------------------------
DROP TABLE IF EXISTS `web_grouprole`;
CREATE TABLE `web_grouprole`  (
  `GroupID` int(0) NOT NULL,
  `RoleID` int(0) NOT NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for web_groupuser
-- ----------------------------
DROP TABLE IF EXISTS `web_groupuser`;
CREATE TABLE `web_groupuser`  (
  `GroupID` int(0) NOT NULL,
  `UserID` int(0) NOT NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for web_role
-- ----------------------------
DROP TABLE IF EXISTS `web_role`;
CREATE TABLE `web_role`  (
  `ID` int(0) NOT NULL AUTO_INCREMENT,
  `RoleName` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `IsShow` int(0) NOT NULL DEFAULT 0,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for web_user
-- ----------------------------
DROP TABLE IF EXISTS `web_user`;
CREATE TABLE `web_user`  (
  `ID` bigint(0) NOT NULL AUTO_INCREMENT,
  `LoginName` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `RealName` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `Pwd` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `RoleIds` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `Phone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `Email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `Salt` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `LastLogin` bigint(0) NOT NULL DEFAULT 0,
  `LastIP` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `State` int(0) NOT NULL DEFAULT 0,
  `CreateID` bigint(0) NOT NULL DEFAULT 0,
  `UpdateID` bigint(0) NOT NULL DEFAULT 0,
  `CreateTime` bigint(0) NOT NULL DEFAULT 0,
  `UpdateTime` bigint(0) NOT NULL DEFAULT 0,
  `UserDes` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of web_user
-- ----------------------------
INSERT INTO `web_user` VALUES (1, 'zhongjin', '', 'e10adc3949ba59abbe56e057f20f883e', '', '', 'zhongjin-19@163.com', '', 0, '', 1, 0, 0, 1596180622, 1596180623, '');

-- ----------------------------
-- Table structure for web_versions
-- ----------------------------
DROP TABLE IF EXISTS `web_versions`;
CREATE TABLE `web_versions`  (
  `ID` bigint(0) NOT NULL AUTO_INCREMENT,
  `VersionTitle` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `VersionAddTime` bigint(0) NULL DEFAULT 0,
  `VersionDes` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `VersionURL` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `VersionUpdateTime` bigint(0) NULL DEFAULT 0,
  `VersionSize` double NULL DEFAULT 0,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of web_versions
-- ----------------------------
INSERT INTO `web_versions` VALUES (5, 'Version v0.0.0.1', 1596533040, 'this is a test version', 'static/upload/20200804172312.7z', 1596533006, 3622274);

SET FOREIGN_KEY_CHECKS = 1;
