package main

import (
	"container/list"
	"fmt"
)

type Player struct {
	name  string
	score int
}

type RankList struct {
	Player
	rank int
}

func main() {
	l := list.New()
	sync(l, Player{
		name:  "a",
		score: 98,
	})
	sync(l, Player{
		name:  "b",
		score: 99,
	})
	sync(l, Player{
		name:  "c",
		score: 99,
	})
	sync(l, Player{
		name:  "d",
		score: 100,
	})
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	query(l, Player{
		name:  "d",
		score: 100,
	})
	for _, node := range query(l, Player{
		name:  "d",
		score: 100,
	}) {
		fmt.Println(node)
	}
}

func sync(root *list.List, player Player) {
	node := root.Front()
	if root.Len() == 0 {
		root.PushFront(player)
		return
	}
	for i := 0; i < root.Len(); i++ {
		if player.score > node.Value.(Player).score {
			root.InsertBefore(player, node)
			return
		} else {
			if node.Next() == nil {
				root.InsertAfter(player, node)
				return
			}
			node = node.Next()
		}
	}
}

func query(root *list.List, player Player) []*RankList {
	node := root.Front()
	rank := 1
	count := 10
	rankList := make([]*RankList, 0)
	for node != nil {
		if node.Value.(Player).name == player.name {
			pre := node.Prev()
			for pre != nil && count > 5 {
				count--
			}
			for i := 0; i < 10-count; i++ {
				rankList = append(rankList, &RankList{
					Player: pre.Value.(Player),
					rank:   rank - (10 - count - i),
				})
			}
			rankList = append(rankList, &RankList{
				Player: node.Value.(Player),
				rank:   rank,
			})
			next := node.Next()
			for count > 0 && next != nil {
				count--
				rank++
				rankList = append(rankList, &RankList{
					Player: next.Value.(Player),
					rank:   rank,
				})
				next = next.Next()
			}
			return rankList
		}
		node = node.Next()
		rank++
	}

	return nil
}
