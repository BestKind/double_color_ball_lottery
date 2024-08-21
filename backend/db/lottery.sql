CREATE TABLE `lottery` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `version` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '期数',
  `open_time` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '开奖时间',
  `week` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '周几开奖',
  `red_1` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '红球1',
  `red_2` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '红球2',
  `red_3` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '红球3',
  `red_4` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '红球4',
  `red_5` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '红球5',
  `red_6` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '红球6',
  `blue` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '篮球',
  `l1_val` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '一等奖金额',
  `l1_cnt` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '一等奖数量',
  `l2_val` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '二等奖金额',
  `l2_cnt` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '二等奖数量',
  `l3_cnt` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '三等奖数量',
  `total` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '总奖池',
  `sale` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '销售额',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 历史同期排序
SELECT blue, COUNT(1)
FROM lottery WHERE version IN ('2024097','2023097','2022097','2021097','2020097','2019097','2018097','2017097','2016097','2015097','2014097','2013097') 
GROUP BY blue ORDER BY COUNT(1) DESC;

-- 固定前三位为每位出现次数最多的数字
SELECT red_4, count(1) FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 = '14' GROUP BY red_4 ORDER BY count(1) DESC;

SELECT red_5, count(1) FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 = '14' AND red_4 = '22' GROUP BY red_5 ORDER BY count(1) DESC;
SELECT red_5, count(1) FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 = '14' AND red_4 = '24' GROUP BY red_5 ORDER BY count(1) DESC;

SELECT red_6, count(1) FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 = '14' AND red_4 = '22' AND red_5 = '25' GROUP BY red_6 ORDER BY count(1) DESC;
SELECT red_6, count(1) FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 = '14' AND red_4 = '24' AND red_5 = '28' GROUP BY red_6 ORDER BY count(1) DESC;

SELECT blue, count(1) FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 = '14' AND red_4 = '22' AND red_5 = '25' AND red_6 = '26' GROUP BY blue ORDER BY count(1) DESC;
SELECT blue, count(1) FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 = '14' AND red_4 = '24' AND red_5 = '28' AND red_6 = '32' GROUP BY blue ORDER BY count(1) DESC;
--  12

-- 固定前两位为每位出现次数最多的数字
SELECT red_3, count(1) FROM lottery WHERE red_1 = '01' AND red_2 = '06' GROUP BY red_3 ORDER BY count(1) DESC;

SELECT red_4, count(1) FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 = '12' GROUP BY red_4 ORDER BY count(1) DESC;
SELECT red_4, COUNT(1) FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 IN (
	SELECT red_3 FROM lottery WHERE red_1 = '01' AND red_2 = '06'
)  GROUP BY red_4 ORDER BY COUNT(1) DESC;

SELECT red_5, count(1) FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 = '12' AND red_4 = '13' GROUP BY red_5 ORDER BY count(1) DESC;
SELECT red_5, COUNT(1) FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 IN (SELECT red_3 FROM lottery WHERE red_1 = '01' AND red_2 = '06') AND red_4 IN (
	SELECT red_4 FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 IN ('12', '07', '09', '11', '08', '17', '14', '13', '18', '27', '19', '15')
) GROUP BY red_5 ORDER BY COUNT(1) DESC;

SELECT red_6, count(1) FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 = '12' AND red_4 = '13' AND red_5 = '15' GROUP BY red_6 ORDER BY count(1) DESC;
SELECT red_6, COUNT(1) FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 IN (SELECT red_3 FROM lottery WHERE red_1 = '01' AND red_2 = '06') AND red_4 IN (
	SELECT red_4 FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 IN ('12', '07', '09', '11', '08', '17', '14', '13', '18', '27', '19', '15')
) AND red_5 IN (
	SELECT red_5 FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 IN (SELECT red_3 FROM lottery WHERE red_1 = '01' AND red_2 = '06') AND red_4 IN (
		SELECT red_4 FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 IN ('12', '07', '09', '11', '08', '17', '14', '13', '18', '27', '19', '15')
	)
) GROUP BY red_6 ORDER BY COUNT(1) DESC;

