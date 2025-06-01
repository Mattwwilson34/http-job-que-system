package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// Generate a UUID compliant string
func GenerateUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)

	if err != nil {
		return "", err
	}

	// Add UUID Spec bits
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant bits

	// According to crypto/rand.Read docs, n == len(uuid) if err == nil
	// This check is defensive programming - should never trigger in practice
	if n < len(uuid) {
		return "", fmt.Errorf("crypto/rand.Read returned %d bytes, expected %d", n, len(uuid))
	}

	buf := make([]byte, 36)
	hexPos := 0
	uuidPos := 0
	uuidSegmentLengths := []int{4, 2, 2, 2, 6}

	for i, segmentLength := range uuidSegmentLengths {
		// Write hex chars to buffer
		hex.Encode(buf[hexPos:], uuid[uuidPos:uuidPos+segmentLength])

		// Update positions
		uuidPos += segmentLength
		hexPos += segmentLength * 2

		// Add dash if not last segment
		if i < len(uuidSegmentLengths)-1 {
			buf[hexPos] = '-'
			hexPos++
		}
	}

	return string(buf[:]), nil
}
