package cm

import (
	"database/sql"
	"errors"
	"regexp"

	_ "github.com/lib/pq"
)

// errors
var (
        errDBInitFailed     = errors.New("failed to init contact db")
        errFailedRegExpComp = errors.New("failed to init regexp checker")
        errWrongNumFormat   = errors.New("phone number invalid format, need eleven digits")
        errFailedAddCont    = errors.New("failed adding a contact")
        errFailedDeleteCont = errors.New("database failed executing 'delete' query")
        errFailedCheckAff   = errors.New("failed to check affected rows")
        ErrContDelNotFound  = errors.New("contact to delete not found")
        errFailedSelectCont = errors.New("failed to select contacts")
        errFailedUpdateCont = errors.New("failed to update contacts")
        errFailedGetRowsAff = errors.New("failed to retrieve rows affected")
        ErrContUpdNotFound  = errors.New("contact to update not found")
        errFailedRowScan    = errors.New("failed to scan one of rows selected")
        ErrDuplicateAdded   = errors.New("tried to add a record with duplicate number")
)

// queries
var (
        queryInit = `
  CREATE TABLE IF NOT EXISTS contacts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(64) NOT NULL,
    number VARCHAR(16) NOT NULL,
    CONSTRAINT unique_number UNIQUE (number)
  )`
        queryAddCont = `
  INSERT INTO contacts (name, number)
  VALUES ($1, $2)
  RETURNING id`
        queryAddContID = `
  INSERT INTO contacts (id, name, number)
  VALUES ($1, $2, $3)`
        queryDelete = `
  DELETE FROM contacts WHERE id = $1`
        querySelect = `
  SELECT id, name, number FROM contacts WHERE id = $1`
        queryUpdate = `
  UPDATE contacts SET name = $1, number = $2 WHERE id = $3`
        querySelectAll = `
  SELECT id, name, number FROM contacts ORDER BY name`
)

type Contact struct {
        ID     string `json:"id"`
        Name   string `json:"name"`
        Number string `json:"number"`
}

type ContMan struct {
        database *sql.DB
}

func New(db *sql.DB) *ContMan {
        return &ContMan{database: db}
}

func (contMan *ContMan) Init() error {
        _, err := contMan.database.Exec(queryInit)
        if err != nil {
                return errors.Join(errDBInitFailed, err)
        }
        
        return nil
}

func (contMan *ContMan) Add(contact Contact) error {
        numberRegExp, err := regexp.Compile(`^\+[0-9]?[0-9]{11}$`)
        if err != nil {
                return errors.Join(errFailedRegExpComp, err)
        }
        
        if !numberRegExp.MatchString(contact.Number) {
                return errWrongNumFormat
        }
        
        err = contMan.database.QueryRow(queryAddCont, contact.Name, contact.Number).Scan(&contact.ID)
        if err != nil {
                if err.Error() == `pq: duplicate key value violates unique constraint "unique_number"` {
                        return errors.Join(ErrDuplicateAdded, err)
                } else {
                        return errors.Join(errFailedAddCont, err)
                }
        }

        return nil
}

func (contMan *ContMan) Get(id string) (Contact, error) {
        var contact Contact
        
        err := contMan.database.QueryRow(querySelect, id).Scan(&contact.ID, &contact.Name, &contact.Number)
        if err != nil {
                return contact, errors.Join(errFailedSelectCont, err)
        }
        
        return contact, nil
}

func (contMan *ContMan) GetAll() ([]Contact, error) {
        rows, err := contMan.database.Query(querySelectAll)
        if err != nil {
                return nil, errors.Join(errFailedSelectCont, err)
        }
        defer rows.Close()

        var contacts []Contact
        for rows.Next() {
                var contact Contact
                err := rows.Scan(&contact.ID, &contact.Name, &contact.Number)
                if err != nil {
                        return nil, errors.Join(errFailedRowScan, err)
                }

                contacts = append(contacts, contact)
        }

        return contacts, nil
}

func (contMan *ContMan) Delete(id string) error {
        result, err := contMan.database.Exec(queryDelete, id)
        if err != nil {
                return errors.Join(errFailedDeleteCont, err)
        }

        rowsAffected, err := result.RowsAffected()
        if err != nil {
                return errors.Join(errFailedCheckAff, err)
        }
        if rowsAffected == 0 {
                return errors.Join(ErrContDelNotFound, err)
        }

        return nil
}

func (contMan *ContMan) Update(contact Contact) error {
        numberRegExp, err := regexp.Compile(`^\+[0-9]?[0-9]{11}$`)
        if err != nil {
                return errors.Join(errFailedRegExpComp, err)
        }

        if !numberRegExp.MatchString(contact.Number) {
                return errWrongNumFormat
        }

        result, err := contMan.database.Exec(queryUpdate, contact.Name, contact.Number, contact.ID)
        if err != nil {
                if err.Error() == `pq: duplicate key value violates unique constraint "unique_number"` {
                        return errors.Join(ErrDuplicateAdded, err)
                } else {
                        return errors.Join(errFailedUpdateCont, err)
                }
        }

        rowsAffected, err := result.RowsAffected()
        if err != nil {
                return errors.Join(errFailedGetRowsAff, err)
        }
        if rowsAffected == 0 {
                return ErrContUpdNotFound
        }

        return nil
}
