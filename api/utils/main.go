package utils

import (
	"crypto/rand"
	"fmt"
	"strings"

	"github.com/ZayenJS/models"
)

func ValidateSqlSortType(sort string) string {
	if sort != "asc" && sort != "desc" {
		return "asc"
	}

	return sort
}

func IndexFoodsByLetter(foods []models.Food) map[string][]models.Food {
	foodsIndex := make(map[string][]models.Food)

	for _, food := range foods {
		firstLetter := strings.ToLower(string(food.Name[0]))
		foodsIndex[firstLetter] = append(foodsIndex[firstLetter], food)
	}

	return foodsIndex
}

func GenerateUUIDv4() (string, error) {
	// Create a 16-byte slice for UUID
	uuid := make([]byte, 16)

	// Read random bytes from the crypto/rand source
	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}

	// Set version (4) and variant (RFC 4122) as per the UUIDv4 specification
	// Set the 6th byte to 0x4 (version 4)
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// Set the 8th byte to one of the variant values (0x80-0xBF)
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// Return the UUID in the standard format (8-4-4-4-12)
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
