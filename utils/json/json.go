package json

import (
	"encoding/json"
	"log"

	"github.com/wildanie12/go-hex-arch-sample/utils/color"
)

// Encode your object, useful for debugging.
func Encode(obj interface{}) string {
	json, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Println(color.ThisF("yellow", "[jsonUtil] fail to encode json: %v", err))
	}
	return string(json)
}