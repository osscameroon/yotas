package app

type Organisations struct {
	Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	GithubId    string `json:"github_id"`
	AvatarUrl   string `json:"avatar_url"`
	WebSite     string `json:"web_site"`
	Description string `json:"description"`
}

type OrganisationsArticles struct {
	Model
	OrganisationId uint `json:"organisation_id"`
	ArticleId      uint `json:"article_id"`
}

type OrganisationsUsers struct {
	Model
	OrganisationId uint `json:"organisation_id"`
	UserId         uint `json:"user_id"`
	Active         bool `json:"active"`
}

type Wallets struct {
	Model
	WalletId       string `json:"wallet_id"`
	UserId         uint   `json:"user_id"`
	OrganisationId uint   `json:"organisation_id"`
	Balance        int64  `json:"balance"`
}

type Operations struct {
	Model
	Amount        int64  `json:"amount"`
	Description   string `json:"description"`
	WalletId      string `json:"wallet_id"`
	OperationType string `json:"operation_type"`
	Approved      bool   `json:"approved"`
	OperationHash string `json:"operation_hash"`
}

type Orders struct {
	Model
	WalletId    string `json:"wallet_id"`
	TotalAmount int64  `json:"total_amount"`
	State       string `json:"state"`
	Decision    string `json:"decision"`
}

type OrdersArticles struct {
	Model
	OrderID      uint  `json:"order_id"`
	ArticleID    uint  `json:"article_id"`
	ArticlePrice int64 `json:"article_price"`
	Quantity     int   `json:"quantity"`
}

type OrderPresenter struct {
	*Orders
	Items []*OrderItemPresenter `json:"items"`
}

type OrderItemPresenter struct {
	*OrdersArticles
	Article ArticlesPresenter `json:"article"`
}

type OrderDecision struct {
	Accepted bool   `json:"accepted"`
	Reason   string `json:"reason"`
}

type Articles struct {
	Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
	Price       int64  `json:"price"`
	Metadata    string `json:"metadata"`
}

type Pictures struct {
	Model
	OrganisationId uint   `json:"organisation_id"`
	AltText        string `json:"alt_text"`
	Original       string `json:"original"`
	Small          string `json:"small"`
	Medium         string `json:"medium"`
	Large          string `json:"large"`
}

type ArticlesPictures struct {
	Model
	PictureId uint `json:"picture_id"`
	ArticleId uint `json:"article_id"`
}

type ArticlesPresenter struct {
	Articles
	Pictures []Pictures `json:"pictures"`
}

type Users struct {
	Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	GithubId    string `json:"github_id"`
	GithubToken string `json:"github_token"`
	AvatarUrl   string `json:"avatar_url"`
}

type UsersPresenter struct {
	Users
}

type Callback struct {
	Code  string `json:"code"`
	State string `json:"state"`
}
