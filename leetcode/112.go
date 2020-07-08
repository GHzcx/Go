package main


func hasPathSum(root *TreeNode, sum int) bool {
	return getResult(root, 0, sum)
}

func getResult(root *TreeNode, val int, sum int) bool {
	if root == nil {return false}

	if root.Right == nil && root.Left == nil {
		return (val + root.Val) == sum
	} else if root.Right == nil {
		return getResult(root.Left, val + root.Val, sum)
	} else if root.Left == nil {
		return getResult(root.Right, val + root.Val, sum)
	} else {
		return getResult(root.Left, val + root.Val, sum) || getResult(root.Right, val + root.Val, sum)
	}
}

func main() {

}
