package persistency

import (
	"database/sql"
	"log"

	"github.com/josue/chalenge_golang/internal/domain/user"
	uuid "github.com/satori/go.uuid"
)

type userMySQLRepository struct {
	db *sql.DB
}

func NewUserMySQLRepository(db *sql.DB) userMySQLRepository {
	return userMySQLRepository{
		db: db,
	}
}

func (ur userMySQLRepository) SaveUser(user user.User) error {
	log.Println("On userMySQLRepository.SaveUser")

	query := "INSERT INTO user (id, user_name, phone_number, email, pass) VALUES (?,?,?,?,?)"
	randomUUID := uuid.NewV4()
	_, err := ur.db.Exec(query,
		randomUUID.String(),
		user.UserName,
		user.PhoneNumber,
		user.Email,
		user.Pass)
	if err != nil {
		log.Println("On userMySQLRepository.SaveUser. Error: ", err.Error())
		return err
	}

	log.Println("On userMySQLRepository.SaveUser -  new user created")

	return nil
}

func (ur userMySQLRepository) ValidateUserUniqueInfo(phone string, email string) (bool, error) {
	log.Println("On userMySQLRepository.ValidateUserUniqueInfo")

	var count int
	err := ur.db.QueryRow("SELECT COUNT(*) FROM user where phone_number=? OR email=?", phone, email).Scan(&count)
	if err != nil {
		log.Println("On userMySQLRepository.ValidateUserUniqueInfo. Error: ", err.Error())
		return false, err
	}

	if count >= 1 {
		return false, nil
	}

	return true, nil
}
