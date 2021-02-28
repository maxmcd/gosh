package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"syscall"
	"text/tabwriter"

	"github.com/spf13/pflag"
)

type lsOptions struct {
	all           bool
	list          bool
	humanReadable bool
	args          []string
	writer        *tabwriter.Writer
}

func newLsOptions(argv []string) (o lsOptions, err error) {
	set := pflag.NewFlagSet("ls", pflag.ContinueOnError)
	set.BoolVarP(&o.all, "all", "a", false, "do not ignore entries starting with .")
	set.BoolVarP(&o.list, "list", "l", false, "use a long list format")
	set.BoolVarP(&o.humanReadable, "human-readable", "h", false, "with -l and -s print sizes like 1k 234M 2G etc.")
	if err = set.Parse(argv); err != nil {
		return
	}
	o.args = set.Args()
	o.writer = tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	return
}

func Ls(argv []string) error {
	o, err := newLsOptions(argv)
	if err != nil {
		return err
	}
	location := "."
	var locations []string
	switch len(o.args) {
	case 0, 1:
	case 2:
		location = o.args[1]
	default:
		locations = o.args[1:]
	}
	if locations != nil {
		for i, location := range locations {
			if err := o.printLocation(location, true); err != nil {
				return nil
			}
			if i != len(locations)-1 {
				fmt.Println()
			}
		}
		return nil
	}
	defer o.writer.Flush()
	return o.printLocation(location, false)
}

func (o lsOptions) printLocation(location string, printDirLabel bool) (err error) {
	fi, err := os.Stat(location)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		o.printFileInfo(fi)
		return nil
	}
	if printDirLabel {
		fmt.Fprintln(o.writer, fi.Name()+"/:")
	}
	files, err := ioutil.ReadDir(location)
	if err != nil {
		return err
	}
	for _, file := range files {
		if !o.all && strings.HasPrefix(file.Name(), ".") {
			continue
		}
		o.printFileInfo(file)
	}
	return nil
}

func (o lsOptions) printFileInfo(fi os.FileInfo) {
	name := fi.Name()
	if fi.IsDir() {
		name += "/"
	}
	if !o.list {
		fmt.Fprintln(o.writer, name)
	} else {
		fmt.Fprintf(o.writer, "%s\t",
			fi.Mode().String(),
		)
		if sys, ok := fi.Sys().(*syscall.Stat_t); ok {
			fmt.Fprintf(o.writer, "%d\t%d\t", sys.Gid, sys.Uid)
		}
		fmt.Fprintf(o.writer, "%d\t%s\t%s\n",
			fi.Size(),
			fi.ModTime().Local().Format("Jan 02 15:04"),
			name,
		)

	}
}
