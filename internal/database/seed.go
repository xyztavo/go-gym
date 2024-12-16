package database

import (
	"fmt"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/configs"
	"github.com/xyztavo/go-gym/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type Exercise struct {
	ID          string
	Name        string
	Description string
	GIF         string
}

func Seed() error {
	user := configs.GetAdminInfo()
	id, _ := gonanoid.New()
	HashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	_, err := db.Exec(`
	INSERT INTO users (id, name, email, role, password) VALUES ($1, $2, $3, 'admin', $4);
	`, id, user.Name, user.Email, HashedPassword)
	if err != nil {
		return err
	}
	query := "INSERT INTO exercises (id, name, description, gif) VALUES "
	values := []string{}
	args := []interface{}{}
	placeholderIdx := 1
	for _, exercise := range SeedExercises() {
		values = append(values, fmt.Sprintf("($%d, $%d, $%d, $%d)", placeholderIdx, placeholderIdx+1, placeholderIdx+2, placeholderIdx+3))
		args = append(args, exercise.Id, exercise.Name, exercise.Description, exercise.Gif)
		placeholderIdx += 4
	}
	query += strings.Join(values, ", ") + " ON CONFLICT (id) DO NOTHING"
	_, err = db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to seed exercises: %v", err)
	}
	_, err = db.Exec("INSERT INTO collections (id, admin_id, name, description, img) VALUES ('0001', $1, 'Chest Triceps', 'This one has chest and triceps with forearms and abs.', 'https://fitnessprogramer.com/wp-content/uploads/2021/02/Barbell-Bench-Press.gif')", id)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO exercises_reps_collections (id, admin_id, collection_id, exercise_id, reps, sets) VALUES ('0001', $1, '0001', '0001', 12, 3), ('0002', $1, '0001', '0002', 12, 3), ('0003', $1, '0001', '0003', 12, 3), ('0004', $1, '0001', '0006', 12, 3), ('0005', $1, '0001', '0007', 12, 3), ('0006', $1, '0001', '0008', 12, 3), ('0007', $1, '0001', '0009', 12, 3)", id)
	if err != nil {
		return err
	}
	return nil
}

func SeedExercises() []models.Exercise {
	return []models.Exercise{
		{
			Id:          "0001",
			Name:        "Barbell Bench Press",
			Description: "Strengthens chest, shoulders, and triceps. Lie on a bench, lower the barbell to your chest, and press upward.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/02/Barbell-Bench-Press.gif",
		},
		{
			Id:          "0002",
			Name:        "Incline Barbell Bench Press",
			Description: "Targets upper chest and shoulders. Press a barbell upwards at an inclined bench angle.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/02/Incline-Barbell-Bench-Press.gif",
		},
		{
			Id:          "0003",
			Name:        "Peck Deck Fly",
			Description: "Isolates the chest muscles. Bring arms together on the pec deck machine in a controlled motion.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/02/Pec-Deck-Fly.gif",
		},
		{
			Id:          "0004",
			Name:        "Hammer Curl",
			Description: "Strengthens biceps and forearms. Hold dumbbells with a neutral grip and curl them upwards.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/02/Hammer-Curl.gif",
		},
		{
			Id:          "0005",
			Name:        "Ez Bar Preacher Curl",
			Description: "Isolates the biceps. Use an EZ bar on a preacher bench, curling the bar upward.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/02/Z-Bar-Preacher-Curl.gif",
		},
		{
			Id:          "0006",
			Name:        "Wrist Curl",
			Description: "Builds forearm strength. Rest your forearms on a bench, curl the barbell using your wrists.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/02/barbell-Wrist-Curl.gif",
		},
		{
			Id:          "0007",
			Name:        "Triceps Pushdown",
			Description: "Targets the triceps. Push the cable bar down until arms are fully extended.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/02/Pushdown.gif",
		},
		{
			Id:          "0008",
			Name:        "Seated One Arm Dumbbell Triceps Extension",
			Description: "Strengthens triceps. Sit, hold a dumbbell overhead, and lower it behind your head.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/06/Seated-One-Arm-Dumbbell-Triceps-Extension.gif",
		},
		{
			Id:          "0009",
			Name:        "Lying Knee Raise",
			Description: "Works the abs and hip flexors. Lie flat, raise your knees toward your chest, and lower them.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/05/Lying-Knee-Raise.gif",
		},
	}
}
