package test

import (
	"fmt"
)

// 幂律科技
// 现在数据库有一张表，用来存储一个多叉树，id为主键，pid 表示父节点的 id，已知 "-1" 表示根节点，现在要求打印出从根节点到每个子节点的路径（可以是无序的）。
//
// | id      | pid    |
// |---------|--------|
// | "A"     | "-1"   |
// | "A-1"   | "A"    |
// | "A-2"   | "A"    |
// | "A-3"   | "A"    |
// | "A-2-1" | "A-2"  |
// | "A-2-2" | "A-2"  |
// | "A-2-3" | "A-2"  |
//
// Input: [
//   {
//       "id": "A",
//       "pid": "-1"
//   },
//   {
//       "id": "A-1",
//       "pid": "A"
//   },
//   {
//       "id": "A-2",
//       "pid": "A"
//   },
//   {
//       "id": "A-3",
//       "pid": "A"
//   },
//   {
//       "id": "A-2-1",
//       "pid": "A-2"
//   },
//   {
//       "id": "A-2-2",
//       "pid": "A-2"
//   },
//   {
//       "id": "A-2-3",
//       "pid": "A-2"
//   }
// ]
// Output: [
//   "/A",
//   "/A/A-1",
//   "/A/A-2",
//   "/A/A-3",
//   "/A/A-2/A-2-1",
//   "/A/A-2/A-2-2",
//   "/A/A-2/A-2-3",
// ]

type Node struct {
	ID  string
	PID string
}

func main1() {
	nodes := []Node{
		{
			"A-1",
			"A",
		},
		{
			"A",
			"-1",
		},
		{
			"A-2",
			"A",
		},
		{
			"A-3",
			"A",
		},
		{
			"A-2-1",
			"A-2",
		},
		{
			"A-2-2",
			"A-2",
		},
		{
			"A-2-3",
			"A-2",
		},
	}
	output := make([]string, 0, len(nodes))
	// 存储数据关系
	m := make(map[string]string, len(output))
	for _, v := range nodes {
		m[v.ID] = v.PID
	}
	//fmt.Println(m)
	root := "-1"
	for _, v := range nodes {
		tmp := v.ID
		tmpArr := []string{}
		for tmp != root {
			tmpArr = append([]string{tmp}, tmpArr...)
			tmp = m[tmp]
		}
		str := ""
		for _, v := range tmpArr {
			str = fmt.Sprintf("%s/%s", str, v)
		}
		output = append(output, str)
	}
	fmt.Println(output)
}
