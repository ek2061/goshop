-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- 主機： 127.0.0.1
-- 產生時間： 2021-05-15 09:47:14
-- 伺服器版本： 10.4.14-MariaDB
-- PHP 版本： 7.4.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 資料庫： `shop`
--

-- --------------------------------------------------------

--
-- 資料表結構 `catalogs`
--

CREATE TABLE `catalogs` (
  `catalog_id` int(11) NOT NULL,
  `name` varchar(20) NOT NULL,
  `hiden` tinyint(1) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 傾印資料表的資料 `catalogs`
--

INSERT INTO `catalogs` (`catalog_id`, `name`, `hiden`) VALUES
(1, '家電', 0),
(2, '3C', 0),
(3, '食品', 0);

-- --------------------------------------------------------

--
-- 資料表結構 `products`
--

CREATE TABLE `products` (
  `product_id` int(11) NOT NULL,
  `catalog_id` int(11) NOT NULL,
  `name` varchar(20) NOT NULL,
  `cost` int(30) NOT NULL,
  `price` int(30) NOT NULL,
  `description` text NOT NULL,
  `on_sale` tinyint(1) NOT NULL DEFAULT 0,
  `start_sell` datetime NOT NULL,
  `end_sell` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 傾印資料表的資料 `products`
--

INSERT INTO `products` (`product_id`, `catalog_id`, `name`, `cost`, `price`, `description`, `on_sale`, `start_sell`, `end_sell`) VALUES
(1, 3, '日本新潟越光米', 120, 385, '防疫期間，宅家就能在自家餐桌享用美味\r\n給你高CP值的日本越光米\r\n這價格你還不買爆?你還在猶豫什麼??', 0, '2021-05-01 00:00:00', '2021-05-31 00:00:00'),
(2, 3, 'Baan泰式綠咖哩拌麵', 150, 499, '★ Baan團隊用心研發 米其林星級名店\r\n★ 2款口味可選 泰式綠咖哩/酸辣麵\r\n★ 泰式綠咖哩口味 椰奶滑順的口感濃郁\r\n★ 數十種蔬菜及辛料 拌炒製成的美味醬汁\r\n★ 嚴選小麥關廟麵 中厚外薄超吸汁\r\n★ 5分鐘輕鬆上菜 星級美食輕鬆完成！', 1, '2021-05-03 10:00:00', '2021-05-16 00:00:00'),
(3, 1, '品諾 14吋DC馬達遙控立扇DF-140', 500, 1080, '■ 無線遙控器\r\n■ DC直流馬達設計\r\n■ 七段可調式風量\r\n■ 五片式扇葉，風量更強勁\r\n■ 台灣製造，品質有保障\r\n■ X寶、東X一線品牌同廠品質', 1, '2021-05-02 08:00:00', '2021-05-26 08:00:00'),
(4, 1, 'Sony BRAVIA 43吋 4K G', 12000, 22900, '► 4K HDR 超極真影像處理器 X1\r\n► 原色顯示PRO\r\n► 極致平衡揚聲器\r\n► 支援杜比視界、杜比全景聲\r\n► Google TV: 全新使用者介面，容易操作，迎接無限串流娛樂', 1, '2021-05-12 00:00:00', '2021-05-31 00:00:00'),
(5, 2, 'Kingston 8GB DDR4 26', 600, 1299, '★桌上型電腦專用\r\n★僅適用第8代以上 CPU\r\n★規格：DDR4-2666\r\n★容量：8GB\r\n★電壓：1.2V', 1, '2021-05-01 12:00:00', '2021-05-31 00:00:00'),
(18, 3, '米', 1, 2, '很好吃', 1, '2021-05-01 00:00:00', '2021-05-25 00:00:00');

--
-- 已傾印資料表的索引
--

--
-- 資料表索引 `catalogs`
--
ALTER TABLE `catalogs`
  ADD PRIMARY KEY (`catalog_id`);

--
-- 資料表索引 `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`product_id`) USING BTREE,
  ADD KEY `catalog_id` (`catalog_id`);

--
-- 在傾印的資料表使用自動遞增(AUTO_INCREMENT)
--

--
-- 使用資料表自動遞增(AUTO_INCREMENT) `catalogs`
--
ALTER TABLE `catalogs`
  MODIFY `catalog_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- 使用資料表自動遞增(AUTO_INCREMENT) `products`
--
ALTER TABLE `products`
  MODIFY `product_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=40;

--
-- 已傾印資料表的限制式
--

--
-- 資料表的限制式 `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`catalog_id`) REFERENCES `catalogs` (`catalog_id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
