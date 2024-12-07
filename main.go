package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Img  *ebiten.Image
	X, Y float64
}

type Player struct {
	*Sprite
	Health uint
}

type Enemy struct {
	*Sprite
	FollowsPlayer bool
}

type Potion struct {
	*Sprite
	AmtHeal uint
}

type Game struct {
	player  *Player
	enemies []*Enemy
	potions []*Potion
}

func (g *Game) Update() error {

	//react to key presses
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.player.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.player.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.player.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.player.Y += 2
	}

	for _, sprite := range g.enemies {
		if sprite.FollowsPlayer {
			if sprite.X < g.player.X {
				sprite.X += 0.5
			} else if sprite.X > g.player.X {
				sprite.X -= 0.5
			}
			if sprite.Y < g.player.Y {
				sprite.Y += 0.5
			} else if sprite.Y > g.player.Y {
				sprite.Y -= 0.5
			}
		}
	}

	for _, potion := range g.potions {
		if g.player.X > potion.X {
			g.player.Health += potion.AmtHeal
			fmt.Printf("Picked up potion! Health: %d", g.player.Health)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.player.X, g.player.Y)
	//draw our player
	screen.DrawImage(
		g.player.Img.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image),
		&opts,
	)

	opts.GeoM.Reset()

	for _, sprite := range g.enemies {
		opts.GeoM.Translate(sprite.X, sprite.Y)

		screen.DrawImage(
			sprite.Img.SubImage(
				image.Rect(0, 0, 16, 16),
			).(*ebiten.Image),
			&opts,
		)

		opts.GeoM.Reset()
	}
	for _, sprite := range g.potions {
		opts.GeoM.Translate(sprite.X, sprite.Y)

		screen.DrawImage(
			sprite.Img.SubImage(
				image.Rect(0, 0, 16, 16),
			).(*ebiten.Image),
			&opts,
		)

		opts.GeoM.Reset()
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Endless Journey")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/ninja.png")
	if err != nil {
		//handle error
		log.Fatal(err)
	}
	skeletonImg, _, err := ebitenutil.NewImageFromFile("assets/images/skeleton.png")
	if err != nil {
		//handle error
		log.Fatal(err)
	}
	potionImg, _, err := ebitenutil.NewImageFromFile("assets/images/potion.png")
	if err != nil {
		//handle error
		log.Fatal(err)
	}

	game := Game{
		player: &Player{
			Sprite: &Sprite{
				Img: playerImg,
				X:   50.0,
				Y:   50.0,
			},
			Health: 3,
		},

		enemies: []*Enemy{
			{
				&Sprite{
					Img: skeletonImg,
					X:   100.0,
					Y:   100.0,
				},
				true,
			},
			{
				&Sprite{
					Img: skeletonImg,
					X:   150.0,
					Y:   150.0,
				},
				true,
			},
			{
				&Sprite{
					Img: skeletonImg,
					X:   200.0,
					Y:   200.0,
				},
				true,
			},
		},
		potions: []*Potion{
			{
				&Sprite{
					Img: potionImg,
					X:   210.0,
					Y:   100.0,
				},
				1.0,
			},
		},
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
