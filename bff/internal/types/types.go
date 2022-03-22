// Code generated by goctl. DO NOT EDIT.
package types

type PageReq struct {
	PageSize   int32   `form:"pageSize"`
	PageNumber *int32  `form:"pageNumber,optional"`
	LastId     *uint64 `form:"lastId,optional"`
	HasCount   bool    `form:"hasCount"`
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserInfo struct {
	Id           string  `json:"id"`
	Nickname     *string `json:"nickname,optional"`
	Favicon      *string `json:"favicon,optional"`
	Phone        *string `json:"phone,optional"`
	Email        string  `json:"email"`
	Coins        int32   `json:"coins"`
	Point        int32   `json:"point"`
	Introduction *string `json:"introduction,optional"`
	Github       *string `json:"github,optional"`
	Website      *string `json:"website,optional"`
	Wechat       *string `json:"wechat,optional"`
	Role         int32   `json:"role"`
}

type LoginResult struct {
	Info      UserInfo `json:"info"`
	IsNewUser bool     `json:"isNewUser"`
	Tokens    Tokens   `json:"tokens"`
}

type LoginByPasswordReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginByCodeReq struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type UserInfoReq struct {
	Email string `form:"email",json:"email"`
}

type VerifyEmailReq struct {
	Email string `json:"email"`
}

type ChangeUserInfoReq struct {
	Nickname     *string `json:"nickname,optional"`
	Favicon      *string `json:"favicon,optional"`
	Phone        *string `json:"phone,optional"`
	Introduction *string `json:"introduction,optional"`
	Github       *string `json:"github,optional"`
	Website      *string `json:"website,optional"`
	Wechat       *string `json:"wechat,optional"`
}

type ChangePasswordReq struct {
	NewPassword string `json:"newPassword"`
	Code        string `json:"code"`
}

type RefreshReq struct {
	RefreshToken string `json:"refreshToken"`
}

type UserRank struct {
	Rank int32 `json:"rank"`
}

type MessageByIdReq struct {
	Id uint64 `form:"id",json:"id"`
}

type Message struct {
	Id       uint64 `json:"id"`
	Receiver uint64 `json:"receiver"`
	Time     int64  `json:"time"`
	Read     bool   `json:"read"`
	Type     byte   `json:"type"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type MessageaPageResp struct {
	Messages   []*Message `json:"messages"`
	TotalCount *int32     `json:"totalCount"`
}

type ProblemSynopsis struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Point int32  `json:"point"`
	Level int32  `json:"level"`
	Tags  string `json:"tags"`
}

type ProblemDetail struct {
	Id          int32  `json:"id,optional"`
	Name        string `json:"name"`
	Point       int32  `json:"point"`
	Level       int32  `json:"level"`
	Tags        string `json:"tags"`
	Description string `json:"description"`
	TimeLimit   uint64 `json:"timeLimit"`
	SpaceLimit  uint64 `json:"spaceLimit"`
}

type ProblemIO struct {
	InText  string `json:"inText"`
	OutText string `json:"outText"`
}

type ProblemAddOrUpdateDto struct {
	Detail  *ProblemDetail `json:"detail"`
	InText  string         `json:"inText"`
	OutText string         `json:"outText"`
}

type ProblemHotVO struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type ProblemHotResp struct {
	Problems []*ProblemHotVO `json:"problems"`
}

type ProblemsByTagReq struct {
	Tags []int32 `json:"tags,omitempty"`
}

type ProblemPage struct {
	TotalCount int32              `json:"totalCount"`
	Problems   []*ProblemSynopsis `json:"problems"`
}

type Integer struct {
	Value int32 `form:"value"`
}

type AllProblemStatistic struct {
	Easy   int32 `json:"easy"`
	Medium int32 `json:"medium"`
	Hard   int32 `json:"hard"`
}

type SolutionAddReq struct {
	Pid      int32  `json:"pid"`
	Nickname string `json:"nickname"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type SolutionSummary struct {
	Id        uint64 `json:"id"`
	Pid       int32  `json:"pid"`
	Time      uint64 `json:"time"`
	Nickname  string `json:"nickname"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	IsTeacher bool   `json:"isTeacher"`
}

type SolutionDetail struct {
	Id        uint64 `json:"id"`
	Uid       uint64 `json:"uid"`
	Pid       int32  `json:"pid"`
	Time      uint64 `json:"time"`
	Nickname  string `json:"nickname"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
	IsTeacher bool   `json:"isTeacher"`
}

type SolutionByIdReq struct {
	Id uint64 `form:"id"`
}

type AllSolutionByPidReq struct {
	Pid int32 `form:"pid"`
}

type SolutionList struct {
	Solutions []*SolutionSummary `json:"solutions"`
}

type RecordDetail struct {
	Id        uint64 `json:"id"`
	Uid       uint64 `json:"uid"`
	Time      uint64 `json:"time"`
	Pid       int32  `json:"pid"`
	Status    uint32 `json:"status"`
	Language  uint32 `json:"language"`
	TimeUsed  uint64 `json:"timeUsed"`
	SpaceUsed uint64 `json:"spaceUsed"`
	Code      string `json:"code"`
}

type RecordAddReq struct {
	Pid      int32  `json:"pid"`
	Language uint32 `json:"language"`
	Code     string `json:"code"`
}

type RecordList struct {
	Records []*RecordDetail `json:"records"`
}

type RecordStatus struct {
	Status uint32 `json:"status"`
}

type RecordByIdReq struct {
	Id uint64 `form:"id"`
}

type Long struct {
	Id uint64 `json:"id"`
}

type SuccessStatistic struct {
	Easy   int32 `json:"easy"`
	Medium int32 `json:"medium"`
	Hard   int32 `json:"hard"`
}

type AskByIdReq struct {
	Id uint64 `form:"id"`
}

type AskDetail struct {
	Id       uint64 `json:"id"`
	Uid      uint64 `json:"uid"`
	Time     int64  `json:"time"`
	Nickname string `json:"nickname"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type AskAddReq struct {
	Nickname string `json:"nickname"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type AskSummary struct {
	Id       uint64 `json:"id"`
	Uid      uint64 `json:"uid"`
	Time     int64  `json:"time"`
	Nickname string `json:"nickname"`
	Title    string `json:"title"`
}

type ReplyByIdReq struct {
	Id uint64 `form:"id"`
}

type ReplyAddReq struct {
	AskId    uint64 `json:"askId"`
	Nickname string `json:"nickname"`
	Content  string `json:"content"`
}

type ReplyDetail struct {
	Id        uint64 `json:"id"`
	AskId     uint64 `json:"askId"`
	Uid       uint64 `json:"uid"`
	Time      int64  `json:"time"`
	Nickname  string `json:"nickname"`
	Content   string `json:"content"`
	IsTeacher bool   `json:"isTeacher"`
}

type AskList struct {
	Asks []*AskSummary `json:"asks"`
}

type ReplyList struct {
	Replies []*ReplyDetail `json:"replies"`
}
