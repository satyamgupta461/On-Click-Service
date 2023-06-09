package app

import (
	"github.com/rsj-rishabh/urbanClapClone/server/app/model"
)

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func (a *App) DBMigrate() {
	// Drop the table if it exists
	a.DB.AutoMigrate().DropTable(&model.User{})
	a.DB.AutoMigrate().DropTable(&model.Service{})
	a.DB.AutoMigrate().DropTable(&model.Booking{})
	a.DB.AutoMigrate().DropTable(&model.CityServiceMapping{})

	// Migrate the schema
	a.DB.AutoMigrate(&model.User{}, &model.Service{}, &model.Booking{}, &model.CityServiceMapping{})

	// Create users table
	a.DB.Create(&model.User{
		Id:       1,
		Name:     "Dummy Duck",
		Username: "dummy",
		Password: "dumdum",
		Email:    "dummy@abes.edu",
		Gender:   "M",
	})

	a.DB.Create(&model.User{
		Id:       2,
		Name:     "Buzz Lightyear",
		Username: "buzz",
		Password: "busybee",
		Email:    "buzz@abes.edu",
		Gender:   "M",
	})
	a.DB.Create(&model.User{
		Id:       3,
		Name:     "Snow White",
		Username: "snow",
		Password: "abc1234",
		Email:    "snow@abes.edu",
		Gender:   "F",
	})
	a.DB.Create(&model.User{
		Id:       4,
		Name:     "Satyam Gupta",
		Username: "Satyam123",
		Password: "Satyam123",
		Email:    "satyam@abes.edu",
		Gender:   "M",
	})
	a.DB.Create(&model.User{
		Id:       5,
		Name:     "Chirag Singh",
		Username: "Chirag123",
		Password: "Chirag123",
		Email:    "chirag@abes.edu",
		Gender:   "M",
	})
	a.DB.Create(&model.User{
		Id:       6,
		Name:     "Shivam Singh",
		Username: "Shivam123",
		Password: "Shivam123",
		Email:    "shivam@abes.edu",
		Gender:   "M",
	})
	a.DB.Create(&model.User{
		Id:       7,
		Name:     "Kunal Chauhan",
		Username: "Kunal123",
		Password: "Kunal123",
		Email:    "kunal@abes.edu",
		Gender:   "M",
	})

	// Create services table
	a.DB.Create(&model.Service{
		Id:          1,
		Name:        "AC Maintanence",
		Description: "Any type of AC maintanence such as filter cleaning, part replacement, etc.",
		Category:    "Electronics",
		ImageName:   "air_conditioning.jpg",
		Price:       80,
	})
	a.DB.Create(&model.Service{
		Id:          2,
		Name:        "Plumbing",
		Description: "Sanitary and household plumbing. No sewage service.",
		Category:    "Household",
		ImageName:   "plumbing.jpg",
		Price:       100,
	})
	a.DB.Create(&model.Service{
		Id:          3,
		Name:        "Saloon",
		Description: "Haricut, massage, nailwork, makeup, etc.",
		Category:    "Personal Care",
		ImageName:   "saloon.jpg",
		Price:       25,
	})
	a.DB.Create(&model.Service{
		Id:          4,
		Name:        "Furniture Repair",
		Description: "Furniture frame repair, drilling, fitting new furniture, etc.",
		Category:    "Household",
		ImageName:   "furniture_repair.jpg",
		Price:       70,
	})
	a.DB.Create(&model.Service{
		Id:          5,
		Name:        "Exterminator",
		Description: "Pest control, wildlife evac, alligator emergency, etc.",
		Category:    "Animal/Pet",
		ImageName:   "pest_control.jpg",
		Price:       150,
	})

	// Create booking table
	a.DB.Create((&model.Booking{
		UserId:      1,
		ServiceId:   1,
		Date:        "2022-02-15",
		StartTime:   "12:30",
		EndTime:     "13:30",
		IsCancelled: false,
	}))
	a.DB.Create((&model.Booking{
		UserId:      1,
		ServiceId:   2,
		Date:        "2022-02-15",
		StartTime:   "16:30",
		EndTime:     "17:30",
		IsCancelled: false,
	}))
	a.DB.Create((&model.Booking{
		UserId:      2,
		ServiceId:   3,
		Date:        "2022-02-15",
		StartTime:   "16:30",
		EndTime:     "17:30",
		IsCancelled: false,
	}))

	// Create CityServiceMapping table
	a.DB.Create((&model.CityServiceMapping{
		CityName:  "Delhi",
		ServiceId: 3,
	}))
	a.DB.Create((&model.CityServiceMapping{
		CityName:  "Delhi",
		ServiceId: 2,
	}))
	a.DB.Create((&model.CityServiceMapping{
		CityName:  "Mumbai",
		ServiceId: 2,
	}))
	a.DB.Create((&model.CityServiceMapping{
		CityName:  "Mumbai",
		ServiceId: 3,
	}))
	a.DB.Create((&model.CityServiceMapping{
		CityName:  "Noida",
		ServiceId: 1,
	}))
	a.DB.Create((&model.CityServiceMapping{
		CityName:  "Noida",
		ServiceId: 2,
	}))
}
