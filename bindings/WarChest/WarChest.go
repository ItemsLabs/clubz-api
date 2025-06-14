// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package WarChest

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AllowListData is an auto generated low-level Go binding around an user-defined struct.
type AllowListData struct {
	MerkleRoot    [32]byte
	PublicKeyURIs []string
	AllowListURI  string
}

// ERC721SeaDropStructsErrorsAndEventsMultiConfigureStruct is an auto generated low-level Go binding around an user-defined struct.
type ERC721SeaDropStructsErrorsAndEventsMultiConfigureStruct struct {
	MaxSupply                            *big.Int
	BaseURI                              string
	ContractURI                          string
	SeaDropImpl                          common.Address
	PublicDrop                           PublicDrop
	DropURI                              string
	AllowListData                        AllowListData
	CreatorPayoutAddress                 common.Address
	ProvenanceHash                       [32]byte
	AllowedFeeRecipients                 []common.Address
	DisallowedFeeRecipients              []common.Address
	AllowedPayers                        []common.Address
	DisallowedPayers                     []common.Address
	TokenGatedAllowedNftTokens           []common.Address
	TokenGatedDropStages                 []TokenGatedDropStage
	DisallowedTokenGatedAllowedNftTokens []common.Address
	Signers                              []common.Address
	SignedMintValidationParams           []SignedMintValidationParams
	DisallowedSigners                    []common.Address
}

// ISeaDropTokenContractMetadataRoyaltyInfo is an auto generated low-level Go binding around an user-defined struct.
type ISeaDropTokenContractMetadataRoyaltyInfo struct {
	RoyaltyAddress common.Address
	RoyaltyBps     *big.Int
}

// PublicDrop is an auto generated low-level Go binding around an user-defined struct.
type PublicDrop struct {
	MintPrice                *big.Int
	StartTime                *big.Int
	EndTime                  *big.Int
	MaxTotalMintableByWallet uint16
	FeeBps                   uint16
	RestrictFeeRecipients    bool
}

// SignedMintValidationParams is an auto generated low-level Go binding around an user-defined struct.
type SignedMintValidationParams struct {
	MinMintPrice                *big.Int
	MaxMaxTotalMintableByWallet *big.Int
	MinStartTime                *big.Int
	MaxEndTime                  *big.Int
	MaxMaxTokenSupplyForStage   *big.Int
	MinFeeBps                   uint16
	MaxFeeBps                   uint16
}

// TokenGatedDropStage is an auto generated low-level Go binding around an user-defined struct.
type TokenGatedDropStage struct {
	MintPrice                *big.Int
	MaxTotalMintableByWallet uint16
	StartTime                *big.Int
	EndTime                  *big.Int
	DropStageIndex           uint8
	MaxTokenSupplyForStage   uint32
	FeeBps                   uint16
	RestrictFeeRecipients    bool
}

// WarChestMetaData contains all meta data concerning the WarChest contract.
var WarChestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"CannotExceedMaxSupplyOfUint64\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"basisPoints\",\"type\":\"uint256\"}],\"name\":\"InvalidRoyaltyBasisPoints\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotNextOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedSeaDrop\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProvenanceHashCannotBeSetAfterMintStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoyaltyAddressCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignersMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowedSeaDrop\",\"type\":\"address[]\"}],\"name\":\"AllowedSeaDropUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newContractURI\",\"type\":\"string\"}],\"name\":\"ContractURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"MaxSupplyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPotentialAdministrator\",\"type\":\"address\"}],\"name\":\"PotentialOwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newHash\",\"type\":\"bytes32\"}],\"name\":\"ProvenanceHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"name\":\"RoyaltyInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SeaDropTokenDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"}],\"name\":\"emitBatchMetadataUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"getMintStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minterNumMinted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"__name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"__symbol\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"allowedSeaDrop\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"mintSeaDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contractURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structPublicDrop\",\"name\":\"publicDrop\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"dropURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"publicKeyURIs\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"allowListURI\",\"type\":\"string\"}],\"internalType\":\"structAllowListData\",\"name\":\"allowListData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"creatorPayoutAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"provenanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"allowedFeeRecipients\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedFeeRecipients\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"allowedPayers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedPayers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenGatedAllowedNftTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structTokenGatedDropStage[]\",\"name\":\"tokenGatedDropStages\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedTokenGatedAllowedNftTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"internalType\":\"structSignedMintValidationParams[]\",\"name\":\"signedMintValidationParams\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedSigners\",\"type\":\"address[]\"}],\"internalType\":\"structERC721SeaDropStructsErrorsAndEvents.MultiConfigureStruct\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"multiConfigure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provenanceHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newBaseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newContractURI\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newProvenanceHash\",\"type\":\"bytes32\"}],\"name\":\"setProvenanceHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"royaltyAddress\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"royaltyBps\",\"type\":\"uint96\"}],\"internalType\":\"structISeaDropTokenContractMetadata.RoyaltyInfo\",\"name\":\"newInfo\",\"type\":\"tuple\"}],\"name\":\"setRoyaltyInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPotentialOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"publicKeyURIs\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"allowListURI\",\"type\":\"string\"}],\"internalType\":\"structAllowListData\",\"name\":\"allowListData\",\"type\":\"tuple\"}],\"name\":\"updateAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateAllowedFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"allowedSeaDrop\",\"type\":\"address[]\"}],\"name\":\"updateAllowedSeaDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payoutAddress\",\"type\":\"address\"}],\"name\":\"updateCreatorPayoutAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"dropURI\",\"type\":\"string\"}],\"name\":\"updateDropURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updatePayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structPublicDrop\",\"name\":\"publicDrop\",\"type\":\"tuple\"}],\"name\":\"updatePublicDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"internalType\":\"structSignedMintValidationParams\",\"name\":\"signedMintValidationParams\",\"type\":\"tuple\"}],\"name\":\"updateSignedMintValidationParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structTokenGatedDropStage\",\"name\":\"dropStage\",\"type\":\"tuple\"}],\"name\":\"updateTokenGatedDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WarChestABI is the input ABI used to generate the binding from.
// Deprecated: Use WarChestMetaData.ABI instead.
var WarChestABI = WarChestMetaData.ABI

// WarChest is an auto generated Go binding around an Ethereum contract.
type WarChest struct {
	WarChestCaller     // Read-only binding to the contract
	WarChestTransactor // Write-only binding to the contract
	WarChestFilterer   // Log filterer for contract events
}

