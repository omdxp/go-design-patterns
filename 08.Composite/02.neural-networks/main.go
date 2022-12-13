package main

type NeuronInterface interface {
	Iter() []*Neuron
}

type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}

type NeuronLayer struct {
	Neurons []Neuron
}

func (l *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)
	for i := range l.Neurons {
		result = append(result, &l.Neurons[i])
	}
	return result
}

func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{make([]Neuron, count)}
}

func Connect(from, to NeuronInterface) {
	for _, source := range from.Iter() {
		for _, dest := range to.Iter() {
			source.ConnectTo(dest)
		}
	}
}

func main() {
	neuron1, neuron2 := &Neuron{}, &Neuron{}
	layer1, layer2 := NewNeuronLayer(3), NewNeuronLayer(4)

	Connect(neuron1, neuron2)
	Connect(neuron1, layer1)
	Connect(layer2, neuron1)
	Connect(layer1, layer2)
}

// Neural networks are a type of composite pattern.
// They are made up of neurons, which are connected to each other.
// Each neuron can be connected to other neurons, or to a neuron layer, which is a collection of neurons.
// A neuron layer can be connected to other neurons or neuron layers.
