// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3

package types

type Base struct {
	ID        uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CreatedAt string `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	IsDeleted bool   `gorm:"column:is_deleted;default:false" json:"is_deleted"`
}

type Bounty struct {
	ApplicantDeposit            int          `json:"applicants_deposit"`
	ChainId                     uint64       `json:"chain_id"`
	ComerId                     uint64       `json:"comer_id"`
	CreatedAt                   string       `json:"created_at"`
	DepositContractAddress      string       `json:"deposit_contract_address"`
	DepositContractTokenDecimal int          `json:"deposit_contract_token_decimal"`
	DepositContractTokenSymbol  string       `json:"deposit_contract_token_symbol"`
	DiscussionLink              string       `json:"discussion_link"`
	ExpiredTime                 string       `json:"expired_time"`
	FounderDeposit              int          `json:"founder_deposit"`
	Id                          uint64       `json:"id"`
	IsLock                      int          `json:"is_lock"`
	PaymentMode                 int          `json:"payment_mode"`
	Reward                      BountyReward `json:"reward"`
	Startup                     StartupBasic `json:"startup"`
	StartupId                   uint64       `json:"startup_id"`
	Status                      int          `json:"status"`
	Title                       string       `json:"title"`
	TxHash                      string       `json:"tx_hash"`
}

type BountyApplicant struct {
	BountyId    uint64 `json:"bountyId"`    // 赏金任务ID
	ComerID     uint64 `json:"comerId"`     // 申请人ID
	ApplyAt     string `json:"applyAt"`     // 申请时间
	RevokeAt    string `json:"revokeAt"`    // 撤销时间
	ApproveAt   string `json:"approveAt"`   // 批准时间
	QuitAt      string `json:"quitAt"`      // 退出时间
	SubmitAt    string `json:"submitAt"`    // 提交时间
	Status      int    `json:"status"`      // 申请状态
	Description string `json:"description"` // 申请描述
}

type BountyComer struct {
	Activation     bool                  `json:"activation"`
	Address        string                `json:"address"`
	Avatar         string                `json:"avatar"`
	Banner         string                `json:"banner"`
	CustomDomain   string                `json:"customDomain"`
	Id             int                   `json:"id"`
	InvitationCode string                `json:"invitationCode"`
	IsConnected    bool                  `json:"isConnected"`
	Location       string                `json:"location"`
	Name           string                `json:"name"`
	Skills         []TagRelationResponse `json:"skills"`
	TimeZone       string                `json:"timeZone"`
}

type BountyContact struct {
	BountyId       uint64 `json:"bountyId"`       // 赏金任务ID
	ContactType    uint8  `json:"contactType"`    // 联系方式类型
	ContactAddress string `json:"contactAddress"` // 联系地址
}

type BountyDepositRecord struct {
	Amount    int                `json:"amount"`
	BountyId  uint64             `json:"bountyId"`
	Comer     ComerBasicResponse `json:"comer"`
	ComerId   uint64             `json:"comerId"`
	CreatedAt string             `json:"createdAt"`
	Id        uint64             `json:"id"`
	Mode      int                `json:"mode"`
	Status    int8               `json:"status"`
	TxHash    string             `json:"txHash"`
}

type BountyPaymentPeriod struct {
	BountyId     uint64 `json:"bountyId"`     // 赏金任务ID（唯一索引）
	PeriodType   int    `json:"periodType"`   // 周期类型
	PeriodAmount uint64 `json:"periodAmount"` // 周期数量
	HoursPerDay  int    `json:"hoursPerDay"`  // 每日小时数
	Token1Symbol string `json:"token1Symbol"` // 代币1符号
	Token1Amount int    `json:"token1Amount"` // 代币1数量
	Token2Symbol string `json:"token2Symbol"` // 代币2符号
	Token2Amount int    `json:"token2Amount"` // 代币2数量
	Target       string `json:"target"`       // 目标描述
}

type BountyPaymentTerms struct {
	BountyId     uint64 `json:"bountyId"`     // 关联的赏金任务ID
	PaymentMode  int8   `json:"paymentMode"`  // 支付方式
	Token1Symbol string `json:"token1Symbol"` // 第一种代币符号
	Token1Amount int    `json:"token1Amount"` // 第一种代币数量
	Token2Symbol string `json:"token2Symbol"` // 第二种代币符号
	Token2Amount int    `json:"token2Amount"` // 第二种代币数量
	Terms        string `json:"terms"`        // 支付条款详情
	SeqNum       int    `json:"seqNum"`       // 排序序号
	Status       int    `json:"status"`       // 状态
}

type BountyReward struct {
	BountyId     uint64 `json:"bountyId"`
	Token1Symbol string `json:"token1Symbol"`
	Token1Amount int    `json:"token1Amount"`
	Token2Symbol string `json:"token2Symbol"`
	Token2Amount int    `json:"token2Amount"`
}

type ChainBasicResponse struct {
	ChainID        uint64          `json:"chain_id"` // 链ID（唯一索引）
	Name           string          `json:"name"`     // 链名称
	Logo           string          `json:"logo"`     // 链Logo
	Status         int8            `json:"status"`   // 状态：1-正常，2-禁用
	ChainContracts []ChainContract `json:"chain_contracts"`
	ChainEndpoints []ChainEndpoint `json:"chain_endpoints"`
}

type ChainContract struct {
	ChainID       uint64 `json:"chain_id"`        // 链ID
	Address       string `json:"address"`         // 合约地址
	Project       int8   `json:"project"`         // 项目类型：1-Startup, 2-Bounty, 3-Crowdfunding, 4-Gover
	Type          int8   `json:"type"`            // 合约类型：1-工厂合约, 2-子合约
	Version       string `json:"version"`         // 合约版本
	ABI           string `json:"abi"`             // ABI JSON
	CreatedTxHash string `json:"created_tx_hash"` // 创建交易哈希
}

type ChainEndpoint struct {
	Protocol int8   `json:"protocol"` // 通信协议：1-rpc 2-wss
	ChainID  uint64 `json:"chain_id"` // 链ID
	URL      string `json:"url"`      // 节点URL
	Status   int8   `json:"status"`   // 状态：1-正常 2-禁用
}

type ComerAccountResponse struct {
	Avatar    string `json:"avatar"`
	ComerId   int    `json:"comer_id"`
	Id        int    `json:"id"`
	IsLinked  bool   `json:"is_linked"`
	IsPrimary bool   `json:"is_primary"`
	Nickname  string `json:"nickname"`
	Oin       string `json:"oin"`
	Type      int    `json:"type"`
}

type ComerConnectedTotalResponse struct {
	BeConnectComerTotal int `json:"be_connect_comer_total"`
	ConnectComerTotal   int `json:"connect_comer_total"`
	ConnectStartupTotal int `json:"connect_startup_total"`
}

type ComerEducationResponse struct {
	ComerId     int    `json:"comer_id"`
	GraduatedAt string `json:"graduated_at"`
	Id          int    `json:"id"`
	Major       string `json:"major"`
	School      string `json:"school"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
	Level       int    `json:"level"`
	Degree      string `json:"degree"`
}

