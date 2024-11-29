package repository

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/mizmorr/songslib/internal/model"
	"github.com/mizmorr/songslib/store/pg"
	"gorm.io/gorm"
)

type SongRepository struct {
	db *pg.DB
}

func NewSongRepository(db *pg.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) Create(ctx context.Context, song *model.Song) (id uint, err error) {
	result := r.db.Create(song)

	if result.Error != nil {
		return 0, result.Error
	}

	return song.ID, nil
}

func (r *SongRepository) Delete(ctx context.Context, song *model.Song) error {
	var result *gorm.DB

	if song.ID == 0 {
		result = r.db.Where("band=?", song.Band).Where("name=?", song.Name).Where("lyrics=?", song.Lyrics).Delete(song)
	} else {
		result = r.db.Delete(song)
	}
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *SongRepository) GetAllFiltredPaginated(ctx context.Context, song *model.Song, pageOpts *model.Page) (int64, []*model.Song, error) {
	var (
		totalSongs int64
		songs      []*model.Song
	)
	conditionString := r.makeStringCondForGetAll(song)

	query := r.db.Model(song).Where(conditionString)
	query.Count(&totalSongs)

	offset := (pageOpts.Number - 1) * pageOpts.Size

	countSongsNeedToHave := int64(pageOpts.Size + offset)
	if totalSongs < countSongsNeedToHave {
		return totalSongs, nil, fmt.Errorf("Do not have enough songs to fulfill the page options")
	}

	result := query.Limit(pageOpts.Size).Offset(offset).Find(&songs)
	if result.Error != nil {
		return 0, nil, result.Error
	}

	return totalSongs, songs, nil
}

func (r *SongRepository) makeStringCondForGetAll(song *model.Song) string {
	if song.ID != 0 {
		return fmt.Sprintf("id = %d", song.ID)
	}
	conditionString := r.conditionStringBuilder(*song)

	replaced := strings.ReplaceAll(conditionString, "=", " like ")

	rgxpForReplaceQuotes := regexp.MustCompile(`'([^']*)'`)

	stringToFilterType := rgxpForReplaceQuotes.ReplaceAllString(replaced, "'%$1%'")

	return strings.ReplaceAll(stringToFilterType, "%''%", "''")
}

func (r *SongRepository) Update(ctx context.Context, song *model.Song) error {
	result := r.db.Model(song).Where("id = ?", song.ID).Updates(song)
	return result.Error
}

func (r *SongRepository) GetVerses(ctx context.Context, song *model.Song, pageOpt *model.Page) (*model.Verse, error) {
	offset := (pageOpt.Number - 1) * pageOpt.Size
	minPos := offset
	maxPos := offset + pageOpt.Size

	verseLines, err := r.getLines(song, minPos, maxPos)
	if err != nil {
		return nil, err
	}

	return r.makeVerse(verseLines, song, pageOpt.Number), nil
}

func (r *SongRepository) getLines(song *model.Song, minPos, maxPos int) ([]string, error) {
	var lines []string
	conditionString := r.conditionStringBuilder(*song)
	query := fmt.Sprintf(`
				SELECT line
				FROM unnest(string_to_array((select lyrics from songs where %s), E'\n')) WITH ORDINALITY AS lines(line, position)
        WHERE position > ? AND position <= ?`, conditionString)
	result := r.db.Raw(query, minPos, maxPos).Scan(&lines)

	if len(lines) == 0 {
		return nil, fmt.Errorf("No lines found with that size")
	}

	if result.Error != nil {
		return nil, result.Error
	}
	return lines, nil
}

func (r *SongRepository) conditionStringBuilder(song model.Song) string {
	var (
		typeStruct  = reflect.TypeOf(song)
		valueStruct = reflect.ValueOf(song)
		condString  strings.Builder
	)

	if song.ID != 0 {
		return fmt.Sprintf("ID ='%d'", song.ID)
	}

	for i := 0; i < typeStruct.NumField(); i++ {
		field := typeStruct.Field(i)
		value := valueStruct.Field(i).Interface()

		if !r.isEmpty(value) {
			if condString.Len() > 0 {
				condString.WriteString(" AND ")
			}
			valueCorrectToSql := strings.ReplaceAll(value.(string), "'", "''")
			condString.WriteString(field.Name + "='" + valueCorrectToSql + "'")
		}

	}
	if condString.Len() > 0 {
		return condString.String()
	}
	return ""
}

func (r *SongRepository) isEmpty(value any) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int:
		return v.Int() == 0
	case reflect.Struct:
		return reflect.DeepEqual(value, reflect.Zero(v.Type()).Interface())
	}
	return false
}

func (r *SongRepository) makeVerse(lines []string, song *model.Song, verseNumber int) *model.Verse {
	var songDB model.Song
	r.db.First(song).Scan(&songDB)

	return &model.Verse{
		Number: verseNumber,
		Song:   songDB.Name,
		Band:   songDB.Band,
		Lines:  lines,
	}
}
