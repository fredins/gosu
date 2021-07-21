package main

var (
	delim     = []byte("   ") // the delimiter that will be used
	shell     = "sh"        // shell used
	cmdstropt = "-c"        // command string opt for shell

	blocks = []block{
		{cmd: "clock", upSig: 1, upInt: 60, color:"^c#eceff4^"},
		// {cmd: "volume", upSig: 10, upInt: 3, color:"^c#a3be8c^"}, need to remove pulseaudio deamon (10M)
		{cmd: "weather", upInt: 18000, upSig: 5, color:"^c#5e81ac^"},
		// {cmd: "recicon", upSig: 9},
		{cmd: "music", upSig: 11},
		// {cmd: "pacpackages", upSig: 8},
		// {cmd: "news", upSig: 6},
		// {cmd: "georona | cut -d' ' -f1,3", inSh: true, upInt: 18000, upSig: 19}, // example of command that is run in shell
		// {cmd: "torrent", upInt: 10, upSig: 7},
		{cmd: "memory", upInt: 3, upSig: 14, color:"^c#ebcb8b^"},
		{cmd: "cpu", upInt: 3, upSig: 13, color:"^c#88c0d0^"},
		// {cmd: "cpubars", upInt: 1, upSig: 22},
		{cmd: "disk sda3", upSig: 15, color:"^c#b48ead^"},
		// {cmd: "astrological", upInt: 18000, upSig: 18},
		// {cmd: "mailbox", upSig: 12},
		// {cmd: "battery", upInt: 5, upSig: 10},
		{cmd: "nettraf", upInt: 5, upSig: 16, color:"^c#81a1c1^"},
		{cmd: "internet", upInt: 5, upSig: 4, color:"^c#eceff4^"},
	}
)
