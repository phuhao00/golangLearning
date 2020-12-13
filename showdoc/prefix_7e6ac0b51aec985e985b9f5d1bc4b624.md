# 1、map
	1）for-range陷阱
```go

type student struct {
	Name string
	Age  int
}
func pase_student() map[string]*student {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	return m
}
func main() {
	students := pase_student()
	for k, v := range students {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}
}
```
for循环使用stu遍历时，stu只是一个临时变量，遍历过程中指针地址不变，所以后面的赋值都是指向了同一个内存区域，导致最后所有的信息都一致。输出key=zhou,value=&{wang 22} 
key=li,value=&{wang 22} 
key=wang,value=&{wang 22}
# 2、slice
	1）append函数
```go
func main() {
	s := make([]int, 5)
	fmt.Println(s)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
```

	2)切片的数据截取
```go
func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := a[2:5]  
	s1[1] = 666   
	fmt.Println("s1 = ", s1)   
	fmt.Println("a = ", a)    

	s2 := s1[2:7]   
	fmt.Println("s2====",s2)
}
```
注意：新的切片和原数组或原切片共用的是一个底层数组，所以当修改的时候，底层数组的值就会被改变，所以原切片的值也改变

	3）copy函数使用
```go
func main() {
	slice1 := []int{1, 2,3,4,5}
	slice2 := []int{5, 4, 3}
	copy(slice1, slice2)
	fmt.Println(slice1)
	copy(slice2, slice1)
	fmt.Println(slice2)
}
```
copy(切片1，切片2)  将切片2中的元素拷贝到切片1中

	4）切片作为函数参数
```go
func main()  {
	s := make([]int, 10)
	InitData(s)
	fmt.Println(s)
}

func InitData(s []int)  {
	for i:=1;i<11;i++{
		s[i-1] = i
	}
}

func main()  {
	var s [10]int
	InitData(s)
	fmt.Println(s)
}

func InitData(s [10]int)  {
	for i:=1;i<11;i++{
		s[i-1] = i
	}
}
```
注意：在GO语言中，数组作为参数进行传递是值传递，而切片作为参数进行传递是引用传递。

	5）切片排序 sort包主要针对[]int、[]float64、[]string、以及其他自定义切片的排序。
sort 包 在内部实现了四种基本的排序算法：插入排序（insertionSort）、归并排序（symMerge）、堆排序（heapSort）和快速排序（quickSort）； sort 包会依据实际数据自动选择最优的排序算法。复杂数据切片要实现less swap len3个函数

```go
type Animals []string

func (a Animals) Len() int           { return len(a) }
func (a Animals) Less(i, j int) bool { return len(a[i]) < len(a[j]) }
func (a Animals) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	animals := []string{"cat", "bird", "zebra", "fox"}
	// Sort by strings.
	sort.Strings(animals)
	fmt.Println(animals)//[bird cat fox zebra]

	//sort by len
	an := Animals{"cat", "bird", "zebra", "fox"}
	sort.Sort(an)
	fmt.Println(an)//[cat fox bird zebra]
}
```
	6）map和切片联系使用demo
		（1）map 有序输出
```go
func main() {
	var testMap = make(map[string]string)
	testMap["Bda"] = "B"
	testMap["Ada"] = "A"
	testMap["Dda"] = "D"
	testMap["Cda"] = "C"
	testMap["Eda"] = "E"

	for key, value := range testMap {
		fmt.Println(key, ":", value)
	}
	var testSlice []string
	testSlice = append(testSlice, "Bda", "Ada", "Dda", "Cda", "Eda")

	/* 对slice数组进行排序，然后就可以根据key值顺序读取map */
	sort.Strings(testSlice)
	fmt.Println("排序后输出:")
	for _, Key := range testSlice {
		/* 按顺序从MAP中取值输出 */
		if Value, ok := testMap[Key]; ok {
			fmt.Println(Key, ":", Value)
		}
	}
}
```
		（2）切片去重
```go
func main() {
	str := []string{"a", "b", "a"}
	d := removeRepByMap(str)
	fmt.Println(d)
}

func removeRepByMap(slc []string) []string {
	//存放返回的不重复切片
	result := []string{}
	// 存放不重复主键
	tempMap := map[string]byte{}
	for _, e := range slc {
		l := len(tempMap)
		//当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		tempMap[e] = 0
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			//当元素不重复时，将元素添加到切片result中
			result = append(result, e)
		}
	}
	return result
}
```

# 3、Iota
	iota是golang语言的常量计数器,只能在常量的表达式中使用。iota在const关键字出现时将被重置为0(const内部的第一行之前)，const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。使用iota能简化定义，在定义枚举时很有用。
  	iota是一个从0开始递增的整形值iota可以用在const定义块的任何位置，并且它的当前值取决于它所在的位置。
	
```go
const (
	a = 4
	b = 5
	c = iota
	d
)

func main() {
	fmt.Println(a,b,c, b)
}
```

