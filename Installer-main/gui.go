package main

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"image"

	g "github.com/AllenDang/giu"
)

//go:embed winres/icon.png
var iconBytes []byte

// Using default font instead of Whitney (removed embed)
// var Whitney []byte

var (
	win *g.MasterWindow
	customDir string

	// Enhanced colors using standard color.RGBA
	enhancedBg     = DiscordDarkBg
	enhancedAccent = DiscordAccent
	enhancedText   = DiscordText
	enhancedGreen  = DiscordGreen
	enhancedRed    = DiscordRed
)

// Discord installations list
var discords []any

// GUI helper functions
func IsValidInstallLocation() bool {
	return len(discords) > 0
}

func UninstallVencord() error {
	if len(discords) == 0 {
		return errors.New("No Discord installation found")
	}
	return discords[0].(*DiscordInstall).unpatch()
}

func IsOpenAsarInstalled() bool {
	if len(discords) == 0 {
		return false
	}
	return discords[0].(*DiscordInstall).IsOpenAsar()
}

func InstallOpenAsar() error {
	if len(discords) == 0 {
		return errors.New("No Discord installation found")
	}
	return discords[0].(*DiscordInstall).InstallOpenAsar()
}

func UninstallOpenAsar() error {
	if len(discords) == 0 {
		return errors.New("No Discord installation found")
	}
	return discords[0].(*DiscordInstall).UninstallOpenAsar()
}

func main() {
	RunGUI()
}

func RunGUI() {
	// Initialize GitHub downloader and Discord installations
	InitGithubDownloader()
	discords = FindDiscords()
	// g.SetDefaultFontFromBytes(Whitney, 20) // Font file not available

	// Enhanced window with modern title
	win = g.NewMasterWindow("Vencord Installer - Enhanced Edition", 1200, 800, 0)

	icon, _, err := image.Decode(bytes.NewReader(iconBytes))
	if err == nil {
		win.SetIcon([]image.Image{icon})
	}

	win.Run(loop)
}

func loop() {
	g.SingleWindow().Layout(
		g.Style().SetColor(g.StyleColorWindowBg, enhancedBg).To(
			g.Column(
				// Header with enhanced styling
				g.Style().SetColor(g.StyleColorText, enhancedText).To(
					g.Row(
						g.Align(g.AlignCenter).To(
							g.Label("Vencord Installer - Enhanced Edition").Font(g.GetDefaultFonts()[0].SetSize(28)),
						),
					),
				),
				g.Separator(),
				g.Spacing(),
				renderEnhancedInstaller(),
			),
		),
	)
}

