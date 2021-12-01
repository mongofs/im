package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*
	1
  2	   3
4 5   6 7

*/


func main(){
	tr :=&TreeNode{Val: 1}
	tr.Left=&TreeNode{Val: 2}
	tr.Left.Left=&TreeNode{Val: 4}
	tr.Left.Right = &TreeNode{Val: 5}
	tr.Right=&TreeNode{Val: 3}
	tr.Right.Left = &TreeNode{Val: 6}
	tr.Right.Right =&TreeNode{Val: 7}

	fmt.Println(inorderTraversal1(tr))
}


func inorderTraversal(root *TreeNode)  (res []int){
	var orderFunc func(node * TreeNode)
	orderFunc = func(node *TreeNode) {
		if root ==nil {return }
		orderFunc(node.Left)
		res = append(res,node.Val)
		orderFunc(node.Right)
	}
	orderFunc(root)
	return
	// out put : 4251637
}


func inorderTraversal1(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}




/*
	1
  2	   3
4 5   6 7
*/
func preTraversal(root *TreeNode)  {
	if root == nil {return }
	fmt.Println(root)
	printpre(root.Left)
	printpre(root.Right)
	// out put : 1245367
}



func printpre(node *TreeNode) {
	if node == nil {return }
	fmt.Println(node)
	print(node.Left)
	print(node.Right)
}




/*
	1
  2	   3
4 5   6 7
*/
func behandTraversal(root *TreeNode)  {
	if root == nil {return }

	behandTraversal(root.Left)
	behandTraversal(root.Right)
	fmt.Println(root)
	// out put : 4526731
}



func printbehand(node *TreeNode) {
	if node == nil {return }
	print(node.Left)
	print(node.Right)
	fmt.Println(node)
}