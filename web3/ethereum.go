package connections

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	GameToken "github.com/gameon-app-inc/laliga-matchfantasy-api/bindings/GameToken"
	LaLiga "github.com/gameon-app-inc/laliga-matchfantasy-api/bindings/Players"
	WarChest "github.com/gameon-app-inc/laliga-matchfantasy-api/bindings/WarChest"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

var (
	rpc              = os.Getenv("ETHEREUM_URL")
	warChestAddress  = common.HexToAddress(os.Getenv("WAR_CHEST_ADDRESS"))
	gameTokenAddress = common.HexToAddress(os.Getenv("GAME_TOKEN_ADDRESS"))
	laLigaAddress    = common.HexToAddress(os.Getenv("LALIGA_ADDRESS"))
	client           *ethclient.Client
)

func init() {
	var err error
	client, err = ethclient.Dial(rpc)
	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v", err)
	}
}

func NewWarChestToken() (*WarChest.WarChest, error) {
	return WarChest.NewWarChest(warChestAddress, client)
}

func NewGameToken() (*GameToken.GAME, error) {
	return GameToken.NewGAME(gameTokenAddress, client)
}

func NewLaLigaToken() (*LaLiga.LaLiga, error) {
	return LaLiga.NewLaLiga(laLigaAddress, client)
}

