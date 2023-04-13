package seed

import (
	"backend/config"
	"backend/internal/schema"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/jaswdr/faker"
	"gorm.io/gorm"
)

type Province struct {
	ID   int    `json:"id"`
	Name string `json:"nama"`
}

type District struct {
	ID         int    `json:"id"`
	ProvinceID string `json:"id_provinsi"`
	Name       string `json:"nama"`
}

type ProvinceData struct {
	Provinces []Province `json:"provinsi"`
}

type DistrictData struct {
	Districts []District `json:"kota_kabupaten"`
}

type SubDistrict struct {
	ID     int    `json:"id"`
	CityID string `json:"id_kota"`
	Name   string `json:"nama"`
}

type SubDistrictData struct {
	SubDistricts []SubDistrict `json:"kecamatan"`
}

func SeedStore(db *gorm.DB) {
	bucketAddress := config.C.BucketAddress

	rand.Seed(time.Now().UnixNano())
	var seedStore []schema.Store

	// Create a map to store unique company names
	uniqueCompanyNames := make(map[string]bool)

	for len(seedStore) < 9 {
		name := faker.New().Company().Name()

		// Check if the company name already exists in the map
		if _, exists := uniqueCompanyNames[name]; !exists {
			uniqueCompanyNames[name] = true

			street := faker.New().Address().StreetName()
			provinces, _ := fetchProvinces()
			randProvince := provinces[rand.Intn(len(provinces))]
			districts, _ := fetchDistrict(randProvince.ID)
			randDistrict := districts[rand.Intn(len(districts))]
			subdistricts, _ := fetchSubDistrict(randDistrict.ID)
			randSubDistrict := subdistricts[rand.Intn(len(subdistricts))]

			seedStore = append(seedStore, schema.Store{
				Name:        name,
				Street:      street,
				Province:    randProvince.Name,
				District:    randDistrict.Name,
				Subdistrict: randSubDistrict.Name,
				Image:       bucketAddress + "assets/store/default.png",
			})
		}
	}

	tx := db.Begin()

	for _, store := range seedStore {
		if err := tx.Create(&store).Error; err != nil {
			tx.Rollback()
			return
		}
	}

	tx.Commit()
}

func fetchProvinces() ([]Province, error) {
	url := "https://dev.farizdotid.com/api/daerahindonesia/provinsi"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data ProvinceData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Provinces, nil
}

func fetchDistrict(provinceID int) ([]District, error) {
	url := fmt.Sprintf("https://dev.farizdotid.com/api/daerahindonesia/kota?id_provinsi=%d", provinceID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data DistrictData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Districts, nil
}

func fetchSubDistrict(districtID int) ([]SubDistrict, error) {
	url := fmt.Sprintf("https://dev.farizdotid.com/api/daerahindonesia/kecamatan?id_kota=%d", districtID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data SubDistrictData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.SubDistricts, nil
}
