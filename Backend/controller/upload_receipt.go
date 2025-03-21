package controller

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"my-backend/graph/model"
)

func (d *Service) UploadReceipt(file graphql.Upload) (*model.ReceiptResult, error) {

	tempFile, err := os.CreateTemp("", "upload-*.png")
	if err != nil {
		return nil, fmt.Errorf("Fehler beim Erstellen der Datei: %w", err)
	}
	defer os.Remove(tempFile.Name()) // Temp-Datei wieder löschen

	_, err = io.Copy(tempFile, file.File)
	if err != nil {
		return nil, fmt.Errorf("Fehler beim Kopieren der Datei: %w", err)
	}
	tempFile.Close()

	// 2. OCR (Tesseract) ausführen
	text, err := runTesseract(tempFile.Name())
	if err != nil {
		return nil, fmt.Errorf("Fehler beim OCR: %w", err)
	}

	// 3. Text parsen
	items, total := parseReceiptText(text)

	// 4. Rückgabe formatieren
	var gqlItems []*model.ReceiptItem
	for _, item := range items {
		gqlItems = append(gqlItems, &model.ReceiptItem{
			Name:  item.Name,
			Price: item.Price,
		})
	}

	return &model.ReceiptResult{
		Items: gqlItems,
		Total: total,
	}, nil
}

// Struktur für intern geparste Items
type item struct {
	Name  string
	Price float64
}

// Führt tesseract aus und gibt Text zurück
func runTesseract(filePath string) (string, error) {
	output := filePath + "-out"
	cmd := exec.Command("tesseract", filePath, output, "-l", "deu") // deutsche Sprache
	if err := cmd.Run(); err != nil {
		return "", err
	}

	textBytes, err := os.ReadFile(output + ".txt")
	if err != nil {
		return "", err
	}
	defer os.Remove(output + ".txt")

	return string(textBytes), nil
}

// Parst den OCR-Text in Produktliste und Gesamtpreis
func parseReceiptText(text string) ([]item, float64) {
	lines := strings.Split(text, "\n")
	var items []item
	var total float64

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Gesamtbetrag erkennen
		if strings.Contains(strings.ToLower(line), "gesamt") || strings.Contains(strings.ToLower(line), "summe") {
			for _, word := range strings.Fields(line) {
				if val, err := parsePrice(word); err == nil {
					total = val
					break
				}
			}
			continue
		}

		// Einzelprodukt parsen: letzte Zahl als Preis
		words := strings.Fields(line)
		if len(words) >= 2 {
			last := words[len(words)-1]
			if val, err := parsePrice(last); err == nil {
				name := strings.Join(words[:len(words)-1], " ")
				items = append(items, item{Name: name, Price: val})
			}
		}
	}

	return items, total
}

// Wandelt String in Float, ersetzt Komma durch Punkt
func parsePrice(s string) (float64, error) {
	s = strings.ReplaceAll(s, ",", ".")
	var f float64
	_, err := fmt.Sscanf(s, "%f", &f)
	return f, err
}
