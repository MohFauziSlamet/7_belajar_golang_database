package belajargolangdatabase

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseTest(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {

	// Ganti dengan hostname, username, password, dan nama database Anda
	hostname := "sql207.infinityfree.com"
	username := "if0_37928602"
	password := "Fauzi222606"
	dbname := "if0_37928602_belajar_golang_database"

	// Format DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		t.Errorf("Gagal membuka koneksi ke basis data: %v", err)
	}

	/// tutup db , jika sudah selesai digunakan
	defer db.Close()

	/// gunakan db

}
