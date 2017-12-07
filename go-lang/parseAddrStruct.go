package main

import (
    "fmt"
)

type Dog struct {
    Name string
    Age int
}

type DogCollection struct {
    Data []*Dog
}

func (this *DogCollection) Init() {
    cloey := &Dog{"Cloey", 1}
    ralph := &Dog{"Ralph", 5}
    jackie := &Dog{"Jackie", 10}
    bella := &Dog{"Bella", 2}
    jamie := &Dog{"Jamie", 6}

    this.Data = []*Dog{cloey, ralph, jackie, bella, jamie}
}

func (this *DogCollection) CollectionChannel() chan *Dog {
    dataChannel := make(chan *Dog, len(this.Data))

    for _, dog := range this.Data {
        dataChannel <- dog
    }

    close(dataChannel)

    return dataChannel
}

func main() {
    dc := DogCollection{}
    dc.Init()

    for dog := range dc.CollectionChannel() {
        fmt.Printf("Channel Name: %s\n", dog.Name)
    }
}
