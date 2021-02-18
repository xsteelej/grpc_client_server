package ports

import (
	"context"
	"io"
	"log"
	"strings"
)

func ReadJsonFile(ctx context.Context, r io.Reader, w io.Writer) error {
	objectLevel := 0

	buf := make([]byte, 1)
	currentObject := ""

	for !contextCancelled(ctx) {
		if _, err := r.Read(buf); err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		currentObject += strings.Trim(string(buf), "\n\t")
		objectLevel = calculateObjectLevel(buf, objectLevel, currentObject)
		if string(buf) != "}" {
			continue
		}
		complete := false
		complete, currentObject = checkCompleteJsonObject(objectLevel, currentObject)
		if !complete {
			continue
		}
		_, err := w.Write([]byte(currentObject))
		if err != nil {
			log.Println("Error writing: " + currentObject)
		}
		currentObject = ""
	}

	return nil
}

func calculateObjectLevel(buf []byte, objectLevel int, currentObject string) int {
	if string(buf) == "{" {
		objectLevel++
	}

	if string(buf) == "}" {
		objectLevel--
	}

	return objectLevel
}

func checkCompleteJsonObject(objectLevel int, currentObject string) (bool, string) {
	if objectLevel != 1 {
		return false, currentObject
	}
	currentObject += "}"
	if strings.HasPrefix(currentObject, ",") {
		currentObject = strings.Replace(currentObject, ",", "{", 1)
	}

	return true, currentObject
}

func contextCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
