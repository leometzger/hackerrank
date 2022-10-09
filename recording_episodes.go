package main

// https://www.hackerrank.com/challenges/episode-recording/problem?isFullScreen=true&h_r=next-challenge&h_v=zen

type Schedule = map[int][]*Episode

type Episode struct {
	Ep    int
	Start int32
	End   int32
	Nexts []*Episode
}

func RecordingEpisodes(input [][]int32) []int32 {
	schedule := buildSchedule(input)
	var start int = 0
	var numberOfEpisodes int = 0

	for epNumber, episodes := range schedule {
		for _, play := range episodes {
			max := getMaxUse(play)

			if max > numberOfEpisodes {
				numberOfEpisodes = max
				start = epNumber
			}
		}
	}

	return []int32{int32(start + 1), int32(numberOfEpisodes)}
}

func getMaxUse(ep *Episode) int {
	if len(ep.Nexts) == 0 {
		return 1
	}

	max := 0

	for _, play := range ep.Nexts {
		playMax := getMaxUse(play)

		if max < playMax {
			max = playMax
		}
	}

	return 1 + max
}

func buildSchedule(input [][]int32) Schedule {
	episodes := make(map[int][]*Episode, len(input))

	for i := 0; i < len(input); i++ {
		live := Episode{
			Ep:    i,
			Start: input[i][0],
			End:   input[i][1],
			Nexts: make([]*Episode, 0),
		}

		repeat := Episode{
			Ep:    i,
			Start: input[i][2],
			End:   input[i][3],
			Nexts: make([]*Episode, 0),
		}

		eps := []*Episode{&live, &repeat}
		episodes[i] = eps

		if i > 0 {
			if episodes[i-1][0].End < live.Start {
				episodes[i-1][0].Nexts = append(episodes[i-1][0].Nexts, &live)
			}

			if episodes[i-1][0].End < repeat.Start {
				episodes[i-1][0].Nexts = append(episodes[i-1][0].Nexts, &repeat)
			}

			if episodes[i-1][1].End < live.Start {
				episodes[i-1][1].Nexts = append(episodes[i-1][1].Nexts, &live)
			}

			if episodes[i-1][1].End < repeat.Start {
				episodes[i-1][1].Nexts = append(episodes[i-1][1].Nexts, &repeat)
			}
		}
	}

	return episodes
}
