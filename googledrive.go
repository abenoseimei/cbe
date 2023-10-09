package main
import (
	"context"
	"google.golang.org/api/option"
	"log"
	"fmt"
	"os"
	drive "google.golang.org/api/drive/v3"
)
func Googledrive(c *Conf, localFile string, op string) {
	ctx := context.Background()
	srv, err := drive.NewService(ctx, option.WithCredentialsFile(c.Googledrive.Credential), option.WithScopes(drive.DriveScope))
	if err != nil {
		log.Fatal(err)
	}
	if op == "list" {
		r, err := srv.Files.List().PageSize(10).Fields("nextPageToken, files(id, name)").Do()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Files:")
		if len(r.Files) == 0 {
			fmt.Println("No files found.")
		} else {
			for _, i := range r.Files {
				fmt.Printf("%s (%s)\n", i.Name, i.Id)
			}
		}
	} else {
		file, err := os.Open(localFile)
		info, _ := file.Stat()
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		parents := []string{c.Googledrive.Parent}
		f := &drive.File{
			Name: info.Name(),
			Parents:  parents,
		}
		_, err = srv.Files.Create(f).Media(file).ProgressUpdater(func(now, size int64) { fmt.Printf("%d, %d\r", now, size) }).Do()
		if err != nil {
			log.Fatal(err)
		}
	}
}