package dinosaur

type Dinosaur struct {
	ID             int64                  `json:"id"`
	Name           string                 `json:"name"`
	Era            Mesozoic               `json:"era"`
	Classification DinosaurClassification `json:"classification"`
}

// https://en.wikipedia.org/wiki/Mesozoic
type Mesozoic int

const (
	Triassic   = 1
	Jurassic   = 2
	Cretaceous = 3
)

// INSERT INTO dinosaur (id, name, era, classification) values (1, "T-Rex", 2, 1);
// func (m Mesozoic) String() string {
// 	switch m {
// 	case Triassic:
// 		return "Triassic"
// 	case Jurassic:
// 		return "Jurassic"
// 	case Cretaceous:
// 		return "Cretaceous"
// 	}
// 	return "unknow"
// }

//https://www.kids-dinosaurs.com/different-types-of-dinosaurs.html
type DinosaurClassification int

const (
	Theropods = iota + 1
	Sauropods
	Cerapods
	Thyreophora
	Ankylosauria
	Ornithopod
)

// func (d DinosaurClassification) String() string {
// 	switch d {
// 	case Theropods:
// 		return "Theropods"
// 	case Sauropods:
// 		return "Sauropods"
// 	case Cerapods:
// 		return "Cerapods"
// 	case Thyreophora:
// 		return "Thyreophora"
// 	case Ankylosauria:
// 		return "Ankylosauria"
// 	case Ornithopod:
// 		return "Ornithopod"
// 	}
// 	return "unknow"
// }

var mesozoic = map[Mesozoic]string{
	Triassic:   "Triassic",
	Jurassic:   "Jurassic",
	Cretaceous: "Cretaceous",
}

var DinosaurClassificationMap = map[DinosaurClassification]string{
	Theropods:    "Theropods",
	Sauropods:    "Sauropods",
	Cerapods:     "Cerapods",
	Thyreophora:  "Thyreophora",
	Ankylosauria: "Ankylosauria",
	Ornithopod:   "Ornithopod",
}
