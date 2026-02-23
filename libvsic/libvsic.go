package vsic

import (
	"bufio"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"net"
	"strings"
	"time"
)

type Config struct {
	MaxMsgSize int
	TimeoutSec int
}

type Conn struct {
	NetConn net.Conn
	R       *bufio.Reader
	W       *bufio.Writer
	Nick    string
	cfg     Config
}

func Wrap(c net.Conn, cfg Config) *Conn {
	if cfg.MaxMsgSize <= 0 {
		cfg.MaxMsgSize = 4096
	}
	if cfg.TimeoutSec <= 0 {
		cfg.TimeoutSec = 120
	}

	return &Conn{
		NetConn: c,
		R:       bufio.NewReaderSize(c, cfg.MaxMsgSize),
		W:       bufio.NewWriter(c),
		cfg:     cfg,
	}
}

func (c *Conn) Close() error {
	return c.NetConn.Close()
}

func (c *Conn) ReadLine() (string, error) {
	_ = c.NetConn.SetReadDeadline(time.Now().Add(time.Duration(c.cfg.TimeoutSec) * time.Second))

	line, err := c.R.ReadString('\n')
	if err != nil {
		return "", err
	}

	if len(line) > c.cfg.MaxMsgSize {
		return "", errors.New("message too big")
	}

	line = strings.TrimRight(line, "\r\n")

	// stop multi line injection type stuff
	if strings.Contains(line, "\n") || strings.Contains(line, "\r") {
		return "", errors.New("invalid control chars")
	}

	return line, nil
}

func (c *Conn) WriteLine(s string) error {
	_ = c.NetConn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	if len(s) > c.cfg.MaxMsgSize {
		return errors.New("message too big")
	}

	if strings.ContainsAny(s, "\n\r") {
		return errors.New("invalid control chars")
	}

	if _, err := c.W.WriteString(s + "\n"); err != nil {
		return err
	}

	return c.W.Flush()
}

func ParseCommand(line string) (cmd string, arg string) {
	if i := strings.IndexByte(line, ' '); i != -1 {
		return line[:i], strings.TrimSpace(line[i+1:])
	}
	return line, ""
}

func ValidNick(n string) bool {
	if len(n) < 3 || len(n) > 20 {
		return false
	}
	for _, r := range n {
		if !(r >= 'a' && r <= 'z' ||
			r >= 'A' && r <= 'Z' ||
			r >= '0' && r <= '9' ||
			r == '_') {
			return false
		}
	}
	return true
}

func RandomSuffix() string {
	var b [2]byte
	_, _ = rand.Read(b[:])
	n := binary.BigEndian.Uint16(b[:]) % 10000
	return "_" + pad4(int(n))
}

func pad4(n int) string {
	if n < 10 {
		return "000" + itoa(n)
	}
	if n < 100 {
		return "00" + itoa(n)
	}
	if n < 1000 {
		return "0" + itoa(n)
	}
	return itoa(n)
}

func itoa(n int) string {
	return strings.TrimPrefix(strings.TrimPrefix(time.Unix(int64(n), 0).UTC().Format("0000"), "1970"), "")
}
