package cte

import (
	"encoding/base64"
	"math"
	"net/url"
	"reflect"

	// "bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

// General

type MetadataMap map[interface{}]interface{}
type Comment string

type Marker string
type Reference string

func asList(values ...interface{}) []interface{} {
	return values
}

func asMap(values ...interface{}) map[interface{}]interface{} {
	result := make(map[interface{}]interface{})
	var key interface{}
	for i, v := range values {
		if i&1 == 0 {
			key = v
		} else {
			result[key] = v
		}
	}
	return result
}

func asMetadataMap(values ...interface{}) MetadataMap {
	result := make(MetadataMap)
	var key interface{}
	for i, v := range values {
		if i&1 == 0 {
			key = v
		} else {
			result[key] = v
		}
	}
	return result
}

func asComment(value string) Comment {
	return Comment(value)
}

func asMarker(value string) Marker {
	return Marker(value)
}

func asIntMarker(value int) Marker {
	return Marker(string(value))
}

func asReference(value string) Reference {
	return Reference(value)
}

func asIntReference(value int) Reference {
	return Reference(string(value))
}

func asURL(str string) *url.URL {
	url, err := url.Parse(str)
	if err != nil {
		url, err = url.Parse("http://parse.error")
	}
	return url
}

func asDate(year int, month int, day int) time.Time {
	location := time.UTC
	hour := 0
	minute := 0
	second := 0
	nanosecond := 0
	return time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, location)
}

func asTime(hour int, minute int, second int, nanosecond int, timezone string) time.Time {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		panic(err)
	}
	baseTime := time.Now()
	return time.Date(baseTime.Year(), baseTime.Month(), baseTime.Day(), hour, minute, second, nanosecond, location)
}

func asTimestamp(year int, month int, day int, hour int, minute int, second int, nanosecond int, timezone string) time.Time {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		panic(err)
	}
	return time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, location)
}

var stringGeneratorChars = [...]byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
	'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
}

func generateString(length int) string {
	var result strings.Builder
	for i := 0; i < length; i++ {
		result.WriteByte(stringGeneratorChars[i%len(stringGeneratorChars)])
	}
	return result.String()
}

func generateBytes(length int) []byte {
	return []byte(generateString(length))
}

func getPanicContents(function func()) (recovered interface{}) {
	defer func() {
		recovered = recover()
	}()
	function()
	return recovered
}

func assertPanics(t *testing.T, function func()) {
	if getPanicContents(function) == nil {
		t.Errorf("Should have panicked but didn't")
	}
}

func assertDoesNotPanic(t *testing.T, function func()) {
	if result := getPanicContents(function); result != nil {
		t.Errorf("Should not have panicked, but did: %v", result)
	}
}

func assertSuccess(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func assertFailure(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Unexpected success")
	}
}

// Decoder

type Nil string

var NilValue Nil = Nil("*NIL*")

type arrayType int

const (
	arrayTypeNone arrayType = iota
	arrayTypeBytes
	arrayTypeString
	arrayTypeURI
	arrayTypeComment
)

type testCallbacks struct {
	nextValue          interface{}
	containerStack     []interface{}
	currentList        []interface{}
	currentMap         map[interface{}]interface{}
	currentMetadataMap MetadataMap
	currentArray       []byte
	currentArrayType   arrayType
}

func (this *testCallbacks) storeValue(value interface{}) {
	if this.currentList != nil {
		this.currentList = append(this.currentList, value)
		return
	}

	if this.currentMap != nil {
		if this.nextValue == nil {
			this.nextValue = value
		} else {
			this.currentMap[this.nextValue] = value
			this.nextValue = nil
		}
		return
	}

	if this.currentMetadataMap != nil {
		if this.nextValue == nil {
			this.nextValue = value
		} else {
			this.currentMetadataMap[this.nextValue] = value
			this.nextValue = nil
		}
		return
	}

	if this.nextValue != nil {
		panic(fmt.Errorf("Top level object already exists: %v", this.nextValue))
	}
	this.nextValue = value
}

