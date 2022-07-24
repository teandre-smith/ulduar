# ulduar

This collection includes functions that make working with AWS DataAPI a bit easier. There are convenient generation functions that can be used when wanting to generate generic raw sql queries or dealing with AWS DataAPI. The collection is still a work in progress, but at this stage, it is a basic ORM for simple statements. As time goes on, support for more complex queries will be implemented.

Please take note that majority of these functions that generate statements using structs/collections will require a struct tag named `datapi` to formulate the column names.
