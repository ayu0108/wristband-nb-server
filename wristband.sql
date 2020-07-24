-- phpMyAdmin SQL Dump
-- version 4.9.2
-- https://www.phpmyadmin.net/
--
-- 主機： 127.0.0.1
-- 產生時間： 
-- 伺服器版本： 10.4.11-MariaDB
-- PHP 版本： 7.4.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 資料庫： `wristband`
--

-- --------------------------------------------------------

--
-- 資料表結構 `device`
--

CREATE TABLE `device` (
  `id` int(16) NOT NULL,
  `device_id` varchar(255) CHARACTER SET utf8mb4 NOT NULL COMMENT '裝置ID',
  `password` text CHARACTER SET utf8mb4 NOT NULL COMMENT '裝置密碼',
  `status` tinyint(1) NOT NULL COMMENT '裝置狀態',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '創建時間',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '更新時間'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- 資料表結構 `receive_data`
--

CREATE TABLE `receive_data` (
  `id` int(128) NOT NULL,
  `device_id` varchar(255) CHARACTER SET utf8mb4 NOT NULL COMMENT '裝置ID',
  `ip_address` varchar(48) CHARACTER SET utf8mb4 NOT NULL COMMENT '裝置IP',
  `beacon_name` text CHARACTER SET utf8mb4 NOT NULL COMMENT '藍芽名稱',
  `beacon_data` text CHARACTER SET utf8mb4 NOT NULL COMMENT '藍芽資料',
  `serial_number` varchar(255) CHARACTER SET utf8mb4 NOT NULL COMMENT '序列號',
  `singal` tinyint(2) NOT NULL COMMENT '裝置訊號',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '接收時間',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '更新時間'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- 傾印資料表的資料 `receive_data`
--

INSERT INTO `receive_data` (`id`, `device_id`, `ip_address`, `beacon_name`, `beacon_data`, `serial_number`, `singal`, `created_at`, `updated_at`) VALUES
(31, 'wb-01', '[::1]:57826', 'beacon_name', 'beacon_data', '0', 68, '2020-07-24 18:19:57', '2020-07-24 18:19:57'),
(32, 'wb-02', '[::1]:58537', 'test_name', 'test_data', '0', 55, '2020-07-24 18:34:40', '2020-07-24 18:34:40'),
(35, 'wb-02', '[::1]:56521', 'test_name', 'test_data', '0', 68, '2020-07-24 20:19:48', '2020-07-24 20:19:48');

--
-- 已傾印資料表的索引
--

--
-- 資料表索引 `device`
--
ALTER TABLE `device`
  ADD PRIMARY KEY (`id`);

--
-- 資料表索引 `receive_data`
--
ALTER TABLE `receive_data`
  ADD PRIMARY KEY (`id`);

--
-- 在傾印的資料表使用自動遞增(AUTO_INCREMENT)
--

--
-- 使用資料表自動遞增(AUTO_INCREMENT) `device`
--
ALTER TABLE `device`
  MODIFY `id` int(16) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- 使用資料表自動遞增(AUTO_INCREMENT) `receive_data`
--
ALTER TABLE `receive_data`
  MODIFY `id` int(128) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=36;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
