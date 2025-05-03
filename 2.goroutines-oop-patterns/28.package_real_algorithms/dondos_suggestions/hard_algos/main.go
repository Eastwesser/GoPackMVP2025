package main

import (
	"fmt"
	"sort"
)

// 1. Бинарный поиск
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 2. Сортировка
func sortingExample() {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6}
	sort.Ints(arr)
	fmt.Println("Sorted Array:", arr)
}

// 3. Два указателя
func twoPointersExample(arr []int, target int) (int, int) {
	left, right := 0, len(arr)-1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			return arr[left], arr[right]
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return -1, -1
}

// 4. Скользящее окно
func slidingWindowExample(arr []int, k int) int {
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}
	maxSum := windowSum
	for i := k; i < len(arr); i++ {
		windowSum += arr[i] - arr[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum
}

// 5. BFS (поиск в ширину)
func bfs(graph map[string][]string, start string) {
	queue := []string{start}
	visited := make(map[string]bool)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if !visited[node] {
			fmt.Println("BFS Visited:", node)
			visited[node] = true
			for _, neighbor := range graph[node] {
				queue = append(queue, neighbor)
			}
		}
	}
}

// 6. DFS (поиск в глубину)
func dfs(graph map[string][]string, node string, visited map[string]bool) {
	if !visited[node] {
		fmt.Println("DFS Visited:", node)
		visited[node] = true
		for _, neighbor := range graph[node] {
			dfs(graph, neighbor, visited)
		}
	}
}

// 7. Бэктрекинг (задача о ферзях)
func solveNQueens(n int) [][]string {
	result := [][]string{}
	board := make([]int, n)
	backtrack(board, 0, &result)
	return result
}

func backtrack(board []int, col int, result *[][]string) {
	if col == len(board) {
		*result = append(*result, formatBoard(board))
		return
	}
	for i := 0; i < len(board); i++ {
		if isSafe(board, i, col) {
			board[col] = i
			backtrack(board, col+1, result)
			board[col] = -1
		}
	}
}

func isSafe(board []int, row, col int) bool {
	for i := 0; i < col; i++ {
		if board[i] == row || abs(board[i]-row) == abs(i-col) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func formatBoard(board []int) []string {
	n := len(board)
	result := make([]string, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			if board[j] == i {
				row[j] = 'Q'
			} else {
				row[j] = '.'
			}
		}
		result[i] = string(row)
	}
	return result
}

func main() {
	// Бинарный поиск
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Binary Search:", binarySearch(arr, 5))

	// Сортировка
	sortingExample()

	// Два указателя
	fmt.Println("Two Pointers:", twoPointersExample(arr, 9))

	// Скользящее окно
	fmt.Println("Sliding Window:", slidingWindowExample(arr, 3))

	// BFS
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"D"},
		"C": {"E"},
		"D": {},
		"E": {},
	}
	bfs(graph, "A")

	// DFS
	visited := make(map[string]bool)
	dfs(graph, "A", visited)

	// Бэктрекинг
	fmt.Println("N-Queens:", solveNQueens(4))
}
