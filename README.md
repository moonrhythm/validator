# validator

Validator is not validator, just an string error list helper :P

## Why

```go
f := flash()

if name == "" {
    f.AddError("name required")
}
if email == "" {
    f.AddError("email required")
}
if !govalidator.IsEmail(email) {
    f.AddError("invalid email")
}
if err := validatePassword(pass); err != nil {
    f.AddError("invalid password")
}
// ...
if f.HasError() {
    // ...
}

// no error
```

## Then

```go
v := validator.New() // this is not the validator :P

v.Must(name != "", "name required")
v.Must(email != "", "email required")
v.Must(govalidator.IsEmail(email), "invalid email")
v.Must(validatePassword(pass), "invalid password")
// ...
if !v.Valid() {
    // has error, ...
}

// no error
```
