package domain

type Sample struct {
	ID   string
	Name string
}
type Samples []Sample

type SampleRepository interface {
	GetSamples() (Samples, error)
}
type SamplesUseCase interface {
	GetSamples() (Samples, error)
}
