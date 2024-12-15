package helper

import (
    // "fmt"
    "github.com/midtrans/midtrans-go"
    "github.com/midtrans/midtrans-go/snap"
)

type MidtransConfig struct {
    ServerKey string
    ClientKey string
    Environment midtrans.EnvironmentType
}

type MidtransClient struct {
    Client *snap.Client
    Config MidtransConfig
}

func NewMidtransClient(config MidtransConfig) *MidtransClient {
    midtrans.ServerKey = config.ServerKey
    midtrans.ClientKey = config.ClientKey
    midtrans.Environment = config.Environment

    return &MidtransClient{
        Client: &snap.Client{},
        Config: config,
    }
}

func (m *MidtransClient) GenerateToken(transactionDetails *midtrans.TransactionDetails) (*snap.Response, error) {
    req := &snap.Request{
        TransactionDetails: *transactionDetails,
    }

    return snap.CreateTransaction(req)
}

func (m *MidtransClient) VerifyCallback(payload map[string]interface{}) bool {
    // signatureKey := fmt.Sprintf("%s%s%d%s", 
    //     midtrans.ServerKey, 
    //     payload["order_id"], 
    //     int64(payload["gross_amount"].(float64)), 
    //     payload["status_code"],
    // )
    
    return true
}