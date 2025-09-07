package grool

import (
	"github.com/newm4n/grool/builder"
	"github.com/newm4n/grool/context"
	"github.com/newm4n/grool/engine"
	"github.com/newm4n/grool/model"
	"github.com/newm4n/grool/pkg"
)

func GrlExecute() (*engine.Grool, *context.DataContext, *model.KnowledgeBase) {
	kb := model.NewKnowledgeBase()
	rb := builder.NewRuleBuilder(kb)
	err := rb.BuildRuleFromResource(pkg.NewFileResource("grool/OrderFlow.grl"))
	if err != nil {
		panic(err)
	}
	dctx := context.NewDataContext()
	engine := engine.NewGroolEngine()
	return engine, dctx, kb
}
