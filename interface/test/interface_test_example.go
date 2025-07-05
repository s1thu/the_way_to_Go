package main

type DB interface {
	Get(id int) string
}

// Real implementation
type RealDB struct{}

func (db RealDB) Get(id int) string {
	return "Real Data"
}

// Fake/mock implementation
type FakeDB struct{}

func (db FakeDB) Get(id int) string {
	return "Fake Data"
}

// Function under test
func FetchData(db DB, id int) string {
	return db.Get(id)
}

func main() {
	realDB := RealDB{}
	fakeDB := FakeDB{}

	// Using real implementation
	data := FetchData(realDB, 1)
	println("Real DB:", data) // Output: Real DB: Real Data

	// Using fake implementation
	data = FetchData(fakeDB, 1)
	println("Fake DB:", data) // Output: Fake DB: Fake Data
}
