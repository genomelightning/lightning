// Package cytomap handles UCSC cytomap file.
package cytomap

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/genomelightning/lightning/tileset"
)

// CytoRule represents the rule of positing human genome.
type CytoRule struct {
	Chr        string
	Start, End int64 // Index, start from 0.
	Section    string
	Color      string
	Tiles      []*tileset.Tile
}

// CytoMap represents cytomap of human genome.
type CytoMap struct {
	Hg    int
	Rules []*CytoRule
}

// ParseCytoMap parses UCSC cytomap file.
func ParseCytoMap(hgNum int, fileName string) (*CytoMap, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	snr := bufio.NewScanner(f)

	// Jump over the first line.
	snr.Scan()

	cm := &CytoMap{Hg: hgNum}
	cm.Rules = make([]*CytoRule, 0, 1000)

	// Parsing lines.
	for snr.Scan() {
		infos := strings.Split(snr.Text(), "\t")
		if len(infos) != 5 || len(infos[3]) == 0 {
			break
		}

		rule := &CytoRule{
			Chr:     infos[0],
			Section: infos[3],
			Color:   infos[4],
		}
		startIdx, err := strconv.ParseInt(infos[1], 10, 64)
		if err != nil {
			return nil, nil
		}
		endIdx, err := strconv.ParseInt(infos[2], 10, 64)
		if err != nil {
			return nil, nil
		}
		rule.Start = startIdx
		rule.End = endIdx
		cm.Rules = append(cm.Rules, rule)
	}
	return cm, nil
}

// getRule finds the rule and set data that fits the range of tile.
// It returns false when no rule found.
func (cm *CytoMap) checkRule(chr string, start, end int64, data []byte) bool {
	for _, rule := range cm.Rules {
		if rule.Chr != chr {
			continue
		}

		if start >= rule.Start && end <= rule.End {
			rule.Tiles = append(rule.Tiles, &tileset.Tile{data})
			return true
		}
	}
	return false
}

func (cm *CytoMap) parseTile(i int) (n int64, err error) {
	f, err := os.Open(fmt.Sprintf("data/tiles/tileset%04d.fa", i))
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var (
		chr        string
		start, end int64
	)
	buf := bytes.NewBufferString("")
	snr := bufio.NewScanner(f)
	for snr.Scan() {
		byts := snr.Bytes()
		if byts[0] != '>' {
			if _, err = buf.Write(byts); err != nil {
				return 0, err
			}
			continue
		}

		// Set for last tile data.
		if len(chr) > 0 {
			if !cm.checkRule(chr, start, end, buf.Bytes()) {
				log.Printf("No rule match(%04d): %s %d-%d\n", i, chr, start, end)
			}
			buf.Reset()
			n++
		}

		// Get information about current tile.
		infos := strings.Split(snr.Text(), ":")
		idxes := strings.Split(infos[1], "-")
		chr = infos[0][1:]
		start, err = strconv.ParseInt(idxes[0], 10, 64)
		if err != nil {
			return 0, nil
		}
		end, err = strconv.ParseInt(idxes[1], 10, 64)
		if err != nil {
			return 0, nil
		}
	}

	// Set for last tile data.
	if len(chr) > 0 {
		if !cm.checkRule(chr, start, end, buf.Bytes()) {
			log.Printf("No rule match(%04d): %s %d-%d\n", i, chr, start, end)
		}
		buf.Reset()
		n++
	}

	return n, nil
}

// PasreTiles parses tile set files by rules.
func (cm *CytoMap) PasreTiles() (n int64, err error) {
	var m int64
	for i := 0; i < 1001; i++ {
		if m, err = cm.parseTile(i); err != nil {
			if strings.Contains(err.Error(), "no such file or directory") {
				return n, nil
			}
			return 0, err
		}
		n += m
	}
	return n, nil
}
