// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *AiDrawingGetOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *AiDrawingGetOK) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("result")
		s.Result.Encode(e)
	}
}

var jsonFieldsNameOfAiDrawingGetOK = [1]string{
	0: "result",
}

// Decode decodes AiDrawingGetOK from json.
func (s *AiDrawingGetOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AiDrawingGetOK to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "result":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.Result.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"result\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode AiDrawingGetOK")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfAiDrawingGetOK) {
					name = jsonFieldsNameOfAiDrawingGetOK[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *AiDrawingGetOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AiDrawingGetOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *AiDrawingGetOKResult) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *AiDrawingGetOKResult) encodeFields(e *jx.Encoder) {
	{
		if s.TopDrawing.Set {
			e.FieldStart("topDrawing")
			s.TopDrawing.Encode(e)
		}
	}
	{
		if s.RightDrawing.Set {
			e.FieldStart("rightDrawing")
			s.RightDrawing.Encode(e)
		}
	}
	{
		if s.BottomDrawing.Set {
			e.FieldStart("bottomDrawing")
			s.BottomDrawing.Encode(e)
		}
	}
	{
		if s.LeftDrawing.Set {
			e.FieldStart("leftDrawing")
			s.LeftDrawing.Encode(e)
		}
	}
}

var jsonFieldsNameOfAiDrawingGetOKResult = [4]string{
	0: "topDrawing",
	1: "rightDrawing",
	2: "bottomDrawing",
	3: "leftDrawing",
}

// Decode decodes AiDrawingGetOKResult from json.
func (s *AiDrawingGetOKResult) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AiDrawingGetOKResult to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "topDrawing":
			if err := func() error {
				s.TopDrawing.Reset()
				if err := s.TopDrawing.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"topDrawing\"")
			}
		case "rightDrawing":
			if err := func() error {
				s.RightDrawing.Reset()
				if err := s.RightDrawing.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"rightDrawing\"")
			}
		case "bottomDrawing":
			if err := func() error {
				s.BottomDrawing.Reset()
				if err := s.BottomDrawing.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"bottomDrawing\"")
			}
		case "leftDrawing":
			if err := func() error {
				s.LeftDrawing.Reset()
				if err := s.LeftDrawing.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"leftDrawing\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode AiDrawingGetOKResult")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *AiDrawingGetOKResult) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AiDrawingGetOKResult) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *HumanDrawingPostBadRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *HumanDrawingPostBadRequest) encodeFields(e *jx.Encoder) {
	{
		if s.Error.Set {
			e.FieldStart("error")
			s.Error.Encode(e)
		}
	}
}

var jsonFieldsNameOfHumanDrawingPostBadRequest = [1]string{
	0: "error",
}

// Decode decodes HumanDrawingPostBadRequest from json.
func (s *HumanDrawingPostBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode HumanDrawingPostBadRequest to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "error":
			if err := func() error {
				s.Error.Reset()
				if err := s.Error.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode HumanDrawingPostBadRequest")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *HumanDrawingPostBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *HumanDrawingPostBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *HumanDrawingPostOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *HumanDrawingPostOK) encodeFields(e *jx.Encoder) {
	{
		if s.Message.Set {
			e.FieldStart("message")
			s.Message.Encode(e)
		}
	}
}

var jsonFieldsNameOfHumanDrawingPostOK = [1]string{
	0: "message",
}

// Decode decodes HumanDrawingPostOK from json.
func (s *HumanDrawingPostOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode HumanDrawingPostOK to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "message":
			if err := func() error {
				s.Message.Reset()
				if err := s.Message.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode HumanDrawingPostOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *HumanDrawingPostOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *HumanDrawingPostOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SavedURLPostBadRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SavedURLPostBadRequest) encodeFields(e *jx.Encoder) {
	{
		if s.Error.Set {
			e.FieldStart("error")
			s.Error.Encode(e)
		}
	}
}

var jsonFieldsNameOfSavedURLPostBadRequest = [1]string{
	0: "error",
}

// Decode decodes SavedURLPostBadRequest from json.
func (s *SavedURLPostBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SavedURLPostBadRequest to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "error":
			if err := func() error {
				s.Error.Reset()
				if err := s.Error.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SavedURLPostBadRequest")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SavedURLPostBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SavedURLPostBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SavedURLPostOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SavedURLPostOK) encodeFields(e *jx.Encoder) {
	{
		if s.Message.Set {
			e.FieldStart("message")
			s.Message.Encode(e)
		}
	}
}

var jsonFieldsNameOfSavedURLPostOK = [1]string{
	0: "message",
}

// Decode decodes SavedURLPostOK from json.
func (s *SavedURLPostOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SavedURLPostOK to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "message":
			if err := func() error {
				s.Message.Reset()
				if err := s.Message.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SavedURLPostOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SavedURLPostOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SavedURLPostOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SavedURLPostReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SavedURLPostReq) encodeFields(e *jx.Encoder) {
	{
		if s.Image.Set {
			e.FieldStart("image")
			s.Image.Encode(e)
		}
	}
}

var jsonFieldsNameOfSavedURLPostReq = [1]string{
	0: "image",
}

// Decode decodes SavedURLPostReq from json.
func (s *SavedURLPostReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SavedURLPostReq to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "image":
			if err := func() error {
				s.Image.Reset()
				if err := s.Image.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"image\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SavedURLPostReq")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SavedURLPostReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SavedURLPostReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UploadURLGetBadRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UploadURLGetBadRequest) encodeFields(e *jx.Encoder) {
	{
		if s.Error.Set {
			e.FieldStart("error")
			s.Error.Encode(e)
		}
	}
}

var jsonFieldsNameOfUploadURLGetBadRequest = [1]string{
	0: "error",
}

// Decode decodes UploadURLGetBadRequest from json.
func (s *UploadURLGetBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UploadURLGetBadRequest to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "error":
			if err := func() error {
				s.Error.Reset()
				if err := s.Error.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UploadURLGetBadRequest")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UploadURLGetBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UploadURLGetBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UploadURLGetOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UploadURLGetOK) encodeFields(e *jx.Encoder) {
	{
		if s.PresignedURL.Set {
			e.FieldStart("presigned_url")
			s.PresignedURL.Encode(e)
		}
	}
}

var jsonFieldsNameOfUploadURLGetOK = [1]string{
	0: "presigned_url",
}

// Decode decodes UploadURLGetOK from json.
func (s *UploadURLGetOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UploadURLGetOK to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "presigned_url":
			if err := func() error {
				s.PresignedURL.Reset()
				if err := s.PresignedURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"presigned_url\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UploadURLGetOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UploadURLGetOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UploadURLGetOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
