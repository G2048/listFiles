package style

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type Colours struct {
	white  lipgloss.Style
	purple lipgloss.Style
	blue   lipgloss.Style
	green  lipgloss.Style
	red    lipgloss.Style
}

func CreateColours() Colours {
	return Colours{
		white:  lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")),
		purple: lipgloss.NewStyle().Foreground(lipgloss.Color("#A020F0")),
		blue:   lipgloss.NewStyle().Foreground(lipgloss.Color("#0077b6")),
		green:  lipgloss.NewStyle().Foreground(lipgloss.Color("#008000")),
		red:    lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")),
	}
}

type PermissionsStyles struct {
	Colours
	Use   bool
	dir   string
	read  string
	write string
	exec  string
}

func (permission PermissionsStyles) DirStyle() string {
	if permission.Use {
		return permission.blue.Render("d")
	}
	return permission.white.Render("d")
}
func (permission PermissionsStyles) ReadStyle() string {
	if permission.Use {
		return permission.red.Render(permission.read)
	}
	return permission.white.Render(permission.read)
}
func (permission PermissionsStyles) WriteStyle() string {
	if permission.Use {
		return permission.purple.Render(permission.write)
	}
	return permission.white.Render(permission.write)
}
func (permission PermissionsStyles) ExecStyle() string {
	if permission.Use {
		return permission.green.Render(permission.exec)
	}
	return permission.white.Render(permission.exec)
}
func (permission PermissionsStyles) CreateMapPermissions() map[string]string {
	permission.read = "r"
	permission.write = "w"
	permission.exec = "x"
	return map[string]string{"r": permission.ReadStyle(), "w": permission.WriteStyle(), "x": permission.ExecStyle()}
}

func CreateTable(headers []string, rows ...[]string) *table.Table {
	HeaderStyle := lipgloss.NewStyle().Padding(0, 1).Align(lipgloss.Center)
	EvenRowStyle := lipgloss.NewStyle().Padding(0, 1)
	OddRowStyle := lipgloss.NewStyle().Padding(0, 1)

	return table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return HeaderStyle
			case row%2 == 0:
				return EvenRowStyle
			default:
				return OddRowStyle
			}
		}).Headers(headers...).Rows(rows...)

}