func (this *testCallbacks) setCurrentContainer() {
	lastEntry := len(this.containerStack) - 1
	this.currentList = nil
	this.currentMap = nil
	if lastEntry >= 0 {
		container := this.containerStack[lastEntry]
		switch container.(type) {
		case []interface{}:
			this.currentList = container.([]interface{})
		case *[]interface{}:
			this.currentList = *(container.(*[]interface{}))
		case map[interface{}]interface{}:
			this.currentMap = container.(map[interface{}]interface{})
		case *map[interface{}]interface{}:
			this.currentMap = *(container.(*map[interface{}]interface{}))
		case MetadataMap:
			this.currentMetadataMap = container.(MetadataMap)
		default:
			panic(fmt.Errorf("Unknown container type: %v", container))
		}
	}
}

func (this *testCallbacks) containerEnd() {
	var item interface{}

	if this.currentList != nil {
		item = this.currentList
		this.currentList = nil
	} else if this.currentMap != nil {
		item = this.currentMap
		this.currentMap = nil
	} else if this.currentMetadataMap != nil {
		item = this.currentMetadataMap
		this.currentMap = nil
	} else {
		panic("Ended unhandled container type")
	}
	length := len(this.containerStack)
	if length > 0 {
		this.containerStack = this.containerStack[:length-1]
		this.setCurrentContainer()
	}
	this.storeValue(item)
}

func (this *testCallbacks) containerBegin(container interface{}) {
	this.containerStack = append(this.containerStack, container)
	this.setCurrentContainer()
}

func (this *testCallbacks) listBegin() {
	this.containerBegin(make([]interface{}, 0))
}

func (this *testCallbacks) mapBegin() {
	this.containerBegin(make(map[interface{}]interface{}))
}

func (this *testCallbacks) metadataMapBegin() {
	this.containerBegin(make(MetadataMap))
}

func (this *testCallbacks) arrayBegin(newArrayType arrayType) {
	this.currentArray = make([]byte, 0, 10)
	this.currentArrayType = newArrayType
}

func (this *testCallbacks) arrayData(data []byte) {
	this.currentArray = append(this.currentArray, data...)
}

func (this *testCallbacks) arrayEnd() error {
	array := this.currentArray
	this.currentArray = nil
	if this.currentArrayType == arrayTypeBytes {
		this.storeValue(array)
	} else if this.currentArrayType == arrayTypeURI {
		uri, err := url.Parse(string(array))
		if err != nil {
			return err
		}
		this.storeValue(uri)
	} else {
		this.storeValue(string(array))
	}
	return nil
}

func (this *testCallbacks) getValue() interface{} {
	if len(this.containerStack) != 0 {
		return this.containerStack[0]
	}
	return this.nextValue
}

func (this *testCallbacks) OnNil() error {
	this.storeValue(NilValue)
	return nil
}

func (this *testCallbacks) OnBool(value bool) error {
	this.storeValue(value)
	return nil
}

func (this *testCallbacks) OnPositiveInt(value uint64) error {
	this.storeValue(value)
	return nil
}

func (this *testCallbacks) OnNegativeInt(value uint64) error {
	this.storeValue(-int64(value))
	return nil
}

func (this *testCallbacks) OnFloat(value float64) error {
	this.storeValue(value)
	return nil
}

func (this *testCallbacks) OnDecimalFloat(significand int64, exponent int) error {
	this.storeValue(float64(significand) * math.Pow10(exponent))
	return nil
}

func (this *testCallbacks) OnDate(year, month, day int) error {
	this.storeValue(asDate(year, month, day))
	return nil
}

func (this *testCallbacks) OnTimeTZ(hour, minute, second, nanosecond int, tz string) error {
	this.storeValue(asTime(hour, minute, second, nanosecond, tz))
	return nil
}

func (this *testCallbacks) OnTimeLoc(hour, minute, second, nanosecond int, latitude, longitude float32) error {
	// TODO: Something?
	baseDate := time.Now()
	loc := time.UTC
	this.storeValue(time.Date(baseDate.Year(), baseDate.Month(), baseDate.Day(), hour, minute, second, nanosecond, loc))
	return nil
}

