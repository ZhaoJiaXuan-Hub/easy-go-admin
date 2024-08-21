import type * as System from './type'
import http from '@/utils/http'

/** @desc 获取菜单数据 */
export function getSystemMenuList() {
  return http.get<System.MenuItem[]>('/system/menu/getMenuTree')
}

/** @desc 获取菜单详情 */
export function getSystemMenuDetail(params: { id: string }) {
  return http.get<System.MenuItem>('/system/menu/detail', params)
}

/** @desc 获取角色分配权限菜单树 */
export function getSystemMenuOptions() {
  return http.get<System.MenuOptionsItem[]>('/system/menu/options')
}

/** @desc 获取动态路由数据 */
export function getSystemAsyncRoutes() {
  return http.get<any[]>('/system/menu/routes')
}

/** @desc 保存/新增 */
export function saveMenu(data: any) {
  return http.post<System.MenuItem>('/system/menu/save', data)
}
