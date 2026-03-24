package L5

// this is a valid interface, but based on the code alone, we can't tell
// what kind of strings should be passed into the copy function
type Copier interface {
  Copy(string, string) int
}

// you can name the argument types and return types for better readability
type Copier_Named interface {
  Copy(sourceFile string, destinationFile string) (bytesCopied int)
}