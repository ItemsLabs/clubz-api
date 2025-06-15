package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"math/big"

	connections "github.com/itemslabs/clubz-api/web3"
	"github.com/labstack/echo/v4"
)

// ValidateUUID checks if a string is a valid UUID v4.
func ValidateUUID(uuid string) bool {
	regex := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89ABab][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	return regex.MatchString(uuid)
}

// MintLaLigaToken godoc
// @Summary Mint a new LaLiga token
// @Description Mint a new LaLiga token
// @ID mint-laliga-token
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param toAddress formData string true "Address to mint token to"
// @Param tokenId formData string true "Token ID to mint"
// @Success 200 {object} string
// @Router /web3/mintLaLiga [post]
func (e *Env) MintLaLigaToken(c echo.Context) error {
	toAddress := c.FormValue("toAddress")
	tokenIdStr := c.FormValue("tokenId")

	if toAddress == "" || tokenIdStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing required form values")
	}

	tokenIdInt, err := strconv.ParseInt(tokenIdStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid tokenId value: %s", tokenIdStr))
	}
	tokenId := big.NewInt(tokenIdInt)

	tx, err := connections.MintLaLiga(toAddress, tokenId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error minting LaLiga token: %v", err))
	}

	return c.JSON(http.StatusOK, map[string]string{"transaction": tx.Hash().Hex()})
}

// BatchMintLaLigaToken godoc
// @Summary Batch mint LaLiga tokens
// @Description Batch mint LaLiga tokens
// @ID batch-mint-laliga-token
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param toAddress formData string true "Address to mint tokens to"
// @Param tokenIds formData string true "Comma-separated list of token IDs to mint"
// @Success 200 {object} string
// @Router /web3/batchMintLaLiga [post]
func (e *Env) BatchMintLaLigaToken(c echo.Context) error {
	toAddress := c.FormValue("toAddress")
	tokenIdsStr := c.FormValue("tokenIds")

	if toAddress == "" || tokenIdsStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing required form values")
	}

	tokenIdsStrArray := strings.Split(tokenIdsStr, ",")
	tokenIds := make([]*big.Int, len(tokenIdsStrArray))

	for i, idStr := range tokenIdsStrArray {
		idInt, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid tokenId value: %s", idStr))
		}
		tokenIds[i] = big.NewInt(idInt)
	}

	tx, err := connections.BatchMintLaLiga(toAddress, tokenIds)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error batch minting LaLiga tokens: %v", err))
	}

	return c.JSON(http.StatusOK, map[string]string{"transaction": tx.Hash().Hex()})
}

// GetBalanceOf godoc
// @Summary Get balance of an address
// @Description Get balance of an address
// @ID get-balance
// @Produce json
// @Param address path string true "Address to get balance of"
// @Param contractType path string true "Contract type"
// @Success 200 {object} string
// @Router /web3/balance/{address}/{contractType} [get]
func (e *Env) GetBalanceOf(c echo.Context) error {
	address := c.Param("address")
	contractType := c.Param("contractType")

	balance, err := connections.GetBalanceOf(address, contractType)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"balance": balance.String()})
}

// ApproveToken godoc
// @Summary Approve a token
// @Description Approve a token
// @ID approve-token
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param toAddress formData string true "Address to approve token to"
// @Param tokenId formData string true "Token ID to approve"
// @Param contractType formData string true "Contract type"
// @Success 200 {object} string
// @Router /web3/approve [post]
func (e *Env) ApproveToken(c echo.Context) error {
	toAddress := c.FormValue("toAddress")
	tokenId, err := strconv.ParseInt(c.FormValue("tokenId"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid tokenId")
	}
	contractType := c.FormValue("contractType")

	tx, err := connections.ApproveToken(toAddress, big.NewInt(tokenId), contractType)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"transaction": tx.Hash().Hex()})
}

// SetApprovalForAll godoc
// @Summary Set approval for all
// @Description Set approval for all
// @ID set-approval-for-all
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param operator formData string true "Operator address"
// @Param approved formData string true "Approved value"
// @Param contractType formData string true "Contract type"
// @Success 200 {object} string
// @Router /web3/setApprovalForAll [post]
func (e *Env) SetApprovalForAll(c echo.Context) error {
	operator := c.FormValue("operator")
	approved, err := strconv.ParseBool(c.FormValue("approved"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid 'approved' value")
	}
	contractType := c.FormValue("contractType")

	tx, err := connections.SetApprovalForAll(operator, approved, contractType)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"transaction": tx.Hash().Hex()})
}

// TransferToken godoc
// @Summary Transfer a token
// @Description Transfer a token
// @ID transfer-token
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param fromAddress formData string true "Address to transfer token from"
// @Param toAddress formData string true "Address to transfer token to"
// @Param tokenId formData string true "Token ID to transfer"
// @Param contractType formData string true "Contract type"
// @Success 200 {object} string
// @Router /web3/transfer [post]
func (e *Env) TransferToken(c echo.Context) error {
	fromAddress := c.FormValue("fromAddress")
	toAddress := c.FormValue("toAddress")
	tokenId, err := strconv.ParseInt(c.FormValue("tokenId"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid tokenId")
	}
	contractType := c.FormValue("contractType")

	tx, err := connections.TransferToken(fromAddress, toAddress, big.NewInt(tokenId), contractType)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"transaction": tx.Hash().Hex()})
}

