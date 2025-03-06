## validator

[![Go Test](https://github.com/xray-family/validator/actions/workflows/go.yml/badge.svg)](https://github.com/xray-family/validator/actions/workflows/go.yml) [![Coverage Statusd][1]][2]

[1]: https://codecov.io/gh/lxzan/validator/branch/main/graph/badge.svg

[2]: https://codecov.io/gh/lxzan/validator

### Install

```bash
go get -v github.com/xray-family/validator@latest
```

### Quick Start

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/xray-family/validator"
    "net/http"
)

type Req struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Roles []int  `json:"roles"`
}

func (c *Req) Validate(r *http.Request) error {
    return validator.NewValidator(r).Validate(
        validator.String("Name", c.Name).Required().Alphabet(),
        validator.Ordered("Age", c.Age).Gte(18),
        validator.Slice("Roles", c.Roles).Required(),
    )
}

func Handle(writer http.ResponseWriter, request *http.Request) {
    var r = &Req{}
    _ = json.NewDecoder(request.Body).Decode(r)
    _ = request.Body.Close()
    var err = r.Validate(request)
    fmt.Printf("%v\n", err)
}

func main() {
    http.HandleFunc("/", Handle)
    http.ListenAndServe(":8080", nil)
}

```

### Advanced

#### Customized Check Functions

```go
package main

import (
    "fmt"
    "github.com/nicksnyder/go-i18n/v2/i18n"
    "github.com/xray-family/validator"
)

func isPhone(s string) bool {
    return len(s) == 11 && s[0] == '1'
}

func main() {
    _ = validator.GetBundle().AddMessages(validator.English, &i18n.Message{
        ID:    "Phone",
        Other: "Failed to verify cell phone number",
    })
    var validator = validator.NewValidator(nil)
    var err = validator.Validate(
        validator.String("phone number", "xyz").Customize("Phone", isPhone),
    )
    fmt.Printf("%v\n", err)
}
```

