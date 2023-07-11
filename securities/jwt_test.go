package securities

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	t.Setenv("JWT_KEY", "someKey")
	id, err := uuid.NewRandom()
	if err != nil {
		t.Fatalf("error in generate uuid: %s", err)
	}

	key, _ := GenerateJWT(id, "someUsername")
	assert.True(t, len(strings.Split(key, ".")) == 3)
}
