package protocol

type ResultCode int

const (
	Success            ResultCode = 0
	Failed             ResultCode = 1   // 요청 처리 실패
	IpInvalid          ResultCode = 2   // 요청 IP가 유효하지 않음
	UserIDNotFound     ResultCode = 13  // 유저 아이디가 존재하지 않음
	AccessTokenInvalid ResultCode = 101 // 액세스 토큰이 유효하지 않음
)

func (r ResultCode) toString() string {
	switch r {
	case Success:
		return "Success"
	case Failed:
		return "Failed"
	case IpInvalid:
		return "IpInvalid"
	case UserIDNotFound:
		return "UserIDNotFound"
	case AccessTokenInvalid:
		return "AccessTokenInvalid"
	}
	return ""
}
