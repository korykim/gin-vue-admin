package demo

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ StudentApi }

var studentService = service.ServiceGroupApp.DemoServiceGroup.StudentService
