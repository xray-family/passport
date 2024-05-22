## Passport

[![Go Test](https://github.com/xray-family/passport/actions/workflows/go.yml/badge.svg)](https://github.com/xray-family/passport/actions/workflows/go.yml) [![Coverage Statusd][1]][2]

[1]: https://codecov.io/gh/lxzan/passport/branch/main/graph/badge.svg

[2]: https://codecov.io/gh/lxzan/passport

### Install

```bash
go get -v github.com/xray-family/passport@latest
```

### Quick Start

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/xray-family/passport"
    "net/http"
)

type Req struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Roles []int  `json:"roles"`
}

func (c *Req) Validate() error {
    return passport.Validate(
        passport.String("Name", c.Name).Required().Alphabet(),
        passport.Ordered("Age", c.Age).Gte(18),
        passport.Slice("Roles", c.Roles).Required(),
    )
}

func Handle(writer http.ResponseWriter, request *http.Request) {
    var r = &Req{}
    _ = json.NewDecoder(request.Body).Decode(r)
    _ = request.Body.Close()
    var err = r.Validate(request.Header)
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
    "github.com/xray-family/passport"
    "github.com/nicksnyder/go-i18n/v2/i18n"
    "golang.org/x/text/language"
)

func isPhone(s string) bool {
    return len(s) == 11 && s[0] == '1'
}

func main() {
    _ = passport.GetBundle().AddMessages(language.Make("en-US"), &i18n.Message{
        ID:    "Phone",
        Other: "Failed to verify cell phone number",
    })
    var validator = passport.NewValidator()
    var err = validator.Validate(
        passport.String("phone number", "xyz").Customize("Phone", isPhone),
    )
    fmt.Printf("%v\n", err)
}
```

#### Field Translations

```go
package main

import (
    "fmt"
    "github.com/xray-family/passport"
    "github.com/nicksnyder/go-i18n/v2/i18n"
    "golang.org/x/text/language"
)

func main() {
    tag := language.Make("zh-CN")
    _ = passport.GetBundle().AddMessages(tag, &i18n.Message{
        ID:    "Req.Name",
        Other: "用户名",
    })
    var validator = passport.NewValidator("zh-CN")
    var err = validator.Validate(
        passport.
            String(validator.Localize("Req.Name"), "").
            Required(),
    )
    fmt.Printf("%v\n", err)
}
```
