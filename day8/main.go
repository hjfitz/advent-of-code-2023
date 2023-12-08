package day8

import (
	"github.com/hjfitz/advent-of-code-2023/utils"
)

/*
const SAMPLE_DATA = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`
*/

const SAMPLE_DATA = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

// create a tree and return the root node
func parseNodes(nodeMap []string) *Node {
	rawNodes := map[string]*Node{}
	var root *Node
	// create the raw nodes
	for _, node := range nodeMap {
		nodeName := node[0:3]
		newNode := &Node{Name: nodeName}
		rawNodes[nodeName] = newNode
		utils.LogDebug("Node name: %s\n", nodeName)
	}

	for _, node := range nodeMap {
		nodeName := node[0:3]
		cur := rawNodes[nodeName]

		if nodeName == "AAA" {
			root = cur
		}

		leftName := node[7:10]
		rightName := node[12:15]

		cur.Left = rawNodes[leftName]
		cur.Right = rawNodes[rightName]

		utils.LogDebug("leftName: %s, rightName: %s\n", leftName, rightName)
	}

	utils.LogDebug("rawNodes: %+v\n", rawNodes)
	utils.LogDebug("\n\n%+v\n\n", root)

	return root

}

func traverseNodes(instructions string, rootNode *Node) int {
	utils.LogDebug("instructions: %s\n", instructions)
	curInstructionIdx := 0
	atEnd := false
	totalIterations := 0
	curNode := rootNode
	for !atEnd {
		curInstruction := string(instructions[curInstructionIdx])
		from := curNode.Name
		switch curInstruction {
		case "L":
			curNode = curNode.Left
		case "R":
			curNode = curNode.Right
		}

		utils.LogDebug("From: %s, To: %s, dir: %s\n", from, curNode.Name, curInstruction)

		if totalIterations >= 100 {
			//return 1
		}

		curInstructionIdx = (curInstructionIdx + 1) % len(instructions)
		totalIterations += 1

		atEnd = curNode.Name == "ZZZ"

	}
	return totalIterations
}

func Part1(lines []string) int {
	rootNode := parseNodes(lines[2:])
	count := traverseNodes(lines[0], rootNode)

	return count
}

func Run() {
	parts := []func([]string) int{
		Part1,
	}

	utils.Runner(8, SAMPLE_DATA, parts, false)
}
