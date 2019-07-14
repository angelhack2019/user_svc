package interfaces

import (
	"database/sql"
	"fmt"
	"github.com/angelhack2019/lib/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	pg               string
	connectionString string
	postgresDB       *sql.DB
)

func init() {
	viper.BindEnv("PG")
	pg = viper.GetString("PG")
	connectionString = fmt.Sprintf(
		"host=%s user=default password=default dbname=dough_you sslmode=disable port=5432",
		pg,
	)
}

func refreshDBConnection() error {
	if postgresDB == nil {
		var err error
		postgresDB, err = sql.Open("postgres", connectionString)
		if err != nil {
			return err
		}
	}

	if err := postgresDB.Ping(); err != nil {
		_ = postgresDB.Close()
		postgresDB = nil
		return err
	}

	return nil
}

// Duplicate from food_svc =========>

func LoginUser(email string, password string) (string, bool) {
	if err := refreshDBConnection(); err != nil {
		return err.Error(), false
	}

	command := `
				SELECT uuid 
				FROM dough_you.users
				WHERE dough_you.users.email = $1
				AND dough_you.users.password = $2
				`
	row, err := postgresDB.Query(command, email, password)

	if err != nil {
		return err.Error(), false
	}

	uuid := ""

	for row.Next() {
		err := row.Scan(&uuid)
		if err != nil {
			return err.Error(), false
		}
		if uuid == "" {
			return "Unautorized", false
		}
	}

	return uuid, true
}

func GetUser(uuid string) (models.User, string) {
	if err := refreshDBConnection(); err != nil {
		return models.User{}, err.Error()
	}

	command := `
				SELECT uuid, first_name, last_name, email, num_ratings,
				sum_ratings, bio, pic_url, school, 
				state, city, phone_number 
				FROM dough_you.users
				WHERE dough_you.users.uuid = $1
				`
	row, err := postgresDB.Query(command, uuid)

	if err != nil {
		return models.User{}, err.Error()
	}

	user := models.User{}

	for row.Next() {
		err := row.Scan(&user.UUID, &user.FirstName, &user.LastName, &user.Email, &user.NumRatings,
			&user.SumRatings, &user.Bio, &user.PicURL,
			&user.School, &user.State, &user.City, &user.PhoneNumber)
		if err != nil {
			return models.User{}, err.Error()
		}
	}

	return user, ""
}

func CreateUser(user *models.User) (string, string) {
	if err := refreshDBConnection(); err != nil {
		return "", err.Error()
	}

	command := `
				INSERT INTO dough_you.users(
				uuid, first_name, last_name, email, password, num_ratings,
				sum_ratings, bio, pic_url, school, 
				state, city, phone_number
				) VALUES($1, $2, $3, $4, 
						 $5, $6, $7, $8,
						 $9, $10, $11, $12, $13)
				`

	newUUID := uuid.New().String()
	_, err := postgresDB.Exec(command, newUUID, user.FirstName, user.LastName, user.Email, user.Password, user.NumRatings,
		user.SumRatings, user.Bio, user.PicURL,
		user.School, user.State, user.City, user.PhoneNumber)

	if err != nil {
		return "", err.Error()
	}

	return newUUID, ""
}

func EditUser(user *models.User) string {
	if err := refreshDBConnection(); err != nil {
		return err.Error()
	}

	command := `
				UPDATE dough_you.users SET
					email=$1, first_name=$2, last_name=$3, num_ratings=$4,
					sum_ratings=$5, bio=$6, pic_url=$7, school=$8, 
					state=$9, city=$10, phone_number=$11 
				WHERE dough_you.users.uuid = $12
				`
	_, err := postgresDB.Exec(command, user.Email, user.FirstName, user.LastName,
		user.NumRatings, user.SumRatings, user.Bio, user.PicURL,
		user.School, user.State, user.City, user.PhoneNumber, user.UUID)

	if err != nil {
		return err.Error()
	}

	return ""
}
