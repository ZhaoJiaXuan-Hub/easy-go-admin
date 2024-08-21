import type * as User from './type'
import http from '@/utils/http'
import type { MenuItem } from '@/apis/system/type'

/** @desc 登录 */
export function login(data: { user_name: string, password: string }) {
  return http.post<User.LoginRes>('/system/account/login', data)
}

/** @desc 退出登录 */
export function logout() {
  return http.post('/user/logout')
}

/** @desc 获取用户信息 */
export const getUserInfo = () => {
  return http.get<User.UserInfo>('/system/account/getDetail')
}

/** @desc 获取用户路由信息 */
export const getUserAsyncRoutes = () => {
  return http.get<MenuItem[]>('/system/account/getRouters')
}
