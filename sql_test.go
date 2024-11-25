package belajargolangdatabase

import (
	"context"
	"log"
	"testing"
)

/*
Eksekusi Perintah SQL
● Saat membuat aplikasi menggunakan database, sudah pasti kita ingin berkomunikasi dengan
database menggunakan perintah SQL.
● Di Golang juga menyediakan function yang bisa kita gunakan untuk mengirim
perintah SQL ke database menggunakan function (DB) ExecContext(context, sql, params).
● Ketika mengirim perintah SQL, kita butuh mengirimkan context, dan seperti yang
sudah pernah kita pelajari di course Golang Context, dengan context, kita bisa
mengirim sinyal cancel jika kita ingin membatalkan pengiriman perintah SQL nya.
*/
func TestExecSql(t *testing.T) {
	/// create connection
	db := GetConnection()

	/// close db , after done.
	defer db.Close()

	/// create new ctx
	ctx := context.Background()

	/// insert data to db -> table customer
	script := "INSERT INTO customer (id, name) VALUES ('fitra', 'Fitra');"
	_, err := db.ExecContext(ctx, script)

	/// check apakah ada error
	if err != nil {
		panic(err)
	}

	log.Println("Berhasil memasukan data")
}

/// ============================================================

/*
/ Query SQL
● Untuk operasi SQL yang tidak membutuhkan hasil, kita bisa menggunakan perintah Exec,
namun jika kita membutuhkan result, seperti SELECT SQL, kita bisa menggunakan function
yang berbeda.
● Function untuk melakukan query ke database, bisa menggunakan function (DB)
QueryContext(context, sql, params).

/ Rows
● Hasil Query function adalah sebuah data structs sql.Rows.
● Rows digunakan untuk melakukan iterasi terhadap hasil dari query.
● Kita bisa menggunakan function (Rows) Next() (boolean) untuk melakukan iterasi
terhadap data hasil query, jika return data false, artinya sudah tidak ada data lagi
didalam result.
● Untuk membaca tiap data, kita bisa menggunakan (Rows) Scan(columns...)
● Dan jangan lupa, setelah menggunakan Rows, jangan lupa untuk menutupnya menggunakan
(Rows) Close().
*/

func TestQuerySql(t *testing.T) {
	/// create connection
	db := GetConnection()

	/// close db , after done.
	defer db.Close()

	/// create new ctx
	ctx := context.Background()

	/// insert data to db -> table customer
	script := "SELECT * FROM customer;"
	rows, err := db.QueryContext(ctx, script)

	/// check apakah ada error
	if err != nil {
		panic(err)
	}

	/// iterasi pada rows
	for rows.Next() {
		/// true -> hasData
		/// false -> no more data.

		var id, name string

		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		log.Printf("id: %v\n", id)
		log.Printf("name: %v\n", name)
	}

	/// tutup rows , jika sudah selesai mengambil data.
	defer rows.Close()

	log.Println("Berhasil membaca data")
}
