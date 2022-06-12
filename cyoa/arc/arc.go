package arc

import (
	"encoding/json"
	"fmt"
)

type StoryArc struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Stories map[string]StoryArc

func (s *StoryArc) String() string {
	return fmt.Sprintf("title:%v\n", s.Title)
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func (o *Option) String() string {
	return fmt.Sprintf(" [%v] %v ", o.Arc, o.Text)
}

func Story_Arcs_From_JSON(b []byte) (Stories, error) {
	var arcs Stories
	err := json.Unmarshal(b, &arcs)
	if err != nil {
		return nil, err
	}
	return arcs, nil

}