func renderEnhancedInstaller() g.Widget {
	return g.Column(
		// Discord installations section
		g.Style().SetColor(g.StyleColorText, enhancedText).To(
			g.Label("Select Discord Installation:").Font(g.GetDefaultFonts()[0].SetSize(18)),
		),
		g.Spacing(),

		// Installation list with enhanced styling
		g.Child().Size(-1, 200).Border(true).Layout(
			g.Style().SetColor(g.StyleColorChildBg, DiscordLightBg).To(
				g.Column(
					renderDiscordInstallations()...,
				),
			),
		),

		g.Spacing(),

		// Custom path section
		g.Style().SetColor(g.StyleColorText, enhancedText).To(
			g.Label("Or specify custom path:"),
		),
		g.Row(
			g.InputText(&customDir).Size(400),
			g.Style().SetColor(g.StyleColorButton, enhancedAccent).To(
				g.Button("Browse").OnClick(func() {
					// File dialog functionality would go here
			// For now, users can manually enter the path
				}),
			),
		),

		g.Spacing(),
		g.Separator(),
		g.Spacing(),

		// Action buttons with enhanced styling
		g.Row(
			// Enhanced Install button with better styling
			g.Style().SetColor(g.StyleColorButton, enhancedGreen).
				SetColor(g.StyleColorButtonHovered, enhancedAccent).
				SetColor(g.StyleColorButtonActive, DiscordSuccessHover).To(
				g.Button("🚀 Install").Size(140, 40).OnClick(func() {
					if !IsValidInstallLocation() {
						showErrorModal("Invalid installation location")
						return
					}
					go func() {
						if err := installLatestBuilds(); err != nil {
							showErrorModal(err.Error())
						}
					}()
				}),
			),
			g.Style().SetColor(g.StyleColorButton, enhancedAccent).To(
				g.Button("Reinstall / Repair").Size(150, 40).OnClick(func() {
					if !IsValidInstallLocation() {
						showErrorModal("Invalid installation location")
						return
					}
					go func() {
						if err := installLatestBuilds(); err != nil {
							showErrorModal(err.Error())
						}
					}()
				}),
			),
			g.Style().SetColor(g.StyleColorButton, enhancedRed).To(
				g.Button("Uninstall").Size(120, 40).OnClick(func() {
					if !IsValidInstallLocation() {
						showErrorModal("Invalid installation location")
						return
					}
					go func() {
						if err := UninstallVencord(); err != nil {
							showErrorModal(err.Error())
						}
					}()
				}),
			),
		),

		g.Spacing(),

		// OpenAsar button
		g.Style().SetColor(g.StyleColorButton, DiscordWarning).To(
			g.Button(Ternary(IsOpenAsarInstalled(), "Uninstall OpenAsar", "Install OpenAsar")).Size(200, 35).OnClick(func() {
				if !IsValidInstallLocation() {
					showErrorModal("Invalid installation location")
					return
				}
				go func() {
					var err error
					if IsOpenAsarInstalled() {
						err = UninstallOpenAsar()
					} else {
						err = InstallOpenAsar()
					}
					if err != nil {
						showErrorModal(err.Error())
					}
				}()
			}),
		),

		// Status and modals
		renderModals(),
	)
}

func renderDiscordInstallations() []g.Widget {
	var widgets []g.Widget

	for _, discord := range discords {
		discordInstall := discord.(*DiscordInstall)
		isSelected := len(widgets) == 0 // Select first one by default

		// Enhanced radio button styling
		widgets = append(widgets, g.Row(
			g.RadioButton(discordInstall.path, isSelected).OnChange(func() {
				// Update selection logic here
				customDir = ""
			}),
			g.Style().SetColor(g.StyleColorText, enhancedText).To(
				g.Label(fmt.Sprintf("%s (%s)", discordInstall.branch, discordInstall.path)),
			),
		))
	}

	return widgets
}

func renderModals() g.Widget {
	return g.Column(
		// Error Modal
		g.PopupModal("Error").Layout(
			g.Style().SetColor(g.StyleColorPopupBg, DiscordCardBg).To(
				g.Column(
					g.Style().SetColor(g.StyleColorText, enhancedRed).To(
						g.Label("Error").Font(g.GetDefaultFonts()[0].SetSize(20)),
					),
					g.Separator(),
					g.Spacing(),
					g.Style().SetColor(g.StyleColorText, enhancedText).To(
						g.Label(errorMessage),
					),
					g.Spacing(),
					g.Style().SetColor(g.StyleColorButton, enhancedAccent).To(
						g.Button("OK").Size(80, 30).OnClick(func() {
							g.CloseCurrentPopup()
						}),
					),
				),
			),
		),

		// Info Modal
		g.PopupModal("Info").Layout(
			g.Style().SetColor(g.StyleColorPopupBg, DiscordCardBg).To(
				g.Column(
					g.Style().SetColor(g.StyleColorText, enhancedAccent).To(
						g.Label("Information").Font(g.GetDefaultFonts()[0].SetSize(20)),
					),
					g.Separator(),
					g.Spacing(),
					g.Style().SetColor(g.StyleColorText, enhancedText).To(
						g.Label(infoMessage),
					),
					g.Spacing(),
					g.Style().SetColor(g.StyleColorButton, enhancedAccent).To(
						g.Button("OK").Size(80, 30).OnClick(func() {
							g.CloseCurrentPopup()
						}),
					),
				),
			),
		),
	)
}

func showErrorModal(message string) {
	errorMessage = message
	g.OpenPopup("Error")
}

func showInfoModal(message string) {
	infoMessage = message
	g.OpenPopup("Info")
}

// Variables for modal messages
var (
	errorMessage string
	infoMessage  string
)