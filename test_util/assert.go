package test_util

import "github.com/google/uuid"

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
