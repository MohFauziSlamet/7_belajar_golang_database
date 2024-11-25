package belajargolangdatabase

import (
	"database/sql"
	"time"
)

/*
/ Database Pooling
● sql.DB di Golang sebenarnya bukanlah sebuah koneksi ke database.
● Melainkan sebuah pool ke database, atau dikenal dengan konsep
Database Pooling.
● Di dalam sql.DB, Golang melakukan management koneksi ke database secara
otomatis. Hal ini menjadikan kita tidak perlu melakukan management koneksi
database secara manual.
● Dengan kemampuan database pooling ini, kita bisa menentukan jumlah minimal
dan maksimal koneksi yang dibuat oleh Golang, sehingga tidak membanjiri
koneksi ke database, karena biasanya ada batas maksimal koneksi yang bisa
ditangani oleh database yang kita gunakan.

*/

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/belajar_golang_database")

	if err != nil {
		panic(err)
	}

	/*
	   /Pengaturan Database Pooling
	   	Method
	   	(DB) SetMaxIdleConns(number) : Pengaturan berapa jumlah koneksi minimal
	   	yang dibuat.
	   	(DB) SetMaxOpenConns(number) : Pengaturan berapa jumlah koneksi
	   	maksimal yang dibuat.
	   	(DB) SetConnMaxIdleTime(duration) Pengaturan berapa lama koneksi yang
	   	sudah tidak digunakan akan dihapus.
	   	(DB) SetConnMaxLifetime(duration) Pengaturan berapa lama koneksi boleh
	   	digunakan.

	*/
	db.SetMaxIdleConns(50)
	db.SetMaxIdleConns(1000)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
