// Use to insert a new contact into the database
package main

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
	"github.com/jinzhu/gorm"
	"github.com/quorumsco/cmd"
	"github.com/quorumsco/contacts/models"
	"github.com/quorumsco/databases"
	"github.com/quorumsco/elastic"
	"github.com/quorumsco/logs"
	"github.com/quorumsco/settings"
)

type result struct {
	_id        string
	_timestamp string
}

var (
	//TIMEOUT time between each try
	TIMEOUT = 5 * time.Second
	//RETRY number of tries
	RETRY = 3
)

func main() {
	cmd := cmd.New()
	cmd.Name = "Indexer"
	cmd.Usage = "Indexes bruteforce added contacts"
	cmd.Version = "0.0.1"
	cmd.Before = script
	cmd.Flags = append(cmd.Flags, []cli.Flag{
		cli.StringFlag{Name: "config, c", Usage: "configuration file", EnvVar: "CONFIG"},
		cli.HelpFlag,
	}...)
	cmd.RunAndExitOnError()
}

func script(ctx *cli.Context) error {
	var err error
	var config settings.Config
	if ctx.String("config") != "" {
		config, err = settings.Parse(ctx.String("config"))
		if err != nil {
			logs.Error(err)
		}
	}

	logs.Level(logs.DebugLevel)

	dialect, args, err := config.SqlDB()
	if err != nil {
		logs.Critical(err)
		os.Exit(1)
	}
	logs.Debug("database type: %s", dialect)

	var db *gorm.DB
	if db, err = databases.InitGORM(dialect, args); err != nil {
		logs.Critical(err)
		os.Exit(1)
	}
	logs.Debug("connected to %s", args)

	if config.Migrate() {
		db.AutoMigrate(models.Models()...)
		logs.Debug("database migrated successfully")
	}

	db.LogMode(true)

	ElasticSettings, err := config.Elasticsearch()
	var client *elastic.Client
	client, err = dialElasticRetry(ElasticSettings.String())
	if err != nil {
		logs.Critical(err)
		os.Exit(1)
	}

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists("contacts").Do()
	if err != nil {
		logs.Critical(err)
		os.Exit(1)
	}
	if !exists {
		logs.Critical("No contacts index")
		os.Exit(1)
	}

	ID, err := findID(client)
	if err != nil {
		logs.Debug(err)
		os.Exit(1)
	}
	logs.Debug("Last ID is : %d", ID)

	contacts, err := findContacts(ID, db)

	for _, contact := range contacts {
		contact, err := addAddresses(contact, db)
		if err != nil {
			logs.Critical(err)
			os.Exit(1)
		}
		logs.Debug("Indexing contact %d : %s %s", contact.ID, contact.Surname, contact.Firstname)
		err = index(contact, client)
		if err != nil {
			logs.Critical(err)
			os.Exit(1)
		}
	}

	return nil
}

func addAddresses(c models.Contact, db *gorm.DB) (*models.Contact, error) {
	if err := db.Where(c.AddressID).First(&c.Address).Error; err != nil {
		if err == gorm.RecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &c, nil
}

func findContacts(ID int, db *gorm.DB) ([]models.Contact, error) {
	var contacts []models.Contact

	err := db.Where("id > ?", ID).Find(&contacts).Error
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

//Find the last indexed Contact's ID
func findID(s *elastic.Client) (int, error) {
	termQuery := elastic.NewMatchAllQuery()
	searchResult, err := s.Search().
		Index("contacts").
		Query(&termQuery).
		Fields("_timestamp").
		Sort("_timestamp", false).
		Size(1).
		Pretty(true).
		Do()
	if err != nil {
		logs.Critical(err)
		return 0, err
	}

	var ID int
	if searchResult.Hits != nil {
		for _, hit := range searchResult.Hits.Hits {
			ID, err = strconv.Atoi(hit.Id)
			if err != nil {
				return 0, err
			}
		}
	} else {
		logs.Debug("No results, ID is 1")
		return 1, nil
	}

	return ID, nil
}

// Index indexes a contact into elasticsearch
func index(Contact *models.Contact, s *elastic.Client) error {
	id := strconv.Itoa(int(Contact.ID))
	if id == "" {
		logs.Error("id is nil")
		return errors.New("id is nil")
	}

	_, err := s.Index().
		Index("contacts").
		Type("contact").
		Id(id).
		BodyJson(Contact).
		Do()
	if err != nil {
		logs.Critical(err)
		return err
	}

	logs.Debug("Indexed")

	return nil
}

// We need a retry because elasticsearch takes a bit of time to be up and running before we can connect to it
func dialElasticRetry(address string) (*elastic.Client, error) {
	var client *elastic.Client
	var err error

	var i int
retry:
	for {
		client, err = elastic.NewClient(elastic.SetURL(address))
		switch {
		case err == nil:
			break retry
		case i >= RETRY:
			return nil, err
		default:
			logs.Error(err)
			i++
		}
		time.Sleep(TIMEOUT)
	}

	return client, nil
}