// WarChestCaller is an auto generated read-only Go binding around an Ethereum contract.
type WarChestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WarChestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WarChestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WarChestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WarChestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WarChestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WarChestSession struct {
	Contract     *WarChest         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WarChestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WarChestCallerSession struct {
	Contract *WarChestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// WarChestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WarChestTransactorSession struct {
	Contract     *WarChestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WarChestRaw is an auto generated low-level Go binding around an Ethereum contract.
type WarChestRaw struct {
	Contract *WarChest // Generic contract binding to access the raw methods on
}

// WarChestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WarChestCallerRaw struct {
	Contract *WarChestCaller // Generic read-only contract binding to access the raw methods on
}

// WarChestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WarChestTransactorRaw struct {
	Contract *WarChestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWarChest creates a new instance of WarChest, bound to a specific deployed contract.
func NewWarChest(address common.Address, backend bind.ContractBackend) (*WarChest, error) {
	contract, err := bindWarChest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WarChest{WarChestCaller: WarChestCaller{contract: contract}, WarChestTransactor: WarChestTransactor{contract: contract}, WarChestFilterer: WarChestFilterer{contract: contract}}, nil
}

// NewWarChestCaller creates a new read-only instance of WarChest, bound to a specific deployed contract.
func NewWarChestCaller(address common.Address, caller bind.ContractCaller) (*WarChestCaller, error) {
	contract, err := bindWarChest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WarChestCaller{contract: contract}, nil
}

// NewWarChestTransactor creates a new write-only instance of WarChest, bound to a specific deployed contract.
func NewWarChestTransactor(address common.Address, transactor bind.ContractTransactor) (*WarChestTransactor, error) {
	contract, err := bindWarChest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WarChestTransactor{contract: contract}, nil
}

// NewWarChestFilterer creates a new log filterer instance of WarChest, bound to a specific deployed contract.
func NewWarChestFilterer(address common.Address, filterer bind.ContractFilterer) (*WarChestFilterer, error) {
	contract, err := bindWarChest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WarChestFilterer{contract: contract}, nil
}

// bindWarChest binds a generic wrapper to an already deployed contract.
func bindWarChest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WarChestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WarChest *WarChestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WarChest.Contract.WarChestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WarChest *WarChestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WarChest.Contract.WarChestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WarChest *WarChestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WarChest.Contract.WarChestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WarChest *WarChestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WarChest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WarChest *WarChestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WarChest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WarChest *WarChestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WarChest.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_WarChest *WarChestCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WarChest.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_WarChest *WarChestSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _WarChest.Contract.BalanceOf(&_WarChest.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_WarChest *WarChestCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _WarChest.Contract.BalanceOf(&_WarChest.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_WarChest *WarChestCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WarChest.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_WarChest *WarChestSession) BaseURI() (string, error) {
	return _WarChest.Contract.BaseURI(&_WarChest.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_WarChest *WarChestCallerSession) BaseURI() (string, error) {
	return _WarChest.Contract.BaseURI(&_WarChest.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_WarChest *WarChestCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WarChest.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_WarChest *WarChestSession) ContractURI() (string, error) {
	return _WarChest.Contract.ContractURI(&_WarChest.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_WarChest *WarChestCallerSession) ContractURI() (string, error) {
	return _WarChest.Contract.ContractURI(&_WarChest.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_WarChest *WarChestCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _WarChest.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_WarChest *WarChestSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _WarChest.Contract.GetApproved(&_WarChest.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_WarChest *WarChestCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _WarChest.Contract.GetApproved(&_WarChest.CallOpts, tokenId)
}

// GetMintStats is a free data retrieval call binding the contract method 0x840e15d4.
//
// Solidity: function getMintStats(address minter) view returns(uint256 minterNumMinted, uint256 currentTotalSupply, uint256 maxSupply)
func (_WarChest *WarChestCaller) GetMintStats(opts *bind.CallOpts, minter common.Address) (struct {
	MinterNumMinted    *big.Int
	CurrentTotalSupply *big.Int
	MaxSupply          *big.Int
}, error) {
	var out []interface{}
	err := _WarChest.contract.Call(opts, &out, "getMintStats", minter)

	outstruct := new(struct {
		MinterNumMinted    *big.Int
		CurrentTotalSupply *big.Int
		MaxSupply          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinterNumMinted = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentTotalSupply = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaxSupply = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetMintStats is a free data retrieval call binding the contract method 0x840e15d4.
//
// Solidity: function getMintStats(address minter) view returns(uint256 minterNumMinted, uint256 currentTotalSupply, uint256 maxSupply)
func (_WarChest *WarChestSession) GetMintStats(minter common.Address) (struct {
	MinterNumMinted    *big.Int
	CurrentTotalSupply *big.Int
	MaxSupply          *big.Int
}, error) {
	return _WarChest.Contract.GetMintStats(&_WarChest.CallOpts, minter)
}

// GetMintStats is a free data retrieval call binding the contract method 0x840e15d4.
//
// Solidity: function getMintStats(address minter) view returns(uint256 minterNumMinted, uint256 currentTotalSupply, uint256 maxSupply)
func (_WarChest *WarChestCallerSession) GetMintStats(minter common.Address) (struct {
	MinterNumMinted    *big.Int
	CurrentTotalSupply *big.Int
	MaxSupply          *big.Int
}, error) {
	return _WarChest.Contract.GetMintStats(&_WarChest.CallOpts, minter)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_WarChest *WarChestCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _WarChest.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_WarChest *WarChestSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _WarChest.Contract.IsApprovedForAll(&_WarChest.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_WarChest *WarChestCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _WarChest.Contract.IsApprovedForAll(&_WarChest.CallOpts, owner, operator)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_WarChest *WarChestCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WarChest.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_WarChest *WarChestSession) MaxSupply() (*big.Int, error) {
	return _WarChest.Contract.MaxSupply(&_WarChest.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_WarChest *WarChestCallerSession) MaxSupply() (*big.Int, error) {
	return _WarChest.Contract.MaxSupply(&_WarChest.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WarChest *WarChestCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WarChest.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WarChest *WarChestSession) Name() (string, error) {
	return _WarChest.Contract.Name(&_WarChest.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WarChest *WarChestCallerSession) Name() (string, error) {
	return _WarChest.Contract.Name(&_WarChest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WarChest *WarChestCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WarChest.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WarChest *WarChestSession) Owner() (common.Address, error) {
	return _WarChest.Contract.Owner(&_WarChest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WarChest *WarChestCallerSession) Owner() (common.Address, error) {
	return _WarChest.Contract.Owner(&_WarChest.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_WarChest *WarChestCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _WarChest.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_WarChest *WarChestSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _WarChest.Contract.OwnerOf(&_WarChest.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_WarChest *WarChestCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _WarChest.Contract.OwnerOf(&_WarChest.CallOpts, tokenId)
}

// ProvenanceHash is a free data retrieval call binding the contract method 0xc6ab67a3.
//
// Solidity: function provenanceHash() view returns(bytes32)
func (_WarChest *WarChestCaller) ProvenanceHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WarChest.contract.Call(opts, &out, "provenanceHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvenanceHash is a free data retrieval call binding the contract method 0xc6ab67a3.
//
// Solidity: function provenanceHash() view returns(bytes32)
func (_WarChest *WarChestSession) ProvenanceHash() ([32]byte, error) {
	return _WarChest.Contract.ProvenanceHash(&_WarChest.CallOpts)
}

// ProvenanceHash is a free data retrieval call binding the contract method 0xc6ab67a3.
//
// Solidity: function provenanceHash() view returns(bytes32)
func (_WarChest *WarChestCallerSession) ProvenanceHash() ([32]byte, error) {
	return _WarChest.Contract.ProvenanceHash(&_WarChest.CallOpts)
}

// RoyaltyAddress is a free data retrieval call binding the contract method 0xad2f852a.
//
// Solidity: function royaltyAddress() view returns(address)
func (_WarChest *WarChestCaller) RoyaltyAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WarChest.contract.Call(opts, &out, "royaltyAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyAddress is a free data retrieval call binding the contract method 0xad2f852a.
//
// Solidity: function royaltyAddress() view returns(address)
func (_WarChest *WarChestSession) RoyaltyAddress() (common.Address, error) {
	return _WarChest.Contract.RoyaltyAddress(&_WarChest.CallOpts)
}

// RoyaltyAddress is a free data retrieval call binding the contract method 0xad2f852a.
//
// Solidity: function royaltyAddress() view returns(address)
func (_WarChest *WarChestCallerSession) RoyaltyAddress() (common.Address, error) {
	return _WarChest.Contract.RoyaltyAddress(&_WarChest.CallOpts)
}

// RoyaltyBasisPoints is a free data retrieval call binding the contract method 0x42260b5d.
//
// Solidity: function royaltyBasisPoints() view returns(uint256)
func (_WarChest *WarChestCaller) RoyaltyBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WarChest.contract.Call(opts, &out, "royaltyBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoyaltyBasisPoints is a free data retrieval call binding the contract method 0x42260b5d.
//
// Solidity: function royaltyBasisPoints() view returns(uint256)
func (_WarChest *WarChestSession) RoyaltyBasisPoints() (*big.Int, error) {
	return _WarChest.Contract.RoyaltyBasisPoints(&_WarChest.CallOpts)
}

// RoyaltyBasisPoints is a free data retrieval call binding the contract method 0x42260b5d.
//
// Solidity: function royaltyBasisPoints() view returns(uint256)
func (_WarChest *WarChestCallerSession) RoyaltyBasisPoints() (*big.Int, error) {
	return _WarChest.Contract.RoyaltyBasisPoints(&_WarChest.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_WarChest *WarChestCaller) RoyaltyInfo(opts *bind.CallOpts, arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _WarChest.contract.Call(opts, &out, "royaltyInfo", arg0, _salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_WarChest *WarChestSession) RoyaltyInfo(arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _WarChest.Contract.RoyaltyInfo(&_WarChest.CallOpts, arg0, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_WarChest *WarChestCallerSession) RoyaltyInfo(arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _WarChest.Contract.RoyaltyInfo(&_WarChest.CallOpts, arg0, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WarChest *WarChestCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _WarChest.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WarChest *WarChestSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WarChest.Contract.SupportsInterface(&_WarChest.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WarChest *WarChestCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WarChest.Contract.SupportsInterface(&_WarChest.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WarChest *WarChestCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WarChest.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WarChest *WarChestSession) Symbol() (string, error) {
	return _WarChest.Contract.Symbol(&_WarChest.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WarChest *WarChestCallerSession) Symbol() (string, error) {
	return _WarChest.Contract.Symbol(&_WarChest.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_WarChest *WarChestCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _WarChest.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_WarChest *WarChestSession) TokenURI(tokenId *big.Int) (string, error) {
	return _WarChest.Contract.TokenURI(&_WarChest.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_WarChest *WarChestCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _WarChest.Contract.TokenURI(&_WarChest.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WarChest *WarChestCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WarChest.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WarChest *WarChestSession) TotalSupply() (*big.Int, error) {
	return _WarChest.Contract.TotalSupply(&_WarChest.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WarChest *WarChestCallerSession) TotalSupply() (*big.Int, error) {
	return _WarChest.Contract.TotalSupply(&_WarChest.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_WarChest *WarChestTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_WarChest *WarChestSession) AcceptOwnership() (*types.Transaction, error) {
	return _WarChest.Contract.AcceptOwnership(&_WarChest.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_WarChest *WarChestTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _WarChest.Contract.AcceptOwnership(&_WarChest.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_WarChest *WarChestTransactor) Approve(opts *bind.TransactOpts, operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "approve", operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_WarChest *WarChestSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WarChest.Contract.Approve(&_WarChest.TransactOpts, operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_WarChest *WarChestTransactorSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WarChest.Contract.Approve(&_WarChest.TransactOpts, operator, tokenId)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_WarChest *WarChestTransactor) CancelOwnershipTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "cancelOwnershipTransfer")
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_WarChest *WarChestSession) CancelOwnershipTransfer() (*types.Transaction, error) {
	return _WarChest.Contract.CancelOwnershipTransfer(&_WarChest.TransactOpts)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_WarChest *WarChestTransactorSession) CancelOwnershipTransfer() (*types.Transaction, error) {
	return _WarChest.Contract.CancelOwnershipTransfer(&_WarChest.TransactOpts)
}

// EmitBatchMetadataUpdate is a paid mutator transaction binding the contract method 0xa4830114.
//
// Solidity: function emitBatchMetadataUpdate(uint256 fromTokenId, uint256 toTokenId) returns()
func (_WarChest *WarChestTransactor) EmitBatchMetadataUpdate(opts *bind.TransactOpts, fromTokenId *big.Int, toTokenId *big.Int) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "emitBatchMetadataUpdate", fromTokenId, toTokenId)
}

// EmitBatchMetadataUpdate is a paid mutator transaction binding the contract method 0xa4830114.
//
// Solidity: function emitBatchMetadataUpdate(uint256 fromTokenId, uint256 toTokenId) returns()
func (_WarChest *WarChestSession) EmitBatchMetadataUpdate(fromTokenId *big.Int, toTokenId *big.Int) (*types.Transaction, error) {
	return _WarChest.Contract.EmitBatchMetadataUpdate(&_WarChest.TransactOpts, fromTokenId, toTokenId)
}

// EmitBatchMetadataUpdate is a paid mutator transaction binding the contract method 0xa4830114.
//
// Solidity: function emitBatchMetadataUpdate(uint256 fromTokenId, uint256 toTokenId) returns()
func (_WarChest *WarChestTransactorSession) EmitBatchMetadataUpdate(fromTokenId *big.Int, toTokenId *big.Int) (*types.Transaction, error) {
	return _WarChest.Contract.EmitBatchMetadataUpdate(&_WarChest.TransactOpts, fromTokenId, toTokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x481a48ec.
//
// Solidity: function initialize(string __name, string __symbol, address[] allowedSeaDrop, address initialOwner) returns()
func (_WarChest *WarChestTransactor) Initialize(opts *bind.TransactOpts, __name string, __symbol string, allowedSeaDrop []common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "initialize", __name, __symbol, allowedSeaDrop, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x481a48ec.
//
// Solidity: function initialize(string __name, string __symbol, address[] allowedSeaDrop, address initialOwner) returns()
func (_WarChest *WarChestSession) Initialize(__name string, __symbol string, allowedSeaDrop []common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _WarChest.Contract.Initialize(&_WarChest.TransactOpts, __name, __symbol, allowedSeaDrop, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x481a48ec.
//
// Solidity: function initialize(string __name, string __symbol, address[] allowedSeaDrop, address initialOwner) returns()
func (_WarChest *WarChestTransactorSession) Initialize(__name string, __symbol string, allowedSeaDrop []common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _WarChest.Contract.Initialize(&_WarChest.TransactOpts, __name, __symbol, allowedSeaDrop, initialOwner)
}

// MintSeaDrop is a paid mutator transaction binding the contract method 0x64869dad.
//
// Solidity: function mintSeaDrop(address minter, uint256 quantity) returns()
func (_WarChest *WarChestTransactor) MintSeaDrop(opts *bind.TransactOpts, minter common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "mintSeaDrop", minter, quantity)
}

// MintSeaDrop is a paid mutator transaction binding the contract method 0x64869dad.
//
// Solidity: function mintSeaDrop(address minter, uint256 quantity) returns()
func (_WarChest *WarChestSession) MintSeaDrop(minter common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _WarChest.Contract.MintSeaDrop(&_WarChest.TransactOpts, minter, quantity)
}

// MintSeaDrop is a paid mutator transaction binding the contract method 0x64869dad.
//
// Solidity: function mintSeaDrop(address minter, uint256 quantity) returns()
func (_WarChest *WarChestTransactorSession) MintSeaDrop(minter common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _WarChest.Contract.MintSeaDrop(&_WarChest.TransactOpts, minter, quantity)
}

// MultiConfigure is a paid mutator transaction binding the contract method 0x911f456b.
//
// Solidity: function multiConfigure((uint256,string,string,address,(uint80,uint48,uint48,uint16,uint16,bool),string,(bytes32,string[],string),address,bytes32,address[],address[],address[],address[],address[],(uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool)[],address[],address[],(uint80,uint24,uint40,uint40,uint40,uint16,uint16)[],address[]) config) returns()
func (_WarChest *WarChestTransactor) MultiConfigure(opts *bind.TransactOpts, config ERC721SeaDropStructsErrorsAndEventsMultiConfigureStruct) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "multiConfigure", config)
}

// MultiConfigure is a paid mutator transaction binding the contract method 0x911f456b.
//
// Solidity: function multiConfigure((uint256,string,string,address,(uint80,uint48,uint48,uint16,uint16,bool),string,(bytes32,string[],string),address,bytes32,address[],address[],address[],address[],address[],(uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool)[],address[],address[],(uint80,uint24,uint40,uint40,uint40,uint16,uint16)[],address[]) config) returns()
func (_WarChest *WarChestSession) MultiConfigure(config ERC721SeaDropStructsErrorsAndEventsMultiConfigureStruct) (*types.Transaction, error) {
	return _WarChest.Contract.MultiConfigure(&_WarChest.TransactOpts, config)
}

// MultiConfigure is a paid mutator transaction binding the contract method 0x911f456b.
//
// Solidity: function multiConfigure((uint256,string,string,address,(uint80,uint48,uint48,uint16,uint16,bool),string,(bytes32,string[],string),address,bytes32,address[],address[],address[],address[],address[],(uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool)[],address[],address[],(uint80,uint24,uint40,uint40,uint40,uint16,uint16)[],address[]) config) returns()
func (_WarChest *WarChestTransactorSession) MultiConfigure(config ERC721SeaDropStructsErrorsAndEventsMultiConfigureStruct) (*types.Transaction, error) {
	return _WarChest.Contract.MultiConfigure(&_WarChest.TransactOpts, config)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WarChest *WarChestTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WarChest *WarChestSession) RenounceOwnership() (*types.Transaction, error) {
	return _WarChest.Contract.RenounceOwnership(&_WarChest.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WarChest *WarChestTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WarChest.Contract.RenounceOwnership(&_WarChest.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_WarChest *WarChestTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_WarChest *WarChestSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WarChest.Contract.SafeTransferFrom(&_WarChest.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_WarChest *WarChestTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WarChest.Contract.SafeTransferFrom(&_WarChest.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_WarChest *WarChestTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_WarChest *WarChestSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _WarChest.Contract.SafeTransferFrom0(&_WarChest.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_WarChest *WarChestTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _WarChest.Contract.SafeTransferFrom0(&_WarChest.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_WarChest *WarChestTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_WarChest *WarChestSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _WarChest.Contract.SetApprovalForAll(&_WarChest.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_WarChest *WarChestTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _WarChest.Contract.SetApprovalForAll(&_WarChest.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_WarChest *WarChestTransactor) SetBaseURI(opts *bind.TransactOpts, newBaseURI string) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "setBaseURI", newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_WarChest *WarChestSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _WarChest.Contract.SetBaseURI(&_WarChest.TransactOpts, newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_WarChest *WarChestTransactorSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _WarChest.Contract.SetBaseURI(&_WarChest.TransactOpts, newBaseURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newContractURI) returns()
func (_WarChest *WarChestTransactor) SetContractURI(opts *bind.TransactOpts, newContractURI string) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "setContractURI", newContractURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newContractURI) returns()
func (_WarChest *WarChestSession) SetContractURI(newContractURI string) (*types.Transaction, error) {
	return _WarChest.Contract.SetContractURI(&_WarChest.TransactOpts, newContractURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newContractURI) returns()
func (_WarChest *WarChestTransactorSession) SetContractURI(newContractURI string) (*types.Transaction, error) {
	return _WarChest.Contract.SetContractURI(&_WarChest.TransactOpts, newContractURI)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 newMaxSupply) returns()
func (_WarChest *WarChestTransactor) SetMaxSupply(opts *bind.TransactOpts, newMaxSupply *big.Int) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "setMaxSupply", newMaxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 newMaxSupply) returns()
func (_WarChest *WarChestSession) SetMaxSupply(newMaxSupply *big.Int) (*types.Transaction, error) {
	return _WarChest.Contract.SetMaxSupply(&_WarChest.TransactOpts, newMaxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 newMaxSupply) returns()
func (_WarChest *WarChestTransactorSession) SetMaxSupply(newMaxSupply *big.Int) (*types.Transaction, error) {
	return _WarChest.Contract.SetMaxSupply(&_WarChest.TransactOpts, newMaxSupply)
}

// SetProvenanceHash is a paid mutator transaction binding the contract method 0x099b6bfa.
//
// Solidity: function setProvenanceHash(bytes32 newProvenanceHash) returns()
func (_WarChest *WarChestTransactor) SetProvenanceHash(opts *bind.TransactOpts, newProvenanceHash [32]byte) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "setProvenanceHash", newProvenanceHash)
}

// SetProvenanceHash is a paid mutator transaction binding the contract method 0x099b6bfa.
//
// Solidity: function setProvenanceHash(bytes32 newProvenanceHash) returns()
func (_WarChest *WarChestSession) SetProvenanceHash(newProvenanceHash [32]byte) (*types.Transaction, error) {
	return _WarChest.Contract.SetProvenanceHash(&_WarChest.TransactOpts, newProvenanceHash)
}

// SetProvenanceHash is a paid mutator transaction binding the contract method 0x099b6bfa.
//
// Solidity: function setProvenanceHash(bytes32 newProvenanceHash) returns()
func (_WarChest *WarChestTransactorSession) SetProvenanceHash(newProvenanceHash [32]byte) (*types.Transaction, error) {
	return _WarChest.Contract.SetProvenanceHash(&_WarChest.TransactOpts, newProvenanceHash)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_WarChest *WarChestTransactor) SetRoyaltyInfo(opts *bind.TransactOpts, newInfo ISeaDropTokenContractMetadataRoyaltyInfo) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "setRoyaltyInfo", newInfo)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_WarChest *WarChestSession) SetRoyaltyInfo(newInfo ISeaDropTokenContractMetadataRoyaltyInfo) (*types.Transaction, error) {
	return _WarChest.Contract.SetRoyaltyInfo(&_WarChest.TransactOpts, newInfo)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_WarChest *WarChestTransactorSession) SetRoyaltyInfo(newInfo ISeaDropTokenContractMetadataRoyaltyInfo) (*types.Transaction, error) {
	return _WarChest.Contract.SetRoyaltyInfo(&_WarChest.TransactOpts, newInfo)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_WarChest *WarChestTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_WarChest *WarChestSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WarChest.Contract.TransferFrom(&_WarChest.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_WarChest *WarChestTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WarChest.Contract.TransferFrom(&_WarChest.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newPotentialOwner) returns()
func (_WarChest *WarChestTransactor) TransferOwnership(opts *bind.TransactOpts, newPotentialOwner common.Address) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "transferOwnership", newPotentialOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newPotentialOwner) returns()
func (_WarChest *WarChestSession) TransferOwnership(newPotentialOwner common.Address) (*types.Transaction, error) {
	return _WarChest.Contract.TransferOwnership(&_WarChest.TransactOpts, newPotentialOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newPotentialOwner) returns()
func (_WarChest *WarChestTransactorSession) TransferOwnership(newPotentialOwner common.Address) (*types.Transaction, error) {
	return _WarChest.Contract.TransferOwnership(&_WarChest.TransactOpts, newPotentialOwner)
}

// UpdateAllowList is a paid mutator transaction binding the contract method 0x3680620d.
//
// Solidity: function updateAllowList(address seaDropImpl, (bytes32,string[],string) allowListData) returns()
func (_WarChest *WarChestTransactor) UpdateAllowList(opts *bind.TransactOpts, seaDropImpl common.Address, allowListData AllowListData) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "updateAllowList", seaDropImpl, allowListData)
}

// UpdateAllowList is a paid mutator transaction binding the contract method 0x3680620d.
//
// Solidity: function updateAllowList(address seaDropImpl, (bytes32,string[],string) allowListData) returns()
func (_WarChest *WarChestSession) UpdateAllowList(seaDropImpl common.Address, allowListData AllowListData) (*types.Transaction, error) {
	return _WarChest.Contract.UpdateAllowList(&_WarChest.TransactOpts, seaDropImpl, allowListData)
}

// UpdateAllowList is a paid mutator transaction binding the contract method 0x3680620d.
//
// Solidity: function updateAllowList(address seaDropImpl, (bytes32,string[],string) allowListData) returns()
func (_WarChest *WarChestTransactorSession) UpdateAllowList(seaDropImpl common.Address, allowListData AllowListData) (*types.Transaction, error) {
	return _WarChest.Contract.UpdateAllowList(&_WarChest.TransactOpts, seaDropImpl, allowListData)
}

// UpdateAllowedFeeRecipient is a paid mutator transaction binding the contract method 0x48a4c101.
//
// Solidity: function updateAllowedFeeRecipient(address seaDropImpl, address feeRecipient, bool allowed) returns()
func (_WarChest *WarChestTransactor) UpdateAllowedFeeRecipient(opts *bind.TransactOpts, seaDropImpl common.Address, feeRecipient common.Address, allowed bool) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "updateAllowedFeeRecipient", seaDropImpl, feeRecipient, allowed)
}

// UpdateAllowedFeeRecipient is a paid mutator transaction binding the contract method 0x48a4c101.
//
// Solidity: function updateAllowedFeeRecipient(address seaDropImpl, address feeRecipient, bool allowed) returns()
func (_WarChest *WarChestSession) UpdateAllowedFeeRecipient(seaDropImpl common.Address, feeRecipient common.Address, allowed bool) (*types.Transaction, error) {
	return _WarChest.Contract.UpdateAllowedFeeRecipient(&_WarChest.TransactOpts, seaDropImpl, feeRecipient, allowed)
}

// UpdateAllowedFeeRecipient is a paid mutator transaction binding the contract method 0x48a4c101.
//
// Solidity: function updateAllowedFeeRecipient(address seaDropImpl, address feeRecipient, bool allowed) returns()
func (_WarChest *WarChestTransactorSession) UpdateAllowedFeeRecipient(seaDropImpl common.Address, feeRecipient common.Address, allowed bool) (*types.Transaction, error) {
	return _WarChest.Contract.UpdateAllowedFeeRecipient(&_WarChest.TransactOpts, seaDropImpl, feeRecipient, allowed)
}

// UpdateAllowedSeaDrop is a paid mutator transaction binding the contract method 0x60c308b6.
//
// Solidity: function updateAllowedSeaDrop(address[] allowedSeaDrop) returns()
func (_WarChest *WarChestTransactor) UpdateAllowedSeaDrop(opts *bind.TransactOpts, allowedSeaDrop []common.Address) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "updateAllowedSeaDrop", allowedSeaDrop)
}

// UpdateAllowedSeaDrop is a paid mutator transaction binding the contract method 0x60c308b6.
//
// Solidity: function updateAllowedSeaDrop(address[] allowedSeaDrop) returns()
func (_WarChest *WarChestSession) UpdateAllowedSeaDrop(allowedSeaDrop []common.Address) (*types.Transaction, error) {
	return _WarChest.Contract.UpdateAllowedSeaDrop(&_WarChest.TransactOpts, allowedSeaDrop)
}

// UpdateAllowedSeaDrop is a paid mutator transaction binding the contract method 0x60c308b6.
//
// Solidity: function updateAllowedSeaDrop(address[] allowedSeaDrop) returns()
func (_WarChest *WarChestTransactorSession) UpdateAllowedSeaDrop(allowedSeaDrop []common.Address) (*types.Transaction, error) {
	return _WarChest.Contract.UpdateAllowedSeaDrop(&_WarChest.TransactOpts, allowedSeaDrop)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x66251b69.
//
// Solidity: function updateCreatorPayoutAddress(address seaDropImpl, address payoutAddress) returns()
func (_WarChest *WarChestTransactor) UpdateCreatorPayoutAddress(opts *bind.TransactOpts, seaDropImpl common.Address, payoutAddress common.Address) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "updateCreatorPayoutAddress", seaDropImpl, payoutAddress)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x66251b69.
//
// Solidity: function updateCreatorPayoutAddress(address seaDropImpl, address payoutAddress) returns()
func (_WarChest *WarChestSession) UpdateCreatorPayoutAddress(seaDropImpl common.Address, payoutAddress common.Address) (*types.Transaction, error) {
	return _WarChest.Contract.UpdateCreatorPayoutAddress(&_WarChest.TransactOpts, seaDropImpl, payoutAddress)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x66251b69.
//
// Solidity: function updateCreatorPayoutAddress(address seaDropImpl, address payoutAddress) returns()
func (_WarChest *WarChestTransactorSession) UpdateCreatorPayoutAddress(seaDropImpl common.Address, payoutAddress common.Address) (*types.Transaction, error) {
	return _WarChest.Contract.UpdateCreatorPayoutAddress(&_WarChest.TransactOpts, seaDropImpl, payoutAddress)
}

// UpdateDropURI is a paid mutator transaction binding the contract method 0x7a05bc82.
//
// Solidity: function updateDropURI(address seaDropImpl, string dropURI) returns()
func (_WarChest *WarChestTransactor) UpdateDropURI(opts *bind.TransactOpts, seaDropImpl common.Address, dropURI string) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "updateDropURI", seaDropImpl, dropURI)
}

// UpdateDropURI is a paid mutator transaction binding the contract method 0x7a05bc82.
//
// Solidity: function updateDropURI(address seaDropImpl, string dropURI) returns()
func (_WarChest *WarChestSession) UpdateDropURI(seaDropImpl common.Address, dropURI string) (*types.Transaction, error) {
	return _WarChest.Contract.UpdateDropURI(&_WarChest.TransactOpts, seaDropImpl, dropURI)
}

// UpdateDropURI is a paid mutator transaction binding the contract method 0x7a05bc82.
//
// Solidity: function updateDropURI(address seaDropImpl, string dropURI) returns()
func (_WarChest *WarChestTransactorSession) UpdateDropURI(seaDropImpl common.Address, dropURI string) (*types.Transaction, error) {
	return _WarChest.Contract.UpdateDropURI(&_WarChest.TransactOpts, seaDropImpl, dropURI)
}

// UpdatePayer is a paid mutator transaction binding the contract method 0xcb743ba8.
//
// Solidity: function updatePayer(address seaDropImpl, address payer, bool allowed) returns()
func (_WarChest *WarChestTransactor) UpdatePayer(opts *bind.TransactOpts, seaDropImpl common.Address, payer common.Address, allowed bool) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "updatePayer", seaDropImpl, payer, allowed)
}

// UpdatePayer is a paid mutator transaction binding the contract method 0xcb743ba8.
//
// Solidity: function updatePayer(address seaDropImpl, address payer, bool allowed) returns()
func (_WarChest *WarChestSession) UpdatePayer(seaDropImpl common.Address, payer common.Address, allowed bool) (*types.Transaction, error) {
	return _WarChest.Contract.UpdatePayer(&_WarChest.TransactOpts, seaDropImpl, payer, allowed)
}

// UpdatePayer is a paid mutator transaction binding the contract method 0xcb743ba8.
//
// Solidity: function updatePayer(address seaDropImpl, address payer, bool allowed) returns()
func (_WarChest *WarChestTransactorSession) UpdatePayer(seaDropImpl common.Address, payer common.Address, allowed bool) (*types.Transaction, error) {
	return _WarChest.Contract.UpdatePayer(&_WarChest.TransactOpts, seaDropImpl, payer, allowed)
}

// UpdatePublicDrop is a paid mutator transaction binding the contract method 0x1b73593c.
//
// Solidity: function updatePublicDrop(address seaDropImpl, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop) returns()
func (_WarChest *WarChestTransactor) UpdatePublicDrop(opts *bind.TransactOpts, seaDropImpl common.Address, publicDrop PublicDrop) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "updatePublicDrop", seaDropImpl, publicDrop)
}

// UpdatePublicDrop is a paid mutator transaction binding the contract method 0x1b73593c.
//
// Solidity: function updatePublicDrop(address seaDropImpl, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop) returns()
func (_WarChest *WarChestSession) UpdatePublicDrop(seaDropImpl common.Address, publicDrop PublicDrop) (*types.Transaction, error) {
	return _WarChest.Contract.UpdatePublicDrop(&_WarChest.TransactOpts, seaDropImpl, publicDrop)
}

// UpdatePublicDrop is a paid mutator transaction binding the contract method 0x1b73593c.
//
// Solidity: function updatePublicDrop(address seaDropImpl, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop) returns()
func (_WarChest *WarChestTransactorSession) UpdatePublicDrop(seaDropImpl common.Address, publicDrop PublicDrop) (*types.Transaction, error) {
	return _WarChest.Contract.UpdatePublicDrop(&_WarChest.TransactOpts, seaDropImpl, publicDrop)
}

// UpdateSignedMintValidationParams is a paid mutator transaction binding the contract method 0x511aa644.
//
// Solidity: function updateSignedMintValidationParams(address seaDropImpl, address signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams) returns()
func (_WarChest *WarChestTransactor) UpdateSignedMintValidationParams(opts *bind.TransactOpts, seaDropImpl common.Address, signer common.Address, signedMintValidationParams SignedMintValidationParams) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "updateSignedMintValidationParams", seaDropImpl, signer, signedMintValidationParams)
}

// UpdateSignedMintValidationParams is a paid mutator transaction binding the contract method 0x511aa644.
//
// Solidity: function updateSignedMintValidationParams(address seaDropImpl, address signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams) returns()
func (_WarChest *WarChestSession) UpdateSignedMintValidationParams(seaDropImpl common.Address, signer common.Address, signedMintValidationParams SignedMintValidationParams) (*types.Transaction, error) {
	return _WarChest.Contract.UpdateSignedMintValidationParams(&_WarChest.TransactOpts, seaDropImpl, signer, signedMintValidationParams)
}

// UpdateSignedMintValidationParams is a paid mutator transaction binding the contract method 0x511aa644.
//
// Solidity: function updateSignedMintValidationParams(address seaDropImpl, address signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams) returns()
func (_WarChest *WarChestTransactorSession) UpdateSignedMintValidationParams(seaDropImpl common.Address, signer common.Address, signedMintValidationParams SignedMintValidationParams) (*types.Transaction, error) {
	return _WarChest.Contract.UpdateSignedMintValidationParams(&_WarChest.TransactOpts, seaDropImpl, signer, signedMintValidationParams)
}

// UpdateTokenGatedDrop is a paid mutator transaction binding the contract method 0x7bc2be76.
//
// Solidity: function updateTokenGatedDrop(address seaDropImpl, address allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage) returns()
func (_WarChest *WarChestTransactor) UpdateTokenGatedDrop(opts *bind.TransactOpts, seaDropImpl common.Address, allowedNftToken common.Address, dropStage TokenGatedDropStage) (*types.Transaction, error) {
	return _WarChest.contract.Transact(opts, "updateTokenGatedDrop", seaDropImpl, allowedNftToken, dropStage)
}

// UpdateTokenGatedDrop is a paid mutator transaction binding the contract method 0x7bc2be76.
//
// Solidity: function updateTokenGatedDrop(address seaDropImpl, address allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage) returns()
func (_WarChest *WarChestSession) UpdateTokenGatedDrop(seaDropImpl common.Address, allowedNftToken common.Address, dropStage TokenGatedDropStage) (*types.Transaction, error) {
	return _WarChest.Contract.UpdateTokenGatedDrop(&_WarChest.TransactOpts, seaDropImpl, allowedNftToken, dropStage)
}

// UpdateTokenGatedDrop is a paid mutator transaction binding the contract method 0x7bc2be76.
//
// Solidity: function updateTokenGatedDrop(address seaDropImpl, address allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage) returns()
func (_WarChest *WarChestTransactorSession) UpdateTokenGatedDrop(seaDropImpl common.Address, allowedNftToken common.Address, dropStage TokenGatedDropStage) (*types.Transaction, error) {
	return _WarChest.Contract.UpdateTokenGatedDrop(&_WarChest.TransactOpts, seaDropImpl, allowedNftToken, dropStage)
}

// WarChestAllowedSeaDropUpdatedIterator is returned from FilterAllowedSeaDropUpdated and is used to iterate over the raw logs and unpacked data for AllowedSeaDropUpdated events raised by the WarChest contract.
type WarChestAllowedSeaDropUpdatedIterator struct {
	Event *WarChestAllowedSeaDropUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WarChestAllowedSeaDropUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WarChestAllowedSeaDropUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WarChestAllowedSeaDropUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WarChestAllowedSeaDropUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WarChestAllowedSeaDropUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WarChestAllowedSeaDropUpdated represents a AllowedSeaDropUpdated event raised by the WarChest contract.
type WarChestAllowedSeaDropUpdated struct {
	AllowedSeaDrop []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAllowedSeaDropUpdated is a free log retrieval operation binding the contract event 0xbbd3b69c138de4d317d0bc4290282c4e1cbd1e58b579a5b4f114b598c237454d.
//
// Solidity: event AllowedSeaDropUpdated(address[] allowedSeaDrop)
func (_WarChest *WarChestFilterer) FilterAllowedSeaDropUpdated(opts *bind.FilterOpts) (*WarChestAllowedSeaDropUpdatedIterator, error) {

	logs, sub, err := _WarChest.contract.FilterLogs(opts, "AllowedSeaDropUpdated")
	if err != nil {
		return nil, err
	}
	return &WarChestAllowedSeaDropUpdatedIterator{contract: _WarChest.contract, event: "AllowedSeaDropUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedSeaDropUpdated is a free log subscription operation binding the contract event 0xbbd3b69c138de4d317d0bc4290282c4e1cbd1e58b579a5b4f114b598c237454d.
//
// Solidity: event AllowedSeaDropUpdated(address[] allowedSeaDrop)
func (_WarChest *WarChestFilterer) WatchAllowedSeaDropUpdated(opts *bind.WatchOpts, sink chan<- *WarChestAllowedSeaDropUpdated) (event.Subscription, error) {

	logs, sub, err := _WarChest.contract.WatchLogs(opts, "AllowedSeaDropUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WarChestAllowedSeaDropUpdated)
				if err := _WarChest.contract.UnpackLog(event, "AllowedSeaDropUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAllowedSeaDropUpdated is a log parse operation binding the contract event 0xbbd3b69c138de4d317d0bc4290282c4e1cbd1e58b579a5b4f114b598c237454d.
//
// Solidity: event AllowedSeaDropUpdated(address[] allowedSeaDrop)
func (_WarChest *WarChestFilterer) ParseAllowedSeaDropUpdated(log types.Log) (*WarChestAllowedSeaDropUpdated, error) {
	event := new(WarChestAllowedSeaDropUpdated)
	if err := _WarChest.contract.UnpackLog(event, "AllowedSeaDropUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WarChestApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WarChest contract.
type WarChestApprovalIterator struct {
	Event *WarChestApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WarChestApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WarChestApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WarChestApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WarChestApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WarChestApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WarChestApproval represents a Approval event raised by the WarChest contract.
type WarChestApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_WarChest *WarChestFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*WarChestApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _WarChest.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &WarChestApprovalIterator{contract: _WarChest.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_WarChest *WarChestFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WarChestApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _WarChest.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WarChestApproval)
				if err := _WarChest.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_WarChest *WarChestFilterer) ParseApproval(log types.Log) (*WarChestApproval, error) {
	event := new(WarChestApproval)
	if err := _WarChest.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WarChestApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the WarChest contract.
type WarChestApprovalForAllIterator struct {
	Event *WarChestApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WarChestApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WarChestApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WarChestApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WarChestApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WarChestApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WarChestApprovalForAll represents a ApprovalForAll event raised by the WarChest contract.
type WarChestApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_WarChest *WarChestFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*WarChestApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _WarChest.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &WarChestApprovalForAllIterator{contract: _WarChest.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_WarChest *WarChestFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *WarChestApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _WarChest.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WarChestApprovalForAll)
				if err := _WarChest.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_WarChest *WarChestFilterer) ParseApprovalForAll(log types.Log) (*WarChestApprovalForAll, error) {
	event := new(WarChestApprovalForAll)
	if err := _WarChest.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WarChestBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the WarChest contract.
type WarChestBatchMetadataUpdateIterator struct {
	Event *WarChestBatchMetadataUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WarChestBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WarChestBatchMetadataUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WarChestBatchMetadataUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WarChestBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WarChestBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WarChestBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the WarChest contract.
type WarChestBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_WarChest *WarChestFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*WarChestBatchMetadataUpdateIterator, error) {

	logs, sub, err := _WarChest.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &WarChestBatchMetadataUpdateIterator{contract: _WarChest.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_WarChest *WarChestFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *WarChestBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _WarChest.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WarChestBatchMetadataUpdate)
				if err := _WarChest.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_WarChest *WarChestFilterer) ParseBatchMetadataUpdate(log types.Log) (*WarChestBatchMetadataUpdate, error) {
	event := new(WarChestBatchMetadataUpdate)
	if err := _WarChest.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WarChestConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the WarChest contract.
type WarChestConsecutiveTransferIterator struct {
	Event *WarChestConsecutiveTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WarChestConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WarChestConsecutiveTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WarChestConsecutiveTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WarChestConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WarChestConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WarChestConsecutiveTransfer represents a ConsecutiveTransfer event raised by the WarChest contract.
type WarChestConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_WarChest *WarChestFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*WarChestConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WarChest.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WarChestConsecutiveTransferIterator{contract: _WarChest.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_WarChest *WarChestFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *WarChestConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WarChest.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WarChestConsecutiveTransfer)
				if err := _WarChest.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_WarChest *WarChestFilterer) ParseConsecutiveTransfer(log types.Log) (*WarChestConsecutiveTransfer, error) {
	event := new(WarChestConsecutiveTransfer)
	if err := _WarChest.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WarChestContractURIUpdatedIterator is returned from FilterContractURIUpdated and is used to iterate over the raw logs and unpacked data for ContractURIUpdated events raised by the WarChest contract.
type WarChestContractURIUpdatedIterator struct {
	Event *WarChestContractURIUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WarChestContractURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WarChestContractURIUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WarChestContractURIUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WarChestContractURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WarChestContractURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WarChestContractURIUpdated represents a ContractURIUpdated event raised by the WarChest contract.
type WarChestContractURIUpdated struct {
	NewContractURI string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContractURIUpdated is a free log retrieval operation binding the contract event 0x905d981207a7d0b6c62cc46ab0be2a076d0298e4a86d0ab79882dbd01ac37378.
//
// Solidity: event ContractURIUpdated(string newContractURI)
func (_WarChest *WarChestFilterer) FilterContractURIUpdated(opts *bind.FilterOpts) (*WarChestContractURIUpdatedIterator, error) {

	logs, sub, err := _WarChest.contract.FilterLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return &WarChestContractURIUpdatedIterator{contract: _WarChest.contract, event: "ContractURIUpdated", logs: logs, sub: sub}, nil
}

// WatchContractURIUpdated is a free log subscription operation binding the contract event 0x905d981207a7d0b6c62cc46ab0be2a076d0298e4a86d0ab79882dbd01ac37378.
//
// Solidity: event ContractURIUpdated(string newContractURI)
func (_WarChest *WarChestFilterer) WatchContractURIUpdated(opts *bind.WatchOpts, sink chan<- *WarChestContractURIUpdated) (event.Subscription, error) {

	logs, sub, err := _WarChest.contract.WatchLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WarChestContractURIUpdated)
				if err := _WarChest.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractURIUpdated is a log parse operation binding the contract event 0x905d981207a7d0b6c62cc46ab0be2a076d0298e4a86d0ab79882dbd01ac37378.
//
// Solidity: event ContractURIUpdated(string newContractURI)
func (_WarChest *WarChestFilterer) ParseContractURIUpdated(log types.Log) (*WarChestContractURIUpdated, error) {
	event := new(WarChestContractURIUpdated)
	if err := _WarChest.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WarChestInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WarChest contract.
type WarChestInitializedIterator struct {
	Event *WarChestInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WarChestInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WarChestInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WarChestInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WarChestInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WarChestInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WarChestInitialized represents a Initialized event raised by the WarChest contract.
type WarChestInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WarChest *WarChestFilterer) FilterInitialized(opts *bind.FilterOpts) (*WarChestInitializedIterator, error) {

	logs, sub, err := _WarChest.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WarChestInitializedIterator{contract: _WarChest.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WarChest *WarChestFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WarChestInitialized) (event.Subscription, error) {

	logs, sub, err := _WarChest.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WarChestInitialized)
				if err := _WarChest.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WarChest *WarChestFilterer) ParseInitialized(log types.Log) (*WarChestInitialized, error) {
	event := new(WarChestInitialized)
	if err := _WarChest.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WarChestMaxSupplyUpdatedIterator is returned from FilterMaxSupplyUpdated and is used to iterate over the raw logs and unpacked data for MaxSupplyUpdated events raised by the WarChest contract.
type WarChestMaxSupplyUpdatedIterator struct {
	Event *WarChestMaxSupplyUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WarChestMaxSupplyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WarChestMaxSupplyUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WarChestMaxSupplyUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WarChestMaxSupplyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WarChestMaxSupplyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WarChestMaxSupplyUpdated represents a MaxSupplyUpdated event raised by the WarChest contract.
type WarChestMaxSupplyUpdated struct {
	NewMaxSupply *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMaxSupplyUpdated is a free log retrieval operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 newMaxSupply)
func (_WarChest *WarChestFilterer) FilterMaxSupplyUpdated(opts *bind.FilterOpts) (*WarChestMaxSupplyUpdatedIterator, error) {

	logs, sub, err := _WarChest.contract.FilterLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return &WarChestMaxSupplyUpdatedIterator{contract: _WarChest.contract, event: "MaxSupplyUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxSupplyUpdated is a free log subscription operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 newMaxSupply)
func (_WarChest *WarChestFilterer) WatchMaxSupplyUpdated(opts *bind.WatchOpts, sink chan<- *WarChestMaxSupplyUpdated) (event.Subscription, error) {

	logs, sub, err := _WarChest.contract.WatchLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WarChestMaxSupplyUpdated)
				if err := _WarChest.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMaxSupplyUpdated is a log parse operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 newMaxSupply)
func (_WarChest *WarChestFilterer) ParseMaxSupplyUpdated(log types.Log) (*WarChestMaxSupplyUpdated, error) {
	event := new(WarChestMaxSupplyUpdated)
	if err := _WarChest.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WarChestOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WarChest contract.
type WarChestOwnershipTransferredIterator struct {
	Event *WarChestOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WarChestOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WarChestOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WarChestOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WarChestOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WarChestOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WarChestOwnershipTransferred represents a OwnershipTransferred event raised by the WarChest contract.
type WarChestOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WarChest *WarChestFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WarChestOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WarChest.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WarChestOwnershipTransferredIterator{contract: _WarChest.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WarChest *WarChestFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WarChestOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WarChest.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WarChestOwnershipTransferred)
				if err := _WarChest.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WarChest *WarChestFilterer) ParseOwnershipTransferred(log types.Log) (*WarChestOwnershipTransferred, error) {
	event := new(WarChestOwnershipTransferred)
	if err := _WarChest.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WarChestPotentialOwnerUpdatedIterator is returned from FilterPotentialOwnerUpdated and is used to iterate over the raw logs and unpacked data for PotentialOwnerUpdated events raised by the WarChest contract.
type WarChestPotentialOwnerUpdatedIterator struct {
	Event *WarChestPotentialOwnerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WarChestPotentialOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WarChestPotentialOwnerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WarChestPotentialOwnerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WarChestPotentialOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WarChestPotentialOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WarChestPotentialOwnerUpdated represents a PotentialOwnerUpdated event raised by the WarChest contract.
type WarChestPotentialOwnerUpdated struct {
	NewPotentialAdministrator common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterPotentialOwnerUpdated is a free log retrieval operation binding the contract event 0x11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da.
//
// Solidity: event PotentialOwnerUpdated(address newPotentialAdministrator)
func (_WarChest *WarChestFilterer) FilterPotentialOwnerUpdated(opts *bind.FilterOpts) (*WarChestPotentialOwnerUpdatedIterator, error) {

	logs, sub, err := _WarChest.contract.FilterLogs(opts, "PotentialOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &WarChestPotentialOwnerUpdatedIterator{contract: _WarChest.contract, event: "PotentialOwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchPotentialOwnerUpdated is a free log subscription operation binding the contract event 0x11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da.
//
// Solidity: event PotentialOwnerUpdated(address newPotentialAdministrator)
func (_WarChest *WarChestFilterer) WatchPotentialOwnerUpdated(opts *bind.WatchOpts, sink chan<- *WarChestPotentialOwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _WarChest.contract.WatchLogs(opts, "PotentialOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WarChestPotentialOwnerUpdated)
				if err := _WarChest.contract.UnpackLog(event, "PotentialOwnerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePotentialOwnerUpdated is a log parse operation binding the contract event 0x11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da.
//
// Solidity: event PotentialOwnerUpdated(address newPotentialAdministrator)
func (_WarChest *WarChestFilterer) ParsePotentialOwnerUpdated(log types.Log) (*WarChestPotentialOwnerUpdated, error) {
	event := new(WarChestPotentialOwnerUpdated)
	if err := _WarChest.contract.UnpackLog(event, "PotentialOwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WarChestProvenanceHashUpdatedIterator is returned from FilterProvenanceHashUpdated and is used to iterate over the raw logs and unpacked data for ProvenanceHashUpdated events raised by the WarChest contract.
type WarChestProvenanceHashUpdatedIterator struct {
	Event *WarChestProvenanceHashUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WarChestProvenanceHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WarChestProvenanceHashUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WarChestProvenanceHashUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WarChestProvenanceHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WarChestProvenanceHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WarChestProvenanceHashUpdated represents a ProvenanceHashUpdated event raised by the WarChest contract.
type WarChestProvenanceHashUpdated struct {
	PreviousHash [32]byte
	NewHash      [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProvenanceHashUpdated is a free log retrieval operation binding the contract event 0x7c22004198bf87da0f0dab623c72e66ca1200f4454aa3b9ca30f436275428b7c.
//
// Solidity: event ProvenanceHashUpdated(bytes32 previousHash, bytes32 newHash)
func (_WarChest *WarChestFilterer) FilterProvenanceHashUpdated(opts *bind.FilterOpts) (*WarChestProvenanceHashUpdatedIterator, error) {

	logs, sub, err := _WarChest.contract.FilterLogs(opts, "ProvenanceHashUpdated")
	if err != nil {
		return nil, err
	}
	return &WarChestProvenanceHashUpdatedIterator{contract: _WarChest.contract, event: "ProvenanceHashUpdated", logs: logs, sub: sub}, nil
}

// WatchProvenanceHashUpdated is a free log subscription operation binding the contract event 0x7c22004198bf87da0f0dab623c72e66ca1200f4454aa3b9ca30f436275428b7c.
//
// Solidity: event ProvenanceHashUpdated(bytes32 previousHash, bytes32 newHash)
func (_WarChest *WarChestFilterer) WatchProvenanceHashUpdated(opts *bind.WatchOpts, sink chan<- *WarChestProvenanceHashUpdated) (event.Subscription, error) {

	logs, sub, err := _WarChest.contract.WatchLogs(opts, "ProvenanceHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WarChestProvenanceHashUpdated)
				if err := _WarChest.contract.UnpackLog(event, "ProvenanceHashUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProvenanceHashUpdated is a log parse operation binding the contract event 0x7c22004198bf87da0f0dab623c72e66ca1200f4454aa3b9ca30f436275428b7c.
//
// Solidity: event ProvenanceHashUpdated(bytes32 previousHash, bytes32 newHash)
func (_WarChest *WarChestFilterer) ParseProvenanceHashUpdated(log types.Log) (*WarChestProvenanceHashUpdated, error) {
	event := new(WarChestProvenanceHashUpdated)
	if err := _WarChest.contract.UnpackLog(event, "ProvenanceHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WarChestRoyaltyInfoUpdatedIterator is returned from FilterRoyaltyInfoUpdated and is used to iterate over the raw logs and unpacked data for RoyaltyInfoUpdated events raised by the WarChest contract.
type WarChestRoyaltyInfoUpdatedIterator struct {
	Event *WarChestRoyaltyInfoUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WarChestRoyaltyInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WarChestRoyaltyInfoUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WarChestRoyaltyInfoUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WarChestRoyaltyInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WarChestRoyaltyInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WarChestRoyaltyInfoUpdated represents a RoyaltyInfoUpdated event raised by the WarChest contract.
type WarChestRoyaltyInfoUpdated struct {
	Receiver common.Address
	Bps      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyInfoUpdated is a free log retrieval operation binding the contract event 0xf21fccf4d64d86d532c4e4eb86c007b6ad57a460c27d724188625e755ec6cf6d.
//
// Solidity: event RoyaltyInfoUpdated(address receiver, uint256 bps)
func (_WarChest *WarChestFilterer) FilterRoyaltyInfoUpdated(opts *bind.FilterOpts) (*WarChestRoyaltyInfoUpdatedIterator, error) {

	logs, sub, err := _WarChest.contract.FilterLogs(opts, "RoyaltyInfoUpdated")
	if err != nil {
		return nil, err
	}
	return &WarChestRoyaltyInfoUpdatedIterator{contract: _WarChest.contract, event: "RoyaltyInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchRoyaltyInfoUpdated is a free log subscription operation binding the contract event 0xf21fccf4d64d86d532c4e4eb86c007b6ad57a460c27d724188625e755ec6cf6d.
//
// Solidity: event RoyaltyInfoUpdated(address receiver, uint256 bps)
func (_WarChest *WarChestFilterer) WatchRoyaltyInfoUpdated(opts *bind.WatchOpts, sink chan<- *WarChestRoyaltyInfoUpdated) (event.Subscription, error) {

	logs, sub, err := _WarChest.contract.WatchLogs(opts, "RoyaltyInfoUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WarChestRoyaltyInfoUpdated)
				if err := _WarChest.contract.UnpackLog(event, "RoyaltyInfoUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoyaltyInfoUpdated is a log parse operation binding the contract event 0xf21fccf4d64d86d532c4e4eb86c007b6ad57a460c27d724188625e755ec6cf6d.
//
// Solidity: event RoyaltyInfoUpdated(address receiver, uint256 bps)
func (_WarChest *WarChestFilterer) ParseRoyaltyInfoUpdated(log types.Log) (*WarChestRoyaltyInfoUpdated, error) {
	event := new(WarChestRoyaltyInfoUpdated)
	if err := _WarChest.contract.UnpackLog(event, "RoyaltyInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WarChestSeaDropTokenDeployedIterator is returned from FilterSeaDropTokenDeployed and is used to iterate over the raw logs and unpacked data for SeaDropTokenDeployed events raised by the WarChest contract.
type WarChestSeaDropTokenDeployedIterator struct {
	Event *WarChestSeaDropTokenDeployed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WarChestSeaDropTokenDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WarChestSeaDropTokenDeployed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WarChestSeaDropTokenDeployed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WarChestSeaDropTokenDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WarChestSeaDropTokenDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WarChestSeaDropTokenDeployed represents a SeaDropTokenDeployed event raised by the WarChest contract.
type WarChestSeaDropTokenDeployed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSeaDropTokenDeployed is a free log retrieval operation binding the contract event 0xd7aca75208b9be5ffc04c6a01922020ffd62b55e68e502e317f5344960279af8.
//
// Solidity: event SeaDropTokenDeployed()
func (_WarChest *WarChestFilterer) FilterSeaDropTokenDeployed(opts *bind.FilterOpts) (*WarChestSeaDropTokenDeployedIterator, error) {

	logs, sub, err := _WarChest.contract.FilterLogs(opts, "SeaDropTokenDeployed")
	if err != nil {
		return nil, err
	}
	return &WarChestSeaDropTokenDeployedIterator{contract: _WarChest.contract, event: "SeaDropTokenDeployed", logs: logs, sub: sub}, nil
}

// WatchSeaDropTokenDeployed is a free log subscription operation binding the contract event 0xd7aca75208b9be5ffc04c6a01922020ffd62b55e68e502e317f5344960279af8.
//
// Solidity: event SeaDropTokenDeployed()
func (_WarChest *WarChestFilterer) WatchSeaDropTokenDeployed(opts *bind.WatchOpts, sink chan<- *WarChestSeaDropTokenDeployed) (event.Subscription, error) {

	logs, sub, err := _WarChest.contract.WatchLogs(opts, "SeaDropTokenDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WarChestSeaDropTokenDeployed)
				if err := _WarChest.contract.UnpackLog(event, "SeaDropTokenDeployed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSeaDropTokenDeployed is a log parse operation binding the contract event 0xd7aca75208b9be5ffc04c6a01922020ffd62b55e68e502e317f5344960279af8.
//
// Solidity: event SeaDropTokenDeployed()
func (_WarChest *WarChestFilterer) ParseSeaDropTokenDeployed(log types.Log) (*WarChestSeaDropTokenDeployed, error) {
	event := new(WarChestSeaDropTokenDeployed)
	if err := _WarChest.contract.UnpackLog(event, "SeaDropTokenDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WarChestTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WarChest contract.
type WarChestTransferIterator struct {
	Event *WarChestTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WarChestTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WarChestTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WarChestTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WarChestTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WarChestTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WarChestTransfer represents a Transfer event raised by the WarChest contract.
type WarChestTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_WarChest *WarChestFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*WarChestTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _WarChest.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &WarChestTransferIterator{contract: _WarChest.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_WarChest *WarChestFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WarChestTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _WarChest.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WarChestTransfer)
				if err := _WarChest.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_WarChest *WarChestFilterer) ParseTransfer(log types.Log) (*WarChestTransfer, error) {
	event := new(WarChestTransfer)
	if err := _WarChest.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
