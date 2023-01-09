-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Waktu pembuatan: 09 Jan 2023 pada 01.57
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
-- Database: `service_course`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `chapters`
--

CREATE TABLE `chapters` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `course_id` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `chapters`
--

INSERT INTO `chapters` (`id`, `name`, `course_id`, `created_at`, `updated_at`) VALUES
(2, 'Pendahuluan', 1, '2022-12-25 22:45:07', '2022-12-25 22:45:07');

-- --------------------------------------------------------

--
-- Struktur dari tabel `courses`
--

CREATE TABLE `courses` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `certificate` int(11) NOT NULL,
  `thumbnail` varchar(255) NOT NULL,
  `type` varchar(255) NOT NULL,
  `status` varchar(255) NOT NULL,
  `price` int(11) NOT NULL,
  `level` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `mentor_id` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `courses`
--

INSERT INTO `courses` (`id`, `name`, `certificate`, `thumbnail`, `type`, `status`, `price`, `level`, `description`, `mentor_id`, `created_at`, `updated_at`) VALUES
(1, 'Belajar Microservices Untuk Beginner', 1, ' ', 'Premium', 'draft', 0, 'Pemula', 'kelas untuk mempelajarin arsitektur microservices bagi pemula', 2, '0000-00-00 00:00:00', '2022-12-12 18:17:39'),
(2, 'Belajar Microservices Untuk Beginner', 1, ' ', 'Premium', 'draft', 0, 'Pemula', 'kelas untuk mempelajarin arsitektur microservices bagi pemula', 2, '0000-00-00 00:00:00', '2022-12-12 18:18:42'),
(3, 'Belajar Microservices Untuk Pemula', 1, ' ', 'Premium', 'draft', 0, 'Pemula', 'kelas untuk mempelajarin arsitektur microservices bagi pemula', 2, '2022-09-14 14:19:05', '2022-09-14 14:19:05'),
(4, 'Belajar Microservices Untuk Pemula', 1, ' ', 'Premium', 'draft', 0, 'Pemula', 'kelas untuk mempelajarin arsitektur microservices bagi pemula', 2, '2022-09-14 14:23:30', '2022-09-14 14:23:30'),
(5, 'Belajar MVC Untuk Pemula', 1, ' ', 'Premium', 'draft', 0, 'Pemula', 'kelas untuk mempelajarin arsitektur microservices bagi pemula', 1, '2022-09-14 18:01:42', '2022-09-14 18:01:42'),
(6, 'Belajar MVC Untuk Pemula', 1, ' ', 'Premium', 'draft', 0, 'Pemula', 'kelas untuk mempelajarin arsitektur microservices bagi pemula', 1, '2023-01-04 11:01:58', '2023-01-04 11:01:58'),
(7, 'kelas jaringan', 1, ' ', 'free', 'publish', 0, 'Pemula', 'kelas untuk mempelajari tentang jaringan dan keamanan', 1, '2023-01-04 11:03:15', '2023-01-04 11:03:15');

-- --------------------------------------------------------

--
-- Struktur dari tabel `image_courses`
--

