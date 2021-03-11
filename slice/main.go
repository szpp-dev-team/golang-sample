package main

func main() {
	ary := [5]int{3, 1, 4, 1, 5} // これは生配列
	slice := []int{}             // これはスライス

	ary = append(ary, 3)     // 生配列だから要素の追加はできない
	slice = append(slice, 3) // スライスだから可能

	exArySlice := ary[:]   // 生配列からスライス
	miniSlice := slice[:2] // スライスをスライス
}