func (this *testCallbacks) OnTimestampTZ(year, month, day, hour, minute, second, nanosecond int, tz string) error {
	this.storeValue(asTimestamp(year, month, day, hour, minute, second, nanosecond, tz))
	return nil
}

func (this *testCallbacks) OnTimestampLoc(year, month, day, hour, minute, second, nanosecond int, latitude, longitude float32) error {
	// TODO: Something?
	loc := time.UTC
	this.storeValue(time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, loc))
	return nil
}

func (this *testCallbacks) OnListBegin() error {
	this.listBegin()
	return nil
}

func (this *testCallbacks) OnMapBegin() error {
	this.mapBegin()
	return nil
}

func (this *testCallbacks) OnMetadataMapBegin() error {
	this.metadataMapBegin()
	return nil
}

func (this *testCallbacks) OnContainerEnd() error {
	this.containerEnd()
	return nil
}

func (this *testCallbacks) OnStringBegin() error {
	this.arrayBegin(arrayTypeString)
	return nil
}

func (this *testCallbacks) OnCommentBegin() error {
	this.arrayBegin(arrayTypeComment)
	return nil
}

func (this *testCallbacks) OnURIBegin() error {
	this.arrayBegin(arrayTypeURI)
	return nil
}

func (this *testCallbacks) OnBytesBegin() error {
	this.arrayBegin(arrayTypeBytes)
	return nil
}

func (this *testCallbacks) OnArrayData(bytes []byte) error {
	this.arrayData(bytes)
	return nil
}

func (this *testCallbacks) OnArrayEnd() error {
	return this.arrayEnd()
}

func (this *testCallbacks) OnMarker(id string) error {
	this.storeValue(asMarker(id))
	return nil
}

func (this *testCallbacks) OnReference(id string) error {
	this.storeValue(asReference(id))
	return nil
}

func decodeDocument(maxDepth int, encoded []byte) (result interface{}, err error) {
	callbacks := new(testCallbacks)
	parser := NewParser(maxDepth)
	var done bool
	done, err = parser.Parse(encoded, callbacks)
	if err != nil {
		return nil, err
	}
	if !done {
		return nil, fmt.Errorf("Unexpected: Incomplete parse")
	}
	result = callbacks.getValue()
	return result, err
}

// func decodeWithBufferSize(maxDepth int, encoded []byte, bufferSize int) (result interface{}, err error) {
// 	unmarshaler := new(Unmarshaler)
// 	decoder := NewCbeDecoder(ContainerTypeNone, maxDepth, unmarshaler)
// 	for offset := 0; offset < len(encoded); offset += bufferSize {
// 		end := offset + bufferSize
// 		if end > len(encoded) {
// 			end = len(encoded)
// 		}
// 		if err := decoder.Feed(encoded[offset:end]); err != nil {
// 			return nil, err
// 		}
// 	}
// 	if err := decoder.End(); err != nil {
// 		return nil, err
// 	}
// 	result = unmarshaler.Unmarshaled()
// 	return result, err
// }

func tryDecode(maxDepth int, encoded []byte) error {
	_, err := decodeDocument(maxDepth, encoded)
	return err
}

func assertDecoded(t *testing.T, encoded string, expected interface{}) {
	actual, err := decodeDocument(100, []byte(encoded))
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}
	if !DeepEquivalence(actual, expected) {
		t.Errorf("Expected type <%v>, value <%v>, actual type <%v>, value <%v>", reflect.TypeOf(expected), expected, reflect.TypeOf(actual), actual)
	}
}

func assertBase64Decoded(t *testing.T, data []byte) {
	asBase64 := base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
	encoded := fmt.Sprintf("b\"%v\"", asBase64)
	assertDecoded(t, encoded, data)
}

func assertDecodeFails(t *testing.T, encoded string) {
	_, err := decodeDocument(100, []byte(encoded))
	if err == nil {
		t.Errorf("Error: Expected error")
	}
}

