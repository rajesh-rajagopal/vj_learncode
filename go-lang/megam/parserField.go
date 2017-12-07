package main

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

type Df struct {
	RawSpace []byte
	RawInode []byte
}

func (self *Df) Prefix() string {
	return "df"
}

const (
	SPACE = "space"
	INODE = "inode"
	DF    = "df"
)

func init() {
	parser.Add(DF, "true", "Collect disk free space metrics")
}

var dfFlgaMapping = map[string]string{
	SPACE: "-k",
	INODE: "-i",
}

func CollectDf(t string, raw []byte, c *MetricsCollection) (e error) {
	if len(raw) == 0 {
		if flag, ok := dfFlgaMapping[t]; ok {
			raw, e = ReadDf(flag)
		} else {
			return errors.New("no mapping for df key " + t + " defined")
		}
	}
	lines := strings.Split(strings.TrimSpace(string(raw)), "\n")
	for _, line := range lines[1:] {
		if !strings.HasPrefix(line, "/") {
			continue
		}
		chunks := strings.Fields(line)
		if len(chunks) >= 6 {
			tags := map[string]string{
				"file_system": chunks[0],
				"mounted_on":  chunks[5],
			}
			if v, e := strconv.ParseInt(chunks[1], 10, 64); e == nil {
				c.AddWithTags(t+".Total", v, tags)
			}
			if v, e := strconv.ParseInt(chunks[2], 10, 64); e == nil {
				c.AddWithTags(t+".Used", v, tags)
			}
			if v, e := strconv.ParseInt(chunks[3], 10, 64); e == nil {
				c.AddWithTags(t+".Available", v, tags)
			}
			if v, e := strconv.ParseInt(strings.Replace(chunks[4], "%", "", 1), 10, 64); e == nil {
				c.AddWithTags(t+".Use", v, tags)
			}
		}
	}
	return
}

func (self *Df) Collect(c *MetricsCollection) (e error) {
	e = CollectDf(SPACE, self.RawSpace, c)
	if e != nil {
		return
	}
	e = CollectDf(INODE, self.RawInode, c)
	return
}

func ReadDf(flag string) (b []byte, e error) {
	dbg.Print("reading df with", flag)
	b, e = exec.Command("df", flag).Output()
	if e != nil {
		return
	}
	if len(b) == 0 {
		e = errors.New("Reading df returned an empty string")
	}
	return
}


func main() {


		mh := new(MetricHandler)
		col := &Df{
			RawSpace: readFile("fixtures/df.txt"),
			RawInode: readFile("fixtures/df_inode.txt"),
		}
		stats, e := mh.Collect(col)
		So(e, ShouldBeNil)

		So(len(stats), ShouldEqual, 8)

		So(stats[0].Key, ShouldEqual, "df.space.Total")
		So(stats[0].Value, ShouldEqual, int64(20511356))
		So(stats[0].Tags["file_system"], ShouldEqual, "/dev/sda")
		So(stats[0].Tags["mounted_on"], ShouldEqual, "/")

		So(stats[4].Key, ShouldEqual, "df.inode.Total")
		So(stats[4].Value, ShouldEqual, int64(1310720))
		So(stats[4].Tags["file_system"], ShouldEqual, "/dev/sda")
		So(stats[4].Tags["mounted_on"], ShouldEqual, "/")


}
