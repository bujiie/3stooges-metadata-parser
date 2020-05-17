package unmarshaler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type EpisodeMetadata struct {
	EpisodeNo int    `json:"episodeNo"`
	Season    string `json:"season"`
	Episode   string `json:"episode"`
	AirDate   string `json:"airDate"`
	Name      string `json:"name"`
}

type EpisodeMetadataUnmarshaler struct{}

func (emu *EpisodeMetadataUnmarshaler) Unmarshal(data []byte) (interface{}, error) {
	obj := make([]*EpisodeMetadata, 0)
	err := json.Unmarshal(data, &obj)

	if err != nil {
		logrus.WithError(err).Error("Failed to unmarshal episode metadata.")
		return nil, err
	}
	return obj, nil
}
