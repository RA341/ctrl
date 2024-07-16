package updater

import (
	"context"
	"github.com/google/go-github/v63/github"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func UpdateBinary() {
	client := github.NewClient(nil)

	ctx := context.Background()

	owner := "RA341"
	repo := "ctrl-srv"

	tag := compareTags(client, ctx, owner, repo)

	if tag != "" {
		log.Info().Msgf("New version detected: %s", tag)
		installBinary(client, ctx, owner, repo, tag)
	}
}

func installBinary(client *github.Client, ctx context.Context, owner string, repo string, tag string) {
	release, _, err := client.Repositories.GetReleaseByTag(ctx, owner, repo, tag)
	if err != nil {
		log.Error().Err(err).Msg("Error getting release by tag")
	}

	assets, _, err := client.Repositories.ListReleaseAssets(ctx, owner, repo, release.GetID(), nil)
	if err != nil {
		log.Error().Err(err).Msgf("Error listing assets")
		return
	}

	for _, asset := range assets {
		rel := strings.Split(*asset.Name, "_")
		// get binary for the os
		if len(rel) >= 2 && rel[1] == runtime.GOOS {
			log.Info().Msgf("Binary found: %s", *asset.Name)
			downloadAndInstallBinary(*asset.Name, *asset.BrowserDownloadURL)
			return
		} else {
			log.Info().Msgf("Skipping asset: %s, incomaptible with os: %s", *asset.Name, runtime.GOOS)
		}
	}
	log.Debug().Msg("No compatible binary found")
}

func downloadAndInstallBinary(assetName string, url string) {
	// Download the asset
	resp, err := http.Get(url)
	if err != nil {
		log.Error().Err(err).Msg("Error downloading binary")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal().Err(err).Msg("Error closing body")
		}
	}(resp.Body)

	// Create the file
	inputFile, err := os.Create(assetName)
	if err != nil {
		log.Error().Err(err).Msgf("Error creating file: %v", err)
		return
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Fatal().Err(err).Msg("Error closing body")
		}
	}(inputFile)

	_, err = io.Copy(inputFile, resp.Body)
	if err != nil {
		log.Error().Err(err).Msgf("Error writing to donwloaded file")
	}

	err = os.Chmod(assetName, 0755)
	if err != nil {
		log.Error().Err(err).Msg("Error making file executable")
	}

	// copy and replace the binary and the remove the binary
	executablePath := os.Args[0]

	outputFile, err := os.OpenFile(executablePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Error().Err(err).Msgf("Error opening executable")
		return
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			log.Fatal().Err(err).Msg("Error closing body")
		}
	}(outputFile)

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		log.Error().Err(err).Msgf("Error writing to current binary")
		return
	}

	err = os.Remove(assetName)
	if err != nil {
		log.Error().Err(err).Msgf("Error removing downloaded file")
		return
	}
}

func compareTags(client *github.Client, ctx context.Context, owner string, repo string) string {
	tags, status, err := client.Repositories.ListTags(ctx, owner, repo, nil)

	if err != nil {
		log.Error().Err(err).Msg("Failed to get the repository")
		return ""
	}

	if status.Status != "200 OK" {
		log.Error().Msgf("Failed to make http request, status: %s", status.Status)
		return ""
	}

	if len(tags) > 0 {
		if Version == "development" {
			log.Warn().Msgf("Development version detected. Install a production version to recieve updates")
			return ""
		}

		latestTag := tags[0]
		log.Info().Msgf("Latest tag: %s, binary tag: %s", *latestTag.Name, Version)

		binVer := convertTagsToNum(Version)
		remoteVer := convertTagsToNum(*latestTag.Name)

		if remoteVer > binVer {
			return *latestTag.Name
		} else if remoteVer == binVer {
			log.Info().Msg("Binary is already at latest version")
			return ""
		}
	}

	log.Info().Msg("No tags found")
	return ""
}

func convertTagsToNum(tag string) int {
	// remove v and dots
	replacer := strings.NewReplacer(".", "", "v", "")
	tag = replacer.Replace(tag)
	log.Debug().Msgf("Removed '.' and 'v': %s", tag)

	res, err := strconv.Atoi(tag)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert tag to int")
		return -1
	}
	return res
}
