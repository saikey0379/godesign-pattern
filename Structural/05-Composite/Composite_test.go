package Composite

import "testing"

func TestComposite(t *testing.T) {
	root := &Composite{ID: 0, Name: "ROOT"}

	folderA := &Composite{ID: 1, Name: "FoderA"}
	root.Add(folderA)
	folderA.Print()

	fileA := &Composite{ID: 2, Name: "FileA"}
	folderA.Add(fileA)
	fileA.Print()

	root.Print()
	root.GetChild(1).Print()
	root.GetChild(1).GetChild(2).Print()

}
