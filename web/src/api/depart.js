import service from '@/utils/request'

export const createDepart = (data) => {
  return service({
    url: '/depart/createDepart',
    method: 'post',
    data
  })
}

export const deleteDepart = (data) => {
  return service({
    url: '/depart/deleteDepart',
    method: 'delete',
    data
  })
}

export const updateDepart = (data) => {
  return service({
    url: '/depart/updateDepart',
    method: 'put',
    data
  })
}

export const getDepartInfo = (params) => {
  return service({
    url: '/depart/getDepartInfo',
    method: 'get',
    params
  })
}


// @Tags User
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
export const getDepartList = (data) => {
  return service({
    url: '/depart/getDepartList',
    method: 'post',
    data: data
  })
}
