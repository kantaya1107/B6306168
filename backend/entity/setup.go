package entity

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Doctor{}, &Type{}, &Evidence{}, &Leave{},
	)

	db = database
	//--- ข้อมูลแพทย์---//
	Phonsak := Doctor{
		DoctorName:  "Phonsak songsang",
		DoctorEmail: "Phonsak@gmail.com",
	}
	db.Model(&Doctor{}).Create(&Phonsak)
	Hanoi := Doctor{
		DoctorName:  "Hanoi slotmachine",
		DoctorEmail: "Hanoi@gmail.com",
	}
	db.Model(&Doctor{}).Create(&Hanoi)
	// db.Model(&Doctor{}).Create(&Doctor{
	// 	DoctorName:  "Phonsak songsang",
	// 	DoctorEmail: "Phonsak@gmail.com",
	// })
	// db.Model(&Doctor{}).Create(&Doctor{
	// 	DoctorName:  "Hanoi slotmachine",
	// 	DoctorEmail: "Hanoi@gmail.com",
	// })

	//--- ประเภทการลา ---//
	Personal := Type{
		Tleave: "Personal leave",
	}
	db.Model(&Type{}).Create(&Personal)

	Sick := Type{
		Tleave: "Sick leave",
	}
	db.Model(&Type{}).Create(&Sick)

	Maternity := Type{
		Tleave: "Maternity leave",
	}
	db.Model(&Type{}).Create(&Maternity)

	//--- ชนิดหลักฐาน ---//
	Document := Evidence{
		Etype: "Document",
	}
	db.Model(&Evidence{}).Create(&Document)

	Medicalcertificate := Evidence{
		Etype: "Medical certificate",
	}
	db.Model(&Evidence{}).Create(&Medicalcertificate)

	ImageVideo := Evidence{
		Etype: "Image/Video",
	}
	db.Model(&Evidence{}).Create(&ImageVideo)
	// db.Model(&Evidence{}).Create(&Evidence{
	// 	etype: "Document",
	// })
	// db.Model(&Evidence{}).Create(&Evidence{
	// 	etype: "Medical certificate",
	// })
	// db.Model(&Evidence{}).Create(&Evidence{
	// 	etype: "Image/Video",
	// })

	// var Phonsak Doctor
	// var Hanoi Doctor
	// db.Raw("SELECT * FROM doctor WHERE doctoremail = ?", "Phonsak@gmail.com").Scan(&Phonsak)
	// db.Raw("SELECT * FROM doctor WHERE doctoremail = ?", "Hanoi@gmail.com").Scan(&Hanoi)

	//leave data1
	db.Model(&Leave{}).Create(&Leave{
		Reason:   "to run errands in the provinces",
		Fdate:    time.Date(2022, 1, 1, 0, 0, 0, 0, time.Now().Location()),
		Ldate:    time.Date(2022, 1, 5, 0, 0, 0, 0, time.Now().Location()),
		Doctor:   Phonsak,
		Type:     Personal,
		Evidence: Document,
	})
	//leave data2
	db.Model(&Leave{}).Create(&Leave{
		Reason:   "COVID-infected",
		Fdate:    time.Date(2022, 2, 1, 0, 0, 0, 0, time.Now().Location()),
		Ldate:    time.Date(2022, 2, 16, 0, 0, 0, 0, time.Now().Location()),
		Doctor:   Hanoi,
		Type:     Sick,
		Evidence: Medicalcertificate,
	})
}
