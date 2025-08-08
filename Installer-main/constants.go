/*
 * SPDX-License-Identifier: GPL-3.0
 * Vencord Installer Enhanced Constants, a cross platform gui/cli app for installing Vencord
 * Copyright (c) 2023 Vendicated and Vencord contributors
 * Enhanced with modern UI color scheme and styling constants
 */

package main

import (
	"image/color"
	"vencordinstaller/buildinfo"
)

const ReleaseUrl = "https://api.github.com/repos/Vendicated/Vencord/releases/latest"
const ReleaseUrlFallback = "https://vencord.dev/releases/vencord"
const InstallerReleaseUrl = "https://api.github.com/repos/Vencord/Installer/releases/latest"
const InstallerReleaseUrlFallback = "https://vencord.dev/releases/installer"

var UserAgent = "VencordInstaller/" + buildinfo.InstallerGitHash + " (https://github.com/Vencord/Installer)"

// Enhanced Discord colors (improved for better visibility)
var (
	DiscordGreen  = color.RGBA{R: 0x57, G: 0xF2, B: 0x87, A: 0xFF} // Enhanced success green
	DiscordRed    = color.RGBA{R: 0xF0, G: 0x43, B: 0x47, A: 0xFF} // Enhanced error red
	DiscordBlue   = color.RGBA{R: 0x5B, G: 0x65, B: 0xEA, A: 0xFF} // Enhanced primary blue
	DiscordYellow = color.RGBA{R: 0xFE, G: 0xE7, B: 0x5C, A: 0xFF} // Enhanced warning yellow

	// Additional enhanced colors
	DiscordPurple    = color.RGBA{R: 0x7C, G: 0x3A, B: 0xED, A: 0xFF} // Purple accent
	DiscordOrange    = color.RGBA{R: 0xFF, G: 0x73, B: 0x3C, A: 0xFF} // Orange accent
	DiscordTextLight = color.RGBA{R: 0xDC, G: 0xDD, B: 0xDE, A: 0xFF} // Light text
)

// Enhanced Modern Color Palette
var (
	// Background Colors
	DiscordDarkBg     = color.RGBA{R: 0x2C, G: 0x2F, B: 0x33, A: 0xFF} // Primary dark background
	DiscordLightBg    = color.RGBA{R: 0x36, G: 0x39, B: 0x3F, A: 0xFF} // Secondary lighter background
	DiscordCardBg     = color.RGBA{R: 0x40, G: 0x43, B: 0x4B, A: 0xFF} // Card/panel background
	DiscordHoverBg    = color.RGBA{R: 0x4A, G: 0x4D, B: 0x57, A: 0xFF} // Hover state background

	// Accent Colors
	DiscordAccent     = color.RGBA{R: 0x5D, G: 0x65, B: 0xF3, A: 0xFF} // Primary accent (modern blue)
	DiscordAccentHover = color.RGBA{R: 0x6D, G: 0x75, B: 0xFF, A: 0xFF} // Accent hover state
	DiscordAccentActive = color.RGBA{R: 0x4D, G: 0x55, B: 0xE3, A: 0xFF} // Accent active state

	// Status Colors
	DiscordSuccess    = color.RGBA{R: 0x3B, G: 0xA5, B: 0x5C, A: 0xFF} // Success green
	DiscordSuccessHover = color.RGBA{R: 0x4B, G: 0xC5, B: 0x6C, A: 0xFF} // Success hover
	DiscordDanger     = color.RGBA{R: 0xED, G: 0x42, B: 0x45, A: 0xFF} // Danger red
	DiscordDangerHover = color.RGBA{R: 0xFF, G: 0x52, B: 0x55, A: 0xFF} // Danger hover
	DiscordWarning    = color.RGBA{R: 0xFF, G: 0xA5, B: 0x00, A: 0xFF} // Warning orange
	DiscordWarningHover = color.RGBA{R: 0xFF, G: 0xB5, B: 0x10, A: 0xFF} // Warning hover
	DiscordInfo       = color.RGBA{R: 0x00, G: 0xB4, B: 0xFF, A: 0xFF} // Info blue

	// Text Colors
	DiscordText       = color.RGBA{R: 0xDC, G: 0xDD, B: 0xDE, A: 0xFF} // Primary text (light)
	DiscordTextMuted  = color.RGBA{R: 0x96, G: 0x99, B: 0x9D, A: 0xFF} // Secondary text (muted)
	DiscordTextDark   = color.RGBA{R: 0x2E, G: 0x31, B: 0x38, A: 0xFF} // Dark text for light backgrounds
	DiscordTextWhite  = color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF} // Pure white text

	// Border and Outline Colors
	DiscordBorder     = color.RGBA{R: 0x4F, G: 0x54, B: 0x5C, A: 0xFF} // Standard border
	DiscordBorderLight = color.RGBA{R: 0x72, G: 0x76, B: 0x7D, A: 0xFF} // Light border
	DiscordBorderDark = color.RGBA{R: 0x1E, G: 0x21, B: 0x25, A: 0xFF} // Dark border

	// Glass Morphism Effects
	GlassOverlay      = color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x0A} // 4% white overlay
	GlassBorder       = color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x1A} // 10% white border
	GlassBackground   = color.RGBA{R: 0x2C, G: 0x2F, B: 0x33, A: 0xE6} // 90% opacity background

	// Gradient Colors
	GradientStart     = color.RGBA{R: 0x5D, G: 0x65, B: 0xF3, A: 0xFF} // Blue start
	GradientMid       = color.RGBA{R: 0x7C, G: 0x3A, B: 0xED, A: 0xFF} // Purple middle
	GradientEnd       = color.RGBA{R: 0xED, G: 0x42, B: 0x95, A: 0xFF} // Pink end

	// Shadow Colors
	ShadowLight       = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x1A} // 10% black shadow
	ShadowMedium      = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x33} // 20% black shadow
	ShadowDark        = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x4D} // 30% black shadow

	// Special Effect Colors
	NeonGlow          = color.RGBA{R: 0x5D, G: 0x65, B: 0xF3, A: 0x80} // Neon glow effect
	ParticleColor     = color.RGBA{R: 0x5D, G: 0x65, B: 0xF3, A: 0x40} // Particle effects
	ShimmerColor      = color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x20} // Shimmer effect
)

