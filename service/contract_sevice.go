package service

import (
	"context"
	"fmt"
	"math/big"
	"github.com/charmbracelet/log"

	"github.com/emenesism/Decentralized-voting-backend/config"
	contract "github.com/emenesism/Decentralized-voting-backend/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var contractInstance *contract.Contract
var ethClient *ethclient.Client

func InitContractService() {
	var err error

	ethClient, err = ethclient.Dial(config.AppConfig.Rpc_url)

	if err != nil {
		log.Fatalf("Failed to connect to the eth node : %v", err)
	}

	contractAddress := common.HexToAddress(config.AppConfig.Contract_address)
	contractInstance, err = contract.NewContract(contractAddress, ethClient)

	if err != nil {
		log.Fatalf("Failed to load contract: %v", err)
	}

	log.Info("Contract connected successfully")

}

func GetVotes(candidate string) (*big.Int, error) {
	result, err := contractInstance.GetVotes(&bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	}, candidate)

	if err != nil {
		return nil, fmt.Errorf("Failed to get votes: %v", err)
	}

	return result, nil

}




// 🔹 Vote submits a vote for a candidate
func Vote(candidate string) (string, error) {
	privateKeyHex := config.AppConfig.Private_key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("invalid private key: %w", err)
	}

	// Get sender address
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// Get nonce
	nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %w", err)
	}

	// Get suggested gas price
	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get gas price: %w", err)
	}

	// Chain ID (Ganache uses 1337 or 5777 — try 1337 first)
	chainID := big.NewInt(1337)

	// Create transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %w", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // no ETH sent
	auth.GasLimit = uint64(300000)  // can adjust as needed
	auth.GasPrice = gasPrice

	tx, err := contractInstance.Vote(auth, candidate)
	if err != nil {
		return "", fmt.Errorf("failed to submit vote: %w", err)
	}

	return tx.Hash().Hex(), nil
}