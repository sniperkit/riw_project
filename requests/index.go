package requests

import (
	"fmt"
	"../utils"
	"../inversers"
)

// Index and BufferIndex are different structure so that to be sure to avoid conflicts between the index used by the user and the one used to build it.
type Index struct {
	folderPath string
	corpusSize int
	postingLists map[string]inversers.PostingList
	docIDToFilePath map[int]string
}

func NewIndex (folderPath string) *Index {
	docIDToFilePath := make(map[int]string)
	err := utils.ReadGob("./saved/IDToPath.meta", &docIDToFilePath)
	if err != nil {
		panic(err)
	}
	postingLists := make(map[string]inversers.PostingList)
	return &Index{
		folderPath: folderPath,
		docIDToFilePath: docIDToFilePath,
		postingLists: postingLists,
	}
}

func (index *Index) GetTerms(terms []string) (error, []inversers.PostingList) {
	output := make([]inversers.PostingList, len(terms))
	err := index.LoadTerms(terms)
	for i, term := range terms {
		output[i] = index.postingLists[term]
	}
	return err, output
}

func (index *Index) LoadTerms(terms []string) error {
	for _, term := range terms {
		err := index.loadTerm(term)
		if err != nil {
			fmt.Printf("%s is not in the index, it won't be taken into account\n", term)
		}
	}
	return nil
}

func (index *Index) loadTerm(term string) error {
	termFile := fmt.Sprintf("./saved/%s.postings", term)
	err, postingList := inversers.PostingListFromFile(termFile)	
	if err != nil {
		return err
	}
	index.postingLists[term] = postingList
	return nil
}

// TODO: make it safe
func (index *Index) unloadTerm(term string) {
	index.postingLists[term] = nil
}

func (index *Index) PrintPostings() {
	for _, postingList := range index.postingLists {
		fmt.Printf("%v\n", postingList)
	}
}