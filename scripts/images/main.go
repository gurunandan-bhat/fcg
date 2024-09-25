package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	toml "github.com/pelletier/go-toml/v2"
)

// bio = "Mihir Bhanage is a film critic and has been reviewing films, majorly Marathi, for Times of India since 2014. Besides reviews, he is also an entertainment correspondent for Pune Times. As a viewer, he loves to watch films across genres and languages"
// tag = "mihirbhanage"
// designations = [ "The Times of India" ]
// name = "Mihir Bhanage"
// img = "/images/members/Mihir-Bhanage.png"

//   [members.soc_media]
//   facebook = "https://www.facebook.com/mihirbhanage"
//   instagram = "https://instagram.com/mihir_03?igshid=YmMyMTA2M2Y="
//   twitter = "https://twitter.com/mihir_bhanage"

type SocMed map[string]string
type MemberIn struct {
	Name         string   `toml:"name,omitempty"`
	Bio          string   `toml:"bio,omitempty"`
	Designations []string `toml:"designations,omitempty"`
	Img          string   `toml:"img,omitempty"`
	Social       SocMed   `toml:"soc_media,omitempty"`
	Tag          string   `toml:"tag,omitempty"`
}

type MemberOut struct {
	Name         string   `toml:"title,omitempty"`
	Bio          string   `toml:"-"`
	Designations []string `toml:"designations,omitempty"`
	Img          string   `toml:"img,omitempty"`
	Social       SocMed   `toml:"soc_media,omitempty"`
	Tag          string   `toml:"tag,omitempty"`
}

type Header struct {
	HdrDate time.Time `toml:"date"`
	Draft   bool      `toml:"draft"`
	Sort    int       `toml:"sort"`
	MemberOut
}

func main() {

	fName := "/Users/nandan/repos/fcg/data/members.toml"
	tomlBytes, err := os.ReadFile(fName)
	if err != nil {
		log.Fatalf("error reading file %s: %s", fName, err)
	}

	var m map[string][]MemberIn

	if err := toml.Unmarshal(tomlBytes, &m); err != nil {
		var derr *toml.DecodeError
		if errors.As(err, &derr) {
			fmt.Println(derr.String())
			row, col := derr.Position()
			fmt.Println("error occurred at row", row, "column", col)
		}
		log.Fatalf("error unmarshaling toml string: %s", err)
	}

	for i, member := range m["members"] {
		mOut := MemberOut(member)
		hdr := Header{
			MemberOut: mOut,
			HdrDate:   time.Now(),
			Draft:     false,
			Sort:      i + 1,
		}
		tomlStr, err := toml.Marshal(&hdr)
		if err != nil {
			log.Fatalf("error marshaling %#v: %s", member, err)
		}

		str := fmt.Sprintf("+++\n%s+++\n\n%s\n", string(tomlStr), member.Bio)
		fPath := "/Users/nandan/repos/fcg/content/members/" + member.Tag + `.md`
		f, err := os.Create(fPath)
		if err != nil {
			log.Fatalf("error creating file %s: %s", fPath, err)
		}
		f.Write([]byte(str))
		f.Close()
	}
	// 	fmt.Printf("%+v\n", m)

}
