package main

import (
	"log"

	"github.com/cockroachdb/pebble"
)

// InitializeDB initializes and returns a PebbleDB instance.
func InitializeDB(path string) *pebble.DB {
    db, err := pebble.Open(path, &pebble.Options{})
    if err != nil {
        log.Fatalf("Failed to open database: %v", err)
    }
    return db
}

// // StorePerson serializes and stores a Person protobuf message in PebbleDB.
// func StorePerson(db *pebble.DB, key string, person *go_vite_app.Person) error {
//     // Serialize the Person message to bytes
//     data, err := proto.Marshal(person)
//     if err != nil {
//         return err
//     }

//     // Store the serialized data in PebbleDB
//     err = db.Set([]byte(key), data, pebble.Sync)
//     return err
// }

// // RetrievePerson retrieves and deserializes a Person protobuf message from PebbleDB.
// func RetrievePerson(db *pebble.DB, key string) (*go_vite_app.Person, error) {
//     // Retrieve the data from PebbleDB
//     data, closer, err := db.Get([]byte(key))
//     if err != nil {
//         return nil, err
//     }
//     defer closer.Close()

//     // Deserialize the data back into a Person message
//     person := &go_vite_app.Person{}
//     if err := proto.Unmarshal(data, person); err != nil {
//         return nil, err
//     }

//     return person, nil
// }

// Example usage
// func main() {
//     db := InitializeDB("path/to/your/db")
//     defer db.Close()

//     person := &go_vite_app.Person{
//         Name:  "John Doe",
//         Id:    1234,
//         Email: "johndoe@example.com",
//     }

//     key := "person:1234"

//     if err := StorePerson(db, key, person); err != nil {
//         log.Fatalf("Failed to store person: %v", err)
//     }

//     retrievedPerson, err := RetrievePerson(db, key)
//     if err != nil {
//         log.Fatalf("Failed to retrieve person: %v", err)
//     }

//     log.Printf("Retrieved Person: %+v", retrievedPerson)
// }
