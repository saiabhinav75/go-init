package filehandler

import (
	"os"
)

/*

os.File contains internal state - things like the current read position, file descriptor (a number the OS uses), and flags
If you copy os.File, you'd have TWO variables trying to manage the SAME file - imagine two people trying to keep track of what page they're on in the same book, but they can't see each other's bookmark!
When you close the file, only one copy would know - the other copy would think the file is still open (disaster!)


*/

type FileReader struct {
	filepath string
	file     (*os.File) // file pointer what happens if this is os.File instead of *os.File
}

// Constructor is a standalone function
func NewFileReader(filePath string) *FileReader {

	return &FileReader{
		filepath: filePath,
		file:     nil, // file:  &os.File{}, this should not be done? because: Never create &os.File{} yourself! USE os.Open()
	}
}

func (f *FileReader) Open() error {
	file, err := os.Open(f.filepath)
	if err != nil {
		// fmt.Println("Error loading file", err.Error()) its better to return the error
		return err
	}
	f.file = file
	return nil
}

func (f *FileReader) Close() error {
	err := f.file.Close()
	if err != nil {
		// fmt.Println("Error Closing File") Return err
		return err
	}
	return nil
}

func (f *FileReader) Read() {
	// f.file.Read(10)
}
