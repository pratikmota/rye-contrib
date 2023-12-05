//go:build b_ebitengine
// +build b_ebitengine

package ebitengine

import (
	// "fmt"
	"rye/env"
	"rye/evaldo"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var onDraw *env.Block

var Ps *env.ProgramState

type Game struct{}

var game Game

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	scr := *env.NewNative(Ps.Idx, screen, "screen")

	// fmt.Println("on--draw")
	// fmt.Println(onDraw)
	ser := Ps.Ser
	Ps.Ser = onDraw.Series
	evaldo.EvalBlockInj(Ps, scr, true)
	Ps.Ser = ser
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

// ** Just a proof of concept module so far **

// ## TODO
// Global Ps and onDraw are undesired but they seemed necesarry for how (statically?) ebitengine semms to work.
// If I overlooked something and there is any solution without the need for this it should be removed!

// ## IDEA
// Should move to external repo rye-alterego, which contrary to main Rye, would focus on desktop / UI / game / windows?
// This would make main Rye and Contrib cleaner and focused on the linux backend tasks, information (pre)processing, ...
// It would also serve as a test if we can move contrib to external module instead of it being a git submodule which
// complicates many things.

var Builtins_ebitengine = map[string]*env.Builtin{

	"ebitengine-run": {
		Argsn: 0,
		Doc:   "TODODOC",
		Fn: func(ps *env.ProgramState, arg0 env.Object, arg1 env.Object, arg2 env.Object, arg3 env.Object, arg4 env.Object) env.Object {
			Ps = ps
			// Specify the window size as you like. Here, a doubled size is specified.
			// fmt.Println(onDraw)
			ebiten.SetWindowSize(320, 240)
			ebiten.SetWindowTitle("Your game's title")
			game := &Game{}

			// Call ebiten.RunGame to start your game loop.
			if err := ebiten.RunGame(game); err != nil {
				// log.Fatal(err)
				return nil
			}
			return nil
		},
	},
	"on-draw": {
		Argsn: 1,
		Doc:   "TODODOC",
		Fn: func(ps *env.ProgramState, arg0 env.Object, arg1 env.Object, arg2 env.Object, arg3 env.Object, arg4 env.Object) env.Object {
			onDraw_ := arg0.(env.Block)
			onDraw = &onDraw_
			// fmt.Println(onDraw)
			Ps = ps
			return nil
		},
	},
	"new-image": {
		Argsn: 1,
		Doc:   "TODODOC",
		Fn: func(ps *env.ProgramState, arg0 env.Object, arg1 env.Object, arg2 env.Object, arg3 env.Object, arg4 env.Object) env.Object {

			img, _, _ := ebitenutil.NewImageFromFile("gopher.png")
			return *env.NewNative(ps.Idx, img, "image")
		},
	},

	"draw-image": {
		Argsn: 2,
		Doc:   "TODODOC",
		Fn: func(ps *env.ProgramState, arg0 env.Object, arg1 env.Object, arg2 env.Object, arg3 env.Object, arg4 env.Object) env.Object {

			arg0.(env.Native).Value.(*ebiten.Image).DrawImage(arg1.(env.Native).Value.(*ebiten.Image), nil)
			img, _, _ := ebitenutil.NewImageFromFile("gopher.png")
			return *env.NewNative(ps.Idx, img, "image")
		},
	},
}
