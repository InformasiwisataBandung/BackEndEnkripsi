# Update enkripsi Password Mongo DB, Deploy function Signup, Memasukan Token jika user berhasil login kedalam Cookies
#### •Update 21-10-2023

```
{
  "username": "ucup",
  "password": "$2a$10$r.Z8w/WHkd7uHcE6ZGlqCOcsNQEQOdXyrYYcDMMY9V4/HLOmXloCq"
}
```
#### Update 23-10-2023
-API clooud functions Signup
```
https://asia-southeast2-bustling-walker-340203.cloudfunctions.net/function-Signup
```
input
```
{
  "username": "username",
  "password": "Password"
}
```
Send Post
```
{
    "message": "Pendaftaran berhasil"
}
```
-Memasukan Token ke Cookies
```
cookie := http.Cookie{
		Name:     "token",     // Nama cookie
		Value:    tokenString, // Token sebagai nilai cookie
		HttpOnly: true,        // Hanya bisa diakses melalui HTTP
		Path:     "/",         // Path di mana cookie berlaku (misalnya, seluruh situs)
		MaxAge:   3600,        // Durasi cookie (dalam detik), sesuaikan sesuai kebutuhan
		// Secure: true, // Jika situs dijalankan melalui HTTPS
	}
```

