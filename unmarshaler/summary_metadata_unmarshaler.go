package unmarshaler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type SummaryMetadata struct {
	EpisodeNo int    `json:"episodeNo"`
	Summary   string `json:"summary"`
}

type EpisodeSummaryMetadataUnmarshaler struct{}

func (esm *EpisodeSummaryMetadataUnmarshaler) Unmarshal(data []byte) (interface{}, error) {
	obj := SummaryMetadata{}
	err := json.Unmarshal(data, &obj)

	if err != nil {
		logrus.WithError(err).Error("Failed to unmarshal summary metadata.")
		return nil, err
	}
	return obj, nil
}
