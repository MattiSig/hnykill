package main

// translationsIS maps each English question Text to its Icelandic translation.
var translationsIS = map[string]QuestionL10n{
	// ---- Kids ----
	"How many countries are hosting the 2026 World Cup?": {
		Text:    "Hversu mörg lönd halda HM 2026?",
		Options: []string{"1", "2", "3", "4"},
		Fact:    "Þrjú! Kanada, Mexíkó og Bandaríkin halda mótið saman í fyrsta sinn.",
	},
	"What sport is the World Cup all about?": {
		Text:    "Um hvaða íþrótt snýst HM?",
		Options: []string{"Körfubolta", "Fótbolta", "Tennis", "Rugby"},
		Fact:    "Það er stærsta fótboltamót í heimi!",
	},
	"How many players from one team are on the pitch at the same time?": {
		Text:    "Hversu margir leikmenn úr einu liði eru á vellinum á sama tíma?",
		Options: []string{"9", "10", "11", "12"},
		Fact:    "Ellefu leikmenn í hverju liði, þar með talinn markvörðurinn.",
	},
	"Canada's 2026 mascot \"Maple\" is which animal?": {
		Text:    "Lukkudýr Kanada 2026, \"Maple\", er hvaða dýr?",
		Options: []string{"Bjór", "Elgur", "Björn", "Úlfur"},
		Fact:    "Maple er elgur — og áhugasamur markvörður sem elskar list og tónlist!",
	},
	"What colour card does the referee show to send a player off?": {
		Text:    "Hvaða lit á spjaldi sýnir dómarinn til að vísa leikmanni af velli?",
		Options: []string{"Gult", "Grænt", "Rautt", "Blátt"},
		Fact:    "Rautt spjald þýðir að leikmaðurinn verður að yfirgefa leikinn.",
	},
	"How often is the World Cup held?": {
		Text:    "Hversu oft er HM haldið?",
		Options: []string{"Á hverju ári", "Á 2 ára fresti", "Á 4 ára fresti", "Á 10 ára fresti"},
		Fact:    "Á fjögurra ára fresti — svo það er virkilega sérstakur viðburður!",
	},
	"Mexico's 2026 mascot \"Zayu\" is which jungle animal?": {
		Text:    "Lukkudýr Mexíkó 2026, \"Zayu\", er hvaða frumskógardýr?",
		Options: []string{"Ljón", "Tígrisdýr", "Jagúar", "Pardusdýr"},
		Fact:    "Zayu er fljótur jagúar-sóknarmaður úr frumskógum suðurhluta Mexíkó.",
	},
	"What do we call the player who guards the goal?": {
		Text:    "Hvað köllum við leikmanninn sem ver markið?",
		Options: []string{"Sóknarmann", "Varnarmann", "Markvörð", "Kantmann"},
		Fact:    "Markvörðurinn er eini leikmaðurinn sem má nota hendurnar (innan vítateigs).",
	},
	"The USA's 2026 mascot is a bald eagle. What is its name?": {
		Text:    "Lukkudýr Bandaríkjanna 2026 er skallaörn. Hvað heitir hann?",
		Options: []string{"Sammy", "Clutch", "Spike", "Buddy"},
		Fact:    "Clutch skallaörninn er miðjumaður sem elskar að sameina fólk.",
	},
	"What do players try to score to win?": {
		Text:    "Hvað reyna leikmenn að skora til að vinna?",
		Options: []string{"Stig", "Mörk", "Hlaup", "Tilraunir"},
		Fact:    "Mörk! Liðið með flest mörk vinnur.",
	},
	"How many minutes is a normal football match (not counting added time)?": {
		Text:    "Hversu margar mínútur er venjulegur fótboltaleikur (ekki talinn uppbótartími)?",
		Options: []string{"45", "60", "90", "120"},
		Fact:    "90 mínútur — tveir hálfleikir, 45 mínútur hvor.",
	},
	"When a player scores three goals in one match, what is it called?": {
		Text:    "Þegar leikmaður skorar þrjú mörk í einum leik, hvað kallast það?",
		Options: []string{"Þrenna", "Hat trick", "Þrefalt skot", "Tríó"},
		Fact:    "Hat trick! Sumir áhorfendur kasta jafnvel höttunum sínum til að fagna.",
	},
	"Cristiano Ronaldo plays for which country?": {
		Text:    "Fyrir hvaða land leikur Cristiano Ronaldo?",
		Options: []string{"Brasilíu", "Spán", "Portúgal", "Ítalíu"},
		Fact:    "Ronaldo er markahæsti leikmaður Portúgals frá upphafi og fyrirliði.",
	},
	"Lionel Messi plays for which country?": {
		Text:    "Fyrir hvaða land leikur Lionel Messi?",
		Options: []string{"Argentínu", "Brasilíu", "Úrúgvæ", "Síle"},
		Fact:    "Messi var fyrirliði Argentínu til sigurs á HM 2022.",
	},
	"Brazil's football shirts are famously which colour?": {
		Text:    "Fótboltatreyjur Brasilíu eru frægar fyrir hvaða lit?",
		Options: []string{"Rauðan", "Gulan", "Bláan", "Grænan"},
		Fact:    "Skærgula treyjan er ein sú frægasta í fótbolta.",
	},
	"England's national team is nicknamed the \"Three…\" what?": {
		Text:    "Landslið Englands er kallað \"Þrjú …\" hvað?",
		Options: []string{"Bolabítar", "Ljón", "Ernir", "Birnir"},
		Fact:    "Ljónin þrjú — eftir merkinu á treyjum þeirra.",
	},
	"If a knockout match is tied and a winner is needed, how is it decided?": {
		Text:    "Ef útsláttarleikur er jafn og sigurvegara er þörf, hvernig er það ákveðið?",
		Options: []string{"Með peningakasti", "Með vítaspyrnukeppni", "Með endurleik daginn eftir", "Með aukaspyrnukeppni"},
		Fact:    "Eftir framlengingu fer það í taugatrekkjandi vítaspyrnukeppni.",
	},
	"Which body part can a goalkeeper use that other players can't (inside the box)?": {
		Text:    "Hvaða líkamshluta má markvörður nota sem aðrir leikmenn mega ekki (innan vítateigs)?",
		Options: []string{"Höfuðið", "Hendurnar", "Bringuna", "Fæturna"},
		Fact:    "Markverðir mega grípa boltann með höndunum innan eigin vítateigs.",
	},
	"How many points does a team get for winning a group game?": {
		Text:    "Hversu mörg stig fær lið fyrir að vinna riðlaleik?",
		Options: []string{"1", "2", "3", "4"},
		Fact:    "3 stig fyrir sigur, 1 fyrir jafntefli, 0 fyrir tap.",
	},
	"What is the shiny gold prize the winning team lifts called?": {
		Text:    "Hvað kallast glansandi gullverðlaunin sem sigurliðið lyftir?",
		Options: []string{"Bikarinn", "Verðlaunapeningurinn", "Skjöldurinn", "Diskurinn"},
		Fact:    "HM-bikar FIFA — gerður úr skíru gulli!",
	},
	"France's football team is known as \"Les …\" what (after their shirt colour)?": {
		Text:    "Fótboltalið Frakklands er þekkt sem \"Les …\" hvað (eftir lit treyju þeirra)?",
		Options: []string{"Bleus (Bláu)", "Rouges (Rauðu)", "Verts (Grænu)", "Blancs (Hvítu)"},
		Fact:    "\"Les Bleus\" — Bláu — urðu heimsmeistarar 1998 og 2018.",
	},
	"Kylian Mbappé, a superstar striker, plays for which country?": {
		Text:    "Kylian Mbappé, stórstjarna í sókn, leikur fyrir hvaða land?",
		Options: []string{"Frakkland", "Spán", "Brasilíu", "England"},
		Fact:    "Númer 10 hjá Frakklandi — og einn fljótasti leikmaður í heimi.",
	},
	"Erling Haaland scores lots and lots of goals for which country?": {
		Text:    "Erling Haaland skorar fullt og fullt af mörkum fyrir hvaða land?",
		Options: []string{"Svíþjóð", "Noreg", "Danmörku", "Ísland"},
		Fact:    "Noreg! Hann er frægur fyrir að skora mörk í bílförmum.",
	},
	"The Netherlands team is famous for wearing which bright colour?": {
		Text:    "Lið Hollands er frægt fyrir að klæðast hvaða skæra lit?",
		Options: []string{"Appelsínugulum", "Fjólubláum", "Bleikum", "Gráum"},
		Fact:    "Skærappelsínugulum — aðdáendur kalla liðið \"Oranje\".",
	},
	"Which country does the famous player Neymar play for?": {
		Text:    "Fyrir hvaða land leikur hinn frægi leikmaður Neymar?",
		Options: []string{"Portúgal", "Spán", "Brasilíu", "Mexíkó"},
		Fact:    "Neymar er ein af stærstu stórstjörnum Brasilíu.",
	},
	"Argentina, the team of Lionel Messi, wears shirts of which two colours?": {
		Text:    "Argentína, lið Lionels Messi, klæðist treyjum í hvaða tveimur litum?",
		Options: []string{"Ljósbláum og hvítum", "Rauðum og svörtum", "Grænum og gulum", "Alsvörtum"},
		Fact:    "Frægu ljósbláu og hvítu rendurnar hjá Argentínu.",
	},
	"How many times has Germany won the World Cup — none or lots?": {
		Text:    "Hversu oft hefur Þýskaland unnið HM — aldrei eða oft?",
		Options: []string{"Aldrei", "Einu sinni", "Fjórum sinnum", "Tuttugu sinnum"},
		Fact:    "Þýskaland er fjórfaldur heimsmeistari!",
	},

	// ---- Adult ----
	"How many teams compete in the expanded 2026 World Cup?": {
		Text:    "Hversu mörg lið keppa á stækkaðri HM 2026?",
		Options: []string{"32", "40", "48", "64"},
		Fact:    "2026 er fyrsta HM með 48 liðum, upp úr 32.",
	},
	"Which country won the 2022 World Cup in Qatar?": {
		Text:    "Hvaða land vann HM 2022 í Katar?",
		Options: []string{"Frakkland", "Brasilía", "Argentína", "Þýskaland"},
		Fact:    "Argentína vann Frakkland í vítaspyrnukeppni — þeir mæta 2026 sem heimsmeistarar.",
	},
	"The 2026 final at MetLife Stadium will be played in which US state?": {
		Text:    "Úrslitaleikurinn 2026 á MetLife Stadium verður leikinn í hvaða ríki Bandaríkjanna?",
		Options: []string{"New York", "New Jersey", "Kaliforníu", "Flórída"},
		Fact:    "MetLife Stadium er í East Rutherford, New Jersey. Úrslitaleikurinn er 19. júlí 2026.",
	},
	"Mexico hosted the men's World Cup before in which years?": {
		Text:    "Mexíkó hélt HM karla áður á hvaða árum?",
		Options: []string{"1970 og 1986", "1966 og 1994", "1958 og 1982", "1978 og 1990"},
		Fact:    "1970 og 1986 — 2026 gerir Mexíkó að fyrsta gestgjafanum þrisvar sinnum.",
	},
	"Who won the Golden Ball (best player) at the 2022 World Cup?": {
		Text:    "Hver vann Golden Ball (besti leikmaður) á HM 2022?",
		Options: []string{"Kylian Mbappé", "Lionel Messi", "Luka Modrić", "Neymar"},
		Fact:    "Lionel Messi vann Golden Ball þegar hann lyfti loksins bikarnum.",
	},
	"How many matches are played in total at the 2026 World Cup?": {
		Text:    "Hversu margir leikir eru spilaðir í heildina á HM 2026?",
		Options: []string{"64", "80", "104", "128"},
		Fact:    "104 leikir — upp úr 64, þökk sé 48 liða fyrirkomulaginu.",
	},
	"Which iconic stadium hosted the 2026 opening match?": {
		Text:    "Hvaða táknræni leikvangur hýsti opnunarleik HM 2026?",
		Options: []string{"Estadio Azteca", "MetLife Stadium", "Wembley", "SoFi Stadium"},
		Fact:    "Estadio Azteca í Mexíkóborg opnaði mótið 11. júní 2026.",
	},
	"How many groups are there in the 2026 group stage?": {
		Text:    "Hversu margir riðlar eru í riðlakeppni HM 2026?",
		Options: []string{"8", "10", "12", "16"},
		Fact:    "12 riðlar með 4 liðum hver.",
	},
	"Which country has won the most World Cups, with five titles?": {
		Text:    "Hvaða land hefur unnið flesta heimsmeistaratitla, með fimm titla?",
		Options: []string{"Þýskaland", "Ítalía", "Brasilía", "Argentína"},
		Fact:    "Brasilía er efst með fimm (1958, 1962, 1970, 1994, 2002).",
	},
	"The official 2026 Adidas match ball is called what?": {
		Text:    "Hvað heitir opinberi Adidas-leikbolti HM 2026?",
		Options: []string{"Al Rihla", "Trionda", "Brazuca", "Jabulani"},
		Fact:    "\"Trionda\" þýðir \"þrjár öldur\" og fagnar gestgjafaþjóðunum þremur.",
	},
	"Who won the 2018 World Cup in Russia?": {
		Text:    "Hver vann HM 2018 í Rússlandi?",
		Options: []string{"Króatía", "Frakkland", "Belgía", "England"},
		Fact:    "Frakkland vann Króatíu 4–2 í úrslitaleiknum.",
	},
	"Which country did France beat in the 2018 World Cup final?": {
		Text:    "Hvaða land sigraði Frakkland í úrslitaleik HM 2018?",
		Options: []string{"Króatíu", "Belgíu", "England", "Brasilíu"},
		Fact:    "Frakkland 4–2 Króatía í Moskvu — fyrsti úrslitaleikur Króatíu nokkurn tímann.",
	},
	"Who won the 2014 World Cup, held in Brazil?": {
		Text:    "Hver vann HM 2014, sem haldið var í Brasilíu?",
		Options: []string{"Argentína", "Holland", "Þýskaland", "Spánn"},
		Fact:    "Þýskaland vann Argentínu 1–0 eftir framlengingu.",
	},
	"Which country won the 2010 World Cup — their first ever title?": {
		Text:    "Hvaða land vann HM 2010 — sinn fyrsta titil nokkurn tímann?",
		Options: []string{"Holland", "Spánn", "Þýskaland", "Ítalía"},
		Fact:    "Spánn vann Holland 1–0 í framlengingu í Suður-Afríku.",
	},
	"How many European (UEFA) teams qualified for the 2026 World Cup?": {
		Text:    "Hversu mörg evrópsk (UEFA) lið komust á HM 2026?",
		Options: []string{"13", "16", "20", "24"},
		Fact:    "16 — stærsta HM-úthlutun UEFA frá upphafi.",
	},
	"Kylian Mbappé plays for which national team?": {
		Text:    "Fyrir hvaða landslið leikur Kylian Mbappé?",
		Options: []string{"Belgíu", "Frakkland", "Sviss", "Marokkó"},
		Fact:    "Fyrirliði og lykilmaður Frakklands.",
	},
	"Who won the 2024 men's Ballon d'Or?": {
		Text:    "Hver vann Ballon d'Or karla 2024?",
		Options: []string{"Vinícius Júnior", "Jude Bellingham", "Rodri", "Erling Haaland"},
		Fact:    "Rodri, miðjumaður Spánar og Manchester City, hafði betur en Vinícius Júnior.",
	},
	"Which country won Euro 2024?": {
		Text:    "Hvaða land vann EM 2024?",
		Options: []string{"England", "Spánn", "Frakkland", "Þýskaland"},
		Fact:    "Spánn vann England 2–1 í Berlín og hlaut fjórða EM-titilinn, sem er met.",
	},
	"Mexico opened the 2026 World Cup against which nation?": {
		Text:    "Mexíkó opnaði HM 2026 gegn hvaða þjóð?",
		Options: []string{"Suður-Afríku", "Marokkó", "Bandaríkjunum", "Suður-Kóreu"},
		Fact:    "Mexíkó vann Suður-Afríku 2–0 í opnunarleiknum.",
	},
	"Who scored a hat-trick in the 2022 World Cup final but still finished as a runner-up?": {
		Text:    "Hver skoraði hat trick í úrslitaleik HM 2022 en endaði samt í öðru sæti?",
		Options: []string{"Lionel Messi", "Kylian Mbappé", "Ángel Di María", "Olivier Giroud"},
		Fact:    "Mbappé skoraði þrjú en Frakkland tapaði vítaspyrnukeppninni.",
	},
	"Which country shares second place for most World Cup titles (four) with Germany?": {
		Text:    "Hvaða land deilir öðru sæti yfir flesta heimsmeistaratitla (fjóra) með Þýskalandi?",
		Options: []string{"Ítalía", "Argentína", "Frakkland", "Úrúgvæ"},
		Fact:    "Ítalía á líka fjóra titla (1934, 1938, 1982, 2006).",
	},
	"Which club has won the European Cup / Champions League the most times?": {
		Text:    "Hvaða félag hefur unnið Evrópukeppni meistaraliða / Champions League oftast?",
		Options: []string{"Barcelona", "AC Milan", "Bayern Munich", "Real Madrid"},
		Fact:    "Real Madrid — langsigursælasta félagið í Evrópu.",
	},
	"Erling Haaland, one of the world's deadliest strikers, plays for which country?": {
		Text:    "Erling Haaland, einn skæðasti sóknarmaður heims, leikur fyrir hvaða land?",
		Options: []string{"Svíþjóð", "Noreg", "Danmörku", "Holland"},
		Fact:    "Noreg — sem hefur sjaldan komist á stórmót þrátt fyrir hæfileika hans.",
	},
	"Which player has won the most Ballon d'Or awards, with eight?": {
		Text:    "Hvaða leikmaður hefur unnið flest Ballon d'Or-verðlaun, með átta?",
		Options: []string{"Cristiano Ronaldo", "Lionel Messi", "Michel Platini", "Johan Cruyff"},
		Fact:    "Messi á átta; Ronaldo er næstur með fimm.",
	},
	"England midfield star Jude Bellingham plays his club football for which team?": {
		Text:    "Enska miðjustjarnan Jude Bellingham leikur félagsfótbolta með hvaða liði?",
		Options: []string{"Manchester City", "Real Madrid", "Liverpool", "Bayern Munich"},
		Fact:    "Real Madrid, sem keypti hann frá Borussia Dortmund árið 2023.",
	},
	"England's top men's football division is called what?": {
		Text:    "Hvað heitir efsta deild karla í fótbolta á Englandi?",
		Options: []string{"La Liga", "Serie A", "Premier League", "Bundesliga"},
		Fact:    "Premier League — mest áhorfaða deild í heimi.",
	},
	"Since 2024, Kylian Mbappé plays his club football for which team?": {
		Text:    "Frá 2024 leikur Kylian Mbappé félagsfótbolta með hvaða liði?",
		Options: []string{"Paris Saint-Germain", "Real Madrid", "Liverpool", "Al-Nassr"},
		Fact:    "Hann gekk til liðs við Real Madrid frá Paris Saint-Germain árið 2024.",
	},
	"Italy's national team is nicknamed the what?": {
		Text:    "Landslið Ítalíu er kallað hvað?",
		Options: []string{"Azzurri", "Oranje", "Les Bleus", "Seleção"},
		Fact:    "\"Gli Azzurri\" — Bláu — eftir himinbláu treyjum þeirra.",
	},
	"Harry Kane, England's record goalscorer, plays his club football for which team?": {
		Text:    "Harry Kane, markahæsti leikmaður Englands frá upphafi, leikur félagsfótbolta með hvaða liði?",
		Options: []string{"Tottenham", "Bayern Munich", "Manchester United", "Chelsea"},
		Fact:    "Hann gekk til liðs við Bayern Munich frá Tottenham árið 2023.",
	},

	// ---- Nerd ----
	"The 48-team format adds a brand-new knockout round. What is it called?": {
		Text:    "48 liða fyrirkomulagið bætir við glænýrri útsláttarumferð. Hvað heitir hún?",
		Options: []string{"16-liða úrslit", "32-liða úrslit", "Forkeppniumferð", "Síðustu 40"},
		Fact:    "32-liða úrslit birtast á HM í fyrsta sinn árið 2026.",
	},
	"How many host cities stage matches across the three 2026 nations?": {
		Text:    "Hversu margar gestgjafaborgir hýsa leiki í þjóðunum þremur árið 2026?",
		Options: []string{"11", "14", "16", "23"},
		Fact:    "16 gestgjafaborgir: 11 í Bandaríkjunum, 3 í Mexíkó og 2 í Kanada.",
	},
	"How many best third-placed teams advance from the 2026 group stage?": {
		Text:    "Hversu mörg bestu liðin í þriðja sæti komast áfram úr riðlakeppni HM 2026?",
		Options: []string{"4", "6", "8", "12"},
		Fact:    "Tvö efstu úr hverjum riðli auk 8 bestu liða í þriðja sæti komast í 32-liða úrslit.",
	},
	"Which of these is NOT a 2026 host city in Mexico?": {
		Text:    "Hver af þessum er EKKI gestgjafaborg í Mexíkó árið 2026?",
		Options: []string{"Guadalajara", "Monterrey", "Cancún", "Mexíkóborg"},
		Fact:    "Gestgjafaborgir Mexíkó eru Mexíkóborg, Guadalajara og Monterrey.",
	},
	"The match-ball name \"Trionda\" (\"three waves\") celebrates what?": {
		Text:    "Nafn leikboltans \"Trionda\" (\"þrjár öldur\") fagnar hverju?",
		Options: []string{"Gestgjafaþjóðunum þremur", "Þremur stigum fyrir sigur", "Þremur ljónum", "Þremur dómurum"},
		Fact:    "Það heiðrar Kanada, Mexíkó og Bandaríkin — fyrstu þriggja þjóða gestgjafana.",
	},
	"In the 2022 final, Mbappé scored a hat-trick but France lost on penalties. What was the score after extra time?": {
		Text:    "Í úrslitaleiknum 2022 skoraði Mbappé hat trick en Frakkland tapaði í vítaspyrnukeppni. Hver var staðan eftir framlengingu?",
		Options: []string{"2–2", "3–3", "4–4", "3–2"},
		Fact:    "Það endaði 3–3; Argentína vann 4–2 í vítaspyrnukeppni.",
	},
	"Which two Canadian cities are 2026 host cities?": {
		Text:    "Hvaða tvær kanadískar borgir eru gestgjafaborgir árið 2026?",
		Options: []string{"Toronto og Vancouver", "Montreal og Calgary", "Ottawa og Edmonton", "Toronto og Montreal"},
		Fact:    "Toronto og Vancouver eru gestgjafaborgir Kanada tvær.",
	},
	"Who is the World Cup's all-time leading goalscorer, with 16 goals?": {
		Text:    "Hver er markahæsti leikmaður HM frá upphafi, með 16 mörk?",
		Options: []string{"Ronaldo (Brasilía)", "Miroslav Klose", "Gerd Müller", "Just Fontaine"},
		Fact:    "Klose frá Þýskalandi (16) fór fram úr Ronaldo frá Brasilíu (15) árið 2014.",
	},
	"On which date is the 2026 World Cup final played?": {
		Text:    "Hvaða dagsetningu er úrslitaleikur HM 2026 leikinn?",
		Options: []string{"4. júlí", "12. júlí", "19. júlí", "26. júlí"},
		Fact:    "19. júlí 2026 á MetLife Stadium.",
	},
	"Who scored the only hat-trick in a World Cup final, for England in 1966?": {
		Text:    "Hver skoraði eina hat trickið í úrslitaleik HM, fyrir England árið 1966?",
		Options: []string{"Bobby Charlton", "Geoff Hurst", "Martin Peters", "Gary Lineker"},
		Fact:    "Geoff Hurst, í 4–2 sigri Englands á Vestur-Þýskalandi.",
	},
	"Who scored Germany's extra-time winner in the 2014 World Cup final?": {
		Text:    "Hver skoraði sigurmark Þýskalands í framlengingu í úrslitaleik HM 2014?",
		Options: []string{"Thomas Müller", "Mario Götze", "André Schürrle", "Bastian Schweinsteiger"},
		Fact:    "Glæsilegt vítaspyrnumark Marios Götze tryggði 1–0 sigur á Argentínu.",
	},
	"Which player was sent off for a headbutt in the 2006 World Cup final?": {
		Text:    "Hvaða leikmaður var rekinn af velli fyrir höfuðhögg í úrslitaleik HM 2006?",
		Options: []string{"Marco Materazzi", "Zinedine Zidane", "Patrick Vieira", "Thierry Henry"},
		Fact:    "Zidane skallaði Materazzi; Ítalía vann vítaspyrnukeppnina.",
	},
	"Who scored the infamous \"Hand of God\" goal at the 1986 World Cup?": {
		Text:    "Hver skoraði hið alræmda \"Hönd Guðs\" mark á HM 1986?",
		Options: []string{"Diego Maradona", "Jorge Valdano", "Gary Lineker", "Michel Platini"},
		Fact:    "Maradona, gegn Englandi í 8-liða úrslitum.",
	},
	"Who holds the record for most goals in a single World Cup — 13 in 1958?": {
		Text:    "Hver á metið yfir flest mörk á einni HM — 13 árið 1958?",
		Options: []string{"Just Fontaine", "Pelé", "Sándor Kocsis", "Gerd Müller"},
		Fact:    "Just Fontaine frá Frakklandi — met sem stendur enn.",
	},
	"Spain beat which country 1–0 in the 2010 World Cup final?": {
		Text:    "Spánn vann hvaða land 1–0 í úrslitaleik HM 2010?",
		Options: []string{"Þýskaland", "Holland", "Ítalíu", "Portúgal"},
		Fact:    "Spánn 1–0 Holland eftir framlengingu.",
	},
	"Who scored Spain's extra-time winner in the 2010 final?": {
		Text:    "Hver skoraði sigurmark Spánar í framlengingu í úrslitaleiknum 2010?",
		Options: []string{"Xavi", "David Villa", "Andrés Iniesta", "Fernando Torres"},
		Fact:    "Síðbúið skot Andrésar Iniesta tryggði Spáni sinn fyrsta heimsmeistaratitil.",
	},
	"How many qualifying places did UEFA (Europe) receive for the 2026 World Cup?": {
		Text:    "Hversu mörg sæti fékk UEFA (Evrópa) til að komast á HM 2026?",
		Options: []string{"13", "15", "16", "18"},
		Fact:    "16 sæti — upp úr 13 á mótinu 2022.",
	},
	"Which is the only nation to have played at every World Cup since 1930?": {
		Text:    "Hvaða þjóð er sú eina sem hefur leikið á öllum heimsmeistaramótum frá 1930?",
		Options: []string{"Þýskaland", "Ítalía", "Brasilía", "Argentína"},
		Fact:    "Brasilía — viðstödd öll 23 mótin.",
	},
	"Italy won the 2006 World Cup final on penalties against which country?": {
		Text:    "Ítalía vann úrslitaleik HM 2006 í vítaspyrnukeppni gegn hvaða landi?",
		Options: []string{"Þýskalandi", "Frakklandi", "Portúgal", "Brasilíu"},
		Fact:    "Ítalía vann Frakkland 5–3 í vítaspyrnukeppni eftir 1–1 jafntefli.",
	},
	"Who won the Golden Boot (top scorer) at the 2022 World Cup, with 8 goals?": {
		Text:    "Hver vann Golden Boot (markahæstur) á HM 2022, með 8 mörk?",
		Options: []string{"Lionel Messi", "Kylian Mbappé", "Julián Álvarez", "Olivier Giroud"},
		Fact:    "8 mörk Mbappé höfðu betur en 7 mörk Messi.",
	},
	"Who was named best player (Golden Ball) at the 2018 World Cup?": {
		Text:    "Hver var valinn besti leikmaður (Golden Ball) á HM 2018?",
		Options: []string{"Kylian Mbappé", "Luka Modrić", "Eden Hazard", "Antoine Griezmann"},
		Fact:    "Luka Modrić frá Króatíu, þrátt fyrir að tapa úrslitaleiknum.",
	},
	"Who is the only player to have won three World Cups?": {
		Text:    "Hver er eini leikmaðurinn sem hefur unnið þrjú heimsmeistaramót?",
		Options: []string{"Diego Maradona", "Pelé", "Franz Beckenbauer", "Cafu"},
		Fact:    "Pelé vann 1958, 1962 og 1970 með Brasilíu.",
	},
	"Germany's stunning 2014 semi-final win over hosts Brazil finished what score?": {
		Text:    "Stórkostlegur sigur Þýskalands á gestgjöfunum Brasilíu í undanúrslitum 2014 endaði með hvaða stöðu?",
		Options: []string{"5–0", "7–1", "6–2", "4–0"},
		Fact:    "Vart trúlegt 7–1 í Belo Horizonte.",
	},
	"The 2026 opening match (Mexico 2–0 South Africa) set an unwanted record for what?": {
		Text:    "Opnunarleikur HM 2026 (Mexíkó 2–0 Suður-Afríka) setti óvelkomið met fyrir hvað?",
		Options: []string{"Flest rauð spjöld í HM-leik (3)", "Minnsta aðsókn", "Lengsta leik", "Flest sjálfsmörk"},
		Fact:    "Þrjú rauð spjöld — þau flestu sem sýnd hafa verið í einum HM-leik.",
	},
	"Mexico City's stadium became the first ground to host how many World Cup opening matches?": {
		Text:    "Leikvangur Mexíkóborgar varð fyrsti völlurinn til að hýsa hversu marga opnunarleiki HM?",
		Options: []string{"Tvo", "Þrjá", "Fjóra", "Einn"},
		Fact:    "Þrjá — hann opnaði líka mótin 1970 og 1986.",
	},
	"Which nation became the first African team to reach a World Cup semi-final, in 2022?": {
		Text:    "Hvaða þjóð varð fyrsta afríska liðið til að komast í undanúrslit HM, árið 2022?",
		Options: []string{"Senegal", "Marokkó", "Kamerún", "Gana"},
		Fact:    "Marokkó vann Spán og Portúgal á leiðinni í undanúrslit.",
	},
	"Who is Brazil's all-time leading World Cup goalscorer, with 15 goals?": {
		Text:    "Hver er markahæsti leikmaður Brasilíu á HM frá upphafi, með 15 mörk?",
		Options: []string{"Pelé", "Ronaldo (Nazário)", "Romário", "Neymar"},
		Fact:    "Ronaldo \"O Fenômeno\" skoraði 15 HM-mörk fyrir Brasilíu.",
	},
	"Which goalkeeper captained Spain to the 2010 World Cup title?": {
		Text:    "Hvaða markvörður var fyrirliði Spánar til heimsmeistaratitilsins 2010?",
		Options: []string{"Iker Casillas", "Víctor Valdés", "Pepe Reina", "David de Gea"},
		Fact:    "Iker Casillas lyfti bikarnum bæði sem fyrirliði og markvörður.",
	},
}
