package abstract_factory

import "testing"

// GetMainAndDetail 调用工厂生成 DAO 并返回结果
func GetMainAndDetail(factory DAOFactory) (mainResult, detailResult string) {
	mainResult = factory.CreateOrderMainDAO().SaveOrderMain()
	detailResult = factory.CreateOrderDetailDAO().SaveOrderDetail()
	return
}

func TestAbstractFactory(t *testing.T) {
	tests := []struct {
		name           string
		factory        DAOFactory
		expectedMain   string
		expectedDetail string
	}{
		{"RDBFactory", &RDBDAOFactory{}, "rdb main save", "rdb detail save"},
		{"XMLFactory", &XMLDAOFactory{}, "xml main save", "xml detail save"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mainResult, detailResult := GetMainAndDetail(tt.factory)
			if mainResult != tt.expectedMain {
				t.Errorf("mainResult = %q, want %q", mainResult, tt.expectedMain)
			}
			if detailResult != tt.expectedDetail {
				t.Errorf("detailResult = %q, want %q", detailResult, tt.expectedDetail)
			}
		})
	}
}
