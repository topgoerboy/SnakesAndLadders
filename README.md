# SnakesAndLadders

## api文档

1、连接

```json
{
    "code": 0,
    "data": {
        "user_id": 10,
        "user_name: "Player-10",
        "room_id": 1,
        "dice_pos": 0
    }
}
```

2、游戏开始
Request

```json
{
    "code": 2,
    "data": ""
}
```

Response

```json
{
    "code": 2,
    "data": {
        "is_finish": "false", // 是否完成
        "dice_points": 2, // 骰子点数
        "general_pos": 2, // 普通位置
        "is_be_jump": false, // 是否跳跃
        "jump_pos": 0 // 跳跃后位置
    }
}
```
