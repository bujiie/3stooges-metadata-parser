package main

import (
	"encoding/json"
	"fmt"
	"github.com/bujiie/3stooges-metadata-parser/unmashaler"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

var log *logrus.Logger

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	log = logrus.StandardLogger()
}

func main() {
	// Removes the first arg which is the command name.
	filesToParse := os.Args[1:]
	episodeMetas := make([][]*unmarshaler.EpisodeMetadata, 0)

	for _, file := range filesToParse {
		data, err := ioutil.ReadFile(file)

		if err != nil {
			log.WithError(err).Error("Error")
			return
		}

		factory := unmarshaler.NewMetadataFactory()
		episodeMetadata := factory.GetEpisodeMetadataUnmarshalerFactory()
		output, unmarshalErr := episodeMetadata.Unmarshal(data)

		if unmarshalErr != nil {
			log.WithError(unmarshalErr).Error("Failed to unmarshal episode metadata list")
			return
		}
		episodeMetas = append(episodeMetas, output.([]*unmarshaler.EpisodeMetadata))
	}

	for _, metas := range episodeMetas {
		bytes, marshalErr := json.Marshal(metas)

		if marshalErr != nil {
			log.WithError(marshalErr).Error("Failed to marshal episode metadata list")
			return
		}
		fmt.Println(string(bytes))
	}
}
