package quotes

import (
	"encoding/xml"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

type AnnotationSet struct {
	Publication Publication  `xml:"publication"`
	Annotations []Annotation `xml:"annotation"`
}

type Publication struct {
	Identifier string `xml:"identifier"`
	Title      string `xml:"title"`
	Creator    string `xml:"creator"`
}

type Annotation struct {
	Identifier string    `xml:"identifier"`
	Date       time.Time `xml:"date"`
	Text       string    `xml:"target>fragment>text"`
}

type Quote struct {
	Text   string    `json:"text"`
	Source string    `json:"source"`
	Author string    `json:"author"`
	Date   time.Time `json:"date"`
}

func (q Quote) String() string {
	return fmt.Sprintf("%s\n%s\n%s\n---\n%s",
		wrapText(q.Source, 80),
		q.Author,
		q.Date.Local().Format(time.RFC822),
		wrapText(q.Text, 80),
	)
}

func (a Annotation) IsBookmark() bool {
	return a.Text == ""
}

func LoadQuotes(annotationsDir string) ([]Quote, error) {
	dirEntries, err := os.ReadDir(annotationsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %s", err)
	}

	var quotes []Quote

	for _, de := range dirEntries {
		if de.IsDir() {
			continue
		}

		b, err := os.ReadFile(path.Join(annotationsDir, de.Name()))
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %s", err)
		}

		var annotationSet AnnotationSet
		err = xml.Unmarshal(b, &annotationSet)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal: %s", err)
		}

		for _, annot := range annotationSet.Annotations {
			if annot.IsBookmark() {
				continue
			}

			// Trim space
			cleanedText := strings.TrimSpace(annot.Text)

			// Replace curly quotes and other variations with standard quotes
			cleanedText = strings.ReplaceAll(cleanedText, "“", `"`)
			cleanedText = strings.ReplaceAll(cleanedText, "”", `"`)
			cleanedText = strings.ReplaceAll(cleanedText, "‘", `'`)
			cleanedText = strings.ReplaceAll(cleanedText, "’", `'`)
			cleanedText = strings.ReplaceAll(cleanedText, "`", "'") // Convert backticks to single quotes

			// Remove special characters
			re := regexp.MustCompile(`[^a-zA-Z0-9\s\-;,._<>()]'"`)
			cleanedText = re.ReplaceAllString(cleanedText, "")

			// Capitalize first letter
			firstChar := string(cleanedText[0])
			if firstChar != strings.ToUpper(firstChar) {
				cleanedText = strings.ToUpper(firstChar) + cleanedText[1:]
			}

			quotes = append(quotes, Quote{
				Text:   cleanedText,
				Source: annotationSet.Publication.Title,
				Author: annotationSet.Publication.Creator,
				Date:   annot.Date,
			})
		}
	}

	return quotes, nil
}

func wrapText(text string, limit int) string {
	words := strings.Fields(text) // Split text into words
	var wrapped strings.Builder
	lineLen := 0

	for _, word := range words {
		// If adding the word exceeds the limit, start a new line
		if lineLen+len(word) > limit {
			wrapped.WriteString("\n")
			lineLen = 0
		}
		// Add space before word unless it's the first word in a line
		if lineLen > 0 {
			wrapped.WriteString(" ")
			lineLen++
		}
		wrapped.WriteString(word)
		lineLen += len(word)
	}

	return wrapped.String()
}
