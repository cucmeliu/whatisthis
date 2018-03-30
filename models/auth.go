// auth.go
package models

import (
	//	"bytes"
	//	"encoding/json"
	"fmt"
	//	"io/ioutil"
	//	"net/http"
	//	"net/url"
	//	"strings"
)

func GetToken() (token string) {
	return "24.b9078a2dc2ab60bc670d1babcf965b90.2592000.1519294427.282335-10736267"
}

func main() {
	fmt.Println(GetToken())
	//	httpGet()
}
