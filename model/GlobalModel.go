package model

type User struct {
	UserId       int64  `json:"user_id"`
	UserName     string `json:"user_name"`
	RoomId       int64  `json:"room_id"`
	OperationCnt int64  `json:"operation_cnt"`
	DicePos      int64  `json:"dice_pos"` // 骰子位置
}

type SnakesAndLaddersResp struct {
	IsFinish   bool  `json:"is_finish"`   // 是否完成
	DicePoints int64 `json:"dice_points"` // 骰子点数
	GeneralPos int64 `json:"general_pos"` // 普通目标位置
	IsBeJump   bool  `json:"is_be_jump"`  // 是否跳跃位置
	JumpPos    int64 `json:"jump_pos"`    // 跳跃后位置
}
