## Json struct generator

Raw json instead of go structs
Store structs without any marshalling and allocations as a []byte

```go

//fields name
var (
  fid = "id" //field name
)


type User []byte

func (a *User) ID() int64{
  jsons.Int64((*a), fid)
}
```

## Generate (example)
1. Create user.go with some user struct
```go
type User struct{
  ID int
  Name string
  Age int
  Tags []string
  Channels []int
  //etc
}
```

2. Run: jsons user.go

3. You will get a userjs.go with []byte user struct and all get/put methods
