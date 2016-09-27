package dbcontext

import (
	"log"

	r "gopkg.in/dancannon/gorethink.v2"
)

type DbContext struct {
	Response string
}

func (db *DbContext) Connect() *DbContext {
	opts := r.ConnectOpts{
		Address:  Host,
		Database: Database,
	}

	session, err := r.Connect(opts)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	res, err := r.Expr("Hello world").Run(session)
	if err != nil {
		log.Fatalln(err)
	}

	err = res.One(&db.Response)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