type ComerInfo struct {
	Bio     string `json:"bio"`
	ComerId int    `json:"comer_id"`
	Id      int    `json:"id"`
}

type ComerLanguageResponse struct {
	ComerId  int    `json:"comer_id"`
	Id       int    `json:"id"`
	Language string `json:"language"`
	Code     string `json:"code"`
	Level    int    `json:"level"`
	IsNative bool   `json:"is_native"`
}

type ComerSkillResponse struct {
	ComerId     int    `json:"comer_id"`
	Id          int    `json:"id"`
	SkillName   string `json:"skill_name"`
	Level       int    `json:"level"`
	Years       int    `json:"years"`
	Description string `json:"description"`
}

type ComerSocialResponse struct {
	ComerId      int    `json:"comer_id"`
	Id           int    `json:"id"`
	PlatformName string `json:"platform_name"`
	UserName     string `json:"user_name"`
	PlatformId   string `json:"platform_id"`
	IsVerified   bool   `json:"is_verified"`
	Url          string `json:"url"`
}

type Contact struct {
	ContactType    uint8  `json:"contactType"` // 1:Email 2:Discord 3:Telegram
	ContactAddress string `json:"contactAddress"`
}

type FollowRelation struct {
	ComerID uint64 `json:comerID`
}

type OauthAccountBindingInfo struct {
	Linked      bool   `json:"linked"`
	AccountType int    `json:"accountType"`
	AccountId   uint64 `json:"accountId"`
}

