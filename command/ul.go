package command

import (
	"fmt"
	"go.dead.blue/cli115/context"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

/*
"ul" is abbr. for "upload".
*/
type UlCommand struct {
	ArgsCommand
}

func (c *UlCommand) Name() string {
	return "ul"
}

func (c *UlCommand) ImplExec(ctx *context.Impl, args []string) (err error) {
	if len(args) < 0 {
		return errArgsNotEnough
	}
	// Check curl command
	exe, err := exec.LookPath("curl")
	if err != nil {
		return errCurlNotExist
	}
	// Check local file
	localFile := args[0]
	fmt.Printf("Upload local file: %s\n", localFile)
	info, err := os.Stat(localFile)
	if err != nil {
		return
	}
	// Create upload ticket
	ticket, err := ctx.Agent.CreateUploadTicket(ctx.Fs.Curr().Id, info)
	if err != nil {
		return
	}
	// Create temp file to receive result
	tmpFile, err := ioutil.TempFile(os.TempDir(), "115-upload-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	// Upload file via curl
	cmd := exec.Command(exe, ticket.Endpoint, "-o", tmpFile.Name(), "-#")
	for name, value := range ticket.Values {
		cmd.Args = append(cmd.Args, "-F", fmt.Sprintf("%s=%s", name, value))
	}
	cmd.Args = append(cmd.Args, "-F", fmt.Sprintf("%s=@%s", ticket.FileField, localFile))
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err = cmd.Run(); err != nil {
		return
	}
	response, _ := ioutil.ReadAll(tmpFile)
	if file, err := ctx.Agent.ParseUploadResult(response); err != nil {
		return err
	} else {
		log.Printf("Uploaded file: %s", file.Name)
	}
	return
}
