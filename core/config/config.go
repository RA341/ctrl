package config

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Config struct {
	General      GeneralConfig
	Network      NetworkConfig
	Qbit         QbitConfig
	DiscordNotif DiscordNotifConfig
}

type GeneralConfig struct {
	AutoUpdate     bool
	UpdateInterval int
	EnableDocker   bool
}

type NetworkConfig struct {
	Host string
	Port int
}

type QbitConfig struct {
	Enable        bool
	Url           string
	User          string
	Pass          string
	SID           string
	ContainerName string
}

type DiscordNotifConfig struct {
	Enable     bool
	WebhookURL string
	Username   string
	AvatarURL  string
}

const defaultConfigFileName = "config.ini"

var config Config

func Get() *Config {
	return &config
}

func Load() {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to get current working directory")
	}
	configPath := filepath.Join(cwd, defaultConfigFileName)

	CreateDefaultConfigIfNotExists(configPath)

	iniCfg, err := ini.Load(configPath)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to load config file")
	}

	err = parseINI(iniCfg)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to parse config file")
	}

}

func parseINI(cfg *ini.File) error {
	// General section
	generalSection := cfg.Section("General")
	config.General = GeneralConfig{
		AutoUpdate:     generalSection.Key("auto_update").MustBool(true),
		UpdateInterval: generalSection.Key("update_interval").MustInt(168),
	}

	// Network section
	networkSection := cfg.Section("Network")
	config.Network = NetworkConfig{
		Host: networkSection.Key("Host").String(),
		Port: networkSection.Key("Port").MustInt(9220),
	}
	validateNetworkSection()

	// Qbit section
	qbitSection := cfg.Section("Qbit")

	config.Qbit = QbitConfig{
		Url:           createFullUrl(qbitSection.Key("host").MustString("NOHOST"), qbitSection.Key("port").MustInt(8085)),
		User:          qbitSection.Key("username").String(),
		Pass:          qbitSection.Key("password").String(),
		SID:           "",
		ContainerName: qbitSection.Key("container_name").String(),
	}
	validateQbitSection()

	// Discord notifications section
	discordSection := cfg.Section("notifications.Discord")
	config.DiscordNotif = DiscordNotifConfig{
		Enable:     discordSection.Key("enable").MustBool(true),
		WebhookURL: discordSection.Key("discord_webhook_url").String(),
		Username:   discordSection.Key("username").MustString("CTRL Bot"),
		AvatarURL:  discordSection.Key("avatar_url").String(),
	}
	validateDiscordSection()

	return nil
}

func createFullUrl(host string, port int) string {
	if strings.HasPrefix(host, "https") {
		// if host is https address do not add the port
		return host
	}
	return fmt.Sprintf("%s:%s", host, strconv.Itoa(port))
}

func CreateDefaultConfigIfNotExists(configPath string) {
	// Check if the file already exists
	if _, err := os.Stat(configPath); err == nil {
		// File exists, no need to create
		log.Print("Config file found")
		return
	} else if !os.IsNotExist(err) {
		// Some other error occurred
		log.Fatal().Err(err).Msgf("error checking config file")
	}

	// Create a new INI file
	cfg := ini.Empty()

	// [General] section
	secGeneral := createSection(cfg, "General")
	createKey(cfg, secGeneral, "auto_update", "true", "")
	createKey(cfg, secGeneral, "update_interval", "168", "how often to check for updates in hours (default is weekly)")
	createKey(cfg, secGeneral, "enable_docker", "false", "Enable support for controlling docker")

	// [Network] section
	secNetwork := createSection(cfg, "Network")
	secNetwork.Comment = "BE CAREFUL WHEN CHANGING THIS, IT MAY CAUSE THE SERVER TO BECOME INACCESSIBLE"
	createKey(cfg, secNetwork, "Host", "0.0.0.0", "Warning do not add 'http://' in front of the ips")
	createKey(cfg, secNetwork, "Port", "9220", "")

	// [Qbit] section
	secQbit := createSection(cfg, "Qbit")
	createKey(cfg, secQbit, "enable", "true", "")
	createKey(cfg, secQbit, "host", "http://127.0.0.1", "ip or hostname and port of your qbittorrent instance (remember to add https or http accordingly)")
	createKey(cfg, secQbit, "port", "8085", "")
	createKey(cfg, secQbit, "username", "", "")
	createKey(cfg, secQbit, "password", "", "Remember to surround the password with '\"' for eg \"password\" and doe not contain '#'")
	createKey(cfg, secQbit, "container_name", "qbittorrent", "Name of your qbit docker container")

	// [notifications.Discord] section
	secDiscord := createSection(cfg, "notifications.Discord")
	createKey(cfg, secDiscord, "enable", "true", "")
	createKey(cfg, secDiscord, "discord_webhook_url", "", "more info https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks")
	createKey(cfg, secDiscord, "username", "CTRL Bot", "")
	createKey(cfg, secDiscord, "avatar_url", "https://i.imgur.com/KEungv8.png", "")

	// Save the file
	if err := cfg.SaveTo(configPath); err != nil {
		log.Fatal().Err(err).Msg("failed to save default config file")
	}

	setPermissions(configPath)

	log.Info().Msgf("Created default config file at: %s", configPath)
	log.Info().
		Str("Then restart the program by running", "sudo systemctl stop ctrl.service").
		Msg("This is a first time run, check the config first")
}

func createSection(cfg *ini.File, section string) *ini.Section {
	createdSection, err := cfg.NewSection(section)
	if err != nil {
		log.Fatal().Err(err).Msgf("error creating section %s", section)
	}
	return createdSection
}

func createKey(cfg *ini.File, section *ini.Section, key string, value string, comment string) {
	createdKey, err := section.NewKey(key, value)
	if err != nil {
		log.Fatal().Err(err).Msgf("error creating key %s", key)
	}
	createdKey.Comment = comment
}

func setPermissions(filepath string) {
	uid, gid := 1000, 1000

	// Create file with current user's permissions
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Error().Err(err).Msg("failed to create file")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Error().Err(err).Msg("failed to close file")
		}
	}(file)

	// Change ownership of the file to the current user
	if err := file.Chown(uid, gid); err != nil {
		log.Error().Err(err).Msg("failed to change ownership of file")
	}

	log.Info().Msgf("Set permissions for %s", filepath)
}
