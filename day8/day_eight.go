package main

import (
	"first/advent_of_code"
	"strings"
	"strconv"
	"fmt"
)

type LicenseNode struct {
	ChildCount int
	EntryCount int
	Children []LicenseNode
	EntryLength int
	MetadataEntries []int
}

func loadLicenseItems() []int {
	res := []int{}
	strs := strings.Split(advent_of_code.LoadInputTxt(8), " ")
	for _, str := range strs {
		r, _ := strconv.ParseInt(str, 10, 32)
		res = append(res, int(r))
	}
	return res
}

func metadataEntrySum(licenseItems []int) int {
	for i := 0; i < len(licenseItems); i++ {

	}

	return -1
}

func parseLicenseNode(licenseItems []int) LicenseNode {
	header := licenseItems[0:2]
	childrenCount := header[0]
	children := []LicenseNode{}
	length := 2 + header[1]
	parsedLength := 2
	for i:= 0;i < childrenCount;i++ {
		child := parseLicenseNode(licenseItems[parsedLength:len(licenseItems) - header[1]])
		length += child.EntryLength
		parsedLength += child.EntryLength
		children = append(children, child)
	}
	metadataEntries := licenseItems[length - header[1]: length]
	return LicenseNode{
		ChildCount: childrenCount,
		EntryCount: header[1],
		Children: children,
		EntryLength: length,
		MetadataEntries: metadataEntries,
	}
}

func getMetaSum(licenseNode LicenseNode) int {
	metaSum := 0
	for i := 0; i < len(licenseNode.MetadataEntries); i++ {
		metaSum += licenseNode.MetadataEntries[i]
	}
	for i := 0; i < licenseNode.ChildCount; i++ {
		metaSum += getMetaSum(licenseNode.Children[i])
	}
	return metaSum
}

func getNodeValue(licenseNode LicenseNode) int {
	value := 0
	if len(licenseNode.Children) == 0 {
		for _, entry := range licenseNode.MetadataEntries {
			value += entry
		}
		return value
	}
	for _, index := range licenseNode.MetadataEntries {
		if len(licenseNode.Children) >= index {
			value += getNodeValue(licenseNode.Children[index - 1])
		}
	}
	return value
}

func main() {
	fmt.Println(loadLicenseItems())
	root := parseLicenseNode(loadLicenseItems())
	fmt.Println(root)
	fmt.Println(getMetaSum(root))
	fmt.Println(getNodeValue(root))
}
