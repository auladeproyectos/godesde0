package users

import (
	"fmt"
	"time"

	"github.com/Auladeproyectos/godesde0/modelos"
)

func Altausuario() {
	u := new(modelos.User)
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u)
}
