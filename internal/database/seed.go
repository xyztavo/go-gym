package database

import (
	"fmt"
	"strings"

	"github.com/xyztavo/go-gym/internal/models"
)

type Exercise struct {
	ID          string
	Name        string
	Description string
	GIF         string
}

func Seed() error {
	query := "INSERT INTO exercises (id, name, description, gif) VALUES "
	values := []string{}
	args := []interface{}{}
	placeholderIdx := 1
	for _, exercise := range Exercises() {
		values = append(values, fmt.Sprintf("($%d, $%d, $%d, $%d)", placeholderIdx, placeholderIdx+1, placeholderIdx+2, placeholderIdx+3))
		args = append(args, exercise.Id, exercise.Name, exercise.Description, exercise.Gif)
		placeholderIdx += 4
	}
	query += strings.Join(values, ", ") + " ON CONFLICT (id) DO NOTHING"
	_, err := db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to seed exercises: %v", err)
	}
	return nil
}

func Exercises() []models.Exercise {
	return []models.Exercise{
		{
			Id:          "0001",
			Name:        "3/4 Sit-Up",
			Description: "Lie flat, knees bent. Lift your torso to 45 degrees, then lower back down. Targets abs, with secondary focus on hip flexors and lower back.",
			Gif:         "https://v2.exercisedb.io/image/T7WQvAxHTweHpX",
		},
		{
			Id:          "0002",
			Name:        "45° Side Bend",
			Description: "Stand with feet apart. Bend torso to one side, reaching hand to knee, then return. Engages abs and obliques.",
			Gif:         "https://v2.exercisedb.io/image/RGjUYUEhcqWWMn",
		},
		{
			Id:          "0003",
			Name:        "Air Bike",
			Description: "Lie on your back, alternating elbows to opposite knees in a pedaling motion. Works abs and hip flexors.",
			Gif:         "https://v2.exercisedb.io/image/26lcWotnBa8xxR",
		},
		{
			Id:          "0004",
			Name:        "Alternate Heel Touchers",
			Description: "Lie back, lift shoulders, and reach right hand to right heel, then left hand to left heel. Focuses on abs and obliques.",
			Gif:         "https://v2.exercisedb.io/image/8oa3xx6gMylOUP",
		},
		{
			Id:          "0005",
			Name:        "Alternate Lateral Pulldown",
			Description: "Sit, pull handles towards chest, squeezing shoulder blades together. Targets lats, with biceps and rhomboids as secondary muscles.",
			Gif:         "https://v2.exercisedb.io/image/o6eLmRSmcqmoHd",
		},
		{
			Id:          "0006",
			Name:        "Assisted Chest Dip (Kneeling)",
			Description: "Kneel, lower body by bending elbows to parallel, then push back up. Targets pectorals, triceps, and shoulders.",
			Gif:         "https://v2.exercisedb.io/image/zyyPEn0HK1Ntdf",
		},
		{
			Id:          "0007",
			Name:        "Assisted Hanging Knee Raise with Throw Down",
			Description: "Hang from bar, raise knees to chest, then explosively throw legs down. Engages abs and hip flexors.",
			Gif:         "https://v2.exercisedb.io/image/cvPUJmq1PYEUQ5",
		},
		{
			Id:          "0008",
			Name:        "Assisted Hanging Knee Raise",
			Description: "Hang from a bar, lift knees to chest, then lower. Works abs and hip flexors.",
			Gif:         "https://v2.exercisedb.io/image/Db9N-MEVyVQpia",
		},
		{
			Id:          "0009",
			Name:        "Assisted Lying Leg Raise with Lateral Throw Down",
			Description: "Lie down, lift legs straight up, then lower to one side before returning to the center. Works abs and obliques.",
			Gif:         "https://v2.exercisedb.io/image/9van8nBGDKi6ZL",
		},
		{
			Id:          "0010",
			Name:        "Assisted Lying Leg Raise with Throw Down",
			Description: "Lie down, lift legs straight up, then throw them back down explosively. Targets abs and hip flexors.",
			Gif:         "https://v2.exercisedb.io/image/Tpe9WOeuVaW9Eq",
		},
		{
			Id:          "0011",
			Name:        "Assisted Motion Russian Twist",
			Description: "Sit, twist torso with medicine ball to each side. Focuses on abs and obliques.",
			Gif:         "https://v2.exercisedb.io/image/DUm7MP4C2QGuc0",
		},
		{
			Id:          "0013",
			Name:        "Assisted Pull-Up",
			Description: "Pull yourself up on a machine with an overhand grip. Works lats, biceps, and forearms.",
			Gif:         "https://v2.exercisedb.io/image/1-4gYKpwuSB34x",
		},
		{
			Id:          "0014",
			Name:        "Assisted Standing Triceps Extension (with Towel)",
			Description: "Stand and hold towel behind head, extending arms upward to work triceps.",
			Gif:         "https://v2.exercisedb.io/image/U5HIP1ecC99eI0",
		},
		{
			Id:          "0015",
			Name:        "Assisted Triceps Dip (Kneeling)",
			Description: "Kneel, lower body by bending elbows, then push back up. Focuses on triceps, chest, and shoulders.",
			Gif:         "https://v2.exercisedb.io/image/7mTxV975pr2tLp",
		},
		{
			Id:          "0016",
			Name:        "Assisted Prone Hamstring",
			Description: "Lie face down, lift legs toward glutes while keeping knees straight. Targets hamstrings.",
			Gif:         "https://v2.exercisedb.io/image/mnACtBJ9RSwwvI",
		},
		{
			Id:          "0017",
			Name:        "Barbell Pullover to Press",
			Description: "Lie on a bench, lower barbell behind head, then press it back above chest. Works lats, chest, shoulders, and triceps.",
			Gif:         "https://v2.exercisedb.io/image/soatJjEAO7OZk0",
		},
		{
			Id:          "0018",
			Name:        "Barbell Alternate Biceps Curl",
			Description: "Alternate curling barbells with palms facing forward to work biceps and forearms.",
			Gif:         "https://v2.exercisedb.io/image/kzuxBTL44O8csj",
		},
		{
			Id:          "0019",
			Name:        "Barbell Bench Front Squat",
			Description: "Squat with barbell on chest, focusing on quads, hamstrings, glutes, and calves.",
			Gif:         "https://v2.exercisedb.io/image/Qlx3LmR98fSVYS",
		},
		{
			Id:          "0020",
			Name:        "Balance Board",
			Description: "Balance on a board with one foot, shifting weight to engage quads, calves, hamstrings, and glutes.",
			Gif:         "https://v2.exercisedb.io/image/cZHD43iTPUW-hB",
		},
		{
			Id:          "0021",
			Name:        "Assisted Parallel Close Grip Pull-Up",
			Description: "Pull yourself upwards on a parallel grip machine, targeting lats and biceps.",
			Gif:         "https://v2.exercisedb.io/image/uGxwioKUwo7lx8",
		},
		{
			Id:          "0025",
			Name:        "Barbell Bench Press",
			Description: "Lie flat on a bench, grasp the barbell with an overhand grip, lower it to your chest, then push back up to the starting position.",
			Gif:         "https://v2.exercisedb.io/image/k3AicBBZ6xgqE2",
		},
		{
			Id:          "0026",
			Name:        "Barbell Bench Squat",
			Description: "Set up a barbell at chest height, squat down keeping your back straight, then push through your heels to return to the starting position.",
			Gif:         "https://v2.exercisedb.io/image/VZQ5mHws2iTlFG",
		},
		{
			Id:          "0027",
			Name:        "Barbell Bent Over Row",
			Description: "Bend forward, grasp the barbell, pull it to your lower chest by retracting your shoulder blades, then slowly lower it back down.",
			Gif:         "https://v2.exercisedb.io/image/k4G5cJbOFQKqvZ",
		},
		{
			Id:          "0028",
			Name:        "Barbell Clean and Press",
			Description: "Lift the barbell from the ground, catch it at shoulder level, and press it overhead before lowering it back to the starting position.",
			Gif:         "https://v2.exercisedb.io/image/Yt3uQb5gmUORGe",
		},
		{
			Id:          "0029",
			Name:        "Barbell Clean-Grip Front Squat",
			Description: "Hold the barbell at chest level, squat down with your chest up and knees behind toes, then return to standing by pushing through your heels.",
			Gif:         "https://v2.exercisedb.io/image/Hnf9AtWMUAsXWz",
		},
		{
			Id:          "0030",
			Name:        "Barbell Close-Grip Bench Press",
			Description: "Lie flat on a bench, grip the barbell close, lower it to your chest, and push it back up by extending your arms.",
			Gif:         "https://v2.exercisedb.io/image/59GY4iCqA5ig7g",
		},
		{
			Id:          "0031",
			Name:        "Barbell Curl",
			Description: "Hold the barbell with an underhand grip, curl the bar up to shoulder level, then lower it back to the starting position.",
			Gif:         "https://v2.exercisedb.io/image/sCV797EsVo84xu",
		},
		{
			Id:          "0032",
			Name:        "Barbell Deadlift",
			Description: "Bend at the hips and knees to grasp the barbell, then lift it by extending your hips and knees, keeping your back straight.",
			Gif:         "https://v2.exercisedb.io/image/N-VD7XzJBoOb7y",
		},
		{
			Id:          "0033",
			Name:        "Barbell Decline Bench Press",
			Description: "Lie on a decline bench, lower the barbell to your chest, then push it back up by extending your arms.",
			Gif:         "https://v2.exercisedb.io/image/bIHk03txWGmVWJ",
		},
		{
			Id:          "0034",
			Name:        "Barbell Decline Bent Arm Pullover",
			Description: "Hold the barbell overhead, lower it behind your head, then pull it back up by contracting your lats.",
			Gif:         "https://v2.exercisedb.io/image/eqb-cUnCssqeio",
		},
		{
			Id:          "0035",
			Name:        "Barbell Decline Close Grip to Skull Press",
			Description: "Lie on a decline bench, lower the barbell to your forehead, then extend your arms to press it back up.",
			Gif:         "https://v2.exercisedb.io/image/r4dIPKkIzDx6nV",
		},
		{
			Id:          "0036",
			Name:        "Barbell Decline Wide-Grip Press",
			Description: "Lie on a decline bench, lower the barbell to your chest with a wide grip, then press it back up by extending your arms.",
			Gif:         "https://v2.exercisedb.io/image/P0JRjYS8lQnf4e",
		},
		{
			Id:          "0037",
			Name:        "Barbell Decline Wide-Grip Pullover",
			Description: "Hold the barbell with a wide grip, lower it behind your head, then pull it back up by contracting your lats.",
			Gif:         "https://v2.exercisedb.io/image/S8GXXQd0yLYTHN",
		},
		{
			Id:          "0038",
			Name:        "Barbell Drag Curl",
			Description: "Hold the barbell with an underhand grip, curl the bar up to your chest while keeping your upper arms stationary, then lower it back.",
			Gif:         "https://v2.exercisedb.io/image/JSvhg1p3Av8Nzv",
		},
		{
			Id:          "0039",
			Name:        "Barbell Front Chest Squat",
			Description: "Hold the barbell in front of your chest, squat down keeping your chest up and knees behind toes, then return to standing.",
			Gif:         "https://v2.exercisedb.io/image/CcVhf-tikDf-7W",
		},
		{
			Id:          "0040",
			Name:        "Barbell Front Raise and Pullover",
			Description: "Lift the barbell in front of you to shoulder height, lower it behind your head, then raise it back up.",
			Gif:         "https://v2.exercisedb.io/image/3XMZwT1Y7PuVYE",
		},
		{
			Id:          "0041",
			Name:        "Barbell Front Raise",
			Description: "Lift the barbell in front of you to shoulder height, pause, and lower it back down to the starting position.",
			Gif:         "https://v2.exercisedb.io/image/fuXefuh-de4Fkk",
		},
		{
			Id:          "0042",
			Name:        "Barbell Front Squat",
			Description: "Hold the barbell in front of your shoulders, squat down while keeping your chest up and knees behind toes, then return to standing.",
			Gif:         "https://v2.exercisedb.io/image/fjjVjeVZxN25HY",
		},
		{
			Id:          "0043",
			Name:        "Barbell Full Squat",
			Description: "Hold the barbell on your upper back, squat down keeping your knees behind toes, then drive through your heels to stand back up.",
			Gif:         "https://v2.exercisedb.io/image/vJIduDipFeCIoj",
		},
		{
			Id:          "0044",
			Name:        "Barbell Good Morning",
			Description: "Stand with the barbell on your back, hinge forward at the hips while keeping your back straight, then return to standing.",
			Gif:         "https://v2.exercisedb.io/image/YwcE931zGHtATR",
		},
		{
			Id:          "0045",
			Name:        "Barbell Guillotine Bench Press",
			Description: "Targets pectorals with secondary focus on shoulders and triceps. Lower the bar to your neck and push it back up.",
			Gif:         "https://v2.exercisedb.io/image/T6RTgxdwhGZnjU",
		},
		{
			Id:          "0046",
			Name:        "Barbell Hack Squat",
			Description: "Focuses on glutes with secondary work for quadriceps, hamstrings, and calves. Squat with the barbell behind your legs.",
			Gif:         "https://v2.exercisedb.io/image/vhC1QU-9p-mLuZ",
		},
		{
			Id:          "0047",
			Name:        "Barbell Incline Bench Press",
			Description: "Targets pectorals, shoulders, and triceps. Perform the press on an incline bench.",
			Gif:         "https://v2.exercisedb.io/image/L2cg0HIdhtfAP9",
		},
		{
			Id:          "0048",
			Name:        "Barbell Incline Reverse-Grip Press",
			Description: "Focuses on triceps with secondary work for chest and shoulders. Perform the press with a reverse grip.",
			Gif:         "https://v2.exercisedb.io/image/dW2SC6sg9e6Wgs",
		},
		{
			Id:          "0049",
			Name:        "Barbell Incline Row",
			Description: "Works the upper back, biceps, and forearms. Perform rows on an incline bench.",
			Gif:         "https://v2.exercisedb.io/image/D3F6xy6TDwAZEr",
		},
		{
			Id:          "0050",
			Name:        "Barbell Incline Shoulder Raise",
			Description: "Focuses on serratus anterior, with secondary work for deltoids and trapezius. Perform shoulder raises on an incline bench.",
			Gif:         "https://v2.exercisedb.io/image/JeXpONCEEzxu87",
		},
		{
			Id:          "0051",
			Name:        "Barbell Jefferson Squat",
			Description: "Targets glutes with secondary muscles being quadriceps, hamstrings, and calves. Perform a squat with a unique stance.",
			Gif:         "https://v2.exercisedb.io/image/qxcVaVmWy6eVyV",
		},
		{
			Id:          "0052",
			Name:        "Barbell JM Bench Press",
			Description: "Focuses on triceps, with secondary work for chest and shoulders. Perform a close grip press on a flat bench.",
			Gif:         "https://v2.exercisedb.io/image/h4hLuwBwLYG3dE",
		},
		{
			Id:          "0053",
			Name:        "Barbell Jump Squat",
			Description: "Primarily targets glutes with secondary benefits for quadriceps, hamstrings, and calves. Explode upwards from a squat.",
			Gif:         "https://v2.exercisedb.io/image/4-f0YtTYUcgbz5",
		},
		{
			Id:          "0054",
			Name:        "Barbell Lunge",
			Description: "Targets glutes with secondary work for quadriceps, hamstrings, and calves. Perform lunges while holding a barbell on your back.",
			Gif:         "https://v2.exercisedb.io/image/BCRLdWtHahSjJy",
		},
		{
			Id:          "0055",
			Name:        "Barbell Lying Close-Grip Press",
			Description: "Focuses on triceps with secondary muscles in chest and shoulders. Perform a close-grip press while lying flat on a bench.",
			Gif:         "https://v2.exercisedb.io/image/vtTeM01WBUcPVY",
		},
		{
			Id:          "0056",
			Name:        "Barbell Lying Close-Grip Triceps Extension",
			Description: "Targets triceps with a focus on shoulders. Perform a triceps extension while lying flat on the bench.",
			Gif:         "https://v2.exercisedb.io/image/7JjXJBE1PHAN6n",
		},
		{
			Id:          "0057",
			Name:        "Barbell Lying Extension",
			Description: "Focuses on triceps with secondary work for shoulders. Perform an extension while lying on a bench.",
			Gif:         "https://v2.exercisedb.io/image/VrY54P0WJ6Pqrk",
		},
		{
			Id:          "0058",
			Name:        "Barbell Lying Lifting (on Hip)",
			Description: "Targets glutes with secondary muscles in hamstrings and quadriceps. Perform a hip raise with the barbell resting on your hips.",
			Gif:         "https://v2.exercisedb.io/image/Z9aB6rAPt0YAPg",
		},
		{
			Id:          "0059",
			Name:        "Barbell Lying Preacher Curl",
			Description: "Focuses on biceps, with secondary work for forearms. Perform curls while seated on a preacher bench.",
			Gif:         "https://v2.exercisedb.io/image/VypZzKmzm3n5Bx",
		},
		{
			Id:          "0060",
			Name:        "Barbell Lying Triceps Extension Skull Crusher",
			Description: "Targets triceps with secondary work for shoulders. Perform a skull crusher with the barbell while lying on a bench.",
			Gif:         "https://v2.exercisedb.io/image/hen94FTLhKfxUL",
		},
		{
			Id:          "0061",
			Name:        "Barbell Lying Triceps Extension",
			Description: "Focuses on triceps with secondary work for shoulders. Perform triceps extensions while lying on a bench.",
			Gif:         "https://v2.exercisedb.io/image/blkhAtrjbYaAJO",
		},
		{
			Id:          "0063",
			Name:        "Barbell Narrow Stance Squat",
			Description: "Targets glutes, with secondary muscles being quadriceps, hamstrings, and calves. Perform a squat with a narrow stance.",
			Gif:         "https://v2.exercisedb.io/image/2mYdkDO3qgdyfg",
		},
		{
			Id:          "0064",
			Name:        "Barbell One Arm Bent Over Row",
			Description: "Works the upper back with secondary muscles in biceps and forearms. Perform rows with one arm at a time.",
			Gif:         "https://v2.exercisedb.io/image/ztPEwGIdlTJ3Tf",
		},
		{
			Id:          "0065",
			Name:        "Barbell One Arm Floor Press",
			Description: "Focuses on triceps with secondary work for chest and shoulders. Perform the floor press with one arm at a time.",
			Gif:         "https://v2.exercisedb.io/image/pQSSpz7b4b5pVY",
		},
		{
			Id:          "0066",
			Name:        "Barbell One Arm Side Deadlift",
			Description: "Targets the glutes with secondary focus on hamstrings, quadriceps, and lower back. Perform deadlifts with one arm.",
			Gif:         "https://v2.exercisedb.io/image/Js4FWrOVacyqSL",
		},
		{
			Id:          "0067",
			Name:        "Barbell One Arm Snatch",
			Description: "Strengthens delts with secondary work on traps, forearms, and core. Snatch with one arm for explosive power.",
			Gif:         "https://v2.exercisedb.io/image/IQIbmzjI5OJcWz",
		},
		{
			Id:          "0068",
			Name:        "Barbell One Leg Squat",
			Description: "Focuses on quads, glutes, and hamstrings. Perform squats on one leg for balance and strength.",
			Gif:         "https://v2.exercisedb.io/image/NnMLqFXa1jWerE",
		},
		{
			Id:          "0069",
			Name:        "Barbell Overhead Squat",
			Description: "Engages quads, glutes, and core with overhead stability. Perform squats with the barbell overhead.",
			Gif:         "https://v2.exercisedb.io/image/SMJGCU9aPXPLG9",
		},
		{
			Id:          "0070",
			Name:        "Barbell Preacher Curl",
			Description: "Isolates the biceps with added forearm activation. Use a preacher bench for strict form.",
			Gif:         "https://v2.exercisedb.io/image/FzRMahMSOWgNH4",
		},
		{
			Id:          "0071",
			Name:        "Barbell Press Sit-Up",
			Description: "Targets abs with secondary work on shoulders and chest. Perform sit-ups while pressing the barbell.",
			Gif:         "https://v2.exercisedb.io/image/167dVr6BCqoNBe",
		},
		{
			Id:          "0072",
			Name:        "Barbell Prone Incline Curl",
			Description: "Targets biceps with added focus on forearms. Perform curls lying face down on an incline bench.",
			Gif:         "https://v2.exercisedb.io/image/N696KPe7pCd8Hr",
		},
		{
			Id:          "0073",
			Name:        "Barbell Pullover",
			Description: "Works lats with secondary activation of chest and triceps. Perform pullovers for upper body stretch and strength.",
			Gif:         "https://v2.exercisedb.io/image/XFZDOP-eHURG64",
		},
		{
			Id:          "0074",
			Name:        "Barbell Rack Pull",
			Description: "Targets glutes with secondary focus on hamstrings and lower back. Perform pulls from a rack for partial range.",
			Gif:         "https://v2.exercisedb.io/image/5FBYm8tcq4NFY6",
		},
		{
			Id:          "0075",
			Name:        "Barbell Rear Delt Raise",
			Description: "Strengthens delts with added focus on traps and rhomboids. Perform raises with a barbell for rear shoulder work.",
			Gif:         "https://v2.exercisedb.io/image/9X955Cl55vTpi1",
		},
		{
			Id:          "0076",
			Name:        "Barbell Rear Delt Row",
			Description: "Engages delts, traps, and biceps. Perform rows for upper back and shoulder strength.",
			Gif:         "https://v2.exercisedb.io/image/46klYY3Z71PPk2",
		},
		{
			Id:          "0077",
			Name:        "Barbell Rear Lunge V. 2",
			Description: "Focuses on glutes with secondary work for quads, hamstrings, and calves. Perform lunges with a barbell.",
			Gif:         "https://v2.exercisedb.io/image/yl6786ZkaQ-nPu",
		},
		{
			Id:          "0078",
			Name:        "Barbell Rear Lunge",
			Description: "Targets glutes, quads, and hamstrings. Perform lunges while stepping backward with a barbell.",
			Gif:         "https://v2.exercisedb.io/image/bdVcKNcctolHqP",
		},
		{
			Id:          "0079",
			Name:        "Barbell Reverse Wrist Curl V. 2",
			Description: "Strengthens forearms with additional bicep engagement. Perform reverse wrist curls with a barbell.",
			Gif:         "https://v2.exercisedb.io/image/zT9rmpLKlgnq-1",
		},
		{
			Id:          "0080",
			Name:        "Barbell Reverse Curl",
			Description: "Works biceps with secondary focus on forearms. Perform reverse curls with a barbell for arm strength.",
			Gif:         "https://v2.exercisedb.io/image/c3rZWxDKdO6tKp",
		},
		{
			Id:          "0081",
			Name:        "Barbell Reverse Preacher Curl",
			Description: "Targets biceps with forearm activation. Use a preacher bench for strict form during reverse curls.",
			Gif:         "https://v2.exercisedb.io/image/sekdnqNsRFITJD",
		},
		{
			Id:          "0082",
			Name:        "Barbell Reverse Wrist Curl",
			Description: "Strengthens forearms with added bicep activation. Perform reverse wrist curls with controlled motion.",
			Gif:         "https://v2.exercisedb.io/image/fBRGi8sANvDAWv",
		},
		{
			Id:          "0083",
			Name:        "Barbell Rollout from Bench",
			Description: "Targets abs with secondary work on shoulders and triceps. Perform rollouts using a barbell on a bench.",
			Gif:         "https://v2.exercisedb.io/image/z4Cfk-Qokm4FnK",
		},
		{
			Id:          "0084",
			Name:        "Barbell Rollout",
			Description: "Strengthens abs and lower back. Perform rollouts with a barbell for core stability and strength.",
			Gif:         "https://v2.exercisedb.io/image/bjXEwDOAS9I6v9",
		},
		{
			Id:          "0085",
			Name:        "Barbell Romanian Deadlift",
			Description: "Focuses on glutes and hamstrings with lower back support. Perform deadlifts with controlled motion.",
			Gif:         "https://v2.exercisedb.io/image/MiP-t4aZveXzJp",
		},
	}
}
