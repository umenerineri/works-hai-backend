// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

// Encode implements json.Marshaler.
func (s *ErrResp) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ErrResp) encodeFields(e *jx.Encoder) {
	{
		if s.Error.Set {
			e.FieldStart("error")
			s.Error.Encode(e)
		}
	}
}

var jsonFieldsNameOfErrResp = [1]string{
	0: "error",
}

// Decode decodes ErrResp from json.
func (s *ErrResp) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ErrResp to nil")
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
		return errors.Wrap(err, "decode ErrResp")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ErrResp) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ErrResp) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes PresignedUrlsGetOKResult as json.
func (o OptPresignedUrlsGetOKResult) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes PresignedUrlsGetOKResult from json.
func (o *OptPresignedUrlsGetOKResult) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptPresignedUrlsGetOKResult to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptPresignedUrlsGetOKResult) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptPresignedUrlsGetOKResult) UnmarshalJSON(data []byte) error {
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
func (s *PresignedUrlsGetOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *PresignedUrlsGetOK) encodeFields(e *jx.Encoder) {
	{
		if s.Result.Set {
			e.FieldStart("result")
			s.Result.Encode(e)
		}
	}
}

var jsonFieldsNameOfPresignedUrlsGetOK = [1]string{
	0: "result",
}

// Decode decodes PresignedUrlsGetOK from json.
func (s *PresignedUrlsGetOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PresignedUrlsGetOK to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "result":
			if err := func() error {
				s.Result.Reset()
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
		return errors.Wrap(err, "decode PresignedUrlsGetOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *PresignedUrlsGetOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PresignedUrlsGetOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *PresignedUrlsGetOKResult) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *PresignedUrlsGetOKResult) encodeFields(e *jx.Encoder) {
	{
		if s.HumanDrawing.Set {
			e.FieldStart("humanDrawing")
			s.HumanDrawing.Encode(e)
		}
	}
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

var jsonFieldsNameOfPresignedUrlsGetOKResult = [5]string{
	0: "humanDrawing",
	1: "topDrawing",
	2: "rightDrawing",
	3: "bottomDrawing",
	4: "leftDrawing",
}

// Decode decodes PresignedUrlsGetOKResult from json.
func (s *PresignedUrlsGetOKResult) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PresignedUrlsGetOKResult to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "humanDrawing":
			if err := func() error {
				s.HumanDrawing.Reset()
				if err := s.HumanDrawing.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"humanDrawing\"")
			}
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
		return errors.Wrap(err, "decode PresignedUrlsGetOKResult")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *PresignedUrlsGetOKResult) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PresignedUrlsGetOKResult) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ResourcePathPostBadRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ResourcePathPostBadRequest) encodeFields(e *jx.Encoder) {
	{
		if s.Error.Set {
			e.FieldStart("error")
			s.Error.Encode(e)
		}
	}
}

var jsonFieldsNameOfResourcePathPostBadRequest = [1]string{
	0: "error",
}

// Decode decodes ResourcePathPostBadRequest from json.
func (s *ResourcePathPostBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ResourcePathPostBadRequest to nil")
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
		return errors.Wrap(err, "decode ResourcePathPostBadRequest")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ResourcePathPostBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ResourcePathPostBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ResourcePathPostOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ResourcePathPostOK) encodeFields(e *jx.Encoder) {
	{
		if s.Message.Set {
			e.FieldStart("message")
			s.Message.Encode(e)
		}
	}
}

var jsonFieldsNameOfResourcePathPostOK = [1]string{
	0: "message",
}

// Decode decodes ResourcePathPostOK from json.
func (s *ResourcePathPostOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ResourcePathPostOK to nil")
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
		return errors.Wrap(err, "decode ResourcePathPostOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ResourcePathPostOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ResourcePathPostOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ResourcePathPostReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ResourcePathPostReq) encodeFields(e *jx.Encoder) {
	{
		if s.Image.Set {
			e.FieldStart("image")
			s.Image.Encode(e)
		}
	}
}

var jsonFieldsNameOfResourcePathPostReq = [1]string{
	0: "image",
}

// Decode decodes ResourcePathPostReq from json.
func (s *ResourcePathPostReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ResourcePathPostReq to nil")
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
		return errors.Wrap(err, "decode ResourcePathPostReq")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ResourcePathPostReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ResourcePathPostReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
