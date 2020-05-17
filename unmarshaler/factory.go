package unmarshaler

type UnmarshalerFactoryIFace interface {
	GetEpisodeMetadataUnmarshalerFactory() UnmarshalerIFace
}

type MetadataFactory struct{}

func (mf *MetadataFactory) GetEpisodeMetadataUnmarshalerFactory() UnmarshalerIFace {
	return &EpisodeMetadataUnmarshaler{}
}

func NewMetadataFactory() UnmarshalerFactoryIFace {
	return &MetadataFactory{}
}
