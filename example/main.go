package main

import (
	"fmt"
	"tokenizer/lib/tokenizer"
)

func main() {
	var documentA = tokenizer.NewDocument("Captain America ist eine US-amerikanische Comicfigur, ein Superheld, der ein Kostüm in den Farben der Flagge der Vereinigten Staaten trägt. Geschaffen wurde sie von Jack Kirby und Joe Simon für den Verlag Timely Publications, einen Vorgänger des heutigen Verlags Marvel Comics. Captain America erschien zuerst 1941 in Heft 1 der Comicserie Captain America Comics. Die zu Kriegszeiten als Propaganda angelegten Comicgeschichten ließen Captain America häufig gegen Nazis, Saboteure und andere Versinnbildlichungen der damaligen Kriegsgegner antreten. In späteren Jahren wurden die Geschichten mit Captain America von vielen Autoren zur Sozialkritik eingesetzt, jedoch gestaltet sich die Rezeption insbesondere in Deutschland schwierig. Im Laufe der Jahre erschienen Figuren mit unterschiedlichen Namen im Kostüm des Captain America. Die ursprüngliche und bekannteste Figur trägt – als Alter Ego – den Namen Steve Rogers.")
	var documentB = tokenizer.NewDocument("Iron Man (deutsch: „Eisenmann“, in den ersten deutschen Veröffentlichungen „Der Eiserne“) ist eine Comicfigur der Marvel Comics. Erschaffen wurde sie von Stan Lee und Larry Lieber sowie den Zeichnern Don Heck und Jack Kirby. Ihr erster Auftritt war in dem Comic Tales of Suspense #39 im März 1963. 1968 erhielt sie eine eigene Reihe mit dem Titel Iron Man, die bis 1996 fortgesetzt wurde.")
	var documentC = tokenizer.NewDocument("Marvel’s The Avengers ist ein US-amerikanischer Action- und Science-Fiction-Spielfilm aus dem Jahr 2012, der als Comicverfilmung auf dem Superheldenteam The Avengers (in früheren deutschen Übersetzungen „Die Rächer“) des Verlages Marvel basiert. Regie führte Joss Whedon, die Hauptrollen spielen Chris Evans, Chris Hemsworth, Mark Ruffalo, Scarlett Johansson, Jeremy Renner und Robert Downey Jr. In den Vereinigten Staaten und in Deutschland ist Walt Disney Pictures für den Verleih zuständig, allerdings wird auf Filmplakaten, in den Trailern sowie im Vor- und Abspann des Films infolge einer Vereinbarung mit Paramount Pictures stattdessen das Logo von Paramount selbst verwendet.[3] Der offizielle Filmstart in den Vereinigten Staaten war am 4. Mai 2012, in Deutschland und Österreich bereits am 26. April.[4] Der Film wurde sowohl in 2D als auch in konvertiertem 3D veröffentlicht.[5]Marvel’s The Avengers ist der sechste Film, der direkt in Eigenfinanzierung von der Produktionsgesellschaft Marvel Studios produziert wurde. Er ist Teil des Franchise Marvel Cinematic Universe. Im April 2015 erschien die Fortsetzung, Avengers: Age of Ultron und im April 2018 der dritte Teil Avengers: Infinity War. Der vierte Teil Avengers: Endgame ist im April 2019 erschienen.[6]")
	var documentD = tokenizer.NewDocument("DC Comics ist neben Marvel Comics einer der größten US-amerikanischen Comicverlage. Besonders bekannt ist er aufgrund der Comicserien Batman und Superman. Weitere bekannte Figuren und Reihen sind u. a. Wonder Woman, The Flash, Aquaman und Green Lantern sowie die Justice League.Gegründet wurde der Verlag 1934 von Malcolm Wheeler-Nicholson als National Allied Publications. Der heutige Name „DC Comics“ leitet sich von einer der ersten, sehr erfolgreichen Comicserien des Verlages, Detective Comics, ab, die seit 1937 erscheint. Der Verlag ist seit 1969 eine Tochtergesellschaft von Warner Bros. Seinen Hauptsitz hat der Verlag in New York City. Seit 2009 ist Diane Nelson Präsidentin.")
	var tfidf = tokenizer.NewTFIDFTokenizer()

	tokenizer.AddDocumentToTokenizer(documentA, tfidf)
	tokenizer.AddDocumentToTokenizer(documentB, tfidf)
	tokenizer.AddDocumentToTokenizer(documentC, tfidf)
	tokenizer.AddDocumentToTokenizer(documentD, tfidf)

	tokenizer.Compute(tfidf)

	//fmt.Println(tfidf.Documents[0].TFIDFValues)
	fmt.Println(tokenizer.ComputeSimiliarityBetween(tfidf.TFIDFVector[0], tfidf.TFIDFVector[3]))
}
