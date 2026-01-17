package store


type SiteStatus struct{
	Address string
	StatusCode int
	Status string
}

type Store struct{
	data map[string] SiteStatus;
}

func NewStore() *Store{
	return &Store{
		data: make(map[string]SiteStatus),
	}
}

func (s *Store) Insert(siteStatus SiteStatus){
	s.data[siteStatus.Address] = SiteStatus{
		Address: 	siteStatus.Address,
		StatusCode: siteStatus.StatusCode,
		Status: 	siteStatus.Status,
	}
}

func (s *Store) Get(address string)(SiteStatus, bool){
	site, ok := s.data[address]
	return site, ok
}

func (s *Store) GetAll() []SiteStatus{
	var entries []SiteStatus;
	for _, val := range s.data{
		entries = append(entries, val)
	}
	return entries
}