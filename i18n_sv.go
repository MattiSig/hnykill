package main

// translationsSV maps each English question Text to its Swedish translation.
var translationsSV = map[string]QuestionL10n{
	"How many countries are hosting the 2026 World Cup?": {
		Text:    "Hur många länder är värdar för VM 2026?",
		Options: []string{"1", "2", "3", "4"},
		Fact:    "Tre! Kanada, Mexiko och USA är värdar tillsammans för allra första gången.",
	},
	"What sport is the World Cup all about?": {
		Text:    "Vilken sport handlar VM om?",
		Options: []string{"Basket", "Fotboll (Soccer)", "Tennis", "Rugby"},
		Fact:    "Det är världens största fotbollsturnering!",
	},
	"How many players from one team are on the pitch at the same time?": {
		Text:    "Hur många spelare från ett lag är på planen samtidigt?",
		Options: []string{"9", "10", "11", "12"},
		Fact:    "Elva spelare per lag, inklusive målvakten.",
	},
	"Canada's 2026 mascot \"Maple\" is which animal?": {
		Text:    "Kanadas maskot 2026 \"Maple\" är vilket djur?",
		Options: []string{"Bäver", "Älg", "Björn", "Varg"},
		Fact:    "Maple är en älg — och en duktig målvakt som älskar konst och musik!",
	},
	"What colour card does the referee show to send a player off?": {
		Text:    "Vilket färgkort visar domaren för att utvisa en spelare?",
		Options: []string{"Gult", "Grönt", "Rött", "Blått"},
		Fact:    "Ett rött kort betyder att spelaren måste lämna matchen.",
	},
	"How often is the World Cup held?": {
		Text:    "Hur ofta hålls VM?",
		Options: []string{"Varje år", "Vartannat år", "Vart fjärde år", "Vart tionde år"},
		Fact:    "Vart fjärde år — så det är en riktigt speciell händelse!",
	},
	"Mexico's 2026 mascot \"Zayu\" is which jungle animal?": {
		Text:    "Mexikos maskot 2026 \"Zayu\" är vilket djungeldjur?",
		Options: []string{"Lejon", "Tiger", "Jaguar", "Panter"},
		Fact:    "Zayu är en snabb jaguar-anfallare från djunglerna i södra Mexiko.",
	},
	"What do we call the player who guards the goal?": {
		Text:    "Vad kallar vi spelaren som vaktar målet?",
		Options: []string{"Anfallare", "Försvarare", "Målvakt", "Ytter"},
		Fact:    "Målvakten är den enda spelaren som får använda händerna (i straffområdet).",
	},
	"The USA's 2026 mascot is a bald eagle. What is its name?": {
		Text:    "USA:s maskot 2026 är en vithövdad havsörn. Vad heter den?",
		Options: []string{"Sammy", "Clutch", "Spike", "Buddy"},
		Fact:    "Clutch havsörnen är en mittfältare som älskar att föra människor samman.",
	},
	"What do players try to score to win?": {
		Text:    "Vad försöker spelarna göra för att vinna?",
		Options: []string{"Poäng", "Mål", "Runs", "Tries"},
		Fact:    "Mål! Laget med flest mål vinner.",
	},
	"How many minutes is a normal football match (not counting added time)?": {
		Text:    "Hur många minuter är en vanlig fotbollsmatch (utan tilläggstid)?",
		Options: []string{"45", "60", "90", "120"},
		Fact:    "90 minuter — två halvlekar på 45 minuter vardera.",
	},
	"When a player scores three goals in one match, what is it called?": {
		Text:    "När en spelare gör tre mål i en match, vad kallas det?",
		Options: []string{"En trippel", "En hat-trick", "Ett trippelskott", "En trio"},
		Fact:    "En hat-trick! Vissa fans kastar till och med sina hattar för att fira.",
	},
	"Cristiano Ronaldo plays for which country?": {
		Text:    "Cristiano Ronaldo spelar för vilket land?",
		Options: []string{"Brasilien", "Spanien", "Portugal", "Italien"},
		Fact:    "Ronaldo är Portugals genom tiderna bäste målskytt och kapten.",
	},
	"Lionel Messi plays for which country?": {
		Text:    "Lionel Messi spelar för vilket land?",
		Options: []string{"Argentina", "Brasilien", "Uruguay", "Chile"},
		Fact:    "Messi var kapten när Argentina vann VM 2022.",
	},
	"Brazil's football shirts are famously which colour?": {
		Text:    "Brasiliens fotbollströjor är berömt vilken färg?",
		Options: []string{"Röda", "Gula", "Blå", "Gröna"},
		Fact:    "Den klargula tröjan är en av de mest berömda inom fotbollen.",
	},
	"England's national team is nicknamed the \"Three…\" what?": {
		Text:    "Englands landslag har smeknamnet \"Three…\" vad?",
		Options: []string{"Bulldogs", "Lions", "Eagles", "Bears"},
		Fact:    "The Three Lions — från märket på deras tröjor.",
	},
	"If a knockout match is tied and a winner is needed, how is it decided?": {
		Text:    "Om en slutspelsmatch är oavgjord och en vinnare behövs, hur avgörs det?",
		Options: []string{"Slantsingling", "Straffläggning", "Omspel nästa dag", "En frisparkstävling"},
		Fact:    "Efter förlängning går det till en nervkittlande straffläggning.",
	},
	"Which body part can a goalkeeper use that other players can't (inside the box)?": {
		Text:    "Vilken kroppsdel får en målvakt använda som andra spelare inte får (inom straffområdet)?",
		Options: []string{"Huvudet", "Händerna", "Bröstet", "Fötterna"},
		Fact:    "Målvakter får fånga bollen med händerna inom sitt eget straffområde.",
	},
	"How many points does a team get for winning a group game?": {
		Text:    "Hur många poäng får ett lag för att vinna en gruppspelsmatch?",
		Options: []string{"1", "2", "3", "4"},
		Fact:    "3 poäng för en vinst, 1 för oavgjort, 0 för en förlust.",
	},
	"What is the shiny gold prize the winning team lifts called?": {
		Text:    "Vad kallas det blanka guldpriset som det vinnande laget lyfter?",
		Options: []string{"Pokalen (Bucklan)", "Medaljen", "Skölden", "Plåten"},
		Fact:    "FIFA World Cup-pokalen — gjord av massivt guld!",
	},
	"France's football team is known as \"Les …\" what (after their shirt colour)?": {
		Text:    "Frankrikes fotbollslag är känt som \"Les …\" vad (efter deras tröjfärg)?",
		Options: []string{"Bleus (Blå)", "Rouges (Röda)", "Verts (Gröna)", "Blancs (Vita)"},
		Fact:    "\"Les Bleus\" — de blå — var världsmästare 1998 och 2018.",
	},
	"Kylian Mbappé, a superstar striker, plays for which country?": {
		Text:    "Kylian Mbappé, en superstjärneanfallare, spelar för vilket land?",
		Options: []string{"Frankrike", "Spanien", "Brasilien", "England"},
		Fact:    "Frankrikes nummer 10 — och en av världens snabbaste spelare.",
	},
	"Erling Haaland scores lots and lots of goals for which country?": {
		Text:    "Erling Haaland gör massor av mål för vilket land?",
		Options: []string{"Sverige", "Norge", "Danmark", "Island"},
		Fact:    "Norge! Han är berömd för att göra mål i mängder.",
	},
	"The Netherlands team is famous for wearing which bright colour?": {
		Text:    "Nederländernas lag är berömt för att bära vilken klar färg?",
		Options: []string{"Orange", "Lila", "Rosa", "Grå"},
		Fact:    "Klarorange — fans kallar laget för \"Oranje\".",
	},
	"Which country does the famous player Neymar play for?": {
		Text:    "Vilket land spelar den berömda spelaren Neymar för?",
		Options: []string{"Portugal", "Spanien", "Brasilien", "Mexiko"},
		Fact:    "Neymar är en av Brasiliens största superstjärnor.",
	},
	"Argentina, the team of Lionel Messi, wears shirts of which two colours?": {
		Text:    "Argentina, Lionel Messis lag, bär tröjor i vilka två färger?",
		Options: []string{"Ljusblått och vitt", "Rött och svart", "Grönt och gult", "Helt svart"},
		Fact:    "Argentinas berömda ljusblå-och-vita ränder.",
	},
	"How many times has Germany won the World Cup — none or lots?": {
		Text:    "Hur många gånger har Tyskland vunnit VM — inga eller många?",
		Options: []string{"Aldrig", "En gång", "Fyra gånger", "Tjugo gånger"},
		Fact:    "Tyskland är fyrfaldiga världsmästare!",
	},
	"How many teams compete in the expanded 2026 World Cup?": {
		Text:    "Hur många lag tävlar i det utökade VM 2026?",
		Options: []string{"32", "40", "48", "64"},
		Fact:    "2026 är det första VM med 48 lag, upp från 32.",
	},
	"Which country won the 2022 World Cup in Qatar?": {
		Text:    "Vilket land vann VM 2022 i Qatar?",
		Options: []string{"Frankrike", "Brasilien", "Argentina", "Tyskland"},
		Fact:    "Argentina slog Frankrike på straffar — de går in i 2026 som mästare.",
	},
	"The 2026 final at MetLife Stadium will be played in which US state?": {
		Text:    "VM-finalen 2026 på MetLife Stadium spelas i vilken amerikansk delstat?",
		Options: []string{"New York", "New Jersey", "Kalifornien", "Florida"},
		Fact:    "MetLife Stadium ligger i East Rutherford, New Jersey. Finalen är den 19 juli 2026.",
	},
	"Mexico hosted the men's World Cup before in which years?": {
		Text:    "Mexiko har varit värd för herrarnas VM tidigare vilka år?",
		Options: []string{"1970 & 1986", "1966 & 1994", "1958 & 1982", "1978 & 1990"},
		Fact:    "1970 och 1986 — 2026 gör Mexiko till den första trefaldiga värden.",
	},
	"Who won the Golden Ball (best player) at the 2022 World Cup?": {
		Text:    "Vem vann Golden Ball (bäste spelare) vid VM 2022?",
		Options: []string{"Kylian Mbappé", "Lionel Messi", "Luka Modrić", "Neymar"},
		Fact:    "Lionel Messi vann Golden Ball när han äntligen lyfte pokalen.",
	},
	"How many matches are played in total at the 2026 World Cup?": {
		Text:    "Hur många matcher spelas totalt vid VM 2026?",
		Options: []string{"64", "80", "104", "128"},
		Fact:    "104 matcher — upp från 64, tack vare formatet med 48 lag.",
	},
	"Which iconic stadium hosted the 2026 opening match?": {
		Text:    "Vilket ikoniskt stadion var värd för öppningsmatchen 2026?",
		Options: []string{"Estadio Azteca", "MetLife Stadium", "Wembley", "SoFi Stadium"},
		Fact:    "Mexico Citys Estadio Azteca öppnade turneringen den 11 juni 2026.",
	},
	"How many groups are there in the 2026 group stage?": {
		Text:    "Hur många grupper finns det i gruppspelet 2026?",
		Options: []string{"8", "10", "12", "16"},
		Fact:    "12 grupper med 4 lag vardera.",
	},
	"Which country has won the most World Cups, with five titles?": {
		Text:    "Vilket land har vunnit flest VM, med fem titlar?",
		Options: []string{"Tyskland", "Italien", "Brasilien", "Argentina"},
		Fact:    "Brasilien leder med fem (1958, 1962, 1970, 1994, 2002).",
	},
	"The official 2026 Adidas match ball is called what?": {
		Text:    "Vad heter den officiella Adidas-matchbollen 2026?",
		Options: []string{"Al Rihla", "Trionda", "Brazuca", "Jabulani"},
		Fact:    "\"Trionda\" betyder \"tre vågor\", och hyllar de tre värdnationerna.",
	},
	"Who won the 2018 World Cup in Russia?": {
		Text:    "Vem vann VM 2018 i Ryssland?",
		Options: []string{"Kroatien", "Frankrike", "Belgien", "England"},
		Fact:    "Frankrike slog Kroatien 4–2 i finalen.",
	},
	"Which country did France beat in the 2018 World Cup final?": {
		Text:    "Vilket land slog Frankrike i VM-finalen 2018?",
		Options: []string{"Kroatien", "Belgien", "England", "Brasilien"},
		Fact:    "Frankrike 4–2 Kroatien i Moskva — Kroatiens första final någonsin.",
	},
	"Who won the 2014 World Cup, held in Brazil?": {
		Text:    "Vem vann VM 2014, som hölls i Brasilien?",
		Options: []string{"Argentina", "Nederländerna", "Tyskland", "Spanien"},
		Fact:    "Tyskland slog Argentina 1–0 efter förlängning.",
	},
	"Which country won the 2010 World Cup — their first ever title?": {
		Text:    "Vilket land vann VM 2010 — deras första titel någonsin?",
		Options: []string{"Nederländerna", "Spanien", "Tyskland", "Italien"},
		Fact:    "Spanien slog Nederländerna 1–0 efter förlängning i Sydafrika.",
	},
	"How many European (UEFA) teams qualified for the 2026 World Cup?": {
		Text:    "Hur många europeiska (UEFA) lag kvalificerade sig till VM 2026?",
		Options: []string{"13", "16", "20", "24"},
		Fact:    "16 — UEFA:s största VM-tilldelning någonsin.",
	},
	"Kylian Mbappé plays for which national team?": {
		Text:    "Kylian Mbappé spelar för vilket landslag?",
		Options: []string{"Belgien", "Frankrike", "Schweiz", "Marocko"},
		Fact:    "Frankrikes kapten och galjonsfigur.",
	},
	"Who won the 2024 men's Ballon d'Or?": {
		Text:    "Vem vann herrarnas Ballon d'Or 2024?",
		Options: []string{"Vinícius Júnior", "Jude Bellingham", "Rodri", "Erling Haaland"},
		Fact:    "Rodri, mittfältaren från Spanien och Manchester City, knep den före Vinícius Júnior.",
	},
	"Which country won Euro 2024?": {
		Text:    "Vilket land vann EM 2024?",
		Options: []string{"England", "Spanien", "Frankrike", "Tyskland"},
		Fact:    "Spanien slog England 2–1 i Berlin för en rekordfjärde EM-titel.",
	},
	"Mexico opened the 2026 World Cup against which nation?": {
		Text:    "Mexiko inledde VM 2026 mot vilken nation?",
		Options: []string{"Sydafrika", "Marocko", "USA", "Sydkorea"},
		Fact:    "Mexiko slog Sydafrika 2–0 i öppningsmatchen.",
	},
	"Who scored a hat-trick in the 2022 World Cup final but still finished as a runner-up?": {
		Text:    "Vem gjorde en hat-trick i VM-finalen 2022 men slutade ändå som tvåa?",
		Options: []string{"Lionel Messi", "Kylian Mbappé", "Ángel Di María", "Olivier Giroud"},
		Fact:    "Mbappé gjorde tre mål men Frankrike förlorade straffläggningen.",
	},
	"Which country shares second place for most World Cup titles (four) with Germany?": {
		Text:    "Vilket land delar andraplatsen för flest VM-titlar (fyra) med Tyskland?",
		Options: []string{"Italien", "Argentina", "Frankrike", "Uruguay"},
		Fact:    "Italien har också fyra titlar (1934, 1938, 1982, 2006).",
	},
	"Which club has won the European Cup / Champions League the most times?": {
		Text:    "Vilken klubb har vunnit Europacupen / Champions League flest gånger?",
		Options: []string{"Barcelona", "AC Milan", "Bayern München", "Real Madrid"},
		Fact:    "Real Madrid — den utan tvekan mest meriterade klubben i Europa.",
	},
	"Erling Haaland, one of the world's deadliest strikers, plays for which country?": {
		Text:    "Erling Haaland, en av världens farligaste anfallare, spelar för vilket land?",
		Options: []string{"Sverige", "Norge", "Danmark", "Nederländerna"},
		Fact:    "Norge — som sällan har nått stora turneringar trots hans talang.",
	},
	"Which player has won the most Ballon d'Or awards, with eight?": {
		Text:    "Vilken spelare har vunnit flest Ballon d'Or-utmärkelser, med åtta?",
		Options: []string{"Cristiano Ronaldo", "Lionel Messi", "Michel Platini", "Johan Cruyff"},
		Fact:    "Messi har åtta; Ronaldo är näst med fem.",
	},
	"England midfield star Jude Bellingham plays his club football for which team?": {
		Text:    "Englands mittfältsstjärna Jude Bellingham spelar klubbfotboll för vilket lag?",
		Options: []string{"Manchester City", "Real Madrid", "Liverpool", "Bayern München"},
		Fact:    "Real Madrid, som värvade honom från Borussia Dortmund 2023.",
	},
	"England's top men's football division is called what?": {
		Text:    "Vad heter Englands högsta herrfotbollsdivision?",
		Options: []string{"La Liga", "Serie A", "Premier League", "Bundesliga"},
		Fact:    "Premier League — den mest sedda ligan i världen.",
	},
	"Since 2024, Kylian Mbappé plays his club football for which team?": {
		Text:    "Sedan 2024 spelar Kylian Mbappé klubbfotboll för vilket lag?",
		Options: []string{"Paris Saint-Germain", "Real Madrid", "Liverpool", "Al-Nassr"},
		Fact:    "Han gick till Real Madrid från Paris Saint-Germain 2024.",
	},
	"Italy's national team is nicknamed the what?": {
		Text:    "Vad har Italiens landslag för smeknamn?",
		Options: []string{"Azzurri", "Oranje", "Les Bleus", "Seleção"},
		Fact:    "\"Gli Azzurri\" — de blå — efter deras himmelsblå tröjor.",
	},
	"Harry Kane, England's record goalscorer, plays his club football for which team?": {
		Text:    "Harry Kane, Englands meste målskytt, spelar klubbfotboll för vilket lag?",
		Options: []string{"Tottenham", "Bayern München", "Manchester United", "Chelsea"},
		Fact:    "Han gick till Bayern München från Tottenham 2023.",
	},
	"The 48-team format adds a brand-new knockout round. What is it called?": {
		Text:    "Formatet med 48 lag lägger till en helt ny slutspelsrunda. Vad kallas den?",
		Options: []string{"Sextondelsfinal", "Trettiotvåondelsfinal", "Play-in-runda", "Sista 40"},
		Fact:    "Trettiotvåondelsfinalen (Round of 32) förekommer i ett VM för första gången 2026.",
	},
	"How many host cities stage matches across the three 2026 nations?": {
		Text:    "Hur många värdstäder arrangerar matcher i de tre 2026-nationerna?",
		Options: []string{"11", "14", "16", "23"},
		Fact:    "16 värdstäder: 11 i USA, 3 i Mexiko och 2 i Kanada.",
	},
	"How many best third-placed teams advance from the 2026 group stage?": {
		Text:    "Hur många bästa trea går vidare från gruppspelet 2026?",
		Options: []string{"4", "6", "8", "12"},
		Fact:    "De två bästa i varje grupp plus de 8 bästa treorna når trettiotvåondelsfinalen.",
	},
	"Which of these is NOT a 2026 host city in Mexico?": {
		Text:    "Vilken av dessa är INTE en värdstad 2026 i Mexiko?",
		Options: []string{"Guadalajara", "Monterrey", "Cancún", "Mexico City"},
		Fact:    "Mexikos värdstäder är Mexico City, Guadalajara och Monterrey.",
	},
	"The match-ball name \"Trionda\" (\"three waves\") celebrates what?": {
		Text:    "Matchbollens namn \"Trionda\" (\"tre vågor\") hyllar vad?",
		Options: []string{"De tre värdnationerna", "Tre poäng för en vinst", "Tre lejon", "Tre domare"},
		Fact:    "Den hyllar Kanada, Mexiko och USA — de första värdarna med tre nationer.",
	},
	"In the 2022 final, Mbappé scored a hat-trick but France lost on penalties. What was the score after extra time?": {
		Text:    "I finalen 2022 gjorde Mbappé en hat-trick men Frankrike förlorade på straffar. Vad var ställningen efter förlängning?",
		Options: []string{"2–2", "3–3", "4–4", "3–2"},
		Fact:    "Det slutade 3–3; Argentina vann 4–2 på straffar.",
	},
	"Which two Canadian cities are 2026 host cities?": {
		Text:    "Vilka två kanadensiska städer är värdstäder 2026?",
		Options: []string{"Toronto & Vancouver", "Montreal & Calgary", "Ottawa & Edmonton", "Toronto & Montreal"},
		Fact:    "Toronto och Vancouver är Kanadas två värdstäder.",
	},
	"Who is the World Cup's all-time leading goalscorer, with 16 goals?": {
		Text:    "Vem är VM:s meste målskytt genom tiderna, med 16 mål?",
		Options: []string{"Ronaldo (Brasilien)", "Miroslav Klose", "Gerd Müller", "Just Fontaine"},
		Fact:    "Tysklands Klose (16) gick om Brasiliens Ronaldo (15) 2014.",
	},
	"On which date is the 2026 World Cup final played?": {
		Text:    "Vilket datum spelas VM-finalen 2026?",
		Options: []string{"4 juli", "12 juli", "19 juli", "26 juli"},
		Fact:    "19 juli 2026 på MetLife Stadium.",
	},
	"Who scored the only hat-trick in a World Cup final, for England in 1966?": {
		Text:    "Vem gjorde den enda hat-tricken i en VM-final, för England 1966?",
		Options: []string{"Bobby Charlton", "Geoff Hurst", "Martin Peters", "Gary Lineker"},
		Fact:    "Geoff Hurst, i Englands 4–2-seger över Västtyskland.",
	},
	"Who scored Germany's extra-time winner in the 2014 World Cup final?": {
		Text:    "Vem gjorde Tysklands avgörande mål i förlängningen i VM-finalen 2014?",
		Options: []string{"Thomas Müller", "Mario Götze", "André Schürrle", "Bastian Schweinsteiger"},
		Fact:    "Mario Götzes lysande volley slog Argentina 1–0.",
	},
	"Which player was sent off for a headbutt in the 2006 World Cup final?": {
		Text:    "Vilken spelare blev utvisad för en skalle i VM-finalen 2006?",
		Options: []string{"Marco Materazzi", "Zinedine Zidane", "Patrick Vieira", "Thierry Henry"},
		Fact:    "Zidane skallade Materazzi; Italien vann straffläggningen.",
	},
	"Who scored the infamous \"Hand of God\" goal at the 1986 World Cup?": {
		Text:    "Vem gjorde det ökända \"Guds hand\"-målet vid VM 1986?",
		Options: []string{"Diego Maradona", "Jorge Valdano", "Gary Lineker", "Michel Platini"},
		Fact:    "Maradona, mot England i kvartsfinalen.",
	},
	"Who holds the record for most goals in a single World Cup — 13 in 1958?": {
		Text:    "Vem har rekordet för flest mål i ett enda VM — 13 mål 1958?",
		Options: []string{"Just Fontaine", "Pelé", "Sándor Kocsis", "Gerd Müller"},
		Fact:    "Frankrikes Just Fontaine — ett rekord som fortfarande står sig.",
	},
	"Spain beat which country 1–0 in the 2010 World Cup final?": {
		Text:    "Spanien slog vilket land 1–0 i VM-finalen 2010?",
		Options: []string{"Tyskland", "Nederländerna", "Italien", "Portugal"},
		Fact:    "Spanien 1–0 Nederländerna efter förlängning.",
	},
	"Who scored Spain's extra-time winner in the 2010 final?": {
		Text:    "Vem gjorde Spaniens avgörande mål i förlängningen i finalen 2010?",
		Options: []string{"Xavi", "David Villa", "Andrés Iniesta", "Fernando Torres"},
		Fact:    "Andrés Iniestas sena fullträff gav Spanien deras första VM-titel.",
	},
	"How many qualifying places did UEFA (Europe) receive for the 2026 World Cup?": {
		Text:    "Hur många kvalplatser fick UEFA (Europa) till VM 2026?",
		Options: []string{"13", "15", "16", "18"},
		Fact:    "16 platser — upp från 13 vid 2022 års upplaga.",
	},
	"Which is the only nation to have played at every World Cup since 1930?": {
		Text:    "Vilken är den enda nationen som har spelat i varje VM sedan 1930?",
		Options: []string{"Tyskland", "Italien", "Brasilien", "Argentina"},
		Fact:    "Brasilien — närvarande i alla 23 turneringar.",
	},
	"Italy won the 2006 World Cup final on penalties against which country?": {
		Text:    "Italien vann VM-finalen 2006 på straffar mot vilket land?",
		Options: []string{"Tyskland", "Frankrike", "Portugal", "Brasilien"},
		Fact:    "Italien slog Frankrike 5–3 på straffar efter 1–1.",
	},
	"Who won the Golden Boot (top scorer) at the 2022 World Cup, with 8 goals?": {
		Text:    "Vem vann Golden Boot (skytteligan) vid VM 2022, med 8 mål?",
		Options: []string{"Lionel Messi", "Kylian Mbappé", "Julián Álvarez", "Olivier Giroud"},
		Fact:    "Mbappés 8 mål knep den före Messis 7.",
	},
	"Who was named best player (Golden Ball) at the 2018 World Cup?": {
		Text:    "Vem utsågs till bäste spelare (Golden Ball) vid VM 2018?",
		Options: []string{"Kylian Mbappé", "Luka Modrić", "Eden Hazard", "Antoine Griezmann"},
		Fact:    "Kroatiens Luka Modrić, trots att han förlorade finalen.",
	},
	"Who is the only player to have won three World Cups?": {
		Text:    "Vem är den enda spelaren som har vunnit tre VM?",
		Options: []string{"Diego Maradona", "Pelé", "Franz Beckenbauer", "Cafu"},
		Fact:    "Pelé vann 1958, 1962 och 1970 med Brasilien.",
	},
	"Germany's stunning 2014 semi-final win over hosts Brazil finished what score?": {
		Text:    "Tysklands häpnadsväckande semifinalseger 2014 över värdnationen Brasilien slutade med vilket resultat?",
		Options: []string{"5–0", "7–1", "6–2", "4–0"},
		Fact:    "Ett knappt trovärdigt 7–1 i Belo Horizonte.",
	},
	"The 2026 opening match (Mexico 2–0 South Africa) set an unwanted record for what?": {
		Text:    "Öppningsmatchen 2026 (Mexiko 2–0 Sydafrika) satte ett oönskat rekord i vad?",
		Options: []string{"Flest röda kort i en VM-match (3)", "Lägst publiksiffra", "Längsta matchen", "Flest självmål"},
		Fact:    "Tre röda kort — det flesta som någonsin visats i en enda VM-match.",
	},
	"Mexico City's stadium became the first ground to host how many World Cup opening matches?": {
		Text:    "Mexico Citys stadion blev den första arenan att vara värd för hur många VM-öppningsmatcher?",
		Options: []string{"Två", "Tre", "Fyra", "En"},
		Fact:    "Tre — den öppnade även turneringarna 1970 och 1986.",
	},
	"Which nation became the first African team to reach a World Cup semi-final, in 2022?": {
		Text:    "Vilken nation blev det första afrikanska laget att nå en VM-semifinal, 2022?",
		Options: []string{"Senegal", "Marocko", "Kamerun", "Ghana"},
		Fact:    "Marocko slog Spanien och Portugal på vägen till semifinalerna.",
	},
	"Who is Brazil's all-time leading World Cup goalscorer, with 15 goals?": {
		Text:    "Vem är Brasiliens meste VM-målskytt genom tiderna, med 15 mål?",
		Options: []string{"Pelé", "Ronaldo (Nazário)", "Romário", "Neymar"},
		Fact:    "Ronaldo \"O Fenômeno\" gjorde 15 VM-mål för Brasilien.",
	},
	"Which goalkeeper captained Spain to the 2010 World Cup title?": {
		Text:    "Vilken målvakt var kapten när Spanien tog VM-titeln 2010?",
		Options: []string{"Iker Casillas", "Víctor Valdés", "Pepe Reina", "David de Gea"},
		Fact:    "Iker Casillas lyfte pokalen som både kapten och målvakt.",
	},
}
