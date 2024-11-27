# Usage

```bash
go run ./cmd/main.go -h
Usage of main.go:
  -f string
        Path to csv or json file containing reports (required).
  -p string
        Period time in YYYYMM format (required).

go run ./cmd/main.go -f ./data/tx.csv -p 202201
go run ./cmd/main.go -f ./data/tx.json -p 202201

# Output
{
    "period": "202201",
    "total_income": 0,
    "total_expenditure": -111000,
    "transactions": [
        {
                "date": "2022/01/05",
                "amount": "-1000",
                "content": "eating out"
        },
        {
                "date": "2022/01/06",
                "amount": "-10000",
                "content": "debit"
        },
        {
                "date": "2022/01/25",
                "amount": "-100000",
                "content": "rent"
        }
    ]
}
```


# Thought Process

Regarding my approach to solving the problem, I carefully reviewed the requirements several times along with the input/output examples. Then, for each requirement, I broke the problem down into smaller, manageable tasks. Afterward, I did a final review to check if I could start coding those parts. If not, I identified any potential obstacles and addressed them before beginning the coding process.

# Technology Choices

For technology selection, I chose the Go programming language because it is the one I am most familiar with.

Since the project involves calculating amount, I used the [decimal](https://github.com/shopspring/decimal) library to ensure precise calculations and avoid issues like `0.1 + 0.2 != 0.3` with decimal amounts. Aside from this library, all other libraries used are either built-in or common libraries, making the coding and debugging process more straightforward.

# Design Decisions

With the given requirements, I aim for a simpler design rather than a complex one. Therefore, I use the Golang common app template to build this project.

This design is straightforward, user-friendly, and easy to modify without difficulties. It also ensures that others can easily read and understand the code.

I also considered adding unit tests, but since all the functions are quite simple, I decided not to include them.

# Requirement Fulfillment

According to the requirements, I had to code a CLI application that accepts two parameters: `periodTime` and `filePath`, with the output being displayed on the screen.

For me, the most challenging part was processing the data in the input file. Since I didn't know how many lines of data the file would contain, I decided to read the file line by line instead of reading it all at once to avoid a sudden increase in memory usage.

Because I couldn't estimate the number of input data lines, I also couldn't optimize the data conversion when outputting to stdout. If the data set is large, this might cause a bottleneck.

Additionally, filtering transactions within a time range has been a challenge. I used `time.Time` to check and parse the time into the correct format, instead of using `regex` or simpler methods like `strings.Contains()`, which made this part of the task more complicated.

# Future Work

I think I can improve the input handling by not limiting it to just local files but also supporting file input via URL. Additionally, I’ve been considering using goroutines to improve data processing time, although I’m not entirely sure about the implementation and its potential impact.

Additionally, I noticed that the `total_income` and `total_expenditure` fields are returned as numbers (currently, I’m using `float64` to return the data). However, for values related to amounts, I believe the backend should return them as strings to avoid precision issues with decimal numbers.
