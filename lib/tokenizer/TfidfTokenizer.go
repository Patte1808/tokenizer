package tokenizer

import (
	"math"
	"tokenizer/lib/utils"
)

type tfidfTokenizer struct {
	Documents                []*document
	AllDocumentsWordCount    map[string]int
	InverseDocumentFrequency map[string]float64
	TFIDFVector              []map[string]float64
}

type document struct {
	WordCount     map[string]int
	TermFrequency map[string]float64
	Words         []string
	TFIDFValues   map[string]float64
}

// NewDocument creates a document
// It requires a string which describes the document content (e.g. a sentence)
// This function takes care of cleaning the document content from unnecessary characters (e.g. punctuation),
// which might otherwise create some noise in the end-result
func NewDocument(documentContent string) *document {
	document := new(document)
	documentContent = utils.CleanDocumentContent(documentContent)
	document.Words = utils.CreateWordsFromString(documentContent)
	document.WordCount = make(map[string]int)
	document.TermFrequency = make(map[string]float64)
	document.TFIDFValues = make(map[string]float64)

	return document
}

// NewTFIDFTokenizer creates a TFIDFTokenizer
func NewTFIDFTokenizer() *tfidfTokenizer {
	tokenizer := new(tfidfTokenizer)
	tokenizer.InverseDocumentFrequency = make(map[string]float64)
	tokenizer.Documents = make([]*document, 0)

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
	computeTFIDFVector(tokenizer)
}

// ComputeSimiliarityBetween takes 2 vectors and computes their similiarity
// In order to do so, we compute the dotproduct of the given vectors and divide it by the product of their magnitude
func ComputeSimiliarityBetween(vectorX map[string]float64, vectorY map[string]float64) float64 {
	return computeDotProduct(vectorX, vectorY) / (computeMagnitude(vectorX) * computeMagnitude(vectorY))
}

// GetFeaturesNames gets the features from all attached documents of the TFIDFTokenizer
func GetFeatureNames(tokenizer *tfidfTokenizer) [] string {
	var features = make([] string, len(tokenizer.AllDocumentsWordCount))

	for featureName := range tokenizer.AllDocumentsWordCount {
		features = append(features, featureName)
	}

	return features
}

// computeTF computes the term frequency for a given document
// It first calculates the overall occurence of a word in a given document
// and then calculates the term frequency
func computeTF(document *document) {
	for _, word := range document.Words {
		document.WordCount[word] += 1
	}

	for word, _ := range document.WordCount {
		document.TermFrequency[word] = float64(document.WordCount[word]) / float64(len(document.Words))
	}
}

// computeIDF returns the inverse document frequency for a given tokenizer.
func computeIDF(tfidf *tfidfTokenizer) {
	allDocumentsWordCount := make(map[string]int)
	inverseDocumentFrequency := make(map[string]float64)

	for _, document := range tfidf.Documents {
		for _, word := range document.Words {
			allDocumentsWordCount[word] += 1
		}
	}

	for word := range allDocumentsWordCount {
		inverseDocumentFrequency[word] = math.Log((1 + float64(len(tfidf.Documents))) / (1 + float64(allDocumentsWordCount[word])))
	}

	tfidf.InverseDocumentFrequency = inverseDocumentFrequency
	tfidf.AllDocumentsWordCount = allDocumentsWordCount
}

func computeTFIDF(tokenizer *tfidfTokenizer) {
	for _, document := range tokenizer.Documents {
		tfidfValues := make(map[string]float64)

		for word := range document.TermFrequency {
			tfidfValues[word] = document.TermFrequency[word] * tokenizer.InverseDocumentFrequency[word]
		}

		document.TFIDFValues = tfidfValues
	}
}

func computeTFIDFVector(tokenizer *tfidfTokenizer) {
	tfidfVector := make([](map[string]float64), len(tokenizer.Documents))

	for i, document := range tokenizer.Documents {
		tfidfVector[i] = make(map[string]float64)

		for word := range tokenizer.AllDocumentsWordCount {
			if value, ok := document.TFIDFValues[word]; ok {
				tfidfVector[i][word] = value
			}
		}
	}

	tokenizer.TFIDFVector = tfidfVector
}

func computeDotProduct(vectorX map[string]float64, vectorY map[string]float64) float64 {
	var dotProduct = 0.0

	for index, value := range vectorX {
		dotProduct += value * vectorY[index]
	}

	return dotProduct
}

func computeMagnitude(vector map[string]float64) float64 {
	var magnitude = 0.0

	for _, value := range vector {
		magnitude += math.Pow(value, 2)
	}

	return math.Sqrt(magnitude)
}