CREATE TABLE `image_courses` (
  `id` int(11) NOT NULL,
  `course_id` int(11) NOT NULL,
  `image` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Struktur dari tabel `lessons`
--

CREATE TABLE `lessons` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `video` varchar(255) NOT NULL,
  `chapter_id` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `lessons`
--

INSERT INTO `lessons` (`id`, `name`, `video`, `chapter_id`, `created_at`, `updated_at`) VALUES
(1, 'Introduction', 'www.youtube.com', 2, '2022-12-25 22:45:37', '2022-12-25 22:45:37'),
(2, 'Requirement', 'www.youtube.com', 2, '2022-12-25 22:48:50', '2022-12-25 22:48:50');

-- --------------------------------------------------------

--
-- Struktur dari tabel `mentors`
--

CREATE TABLE `mentors` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `profile` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `profession` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `mentors`
--

INSERT INTO `mentors` (`id`, `name`, `profile`, `email`, `profession`, `created_at`, `updated_at`) VALUES
(1, 'Tarto Rizaldi S.kom', 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/49/Jefri_Nichol_in_2019.png/800px-Jefri_Nichol_in_2019.png', 'tartorizaldi@gmail.com', 'Backend Programmer', '2022-09-05 15:17:03', '2022-09-05 15:17:03'),
(2, 'Muhammad Ilham Kusumawardhana S.kom', 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/49/Jefri_Nichol_in_2019.png/800px-Jefri_Nichol_in_2019.png', 'Kawekaweha123@gmail.com', 'Backend Programmer', '2022-09-05 15:18:31', '2022-09-05 15:18:31'),
(3, 'Dhymas Wahyu S.kom', 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/49/Jefri_Nichol_in_2019.png/800px-Jefri_Nichol_in_2019.png', 'dws@gmail.com', 'Backend Programmer', '2022-09-05 16:01:46', '2022-09-05 16:01:46'),
(4, 'Ozzi Ze Dong', 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/49/Jefri_Nichol_in_2019.png/800px-Jefri_Nichol_in_2019.png', 'ozzigans@gmail.com', 'Petinju', '0000-00-00 00:00:00', '0000-00-00 00:00:00');

-- --------------------------------------------------------

--
-- Struktur dari tabel `my_courses`
--

CREATE TABLE `my_courses` (
  `id` int(11) NOT NULL,
  `course_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `my_courses`
--

INSERT INTO `my_courses` (`id`, `course_id`, `user_id`, `created_at`, `updated_at`) VALUES
(1, 1, 1, '2022-12-26 23:36:15', '2022-12-26 23:36:15');

-- --------------------------------------------------------

--
-- Struktur dari tabel `reviews`
--

CREATE TABLE `reviews` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `course_id` int(11) NOT NULL,
  `rating` int(11) NOT NULL,
  `note` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `chapters`
--
ALTER TABLE `chapters`
  ADD PRIMARY KEY (`id`),
  ADD KEY `course_id` (`course_id`);

--
-- Indeks untuk tabel `courses`
--
ALTER TABLE `courses`
  ADD PRIMARY KEY (`id`),
  ADD KEY `mentor_id` (`mentor_id`);

--
-- Indeks untuk tabel `image_courses`
--
ALTER TABLE `image_courses`
  ADD PRIMARY KEY (`id`),
  ADD KEY `course_id` (`course_id`);

--
-- Indeks untuk tabel `lessons`
--
ALTER TABLE `lessons`
  ADD PRIMARY KEY (`id`),
  ADD KEY `chapter_id` (`chapter_id`);

--
-- Indeks untuk tabel `mentors`
--
ALTER TABLE `mentors`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `my_courses`
--
ALTER TABLE `my_courses`
  ADD PRIMARY KEY (`id`),
  ADD KEY `course_id` (`course_id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indeks untuk tabel `reviews`
--
ALTER TABLE `reviews`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `course_id` (`course_id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `chapters`
--
ALTER TABLE `chapters`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `courses`
--
ALTER TABLE `courses`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT untuk tabel `image_courses`
--
ALTER TABLE `image_courses`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `lessons`
--
ALTER TABLE `lessons`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `mentors`
--
ALTER TABLE `mentors`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `my_courses`
--
ALTER TABLE `my_courses`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT untuk tabel `reviews`
--
ALTER TABLE `reviews`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `chapters`
--
ALTER TABLE `chapters`
  ADD CONSTRAINT `chapters_ibfk_1` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON UPDATE CASCADE;

--
-- Ketidakleluasaan untuk tabel `courses`
--
ALTER TABLE `courses`
  ADD CONSTRAINT `courses_ibfk_1` FOREIGN KEY (`mentor_id`) REFERENCES `mentors` (`id`) ON UPDATE CASCADE;

--
-- Ketidakleluasaan untuk tabel `image_courses`
--
ALTER TABLE `image_courses`
  ADD CONSTRAINT `image_courses_ibfk_1` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON UPDATE CASCADE;

--
-- Ketidakleluasaan untuk tabel `lessons`
--
ALTER TABLE `lessons`
  ADD CONSTRAINT `lessons_ibfk_1` FOREIGN KEY (`chapter_id`) REFERENCES `chapters` (`id`) ON UPDATE CASCADE;

--
-- Ketidakleluasaan untuk tabel `my_courses`
--
ALTER TABLE `my_courses`
  ADD CONSTRAINT `my_courses_ibfk_1` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON UPDATE CASCADE;

--
-- Ketidakleluasaan untuk tabel `reviews`
--
ALTER TABLE `reviews`
  ADD CONSTRAINT `reviews_ibfk_1` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
