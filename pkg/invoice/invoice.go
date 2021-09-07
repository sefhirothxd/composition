package invoice

//importando estuctura de customer
import (
	"github.com/sefhirothxd/composition/pkg/customer"
	"github.com/sefhirothxd/composition/pkg/invoiceitem"
)

type Invoice struct {
	contry string
	city   string
	total  float64
	client customer.Customer
	items  []invoiceitem.Item
}
