package belajargolangdatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseTest(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/belajar_golang_database")

	if err != nil {
		t.Errorf("Gagal membuka koneksi ke basis data: %v", err)
	}

	/// tutup db , jika sudah selesai digunakan
	defer db.Close()

	/// gunakan db

}
