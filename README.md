# Silent Errors with Missing Map Keys in Go

This repository demonstrates a common, yet subtle, error in Go related to accessing non-existent keys in maps.  Go's map implementation doesn't throw an error when you try to access a key that's not present. Instead, it returns the zero value for the map's value type. This can be problematic because it leads to silent failures; your program continues running without alerting you to the missing key.

## The Bug
The `bug.go` file showcases this behavior.  Accessing keys that aren't in the map simply returns 0, which is the zero value for the `int` type. This may not always be readily apparent, especially in larger, more complex programs.

## The Solution
The `bugSolution.go` file demonstrates a better approach. It uses the comma ok idiom to explicitly check if a key exists before accessing its value.  This idiom eliminates the silent failure and makes code more robust.

## How to Run
1. Clone the repository.
2. Navigate to the project directory.
3. Run the files using `go run bug.go` and `go run bugSolution.go` to see the different behavior.