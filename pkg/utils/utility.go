package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

func ClearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // Windows
	} else {
		cmd = exec.Command("clear") // Linux/macOS
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// func GetAllChaincodeDefenitions(domainName string) {
// 	filePath := fmt.Sprintf("./fabrix/%v/Network", domainName)

// 	// Set the file you want to read
// 	viper.SetConfigName("network_info")
// 	viper.SetConfigType("json")
// 	viper.AddConfigPath(filePath)

// 	// Read the config file
// 	if err := viper.ReadInConfig(); err != nil {
// 		fmt.Printf("Error reading config file, %s", err)
// 		return
// 	}

// 	path := filepath.Join("fabrix", domainName, "Network")

// 	channelName := viper.GetString("channelName")
// 	organizations := viper.Get("organisations").([]interface{})

// 	org := organizations[0]
// 	orgMap := org.(map[string]interface{})
// 	peers := orgMap["peers"].([]interface{})
// 	firstPeer := peers[0].(map[string]interface{})
// 	peerName := firstPeer["name"].(string)
// 	peerPort := uint16(firstPeer["port"].(float64))
// 	orderMap := viper.Get("orderer").(map[string]interface{})
// 	ordererPeers := orderMap["peers"].([]interface{})
// 	firstOrdererPeer := ordererPeers[0].(map[string]interface{})
// 	firstOrdererPeerName := firstOrdererPeer["name"].(string)
// 	ordererPort := uint16(firstOrdererPeer["port"].(float64))

// 	// Create the content of the script
// 	queryChaincodeDefinitions := fmt.Sprintf(`#!/bin/bash
// echo 'query Chaincode Definitions'
// export DOMAIN_NAME=%s
// export CHANNEL_NAME=%s
// export ORDERER_PEER=%s
// export ORDERER_PORT=%d
// export PEER_NAME=%s
// export PEER_PORT=%d

// export ORDERER_CA=${PWD}/organizations/ordererOrganizations/${DOMAIN_NAME}/orderers/orderer.${DOMAIN_NAME}/msp/tlscacerts/tlsca.${DOMAIN_NAME}-cert.pem
// export FABRIC_CFG_PATH=${PWD}/peercfg
// export CORE_PEER_TLS_ENABLED=true
// peer lifecycle chaincode querycommitted -o ${ORDERER_PEER}.${DOMAIN_NAME}:${ORDERER_PORT} --channelID ${CHANNEL_NAME} --tls --cafile $ORDERER_CA --peerAddresses ${PEER_NAME}:${PEER_PORT} --output json

// `, domainName, channelName, firstOrdererPeerName, ordererPort, peerName, peerPort)

// 	// Create a temporary file for the script
// 	tmpFile, err := os.CreateTemp("", "upgrade_chaincode_*.sh")
// 	if err != nil {
// 		fmt.Printf("Error creating temporary script file: %s\n", err)
// 		return
// 	}
// 	defer os.Remove(tmpFile.Name())

// 	// Write the initial content to the temp file
// 	_, err = tmpFile.Write([]byte(queryChaincodeDefinitions))
// 	if err != nil {
// 		fmt.Printf("Error writing initial content to temporary script file: %s\n", err)
// 		return
// 	}

// 	// Display file contents
// 	fmt.Println("Script file contents:")
// 	content, _ := os.ReadFile(tmpFile.Name())
// 	fmt.Println(string(content))

// 	// Make the script executable
// 	err = os.Chmod(tmpFile.Name(), 0755)
// 	if err != nil {
// 		fmt.Printf("Error making script executable: %s\n", err)
// 		return
// 	}

// 	// Run the script
// 	cmd := exec.Command("bash", tmpFile.Name())
// 	cmd.Dir = path
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	// cmd.Stdout = io.Discard
// 	// cmd.Stderr = io.Discard

// 	err = cmd.Run()
// 	if err != nil {
// 		fmt.Printf("Error executing script: %v\n", err)
// 		return
// 	}
// }

func GetAllChaincodeDefenitions(domainName string) string {

	filePath := fmt.Sprintf("./fabrix/%v/Network", domainName)

	// Set the file you want to read
	viper.SetConfigName("network_info")
	viper.SetConfigType("json")
	viper.AddConfigPath(filePath)

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return ""
	}

	path := filepath.Join("fabrix", domainName, "Network")

	channelName := viper.GetString("channelName")
	// Create the initial content of the script
	packageChaincode := fmt.Sprintf(`#!/bin/bash
echo 'Packaging Chaincode'
export DOMAIN_NAME=%s
export CHANNEL_NAME=%s


export ORDERER_CA=${PWD}/organizations/ordererOrganizations/${DOMAIN_NAME}/orderers/orderer.${DOMAIN_NAME}/msp/tlscacerts/tlsca.${DOMAIN_NAME}-cert.pem
export FABRIC_CFG_PATH=${PWD}/peercfg
export CORE_PEER_TLS_ENABLED=true
`, domainName, channelName)

	// Create a temporary file for the script
	tmpFile, err := os.CreateTemp("", "install_chaincode_*.sh")
	if err != nil {
		fmt.Printf("Error creating temporary script file: %s\n", err)
		return ""
	}
	defer os.Remove(tmpFile.Name())

	// Write the initial content to the temp file
	_, err = tmpFile.Write([]byte(packageChaincode))
	if err != nil {
		fmt.Printf("Error writing initial content to temporary script file: %s\n", err)
		return ""
	}

	// Close and reopen the temp file in append mode
	tmpFile.Close()
	tmpFile, err = os.OpenFile(tmpFile.Name(), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error reopening temporary script file for appending: %s\n", err)
		return ""
	}
	defer tmpFile.Close()

	organizations := viper.Get("organisations").([]interface{})
	peerStringList := []string{}

	for _, org := range organizations {
		orgMap := org.(map[string]interface{})

		orgName := orgMap["name"].(string)

		upperOrg := strings.ToUpper(orgName)

		orgMSP := orgMap["mspid"].(string)
		peers := orgMap["peers"].([]interface{})
		firstPeer := peers[0].(map[string]interface{})
		peerName := firstPeer["name"].(string)
		peerPort := uint16(firstPeer["port"].(float64))
		orderMap := viper.Get("orderer").(map[string]interface{})
		// ordererName := orderMap["name"].(string)
		ordererPeers := orderMap["peers"].([]interface{})
		firstOrdererPeer := ordererPeers[0].(map[string]interface{})
		ordererPort := uint16(firstOrdererPeer["port"].(float64))
		ordererName := firstOrdererPeer["name"].(string)

		// Append install and approve commands
		installApproveCommands := fmt.Sprintf(`export ORG_NAME_DOMAIN=%s.%s
	export ORDERER_CA=${PWD}/organizations/ordererOrganizations/${DOMAIN_NAME}/orderers/orderer.${DOMAIN_NAME}/msp/tlscacerts/tlsca.${DOMAIN_NAME}-cert.pem
	export ORG_NAME=%s
	export PEER=%s
	export PEER_PORT=%d
	export ORG_CAP=%s
	export ORDERER_PORT=%d
	export CORE_PEER_LOCALMSPID=%s
	export ORDERER_NAME=%s
	export CORE_PEER_ADDRESS=localhost:${PEER_PORT}
	export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/${ORG_NAME_DOMAIN}/peers/${PEER}/tls/ca.crt
	export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/${ORG_NAME_DOMAIN}/users/Admin@${ORG_NAME_DOMAIN}/msp
	export ${ORG_CAP}_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/${ORG_NAME_DOMAIN}/peers/${PEER}/tls/ca.crt

	
	`, orgName, domainName, orgName, peerName, peerPort, upperOrg, ordererPort, orgMSP, ordererName)

		_, err = tmpFile.WriteString(installApproveCommands)
		if err != nil {
			fmt.Printf("Error writing additional commands to temporary script file: %s\n", err)
			return ""
		}
		// for committing chaincode, need peer details
		peerString := fmt.Sprintf("--peerAddresses localhost:%d --tlsRootCertFiles $%s_PEER_TLSROOTCERT", peerPort, upperOrg)
		peerStringList = append(peerStringList, peerString)
	}
	peerStrings := strings.Join(peerStringList, "  ")

	// Append commit commands
	commitCommands := fmt.Sprintf(`
peer lifecycle chaincode querycommitted -o ${ORDERER_NAME}.${DOMAIN_NAME}:7050 --channelID $CHANNEL_NAME --tls --cafile $ORDERER_CA %s --output json

	`, peerStrings)

	_, err = tmpFile.WriteString(commitCommands)
	if err != nil {
		fmt.Printf("Error writing additional commands to temporary script file: %s\n", err)
		return ""
	}

	// // Display file contents
	// fmt.Println("Script file contents:")
	// content, _ := os.ReadFile(tmpFile.Name())
	// fmt.Println(string(content))

	// Make the script executable
	err = os.Chmod(tmpFile.Name(), 0755)
	if err != nil {
		fmt.Printf("Error making script executable: %s\n", err)
		return ""
	}

	cmd := exec.Command("bash", tmpFile.Name())
	cmd.Dir = path

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing script: %v\n", err)
		return ""
	}

	scriptOutput := string(output)
	return scriptOutput
}
