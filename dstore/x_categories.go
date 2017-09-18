/*
 * Copyright (C) 2015 ~ 2017 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package dstore

var xcategoriesFallback = map[string]string{
	"2dgraphics":       "graphics",
	"3dgraphics":       "graphics",
	"accessibility":    "system",
	"accessories":      "others",
	"actiongame":       "game",
	"advancedsettings": "system",
	"adventuregame":    "game",
	"amusement":        "game",
	"applet":           "others",
	"arcadegame":       "game",
	"archiving":        "system",
	"art":              "office",
	"artificialintelligence": "office",
	"astronomy":              "office",
	"audio":                  "music",
	"audiovideo":             "video",
	"audiovideoediting":      "video",
	"biology":                "office",
	"blocksgame":             "game",
	"boardgame":              "game",
	"building":               "development",
	"calculator":             "system",
	"calendar":               "system",
	"cardgame":               "game",
	"cd":                     "music",
	"chart":                  "office",
	"chat":                   "chat",
	"chemistry":              "office",
	"clock":                  "system",
	"compiz":                 "system",
	"compression":            "system",
	"computerscience":        "office",
	"consoleonly":            "others",
	"contactmanagement":      "chat",
	"core":                   "others",
	"debugger":               "development",
	"desktopsettings":        "system",
	"desktoputility":         "system",
	"development":            "development",
	"dialup":                 "system",
	"dictionary":             "office",
	"discburning":            "system",
	"documentation":          "office",
	"editors":                "others",
	"education":              "office",
	"electricity":            "office",
	"electronics":            "office",
	"email":                  "internet",
	"emulator":               "game",
	"engineering":            "system",
	"favorites":              "others",
	"filemanager":            "system",
	"filesystem":             "system",
	"filetools":              "system",
	"filetransfer":           "internet",
	"finance":                "office",
	"game":                   "game",
	"geography":              "office",
	"geology":                "office",
	"geoscience":             "others",
	"gnome":                  "system",
	"gpe":                    "others",
	"graphics":               "graphics",
	"guidesigner":            "development",
	"hamradio":               "office",
	"hardwaresettings":       "system",
	"ide":                    "development",
	"imageprocessing":        "graphics",
	"instantmessaging":       "chat",
	"internet":               "internet",
	"ircclient":              "chat",
	"kde":                    "system",
	"kidsgame":               "game",
	"literature":             "office",
	"logicgame":              "game",
	"math":                   "office",
	"medicalsoftware":        "office",
	"meteorology":            "others",
	"midi":                   "music",
	"mixer":                  "music",
	"monitor":                "system",
	"motif":                  "others",
	"multimedia":             "video",
	"music":                  "music",
	"network":                "internet",
	"news":                   "reading",
	"numericalanalysis":      "office",
	"ocr":                    "graphics",
	"office":                 "office",
	"p2p":                    "internet",
	"packagemanager":         "system",
	"panel":                  "system",
	"pda":                    "system",
	"photography":            "graphics",
	"physics":                "office",
	"pim":                    "others",
	"player":                 "music",
	"playonlinux":            "others",
	"presentation":           "office",
	"printing":               "office",
	"profiling":              "development",
	"projectmanagement":      "office",
	"publishing":             "office",
	"puzzlegame":             "game",
	"rastergraphics":         "graphics",
	"recorder":               "music",
	"remoteaccess":           "system",
	"revisioncontrol":        "development",
	"robotics":               "office",
	"roleplaying":            "game",
	"scanning":               "office",
	"science":                "office",
	"screensaver":            "others",
	"sequencer":              "music",
	"settings":               "system",
	"simulation":             "game",
	"sportsgame":             "game",
	"spreadsheet":            "office",
	"strategygame":           "game",
	"system":                 "system",
	"systemsettings":         "system",
	"technical":              "others",
	"telephony":              "system",
	"telephonytools":         "system",
	"terminalemulator":       "system",
	"texteditor":             "office",
	"texttools":              "office",
	"transiation":            "development",
	"translation":            "reading",
	"trayicon":               "system",
	"tuner":                  "music",
	"tv":                     "video",
	"utility":                "system",
	"vectorgraphics":         "graphics",
	"video":                  "video",
	"videoconference":        "internet",
	"viewer":                 "graphics",
	"webbrowser":             "internet",
	"webdevelopment":         "development",
	"wine":                   "others",
	"wine-programs-accessories":               "others",
	"wordprocessor":                           "office",
	"x-alsa":                                  "music",
	"x-bible":                                 "reading",
	"x-bluetooth":                             "system",
	"x-debian-applications-emulators":         "game",
	"x-digital_processing":                    "system",
	"x-enlightenment":                         "system",
	"x-geeqie":                                "graphics",
	"x-gnome-networksettings":                 "system",
	"x-gnome-personalsettings":                "system",
	"x-gnome-settings-panel":                  "system",
	"x-gnome-systemsettings":                  "system",
	"x-gnustep":                               "system",
	"x-islamic-software":                      "reading",
	"x-jack":                                  "music",
	"x-kde-edu-misc":                          "reading",
	"x-kde-internet":                          "system",
	"x-kde-more":                              "system",
	"x-kde-utilities-desktop":                 "system",
	"x-kde-utilities-file":                    "system",
	"x-kde-utilities-peripherals":             "system",
	"x-kde-utilities-pim":                     "system",
	"x-lxde-settings":                         "system",
	"x-mandriva-office-publishing":            "others",
	"x-mandrivalinux-internet-other":          "system",
	"x-mandrivalinux-office-other":            "office",
	"x-mandrivalinux-system-archiving-backup": "system",
	"x-midi":                           "music",
	"x-misc":                           "system",
	"x-multitrack":                     "music",
	"x-novell-main":                    "system",
	"x-quran":                          "reading",
	"x-red-hat-base":                   "system",
	"x-red-hat-base-only":              "system",
	"x-red-hat-extra":                  "system",
	"x-red-hat-serverconfig":           "system",
	"x-religion":                       "reading",
	"x-sequencers":                     "music",
	"x-sound":                          "music",
	"x-sun-supported":                  "system",
	"x-suse-backup":                    "system",
	"x-suse-controlcenter-lookandfeel": "system",
	"x-suse-controlcenter-system":      "system",
	"x-suse-core":                      "system",
	"x-suse-core-game":                 "game",
	"x-suse-core-office":               "office",
	"x-suse-sequencer":                 "music",
	"x-suse-yast":                      "system",
	"x-suse-yast-high_availability":    "system",
	"x-synthesis":                      "system",
	"x-turbolinux-office":              "office",
	"x-xfce":                           "system",
	"x-xfce-toplevel":                  "system",
	"x-xfcesettingsdialog":             "system",
	"x-ximian-main":                    "system",
}
