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
	adminId, _ := gonanoid.New()
	HashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	_, err := db.Exec(`
		INSERT INTO users (id, name, email, role, password) VALUES ($1, $2, $3, 'admin', $4);
	`, adminId, user.Name, user.Email, HashedPassword)
	if err != nil {
		return err
	}

	query := "INSERT INTO exercises (id, admin_id, name, description, gif) VALUES "
	values := []string{}
	args := []interface{}{}
	placeholderIdx := 1
	for _, exercise := range SeedExercises() {
		values = append(values, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", placeholderIdx, placeholderIdx+1, placeholderIdx+2, placeholderIdx+3, placeholderIdx+4))
		args = append(args, exercise.Id, adminId, exercise.Name, exercise.Description, exercise.Gif)
		placeholderIdx += 5
	}
	query += strings.Join(values, ", ")
	_, err = db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to seed exercises: %v", err)
	}

	_, err = db.Exec(`
    INSERT INTO collections (id, admin_id, name, description, img) VALUES
        ('0001', $1, 'Chest Triceps', 'This one has chest and triceps with forearms and abs.', 'https://fitnessprogramer.com/wp-content/uploads/2021/02/Barbell-Bench-Press.gif'),
        ('0002', $1, 'Back n Biceps', 'This one has back and biceps with forearms and abs.', 'https://fitnessprogramer.com/wp-content/uploads/2021/02/Lat-Pulldown.gif'),
        ('0003', $1, 'Leg Day', 'This collection focuses on legs and also includes abs exercise.', 'https://fitnessprogramer.com/wp-content/uploads/2021/02/LEG-EXTENSION.gif'),
        ('0004', $1, 'Chest n Back', 'Arnolds split focusing on chest and back.', 'https://fitnessprogramer.com/wp-content/uploads/2021/02/Seated-Cable-Row.gif'),
        ('0005', $1, 'Arms n Shoulders', 'Arnolds split focusing on arms and shoulders.', 'https://fitnessprogramer.com/wp-content/uploads/2021/02/Barbell-Shoulder-Press.gif'),
        ('0006', $1, 'Complete legs', 'Arnolds split focusing on legs.', 'https://fitnessprogramer.com/wp-content/uploads/2021/02/Sled-Hack-Squat.gif')
	`, adminId)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
	INSERT INTO exercises_reps_collections (id, admin_id, collection_id, exercise_id, reps, sets) VALUES
	-- ABC Routine
	('0001', $1, '0001', '0001', 12, 3),
	('0002', $1, '0001', '0002', 12, 3),
	('0003', $1, '0001', '0003', 12, 3),
	('0004', $1, '0001', '0006', 12, 3),
	('0005', $1, '0001', '0007', 12, 3),
	('0006', $1, '0001', '0008', 12, 3),
	('0007', $1, '0001', '0009', 12, 3),
	('0008', $1, '0002', '0004', 12, 3),
	('0009', $1, '0002', '0005', 12, 3),
	('0010', $1, '0002', '0010', 12, 3),
	('0011', $1, '0002', '0011', 12, 3),
	('0012', $1, '0002', '0012', 12, 3),
	('0013', $1, '0002', '0013', 12, 3),
	('0014', $1, '0002', '0014', 12, 3),
	('0021', $1, '0003', '0015', 12, 3),
	('0022', $1, '0003', '0016', 12, 3),
	('0023', $1, '0003', '0017', 12, 3),
	('0024', $1, '0003', '0018', 12, 3),
	('0025', $1, '0003', '0019', 12, 3),
	('0026', $1, '0003', '0020', 12, 3),
	('0027', $1, '0003', '0021', 12, 3),
	('0028', $1, '0001', '0022', 12, 3),
	('0029', $1, '0002', '0023', 12, 3),
	
	-- New Exercises for Arnold Split
	('0030', $1, '0004', '0024', 10, 3), -- Dumbbell Bench Press
	('0032', $1, '0004', '0025', 10, 3), -- Incline Dumbbell Bench Press
	('0034', $1, '0004', '0003', 10, 3), -- Peck Deck Fly
	('0035', $1, '0004', '0010', 10, 3), -- Reverse Lat Pulldown
	('0037', $1, '0004', '0011', 10, 3), -- Lat Pull Down
	('0039', $1, '0004', '0012', 10, 3), -- Seated Cable Row
	('0041', $1, '0004', '0028', 10, 3), -- Face Pull
	('0042', $1, '0004', '0029', 10, 3), -- Shrugs
	('0043', $1, '0005', '0022', 10, 3), -- Leaning Cable Lateral Raise
	('0044', $1, '0005', '0023', 10, 3), -- Standing Dumbbell Overhead Press
	('0045', $1, '0005', '0032', 10, 3), -- Concentration Curl
	('0047', $1, '0005', '0005', 8, 4), -- Preacher Curl
	('0049', $1, '0005', '0034', 8, 4), -- Overhead Triceps Extension
	('0050', $1, '0005', '0007', 10, 4), -- Triceps Pushdown
	('0051', $1, '0006', '0015', 10, 3), -- Leg Press
	('0052', $1, '0006', '0016', 10, 3), -- Smith Machine Squat
	('0054', $1, '0006', '0038', 10, 3), -- Hack Squat
	('0056', $1, '0006', '0018', 10, 3), -- Seated Leg Curl
	('0058', $1, '0006', '0020', 10, 3), -- Standing Calf Raise
	('0060', $1, '0006', '0040', 10, 3); -- Seated Calf Raise
	`, adminId)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
    INSERT INTO routines (id, admin_id, name, description, img) VALUES
        ('0001', $1, 'ABC Routine', 'This routine has chest, triceps, back, biceps and leg day, including forearms and abs.', 'https://fitnessprogramer.com/wp-content/uploads/2023/06/spartan-300-workout.png'),
        ('0002', $1, 'Arnold Split', 'This routine follows the Arnold Split with chest and back, arms and shoulders with complete legs', 'https://cdn.prod.website-files.com/66429cc5ff47fb0567fcd516/666a1ad7c03bf72ed8497564_Arnold-Schwarzenegger-split.jpg')
	`, adminId)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
    INSERT INTO routines_collections (id, admin_id, routine_id, collection_id) VALUES
        ('0001', $1, '0001', '0001'), ('0002', $1, '0001', '0002'),
        ('0003', $1, '0001', '0003'), ('0004', $1, '0002', '0004'),
        ('0005', $1, '0002', '0005'), ('0006', $1, '0002', '0006')
	`, adminId)
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
		{
			Id:          "0010",
			Name:        "Reverse Lat Pulldown",
			Description: "Grip the bar palms facing you, pull it to chest level, keeping back straight. Targets biceps, lats, and middle back.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/04/Reverse-Lat-Pulldown.gif",
		},
		{
			Id:          "0011",
			Name:        "Lat Pull Down",
			Description: "Grip bar wide, pull down to chest while squeezing shoulder blades. Targets lats, traps, and rear delts.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/02/Lat-Pulldown.gif",
		},
		{
			Id:          "0012",
			Name:        "Seated Cable Row",
			Description: "Sit upright, pull handle to torso while squeezing shoulder blades. Targets lats, rhomboids, and traps.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/02/Seated-Cable-Row.gif",
		},
		{
			Id:          "0013",
			Name:        "Barbell Reverse Wrist Curl",
			Description: "Hold barbell overhand, curl wrists upward while keeping forearms still. Targets forearm extensors.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/06/Barbell-Reverse-Wrist-Curl.gif",
		},
		{
			Id:          "0014",
			Name:        "Captains Chair Leg Raise",
			Description: "Focuses on lower abs. Hang from the chair, raise legs by flexing hips and knees, and slowly return to the starting position.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/05/Captains-Chair-Leg-Raise.gif",
		},
		{
			Id:          "0015",
			Name:        "Leg Press",
			Description: "Targets quads, hamstrings, and glutes. Sit on the machine with feet shoulder-width apart, push the weight upward, and slowly lower it back down.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2015/11/Leg-Press.gif",
		},
		{
			Id:          "0016",
			Name:        "Hip Abduction Machine",
			Description: "Strengthens hip abductors and glutes. Sit on the machine, adjust pads to thighs, and push legs outward against resistance.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/02/HiP-ABDUCTION-MACHINE.gif",
		},
		{
			Id:          "0017",
			Name:        "Hip Adduction Machine",
			Description: "Strengthens hip adductors and glutes. Sit on the machine, adjust pads to thighs, and pull legs inward against resistance.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/02/HIP-ADDUCTION-MACHINE.gif",
		},
		{
			Id:          "0018",
			Name:        "Seated Leg Curl",
			Description: "Targets hamstrings. Sit on the machine, position feet on the rollers, and curl legs backward.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/08/Seated-Leg-Curl.gif",
		},
		{
			Id:          "0019",
			Name:        "Leg Extension",
			Description: "Works on quadriceps. Sit on the machine, position feet under the pad, and extend legs upward.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/02/LEG-EXTENSION.gif",
		},
		{
			Id:          "0020",
			Name:        "Calf Raise",
			Description: "Develops calves. Stand with feet on an elevated platform, hold dumbbells, and raise heels.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/02/Dumbbell-Calf-Raise.gif",
		},
		{
			Id:          "0021",
			Name:        "Crunch",
			Description: "Strengthens abs. Lie on your back, bend knees, and lift your shoulders towards your knees.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2015/11/Crunch.gif",
		},
		{
			Id:          "0022",
			Name:        "Leaning Cable Lateral Raise",
			Description: "Targets shoulders. Use a cable machine, lean away, and raise the handle to shoulder height.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/09/Leaning-Cable-Lateral-Raise.gif",
		},
		{
			Id:          "0023",
			Name:        "Standing Dumbbell Overhead Press",
			Description: "Strengthens shoulders and triceps. Stand, press dumbbells overhead, and lower slowly.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2023/09/Standing-Dumbbell-Overhead-Press.gif",
		},
		{
			Id:          "0024",
			Name:        "Dumbbell Bench Press",
			Description: "Flat bench dumbbell press for chest development.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/02/Dumbbell-Press.gif",
		},
		{
			Id:          "0025",
			Name:        "Incline Dumbbell Bench Press",
			Description: "Incline dumbbell press targeting upper chest.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/02/Incline-Dumbbell-Press.gif",
		},
		{
			Id:          "0026",
			Name:        "High Pull",
			Description: "High pull targeting the upper back.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/04/Dumbbell-Upright-Row.gif",
		},
		{
			Id:          "0028",
			Name:        "Face Pull",
			Description: "Face back pull exercise for shoulder and back development.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/02/Face-Pull.gif",
		},
		{
			Id:          "0029",
			Name:        "Dumbbell Shrug",
			Description: "Dumbbell shoulder shrugs for trap muscles.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/04/Dumbbell-Shrug.gif",
		},
		{
			Id:          "0032",
			Name:        "Concentration Curl",
			Description: "Biceps concentration curl for arm development.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/02/Concentration-Curl.gif",
		},
		{
			Id:          "0034",
			Name:        "Overhead Triceps Extension",
			Description: "Overhead triceps extension with dumbbell.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/06/Seated-Dumbbell-Triceps-Extension.gif",
		},
		{
			Id:          "0037",
			Name:        "Smith Machine Squat",
			Description: "Squat using the Smith machine.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2024/10/smith-machine-squat.gif",
		},
		{
			Id:          "0038",
			Name:        "Hack Squat",
			Description: "Hack squat for quadriceps and glutes.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/02/Sled-Hack-Squat.gif",
		},
		{
			Id:          "0040",
			Name:        "Seated Calf Raise",
			Description: "Seated calf raises to target calf muscles.",
			Gif:         "https://fitnessprogramer.com/wp-content/uploads/2021/10/Weighted-Seated-Calf-Raise.gif",
		},
	}
}
