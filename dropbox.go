package main

import (
	"fmt"

	"github.com/stacktic/dropbox"
)

func main() {

	var db *dropbox.Dropbox

	var clientid, clientsecret string
	var token string

	clientid = "31rmr26bffk3ij8"
	clientsecret = "n0rlqt27iuf7scp"
	token = "KeymFkX_8yAAAAAAAAACSBcyXS5BbpSCBxa4wf7ejZAdhEyt201sno3he5lImvl4"

	db = dropbox.NewDropbox()

	db.SetAppInfo(clientid, clientsecret)

	db.SetAccessToken(token)
	p, _ := db.GetAccountInfo()
	fmt.Println(p.DisplayName)
	fmt.Println("Root dir is,", db.RootDirectory)
	new1, _ := db.Metadata("/", true, true, "", "", 1000)
	dirs := []string{}
	fmt.Println("contents")
	for n, _ := range new1.Contents {
		fmt.Println(new1.Contents[n].Path)
		dirs = append(dirs, new1.Contents[n].Path)

	}
	fmt.Println(dirs)
	fmt.Println(new1.Path, new1.Root, new1.ParentSharedFolderID, new1.IsDir)
	fmt.Println(new1.Contents)
}
