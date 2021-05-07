package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。

	map1 := make(map[int]int, 10)
	var map2 = make(map[int]int, 20)
	var map3 map[int]int
	var map4 = map[int]int{}
	map5 := map[int]int{
		0:0,
		1:1,
	}

	// map类型的变量默认初始值为nil，需要使用make()函数来分配内存。
	// 其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。
	//map1[1] = 1
	//map2[1] = 1
	// map3[1] = 1
	fmt.Println(map1, map2, map3, map4, map5)

	// Go语言中有个判断map中键是否存在的特殊写法
	// value, ok := map[key]
	// Go语言中使用for range遍历map。
	// for k, v := range map {
	// }
	// 注意： 遍历map时的元素顺序与添加键值对的顺序无关。

	// 使用delete()函数删除键值对
	// 使用delete()内建函数从map中删除一组键值对，delete()函数的格式如下：
	// delete(map, key)
	// map:表示要删除键值对的map
	// key:表示要删除的键值对的键


	// 按照指定顺序遍历map
	// 随机数种子
	rand.Seed(time.Now().UnixNano())
	// 创建一个长度为2000的map[string]int
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(1000)
		scoreMap[key] = value
	}

	keys := make([]string, 0, 100)
	for k, _ := range scoreMap {
		keys = append(keys, k)
	}

	// 对切片进行排序
	sort.Strings(keys)
	// 按照排序 遍历map
	for _, key := range keys {
		fmt.Printf("key:%s, value:%d\n", key, scoreMap[key])
	}

	// 元素为map类型的切片
	// 新建一个 mapSlice
	var mapSlice = make([]map[string]string, 2)
	mapSlice[0] = make(map[string]string, 3)
	mapSlice[0]["key0"] = "value0"
	mapSlice[0]["key1"] = "value1"
	mapSlice[0]["key2"] = "value2"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	// 值为切片类型的map
	var sliceMap = make(map[string]([]string), 3)
	value := make([]string, 0, 2)
	value = append(value, "guangzhou", "shenzhen")
	key := "china"
	sliceMap[key] = value
	fmt.Println(sliceMap)

	// 初始化一个map
	var mapInit = map[string]string{
		"xiaoli" : "湖南",
		"xiaoliu" : "天津",
	}

	var mapTemp map[string]string
	mapTemp = make(map[string]string, 10)
	mapTemp["xiaoli"] = "湖南"
	mapTemp["xiaoliu"] = "天津"
	fmt.Println(mapInit, mapTemp)



	// key，value存储
	// 最通俗的话说Map是一种通过key来获取value的一个数据结构，其底层存储方式为数组，
	// 在存储时key不能重复，当key重复时，value进行覆盖，我们通过key进行hash运算
	//（可以简单理解为把key转化为一个整形数字）然后对数组的长度取余，得到key存储在数组的哪个下标位置，
	// 最后将key和value组装为一个结构体，放入数组下标处，看下图：

	// length = len(array) = 4
	// hashkey1 = hash(xiaoming) = 4
	// index1  = hashkey1 % length= 0
	// hashkey2 = hash(xiaoli) = 6
	// index2  = hashkey2 % length= 2


	// 开放定址法：也就是说当我们存储一个key，value时，发现hashkey(key)的下标已经被别key占用，
	// 那我们在这个数组中空间中重新找一个没被占用的存储这个冲突的key，那么没被占用的有很多，找哪个好呢？
	// 常见的有线性探测法，线性补偿探测法，随机探测法，这里我们主要说一下线性探测法
	// 线性探测，字面意思就是按照顺序来，从冲突的下标处开始往后探测，到达数组末尾时，从数组开始处探测，直到找到一个空位置存储这个key，
	// 当数组都找不到的情况下回扩容（事实上当数组容量快满的时候就会扩容了）；
	// 查找某一个key的时候，找到key对应的下标，比较key是否相等，如果相等直接取出来，否则按照顺寻探测直到碰到一个空位置，说明key不存在。

	// 拉链法：何为拉链，简单理解为链表，当key的hash冲突时，我们在冲突位置的元素上形成一个链表，
	// 通过指针互连接，当查找时，发现key冲突，顺着链表一直往下找，直到链表的尾节点，找不到则返回空，如下图：


	// Go中Map的实现原理
	// map同样也是数组存储的的，每个数组下标处存储的是一个bucket，
	// 每个bucket中可以存储8个kv键值对，当每个bucket存储的kv对到达8个之后，会通过overflow指针指向一个新的bucket，
	// 从而形成一个链表, kv的结构和overflow指针啊，事实上，这两个结构体并没有显示定义，是通过指针运算进行访问的。




	}


