package util

import (
	"strconv"
	"strings"
)

func GetSequence(lastID string, strRemove string) string {
	strSequence := strings.Replace(lastID, strRemove, "", 1)
	intSequence, _ := strconv.Atoi(strSequence)
	strSequence = "000000" + strconv.Itoa(intSequence+1)

	return strSequence[len(strSequence)-6:]
}

func GenerateID(lastID string, prefix string) string {
	sequence := "000001"
	if lastID != "" {
		sequence = GetSequence(lastID, prefix)
	}
	return prefix + sequence
}
