package main

import (
	"fmt"
	"tf-idf/lib/tokenizer"
)

func main() {
	var documentA = tokenizer.NewDocument("Lorem ipsum dolor amet street art tilde pickled umami echo park. Vape craft beer ramps raw denim enamel pin brooklyn squid keytar kitsch blue bottle austin sartorial mumblecore tacos artisan. Succulents meh activated charcoal venmo quinoa. Jean shorts 3 wolf moon man braid master cleanse shaman vegan")
	var documentB = tokenizer.NewDocument("Stumptown crucifix yuccie jianbing ethical. Glossier vaporware lo-fi cornhole brunch retro typewriter. Ramps letterpress succulents synth. Tattooed taiyaki man braid, tbh art party unicorn woke skateboard.")
	var documentC = tokenizer.NewDocument("viral roof party cornhole. Cornhole art party pug selvage, hashtag cold-pressed pok pok plaid chambray fam narwhal letterpress venmo. Seitan stumptown forage live-edge intelligentsia try-hard letterpress cloud bread vexillologist fixie normcore godard mustache")
	var tfidf = tokenizer.NewTFIDFTokenizer()

	tokenizer.AddDocumentToTokenizer(documentA, tfidf)
	tokenizer.AddDocumentToTokenizer(documentB, tfidf)
	tokenizer.AddDocumentToTokenizer(documentC, tfidf)

	tokenizer.Compute(tfidf)

	fmt.Println(tfidf.Documents[0].TFIDFValues)
	fmt.Println(tokenizer.ComputeSimiliarityBetween(tfidf.TFIDFVector[0], tfidf.TFIDFVector[2]))
}
