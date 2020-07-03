package home

import (
	"encoding/binary"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAuthGL(t *testing.T) {
	GLMode = true
	tval := uint32(1)
	data := make([]byte, 4)
	if archIsLittleEndian() {
		binary.LittleEndian.PutUint32(data, tval)
	} else {
		binary.BigEndian.PutUint32(data, tval)
	}
	assert.Nil(t, ioutil.WriteFile("/tmp/gl_token_"+"test", data, 0644))
	assert.False(t, glCheckToken("test"))

	tval = uint32(time.Now().UTC().Unix() + 60)
	data = make([]byte, 4)
	if archIsLittleEndian() {
		binary.LittleEndian.PutUint32(data, tval)
	} else {
		binary.BigEndian.PutUint32(data, tval)
	}
	assert.Nil(t, ioutil.WriteFile("/tmp/gl_token_"+"test", data, 0644))
	r, _ := http.NewRequest("GET", "http://localhost/", nil)
	r.AddCookie(&http.Cookie{Name: glCookieName, Value: "test"})
	assert.True(t, glProcessCookie(r))
	GLMode = false
}
