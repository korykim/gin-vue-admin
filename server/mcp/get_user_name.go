package mcpTool

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/komo/model"
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&GetUserName{})
}

type GetUserName struct {
}

// 根据昵称返回用户的登录名称
func (t *GetUserName) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 1. 参数验证
	nickname, ok := request.Params.Arguments["nickname"].(string)
	if !ok {
		return nil, errors.New("参数错误：nickname 必须是字符串类型")
	}

	if nickname == "" {
		return nil, errors.New("参数错误：nickname 不能为空")
	}

	// 2. 记录操作日志
	global.GVA_LOG.Info("getUserName 工具被调用")

	// 3. 优化查询，使用LIKE模糊查询，获取所有匹配结果
	var users []struct {
		Username string
		Nickname string
	}

	err := global.GVA_DB.Model(&model.QmUser{}).
		Select("username, nickname").
		Where("nickname LIKE ?", "%"+nickname+"%").
		Limit(10). // 限制返回数量，防止返回过多结果
		Find(&users).Error

	// 4. 优化错误处理
	if err != nil {
		global.GVA_LOG.Error("数据库查询错误")
		return nil, errors.New("系统错误，请稍后再试")
	}

	// 5. 处理查询结果
	if len(users) == 0 {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("未找到昵称包含\"%s\"的用户", nickname),
				},
			},
		}, nil
	}

	// 多个结果时，格式化输出
	var resultText string
	if len(users) == 1 {
		resultText = fmt.Sprintf("昵称为\"%s\"的用户名是 %s", users[0].Nickname, users[0].Username)
	} else {
		var matchResults []string
		for _, user := range users {
			matchResults = append(matchResults, fmt.Sprintf("- 昵称：%s，用户名：%s", user.Nickname, user.Username))
		}
		resultText = fmt.Sprintf("找到%d个昵称包含\"%s\"的用户：\n\n%s", len(users), nickname, strings.Join(matchResults, "\n"))
	}

	// 构造回复信息
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: resultText,
			},
		},
	}, nil
}

func (t *GetUserName) New() mcp.Tool {
	return mcp.NewTool("getUserName",
		mcp.WithDescription("根据昵称返回用户的登录名称"),
		mcp.WithString("nickname", mcp.Required(),
			mcp.Description("用户传入的昵称"),
		),
	)
}
