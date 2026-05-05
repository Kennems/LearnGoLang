package chain_of_responsibility
import "testing"
func TestChain(t *testing.T) {
	jinli := NewCommonManager("jinli")
	zongjian := NewMajordomo("zongjian")
	zongjingli := NewGeneralManager("zongjingli")
	jinli.SetSuperior(zongjian)
	zongjian.SetSuperior(zongjingli)

	req := Request{RequestType: "请假", Number: 1}
	res := jinli.RequestApplications(req)
	if res != "jinli:请假 数量1 被批准" { t.Errorf("Unexpected res: %s", res) }
	
	req2 := Request{RequestType: "加薪", Number: 1000}
	res2 := jinli.RequestApplications(req2)
	if res2 != "zongjingli:加薪 数量1000 再说吧" { t.Errorf("Unexpected res: %s", res2) }
}
