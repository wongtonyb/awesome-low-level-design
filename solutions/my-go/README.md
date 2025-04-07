## notes
>copilot: enable/disable completion

1. type - defines concrete types, representing structure or values

!! Go seperates data(structs) from behavior (methods/inteface) - differs from classes storing both
2. type struct - a type for data structure (represents data)
3. type interface - a type for method signatures that a type must implement (represents behavior)

4. multiple const - const(var1 = val1...) used to declar multiple constants in a single block

Function/Method declaration
5. Functions are standalone block of code to perform a task
6. Methods are functions associated witha specific type (a struct) - defined using a reciever
   1. value reciever - for "read-only" methods. or small struct
   2. pointer reciever (*) - modifies struct. or large struct (inefficiant to be copied)

Return options
7. Pointer (*BaseVehicle / &Vehicle) - does not copy the entire struct. meaning it cna also modify the fields on the original struct
8. Value (BaseVehicle / Vehicle) - copy. ensures caller cannot modify the original struct.
9. Interface (Vehicle / (&)BaseVehicle) - hides implementation data. only shows methods, and hides the data(struct)
10. Multiple value ((Vehicle, error)) - error handling
Go automatically dereferences pointer from return if not needed

nil
11. nil - predefined identifier that represents zero value. essentially "no value" or "not initialized"

12. instance = &ParkingLot{levels: []*Level{}}.
   1. the {} after Level initializes it as an empty []
   2. else it'd be a nil value
13. := - shorthand for declaring and initializing variable. type inferred.