type PostUpdate struct {
	SourceType int    `json:"sourceType"`
	SourceID   uint64 `json:"sourceId"`
	ComerID    uint64 `json:"comerId"`
	Content    string `json:"content"`
	TimeStamp  string `json:"timeStamp"` // post time
}

type SimpleStartupInfo struct {
	Avatar  string `json:"avatar"`
	Id      int    `json:"id"`
	Name    string `json:"name"`
	OnChain bool   `json:"on_chain"`
}

type SocialBookResponse struct {
	Id           int                `json:"id"`
	SocialTool   SocialToolResponse `json:"socialTool"`
	SocialToolId int                `json:"socialToolId"`
	TargetId     int                `json:"targetId"`
	Type         int                `json:"type"`
	Value        string             `json:"value"`
}

type SocialToolResponse struct {
	Id   int    `json:"id"`
	Logo string `json:"logo"`
	Name string `json:"name"`
}

type Startup struct {
	ComerID              uint64 `json:"comerID" db:"comer_id"`
	Name                 string `json:"name" db:"name"`
	Mode                 uint8  `json:"mode" db:"mode"`
	Logo                 string `json:"logo" db:"logo"`
	Cover                string `json:"cover" db:"cover"`
	Mission              string `json:"mission" db:"mission"`
	TokenContractAddress string `json:"tokenContractAddress" db:"token_contract_address"`
	Overview             string `json:"overview" db:"overview"`
	ChainID              uint64 `json:"chainID" db:"chain_id"`
	TxHash               string `json:"blockChainAddress" db:"tx_hash"`
	OnChain              bool   `json:"onChain" db:"on_chain"`
	KYC                  string `json:"kyc" db:"kyc"`
	ContractAudit        string `json:"contractAudit" db:"contract_audit"`
}

type StartupBasic struct {
	Banner        string `json:"banner"`
	ChainId       int    `json:"chainId"`
	ComerId       int    `json:"comerId"`
	ContractAudit string `json:"contractAudit"`
	Id            uint64 `json:"id"`
	IsConnected   bool   `json:"isConnected"`
	Kyc           string `json:"kyc"`
	Logo          string `json:"logo"`
	Mission       string `json:"mission"`
	Name          string `json:"name"`
	OnChain       bool   `json:"onChain"`
	TxHash        string `json:"txHash"`
	Type          int    `json:"type"`
}

type StartupCardResponse struct {
	Banner        string                `json:"banner"`
	ChainId       int                   `json:"chainId"`
	ComerId       int                   `json:"comerId"`
	ContractAudit string                `json:"contractAudit"`
	Id            int                   `json:"id"`
	IsConnected   bool                  `json:"isConnected"`
	Kyc           string                `json:"kyc"`
	Logo          string                `json:"logo"`
	Mission       string                `json:"mission"`
	Name          string                `json:"name"`
	OnChain       bool                  `json:"onChain"`
	Socials       []SocialBookResponse  `json:"socials"`
	Tags          []TagRelationResponse `json:"tags"`
	TxHash        string                `json:"txHash"`
	Type          int                   `json:"type"`
}

type StartupTeamMember struct {
	ComerID      uint64 `json:"comerID"`
	StartupID    uint64 `json:"startupID"`
	Position     string `json:"position"`
	Comer        string `json:"comer"`
	FollowedByMe bool   `json:"followedByMe"`
}

type Tag struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	IsIndex  bool   `json:"isIndex"`
}

type TagListResponse struct {
	List []TagResponse `json:"list"`
}

type TagRelationResponse struct {
	Id       int         `json:"id"`
	Tag      TagResponse `json:"tag"`
	TagId    int         `json:"tag_id"`
	TargetId int         `json:"target_id"`
	Type     int         `json:"type"`
}

type TagResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type Wallet struct {
	Address string `json:address`
}
