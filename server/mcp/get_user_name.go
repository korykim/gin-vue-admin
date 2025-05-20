package mcpTool

import (
	"context"
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/komo/model"
	"github.com/mark3labs/mcp-go/mcp"
	"gorm.io/gorm"
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

	// 3. 优化查询，只选择需要的字段
	var user struct {
		Username string
	}

	err := global.GVA_DB.Model(&model.QmUser{}).
		Select("username").
		Where("nickname = ?", nickname).
		First(&user).Error

	// 4. 优化错误处理
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf("昵称为 %s 的用户不存在", nickname),
					},
				},
			}, nil
		}
		global.GVA_LOG.Error("数据库查询错误")
		return nil, errors.New("系统错误，请稍后再试")
	}

	// 构造回复信息
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf("昵称为 %s 的用户名是 %s", nickname, user.Username),
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
