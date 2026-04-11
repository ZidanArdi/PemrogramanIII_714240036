package repository_test

import (
	"fmt"
	"time"

	"be_latihan/config"
	"be_latihan/model"
	"be_latihan/repository"
	"testing"

	"gorm.io/gorm/clause"
)

func seedMahasiswa(t *testing.T, npm int64) *model.Mahasiswa {
	t.Helper()

	mhs := &model.Mahasiswa{
		NPM:    npm,
		Nama:   "Test User",
		Prodi:  "Teknik Informatika",
		Alamat: "Jl. Test No. 123",
		NoHP:   "081234567890",
		Email:  "test.user@example.com",
		Hobi:   []string{"Membaca", "Coding"},
	}

	if err := config.GetDB().Clauses(clause.OnConflict{UpdateAll: true}).Create(mhs).Error; err != nil {
		t.Fatalf("Gagal menyiapkan data mahasiswa: %v", err)
	}

	return mhs
}

func setupTest(t *testing.T ) {
	config.InitDB()

	// Bersihkan data sebelum setiap pengujian
	err := config.GetDB().AutoMigrate(&model.Mahasiswa{})
	if err != nil {
		t.Fatalf("Gagal melakukan migrasi database: %v", err)
	}
}	

func TestInsertMahasiswa(t *testing.T) {
	setupTest(t)	
	
	npm := time.Now().Unix() // Gunakan timestamp sebagai NPM unik

	mhs := &model.Mahasiswa{
		NPM:    npm,
		Nama:   "Test User",
		Prodi:  "Teknik Informatika",
		Alamat: "Jl. Test No. 123",
		NoHP:   "081234567890",
		Email:  "test.user@example.com",
		Hobi:   []string{"Membaca", "Coding"},
	}
	
	_, err := repository.InsertMahasiswa(mhs)
	if err != nil {
		t.Fatalf("Gagal memasukkan mahasiswa: %v", err)
	}
	fmt.Printf("InsertMahasiswa berhasil")
}

func TestGetAllMahasiswa(t *testing.T) {
	setupTest(t)
	
	data, err := repository.GetAllMahasiswa()
	if err != nil {
		t.Fatalf("Gagal mendapatkan semua mahasiswa: %v", err)
	}

	if len(data) == 0 {
		t.Fatalf("Data mahasiswa kosong, seharusnya ada data")
	}
}

func TestGetMahasiswaByNPM(t *testing.T) {
	setupTest(t)	

	npm := int64(1775465952)
	seedMahasiswa(t, npm)

	mhs, err := repository.GetMahasiswaByNPM(npm)
	if err != nil {
		t.Errorf("Gagal mendapatkan mahasiswa dengan NPM : %v", err)
	}

	if mhs.NPM != npm {
		t.Errorf("Expected NPM %d, got %d", npm, mhs.NPM)
	}
	fmt.Printf("Data Ditemukan : %+v\n", mhs)
}	

func TestUpdateMahasiswa(t *testing.T) {
	setupTest(t)

	npm := int64(1775465952)
	seedMahasiswa(t, npm)

	_, err := repository.UpdateMahasiswa(npm, &model.Mahasiswa{
		Nama:   "Updated User",
		Prodi:  "Sistem Informasi",
		Alamat: "Jl. Updated No. 456",
		NoHP:   "089876543210",
		Email:  "updated.user@example.com",
	})

	if err != nil {
		t.Fatalf("Gagal memperbarui mahasiswa: %v", err)
	}
}	

func TestDeleteMahasiswa(t *testing.T) {
	setupTest(t)

	npm := int64(1775465952)
	seedMahasiswa(t, npm)

	err := repository.DeleteMahasiswa(npm)
	if err != nil {
		t.Fatalf("Gagal menghapus mahasiswa: %v", err)
	}
}