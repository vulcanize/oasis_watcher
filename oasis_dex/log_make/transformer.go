// Copyright 2018 Vulcanize
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package log_make

import (
	"log"

	"github.com/8thlight/oasis_watcher/oasis_dex/constants"
	"github.com/vulcanize/vulcanizedb/libraries/shared"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
)

type Transformer struct {
	Converter              Converter
	WatchedEventRepository datastore.WatchedEventRepository
	FilterRepository       datastore.FilterRepository
	Repository             Datastore
}

func NewTransformer(db *postgres.DB, blockchain core.Blockchain) shared.Transformer {
	var transformer shared.Transformer
	cnvtr := LogMakeConverter{}
	wer := repositories.WatchedEventRepository{DB: db}
	fr := repositories.FilterRepository{DB: db}
	repo := Repository{DB: db}
	transformer = &Transformer{Converter: cnvtr, WatchedEventRepository: wer, FilterRepository: fr, Repository: repo}
	for _, filter := range constants.LogMakeFilters {
		fr.CreateFilter(filter)
	}
	return transformer
}

func (logMakeTransformer *Transformer) Execute() error {
	for _, filter := range constants.LogMakeFilters {
		watchedEvents, err := logMakeTransformer.WatchedEventRepository.GetWatchedEvents(filter.Name)
		if err != nil {
			log.Println("Error fetching events for LogMake filter:", err)
			return err
		}

		for _, we := range watchedEvents {
			err = createLogMakeData(we, logMakeTransformer)
			if err != nil {
				log.Printf("Error persisting data for LogMake (watchedEvent.LogID %s):\n %s", we.LogID, err)
			}
		}
	}
	return nil
}

func createLogMakeData(watchedEvent *core.WatchedEvent, transformer *Transformer) error {
	logMake, err := transformer.Converter.ToEntity(*watchedEvent)
	if err != nil {
		log.Println("Error converting LogMake watched event to log: ", err)
		return err
	}
	model := transformer.Converter.ToModel(*logMake)
	err = transformer.Repository.Create(model, watchedEvent.LogID)
	if err != nil {
		log.Println("Error persisting data for LogMake event:", err)
		return err
	}
	return nil
}
