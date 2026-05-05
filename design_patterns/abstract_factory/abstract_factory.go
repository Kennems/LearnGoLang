package abstract_factory

// OrderMainDAO 为订单主记录
type OrderMainDAO interface {
	SaveOrderMain() string
}

// OrderDetailDAO 为订单详情纪录
type OrderDetailDAO interface {
	SaveOrderDetail() string
}

// DAOFactory DAO 抽象工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

// ------------------ RDB 实现 ------------------

type RDBMainDAO struct{}

func (*RDBMainDAO) SaveOrderMain() string {
	return "rdb main save"
}

type RDBDetailDAO struct{}

func (*RDBDetailDAO) SaveOrderDetail() string {
	return "rdb detail save"
}

type RDBDAOFactory struct{}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

// ------------------ XML 实现 ------------------

type XMLMainDAO struct{}

func (*XMLMainDAO) SaveOrderMain() string {
	return "xml main save"
}

type XMLDetailDAO struct{}

func (*XMLDetailDAO) SaveOrderDetail() string {
	return "xml detail save"
}

type XMLDAOFactory struct{}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}
