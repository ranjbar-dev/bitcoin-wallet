package test

import (
	BitcoinWallet "github.com/Amirilidan78/bitcoin-wallet"
	"testing"
)

// GenerateBitcoinWallet test
func TestGenerateWallet(t *testing.T) {
	w := BitcoinWallet.GenerateBitcoinWallet(node)
	if w == nil {
		t.Errorf("GenerateBitcoinWallet res was incorect, got: %q, want: %q.", "wallet", "*BitcoinWallet")
	}
	if len(w.PrivateKey) == 0 {
		t.Errorf("GenerateBitcoinWallet PrivateKey was incorect, got: %q, want: %q.", w.PrivateKey, "valid PrivateKey")
	}
	if len(w.PublicKey) == 0 {
		t.Errorf("GenerateBitcoinWallet PublicKey was incorect, got: %q, want: %q.", w.PublicKey, "valid PublicKey")
	}
	if len(w.Address) == 0 {
		t.Errorf("GenerateBitcoinWallet Address was incorect, got: %q, want: %q.", w.Address, "valid Address")
	}
}

// CreateBitcoinWallet test
func TestCreateWallet(t *testing.T) {
	_, err := BitcoinWallet.CreateBitcoinWallet(node, invalidPrivateKey)
	if err == nil {
		t.Errorf("CreateBitcoinWallet error was incorect, got: %q, want: %q.", err, "not nil")
	}

	w, err := BitcoinWallet.CreateBitcoinWallet(node, validPrivateKey)
	if err != nil {
		t.Errorf("CreateBitcoinWallet error was incorect, got: %q, want: %q.", err, "nil")
	}
	if len(w.PrivateKey) == 0 {
		t.Errorf("CreateBitcoinWallet PrivateKey was incorect, got: %q, want: %q.", w.PrivateKey, "valid PrivateKey")
	}
	if len(w.PublicKey) == 0 {
		t.Errorf("CreateBitcoinWallet PublicKey was incorect, got: %q, want: %q.", w.PublicKey, "valid PublicKey")
	}
	if len(w.Address) == 0 {
		t.Errorf("CreateBitcoinWallet Address was incorect, got: %q, want: %q.", w.Address, "valid Address")
	}
	if len(w.Address) == 0 {
		t.Errorf("CreateBitcoinWallet AddressBase58 was incorect, got: %q, want: %q.", w.Address, "valid Address")
	}
}

// PrivateKeyRCDSA test
func TestPrivateKeyRCDSA(t *testing.T) {
	w := wallet()

	_, err := w.PrivateKeyRCDSA()
	if err != nil {
		t.Errorf("PrivateKeyRCDSA error was incorect, got: %q, want: %q.", err, "nil")
	}
}

// PrivateKeyBTCE test
func TestPrivateKeyBTCE(t *testing.T) {
	w := wallet()

	_, err := w.PrivateKeyBTCE()
	if err != nil {
		t.Errorf("PrivateKeyRCDSA error was incorect, got: %q, want: %q.", err, "nil")
	}
}

// PrivateKeyBytes test
func TestPrivateKeyBytes(t *testing.T) {
	w := wallet()

	bytes, err := w.PrivateKeyBytes()
	if err != nil {
		t.Errorf("PrivateKeyBytes error was incorect, got: %q, want: %q.", err, "nil")
	}
	if len(bytes) == 0 {
		t.Errorf("PrivateKeyBytes bytes len was incorect, got: %q, want: %q.", len(bytes), "more than 0")
	}
}

// Balance test
func TestBalance(t *testing.T) {
	w := wallet()

	_, err := w.Balance()
	if err != nil {
		t.Errorf("Balance error was incorect, got: %q, want: %q.", err, "nil")
	}
}

// Transfer test
func TestTransfer(t *testing.T) {
	w := wallet()

	/// TODO : uncomment this after checking to Address added
	//_, err := w.Transfer(invalidToAddress, ethAmount)
	//if err == nil {
	//	t.Errorf("Transfer error was incorect, got: %q, want: %q.", err, "not nil becuase to address is invalid")
	//}

	txId, err := w.Transfer(validToAddress, btcAmount, feeAmount)
	if err != nil {
		t.Errorf("Transfer error was incorect, got: %q, want: %q.", err, "nil")
	}
	if len(txId) == 0 {
		t.Errorf("Transfer txId was incorect, got: %q, want: %q.", txId, "not nil")
	}
}
