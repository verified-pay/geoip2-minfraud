package geoip2

import (
	"fmt"
	"log"
	"net"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/verified-pay/geoip2-minfraud/model"
)

// Example provides a basic example of using the API. Use of the Country
// method is analogous to that of the City method.
func Example() {
	db, err := Open(testDatabase)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP("81.2.69.142")
	record, err := db.City(ip)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("City name: %v\n", record.City.Names["en"])
	if len(record.Subdivisions) > 0 {
		fmt.Printf("Subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	}
	fmt.Printf("Country name: %v\n", record.Country.Names["en"])
	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
}

func TestMinfraud(t *testing.T) {
	_ = godotenv.Load()

	accountID := os.Getenv("MAXMIND_ACCOUNT_ID")
	licenseKey := os.Getenv("MAXMIND_LICENSE_KEY")
	if accountID == "" || licenseKey == "" {
		t.Skip("Skipping: MAXMIND_ACCOUNT_ID and MAXMIND_LICENSE_KEY environment variables not set")
	}

	mf, err := NewMinFraud(accountID, licenseKey)
	require.NoError(t, err)

	mf.WithDevice(&model.DeviceReq{
		IPAddress: "81.2.69.142",
	}).WithEvent(&model.EventReq{
		TransactionID: "test-txn-123",
		Type:          "purchase",
		Time:          time.Now(),
	})

	insights, err, apiErr := mf.Insights()
	if err == ErrReceived {
		if apiErr.Code == "AUTHORIZATION_INVALID" {
			t.Skip("Skipping: Invalid API credentials")
		}
		t.Fatalf("API error: %s - %s", apiErr.Code, apiErr.Error)
	}
	require.NoError(t, err)

	assert.NotEmpty(t, insights.ID)
	assert.GreaterOrEqual(t, insights.RiskScore, float32(0))
	assert.LessOrEqual(t, insights.RiskScore, float32(99))
	t.Logf("Risk Score: %.2f", insights.RiskScore)
	t.Logf("IP Risk: %.2f", insights.IPAddress.Risk)
	t.Logf("IP Country: %s", insights.IPAddress.Country.IsoCode)
}