func MintLaLiga(toAddress string, tokenId *big.Int) (*types.Transaction, error) {
	auth, err := getTransactor()
	if err != nil {
		return nil, err
	}
	if toAddress == "" {
		return nil, fmt.Errorf("toAddress is required")
	}
	if tokenId == nil {
		return nil, fmt.Errorf("tokenId is required")
	}

	laLigaTokenInstance, err := NewLaLigaToken()
	if err != nil {
		return nil, fmt.Errorf("failed to create LaLiga instance: %v", err)
	}
	to := common.HexToAddress(toAddress)

	tx, err := laLigaTokenInstance.SafeMint(auth, to, tokenId)

	if err != nil {
		return nil, fmt.Errorf("SafeMint failed: %v", err)
	}

	receipt, err := waitForTransactionReceipt(client, tx.Hash())
	if err != nil {
		return tx, fmt.Errorf("error while waiting for transaction receipt: %v", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return tx, fmt.Errorf("transaction failed to execute successfully, status: %v", receipt.Status)
	}
	return tx, nil
}

func BatchMintLaLiga(toAddress string, tokenIds []*big.Int) (*types.Transaction, error) {
	auth, err := getTransactor()
	if err != nil {
		return nil, err
	}
	if toAddress == "" {
		return nil, fmt.Errorf("toAddress is required")
	}
	if len(tokenIds) == 0 {
		return nil, fmt.Errorf("tokenIds are required")
	}

	laLigaTokenInstance, err := NewLaLigaToken()
	if err != nil {
		return nil, fmt.Errorf("failed to create LaLiga instance: %v", err)
	}
	to := common.HexToAddress(toAddress)

	tx, err := laLigaTokenInstance.BatchMint(auth, to, tokenIds)
	if err != nil {
		return nil, fmt.Errorf("BatchMint failed: %v", err)
	}

	receipt, err := waitForTransactionReceipt(client, tx.Hash())
	if err != nil {
		return tx, fmt.Errorf("error while waiting for transaction receipt: %v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return tx, fmt.Errorf("transaction failed to execute successfully, status: %v", receipt.Status)
	}
	return tx, nil
}

func GetBalanceOf(address string, contractType string) (*big.Int, error) {
	switch contractType {
	case "WarChest":
		warChestInstance, err := NewWarChestToken()
		if err != nil {
			return nil, fmt.Errorf("failed to create new war chest token: %v", err)
		}
		account := common.HexToAddress(address)
		return warChestInstance.BalanceOf(nil, account)

	case "GameToken":
		gameTokenInstance, err := NewGameToken()
		if err != nil {
			return nil, fmt.Errorf("failed to create new game token: %v", err)
		}
		account := common.HexToAddress(address)
		return gameTokenInstance.BalanceOf(nil, account)

	case "LaLigaToken":
		laLigaTokenInstance, err := NewLaLigaToken()
		if err != nil {
			return nil, fmt.Errorf("failed to create new LaLiga token: %v", err)
		}
		account := common.HexToAddress(address)
		return laLigaTokenInstance.BalanceOf(nil, account)

	default:
		return nil, fmt.Errorf("unknown contract type: %s", contractType)
	}
}

func ApproveToken(toAddress string, tokenId *big.Int, contractType string) (*types.Transaction, error) {
	auth, err := getTransactor()
	if err != nil {
		return nil, err
	}
	var tx *types.Transaction

	switch contractType {

	case "WarChest":
		warChestInstance, err := NewWarChestToken()
		if err != nil {
			return nil, fmt.Errorf("failed to create new WarChest instance: %v", err)
		}
		tx, err = warChestInstance.Approve(auth, common.HexToAddress(toAddress), tokenId)
		if err != nil {
			return nil, err
		}

	case "GameToken":
		gameTokenInstance, err := NewGameToken()
		if err != nil {
			return nil, fmt.Errorf("failed to create new GameToken instance: %v", err)
		}
		tx, err = gameTokenInstance.Approve(auth, common.HexToAddress(toAddress), tokenId)
		if err != nil {
			return nil, err
		}

	case "LaLigaToken":
		laLigaTokenInstance, err := NewLaLigaToken()
		if err != nil {
			return nil, fmt.Errorf("failed to create new LaLiga instance: %v", err)
		}
		tx, err = laLigaTokenInstance.Approve(auth, common.HexToAddress(toAddress), tokenId)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("unknown contract type: %s", contractType)
	}

	receipt, err := waitForTransactionReceipt(client, tx.Hash())
	if err != nil {
		return tx, fmt.Errorf("error while waiting for transaction receipt: %v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return tx, fmt.Errorf("transaction failed to execute successfully, status: %v", receipt.Status)
	}

	return tx, nil
}

func SetApprovalForAll(operator string, approved bool, contractType string) (*types.Transaction, error) {
	auth, err := getTransactor()
	if err != nil {
		return nil, err
	}
	var tx *types.Transaction

	switch contractType {

	case "WarChest":
		instance, err := NewWarChestToken()
		if err != nil {
			return nil, fmt.Errorf("failed to create new WarChest instance: %v", err)
		}
		tx, err = instance.SetApprovalForAll(auth, common.HexToAddress(operator), approved)
		if err != nil {
			return nil, err
		}

	case "LaLigaToken":
		instance, err := NewLaLigaToken()
		if err != nil {
			return nil, fmt.Errorf("failed to create new LaLiga instance: %v", err)
		}
		tx, err = instance.SetApprovalForAll(auth, common.HexToAddress(operator), approved)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("unknown contract type: %s", contractType)
	}

	// Wait for the transaction to be mined
	receipt, err := waitForTransactionReceipt(client, tx.Hash())
	if err != nil {
		return tx, fmt.Errorf("error while waiting for transaction receipt: %v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return tx, fmt.Errorf("transaction failed to execute successfully, status: %v", receipt.Status)
	}

	return tx, nil
}

func TransferToken(fromAddress, toAddress string, tokenId *big.Int, contractType string) (*types.Transaction, error) {
	auth, err := getTransactor()
	if err != nil {
		return nil, err
	}
	var tx *types.Transaction

	switch contractType {

	case "WarChest":
		instance, err := NewWarChestToken()
		if err != nil {
			return nil, fmt.Errorf("failed to create new WarChest instance: %v", err)
		}
		tx, err = instance.TransferFrom(auth, common.HexToAddress(fromAddress), common.HexToAddress(toAddress), tokenId)
		if err != nil {
			return nil, err
		}

	case "GameToken":
		instance, err := NewGameToken()
		if err != nil {
			return nil, fmt.Errorf("failed to create new GameToken instance: %v", err)
		}
		tx, err = instance.TransferFrom(auth, common.HexToAddress(fromAddress), common.HexToAddress(toAddress), tokenId)
		if err != nil {
			return nil, err
		}

	case "LaLigaToken":
		instance, err := NewLaLigaToken()
		if err != nil {
			return nil, fmt.Errorf("failed to create new LaLiga instance: %v", err)
		}
		tx, err = instance.TransferFrom(auth, common.HexToAddress(fromAddress), common.HexToAddress(toAddress), tokenId)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("unknown contract type: %s", contractType)
	}

	// Wait for the transaction to be mined
	receipt, err := waitForTransactionReceipt(client, tx.Hash())
	if err != nil {
		return tx, fmt.Errorf("error while waiting for transaction receipt: %v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return tx, fmt.Errorf("transaction failed to execute successfully, status: %v", receipt.Status)
	}

	return tx, nil
}

func BurnToken(tokenId *big.Int, contractType string) (*types.Transaction, error) {
	auth, err := getTransactor()
	if err != nil {
		return nil, err
	}
	var tx *types.Transaction

	switch contractType {

	case "GameToken":
		instance, err := NewGameToken()
		if err != nil {
			return nil, fmt.Errorf("failed to create new GameToken instance: %v", err)
		}
		tx, err = instance.Burn(auth, tokenId)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("unknown contract type: %s", contractType)
	}

	receipt, err := waitForTransactionReceipt(client, tx.Hash())
	if err != nil {
		return tx, fmt.Errorf("error while waiting for transaction receipt: %v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return tx, fmt.Errorf("transaction failed to execute successfully, status: %v", receipt.Status)
	}

	return tx, nil
}

func GetTokenURI(tokenId *big.Int, contractType string) (string, error) {
	switch contractType {
	case "WarChest":
		instance, err := NewWarChestToken()
		if err != nil {
			return "", fmt.Errorf("failed to create new war chest token: %v", err)
		}
		return instance.TokenURI(nil, tokenId)

	case "LaLigaToken":
		instance, err := NewLaLigaToken()
		if err != nil {
			return "", fmt.Errorf("failed to create new LaLiga token: %v", err)
		}
		return instance.TokenURI(nil, tokenId)

	default:
		return "", fmt.Errorf("unknown contract type: %s", contractType)
	}
}

func OwnerOfToken(tokenId *big.Int, contractType string) (common.Address, error) {
	switch contractType {
	case "WarChest":
		instance, err := NewWarChestToken()
		if err != nil {
			return common.Address{}, fmt.Errorf("failed to create new war chest token: %v", err)
		}
		return instance.OwnerOf(nil, tokenId)

	case "LaLigaToken":
		instance, err := NewLaLigaToken()
		if err != nil {
			return common.Address{}, fmt.Errorf("failed to create new LaLiga token: %v", err)
		}
		return instance.OwnerOf(nil, tokenId)

	default:
		return common.Address{}, fmt.Errorf("unknown contract type: %s", contractType)
	}
}

func IsApprovedForAll(ownerAddress, operatorAddress string, contractType string) (bool, error) {
	owner := common.HexToAddress(ownerAddress)
	operator := common.HexToAddress(operatorAddress)

	switch contractType {
	case "WarChest":
		instance, err := NewWarChestToken()
		if err != nil {
			return false, err
		}
		return instance.IsApprovedForAll(nil, owner, operator)

	case "LaLigaToken":
		instance, err := NewLaLigaToken()
		if err != nil {
			return false, err
		}
		return instance.IsApprovedForAll(nil, owner, operator)

	default:
		return false, fmt.Errorf("unknown contract type: %s", contractType)
	}
}

func GetOwner(contractType string) (common.Address, error) {
	switch contractType {

	case "WarChest":
		instance, err := NewWarChestToken()
		if err != nil {
			return common.Address{}, err
		}
		return instance.Owner(nil)

	case "LaLigaToken":
		instance, err := NewLaLigaToken()
		if err != nil {
			return common.Address{}, err
		}
		return instance.Owner(nil)

	default:
		return common.Address{}, fmt.Errorf("unknown contract type: %s", contractType)
	}
}

func GetApproved(tokenId *big.Int, contractType string) (common.Address, error) {
	switch contractType {

	case "WarChest":
		instance, err := NewWarChestToken()
		if err != nil {
			return common.Address{}, err
		}
		return instance.GetApproved(nil, tokenId)

	case "LaLigaToken":
		instance, err := NewLaLigaToken()
		if err != nil {
			return common.Address{}, err
		}
		return instance.GetApproved(nil, tokenId)

	default:
		return common.Address{}, fmt.Errorf("unknown contract type: %s", contractType)
	}
}

func SafeTransferFrom(fromAddress, toAddress string, tokenId *big.Int, contractType string) (*types.Transaction, error) {
	auth, err := getTransactor()
	if err != nil {
		return nil, err
	}
	var tx *types.Transaction

	switch contractType {
	case "WarChest":
		instance, err := NewWarChestToken()
		if err != nil {
			return nil, err
		}
		tx, err = instance.SafeTransferFrom(auth, common.HexToAddress(fromAddress), common.HexToAddress(toAddress), tokenId)
		if err != nil {
			return nil, err
		}

	case "LaLigaToken":
		instance, err := NewLaLigaToken()
		if err != nil {
			return nil, err
		}
		tx, err = instance.SafeTransferFrom(auth, common.HexToAddress(fromAddress), common.HexToAddress(toAddress), tokenId)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("unknown contract type: %s", contractType)
	}

	receipt, err := waitForTransactionReceipt(client, tx.Hash())
	if err != nil {
		return tx, fmt.Errorf("error while waiting for transaction receipt: %v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return tx, fmt.Errorf("transaction failed to execute successfully, status: %v", receipt.Status)
	}

	return tx, nil
}

func SendSignedTransaction(signedTxHex string) (*types.Transaction, error) {
	signedTxBytes, err := hex.DecodeString(signedTxHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode transaction: %v", err)
	}

	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(signedTxBytes, tx); err != nil {
		return nil, fmt.Errorf("failed to decode transaction: %v", err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %v", err)
	}

	// Wait for the transaction to be mined
	receipt, err := waitForTransactionReceipt(client, tx.Hash())
	if err != nil {
		return tx, fmt.Errorf("error while waiting for transaction receipt: %v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return tx, fmt.Errorf("transaction failed to execute successfully, status: %v", receipt.Status)
	}

	return tx, nil
}

func getTransactor() (*bind.TransactOpts, error) {
	walletKey := os.Getenv("WALLET_KEY")
	privateKey, err := crypto.HexToECDSA(walletKey)
	if err != nil {
		return nil, err
	}

	client, err := ethclient.Dial(os.Getenv("ETHEREUM_URL"))
	if err != nil {
		return nil, err
	}

	gasTip, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, err
	}

	head, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	baseFee := head.BaseFee
	feeCap := new(big.Int).Add(baseFee, new(big.Int).Mul(big.NewInt(2), gasTip))

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(getNonce(privateKey)))
	auth.GasFeeCap = feeCap
	auth.GasTipCap = gasTip
	auth.GasLimit = 5000000

	return auth, nil
}

func getNonce(privatekey *ecdsa.PrivateKey) uint64 {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privatekey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)

	return nonce
}

func waitForTransactionReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	queryTicker := time.NewTicker(5 * time.Second)
	defer queryTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("transaction receipt retrieval timed out")
		case <-queryTicker.C:
			receipt, err := client.TransactionReceipt(ctx, txHash)
			if err != nil {
				if receipt == nil {
					continue
				}
				return nil, fmt.Errorf("failed to retrieve transaction receipt: %v", err)
			}
			if receipt == nil {
				continue
			}
			return receipt, nil
		}
	}
}

// // Improved waitForTransactionReceipt with specific context timeout and error handling
// func waitForTransactionReceiptEnhanced(txHash common.Hash) (*types.Receipt, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
// 	defer cancel()

// 	queryTicker := time.NewTicker(5 * time.Second)
// 	defer queryTicker.Stop()

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return nil, ctx.Err() // Use ctx.Err() to return the specific error, e.g., DeadlineExceeded
// 		case <-queryTicker.C:
// 			receipt, err := client.TransactionReceipt(ctx, txHash)
// 			if err != nil {
// 				if receipt == nil {
// 					continue // Continue polling until a receipt is found or context deadline is exceeded
// 				}
// 				// Return more specific error info if possible
// 				return nil, fmt.Errorf("failed to retrieve transaction receipt: %v", err)
// 			}
// 			return receipt, nil
// 		}
// 	}
// }

// func getTransactor2() (*bind.TransactOpts, error) {
// 	walletKey := os.Getenv("WALLET_KEY")
// 	privateKey, err := crypto.HexToECDSA(walletKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	client, err := ethclient.Dial(os.Getenv("ETHEREUM_URL"))
// 	if err != nil {
// 		return nil, err
// 	}

// 	gasTip, err := client.SuggestGasTipCap(context.Background())
// 	if err != nil {
// 		return nil, err
// 	}

// 	head, err := client.HeaderByNumber(context.Background(), nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	baseFee := head.BaseFee
// 	feeCap := new(big.Int).Add(baseFee, new(big.Int).Mul(big.NewInt(2), gasTip))

// 	chainID, err := client.NetworkID(context.Background())
// 	if err != nil {
// 		return nil, err
// 	}

// 	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
// 	if err != nil {
// 		return nil, err
// 	}
// 	auth.Nonce = big.NewInt(int64(nonce))
// 	auth.GasFeeCap = feeCap
// 	auth.GasTipCap = gasTip
// 	auth.GasLimit = 15000000 // Example gas limit, adjust as needed

// 	return auth, nil
// }
