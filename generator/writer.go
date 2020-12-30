package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Writer struct {
	writer       *os.File
	format       string
	baseFileName string
}

func NewWriter(writerType, format string) (*Writer, error) {
	writer := os.Stdout
	fileName := ""
	baseFileName := ""
	var err error

	if writerType == "file" {
		fileName = format
		format = filepath.Ext(fileName)
		baseFileName = fileName[:len(fileName)-len(format)]
		writer, err = os.Create(fileName)
		if err != nil {
			return nil, err
		}
	} else {
		format = fmt.Sprintf(".%s", format)
	}
	return &Writer{
		writer:       writer,
		format:       format,
		baseFileName: baseFileName,
	}, nil
}

func (w *Writer) Write(header []string, data []map[string]interface{}) error {
	if w.format == ".json" {
		return w.writeJson(header, data)
	} else {
		if w.format == ".csv" {
			return w.writeCSV(header, data)
		} else {
			if w.format == ".tsv" {
				return w.writeTSV(header, data)
			} else {
				if w.format == ".sql" {
					return w.writeSQL(header, data)
				}
			}
		}
	}
	return nil
}

func (w *Writer) Close() error {
	return w.writer.Close()
}

func (w *Writer) writeJson(header []string, data []map[string]interface{}) error {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	_, err = w.writer.WriteString(string(b))
	return err
}

func (w *Writer) writeCSV(header []string, data []map[string]interface{}) error {
	return w.writeSomethingSV(header, data, ",")
}

func (w *Writer) writeTSV(header []string, data []map[string]interface{}) error {
	return w.writeSomethingSV(header, data, "\t")
}

func (w *Writer) writeSomethingSV(header []string, data []map[string]interface{}, sep string) error {
	_, err := w.writer.WriteString(fmt.Sprintf("%s\n", strings.Join(header, sep)))
	if err != nil {
		return err
	}
	for _, row := range data {
		buffer := []string{}
		for _, col := range header {
			intVal, intOk := row[col].(int64)
			if intOk {
				buffer = append(buffer, fmt.Sprintf("%d", intVal))
			} else {
				floatVal, floatOk := row[col].(float64)
				if floatOk {
					buffer = append(buffer, fmt.Sprintf("%f", floatVal))
				} else {
					buffer = append(buffer, fmt.Sprintf("\"%s\"", row[col].(string)))
				}
			}
		}
		_, err = w.writer.WriteString(fmt.Sprintf("%s\n", strings.Join(buffer, sep)))
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *Writer) writeSQL(header []string, data []map[string]interface{}) error {
	initialSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES", w.baseFileName, strings.Join(header, ","))

	inserts := []string{}

	for _, row := range data {
		buffer := []string{}
		for _, col := range header {
			intVal, intOk := row[col].(int64)
			if intOk {
				buffer = append(buffer, fmt.Sprintf("%d", intVal))
			} else {
				floatVal, floatOk := row[col].(float64)
				if floatOk {
					buffer = append(buffer, fmt.Sprintf("%f", floatVal))
				} else {
					buffer = append(buffer, fmt.Sprintf("'%s'", w.escapeQuote(row[col].(string))))
				}
			}
		}
		inserts = append(inserts, fmt.Sprintf("(%s)", strings.Join(buffer, ",")))
	}

	sql := fmt.Sprintf("%s %s;", initialSQL, strings.Join(inserts, ","))
	_, err := w.writer.WriteString(sql)
	if err != nil {
		return err
	}

	return nil
}

func (w *Writer) escapeQuote(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}