// BurnToken godoc
// @Summary Burn a token
// @Description Burn a token
// @ID burn-token
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param tokenId formData string true "Token ID to burn"
// @Param contractType formData string true "Contract type"
// @Success 200 {object} string
// @Router /web3/burn [post]
func (e *Env) BurnToken(c echo.Context) error {
	tokenId, err := strconv.ParseInt(c.FormValue("tokenId"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid tokenId")
	}
	contractType := c.FormValue("contractType")

	tx, err := connections.BurnToken(big.NewInt(tokenId), contractType)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"transaction": tx.Hash().Hex()})
}

// GetTokenURI godoc
// @Summary Get token URI
// @Description Get token URI
// @ID get-token-uri
// @Produce json
// @Param tokenId path string true "Token ID"
// @Param contractType path string true "Contract type"
// @Success 200 {object} string
// @Router /web3/tokenURI/{tokenId}/{contractType} [get]
func (e *Env) GetTokenURI(c echo.Context) error {
	tokenId, err := strconv.ParseInt(c.Param("tokenId"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid tokenId")
	}
	contractType := c.Param("contractType")

	uri, err := connections.GetTokenURI(big.NewInt(tokenId), contractType)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"tokenURI": uri})
}

// OwnerOfToken godoc
// @Summary Get owner of a token
// @Description Get owner of a token
// @ID owner-of-token
// @Produce json
// @Param tokenId path string true "Token ID"
// @Param contractType path string true "Contract type"
// @Success 200 {object} string
// @Router /web3/ownerOfToken/{tokenId}/{contractType} [get]
func (e *Env) OwnerOfToken(c echo.Context) error {
	tokenId, err := strconv.ParseInt(c.Param("tokenId"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid tokenId")
	}
	contractType := c.Param("contractType")

	owner, err := connections.OwnerOfToken(big.NewInt(tokenId), contractType)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"owner": owner.Hex()})
}

// IsApprovedForAll godoc
// @Summary Check if an address is approved for all
// @Description Check if an address is approved for all
// @ID is-approved-for-all
// @Produce json
// @Param ownerAddress path string true "Owner address"
// @Param operatorAddress path string true "Operator address"
// @Param contractType path string true "Contract type"
// @Success 200 {object} bool
// @Router /web3/isApprovedForAll/{ownerAddress}/{operatorAddress}/{contractType} [get]
func (e *Env) IsApprovedForAll(c echo.Context) error {
	ownerAddress := c.Param("ownerAddress")
	operatorAddress := c.Param("operatorAddress")
	contractType := c.Param("contractType")

	isApproved, err := connections.IsApprovedForAll(ownerAddress, operatorAddress, contractType)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]bool{"isApprovedForAll": isApproved})
}

// GetOwner godoc
// @Summary Get owner of a contract
// @Description Get owner of a contract
// @ID get-owner
// @Produce json
// @Param contractType path string true "Contract type"
// @Success 200 {object} string
// @Router /web3/owner/{contractType} [get]
func (e *Env) GetOwner(c echo.Context) error {
	contractType := c.Param("contractType")

	owner, err := connections.GetOwner(contractType)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"owner": owner.Hex()})
}

// GetApproved godoc
// @Summary Get approved address
// @Description Get approved address
// @ID get-approved
// @Produce json
// @Param tokenId path string true "Token ID"
// @Param contractType path string true "Contract type"
// @Success 200 {object} string
// @Router /web3/approved/{tokenId}/{contractType} [get]
func (e *Env) GetApproved(c echo.Context) error {
	tokenId, err := strconv.ParseInt(c.Param("tokenId"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid tokenId")
	}
	contractType := c.Param("contractType")

	approvedAddress, err := connections.GetApproved(big.NewInt(tokenId), contractType)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"approvedAddress": approvedAddress.Hex()})
}

// SafeTransferFrom godoc
// @Summary Safe transfer from
// @Description Safe transfer from
// @ID safe-transfer-from
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param fromAddress formData string true "Address to transfer token from"
// @Param toAddress formData string true "Address to transfer token to"
// @Param tokenId formData string true "Token ID to transfer"
// @Param contractType formData string true "Contract type"
// @Success 200 {object} string
// @Router /web3/safeTransferFrom [post]
func (e *Env) SafeTransferFrom(c echo.Context) error {
	fromAddress := c.FormValue("fromAddress")
	toAddress := c.FormValue("toAddress")
	tokenId, err := strconv.ParseInt(c.FormValue("tokenId"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid tokenId")
	}
	contractType := c.FormValue("contractType")

	tx, err := connections.SafeTransferFrom(fromAddress, toAddress, big.NewInt(tokenId), contractType)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"transaction": tx.Hash().Hex()})
}

// SendSignedTransaction godoc
// @Summary Send signed transaction
// @Description Send signed transaction
// @ID send-signed-transaction
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param signedTxHex formData string true "Signed transaction hex string"
// @Success 200 {object} string
// @Router /web3/sendSignedTransaction [post]
func (e *Env) SendSignedTransaction(c echo.Context) error {
	signedTxHex := c.FormValue("signedTxHex")
	if signedTxHex == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Signed transaction hex string is required")
	}

	tx, err := connections.SendSignedTransaction(signedTxHex)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"transactionHash": tx.Hash().Hex()})
}