// Animation Constants
const (
	AnimationDuration     = 0.3  // Default animation duration in seconds
	FastAnimationDuration = 0.15 // Fast animation duration
	SlowAnimationDuration = 0.5  // Slow animation duration
	HoverAnimationSpeed   = 0.1  // Hover animation increment per frame
	FadeAnimationSpeed    = 0.05 // Fade animation increment per frame
	RotationSpeed         = 2.0  // Rotation speed for loading spinners
	ParticleSpeed         = 0.016 // Particle animation speed
)

// UI Constants
const (
	DefaultBorderRadius   = 8.0  // Default border radius
	LargeBorderRadius     = 12.0 // Large border radius for cards/modals
	SmallBorderRadius     = 4.0  // Small border radius for small elements
	ButtonHeight          = 40.0 // Standard button height
	LargeButtonHeight     = 55.0 // Large button height
	SmallButtonHeight     = 30.0 // Small button height
	DefaultPadding        = 16.0 // Default padding
	LargePadding          = 24.0 // Large padding
	SmallPadding          = 8.0  // Small padding
)

// Typography Constants
const (
	HeaderFontSize        = 48.0 // Main header font size
	SubHeaderFontSize     = 36.0 // Sub header font size
	TitleFontSize         = 32.0 // Title font size
	LargeFontSize         = 20.0 // Large text font size
	DefaultFontSize       = 16.0 // Default font size
	SmallFontSize         = 14.0 // Small font size
	CaptionFontSize       = 12.0 // Caption font size
)

// Linux Discord Names (unchanged)
var LinuxDiscordNames = []string{
	"Discord",
	"DiscordPTB",
	"DiscordCanary",
	"DiscordDevelopment",
	"discord",
	"discordptb",
	"discordcanary",
	"discorddevelopment",
	"discord-ptb",
	"discord-canary",
	"discord-development",
	// Flatpak
	"com.discordapp.Discord",
	"com.discordapp.DiscordPTB",
	"com.discordapp.DiscordCanary",
	"com.discordapp.DiscordDevelopment",
}

// Theme Presets
type ThemePreset struct {
	Name            string
	PrimaryColor    color.RGBA
	SecondaryColor  color.RGBA
	BackgroundColor color.RGBA
	TextColor       color.RGBA
	AccentColor     color.RGBA
}

var (
	// Default Discord Theme
	DiscordTheme = ThemePreset{
		Name:            "Discord",
		PrimaryColor:    DiscordAccent,
		SecondaryColor:  DiscordLightBg,
		BackgroundColor: DiscordDarkBg,
		TextColor:       DiscordText,
		AccentColor:     DiscordAccent,
	}

	// Dark Theme
	DarkTheme = ThemePreset{
		Name:            "Dark",
		PrimaryColor:    color.RGBA{R: 0x1A, G: 0x1A, B: 0x1A, A: 0xFF},
		SecondaryColor:  color.RGBA{R: 0x2D, G: 0x2D, B: 0x2D, A: 0xFF},
		BackgroundColor: color.RGBA{R: 0x0D, G: 0x0D, B: 0x0D, A: 0xFF},
		TextColor:       color.RGBA{R: 0xF0, G: 0xF0, B: 0xF0, A: 0xFF},
		AccentColor:     color.RGBA{R: 0x00, G: 0x7A, B: 0xCC, A: 0xFF},
	}

	// Light Theme
	LightTheme = ThemePreset{
		Name:            "Light",
		PrimaryColor:    color.RGBA{R: 0xF8, G: 0xF9, B: 0xFA, A: 0xFF},
		SecondaryColor:  color.RGBA{R: 0xE3, G: 0xE5, B: 0xE8, A: 0xFF},
		BackgroundColor: color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
		TextColor:       color.RGBA{R: 0x2E, G: 0x31, B: 0x38, A: 0xFF},
		AccentColor:     color.RGBA{R: 0x00, G: 0x7A, B: 0xCC, A: 0xFF},
	}

	// Neon Theme
	NeonTheme = ThemePreset{
		Name:            "Neon",
		PrimaryColor:    color.RGBA{R: 0x0A, G: 0x0A, B: 0x0A, A: 0xFF},
		SecondaryColor:  color.RGBA{R: 0x1A, G: 0x1A, B: 0x1A, A: 0xFF},
		BackgroundColor: color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF},
		TextColor:       color.RGBA{R: 0x00, G: 0xFF, B: 0xFF, A: 0xFF},
		AccentColor:     color.RGBA{R: 0xFF, G: 0x00, B: 0xFF, A: 0xFF},
	}
)