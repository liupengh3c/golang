package ch4

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Ch4 第四章测试函数
func Ch4() {
	// s := [2]int{4, 8}
	// sha()
	// 数组指针测试
	// arrayP(&s)
	// popcount(14)
	// appendInt()
	// noempty()
	sort()
	makeWheel()
	// search()
	searchMovie()
}

// Sha 加密
func sha() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n", c1)
	fmt.Printf("%x\n", c2)
}

// arrayP 数组指针
func arrayP(p *[2]int) {
	for i := range p {
		p[i] = 32
	}
	fmt.Println(p)
}

func popcount(x int64) {
	var pc [256]byte
	cnt := 0
	for k := range pc {
		pc[k] = pc[k/2] + byte(k&1)
		fmt.Printf("%d-%d\n", k, pc[k])
	}
	cnt = int(pc[byte(x>>(0*8))] + pc[byte(x>>(1*8))] + pc[byte(x>>(2*8))] + pc[byte(x>>(3*8))] + pc[byte(x>>(4*8))] + pc[byte(x>>(5*8))] + pc[byte(x>>(6*8))] + pc[byte(x>>(7*8))])
	fmt.Println(cnt)
	// fmt.Printf("%T %[1]d", pc)
}

func appendInt() {
	var x []int
	fmt.Println(len(x), cap(x))
	for i := 0; i < 20; i++ {
		x = append(x, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(x), x)
	}
}

func noempty() {
	s := []string{"ab", "", "art"}
	// fmt.Println(nonempty(s))
	// fmt.Printf("%q\n", nonempty(s))
	fmt.Println("nonempt2 function test,appen func")
	fmt.Printf("%q\n", nonempty2(s))
}
func nonempty(str []string) []string {
	i := 0
	for _, v := range str {
		if v != "" {
			str[i] = v
			i++
		}
	}
	return str[:i]
}

func nonempty2(str []string) []string {
	var s []string
	for _, v := range str {
		if v != "" {
			fmt.Println(v)
			s = append(s, v)
		}
	}
	fmt.Println(s)
	return s
}

// 二叉树实现实现插入排序，左小，右大
type tree struct {
	value int32
	left  *tree
	right *tree
}

// sort 排序
func sort() {
	var data []int32
	var s []int32
	var root *tree
	// rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		data = append(data, rand.Int31n(100))
	}
	fmt.Println(data)
	for _, val := range data {
		root = add(root, val)
		fmt.Println(root.value)
	}
	// fmt.Println(root.value)
	// fmt.Println(root.left.value)
	// fmt.Println(root.right.value)
	s = appendVal(s, root)
	fmt.Println(s)
}

// add 将value插入二叉树中，左小右大
func add(t *tree, value int32) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		t.left = nil
		t.right = nil
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// appendVal 将二叉树中的数据放入slice中
func appendVal(values []int32, t *tree) []int32 {
	if t != nil {
		values = appendVal(values, t.left)
		values = append(values, t.value)
		values = appendVal(values, t.right)
	}
	return values
}

/*************************结构体嵌套/匿名成员**************************/
// success is the ability to go from one failure to another with no loss of enthusiasm
// 定义一个点
type point struct {
	X, Y int32
}

// 定义一个圆
type circle struct {
	point
	Radius int32
}

// 定义一个轮子
type wheel struct {
	circle
	spokes int32 // 条辅个数
}

// makeWheel 实例一个5辐条的轮子
func makeWheel() {
	var wh wheel
	wh.X = 0
	wh.Y = 0
	wh.Radius = 10
	wh.spokes = 5
	fmt.Printf("%#v\n", wh)
	fmt.Printf("%+v\n", wh)
	// json 打印出来
	s, _ := json.Marshal(wh)
	fmt.Printf("%s\n", s)

	s, _ = json.MarshalIndent(wh, "", "	")
	fmt.Printf("%s\n", s)
}

const issueUrl = "https://api.github.com/search/issues"

type issuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*issues
}

//
type issues struct {
	Number   int
	HTMLURL  string `json:"html_url"`
	Title    string
	State    string
	User     *user
	CreateAt time.Time `json:"created_at"`
	Body     string
}
type user struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func searchIssues(terms []string) (*issuesSearchResult, error) {
	var result issuesSearchResult
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(issueUrl + "?=" + q + "&sort=created&order=asc")
	if err != nil {
		fmt.Println("req github error")
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println()
	json.NewDecoder(resp.Body).Decode(&result)
	resp.Body.Close()
	return &result, nil
}

type movie struct {
	Month string `json:"month"`
	// Num        int32  `json:"num"`
	// Link       string `json:"link"`
	// Year       string `json:"year"`
	// News       string `json:"news"`
	// SafeTitle  string `json:"safe_title"`
	// Transcript string `json:"transcript"`
	// Alt        string `json:"alt"`
	// Img        string `json:"img"`
	Title string `json:"tilte"`
	// Day        string `json:"day"`
}

func searchMovie() {
	mov := new(movie)
	strUrl := "http://xkcd.com/571/info.0.json"
	resp, err := http.Get(strUrl)
	if err != nil {
		fmt.Println("req url err " + err.Error())
		return
	}
	err = json.NewDecoder(resp.Body).Decode(mov)
	fmt.Println(resp.Status)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(mov.Title)
}
func search() {
	result, err := searchIssues(os.Args[1:])
	if err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("%d issues", result.TotalCount))
	for _, val := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", val.Number, val.User.Login, val.Title)
	}
	return
}
