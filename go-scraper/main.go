package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// This is very ugly with the expectation that no matter what I do,
// I will probably have to re-write this next time I need it because
// the API will probably change

type entry struct {
	EndLocation    endLocation     `json:"endLocation"`
	StartLocations []startLocation `json:"startPositions"`
}
type endLocation struct {
	ID string `json:"id"`
}
type startLocation struct {
	ID string `json:"id"`
}
type mapInfo []entry

type videoInfo struct {
	TeamSide  string `json: "teamSide"`
	StartPos  string `json:"startPosition"`
	EndPos    string `json:"endPosition"`
	Type      string `json:"type"`
	YoutubeID string `json:"youTubeId"`
}

type resMapFile struct {
	T    []videoInfo `json:"T"`
	CT   []videoInfo `json:"CT"`
	Both []videoInfo `json:"Both"`
}

func main() {
	nades := []string{"smoke", "flash", "molotov", "hegrenade"}
	gameMap := os.Args[1]
	fmt.Printf("Map: %s", gameMap)
	mapResult := resMapFile{}
	for _, nadeType := range nades {
		gameMapInfo := mapInfo{}
		url := makeMainUrl(gameMap, nadeType)

		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("reqing: %w", err)
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&gameMapInfo)
		if err != nil {
			log.Fatalf("deser: %w", err)
		}
		for _, e := range gameMapInfo {
			for _, start := range e.StartLocations {
				time.Sleep(1 * time.Second)
				pairs := handlePair(start.ID, e.EndLocation.ID)
				fmt.Println(pairs)
				for _, pair := range pairs {
					if pair.TeamSide == "terrorist" {
						mapResult.T = append(mapResult.T, pair)
					} else if pair.TeamSide == "counterTerrorist" {
						mapResult.CT = append(mapResult.CT, pair)
					} else if pair.TeamSide == "both" {
						mapResult.Both = append(mapResult.Both, pair)
					} else {
						log.Fatal("unknown team side " + pair.TeamSide + " " + fmt.Sprint(pair))
					}
				}
			}
		}
	}
	output, err := json.MarshalIndent(mapResult, "\t", "\t")
	if err != nil {
		log.Fatal("outputting ", err)
	}
	err = os.WriteFile(fmt.Sprintf("%s.json", gameMap), output, 0o644)
	if err != nil {
		log.Fatal("writing ", err)
	}

}

func handlePair(start, end string) []videoInfo {
	url := makeSubUrl(start, end)
	fmt.Println(url, end, start)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("reqing: %w", err)
	}
	defer resp.Body.Close()

	vids := []videoInfo{}
	err = json.NewDecoder(resp.Body).Decode(&vids)
	if err != nil {
		log.Fatalf("deser: %w", err)
	}
	return vids
}

func makeMainUrl(gameMap string, nadeType string) string {
	// return fmt.Sprintf("https://api.csgonades.com/nademap/%s\?nadeType\=%s\&gameMode\=cs2\&tickRate\=any\&teamSide\=both\&favorites\=0", gameMap, nadeType)
	return fmt.Sprintf("https://api.csgonades.com/nademap/%s?nadeType=%s&gameMode=cs2&tickRate=any&teamSide=both&favorites=0", gameMap, nadeType)
}

func makeSubUrl(start, end string) string {
	return fmt.Sprintf("https://api.csgonades.com/nades/start/%s/end/%s", start, end)
}
