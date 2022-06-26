package game

import (
	tea "github.com/charmbracelet/bubbletea"
	"pong/pkg/consts"
	"time"
)

type Game struct {
	screenWidth  int
	screenHeight int
	ball         *Ball
	player1      *Player
	player2      *Player
}

type tickMsg time.Time

func updateBallPosition() tea.Cmd {
	now := time.Now()
	timeSince := time.Second * 2 / consts.FPS
	return tea.Tick(timeSince, func(t time.Time) tea.Msg {
		return tickMsg(now)
	})
}
func (g *Game) Init() tea.Cmd {
	return updateBallPosition()
}

func NewGame() *Game {
	return &Game{
		screenWidth:  consts.ScreenWidth,
		screenHeight: consts.ScreenHeight,
		ball:         NewBall(),
		player1:      NewPlayer("Player 1", 1),
		player2:      NewPlayer("Player 2", 2),
	}
}

func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		g.ball.Move(g.player1, g.player2)
		_, hasEnded := g.GetState()
		if hasEnded {
			return g, nil
		}
		return g, updateBallPosition()

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return g, tea.Quit
		case "up":
			g.player2.MoveUP()
		case "down":
			g.player2.MoveDown()
		case "w":
			g.player1.MoveUP()
		case "s":
			g.player1.MoveDown()
		case "r":
			g.Reset()
			return g, updateBallPosition()
		}

	}
	return g, nil
}

func (g *Game) View() string {
	screen := ""
	for i := 0; i < g.screenHeight; i++ {
		for j := 0; j < g.screenWidth; j++ {
			screen += g.render(i, j)
		}
		screen += "\n"
	}

	winner, hasEnded := g.GetState()
	if hasEnded {
		return winner
	}

	return screen
}

func (g *Game) GetState() (string, bool) {

	if g.ball.PositionX == g.screenWidth-1 {
		return "Player 1 won \nPress R to restart", true
	}
	if g.ball.PositionX == 0 {
		return "Player 2 won \nPress R to restart", true
	}
	return "", false
}

func (g *Game) Reset() {
	g.ball.Reset()
	g.player1.Reset()
	g.player2.Reset()
}
