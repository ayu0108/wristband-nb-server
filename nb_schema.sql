-- phpMyAdmin SQL Dump
-- version 4.7.4
-- https://www.phpmyadmin.net/
--
-- 主機: 127.0.0.1
-- 產生時間： 2020-05-22 06:51:05
-- 伺服器版本: 10.1.30-MariaDB
-- PHP 版本： 5.6.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 資料庫： `nb_schema`
--

-- --------------------------------------------------------

--
-- 資料表結構 `device`
--

CREATE TABLE `device` (
  `id` int(16) NOT NULL,
  `account` varchar(24) NOT NULL,
  `password` text NOT NULL,
  `name` varchar(255) NOT NULL,
  `mac` varchar(48) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- 資料表的匯出資料 `device`
--

INSERT INTO `device` (`id`, `account`, `password`, `name`, `mac`) VALUES
(10, 'user', 'e10adc3949ba59abbe56e057f20f883e', 'JinWei', '00-15-AF-5A-F8-42');

-- --------------------------------------------------------

--
-- 資料表結構 `recive`
--

CREATE TABLE `recive` (
  `id` int(16) NOT NULL,
  `name` varchar(255) NOT NULL,
  `addr` varchar(48) NOT NULL,
  `mac` varchar(48) DEFAULT NULL,
  `distance` varchar(255) NOT NULL,
  `temperature` varchar(255) NOT NULL,
  `humidity` varchar(255) NOT NULL,
  `date` varchar(255) NOT NULL,
  `connStatus` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- 資料表的匯出資料 `recive`
--

INSERT INTO `recive` (`id`, `name`, `addr`, `mac`, `distance`, `temperature`, `humidity`, `date`, `connStatus`) VALUES
(2, 'JinWei', '223.141.30.253:22764', '00-15-AF-5A-F8-42', '-80', '32.20', '62.00', '20200519_18:09:15', 1),
(7, 'defaultID', '[::1]:62223', '', '-60', '31.20', '56.00', '20200520_01:09:15', 1);

--
-- 已匯出資料表的索引
--

--
-- 資料表索引 `device`
--
ALTER TABLE `device`
  ADD PRIMARY KEY (`id`);

--
-- 資料表索引 `recive`
--
ALTER TABLE `recive`
  ADD PRIMARY KEY (`id`);

--
-- 在匯出的資料表使用 AUTO_INCREMENT
--

--
-- 使用資料表 AUTO_INCREMENT `device`
--
ALTER TABLE `device`
  MODIFY `id` int(16) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- 使用資料表 AUTO_INCREMENT `recive`
--
ALTER TABLE `recive`
  MODIFY `id` int(16) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
