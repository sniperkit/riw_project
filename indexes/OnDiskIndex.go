package indexes

import (
	"fmt"
	"../utils"
)
type OnDiskIndex struct {
	folderPath string
	index *Index
}

// Only for indexes persisted to disk

func OnDiskIndexFromFolder(folderPath string) *OnDiskIndex {
	docIDToFilePath := make(map[int]string)
	err := utils.ReadGob("./saved/meta/iDToPath", &docIDToFilePath)
	if err != nil {
		panic(err)
	}
	postingLists := make(map[string]PostingList)
	index := &Index{
		postingLists: postingLists,
		docIDToFilePath: docIDToFilePath,
	}
	return &OnDiskIndex{
		folderPath: folderPath,
		index: index,
	}
}

func (odi *OnDiskIndex) loadTerm(term string) error {
	termFile := fmt.Sprintf("./saved/postings/%s", term)
	_, exists := odi.index.postingLists[term]
	if !exists {
		err, postingList := PostingListFromFile(termFile)	
		if err != nil {
			return err
		}
		odi.index.postingLists[term] = postingList
	}
	return nil
}

// TODO: make it safe
func (odi *OnDiskIndex) unloadTerm(term string) {
	odi.index.postingLists[term] = nil
}

func (odi *OnDiskIndex) GetPostingListsForTerms(terms []string) map[string]PostingList {
	postingListsForTerms := make(map[string]PostingList)
	err := odi.LoadTerms(terms)
	if err != nil {
		panic(err)
	}
	for _, term := range terms {
		postingListsForTerms[term] = odi.index.postingLists[term]
	}
	return postingListsForTerms
}

func (odi *OnDiskIndex) LoadTerms(terms []string) error {
	for _, term := range terms {
		err := odi.loadTerm(term)
		if err != nil {
			fmt.Printf("%s is not in the index, it won't be taken into account\n", term)
		}
	}
	return nil
}