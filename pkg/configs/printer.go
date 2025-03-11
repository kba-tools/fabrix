package configs

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
	"github.com/spf13/viper"
)


func PrintNetworkInfo(domainName string) {
	filePath := fmt.Sprintf("./fabrix/%v/Network", domainName)

	// Set the file you want to read
	viper.SetConfigName("network_info")
	viper.SetConfigType("json")
	viper.AddConfigPath(filePath)

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return
	}

	// Fetch channel name
	channelName := viper.GetString("channelName")

	// Fetch orderer details
	orderMap := viper.Get("orderer").(map[string]interface{})
	ordererPeers := orderMap["peers"].([]interface{})

	ordererNode := tree.New().Root("Orderer")
	ca := orderMap["ca"].(map[string]interface{})
	caName := ca["name"].(string)
	caPort := ca["port"].(float64)

	ordererCA := tree.New().Root("CA")
	ordererNode.Child(ordererCA)

	ordererCA.Child(fmt.Sprintf("Name: %s", caName))
	ordererCA.Child(fmt.Sprintf("Port: %d", int(caPort)))

	ordererPeersTree := tree.New().Root("Peers")
	ordererNode.Child(ordererPeersTree)

	for _, ordererPeer := range ordererPeers {
		ordererPeerMap := ordererPeer.(map[string]interface{})
		ordererPeerName := ordererPeerMap["name"].(string)
		ordererPeerPort := ordererPeerMap["port"].(float64)

		ordererPeersTree.Child(fmt.Sprintf("Name: %s", ordererPeerName))
		ordererPeersTree.Child(fmt.Sprintf("Port: %d", int(ordererPeerPort)))

	}

	// Fetch organizations
	organizations := viper.Get("organisations").([]interface{})

	// Create a parent node for organizations
	orgNodes := tree.New().Root("Organizations")

	for _, org := range organizations {
		orgMap := org.(map[string]interface{})
		orgName := orgMap["name"].(string)
		orgMSP := orgMap["mspid"].(string)
		peers := orgMap["peers"].([]interface{})
		ca := orgMap["ca"].(map[string]interface{})

		caName := ca["name"].(string)
		caPort := ca["port"].(float64)

		currOrg := tree.New().Root(orgName)
		orgNodes.Child(currOrg)

		orgCA := tree.New().Root("CA")
		orgCA.Child(fmt.Sprintf("Name: %s", caName))
		orgCA.Child(fmt.Sprintf("Port: %d", int(caPort)))

		orgNode := currOrg.Child(fmt.Sprintf("MSP Id: %s", orgMSP))
		orgNode.Child(orgCA)
		orgPeers := tree.New().Root("Peers")
		orgNode.Child(orgPeers)

		for _, peer := range peers {
			peerMap := peer.(map[string]interface{})
			peerName := peerMap["name"].(string)
			peerPort := peerMap["port"].(float64)
			couchdbName := peerMap["couchdbname"].(string)
			couchdbPort := peerMap["couchdbport"].(float64)

			currPeer := tree.New().Root(peerName)
			orgPeers.Child(currPeer)

			currPeer.Child(fmt.Sprintf("Name : %s", peerName))
			currPeer.Child(fmt.Sprintf("Port : %d", int(peerPort)))
			currPeer.Child(fmt.Sprintf("Couchdb: %s", couchdbName))
			currPeer.Child(fmt.Sprintf("Couchdb Port : %d", int(couchdbPort)))

		}
	}

	// Styling for the tree
	enumeratorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#1c71e8")).MarginRight(1).Bold(true)
	rootStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#266ed4")).Bold(true)
	itemStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#7fe81c"))

	// Create the final tree with network info
	t := tree.
		Root(fmt.Sprintf("Domain: %s", domainName)).
		Child(fmt.Sprintf("Channel: %s", channelName)).
		Child(orgNodes).
		Child(ordererNode).
		Enumerator(tree.RoundedEnumerator).
		EnumeratorStyle(enumeratorStyle).
		RootStyle(rootStyle).
		ItemStyle(itemStyle)

	styledTree := lipgloss.NewStyle().
		MarginLeft(5).
		Render(t.String())

	fmt.Println(styledTree)

}

// Save NetworkInfo to a JSON file
func SaveNetworkInfoToFile(info *NetworkInfo, filePath string) error {
	// Convert the NetworkInfo struct to JSON
	data, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal network info: %v", err)
	}

	// Create or open the file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Write the JSON data to the file
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}
	return nil
}
