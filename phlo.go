package plivo

type PhloService struct {
	client *Client
}

type Phlo struct {
	PhloId     string `json:"phlo_id" url:"phlo_id"`
	Name     string `json:"name" url:"name"`
	CreatedOn string `json:"created_on" url:"created_on"`
}


func (service *PhloService) Get(phlo_id string) (response *Phlo, err error) {
	req, err := service.client.NewRequestPhlo("GET", nil,"%s", phlo_id)
	response = &Phlo{}
	err = service.client.ExecuteRequest(req, response)
	return
}