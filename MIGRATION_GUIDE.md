## 🧭 Migration Guide (v1 → v2)

Starting from **version 2**, the library replaces the legacy `Primitives` type with a more flexible and extensible `TypeConverter`. This change simplifies how custom types are handled and improves maintainability.

---

### Replace `Primitives` with `TypeConverter`

**Now use:**
```go
type TypeConverter struct{}
```

**Instead of**:
```go
type Primitives struct{}
```

### Replace `ObjectID` method with `Convert`

**Now use:**
```go
func (c *TypeConverter) Convert(value string) (interface{}, error) {
	if id, err := primitive.ObjectIDFromHex(value); err == nil {
		return id, err
	}
	return nil, fmt.Errorf("no conversion matched")
}
```

**Instead of**:
```go
func (primitives *Primitives) ObjectID(oidStr string) (interface{}, error) {
	return ObjectID() // invoke ObjectID from MongoDB Driver
}
```

### Inject `TypeConverter` into `NewQueryHandler`

**Now use:**
```go
queryHandler := mgs.NewQueryHandler(&TypeConverter{})
```

**Instead of**:
```go
queryHandler := mgs.NewQueryHandler(&Primitives{})
```