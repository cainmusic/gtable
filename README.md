# gtable
go tool to make table

## 一般用法

``` go
func normal() {
	table := gtable.NewTable()

	table.AppendTitle("this is a title")
	table.AppendHead([]string{"h1", "h2", "h3"})
	table.AppendBody([]string{"1", "2", "3"})
	table.AppendBody([]string{"123456", "7890abc", "defghijklmnopqrstuvwxyz"})

	table.PrintData()
}
```

```
stdout

+--------------------------------------+
|this is a title                       |
+------+-------+-----------------------+
|h1    |h2     |h3                     |
+------+-------+-----------------------+
|1     |2      |3                      |
|123456|7890abc|defghijklmnopqrstuvwxyz|
+------+-------+-----------------------+
```

查看`./gtable_test.go`了解更多用法。

## 写文件

``` go
func writeFile() {
	table := gtable.NewTable()

	table.AppendTitle("this is a title")
	table.AppendTitle("subtitle")
	table.AppendHead([]string{"h1", "h2", "h3"})
	table.AppendHead([]string{"h21", "h22", "h23"})
	table.AppendBody([]string{"123456", "7890abc", "defghijklmnopqrstuvwxyz"})
	table.AppendBody([]string{"1", "2", "3"})
	table.AppendBody([]string{"4", "5", "6"})
	table.AppendBody([]string{"7", "8", "9"})

	table.SetOutputFile("./out.txt")

	table.PrintData()
}
```

```
cat out.txt

+--------------------------------------+
|this is a title                       |
|subtitle                              |
+------+-------+-----------------------+
|h1    |h2     |h3                     |
|h21   |h22    |h23                    |
+------+-------+-----------------------+
|123456|7890abc|defghijklmnopqrstuvwxyz|
|1     |2      |3                      |
|4     |5      |6                      |
|7     |8      |9                      |
+------+-------+-----------------------+
```

## 读csv文件

```go
func readCsvFile() {
	table := gtable.NewTable()

	table.InitInputFromFile("./test.csv", gtable.DataTypeCsv)
	table.ReadFromInput()

	table.PrintData()
}
```

```
stdout

+-------+-----+-----+-----+
|type   |count|price|limit|
+-------+-----+-----+-----+
|type   |count|price|limit|
|chicken|15   |25   |80   |
|duck   |10   |20   |100  |
|goose  |5    |30   |50   |
+-------+-----+-----+-----+
```

> 注：  
> 初始化`input`时不论读文件还是下面的读字符串最终都会初始化为`io.Reader`。  
> 于是，读csv文件可以，读json文件也是可以的。  
> 同理，读json字符串可以，读csv字符串也是可以的。


## 读json字符串

``` go
const s = `
	{"No.": 1, "Name": "Ed", "Text": "Knock knock."}
	{"No.": 2, "Name": "Sam", "Text": "Who's there?"}
	{"No.": 3, "Name": "Ed", "Text": "Go fmt."}
	{"No.": 4, "Name": "Sam", "Text": "Go fmt who?"}
	{"No.": 5, "Name": "Ed", "Text": "Go fmt yourself!"}
`

func readJsonString() {
	table := gtable.NewTable()

	table.InitInputFromString(s, gtable.DataTypeJson)
	table.ReadFromInput()

	table.PrintData()
}
```

```
stdout

+---+----+----------------+
|No.|Name|Text            |
+---+----+----------------+
|1  |Ed  |Knock knock.    |
|2  |Sam |Who's there?    |
|3  |Ed  |Go fmt.         |
|4  |Sam |Go fmt who?     |
|5  |Ed  |Go fmt yourself!|
+---+----+----------------+

也可能是

+----------------+---+----+
|Text            |No.|Name|
+----------------+---+----+
|Knock knock.    |1  |Ed  |
|Who's there?    |2  |Sam |
|Go fmt.         |3  |Ed  |
|Go fmt who?     |4  |Sam |
|Go fmt yourself!|5  |Ed  |
+----------------+---+----+

也可能是

+----+----------------+---+
|Name|Text            |No.|
+----+----------------+---+
|Ed  |Knock knock.    |1  |
|Sam |Who's there?    |2  |
|Ed  |Go fmt.         |3  |
|Sam |Go fmt who?     |4  |
|Ed  |Go fmt yourself!|5  |
+----+----------------+---+
```

由于`json`解码的时候生成`map`的`key`顺序不确定，所以会有多种结果，下一节解决这个问题。

## 设定title和head

### json

为解决上面的问题我们可以提前设定head。(同时也可以设定title)

``` go
func readJsonStringAndSetHead() {
	table := gtable.NewTable()

	table.InitInputFromString(s, gtable.DataTypeJson)
	table.SetTitle("some dialog")
	// SetHead会修改input，需要在Init方法后执行
	table.SetHead([]string{"No.", "Name", "Text"})
	table.ReadFromInput()

	table.PrintData()
}
```

```
stdout

+-------------------------+
|some dialog              |
+---+----+----------------+
|No.|Name|Text            |
+---+----+----------------+
|1  |Ed  |Knock knock.    |
|2  |Sam |Who's there?    |
|3  |Ed  |Go fmt.         |
|4  |Sam |Go fmt who?     |
|5  |Ed  |Go fmt yourself!|
+---+----+----------------+
```

尝试把`table.SetHead([]string{"No.", "Name", "Text"})`改为`table.SetHead([]string{Name", "Text"})`，会得到

```
stdout

+---------------------+
|some dialog          |
+----+----------------+
|Name|Text            |
+----+----------------+
|Ed  |Knock knock.    |
|Sam |Who's there?    |
|Ed  |Go fmt.         |
|Sam |Go fmt who?     |
|Ed  |Go fmt yourself!|
+----+----------------+
```

### csv

``` go
func readCsvFileAndSetHead() {
	table := gtable.NewTable()

	table.InitInputFromFile("./test_no_head.csv", gtable.DataTypeCsv)
	table.SetHead([]string{"type", "count", "price", "limit"})
	table.ReadFromInput()

	table.PrintData()
}
```

```
stdout

+-------+-----+-----+-----+
|type   |count|price|limit|
+-------+-----+-----+-----+
|chicken|15   |25   |80   |
|duck   |10   |20   |100  |
|goose  |5    |30   |50   |
+-------+-----+-----+-----+
```
