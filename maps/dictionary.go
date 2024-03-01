package main

type Dictionary map[string]string

const (
	ErrNotFound     = DictionaryErr("could not find the word you're looking for")
	ErrDupKey       = DictionaryErr("duplicate key found")
	ErrUpdateFailed = DictionaryErr("unable to update: key not found")
	ErrDeleteFailed = DictionaryErr("unable to delete: key not found")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrDeleteFailed
	case nil:
		delete(d, key)
		return nil
	default:
		return nil
	}
}

func (d Dictionary) Update(key, val string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrUpdateFailed
	case nil:
		d[key] = val
		return nil
	default:
		return nil
	}
}

func (d Dictionary) Add(key, val string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = val
	case nil:
		return ErrDupKey
	default:
		return err
	}
	return nil
}
