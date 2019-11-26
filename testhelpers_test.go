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

// ============================================================================
// Types
// ============================================================================

type NilType string
type MetaType map[interface{}]interface{}
type CommentType string
type ReferenceType interface{}
type SNanType float64

type MarkerType struct {
	id        interface{}
	reference interface{}
}

func Nil() NilType {
	return NilType("*NIL*")
}

func Nan() float64 {
	return math.NaN()
}

func SNan() SNanType {
	return SNanType(math.NaN())
}

func Inf() float64 {
	return math.Inf(1)
}

func NInf() float64 {
	return math.Inf(-1)
}

func List(values ...interface{}) []interface{} {
	return values
}

func Map(values ...interface{}) map[interface{}]interface{} {
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

func Meta(values ...interface{}) MetaType {
	result := make(MetaType)
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

func Bytes(values ...byte) []byte {
	return values
}

func Comment(value string) CommentType {
	return CommentType(value)
}

func PartialMarker(id interface{}) (marker *MarkerType) {
	switch id.(type) {
	case int:
	case string:
	default:
		panic(fmt.Errorf("Unknown marker ID type: %v", reflect.TypeOf(id)))
	}

	marker = new(MarkerType)
	marker.id = id

	return marker
}

func Marker(id interface{}, reference interface{}) (marker *MarkerType) {
	marker = PartialMarker(id)
	marker.reference = reference

	return marker
}

func Reference(id interface{}) ReferenceType {
	switch id.(type) {
	case int:
	case string:
	default:
		panic(fmt.Errorf("Unknown reference ID type: %v", reflect.TypeOf(id)))
	}
	return ReferenceType(id)
}

func URI(str string) *url.URL {
	url, err := url.Parse(str)
	if err != nil {
		url, err = url.Parse("http://parse.error")
	}
	return url
}

func Date(year int, month int, day int) time.Time {
	location := time.UTC
	hour := 0
	minute := 0
	second := 0
	nanosecond := 0
	return time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, location)
}

func Time(hour int, minute int, second int, nanosecond int, timezone string) time.Time {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		panic(err)
	}
	baseTime := time.Now()
	return time.Date(baseTime.Year(), baseTime.Month(), baseTime.Day(), hour, minute, second, nanosecond, location)
}

func TS(year int, month int, day int, hour int, minute int, second int, nanosecond int, timezone string) time.Time {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		panic(err)
	}
	return time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, location)
}

// ============================================================================
// General
// ============================================================================

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

func ShortCircuit(errors ...error) error {
	for _, err := range errors {
		if err != nil {
			return err
		}
	}
	return nil
}

// ============================================================================
// Decoder
// ============================================================================

type testCallbacks struct {
	nextValue            interface{}
	containerStack       []interface{}
	currentList          []interface{}
	currentMap           map[interface{}]interface{}
	currentMetadataMap   MetaType
	currentArray         []byte
	currentArrayType     ArrayType
	currentMarker        *MarkerType
	currentContainerType ContainerType
	isInArray            bool
	isInContainer        bool
	isInMarker           bool
}

func (this *testCallbacks) storeValue(value interface{}) error {
	if this.isInMarker {
		this.currentMarker.reference = value
		this.currentMarker = nil
		this.isInMarker = false
		return nil
	}

	if this.isInContainer {
		switch this.currentContainerType {
		case ContainerTypeList:
			this.currentList = append(this.currentList, value)
		case ContainerTypeMap:
			if this.nextValue == nil {
				this.nextValue = value
			} else {
				this.currentMap[this.nextValue] = value
				this.nextValue = nil
			}
		case ContainerTypeMetadataMap:
			if this.nextValue == nil {
				this.nextValue = value
			} else {
				this.currentMetadataMap[this.nextValue] = value
				this.nextValue = nil
			}
		}
		return nil
	}

	if this.nextValue != nil {
		return fmt.Errorf("Top level object already exists: %v", this.nextValue)
	}
	this.nextValue = value
	return nil
}

func (this *testCallbacks) setCurrentContainer() error {
	lastEntry := len(this.containerStack) - 1
	this.currentList = nil
	this.currentMap = nil
	this.currentMetadataMap = nil

	this.isInContainer = lastEntry >= 0
	if this.isInContainer {
		container := this.containerStack[lastEntry]
		switch container.(type) {
		case []interface{}:
			this.currentList = container.([]interface{})
			this.currentContainerType = ContainerTypeList
		case *[]interface{}:
			this.currentList = *(container.(*[]interface{}))
			this.currentContainerType = ContainerTypeList
		case map[interface{}]interface{}:
			this.currentMap = container.(map[interface{}]interface{})
			this.currentContainerType = ContainerTypeMap
		case *map[interface{}]interface{}:
			this.currentMap = *(container.(*map[interface{}]interface{}))
			this.currentContainerType = ContainerTypeMap
		case MetaType:
			this.currentMetadataMap = container.(MetaType)
			this.currentContainerType = ContainerTypeMetadataMap
		default:
			return fmt.Errorf("Unknown container type: %v", container)
		}
	}
	return nil
}

