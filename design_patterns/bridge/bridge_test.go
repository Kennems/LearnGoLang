package bridge
import "testing"
func TestBridge(t *testing.T) {
	sa := &SoftwareA{}
	pa := NewPhoneA("pa")
	pa.SetSoft(sa)
	if res := pa.Run(); res != "PhoneA pa run: run soft a" {
		t.Errorf("Unexpected result: %s", res)
	}
}
