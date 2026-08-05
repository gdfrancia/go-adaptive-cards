[![GoDoc](https://godoc.org/github.com/DanielTitkov/go-adaptive-cards?status.svg)](https://godoc.org/github.com/DanielTitkov/go-adaptive-cards)

# go-adaptive-cards

Go package for creating adaptive cards in Go.

> [!NOTE]
> It seems the original `DanielTitkov/go-adaptive-cards` seem to have an issue with importing due to mismatch with `go.mod`, I forked this for the sole purpose of fixing this issue, and has also merged changes from `kubeshop/go-adaptive-cards` in lieu to combining their changes by adding the `Table` Type and `Action.Execute` as well, with the addition of adding schema `v1.4` too (including several comment typos). This seems to not be updated for the past 4-6 years, and it seems like a viable tool, but it just goes to waste in terms of being unimportable, so I just made this repo for it.
> I did not update anything from this repo, and credits goes to the original writers of this repo. I just merely merged all the updates from @kubeshop to the original repo by @dani-polani, as I do need this repo for AdaptiveCard v1.4, the one with tables. I do commend the effort made in making this repo for convenience in generating Adaptive Cards in go.

## Purpose

Adaptive cards are JSON and can be very complex. Operating raw JSON is unconvinient, so this packages allows to build cards in Golang. Then they can be converted to JSON, filling in types and validating required fields. More on adaptive cards: https://adaptivecards.io/.

## Usage

Import package:

```go
import (
    cards "github.com/gdfrancia/go-adaptive-cards"
)
```
And define your card:

```go
c := cards.New([]cards.Node{
    &cards.Container{
        Items: []cards.Node{
            &cards.TextBlock{
                Text:     "foo",
                IsSubtle: cards.FalsePtr(),
            },
        },
    },
}, []cards.Node{}).
    WithSchema(cards.DefaultSchema).
    WithVersion(cards.Version12)

s, err := c.StringIndent("", "    ")
```

This will give you JSON that you can use with bot framework and so on: 

```json
{
    "type": "AdaptiveCard",
    "version": "1.2",
    "$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
    "body": [
        {
            "type": "Container",
            "items": [
                {
                    "type": "TextBlock",
                    "text": "foo",
                    "isSubtle": false
                }
            ]
        }
    ]
}
```

See more in examples directory and cards_test.go.

## Limitations

* As yet package does not validate string values except for types (e.g. "bolder" for text weight). Look for supported values in adaptive cards [schema explorer](https://adaptivecards.io/explorer/).
* Not all card attributes are supported by the constructor. By now you can define a card as struct in order to access all attributes.
* Some fields can be both a JSON object or a string (e.g. Inlines of the RichTextBlock). This is not supported. Such field can only be provided as structs.