// func assertDecodedPiecemeal(t *testing.T, encoded []byte, minBufferSize int, maxBufferSize int, expected interface{}) {
// 	for i := minBufferSize; i < maxBufferSize; i++ {
// 		actual, err := decodeWithBufferSize(100, encoded, i)
// 		if err != nil {
// 			t.Errorf("Error: %v", err)
// 			return
// 		}
// 		if !DeepEquivalence(actual, expected) {
// 			t.Errorf("Expected [%v], actual [%v]", expected, actual)
// 		}
// 	}
// }

// // Encoder

// func assertEncoded(t *testing.T, containerType ContainerType, function func(*CbeEncoder) error, expected []byte) {
// 	encoder := NewCbeEncoder(containerType, nil, 100)
// 	err := function(encoder)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	actual := encoder.EncodedBytes()
// 	if !bytes.Equal(actual, expected) {
// 		t.Errorf("Expected %v, actual %v", expected, actual)
// 	}
// }

// // Marshal / Unmarshal

// func assertMarshaled(t *testing.T, containerType ContainerType, value interface{}, expected []byte) {
// 	encoder := NewCbeEncoder(containerType, nil, 100)
// 	Marshal(encoder, containerType, value)
// 	actual := encoder.EncodedBytes()
// 	if !bytes.Equal(actual, expected) {
// 		t.Errorf("Expected %v, actual %v", expected, actual)
// 	}
// }

// func assertEncodesToExternalBuffer(t *testing.T, containerType ContainerType, value interface{}, bufferSize int) {
// 	buffer := make([]byte, bufferSize)
// 	encoder := NewCbeEncoder(containerType, buffer, 100)
// 	if err := Marshal(encoder, containerType, value); err != nil {
// 		t.Errorf("Unexpected error while marshling: %v", err)
// 		return
// 	}

// 	encoder2 := NewCbeEncoder(containerType, nil, 100)
// 	Marshal(encoder2, containerType, value)
// 	expected := encoder2.EncodedBytes()
// 	if !bytes.Equal(buffer, expected) {
// 		t.Errorf("Expected %v, actual %v", expected, buffer)
// 	}
// }

// func assertFailsEncodingToExternalBuffer(t *testing.T, containerType ContainerType, value interface{}, bufferSize int) {
// 	buffer := make([]byte, bufferSize)
// 	encoder := NewCbeEncoder(containerType, buffer, 100)
// 	assertPanics(t, func() {
// 		Marshal(encoder, containerType, value)
// 	})
// }

// func assertMarshaledSize(t *testing.T, containerType ContainerType, value interface{}, expectedSize int) {
// 	actualSize := EncodedSize(containerType, value)
// 	if actualSize != expectedSize {
// 		t.Errorf("Expected size %v but got %v", expectedSize, actualSize)
// 	}
// }

// func assertMarshalUnmarshal(t *testing.T, containerType ContainerType, expected interface{}) {
// 	assertMarshalUnmarshalProduces(t, containerType, expected, expected)
// }

// func assertMarshalUnmarshalProduces(t *testing.T, containerType ContainerType, input interface{}, expected interface{}) {
// 	encoder := NewCbeEncoder(containerType, nil, 100)
// 	if err := Marshal(encoder, containerType, input); err != nil {
// 		t.Errorf("Unexpected error while marshling: %v", err)
// 		return
// 	}
// 	document := encoder.EncodedBytes()
// 	unmarshaler := new(Unmarshaler)
// 	decoder := NewCbeDecoder(containerType, 100, unmarshaler)
// 	if err := decoder.Decode(document); err != nil {
// 		t.Errorf("Unexpected error while decoding: %v", err)
// 		return
// 	}
// 	actual := unmarshaler.Unmarshaled()

// 	if !DeepEquivalence(actual, expected) {
// 		t.Errorf("Expected %t: <%v>, actual %t: <%v>", expected, expected, actual, actual)
// 	}
// }

func ShortCircuit(errors ...error) error {
	for _, err := range errors {
		if err != nil {
			return err
		}
	}
	return nil
}
