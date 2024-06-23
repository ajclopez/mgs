[![Go Reference](https://pkg.go.dev/badge/github.com/ajclopez/mgs.svg)](https://pkg.go.dev/github.com/ajclopez/mgs) [![GitHub release (with filter)](https://img.shields.io/github/v/release/ajclopez/mgs)](https://github.com/ajclopez/mgs/releases) [![Go Report](https://goreportcard.com/badge/ajclopez/mgs)](https://goreportcard.com/report/ajclopez/mgs) [![codecov](https://codecov.io/gh/ajclopez/mgs/graph/badge.svg?token=PGKFDNP677)](https://codecov.io/gh/ajclopez/mgs) [![MIT License][license-shield]][license-url]

<p align="center">
  <h3 align="center">Mongo Golang Search</h3>
  <p align="center">
    Mongo Golang Search provides a query language to a MongoDB database.
    <br />
    <a href="https://github.com/ajclopez/mgs#usage"><strong>Explore the docs</strong></a>
    <br />
    <br />
    <a href="https://github.com/ajclopez/mgs/issues">Report Bug</a>
    ·
    <a href="https://github.com/ajclopez/mgs/issues">Request Feature</a>
  </p>
</p>

-------------------------

# Mongo Golang Search (mgs)

### Content index

- [What is this?](#what-is-this)
- [Getting Started](#getting-started)
  - [Installation](#installation)
- [Usage](#usage)
- [Supported features](#supported-features)
  - [Filtering](#filtering)
  - [Pagination](#pagination)
  - [Sorting](#sorting)
  - [Projection](#projection)
  - [Advanced queries](#advanced-queries)
- [Available options](#available-options)
  - [Customize limit value](#customize-limit-value)
  - [Specify casting per param keys](#specify-casting-per-param-keys)
- [Contributing](#contributing)
- [License](#license)

## What is this?

Mongo Golang Search provides a simple query language to perform advanced searches for your collections in **MongoDB**.

You could also use **Mongo Golang Search** to searching, sorting, pagination and combining logical operators.

## Getting Started

### Installation

The recommended way to get started using the Mongo Golang Search is by using Go modules to install the dependency in your project. This can be done either by importing packages from github.com/ajclopez/mgs and having the build step install the dependency or by explicitly running

```go
go get github.com/ajclopez/mgs
```

## Usage

To get started with mgs, import the mgs package and implement the Primitives ObjectID function which is used to convert strings to ObjectID.

```go
type Primitives struct{}

func (primitives *Primitives) ObjectID(oidStr string) (interface{}, error) {
	return ObjectID() // invoke ObjectID from MongoDB Driver
}
```

Then create an instance of `QueryHandler`:

```go
queryHandler := mgs.NewQueryHandler(&Primitives{})
```

Finally use a mgs.MongoGoSearch function:

```go
queryHandler.MongoGoSearch(query string, opts *FindOptions)
```

##### Arguments

`query`: query string part of the requested API URL.

`opts`: object for advanced configuration [See below](#available-options) [optional].

Using `mgs.MongoGoSearch` function with [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) library to filter, sort, limit and skip in MongoDB:

```go
import (
    "context"

    "github.com/ajclopez/mgs"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Primitives struct{}

func (primitives *Primitives) ObjectID(oidStr string) (interface{}, error) {
	return primitive.ObjectIDFromHex(oidStr)
}

queryHandler := mgs.NewQueryHandler(&Primitives{})

opts := mgs.FindOption()
result, err := queryHandler.MongoGoSearch(query, opts)

...

findOpts := options.Find()
findOpts.SetLimit(result.Limit)
findOpts.SetSkip(result.Skip)
findOpts.SetSort(result.Sort)
findOpts.SetProjection(result.Projection)

cur, err := collection.Find(context.TODO(), result.Filter, findOpts)

...
```

Using optional configurations:

```go
import (
    "context"

    "github.com/ajclopez/mgs"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Primitives struct{}

func (primitives *Primitives) ObjectID(oidStr string) (interface{}, error) {
	return primitive.ObjectIDFromHex(oidStr)
}

queryHandler := mgs.NewQueryHandler(&Primitives{})

opts := mgs.FindOption()
opts.SetCaster(map[string]mgs.CastType{
	"mobile": mgs.STRING,
})
opts.SetMaxLimit(1000)
opts.SetDefaultLimit(10)
result, err := queryHandler.MongoGoSearch(query, opts)

...

findOpts := options.Find()
findOpts.SetLimit(result.Limit)
findOpts.SetSkip(result.Skip)
findOpts.SetSort(result.Sort)
findOpts.SetProjection(result.Projection)

cur, err := collection.Find(context.TODO(), result.Filter, findOpts)

...
```

##### Example

A request of the form:

```js
'employees?status=sent&date>2020-01-06T14:00:00.000Z&author.firstname=Jhon&skip=50&limit=100&sort=-date&fields=id,date';
```

Is translated to:

```Go
Query{
    Filter: map[string]interface{}{
        "author.firstname":     "John",
        "date":                 map[string]interface{}{"$gt": "2020-01-06T14:00:00.000Z"},
        "status":               "SENT",
    },
    Sort:  map[string]int{
        "date": 1,
        "id": 1
    },
    Limit: 100,
    Skip:  50,
}
```

## Supported features

### Filtering

| Operator  | URI                  | Example                         |
| --------- | -------------------- | ------------------------------- |
| `$eq`     | `key=val`            | `type=public`                   |
| `$ne`     | `key!=val`           | `status!=SENT`                  |
| `$gt`     | `key>val`            | `price>5`                       |
| `$gte`    | `key>=val`           | `price>=9`                      |
| `$lt`     | `key<val`            | `date<2020-01-01T14:00:00.000Z` |
| `$lte`    | `key<=val`           | `priority<=-5`                  |
| `$in`     | `key=val1,val2`      | `status=QUEUED,DEQUEUED`        |
| `$nin`    | `key!=val1,val2`     | `status!=QUEUED,DEQUEUED`       |
| `$exists` | `key`                | `email`                         |
| `$exists` | `!key`               | `!email`                        |
| `$regex`  | `key=/value/<opts>`  | `email=/@gmail\.com$/`          |
| `$regex`  | `key!=/value/<opts>` | `phone!=/^58/`                  |

### Pagination

Useful to limit the number of records returned.

- Operator keys are `skip` and `limit`.
- Use `limit` operator to limit the number of records returned.
- Use `skip` operator to skip the specified number of records.

```json
skip=20&limit=10
```

### Sorting

Useful to sort returned records.

- Operator key is `sort`.
- It accepts a comma-separated list of fields.
- Use `-` prefixes to sort in descending order.
- Use `+` prefixes to sort in ascedending order.

```json
sort=id,-date
```

### Projection

Useful to limit fields to return in each records.

- Operator key is `fields`.
- It accepts a comma-separated list of fields.

```json
fields=firstname,lastname,phone,email
```

**Note:**

- The `_id` field (returned by default).

### Advanced queries

For more advanced usage (`and`, `or` logic operations), pass query `filter` as string with the logical operations, for example:

```json
filter=(country=Mexico OR country=Spain) and gender=female
```

##### What operations are possible?

- Filtering operations.
- The `AND/and` operator.
- The `OR/or` operator.
- Parenthesis can be used for grouping.

## Available options

You can use advanced options:

```go
opts := mgs.FindOption()
opts.SetCaster(map[string]mgs.CastType{
	"mobile": mgs.STRING,
})
opts.SetMaxLimit(100)
opts.SetDefaultLimit(10)
```

- `FindOption` creates a new FindOptions instance.
- `SetCaster` object to specify custom casters, key is the caster name, and value is a type (`BOOLEAN, NUMBER, PATTERN, DATE, STRING`).
- `SetDefaultLimit` which contains custom value to return records.
- `SetMaxLimit` which contains custom value to return a maximum of records.

### Customize limit value

You can specify your own maximum or default limit value.

- `defaultLimit`: custom value to return records.
- `maxLimit`: custom value to return a maximum of records.

```go
opts := mgs.FindOption()
opts.SetMaxLimit(1000)
opts.SetDefaultLimit(10)

result, err := mgs.MongoGoSearch("city=Madrid&skip=10&limit=1000", opts)
```

### Specify casting per param keys

You can specify how query parameter values are casted by passing an object.

- `casters`: object which map keys to casters.

```go
opts := mgs.FindOption()
opts.SetCaster(map[string]mgs.CastType{
	"key1": mgs.STRING,
    "key2": mgs.NUMBER,
    "key3": mgs.STRING,
    "key4": mgs.BOOLEAN
})

result, err := mgs.MongoGoSearch("key1=VALUE&key2=10&key3=20&key4=true", opts)
```

## Contributing

Should you like to provide any feedback, please open up an Issue, I appreciate feedback and comments. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing-feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This software is released under the MIT license. See `LICENSE` for more information.

[license-shield]: https://img.shields.io/badge/License-MIT-yellow.svg
[license-url]: https://github.com/ajclopez/mgs/blob/master/LICENSE
