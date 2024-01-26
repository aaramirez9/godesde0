package users

import (
	"fmt"
	"time"

	"github.com/aaramirez9/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(65, "Alex", time.Now(), true)
	fmt.Println(u)
}
