package email

import (
	"fmt"
	"math/rand"
	"time"
)

func RandCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