func (this *testCallbacks) containerEnd() error {
	var item interface{}

	if !this.isInContainer {
		return fmt.Errorf("Called containerEnd() while not in a container")
	}

	switch this.currentContainerType {
	case ContainerTypeList:
		item = this.currentList
		this.currentList = nil
	case ContainerTypeMap:
		item = this.currentMap
		this.currentMap = nil
	case ContainerTypeMetadataMap:
		item = this.currentMetadataMap
		this.currentMetadataMap = nil
	default:
		return fmt.Errorf("Unhandled container type: %v", this.currentContainerType)
	}
	length := len(this.containerStack)
	if length > 0 {
		this.containerStack = this.containerStack[:length-1]
		this.setCurrentContainer()
	}
	this.storeValue(item)
	return nil
}

func (this *testCallbacks) containerBegin(containerType ContainerType) error {
	var container interface{}
	switch containerType {
	case ContainerTypeList:
		container = make([]interface{}, 0)
	case ContainerTypeMap:
		container = make(map[interface{}]interface{})
	case ContainerTypeMetadataMap:
		container = make(MetaType)
	default:
		return fmt.Errorf("Unhandled container type: %v", containerType)
	}
	this.containerStack = append(this.containerStack, container)
	return this.setCurrentContainer()
}

func (this *testCallbacks) arrayBegin(newArrayType ArrayType) error {
	this.currentArray = make([]byte, 0, 100)
	this.currentArrayType = newArrayType
	this.isInArray = true
	return nil
}

func (this *testCallbacks) arrayData(data []byte) error {
	this.currentArray = append(this.currentArray, data...)
	return nil
}

func (this *testCallbacks) arrayEnd() error {
	array := this.currentArray
	this.currentArray = nil
	this.isInArray = false
	switch this.currentArrayType {
	case ArrayTypeBinary:
		return this.storeValue(array)
	case ArrayTypeString:
		return this.storeValue(string(array))
	case ArrayTypeURI:
		uri, err := url.Parse(string(array))
		if err != nil {
			return err
		}
		return this.storeValue(uri)
	case ArrayTypeComment:
		return this.storeValue(Comment(string(array)))
	default:
		return fmt.Errorf("Unhandled array type: %v", this.currentArrayType)
	}
}

func (this *testCallbacks) getValue() interface{} {
	if len(this.containerStack) != 0 {
		return this.containerStack[0]
	}
	return this.nextValue
}

func (this *testCallbacks) OnNil() error {
	return this.storeValue(Nil())
}

func (this *testCallbacks) OnBool(value bool) error {
	return this.storeValue(value)
}

func (this *testCallbacks) OnPositiveInt(value uint64) error {
	return this.storeValue(value)
}

func (this *testCallbacks) OnNegativeInt(value uint64) error {
	return this.storeValue(-int64(value))
}

func (this *testCallbacks) OnFloat(value float64) error {
	return this.storeValue(value)
}

func (this *testCallbacks) OnDecimalFloat(significand int64, exponent int) error {
	return this.storeValue(float64(significand) * math.Pow10(exponent))
}

func (this *testCallbacks) OnDate(year, month, day int) error {
	return this.storeValue(Date(year, month, day))
}

func (this *testCallbacks) OnTimeTZ(hour, minute, second, nanosecond int, tz string) error {
	return this.storeValue(Time(hour, minute, second, nanosecond, tz))
}

func (this *testCallbacks) OnTimeLoc(hour, minute, second, nanosecond int, latitude, longitude float32) error {
	// TODO: Something?
	baseDate := time.Now()
	loc := time.UTC
	return this.storeValue(time.Date(baseDate.Year(), baseDate.Month(), baseDate.Day(), hour, minute, second, nanosecond, loc))
}

func (this *testCallbacks) OnTimestampTZ(year, month, day, hour, minute, second, nanosecond int, tz string) error {
	return this.storeValue(TS(year, month, day, hour, minute, second, nanosecond, tz))
}

func (this *testCallbacks) OnTimestampLoc(year, month, day, hour, minute, second, nanosecond int, latitude, longitude float32) error {
	// TODO: Something?
	loc := time.UTC
	return this.storeValue(time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, loc))
}

func (this *testCallbacks) OnContainerBegin(containerType ContainerType) error {
	return this.containerBegin(containerType)
}

func (this *testCallbacks) OnContainerEnd() error {
	return this.containerEnd()
}

func (this *testCallbacks) OnArrayBegin(arrayType ArrayType) error {
	return this.arrayBegin(arrayType)
}

func (this *testCallbacks) OnArrayData(bytes []byte) error {
	return this.arrayData(bytes)
}

func (this *testCallbacks) OnArrayEnd() error {
	return this.arrayEnd()
}

// func (this *testCallbacks) OnMarker(id string) error {
// 	if this.currentMarker != nil {
// 		return fmt.Errorf("Already in a marker")
// 	}
//  this.isInMarker = true
// 	this.currentMarker = PartialMarker(id)
// 	return this.storeValue(this.currentMarker)
// }

// func (this *testCallbacks) OnReference(id string) error {
// 	return this.storeValue(ReferenceType(id))
// }

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

// ============================================================================
// Assertions
// ============================================================================

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
