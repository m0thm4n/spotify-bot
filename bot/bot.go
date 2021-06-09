package bot

import (
	"github.com/bwmarrin/discordgo"
)

var ip string
var discordBot *discordgo.Session

// BotID is bot ID
var BotID string

// Start starts bot

/*func messageHandler(sess *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == BotID {
		return
	}
	if msg.Content == "ross" {
		_, _ = sess.ChannelMessageSend(msg.ChannelID, "Is a dumbass")
	}
}*/

/*func addScanHandler(sess *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Content == config.BotPrefix+"scan" {
		//fmt.Println(string(msg.Content[1]))
		stringMsg := msg.Content
		stringSplit := strings.Split(stringMsg, " ")
		fmt.Println(stringSplit)
		_, _ = sess.ChannelMessageSend(msg.ChannelID, "Starting async port scan now Chief")
		// TODO add user input to select target to scan
		ports := make(chan int, 100)
		results := make(chan int)
		var openports []int

		fmt.Println("Starting workers")
		for i := 0; i < cap(ports); i++ {
			go worker.Worker(ports, results)
		}

		go func() {
			for i := 1; i <= 1024; i++ {
				ports <- i
			}
		}()

		for i := 0; i < 1024; i++ {
			port := <-results
			if port != 0 {
				openports = append(openports, port)
			}
		}

		close(ports)
		close(results)
		fmt.Println("Sorting the ports")
		sort.Ints(openports)
		for _, port := range openports {
			fmt.Printf("%d open\n", port)
			_, _ = sess.ChannelMessageSend(msg.ChannelID, "Port: "+strconv.Itoa(port)+" is open")
		}
		_, _ = sess.ChannelMessageSend(msg.ChannelID, "Scan is complete")
	}
}
*/
