// Copyright (c) 2014 Canonical Ltd.
// Licensed under the GPLv3, see the COPYING file for details.

package textsecure

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Contact contains information about a contact.
type Contact struct {
	Name  string
	Tel   string
	Photo string
}

type yamlContacts struct {
	Contacts []Contact
}

// ReadContacts reads a YAML contacts file
func ReadContacts(fileName string) ([]Contact, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	contacts := &yamlContacts{}
	err = yaml.Unmarshal(b, contacts)
	if err != nil {
		return nil, err
	}
	return contacts.Contacts, nil
}

// WriteContacts saves a list of contacts to a file
func WriteContacts(filename string, contacts []Contact) error {
	c := &yamlContacts{contacts}
	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, b, 0600)
}
