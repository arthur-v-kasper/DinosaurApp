package dinosaur

type Dinosaur struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Era            int64  `json:"era"`
	Classification int64  `json:"classification"`
}

// https://en.wikipedia.org/wiki/Mesozoic
type Mesozoic struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

const (
	Triassic   = 1
	Jurassic   = 2
	Cretaceous = 3
)

// INSERT INTO dinosaur (id, name, era, classification) values (1, "T-Rex", 2, 1);

//https://www.kids-dinosaurs.com/different-types-of-dinosaurs.html
type DinosaurClassification struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

const (
	Theropods = iota + 1
	Sauropods
	Cerapods
	Thyreophora
	Ankylosauria
	Ornithopod
)

var MesozoicEras = []Mesozoic{
	{ID: Triassic, Name: "Triassic"},
	{ID: Jurassic, Name: "Jurassic"},
	{ID: Cretaceous, Name: "Cretaceous"},
}

var DinosaurClassificationMap = []DinosaurClassification{
	{ID: Theropods, Name: "Theropods"},
	{ID: Sauropods, Name: "Sauropods"},
	{ID: Thyreophora, Name: "Thyreophora"},
	{ID: Ankylosauria, Name: "Ankylosauria"},
	{ID: Ornithopod, Name: "Ornithopod"},
}
