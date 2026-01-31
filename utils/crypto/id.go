package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const aesKeySize = 32

var encryptionKey []byte

// SetKey sets the key for ID encryption/decryption. Key is derived to 32 bytes via SHA-256.
// Must be called at bootstrap (e.g. from config.App.IDEncryptionKey).
func SetKey(key string) {
	if key == "" {
		encryptionKey = nil
		return
	}
	encryptionKey = deriveKey(key)
}

func deriveKey(key string) []byte {
	b := []byte(key)
	if len(b) == aesKeySize {
		return b
	}
	h := sha256.Sum256(b)
	return h[:]
}

// EncryptUUID encrypts a UUID and returns a URL-safe base64 string.
func EncryptUUID(id uuid.UUID) (string, error) {
	if len(encryptionKey) == 0 {
		return id.String(), nil
	}
	plain := []byte(id.String())
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, plain, nil)
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// DecryptUUID decrypts a string produced by EncryptUUID and returns the UUID.
func DecryptUUID(cipherText string) (uuid.UUID, error) {
	if cipherText == "" {
		return uuid.Nil, errors.New("crypto: empty cipher text")
	}
	if len(encryptionKey) == 0 {
		return uuid.Parse(cipherText)
	}
	data, err := base64.URLEncoding.DecodeString(cipherText)
	if err != nil {
		return uuid.Nil, err
	}
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return uuid.Nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return uuid.Nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return uuid.Nil, errors.New("crypto: cipher text too short")
	}
	nonce, cipherBytes := data[:nonceSize], data[nonceSize:]
	plain, err := gcm.Open(nil, nonce, cipherBytes, nil)
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.Parse(string(plain))
}

// EncryptedUUID is a UUID that unmarshals from an encrypted string in JSON (request body).
// Use in request structs for id and *_id fields so client can send encrypted IDs.
type EncryptedUUID struct {
	uuid.UUID
}

// UnmarshalJSON expects an encrypted string (or plain UUID if encryption is disabled), decrypts and sets the UUID.
func (e *EncryptedUUID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		e.UUID = uuid.Nil
		return nil
	}
	id, err := DecryptUUID(s)
	if err != nil {
		return err
	}
	e.UUID = id
	return nil
}

// MarshalJSON outputs the encrypted string so request DTOs that are re-serialized stay consistent.
func (e EncryptedUUID) MarshalJSON() ([]byte, error) {
	s, err := EncryptUUID(e.UUID)
	if err != nil {
		return nil, err
	}
	return json.Marshal(s)
}

// DecryptParam reads the route param, decrypts it and returns the UUID.
func DecryptParam(c *fiber.Ctx, name string) (uuid.UUID, error) {
	raw := c.Params(name)
	return DecryptUUID(raw)
}

// DecryptQuery reads the query param, decrypts it and returns the UUID.
func DecryptQuery(c *fiber.Ctx, name string) (uuid.UUID, error) {
	raw := c.Query(name)
	return DecryptUUID(raw)
}

// isIDField returns true if the struct field name is "id", "ID", or ends with "_id" (e.g. role_id).
func isIDField(name string) bool {
	n := strings.TrimSpace(name)
	if n == "" {
		return false
	}
	lower := strings.ToLower(n)
	if lower == "id" {
		return true
	}
	return strings.HasSuffix(lower, "_id")
}

// EncryptIDsInResponse recursively walks data and encrypts uuid.UUID values in fields named id or *_id.
// Returns a new structure suitable for JSON response. Safe to call with nil or when key is not set.
func EncryptIDsInResponse(data any) (any, error) {
	if data == nil || len(encryptionKey) == 0 {
		return data, nil
	}
	return encryptIDsInValue(reflect.ValueOf(data))
}

func encryptIDsInValue(v reflect.Value) (any, error) {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil, nil
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		return encryptIDsInStruct(v)
	case reflect.Slice, reflect.Array:
		return encryptIDsInSlice(v)
	case reflect.Map:
		return encryptIDsInMap(v)
	default:
		return v.Interface(), nil
	}
}

func encryptIDsInStruct(v reflect.Value) (any, error) {
	out := make(map[string]any)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		sf := t.Field(i)
		fv := v.Field(i)
		jsonTag := sf.Tag.Get("json")
		name := jsonName(jsonTag, sf.Name)
		if name == "-" {
			continue
		}
		if isIDField(sf.Name) && fv.Type() == reflect.TypeOf(uuid.UUID{}) {
			id := fv.Interface().(uuid.UUID)
			enc, err := EncryptUUID(id)
			if err != nil {
				return nil, err
			}
			out[name] = enc
			continue
		}
		child, err := encryptIDsInValue(fv)
		if err != nil {
			return nil, err
		}
		out[name] = child
	}
	return out, nil
}

func jsonName(tag, fieldName string) string {
	if tag == "" {
		return strings.ToLower(fieldName)
	}
	parts := strings.Split(tag, ",")
	name := strings.TrimSpace(parts[0])
	if name == "" {
		return fieldName
	}
	return name
}

func encryptIDsInSlice(v reflect.Value) (any, error) {
	n := v.Len()
	out := make([]any, n)
	for i := 0; i < n; i++ {
		el, err := encryptIDsInValue(v.Index(i))
		if err != nil {
			return nil, err
		}
		out[i] = el
	}
	return out, nil
}

func encryptIDsInMap(v reflect.Value) (any, error) {
	out := make(map[string]any)
	iter := v.MapRange()
	for iter.Next() {
		k := iter.Key()
		val := iter.Value()
		keyStr, _ := k.Interface().(string)
		child, err := encryptIDsInValue(val)
		if err != nil {
			return nil, err
		}
		out[keyStr] = child
	}
	return out, nil
}
