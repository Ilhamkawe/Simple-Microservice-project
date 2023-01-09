-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Waktu pembuatan: 09 Jan 2023 pada 01.59
-- Versi server: 5.7.24
-- Versi PHP: 7.4.15

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `service_user`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `refresh_tokens`
--

CREATE TABLE `refresh_tokens` (
  `id` int(11) NOT NULL,
  `token` text NOT NULL,
  `user_id` int(11) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `refresh_tokens`
--

INSERT INTO `refresh_tokens` (`id`, `token`, `user_id`, `created_at`, `updated_at`) VALUES
(1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImlkIjoxLCJuYW1lIjoiTXVoYW1tYWQgSWxoYW0gS3VzdW1hd2FyZGhhbmEiLCJlbWFpbCI6Imthd2VrYXdlaGEzMjFAZ21haWwuY29tIiwicm9sZSI6ImFkbWluIiwicHJvZmVzc2lvbiI6IkZ1bGxzdGFjayBEZXZlbG9wZXIifSwiaWF0IjoxNjcxMTk3Njk5LCJleHAiOjE2NzEyODQwOTl9.y6XKHRsI3lkcCO_9LIH1y6cVCGWdIN2Ro-nyIfoekzE', 1, '2022-12-16 13:34:59', '2022-12-16 13:34:59');

-- --------------------------------------------------------

--
-- Struktur dari tabel `sequelizemeta`
--

CREATE TABLE `sequelizemeta` (
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data untuk tabel `sequelizemeta`
--

INSERT INTO `sequelizemeta` (`name`) VALUES
('20220813032009-create_table_users.js'),
('20220813032233-create_table_refresh_tokens.js');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `profession` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `role` enum('admin','student') NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `profession`, `avatar`, `role`, `email`, `password`, `created_at`, `updated_at`) VALUES
(1, 'Muhammad Ilham Kusumawardhana', 'Fullstack Developer', NULL, 'admin', 'kawekaweha321@gmail.com', '$2b$10$9BpkZN6ZFNYmzBQ8Bcbjne0jlkvZW5jBVeX6s6lIIQc.pMyg3c1P.', '2022-08-14 06:31:11', '2022-08-14 06:31:11'),
(2, 'Muhamamd Adityo Kusumawardhana', 'Mahasiswa', NULL, 'student', 'kawekaweha123@gmail.com', '$2b$10$nGXBySk67eoWgE./VmmqBO3f79BKMKRO.bognPzZFU9iLEAfDqb9O', '2022-08-14 06:31:11', '2022-08-14 06:31:11'),
(3, 'Muhammad Ilham Kusumawardhana', 'FullStack Developer Jr.', 'https://image.com', 'student', 'kawekaweha00@gmail.com', '$2b$10$voStNRw8Gajd.rwXpIQ8dOuE50MnkP8YAEW.zD5UC.IB4GI9XJvTO', '2022-08-19 07:38:15', '2022-09-04 09:22:57'),
(4, 'Muhammad Ilham Kusumawardhana', 'FullStack Developer', NULL, 'student', 'kawekaweha987@gmail.com', '$2b$10$LV0pcuOD2VukEPA8OUm5ZeaSVvuTpXdLuIaRAk0m/wlWjtHEIUUrK', '2022-09-03 02:25:23', '2022-09-03 02:25:23'),
(5, 'Muhammad Ilham Kusumawardhana', 'FullStack Developer', NULL, 'student', 'kawekaweha666@gmail.com', '$2b$10$kZJku4C6/X89I.L6xLyZUOhck6IXNPQayZ7LxyGrn.34D1HP2qfj2', '2023-01-04 04:42:06', '2023-01-04 04:42:06'),
(6, 'ozzi', 'teacher', NULL, 'student', 'ozz1212@mail.com', '$2b$10$TtZ2qVyVcKo.brWhq/Uz3uE4x1bGRFaTBSVaRLRqZz77OP2k5BBT.', '2023-01-04 04:46:11', '2023-01-04 04:46:11');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `refresh_tokens`
--
ALTER TABLE `refresh_tokens`
  ADD PRIMARY KEY (`id`),
  ADD KEY `REFRESH_TOKEN__USER_ID` (`user_id`);

--
-- Indeks untuk tabel `sequelizemeta`
--
ALTER TABLE `sequelizemeta`
  ADD PRIMARY KEY (`name`),
  ADD UNIQUE KEY `name` (`name`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `UNIQUE_USER_EMAIL` (`email`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `refresh_tokens`
--
ALTER TABLE `refresh_tokens`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `refresh_tokens`
--
ALTER TABLE `refresh_tokens`
  ADD CONSTRAINT `REFRESH_TOKEN__USER_ID` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