SELECT blue, count(1) FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 = '12' AND red_4 = '13' AND red_5 = '15' AND red_6 = '24' GROUP BY blue ORDER BY count(1) DESC;
SELECT blue, COUNT(1) FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 IN (SELECT red_3 FROM lottery WHERE red_1 = '01' AND red_2 = '06') AND red_4 IN (
	SELECT red_4 FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 IN ('12', '07', '09', '11', '08', '17', '14', '13', '18', '27', '19', '15')
) AND red_5 IN (
	SELECT red_5 FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 IN (SELECT red_3 FROM lottery WHERE red_1 = '01' AND red_2 = '06') AND red_4 IN (
		SELECT red_4 FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 IN ('12', '07', '09', '11', '08', '17', '14', '13', '18', '27', '19', '15')
	)
	AND red_6 IN (
		SELECT red_6 FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 IN (SELECT red_3 FROM lottery WHERE red_1 = '01' AND red_2 = '06') AND red_4 IN (
			SELECT red_4 FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 IN ('12', '07', '09', '11', '08', '17', '14', '13', '18', '27', '19', '15')
		) AND red_5 IN (
			SELECT red_5 FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 IN (SELECT red_3 FROM lottery WHERE red_1 = '01' AND red_2 = '06') AND red_4 IN (
				SELECT red_4 FROM lottery WHERE red_1 = '01' AND red_2 = '06' AND red_3 IN ('12', '07', '09', '11', '08', '17', '14', '13', '18', '27', '19', '15')
			)
		)
	)
) GROUP BY blue ORDER BY COUNT(1) DESC;

-- 查询每个位置的数字出现的次数(总数)
SELECT red_1, COUNT(1) FROM lottery GROUP BY red_1 ORDER BY count(1) DESC;
SELECT red_2, COUNT(1) FROM lottery GROUP BY red_2 ORDER BY count(1) DESC;
SELECT red_3, COUNT(1) FROM lottery GROUP BY red_3 ORDER BY count(1) DESC;
SELECT red_4, COUNT(1) FROM lottery GROUP BY red_4 ORDER BY count(1) DESC;
SELECT red_5, COUNT(1) FROM lottery GROUP BY red_5 ORDER BY count(1) DESC;
SELECT red_6, COUNT(1) FROM lottery GROUP BY red_6 ORDER BY count(1) DESC;
SELECT blue, COUNT(1) FROM lottery GROUP BY blue ORDER BY count(1) DESC;

-- 查询每个位置的数字出现的次数(2024)
SELECT red_1, COUNT(1) FROM lottery WHERE id >= 1657 GROUP BY red_1 ORDER BY count(1) DESC;
SELECT red_2, COUNT(1) FROM lottery WHERE id >= 1657 GROUP BY red_2 ORDER BY count(1) DESC;
SELECT red_3, COUNT(1) FROM lottery WHERE id >= 1657 GROUP BY red_3 ORDER BY count(1) DESC;
SELECT red_4, COUNT(1) FROM lottery WHERE id >= 1657 GROUP BY red_4 ORDER BY count(1) DESC;
SELECT red_5, COUNT(1) FROM lottery WHERE id >= 1657 GROUP BY red_5 ORDER BY count(1) DESC;
SELECT red_6, COUNT(1) FROM lottery WHERE id >= 1657 GROUP BY red_6 ORDER BY count(1) DESC;
SELECT blue, COUNT(1) FROM lottery WHERE id >= 1657 GROUP BY blue ORDER BY count(1) DESC;

-- 查询每个位置的数字出现的次数(2023)
SELECT red_1, COUNT(1) FROM lottery WHERE id >= 1506 AND id < 1657 GROUP BY red_1 ORDER BY count(1) DESC;
SELECT red_2, COUNT(1) FROM lottery WHERE id >= 1506 AND id < 1657 GROUP BY red_2 ORDER BY count(1) DESC;
SELECT red_3, COUNT(1) FROM lottery WHERE id >= 1506 AND id < 1657 GROUP BY red_3 ORDER BY count(1) DESC;
SELECT red_4, COUNT(1) FROM lottery WHERE id >= 1506 AND id < 1657 GROUP BY red_4 ORDER BY count(1) DESC;
SELECT red_5, COUNT(1) FROM lottery WHERE id >= 1506 AND id < 1657 GROUP BY red_5 ORDER BY count(1) DESC;
SELECT red_6, COUNT(1) FROM lottery WHERE id >= 1506 AND id < 1657 GROUP BY red_6 ORDER BY count(1) DESC;
SELECT blue, COUNT(1) FROM lottery WHERE id >= 1506 AND id < 1657 GROUP BY blue ORDER BY count(1) DESC;