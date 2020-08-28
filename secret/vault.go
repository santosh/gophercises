package secret

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"github.com/santosh/gophercises/secret/cipher"
)

// File takes in an encodingKey and filepath. This encryption key
// is then used in subsequent calls via Set and Get methods.
func File(encodingKey, filepath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filepath:    filepath,
	}
}

// Vault represent an entity which stores key value pair.
type Vault struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

// load decrypt the encrypted file and fetch the data.
func (v *Vault) load() error {
	f, err := os.Open(v.filepath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}
	defer f.Close()
	r, err := cipher.DecryptReader(v.encodingKey, f)
	if err != nil {
		return err
	}
	return v.readKeyValues(r)
}

func (v *Vault) readKeyValues(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&v.keyValues)
}

// save encrypts the key-value pair with given encodingKey
func (v *Vault) save() error {
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	w, err := cipher.EncryptWriter(v.encodingKey, f)
	if err != nil {
		return err
	}
	return v.writeKeyValues(w)
}

func (v *Vault) writeKeyValues(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&v.keyValues)
}

// Get is a threadsafe method to Vault to get a key from the Vault
func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return "", err
	}
	value, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("secret: no value for that key")
	}
	return value, nil
}

// Set is a threadsafe method to Vault to set a key to the
// value in the Vault
func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return err
	}
	v.keyValues[key] = value
	err = v.save()
	return err
}
