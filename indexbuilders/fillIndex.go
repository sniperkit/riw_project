package indexbuilders

import (
	"log"
	"sync"

	"github.com/rmulton/riw_project/indexes"
	"github.com/rmulton/riw_project/indexes/buildingIndexes"
)

// FillIndex fills in an index using parsed documents sent on the reading channel
func FillIndex(index buildingIndexes.BuildingIndex, readingChannel indexes.ReadingChannel, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for doc := range readingChannel {
		// Add the path to the doc to the map
		index.AddDocToIndex(doc.ID, doc.Path)

		for _, term := range doc.NormalizedTokens {
			//	Add this document to the posting list of this term
			index.AddDocToTerm(doc.ID, term)
		}
	}
	log.Printf("Done getting %d documents", index.GetDocCounter())
}
