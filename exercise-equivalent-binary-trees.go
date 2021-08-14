package main

import "golang.org/x/tour/tree"

//ウォークはツリーを歩き、すべての値を送信します
//ツリーからチャネルchへ。
func Walk(t *tree.Tree, ch chan int)

//同じことが木かどうかを決定します
// t1とt2には同じ値が含まれています。
func Same(t1, t2 *tree.Tree) bool

func main() {
}
