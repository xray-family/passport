## Passport

[![Go Test](https://github.com/lxzan/passport/actions/workflows/go.yml/badge.svg)](https://github.com/lxzan/passport/actions/workflows/go.yml) [![Coverage Statusd][1]][2]

[1]: https://codecov.io/gh/lxzan/passport/branch/main/graph/badge.svg
[2]: https://codecov.io/gh/lxzan/passport

### Install

```bash
go get -v github.com/lxzan/passport@latest
```

#### Usage

```go
type Req struct {
	Name  string `json:"name"`
	Typ   int    `json:"typ"`
	Age   int    `json:"age"`
	Roles []int  `json:"roles"`
}

func (c *Req) Validate() error {
	return passport.Validate(
		passport.NewString("Name", c.Name).Required().Alphabet(),
		passport.NewOrdered("Typ", c.Typ).IncludeBy(1, 3, 5),
		passport.NewOrdered("Age", c.Age).Gte(18),
		passport.NewSlice("Roles", c.Roles).Required(),
	)
}
```
