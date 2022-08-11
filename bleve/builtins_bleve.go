// +build b_bleve

package bleve

import (
	"rye/env"
	"rye/evaldo"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search/query"

	"fmt"
)

var Builtins_bleve = map[string]*env.Builtin{

	"new-bleve-mapping": {
		Argsn: 0,
		Fn: func(ps *env.ProgramState, arg0 env.Object, arg1 env.Object, arg2 env.Object, arg3 env.Object, arg4 env.Object) env.Object {
			mapping := bleve.NewIndexMapping()
			return *env.NewNative(ps.Idx, mapping, "bleve-mapping")
		},
	},

	"bleve-mapping//open": {
		Argsn: 2,
		Fn: func(ps *env.ProgramState, arg0 env.Object, arg1 env.Object, arg2 env.Object, arg3 env.Object, arg4 env.Object) env.Object {
			switch mpi := arg0.(type) {
			case env.Native:
				switch s := arg1.(type) {
				case env.Uri:
					path := strings.Split(s.Path, "://")
					index, err := bleve.New(path[1], mpi.Value.(mapping.IndexMapping))
					if err != nil {
						return evaldo.MakeError(ps, err.Error())
					}
					return *env.NewNative(ps.Idx, index, "bleve-index")
				default:
					return evaldo.MakeError(ps, "Arg 2 not file Uri.")
				}
			default:
				return evaldo.MakeError(ps, "Arg 1 not native.")
			}
		},
	},

	"bleve-index//index": {
		Argsn: 3,
		Doc:   "[ ses-session* gomail-message from-email recipients ]",
		Fn: func(ps *env.ProgramState, arg0 env.Object, arg1 env.Object, arg2 env.Object, arg3 env.Object, arg4 env.Object) env.Object {
			switch idx := arg0.(type) {
			case env.Native:
				switch ident := arg1.(type) { // gomail-message
				case env.String:
					switch text := arg2.(type) { // recipients
					case env.String:
						err := idx.Value.(bleve.Index).Index(ident.Value, text.Value)
						if err != nil {
							return evaldo.MakeError(ps, err.Error())
						}
						return arg0
					default:
						return evaldo.MakeError(ps, "A3 not String")
					}
				default:
					return evaldo.MakeError(ps, "A2 not String")
				}
			default:
				return evaldo.MakeError(ps, "A1 not Native")
			}
			return nil
		},
	},

	"new-match-query": {
		Argsn: 1,
		Fn: func(ps *env.ProgramState, arg0 env.Object, arg1 env.Object, arg2 env.Object, arg3 env.Object, arg4 env.Object) env.Object {
			switch text := arg0.(type) {
			case env.String:
				query := bleve.NewMatchQuery(text.Value)
				return *env.NewNative(ps.Idx, query, "bleve-query")
			default:
				return evaldo.MakeError(ps, "Arg 1 not String.")
			}
		},
	},
	"bleve-query//new-search-request": {
		Argsn: 1,
		Fn: func(ps *env.ProgramState, arg0 env.Object, arg1 env.Object, arg2 env.Object, arg3 env.Object, arg4 env.Object) env.Object {
			switch qry := arg0.(type) {
			case env.Native:
				search := bleve.NewSearchRequest(qry.Value.(query.Query))
				return *env.NewNative(ps.Idx, search, "bleve-search")
			default:
				return evaldo.MakeError(ps, "Arg 1 not native.")
			}
		},
	},
	"bleve-search//search": {
		Argsn: 2,
		Fn: func(ps *env.ProgramState, arg0 env.Object, arg1 env.Object, arg2 env.Object, arg3 env.Object, arg4 env.Object) env.Object {
			switch search := arg0.(type) {
			case env.Native:
				switch index := arg1.(type) {
				case env.Native:
					searchResults, err := index.Value.(bleve.Index).Search(search.Value.(*bleve.SearchRequest))
					if err != nil {
						return evaldo.MakeError(ps, err.Error())
					}
					fmt.Println(searchResults)
					return *env.NewNative(ps.Idx, searchResults, "bleve-results")
				default:
					return evaldo.MakeError(ps, "Arg 1 not native.")
				}
			default:
				return evaldo.MakeError(ps, "Arg 1 not native.")
			}
		},
	},
}
