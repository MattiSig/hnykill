package main

// Difficulty levels for the quiz.
const (
	DiffKids  = "kids"
	DiffAdult = "adult"
	DiffNerd  = "nerd"
)

// Question is a single multiple-choice question.
type Question struct {
	Text    string   `json:"text"`
	Options []string `json:"options"`
	Answer  int      `json:"-"` // index into Options of the correct answer
	Fact    string   `json:"-"` // shown after the answer is revealed
}

// difficultyMeta describes a difficulty level for the UI.
type difficultyMeta struct {
	Key       string
	Title     string
	Tagline   string
	Emoji     string
	Points    int
	Accent    string // tailwind colour name used for theming
	Questions []Question
}

// difficulties holds every difficulty and its curated question bank for the
// FIFA World Cup 2026 (hosted by Canada, Mexico and the USA).
var difficulties = map[string]difficultyMeta{
	DiffKids: {
		Key:     DiffKids,
		Title:   "Kids",
		Tagline: "Fun and friendly — perfect for the little ones",
		Emoji:   "🧒",
		Points:  100,
		Accent:  "emerald",
		Questions: []Question{
			{
				Text:    "How many countries are hosting the 2026 World Cup?",
				Options: []string{"1", "2", "3", "4"},
				Answer:  2,
				Fact:    "Three! Canada, Mexico and the USA are hosting together for the very first time.",
			},
			{
				Text:    "What sport is the World Cup all about?",
				Options: []string{"Basketball", "Soccer (Football)", "Tennis", "Swimming"},
				Answer:  1,
				Fact:    "It's the biggest soccer (football) tournament in the world!",
			},
			{
				Text:    "How many players from one team are on the field at the same time?",
				Options: []string{"9", "10", "11", "12"},
				Answer:  2,
				Fact:    "Eleven players per team, including the goalkeeper.",
			},
			{
				Text:    "The 2026 World Cup has three animal mascots. What animal is \"Maple\" from Canada?",
				Options: []string{"Beaver", "Moose", "Bear", "Wolf"},
				Answer:  1,
				Fact:    "Maple is a moose — a goalkeeper who loves art and music!",
			},
			{
				Text:    "What colour card does the referee show to send a player off?",
				Options: []string{"Yellow", "Green", "Red", "Blue"},
				Answer:  2,
				Fact:    "A red card means the player has to leave the game.",
			},
			{
				Text:    "How often does the World Cup happen?",
				Options: []string{"Every year", "Every 2 years", "Every 4 years", "Every 10 years"},
				Answer:  2,
				Fact:    "Every four years — so it's a really special event!",
			},
			{
				Text:    "Mexico's mascot \"Zayu\" is which jungle animal?",
				Options: []string{"Lion", "Tiger", "Jaguar", "Leopard"},
				Answer:  2,
				Fact:    "Zayu is a jaguar from the jungles of southern Mexico — and a speedy striker.",
			},
			{
				Text:    "What do we call the player who guards the goal?",
				Options: []string{"Striker", "Defender", "Goalkeeper", "Captain"},
				Answer:  2,
				Fact:    "The goalkeeper is the only player allowed to use their hands (inside the box).",
			},
			{
				Text:    "The USA mascot is a bald eagle. What is its name?",
				Options: []string{"Sammy", "Clutch", "Spike", "Buddy"},
				Answer:  1,
				Fact:    "Clutch the bald eagle is a midfielder who loves bringing people together.",
			},
			{
				Text:    "What do players try to score to win the game?",
				Options: []string{"Points", "Goals", "Runs", "Baskets"},
				Answer:  1,
				Fact:    "Goals! The team with the most goals wins.",
			},
		},
	},
	DiffAdult: {
		Key:     DiffAdult,
		Title:   "Adult",
		Tagline: "A solid challenge for the whole family table",
		Emoji:   "🧑",
		Points:  200,
		Accent:  "indigo",
		Questions: []Question{
			{
				Text:    "How many teams are competing in the expanded 2026 World Cup?",
				Options: []string{"32", "40", "48", "64"},
				Answer:  2,
				Fact:    "2026 is the first 48-team World Cup, up from 32.",
			},
			{
				Text:    "Which country won the 2022 World Cup in Qatar?",
				Options: []string{"France", "Brazil", "Argentina", "Germany"},
				Answer:  2,
				Fact:    "Argentina beat France on penalties — they're the defending champions.",
			},
			{
				Text:    "The 2026 final will be played at MetLife Stadium. In which US state is it?",
				Options: []string{"New York", "New Jersey", "California", "Texas"},
				Answer:  1,
				Fact:    "MetLife Stadium is in East Rutherford, New Jersey. The final is on 19 July 2026.",
			},
			{
				Text:    "In 2026 Mexico becomes the first nation to host the men's World Cup three times. Which years did it host before?",
				Options: []string{"1970 & 1986", "1966 & 1994", "1958 & 1982", "1978 & 1990"},
				Answer:  0,
				Fact:    "Mexico previously hosted in 1970 and 1986.",
			},
			{
				Text:    "Who won the Golden Ball (best player) at the 2022 World Cup?",
				Options: []string{"Kylian Mbappé", "Lionel Messi", "Luka Modrić", "Neymar"},
				Answer:  1,
				Fact:    "Lionel Messi won the Golden Ball as he finally lifted the trophy.",
			},
			{
				Text:    "How many matches will be played in total at the 2026 World Cup?",
				Options: []string{"64", "80", "104", "128"},
				Answer:  2,
				Fact:    "104 matches — up from 64, thanks to the 48-team format.",
			},
			{
				Text:    "Which legendary stadium hosts the 2026 opening match?",
				Options: []string{"Estadio Azteca", "MetLife Stadium", "SoFi Stadium", "BMO Field"},
				Answer:  0,
				Fact:    "The Estadio Azteca in Mexico City opened the tournament on 11 June 2026.",
			},
			{
				Text:    "How many groups are there in the 2026 group stage?",
				Options: []string{"8", "10", "12", "16"},
				Answer:  2,
				Fact:    "12 groups of 4 teams each.",
			},
			{
				Text:    "Which country has won the most World Cups, with five titles?",
				Options: []string{"Germany", "Italy", "Brazil", "Argentina"},
				Answer:  2,
				Fact:    "Brazil leads with five titles (1958, 1962, 1970, 1994, 2002).",
			},
			{
				Text:    "The official 2026 match ball made by Adidas is called what?",
				Options: []string{"Al Rihla", "Trionda", "Brazuca", "Jabulani"},
				Answer:  1,
				Fact:    "\"Trionda\" means \"three waves\", celebrating the three host nations.",
			},
		},
	},
	DiffNerd: {
		Key:     DiffNerd,
		Title:   "Nerd",
		Tagline: "Deep cuts for the true football obsessive",
		Emoji:   "🤓",
		Points:  300,
		Accent:  "rose",
		Questions: []Question{
			{
				Text:    "The 48-team format adds a brand-new knockout round. What is it called?",
				Options: []string{"Round of 24", "Round of 32", "Play-in Round", "Last 40"},
				Answer:  1,
				Fact:    "The Round of 32 appears at a World Cup for the first time in 2026.",
			},
			{
				Text:    "How many host cities will stage matches across the three nations?",
				Options: []string{"11", "14", "16", "23"},
				Answer:  2,
				Fact:    "16 host cities: 11 in the USA, 3 in Mexico and 2 in Canada.",
			},
			{
				Text:    "How many best third-placed teams advance from the group stage?",
				Options: []string{"4", "6", "8", "12"},
				Answer:  2,
				Fact:    "The top two of each group plus the 8 best third-placed teams reach the Round of 32.",
			},
			{
				Text:    "Which of these cities will NOT host 2026 matches in Mexico?",
				Options: []string{"Guadalajara", "Monterrey", "Cancún", "Mexico City"},
				Answer:  2,
				Fact:    "Mexico's host cities are Mexico City, Guadalajara and Monterrey — not Cancún.",
			},
			{
				Text:    "The match ball name \"Trionda\" (\"three waves\") celebrates what?",
				Options: []string{"The three host nations", "Three points for a win", "Three lions", "Three referees"},
				Answer:  0,
				Fact:    "It honours Canada, Mexico and the USA — the first three-nation hosts.",
			},
			{
				Text:    "In the 2022 final, Mbappé scored a hat-trick but France lost on penalties. What was the score after extra time?",
				Options: []string{"2–2", "3–3", "4–4", "3–2"},
				Answer:  1,
				Fact:    "It finished 3–3 after extra time; Argentina won 4–2 on penalties.",
			},
			{
				Text:    "Which two Canadian cities are 2026 host cities?",
				Options: []string{"Toronto & Vancouver", "Montreal & Calgary", "Ottawa & Edmonton", "Toronto & Montreal"},
				Answer:  0,
				Fact:    "Toronto and Vancouver are Canada's two host cities.",
			},
			{
				Text:    "Who is the defending men's World Cup champion entering the 2026 tournament?",
				Options: []string{"France", "Argentina", "Spain", "Germany"},
				Answer:  1,
				Fact:    "Argentina, 2022 champions, enter as the holders.",
			},
			{
				Text:    "The 2026 tournament runs from 11 June to which closing date?",
				Options: []string{"4 July", "12 July", "19 July", "26 July"},
				Answer:  2,
				Fact:    "The final is on 19 July 2026 at MetLife Stadium.",
			},
			{
				Text:    "Maple the mascot plays which position, fitting its personality?",
				Options: []string{"Striker", "Midfielder", "Goalkeeper", "Defender"},
				Answer:  2,
				Fact:    "Per FIFA, Maple is a dedicated goalkeeper, Zayu a striker and Clutch a midfielder.",
			},
		},
	},
}

// difficultyOrder is the display order on the landing page.
var difficultyOrder = []string{DiffKids, DiffAdult, DiffNerd}
