[![MIT License][license-shield]][license-url]

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

# Mongo Golang Search (mgs)


### Content index

* [What is this?](#what-is-this)
* [Getting Started](#getting-started)
    *  [Installation](#installation)
* [Usage](#usage)
* [Supported features](#supported-features)
    * [Filtering](#filtering)
    * [Pagination](#pagination)
    * [Sorting](#sorting)
    * [Projection](#projection)
* [Available options](#available-options)
    * [Customize limit value](#customize-limit-value)
    * [Specify casting per param keys](#specify-casting-per-param-keys)
* [Contributing](#contributing)
* [License](#license)

## What is this?

Mongo Golang Search provides a simple query language to perform advanced searches for your collections in **MongoDB**.

You could also use **Mongo Golang Search** to searching, sorting, pagination and combining logical operators.

## Getting Started

### Installation


## Usage


## Supported features

### Filtering

| Operator      	| URI               	| Example                        	|
| -----------------	| ---------------------	| ---------------------------------	|
| `$eq`          	| `key=val`				| `type=public`        				|
| `$ne`          	| `key!=val`        	| `status!=SENT`                    |
| `$gt`          	| `key>val`             | `price>5`                        |
| `$gte`         	| `key>=val`            | `price>=9`                       |
| `$lt`          	| `key<val`             | `date<2020-01-01T14:00:00.000Z`   |
| `$lte`         	| `key<=val`            | `priority<=-5`                    |
| `$in`          	| `key=val1,val2`       | `status=QUEUED,DEQUEUED`          |
| `$nin`          	| `key!=val1,val2`      | `status!=QUEUED,DEQUEUED`         |
| `$exists`         | `key`          		| `email`              				|
| `$exists`         | `!key`         		| `!email`                    		|
| `$regex`      	| `key=/value/<opts>`	| `email=/@gmail\.com$/`			|
| `$regex`        	| `key!=/value/<opts>`  | `phone!=/^58/`                    |


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
* The `_id` field (returned by default).

## Available options

You can use advanced options:

### Customize limit value


### Specify casting per param keys


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
