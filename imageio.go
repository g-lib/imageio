package imageio

/*

Imageio is a Golang library that provides an easy interface to read and
write a wide range of image data, including animated images, volumetric
data, and scientific formats. It is cross-platform.

Main website: https://github.com/g-lib/imageio
*/

// ReadConfig config for reading image
type ReadConfig struct {
}

// ReadImage Reads an image from the specified file. Returns a numpy array, which
// comes with a dict of meta data at its 'meta' attribute.
func ReadImage(uri string, config ...ReadConfig) error {
	return nil
}

// ReadMImage Reads multiple images from the specified file. Returns a list of
// numpy arrays, each with a dict of meta data at its 'meta' attribute.
func ReadMImage() error {
	return nil
}

// ReadVolImage Reads a volume from the specified file. Returns a numpy array, which
// comes with a dict of meta data at its 'meta' attribute.
func ReadVolImage() error {
	return nil
}

// ReadMVolImage Reads multiple volumes from the specified file. Returns a list of
// numpy arrays, each with a dict of meta data at its 'meta' attribute.
func ReadMVolImage() error {
	return nil
}

// WriteConfig config for writing image
type WriteConfig struct {
}

// WriteImage Write an image to the specified file.
func WriteImage(uri string, image interface{}, config ...WriteConfig) error {
	return nil
}

// WriteMImage Reads multiple images from the specified file. Returns a list of
// numpy arrays, each with a dict of meta data at its 'meta' attribute.
func WriteMImage(uri string, image []interface{}, config ...WriteConfig) error {
	return nil
}

// WriteVolImage Write a volume to the specified file.
func WriteVolImage() error {
	return nil
}

// WriteMVolImage  Write multiple volumes to the specified file.
func WriteMVolImage() error {
	return nil
}

// GetReader Returns a :class:`.Reader` object which can be used to read data
// and meta data from the specified file.
func GetReader() error {
	return nil
}

// GetWriter Returns a :class:`.Writer` object which can be used to write data
// and meta data to the specified file.
func GetWriter() error {
	return nil
}

// Aliases

// Read aliases GetReader
var Read = GetReader

// Save aliases GetWriter
var Save = GetWriter

// SaveImage aliases WriteImage
var SaveImage = WriteImage

// SaveMImage aliases WriteMImage
var SaveMImage = WriteMImage

// SaveVolImage aliases WriteVolImage
var SaveVolImage = WriteVolImage

// SaveMVolImage aliases WriteMVolImage
var SaveMVolImage = WriteMVolImage
