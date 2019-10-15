// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/moov-io/base"
	moovhttp "github.com/moov-io/base/http"
	client "github.com/moov-io/customers/client"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

var (
	errNoDisclaimerID = errors.New("no Disclaimer ID found")
)

func addDisclaimerRoute(logger log.Logger, r *mux.Router, repo disclaimerRepository) {
	r.Methods("GET").Path("/customers/{customerID}/disclaimers").HandlerFunc(getCustomerDisclaimers(logger, repo))
	r.Methods("POST").Path("/customers/{customerID}/disclaimers/{disclaimerID}").HandlerFunc(acceptDisclaimer(logger, repo))
}

func getDisclaimerID(w http.ResponseWriter, r *http.Request) string {
	v, ok := mux.Vars(r)["disclaimerID"]
	if !ok || v == "" {
		moovhttp.Problem(w, errNoDisclaimerID)
		return ""
	}
	return v
}

func getCustomerDisclaimers(logger log.Logger, repo disclaimerRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		customerID := getCustomerID(w, r)
		if customerID == "" {
			return
		}

		disclaimers, err := repo.getCustomerDisclaimers(customerID)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(disclaimers)
	}
}

func acceptDisclaimer(logger log.Logger, repo disclaimerRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		customerID, disclaimerID := getCustomerID(w, r), getDisclaimerID(w, r)
		if customerID == "" || disclaimerID == "" {
			return
		}

		if err := repo.acceptDisclaimer(customerID, disclaimerID); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

type disclaimerRepository interface {
	getCustomerDisclaimers(customerID string) ([]*client.Disclaimer, error)
	acceptDisclaimer(customerID, disclaimerID string) error
}

type sqlDisclaimerRepository struct {
	db     *sql.DB
	logger log.Logger
}

func (r *sqlDisclaimerRepository) close() error {
	return r.db.Close()
}

func (r *sqlDisclaimerRepository) getCustomerDisclaimers(customerID string) ([]*client.Disclaimer, error) {
	query := `select d.disclaimer_id, d.text, d.document_id, da.accepted_at from disclaimers as d
left outer join disclaimer_acceptances as da on d.disclaimer_id = da.disclaimer_id
where d.deleted_at is null;`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*client.Disclaimer
	for rows.Next() {
		disc := &client.Disclaimer{}
		var acceptedAt *time.Time
		if err := rows.Scan(&disc.ID, &disc.Text, &disc.DocumentID, &acceptedAt); err != nil {
			return nil, err
		}
		if acceptedAt != nil && !acceptedAt.IsZero() {
			disc.AcceptedAt = *acceptedAt
		}
		out = append(out, disc)
	}
	return out, rows.Err()
}

func (r *sqlDisclaimerRepository) acceptDisclaimer(customerID, disclaimerID string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `select disclaimer_id from disclaimers where disclaimer_id = ? and deleted_at is null limit 1;`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		tx.Rollback()
		return err
	}

	row := stmt.QueryRow(disclaimerID)
	var discID string
	if err := row.Scan(&discID); discID != disclaimerID || err != nil {
		stmt.Close()
		return fmt.Errorf("acceptDisclaimer: missing disclaimer: %v rollback=%v", err, tx.Rollback())
	}
	stmt.Close()

	// write the acceptance row now
	query = `insert into disclaimer_acceptances (disclaimer_id, customer_id, accepted_at) values (?, ?, ?);`
	stmt, err = tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(disclaimerID, customerID, time.Now())
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *sqlDisclaimerRepository) insertDisclaimer(text, documentID string) (*client.Disclaimer, error) {
	query := `insert into disclaimers (disclaimer_id, text, document_id, created_at) values (?, ?, ?, ?);`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	disc := &client.Disclaimer{
		ID:         base.ID(),
		Text:       text,
		DocumentID: documentID,
	}
	_, err = stmt.Exec(disc.ID, disc.Text, disc.DocumentID, time.Now())
	return disc, err
}