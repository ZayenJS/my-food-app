package models

import (
	"fmt"

	"github.com/ZayenJS/database"
)

type Tag struct {
	TagId     int     `json:"tag_id"`
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func TagTableName() string {
	return "tags"
}

func GetAllTags() ([]*Tag, error) {
	stmt, err := database.Db.Prepare(fmt.Sprintf("SELECT * FROM %s", TagTableName()))

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	tags := []*Tag{}

	for rows.Next() {
		var tag Tag
		err := rows.Scan(&tag.TagId, &tag.Name, &tag.Color, &tag.CreatedAt, &tag.UpdatedAt)

		if err != nil {
			return nil, err
		}

		tags = append(tags, &tag)
	}

	return tags, nil
}
