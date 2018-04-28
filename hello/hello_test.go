package hello

import (
	"github.com/golang/mock/gomock"
	"github.com/golang-demos/go-mock-demo/mock/mock_hello"
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestCompany_Meeting(t *testing.T) {
	ctl := gomock.NewController(t)
	mock_talker := mock_hello.NewMockTalker(ctl)
	mock_talker.EXPECT().SayHello(gomock.Eq("王尼玛")).Return("欢迎你")

	company := NewCompany(mock_talker)
	assert.Equal(t, company.Meeting("王尼玛"), "欢迎你")
}
