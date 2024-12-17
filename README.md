# Backend Relawanku - Capstone Project MSIB 7 Alterra Academy

## Deskripsi
Relawanku adalah API yang dirancang sebagai bentuk kampanye sosial guna mengajak masyarakat untuk saling membantu sesama dengan cara menjadi relawan untuk beberapa program sosial kemanusiaan yang tersedia. API ini menyediakan beberapa fitur interaktif bagi pengguna seperti melihat semua program relawan yang ada, melihat artikel terkait kegiatan sosial maupun lingkungan, mendaftar ke program yang disediakan,serta melakukan donasi untuk korban bencana maupun pelestarian lingkungan. API ini menyediakan berita serta informasi terkait kegiatan sosial lingkungan dengan harapan dapat menumbuhkan rasa kepedulian masyarakat terhadap kondisi sekitarnya lebih baik lagi.

## Spesifikasi Fitur Product
### Fitur Umum : 
- Melakukan login sebagai user ataupun sebagai admin menggunakan username dan password.
- Untuk user yang belum memiliki akun dapat melakukan regist terlebih dahulu.
- Menerapkan fitur "update password" dan "update user" untuk memudahkan user mengedit data pribadi nya.
- Menerapkan implementasi API Eksternal untuk membantu proses pembayaran atau proses donasi agar real time dan lebih interaktif.

### Admin : 
- Dapat menambah, mengedit, serta menghapus artikel.
- Dapat menambah, mengedit, serta menghapus "event donasi" berdasarkan kategori program yang ada.
- Dapat menambah, mengedit, serta menghapus program relawan yang disediakan sesuai kategori program.
- Dapat melihat dan menghapus data client atau user.
- Dapat memantau baik semua transaksi yang dilakukan oleh user ke "event donasi" termasuk jumlah dan status transaksi maupun menampilkan secara rinci transaksi yang telah dilakukan.

### User : 
- Dapat melihat semua artikel baik artikel tarbaru maupun artikel berdasarkan kategori social maupun lingkungan yang ada.
- Dapat melihat semua program relawan yang tersedia dan dapat melakukan pendaftaran ke program yang diinginkan.
- Dapat melihat semua program yang telah didaftarkan atau diikuti.
- Dapat melihat "event donasi" yang tersedia dan langsung melakukan transaksi donasi ke program yang diinginkan.
- Dapat mengedit data diri dan password.


## Tech Stack
1. App Framework :Echo Golang
2. ORM library : GORM
3. Database : MySQL, RDS
4. Deployment : EC2 AWS (Amazon Web Service)
5. Code Structure : Clean Architecture
6. Authentication : JWT
7. Containerization : Docker
8. Version Control : Git
9. Other Tools : Firebase untuk upload foto

## ERD
[ERD Relawanku](https://drive.google.com/file/d/13pWb7eQVOOutY5zNUfvojONACUlahN6r/view?usp=sharing)
![ERD Relawanku](https://imgur.com/a/WesB0zg)


## API Documentation
[Swagger](https://relawanku.xyz/swagger/index.html)


