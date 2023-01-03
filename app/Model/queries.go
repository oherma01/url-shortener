package model

import (
	"fmt"
)

// get all a short URLs, return error if any error occurs in the process of populating the struct
func GetAllShortened() ([]ShortURL, error) {

	var shortened []ShortURL

	tx := db.Find(&shortened)

	if tx.Error != nil {
		return []ShortURL{}, tx.Error
	}

	return shortened, nil
}

// get a specific short URL, return error if any error occurs in the process of populating the struct
func GetShortened(id uint64) (ShortURL, error) {

	var shortened ShortURL

	tx := db.Where("id = ?", id).First(&shortened) // find first Shortened with given id

	if tx.Error != nil {
		return ShortURL{}, tx.Error
	}

	return shortened, nil

}

// create a short URL, return error if any error occurs in the process of creating the struct
func CreateShortened(Shortened ShortURL) error {

	// check for duplicates, if there is already a shortened URL with the same redirect, return nil
	var shortened ShortURL

	tx := db.Where("redirect = ?", Shortened.Redirect).First(&shortened)

	if tx.Error != nil {
		return tx.Error
	}

	if shortened.ID != 0 {

		// print to console
		fmt.Println("There is already a shortened URL with the same redirect, returning nil")

		return nil
	}

	tx = db.Create(&Shortened)

	return tx.Error

}

// update a short URL, return error if any error occurs in the process of updating the struct
func UpdateShortened(Shortened ShortURL) error {

	tx := db.Save(&Shortened)
	return tx.Error

}

// delete a short URL, return error if any error occurs in the process of deleting the struct
func DeleteShortened(id uint64) error {

	// use unscoped to delete permanently
	tx := db.Unscoped().Delete(&ShortURL{}, id)
	return tx.Error

}

// get a specific short URL, return error if any error occurs in the process of populating the struct
func GetShortenedByShortURL(shortURL string) (ShortURL, error) {

	var shortened ShortURL

	tx := db.Where("url_shortener = ?", shortURL).First(&shortened) // find first Shortened with given id

	if tx.Error != nil {
		return ShortURL{}, tx.Error
	}

	return shortened, nil

}
