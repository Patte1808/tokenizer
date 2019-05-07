package tokenizer

import (
	"math"
	"tf-idf/lib/utils"
)

type tfidfTokenizer struct {
	Documents [] *document
	InverseDocumentFrequency map[string] float64
}

type document struct {
	WordCount map[string] int
	TermFrequency map[string] float64
	Words [] string
	TFIDFValues map[string] float64
}

// NewDocument creates a document
// It requires a string which describes the document content (e.g. a sentence)
// This function takes care of cleaning the document content from unnecessary characters (e.g. punctuation),
// which might otherwise create some noise in the end-result
func NewDocument(documentContent string) *document {
	document := new(document)
	documentContent = utils.CleanDocumentContent(documentContent)
	document.Words = utils.CreateWordsFromString(documentContent)
	document.WordCount = make(map[string] int)
	document.TermFrequency = make(map[string] float64)
	document.TFIDFValues = make(map[string] float64)

	return document
}

// NewTFIDFTokenizer creates a TFIDFTokenizer
func NewTFIDFTokenizer() *tfidfTokenizer {
	tokenizer := new(tfidfTokenizer)
	tokenizer.InverseDocumentFrequency = make(map[string] float64)
	tokenizer.Documents = make([] *document, 0)

	return tokenizer
}

// AddDocumentToTokenizer takes a document pointer and a tokenizer pointer
// and adds the document to the tokenizers document collection
func AddDocumentToTokenizer(doc *document, tokenizer *tfidfTokenizer) {
	tokenizer.Documents = append(tokenizer.Documents, doc)
}

// Compute takes a pointer to a tokenizer
// It calls all necessary functions to retrieve a TFIDF score for all documents in the tokenizers collection
func Compute(tokenizer *tfidfTokenizer) {
	for _, document := range tokenizer.Documents {
		computeTF(document)
	}

	computeIDF(tokenizer)
	computeTFIDF(tokenizer)
}

func computeTF(document *document) {
	for _, word := range document.Words {
		document.WordCount[word] += 1
	}

	for word, _ := range document.WordCount {
		document.TermFrequency[word] = float64(document.WordCount[word]) / float64(len(document.Words))
	}
}

func computeIDF(tfidf *tfidfTokenizer) {
	allDocumentsWordCount := make(map[string] int)
	inverseDocumentFrequency := make(map[string] float64)

	for _, document := range tfidf.Documents {
		for _, word := range document.Words {
			allDocumentsWordCount[word] += 1
		}
	}

	for word := range allDocumentsWordCount {
		inverseDocumentFrequency[word] = 1 + math.Log(float64(len(tfidf.Documents)) / float64(allDocumentsWordCount[word]))
	}

	tfidf.InverseDocumentFrequency = inverseDocumentFrequency
}

func computeTFIDF(tokenizer *tfidfTokenizer) {
	for _, document := range tokenizer.Documents {
		tfidfValues := make(map[string] float64)

		for word := range document.TermFrequency {
			tfidfValues[word] = document.TermFrequency[word] * tokenizer.InverseDocumentFrequency[word]
		}

		document.TFIDFValues = tfidfValues
	}
}