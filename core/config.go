package core

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

type GameInfo struct {
	Path    string `json:"path"`
	Version string `json:"version"`
}

type Mod struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Author  string `json:"author"`
	Enabled bool   `json:"enabled"`
}

type Modlist struct {
	Mods []Mod `json:"mods"`
}

type Config struct {
	IsCurrent  bool     `json:"iscurrent"`
	IsDefault  bool     `json:"isdefault"`
	Version    string   `json:"version"`
	ConfigName string   `json:"configname"`
	Game       GameInfo `json:"game"`
	CfgMods    Modlist  `json:"cfgmods"`
}

func (c *Config) Prep() {
	fmt.Println(`Found following configs:`)
	var cfgList []string = SuffixFind(CFG_DIR, ".conf")

	log.Print(cfgList)

	for i, file := range cfgList {
		file = strings.TrimSpace(file)
		if file == "" {
			continue
		}
		fmt.Printf("Found Config: #%v. %s\n", i+1, file)
	}

	fmt.Println(`Select your config file, or hit ENTER to make a new one.`)
	var uInput string
	fmt.Scanln(&uInput)

	switch uInput {
	case "":
		fmt.Println(`No input, making new config.`)
		c.Create()
		return
	default:
		c.Load(uInput)
	}
}

func (c *Config) Load(uInput string) {

	/* Open config */
	file, err := os.OpenFile(filepath.Join(CFG_DIR, uInput+".conf"), os.O_RDWR, 0666)
	if err != nil {
		log.Fatal().Err(err).Msgf("Error opening config %s", uInput)
	}
	defer file.Close()

	/* Read file */
	jData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal().Err(err).Msg("Error reading config file.")
	}

	/* Load saved config to mem config */
	err = json.Unmarshal(jData, c)
	if err != nil {
		log.Fatal().Err(err).Msg(`Failed to read Json data in config file.`)
	}

	log.Info().Msgf("Loaded config from %s config file.", uInput)
}

func (c *Config) Create() {
	fmt.Println(`Please select a name for your config file:`)
	var userInput string
	fmt.Scanln(&userInput)

	c.IsCurrent = true
	c.IsDefault = false
	c.Version = "not-set"
	c.Game.Path = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Cyberpunk 2077"
	c.Game.Version = "latest"
	c.ConfigName = userInput

	os.MkdirAll(CFG_DIR, os.ModePerm)

	jData, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		log.Fatal().Err(err).Msg(`Error marshaling Json for Config.`)
		return
	}

	err = os.WriteFile(filepath.Join(CFG_DIR, userInput+".conf"), jData, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("Error writing config file.")
		return
	}

	log.Info().Msgf("Config written to %s.conf", userInput)
}
