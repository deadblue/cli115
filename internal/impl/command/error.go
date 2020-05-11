package command

type Error string

func (e Error) Error() string {
	return string(e)
}

var (
	errArgsNotEnough Error = "Arguments not enough"

	errDirNotExist Error = "No such directory"

	errFileNotExist Error = "No such file"
	errNotFile      Error = "Not a regular file"

	errNoDownloader Error = "No available downloader"
	errMpvNotExist  Error = "Can not find mpv executable"
	errCurlNotExist Error = "Can not find curl executable"
)
