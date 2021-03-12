package search

import (
	"encoding/json"
	"os"
)

// TODO >> Go 编译器可以根据赋值运算符右边的值来推导类型，声明常量的时候不需要指定类型。
// TODO >> Go 此外，这个常量的名称使用小写字母开头，表示它只能在 search 包内的代码里直接访问，而不暴露到包外面。
const dataFile = "data/data.json"

// TODO >> 这些数据文档需要解码到一个结构组成的切片里，以便我们能在程序里使用这些数据。来看看用于解码数据文档的结构类型
// TODO >> 我们声明了一个名叫 Feed 的结构类型。这个类型会对外暴露。这个类型里面声明了 3 个字段，每个字段的类型都是字符串，对应于数据文件中各个文档的不同字段。
// TODO >> 每个字段的声明最后 ` 引号里的部分被称作标记（tag）。这个标记里描述了 JSON 解码的元数据，用于创建 Feed 类型值的切片。每个标记将结构类型里字段对应到 JSON 文档里指定名字的字段。
type Feed struct {
	// 名称
	Name string `json:"site"`
	// url
	URI  string `json:"link"`
	// 类型
	Type string `json:"type"`
}

/**
 * 读取 data.json 文件并返回数据源的切片。
 * 这些数据源会输出内容，随后使用各自的匹配器进行搜索。
 * 这个函数读取data.json数据文件，并将每个 JSON 文档解码，存入一个 Feed 类型值的切片里
 */
func RetrieveFeeds() ([]*Feed, error) {
	// TODO >> 打开文件
	// open the file
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	// TODO >> 当函数执行结束时 关闭文件
	// file close
	// function return.
	defer file.Close()

	// TODO >> 将文件解码到一个切片里
	// TODO >> 这个切片的每一项是一个指向一个 Feed 类型值的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
