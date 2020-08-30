# Jsons
Based on cjson &amp; gjson. Fast works with json withou an allocations.


## Get a values

```go
b := []byte(`{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}`) 

jsons.String(b, "name.first") //Tom
jsons.Int(b, "age") //37
jsons.String(b, "friends[0].first") //Dale
jsons.StringArray(b, "friends[0].nets") //["ig", "fb", "tw"]
```

## Create a json
```go
b := jsons.Create().
  Add("name","Tom").
  Add("Age", 37).
  Add("friends[0].nets", ["ig", "fb", "tw"]).
  Bytes()
```
