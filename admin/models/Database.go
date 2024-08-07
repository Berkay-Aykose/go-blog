package models

var dsn string = "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

/*
GORM 'MySQL' DB BAĞLANTISI

Bağlantı dizesi açıklaması:
- user: Veritabanı kullanıcı adı
- pass: Veritabanı şifresi
- 127.0.0.1:3306: Veritabanı sunucusunun IP adresi ve portu
- dbname: Kullanılacak veritabanının adı
- charset=utf8mb4: Karakter seti
- parseTime=True: Tarih ve saat verilerini parse etmek için
- loc=Local: Yerel zaman dilimini kullanmak için
*/
