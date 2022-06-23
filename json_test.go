package json

import (
	"ej/user"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/bytedance/sonic"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
)

func BenchmarkMarshalStd(b *testing.B) {
	jsonByte, _ := ioutil.ReadFile("/home/liuli1/devspace/go_test/json/data.json")
	ss := user.Users{}
	for n := 0; n < b.N; n++ {
		json.Unmarshal(jsonByte, &ss)
		json.Marshal(ss)
	}
}

func BenchmarkMarshalEasyjson(b *testing.B) {
	jsonByte, _ := ioutil.ReadFile("/home/liuli1/devspace/go_test/json/data.json")
	ss := user.Users{}
	for n := 0; n < b.N; n++ {
		json.Unmarshal(jsonByte, &ss)
		easyjson.Marshal(ss)
	}
}

func BenchmarkMarshalSonic(b *testing.B) {
	jsonByte, _ := ioutil.ReadFile("/home/liuli1/devspace/go_test/json/data.json")
	ss := user.Users{}
	for n := 0; n < b.N; n++ {
		json.Unmarshal(jsonByte, &ss)
		sonic.Marshal(ss)
	}
}

func BenchmarkMarshalJsoniter(b *testing.B) {
	jsonByte, _ := ioutil.ReadFile("/home/liuli1/devspace/go_test/json/data.json")
	ss := user.Users{}
	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
	for n := 0; n < b.N; n++ {
		json.Unmarshal(jsonByte, &ss)
		jsonIterator.Marshal(&ss)
	}
}

func BenchmarkUnmarshalStd(b *testing.B) {
	jsonByte, _ := ioutil.ReadFile("/home/liuli1/devspace/go_test/json/data.json")
	ss := user.Users{}
	for n := 0; n < b.N; n++ {
		json.Unmarshal(jsonByte, &ss)
	}
}

func BenchmarkUnmarshalEasyjson(b *testing.B) {
	jsonByte, _ := ioutil.ReadFile("/home/liuli1/devspace/go_test/json/data.json")
	ss := user.Users{}
	for n := 0; n < b.N; n++ {
		easyjson.Unmarshal(jsonByte, &ss)
	}
}

func BenchmarkUnmarshalSonic(b *testing.B) {
	jsonByte, _ := ioutil.ReadFile("/home/liuli1/devspace/go_test/json/data.json")
	ss := user.Users{}
	for n := 0; n < b.N; n++ {
		sonic.Unmarshal(jsonByte, &ss)
	}
}

func BenchmarkUnmarshalJsoniter(b *testing.B) {
	jsonByte, _ := ioutil.ReadFile("/home/liuli1/devspace/go_test/json/data.json")
	ss := user.Users{}
	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
	for n := 0; n < b.N; n++ {
		jsonIterator.Unmarshal(jsonByte, &ss)
	}
}