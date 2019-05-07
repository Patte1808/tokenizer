package utils

import "strings"

// CleanDocumentContent tries to reduce the noise of a string by removing unnecessary characters
// This function is not complete yet
func CleanDocumentContent(documentContent string) string {
	documentContent = strings.ReplaceAll(documentContent, ",", "")
	documentContent = strings.ReplaceAll(documentContent, ".", "")
	documentContent = strings.ReplaceAll(documentContent, ":", "")

	return documentContent
}

// CreateWordsFromString takes care of splitting a string into a collection of words and also makes sure that every
// word is lowercased
func CreateWordsFromString(documentContent string) []string {
	var seperatedDocumentContent = strings.Split(documentContent, " ")
	var words = make([]string, len(seperatedDocumentContent))

	for i := range seperatedDocumentContent {
		words[i] = strings.ToLower(seperatedDocumentContent[i])
	}

	return words
}
