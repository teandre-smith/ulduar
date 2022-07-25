# ulduar

## What is 'ulduar'

This is a repository that attempts to make working with AWS DataAPI a bit easier. This package can be used as a basic ORM or as a very basic raw sql statement generator. This package is very much still a work in progress, but as time goes on, support for more complex queries will be implemented.

Please take note that majority of these functions that generate statements using structs/collections will require a struct tag named `datapi` to formulate the column names.

## Why Ulduar?

Ulduar? Well, as a former player of World of Warcraft, us participants use to enter Ulduar and face only the most trusted Watchers of the ancient world. But Ulduar also housed the most technological assets, and in addition to that, Thorim. Who can also be known as Thor. He is the God associated with thunder, storms and other energy related superstition.

And as a current employee at Tensor Energy, a company looking to build to make managing renewable assets easier, it only makes sense to name this basic package 'ulduar' since it is the home/manager of data that can be associated with energy.

Not literally, this is only a general package for AWS DataAPI. But, this package is the result of working with AWS DataAPI often in Golang.

## Examples

### Set Up

```go
    api := ulduar.DataApi{
        RDSClient: {rdsClient},
        ResourceArn: {YourArn},
        SecretArn: {YourArn},
        Dbname: {YourDB},
    }
```

### Struct Tags

Ulduar uses the provided `datapi` struct tags as each struct field's column name. Below is an example

```go
    type SomeStruct struct{
        SomeString string `datapi:"someString"`
        SomeInt int64 `datapi:"some_int"`
        SomeTime time.Time `datapi:"date"`
        SomeFloat float64 `datapi:"myNumber"`
    }
```

### Table

```go
    err := api.CreateTable(&ulduar.Table{
        Options: &ulduar.TableOptions{
            Collections: SomeStruct{},
            TableName: &someTable,
        }
    })
```

### Insert

```go
    err := api.InsertRecord(&ulduar.Insert{
        Options: &ulduar.InsertOptions{
            Collection: SomeStruct,
        }
    })
```

### Update

```go
    condition := "id = 1"

    err := api.UpdateRecord(&ulduar.Update{
        Options: &ulduar.UpdateOptions{
            Collection: SomeStruct,
            Condition: &condition
        }
    })
```

### Upsert

```go
    target := "some_column"

    err := api.UpsertRecord(&ulduar.Update{
        Options: &ulduar.UpdateOptions{
            Collection: SomeStruct,
            Target: &target
        }
    })
```

## Want to Support?

I am happy to accept any help from others. If you would like to help with the project, please feel free to create a pull request, but with each pull request, it is expected to have proper testing of each functionality. Any requests without any tests, will be denied.
