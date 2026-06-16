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
	PlayCount int    // how many questions a single game draws from the pool
	Questions []Question
}

// difficulties holds every difficulty and its curated question bank for the
// FIFA World Cup 2026 (hosted by Canada, Mexico and the USA). Questions lean
// towards mainstream/European football to suit a Western/European audience.
// Each game draws a random subset of PlayCount questions for replayability.
var difficulties = map[string]difficultyMeta{
	DiffKids: {
		Key:       DiffKids,
		Title:     "Kids",
		Tagline:   "Fun and friendly — perfect for the little ones",
		Emoji:     "🧒",
		Points:    100,
		Accent:    "emerald",
		PlayCount: 12,
		Questions: []Question{
			{
				Text:    "How many countries are hosting the 2026 World Cup?",
				Options: []string{"1", "2", "3", "4"},
				Answer:  2,
				Fact:    "Three! Canada, Mexico and the USA host together for the very first time.",
			},
			{
				Text:    "What sport is the World Cup all about?",
				Options: []string{"Basketball", "Football (Soccer)", "Tennis", "Rugby"},
				Answer:  1,
				Fact:    "It's the biggest football (soccer) tournament in the world!",
			},
			{
				Text:    "How many players from one team are on the pitch at the same time?",
				Options: []string{"9", "10", "11", "12"},
				Answer:  2,
				Fact:    "Eleven players per team, including the goalkeeper.",
			},
			{
				Text:    "Canada's 2026 mascot \"Maple\" is which animal?",
				Options: []string{"Beaver", "Moose", "Bear", "Wolf"},
				Answer:  1,
				Fact:    "Maple is a moose — and a keen goalkeeper who loves art and music!",
			},
			{
				Text:    "What colour card does the referee show to send a player off?",
				Options: []string{"Yellow", "Green", "Red", "Blue"},
				Answer:  2,
				Fact:    "A red card means the player has to leave the game.",
			},
			{
				Text:    "How often is the World Cup held?",
				Options: []string{"Every year", "Every 2 years", "Every 4 years", "Every 10 years"},
				Answer:  2,
				Fact:    "Every four years — so it's a really special event!",
			},
			{
				Text:    "Mexico's 2026 mascot \"Zayu\" is which jungle animal?",
				Options: []string{"Lion", "Tiger", "Jaguar", "Panther"},
				Answer:  2,
				Fact:    "Zayu is a speedy jaguar striker from the jungles of southern Mexico.",
			},
			{
				Text:    "What do we call the player who guards the goal?",
				Options: []string{"Striker", "Defender", "Goalkeeper", "Winger"},
				Answer:  2,
				Fact:    "The goalkeeper is the only player allowed to use their hands (in the box).",
			},
			{
				Text:    "The USA's 2026 mascot is a bald eagle. What is its name?",
				Options: []string{"Sammy", "Clutch", "Spike", "Buddy"},
				Answer:  1,
				Fact:    "Clutch the bald eagle is a midfielder who loves bringing people together.",
			},
			{
				Text:    "What do players try to score to win?",
				Options: []string{"Points", "Goals", "Runs", "Tries"},
				Answer:  1,
				Fact:    "Goals! The team with the most goals wins.",
			},
			{
				Text:    "How many minutes is a normal football match (not counting added time)?",
				Options: []string{"45", "60", "90", "120"},
				Answer:  2,
				Fact:    "90 minutes — two halves of 45 minutes each.",
			},
			{
				Text:    "When a player scores three goals in one match, what is it called?",
				Options: []string{"A triple", "A hat-trick", "A treble strike", "A trio"},
				Answer:  1,
				Fact:    "A hat-trick! Some fans even throw their hats in celebration.",
			},
			{
				Text:    "Cristiano Ronaldo plays for which country?",
				Options: []string{"Brazil", "Spain", "Portugal", "Italy"},
				Answer:  2,
				Fact:    "Ronaldo is Portugal's all-time top scorer and captain.",
			},
			{
				Text:    "Lionel Messi plays for which country?",
				Options: []string{"Argentina", "Brazil", "Uruguay", "Chile"},
				Answer:  0,
				Fact:    "Messi captained Argentina to victory at the 2022 World Cup.",
			},
			{
				Text:    "Brazil's football shirts are famously which colour?",
				Options: []string{"Red", "Yellow", "Blue", "Green"},
				Answer:  1,
				Fact:    "The bright yellow shirt is one of the most famous in football.",
			},
			{
				Text:    "England's national team is nicknamed the \"Three…\" what?",
				Options: []string{"Bulldogs", "Lions", "Eagles", "Bears"},
				Answer:  1,
				Fact:    "The Three Lions — from the badge on their shirts.",
			},
			{
				Text:    "If a knockout match is tied and a winner is needed, how is it decided?",
				Options: []string{"A coin toss", "A penalty shootout", "A re-match the next day", "A free-kick contest"},
				Answer:  1,
				Fact:    "After extra time, it goes to a nerve-jangling penalty shootout.",
			},
			{
				Text:    "Which body part can a goalkeeper use that other players can't (inside the box)?",
				Options: []string{"Their head", "Their hands", "Their chest", "Their feet"},
				Answer:  1,
				Fact:    "Goalkeepers can catch the ball with their hands inside their penalty area.",
			},
			{
				Text:    "How many points does a team get for winning a group game?",
				Options: []string{"1", "2", "3", "4"},
				Answer:  2,
				Fact:    "3 points for a win, 1 for a draw, 0 for a loss.",
			},
			{
				Text:    "What is the shiny gold prize the winning team lifts called?",
				Options: []string{"The Trophy (Cup)", "The Medal", "The Shield", "The Plate"},
				Answer:  0,
				Fact:    "The FIFA World Cup Trophy — made of solid gold!",
			},
			{
				Text:    "France's football team is known as \"Les …\" what (after their shirt colour)?",
				Options: []string{"Bleus (Blues)", "Rouges (Reds)", "Verts (Greens)", "Blancs (Whites)"},
				Answer:  0,
				Fact:    "\"Les Bleus\" — the Blues — were world champions in 1998 and 2018.",
			},
		},
	},
	DiffAdult: {
		Key:       DiffAdult,
		Title:     "Adult",
		Tagline:   "A solid challenge for the whole family table",
		Emoji:     "🧑",
		Points:    200,
		Accent:    "indigo",
		PlayCount: 14,
		Questions: []Question{
			{
				Text:    "How many teams compete in the expanded 2026 World Cup?",
				Options: []string{"32", "40", "48", "64"},
				Answer:  2,
				Fact:    "2026 is the first 48-team World Cup, up from 32.",
			},
			{
				Text:    "Which country won the 2022 World Cup in Qatar?",
				Options: []string{"France", "Brazil", "Argentina", "Germany"},
				Answer:  2,
				Fact:    "Argentina beat France on penalties — they enter 2026 as champions.",
			},
			{
				Text:    "The 2026 final at MetLife Stadium will be played in which US state?",
				Options: []string{"New York", "New Jersey", "California", "Florida"},
				Answer:  1,
				Fact:    "MetLife Stadium is in East Rutherford, New Jersey. The final is on 19 July 2026.",
			},
			{
				Text:    "Mexico hosted the men's World Cup before in which years?",
				Options: []string{"1970 & 1986", "1966 & 1994", "1958 & 1982", "1978 & 1990"},
				Answer:  0,
				Fact:    "1970 and 1986 — 2026 makes Mexico the first three-time host.",
			},
			{
				Text:    "Who won the Golden Ball (best player) at the 2022 World Cup?",
				Options: []string{"Kylian Mbappé", "Lionel Messi", "Luka Modrić", "Neymar"},
				Answer:  1,
				Fact:    "Lionel Messi won the Golden Ball as he finally lifted the trophy.",
			},
			{
				Text:    "How many matches are played in total at the 2026 World Cup?",
				Options: []string{"64", "80", "104", "128"},
				Answer:  2,
				Fact:    "104 matches — up from 64, thanks to the 48-team format.",
			},
			{
				Text:    "Which iconic stadium hosted the 2026 opening match?",
				Options: []string{"Estadio Azteca", "MetLife Stadium", "Wembley", "SoFi Stadium"},
				Answer:  0,
				Fact:    "Mexico City's Estadio Azteca opened the tournament on 11 June 2026.",
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
				Fact:    "Brazil leads with five (1958, 1962, 1970, 1994, 2002).",
			},
			{
				Text:    "The official 2026 Adidas match ball is called what?",
				Options: []string{"Al Rihla", "Trionda", "Brazuca", "Jabulani"},
				Answer:  1,
				Fact:    "\"Trionda\" means \"three waves\", celebrating the three host nations.",
			},
			{
				Text:    "Who won the 2018 World Cup in Russia?",
				Options: []string{"Croatia", "France", "Belgium", "England"},
				Answer:  1,
				Fact:    "France beat Croatia 4–2 in the final.",
			},
			{
				Text:    "Which country did France beat in the 2018 World Cup final?",
				Options: []string{"Croatia", "Belgium", "England", "Brazil"},
				Answer:  0,
				Fact:    "France 4–2 Croatia in Moscow — Croatia's first ever final.",
			},
			{
				Text:    "Who won the 2014 World Cup, held in Brazil?",
				Options: []string{"Argentina", "Netherlands", "Germany", "Spain"},
				Answer:  2,
				Fact:    "Germany beat Argentina 1–0 after extra time.",
			},
			{
				Text:    "Which country won the 2010 World Cup — their first ever title?",
				Options: []string{"Netherlands", "Spain", "Germany", "Italy"},
				Answer:  1,
				Fact:    "Spain beat the Netherlands 1–0 in extra time in South Africa.",
			},
			{
				Text:    "How many European (UEFA) teams qualified for the 2026 World Cup?",
				Options: []string{"13", "16", "20", "24"},
				Answer:  1,
				Fact:    "16 — UEFA's biggest-ever World Cup allocation.",
			},
			{
				Text:    "Kylian Mbappé plays for which national team?",
				Options: []string{"Belgium", "France", "Switzerland", "Morocco"},
				Answer:  1,
				Fact:    "France's captain and talisman.",
			},
			{
				Text:    "Who won the 2024 men's Ballon d'Or?",
				Options: []string{"Vinícius Júnior", "Jude Bellingham", "Rodri", "Erling Haaland"},
				Answer:  2,
				Fact:    "Rodri, the Spain & Manchester City midfielder, edged out Vinícius Júnior.",
			},
			{
				Text:    "Which country won Euro 2024?",
				Options: []string{"England", "Spain", "France", "Germany"},
				Answer:  1,
				Fact:    "Spain beat England 2–1 in Berlin for a record fourth Euro title.",
			},
			{
				Text:    "Mexico opened the 2026 World Cup against which nation?",
				Options: []string{"South Africa", "Morocco", "USA", "South Korea"},
				Answer:  0,
				Fact:    "Mexico beat South Africa 2–0 in the opening match.",
			},
			{
				Text:    "Who scored a hat-trick in the 2022 World Cup final but still finished as a runner-up?",
				Options: []string{"Lionel Messi", "Kylian Mbappé", "Ángel Di María", "Olivier Giroud"},
				Answer:  1,
				Fact:    "Mbappé scored three but France lost the shootout.",
			},
			{
				Text:    "Which country shares second place for most World Cup titles (four) with Germany?",
				Options: []string{"Italy", "Argentina", "France", "Uruguay"},
				Answer:  0,
				Fact:    "Italy also has four titles (1934, 1938, 1982, 2006).",
			},
			{
				Text:    "Which club has won the European Cup / Champions League the most times?",
				Options: []string{"Barcelona", "AC Milan", "Bayern Munich", "Real Madrid"},
				Answer:  3,
				Fact:    "Real Madrid — comfortably the most decorated club in Europe.",
			},
			{
				Text:    "Erling Haaland, one of the world's deadliest strikers, plays for which country?",
				Options: []string{"Sweden", "Norway", "Denmark", "Netherlands"},
				Answer:  1,
				Fact:    "Norway — who have rarely reached major tournaments despite his talent.",
			},
			{
				Text:    "Which player has won the most Ballon d'Or awards, with eight?",
				Options: []string{"Cristiano Ronaldo", "Lionel Messi", "Michel Platini", "Johan Cruyff"},
				Answer:  1,
				Fact:    "Messi has eight; Ronaldo is next with five.",
			},
			{
				Text:    "England midfield star Jude Bellingham plays his club football for which team?",
				Options: []string{"Manchester City", "Real Madrid", "Liverpool", "Bayern Munich"},
				Answer:  1,
				Fact:    "Real Madrid, who signed him from Borussia Dortmund in 2023.",
			},
			{
				Text:    "England's top men's football division is called what?",
				Options: []string{"La Liga", "Serie A", "Premier League", "Bundesliga"},
				Answer:  2,
				Fact:    "The Premier League — the most-watched league in the world.",
			},
		},
	},
	DiffNerd: {
		Key:       DiffNerd,
		Title:     "Nerd",
		Tagline:   "Deep cuts for the true football obsessive",
		Emoji:     "🤓",
		Points:    300,
		Accent:    "rose",
		PlayCount: 14,
		Questions: []Question{
			{
				Text:    "The 48-team format adds a brand-new knockout round. What is it called?",
				Options: []string{"Round of 24", "Round of 32", "Play-in Round", "Last 40"},
				Answer:  1,
				Fact:    "The Round of 32 appears at a World Cup for the first time in 2026.",
			},
			{
				Text:    "How many host cities stage matches across the three 2026 nations?",
				Options: []string{"11", "14", "16", "23"},
				Answer:  2,
				Fact:    "16 host cities: 11 in the USA, 3 in Mexico and 2 in Canada.",
			},
			{
				Text:    "How many best third-placed teams advance from the 2026 group stage?",
				Options: []string{"4", "6", "8", "12"},
				Answer:  2,
				Fact:    "Top two of each group plus the 8 best third-placed teams reach the Round of 32.",
			},
			{
				Text:    "Which of these is NOT a 2026 host city in Mexico?",
				Options: []string{"Guadalajara", "Monterrey", "Cancún", "Mexico City"},
				Answer:  2,
				Fact:    "Mexico's host cities are Mexico City, Guadalajara and Monterrey.",
			},
			{
				Text:    "The match-ball name \"Trionda\" (\"three waves\") celebrates what?",
				Options: []string{"The three host nations", "Three points for a win", "Three lions", "Three referees"},
				Answer:  0,
				Fact:    "It honours Canada, Mexico and the USA — the first three-nation hosts.",
			},
			{
				Text:    "In the 2022 final, Mbappé scored a hat-trick but France lost on penalties. What was the score after extra time?",
				Options: []string{"2–2", "3–3", "4–4", "3–2"},
				Answer:  1,
				Fact:    "It finished 3–3; Argentina won 4–2 on penalties.",
			},
			{
				Text:    "Which two Canadian cities are 2026 host cities?",
				Options: []string{"Toronto & Vancouver", "Montreal & Calgary", "Ottawa & Edmonton", "Toronto & Montreal"},
				Answer:  0,
				Fact:    "Toronto and Vancouver are Canada's two host cities.",
			},
			{
				Text:    "Who is the World Cup's all-time leading goalscorer, with 16 goals?",
				Options: []string{"Ronaldo (Brazil)", "Miroslav Klose", "Gerd Müller", "Just Fontaine"},
				Answer:  1,
				Fact:    "Germany's Klose (16) overtook Brazil's Ronaldo (15) in 2014.",
			},
			{
				Text:    "On which date is the 2026 World Cup final played?",
				Options: []string{"4 July", "12 July", "19 July", "26 July"},
				Answer:  2,
				Fact:    "19 July 2026 at MetLife Stadium.",
			},
			{
				Text:    "Who scored the only hat-trick in a World Cup final, for England in 1966?",
				Options: []string{"Bobby Charlton", "Geoff Hurst", "Martin Peters", "Gary Lineker"},
				Answer:  1,
				Fact:    "Geoff Hurst, in England's 4–2 win over West Germany.",
			},
			{
				Text:    "Who scored Germany's extra-time winner in the 2014 World Cup final?",
				Options: []string{"Thomas Müller", "Mario Götze", "André Schürrle", "Bastian Schweinsteiger"},
				Answer:  1,
				Fact:    "Mario Götze's superb volley beat Argentina 1–0.",
			},
			{
				Text:    "Which player was sent off for a headbutt in the 2006 World Cup final?",
				Options: []string{"Marco Materazzi", "Zinedine Zidane", "Patrick Vieira", "Thierry Henry"},
				Answer:  1,
				Fact:    "Zidane headbutted Materazzi; Italy won the shootout.",
			},
			{
				Text:    "Who scored the infamous \"Hand of God\" goal at the 1986 World Cup?",
				Options: []string{"Diego Maradona", "Jorge Valdano", "Gary Lineker", "Michel Platini"},
				Answer:  0,
				Fact:    "Maradona, against England in the quarter-final.",
			},
			{
				Text:    "Who holds the record for most goals in a single World Cup — 13 in 1958?",
				Options: []string{"Just Fontaine", "Pelé", "Sándor Kocsis", "Gerd Müller"},
				Answer:  0,
				Fact:    "France's Just Fontaine — a record that still stands.",
			},
			{
				Text:    "Spain beat which country 1–0 in the 2010 World Cup final?",
				Options: []string{"Germany", "Netherlands", "Italy", "Portugal"},
				Answer:  1,
				Fact:    "Spain 1–0 Netherlands after extra time.",
			},
			{
				Text:    "Who scored Spain's extra-time winner in the 2010 final?",
				Options: []string{"Xavi", "David Villa", "Andrés Iniesta", "Fernando Torres"},
				Answer:  2,
				Fact:    "Andrés Iniesta's late strike won Spain their first World Cup.",
			},
			{
				Text:    "How many qualifying places did UEFA (Europe) receive for the 2026 World Cup?",
				Options: []string{"13", "15", "16", "18"},
				Answer:  2,
				Fact:    "16 places — up from 13 at the 2022 edition.",
			},
			{
				Text:    "Which is the only nation to have played at every World Cup since 1930?",
				Options: []string{"Germany", "Italy", "Brazil", "Argentina"},
				Answer:  2,
				Fact:    "Brazil — present at all 23 tournaments.",
			},
			{
				Text:    "Italy won the 2006 World Cup final on penalties against which country?",
				Options: []string{"Germany", "France", "Portugal", "Brazil"},
				Answer:  1,
				Fact:    "Italy beat France 5–3 on penalties after a 1–1 draw.",
			},
			{
				Text:    "Who won the Golden Boot (top scorer) at the 2022 World Cup, with 8 goals?",
				Options: []string{"Lionel Messi", "Kylian Mbappé", "Julián Álvarez", "Olivier Giroud"},
				Answer:  1,
				Fact:    "Mbappé's 8 goals edged Messi's 7.",
			},
			{
				Text:    "Who was named best player (Golden Ball) at the 2018 World Cup?",
				Options: []string{"Kylian Mbappé", "Luka Modrić", "Eden Hazard", "Antoine Griezmann"},
				Answer:  1,
				Fact:    "Croatia's Luka Modrić, despite losing the final.",
			},
			{
				Text:    "Who is the only player to have won three World Cups?",
				Options: []string{"Diego Maradona", "Pelé", "Franz Beckenbauer", "Cafu"},
				Answer:  1,
				Fact:    "Pelé won in 1958, 1962 and 1970 with Brazil.",
			},
			{
				Text:    "Germany's stunning 2014 semi-final win over hosts Brazil finished what score?",
				Options: []string{"5–0", "7–1", "6–2", "4–0"},
				Answer:  1,
				Fact:    "A barely believable 7–1 in Belo Horizonte.",
			},
			{
				Text:    "The 2026 opening match (Mexico 2–0 South Africa) set an unwanted record for what?",
				Options: []string{"Most red cards in a WC match (3)", "Lowest attendance", "Longest match", "Most own goals"},
				Answer:  0,
				Fact:    "Three red cards — the most ever shown in a single World Cup match.",
			},
			{
				Text:    "Mexico City's stadium became the first ground to host how many World Cup opening matches?",
				Options: []string{"Two", "Three", "Four", "One"},
				Answer:  1,
				Fact:    "Three — it also opened the 1970 and 1986 tournaments.",
			},
		},
	},
}

// difficultyOrder is the display order on the landing page.
var difficultyOrder = []string{DiffKids, DiffAdult, DiffNerd}
