# MyGram - Aplikasi Media Sosial untuk Berbagi Foto dan Komentar

MyGram adalah sebuah aplikasi media sosial yang memungkinkan pengguna untuk menyimpan foto serta membuat komentar untuk foto orang lain. Aplikasi ini dilengkapi dengan proses CRUD (Create, Read, Update, Delete) yang memungkinkan pengguna untuk berinteraksi dengan data secara penuh. Dalam pengembangan aplikasi ini, digunakan framework Gin Gonic dan ORM Gorm untuk mempermudah proses pengembangan.

## Daftar Table
Dalam proyek MyGram, terdapat empat table utama yang diperlukan. Berikut merupakan detailnya:

#### 1. User
```
- Field-field:
  - Email: Validasi format email yang valid, unique index, tidak boleh kosong.
  - Username: Unique index, tidak boleh kosong.
  - Password: Minimal panjang 6 karakter, tidak boleh kosong.
  - Age: Minimal nilai di atas 8, tidak boleh kosong.
```

#### 2. Photo
```
= Field-field:
  - Title: Tidak boleh kosong.
  - Photo_url: Tidak boleh kosong.
```

#### 3. Social Media
```
- Field-field:
  - Name: Tidak boleh kosong.
  - Social_media_url: Tidak boleh kosong.
```
#### 4. Comment
```
- Field-field:
  - Text: Tidak boleh kosong.
```

## Validasi
Setiap field pada keempat table tersebut harus melewati validasi sesuai dengan kriteria yang ditentukan. Validasi dapat dibuat sendiri atau menggunakan package seperti Go Validator.

## Alur Aplikasi
- Pengguna harus melewati proses autentikasi menggunakan Json Web Token (JWT) untuk mengakses data pada table Social Media, Photo, dan Comment.
- Untuk modifikasi data kepemilikan seperti Update atau Delete, pengguna harus melewati proses autorisasi.

Dengan MyGram, pengguna dapat dengan mudah berbagi foto, membuat komentar, dan berinteraksi dengan pengguna lainnya dalam platform media sosial yang aman dan terstruktur.

## Endpoint-endpoint yang ada
Berikut adalah daftar endpoint-endpoint yang disediakan dalam aplikasi MyGram:

#### 1. User
```
Register [POST]: Mendaftarkan pengguna baru.
Login [POST]: Proses masuk pengguna.
```

#### 2. Photo
```
GetAll [GET]: Mendapatkan daftar semua foto.
GetOne [GET]: Mendapatkan satu foto berdasarkan ID.
CreatePhoto [POST]: Membuat foto baru.
UpdatePhoto [PUT]: Memperbarui informasi foto.
DeletePhoto [DELETE]: Menghapus foto.
```

#### 3. Comment
```
GetAll [GET]: Mendapatkan daftar semua komentar.
GetOne [GET]: Mendapatkan satu komentar berdasarkan ID.
CreateComment [POST]: Membuat komentar baru.
UpdateComment [PUT]: Memperbarui komentar.
DeleteComment [DELETE]: Menghapus komentar.
```

#### 3. Social Media
```
GetAll [GET]: Mendapatkan daftar semua media sosial.
GetOne [GET]: Mendapatkan satu media sosial berdasarkan ID.
CreateSocialMedia [POST]: Membuat media sosial baru.
UpdateSocialMedia [PUT]: Memperbarui informasi media sosial.
DeleteSocialMedia [DELETE]: Menghapus media sosial.
```

Dengan menyediakan endpoint-endpoint ini, MyGram memungkinkan pengguna untuk melakukan berbagai operasi CRUD serta berinteraksi dengan data foto, komentar, pengguna, dan media sosial dengan mudah dan aman

#### Detail tutorial :  https://github.com/septyancandra/golang-final-project-hacktiv8/tree/main/MyGram
